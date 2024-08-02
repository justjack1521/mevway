package web

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
)

type ServerUpdatePublisher struct {
	publisher *mevrabbit.StandardPublisher
}

func NewServerUpdatePublisher(server *Server, connection *rabbitmq.Conn) (*ServerUpdatePublisher, error) {
	service := &ServerUpdatePublisher{
		publisher: mevrabbit.NewClientPublisher(connection),
	}
	server.publisher.Subscribe(service, ClientConnectedEvent{}, ClientHeartbeatEvent{}, ClientDisconnectedEvent{})
	return service, nil

}

func (s *ServerUpdatePublisher) Notify(evt mevent.Event) {
	switch actual := evt.(type) {
	case ClientConnectedEvent:
		_ = s.publishClientConnected(actual.Context(), actual)
	case ClientDisconnectedEvent:
		_ = s.publishClientDisconnected(actual.Context(), actual)
	case ClientHeartbeatEvent:
		_ = s.publishClientHeartbeat(actual.Context(), actual)
	}
}

func (s *ServerUpdatePublisher) publishClientConnected(ctx context.Context, evt ClientConnectedEvent) error {
	var message = &protocommon.ClientConnected{RemoteAddress: evt.RemoteAddress().String()}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientConnected); err != nil {
		return err
	}
	return nil
}

func (s *ServerUpdatePublisher) publishClientHeartbeat(ctx context.Context, evt ClientHeartbeatEvent) error {
	var message = &protocommon.ClientHeartbeat{
		RemoteAddress: evt.RemoteAddress().String(),
	}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientHeartbeat); err != nil {
		return err
	}
	return nil
}

func (s *ServerUpdatePublisher) publishClientDisconnected(ctx context.Context, evt ClientDisconnectedEvent) error {
	var message = &protocommon.ClientDisconnected{
		SessionId:     evt.SessionID().String(),
		RemoteAddress: evt.RemoteAddress().String(),
	}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return err
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientDisconnected); err != nil {
		return err
	}
	return nil
}
