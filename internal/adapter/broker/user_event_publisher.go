package broker

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/core/application"
	"mevway/internal/domain/user"
)

type UserEventPublisher struct {
	publisher  *mevrabbit.StandardPublisher
	translator application.UserEventTranslator
}

func NewUserEventPublisher(connection *rabbitmq.Conn, publisher *mevent.Publisher, translator application.UserEventTranslator) *UserEventPublisher {
	var service = &UserEventPublisher{publisher: mevrabbit.NewUserPublisher(connection), translator: translator}
	publisher.Subscribe(service, user.CreatedEvent{})
	return service
}

func (s *UserEventPublisher) Notify(evt mevent.Event) {
	switch actual := evt.(type) {
	case user.CreatedEvent:
		s.publishUserCreated(actual.Context(), actual)
	}
}

func (s *UserEventPublisher) publishUserCreated(ctx context.Context, evt user.CreatedEvent) {
	bytes, err := s.translator.Created(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.UserCreated); err != nil {
		return
	}
}

func (s *UserEventPublisher) Close() error {
	s.publisher.Close()
	return nil
}
