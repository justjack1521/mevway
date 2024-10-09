package rmq

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/user"
)

type UserEventPublisher struct {
	publisher  *mevrabbit.StandardPublisher
	translator application.UserEventTranslator
}

func NewUserEventPublisher(app *ApplicationConnection, publisher *mevent.Publisher, translator application.UserEventTranslator) *UserEventPublisher {
	var pub = mevrabbit.NewUserPublisher(app.conn).WithSlogging(app.slogger).WithTracing(app.tracer)
	var service = &UserEventPublisher{publisher: pub, translator: translator}
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

func (s *UserEventPublisher) Close() {
	s.publisher.Close()
}
