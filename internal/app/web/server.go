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
	logger          *logrus.Logger
	Register        chan *Client
	Unregister      chan *Client
	Clients         map[*Client]bool
	Broadcast       chan []byte
	NotifyClient    chan *ClientNotification
	Services        map[RoutingKey]ServiceClientRouter
	publisher       *mevent.Publisher
	updateConsumer  *ServerUpdateConsumer
	updatePublisher *ServerUpdatePublisher
	relic           *newrelic.Application
}

func NewServer(logger *logrus.Logger, relic *newrelic.Application) *Server {
	svr := &Server{
		logger:       logger,
		relic:        relic,
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		Broadcast:    make(chan []byte),
		NotifyClient: make(chan *ClientNotification),
		Services:     make(map[RoutingKey]ServiceClientRouter),
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
			s.publisher.Notify(ClientConnectedEvent{clientID: client.UserID, remoteAddr: client.connection.RemoteAddr()})
			s.Clients[client] = true
		//Unregister
		case client := <-s.Unregister:
			if _, ok := s.Clients[client]; ok {
				s.publisher.Notify(ClientDisconnectedEvent{clientID: client.UserID, remoteAddr: client.connection.RemoteAddr()})
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
				if client.UserID == notification.ClientID {
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

func (s *Server) RouteClientRequest(ctx context.Context, wc *Client, request *protocommon.BaseRequest) (err error) {

	fmt.Println(wc.UserID.String())
	fmt.Println(wc.PlayerID.String())

	wcc := wc.NewClientContext(metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"X-API-USER":   wc.UserID.String(),
		"X-API-PLAYER": wc.PlayerID.String(),
	})), request)

	service, exists := s.Services[RoutingKey(request.Service)]
	if exists == false {
		return fmt.Errorf("service not found at key: %d", request.Service)
	}

	response, err := service.Route(wcc, int(request.Operation), request.Data)
	if err != nil {
		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *protocommon.ApplicationError:
				if err = s.SendClientError(wcc, t.ErrorCode, t.ErrorMessage); err != nil {
					return ErrSendingClientResponse(wcc.client.UserID, err)
				}
				return nil
			}
		}
		if err := s.SendClientError(wcc, 9, err.Error()); err != nil {
			return ErrSendingClientResponse(wcc.client.UserID, err)
		}
		return nil
	}

	if err := s.SendClientResponse(wcc, response); err != nil {
		return ErrSendingClientResponse(wcc.client.UserID, err)
	}

	return nil

}

func (s *Server) SendClientError(ctx *ClientContext, code int32, message string) error {
	response := ctx.NewError(code, message)
	msg, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	ctx.client.send <- msg
	return nil
}

func (s *Server) SendClientResponseBytes(ctx *ClientContext, data []byte) error {
	response := ctx.NewResponse(data)
	message, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	ctx.client.send <- message
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
