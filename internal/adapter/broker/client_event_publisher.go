package broker

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/core/application"
	"mevway/internal/domain/socket"
)

type ClientEventPublisher struct {
	publisher  *mevrabbit.StandardPublisher
	translator application.EventTranslator
}

func NewClientEventPublisher(connection *rabbitmq.Conn, publisher *mevent.Publisher) *ClientEventPublisher {
	var service = &ClientEventPublisher{publisher: mevrabbit.NewClientPublisher(connection)}
	publisher.Subscribe(service, socket.ClientConnectedEvent{}, socket.ClientDisconnectedEvent{})
	return service
}

func (s *ClientEventPublisher) Notify(evt mevent.Event) {
	switch actual := evt.(type) {
	case socket.ClientConnectedEvent:
		s.publishClientConnected(actual.Context(), actual)
	case socket.ClientDisconnectedEvent:
		s.publishClientDisconnected(actual.Context(), actual)
	}
}

func (s *ClientEventPublisher) publishClientConnected(ctx context.Context, evt socket.ClientConnectedEvent) {
	bytes, err := s.translator.Connected(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientConnected); err != nil {
		return
	}
}

func (s *ClientEventPublisher) publishClientDisconnected(ctx context.Context, evt socket.ClientDisconnectedEvent) {
	bytes, err := s.translator.Disconnected(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.ClientDisconnected); err != nil {
		return
	}
}

func (s *ClientEventPublisher) Close() error {
	s.publisher.Close()
	return nil
}
