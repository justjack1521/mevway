package web

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/rabbitmv"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type ServerUpdatePublisher struct {
	logger    *logrus.Logger
	publisher *rabbitmv.StandardPublisher
}

func NewServerUpdatePublisher(server *Server, connection *rabbitmq.Conn) (*ServerUpdatePublisher, error) {
	service := &ServerUpdatePublisher{
		logger: server.logger,
	}
	server.publisher.Subscribe(service, ClientConnectedEvent{}, ClientHeartbeatEvent{}, ClientDisconnectedEvent{})
	service.publisher = rabbitmv.NewClientPublisher(connection)
	return service, nil

}

func (s *ServerUpdatePublisher) publishClientConnected(ctx context.Context, evt ClientConnectedEvent) error {
	request := &protocommon.ClientConnected{RemoteAddress: evt.RemoteAddress().String()}
	bytes, err := request.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.PublishWithContext(ctx, bytes, evt.ClientID(), rabbitmv.ClientConnected); err != nil {
		return err
	}
	return nil
}

func (s *ServerUpdatePublisher) publishClientHeartbeat(ctx context.Context, evt ClientHeartbeatEvent) error {
	request := &protocommon.ClientHeartbeat{RemoteAddress: evt.RemoteAddress().String()}
	bytes, err := request.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.PublishWithContext(ctx, bytes, evt.ClientID(), rabbitmv.ClientHeartbeat); err != nil {
		return err
	}
	return nil
}

func (s *ServerUpdatePublisher) publishClientDisconnected(ctx context.Context, evt ClientDisconnectedEvent) error {
	request := &protocommon.ClientDisconnected{RemoteAddress: evt.RemoteAddress().String()}
	bytes, err := request.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.PublishWithContext(ctx, bytes, evt.ClientID(), rabbitmv.ClientDisconnected); err != nil {
		return err
	}
	return nil
}

func (s *ServerUpdatePublisher) Notify(evt mevent.Event) {
	var err error
	switch actual := evt.(type) {
	case ClientConnectedEvent:
		err = s.publishClientConnected(actual.Context(), actual)
	case ClientDisconnectedEvent:
		err = s.publishClientDisconnected(actual.Context(), actual)
	case ClientHeartbeatEvent:
		err = s.publishClientHeartbeat(actual.Context(), actual)
	}
	if err != nil {
		s.logger.WithFields(evt.ToLogFields()).WithError(err).Error("Server Update Publisher Failed Processing client Event")
	}
}
