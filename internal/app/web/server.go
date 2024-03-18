package web

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ClientNotification struct {
	ClientID     uuid.UUID
	Notification *protocommon.Notification
}

type Server struct {
	logger              *logrus.Logger
	Register            chan *Client
	Unregister          chan *Client
	Clients             map[*Client]bool
	Broadcast           chan []byte
	NotifyClient        chan *ClientNotification
	Services            map[RoutingKey]ServiceClientRouter
	publisher           *mevent.Publisher
	updateConsumer      *ServerUpdateConsumer
	updatePublisher     *ServerUpdatePublisher
	newRelicApplication *newrelic.Application
}

func NewServer(logger *logrus.Logger, relic *newrelic.Application) *Server {
	svr := &Server{
		logger:              logger,
		newRelicApplication: relic,
		Register:            make(chan *Client),
		Unregister:          make(chan *Client),
		Clients:             make(map[*Client]bool),
		Broadcast:           make(chan []byte),
		NotifyClient:        make(chan *ClientNotification),
		Services:            make(map[RoutingKey]ServiceClientRouter),
	}
	return svr.WithEventBus(mevent.PublisherWithLogger(logger))
}

func (s *Server) RegisterServiceClient(key RoutingKey, router ServiceClientRouter) {
	s.logger.WithFields(logrus.Fields{
		"routing.key": key,
	}).Info("Service Client Router Attached")
	s.Services[key] = router
}

func (s *Server) WithUpdatePublisher(mq *rabbitmq.Conn) *Server {
	publisher, err := NewServerUpdatePublisher(s, mq)
	if err != nil {
		panic(err)
	}
	s.updatePublisher = publisher
	return s
}

func (s *Server) WithUpdateConsumer(mq *rabbitmq.Conn) *Server {
	consumer, err := NewServerUpdateConsumer(s, mq)
	if err != nil {
		panic(err)
	}
	s.updateConsumer = consumer
	return s
}

func (s *Server) WithEventBus(options ...mevent.EventPublisherConfiguration) *Server {
	s.publisher = mevent.NewPublisher(options...)
	s.publisher.Subscribe(s, ServerStartEvent{}, ClientMessageErrorEvent{})
	return s
}

func (s *Server) Notify(event mevent.Event) {
	switch actual := event.(type) {
	case ServerStartEvent:
		break
	case ClientMessageErrorEvent:
		s.logger.WithFields(actual.ToLogFields()).WithError(actual.Error()).Error("Client Message Error")
	}
}

func (s *Server) Run() {
	s.publisher.Notify(ServerStartEvent{})
	for {
		select {
		//Register
		case client := <-s.Register:
			s.publisher.Notify(ClientConnectedEvent{clientID: client.ClientID, remoteAddr: client.connection.RemoteAddr()})
			s.Clients[client] = true
		//Unregister
		case client := <-s.Unregister:
			if _, ok := s.Clients[client]; ok {
				s.publisher.Notify(ClientDisconnectedEvent{clientID: client.ClientID, remoteAddr: client.connection.RemoteAddr()})
				delete(s.Clients, client)
				close(client.send)
			}
		//Broadcast
		case message := <-s.Broadcast:
			for client := range s.Clients {
				select {
				case client.send <- message:
				default:
					delete(s.Clients, client)
					close(client.send)
				}
			}
		//Notify
		case notification := <-s.NotifyClient:
			bytes, err := notification.Notification.MarshallBinary()
			if err != nil {
				s.publisher.Notify(ClientMessageErrorEvent{clientID: notification.ClientID, err: err})
			}
			for client := range s.Clients {
				if client.ClientID == notification.ClientID {
					select {
					case client.send <- bytes:
					default:
						delete(s.Clients, client)
						close(client.send)
					}
					break
				}
			}
		}
	}
}

func (s *Server) RouteClientRequest(ctx context.Context, wc *Client, request *protocommon.BaseRequest) error {

	s.logger.WithFields(logrus.Fields{
		"service":   request.Service,
		"operation": request.Operation,
	}).Info("Executing Request")

	segmentRoute := newrelic.FromContext(ctx).StartSegment("socket/route")
	md := metadata.New(map[string]string{"X-API-CLIENT": wc.ClientID.String()})
	wcc := wc.NewClientContext(metadata.NewOutgoingContext(ctx, md), request)

	fmt.Println(RoutingKey(request.Service))

	service, exists := s.Services[RoutingKey(request.Service)]
	if exists == false {
		return ErrFailedRoutingClientRequest(ErrMalformedRequest)
	}

	result, err := service.Route(wcc, int(request.Operation), request.Data)
	segmentRoute.End()

	segmentResponse := newrelic.FromContext(ctx).StartSegment("socket/response")
	if err != nil {
		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *protocommon.ApplicationError:
				if err = s.SendClientError(wcc, t.ErrorCode, t.ErrorMessage); err != nil {
					return ErrSendingClientResponse(wcc.client.ClientID, err)
				}
				return nil
			}
		}
		return s.SendClientError(wcc, 9, err.Error())
	}

	if err := s.SendClientResponse(wcc, result); err != nil {
		return ErrSendingClientResponse(wcc.client.ClientID, err)
	}

	s.logger.WithFields(logrus.Fields{
		"service":   request.Service,
		"operation": request.Operation,
	}).Info("Request Complete")

	segmentResponse.End()

	return nil

}

func (s *Server) SendClientError(ctx *ClientContext, code int32, message string) error {

	s.logger.WithFields(logrus.Fields{
		"service":   ctx.request.Service,
		"operation": ctx.request.Operation,
	}).Info("request Failed")

	response := ctx.NewError(code, message)

	msg, err := proto.Marshal(response)

	if err != nil {
		return err
	}

	ctx.client.send <- msg

	return nil

}

func (s *Server) SendClientResponse(ctx *ClientContext, res ClientResponse) error {

	bytes, err := res.MarshallBinary()

	if err != nil {
		return err
	}

	response := ctx.NewResponse(bytes)

	message, err := proto.Marshal(response)

	if err != nil {
		return err
	}

	ctx.client.send <- message

	return nil

}
