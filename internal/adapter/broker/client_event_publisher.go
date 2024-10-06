package broker

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/core/application"
	"mevway/internal/domain/socket"
)

type SocketClientEventPublisher struct {
	publisher  *mevrabbit.StandardPublisher
	translator application.SocketEventTranslator
}

func NewClientEventPublisher(connection *rabbitmq.Conn, publisher *mevent.Publisher, translator application.SocketEventTranslator) *SocketClientEventPublisher {
	var service = &SocketClientEventPublisher{publisher: mevrabbit.NewClientPublisher(connection), translator: translator}
	publisher.Subscribe(service, socket.ClientConnectedEvent{}, socket.ClientDisconnectedEvent{})
	return service
}

func (s *SocketClientEventPublisher) Notify(evt mevent.Event) {
	switch actual := evt.(type) {
	case socket.ClientConnectedEvent:
		s.publishClientConnected(actual.Context(), actual)
	case socket.ClientDisconnectedEvent:
		s.publishClientDisconnected(actual.Context(), actual)
	}
}

func (s *SocketClientEventPublisher) publishClientConnected(ctx context.Context, evt socket.ClientConnectedEvent) {
	bytes, err := s.translator.Connected(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientConnected); err != nil {
		return
	}
}

func (s *SocketClientEventPublisher) publishClientDisconnected(ctx context.Context, evt socket.ClientDisconnectedEvent) {
	bytes, err := s.translator.Disconnected(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientDisconnected); err != nil {
		return
	}
}

func (s *SocketClientEventPublisher) Close() error {
	s.publisher.Close()
	return nil
}
