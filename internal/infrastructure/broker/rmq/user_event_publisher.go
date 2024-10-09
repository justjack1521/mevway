package rmq

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	uuid "github.com/satori/go.uuid"
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
	publisher.Subscribe(service, user.CreatedEvent{}, user.DeleteEvent{})
	return service
}

func (s *UserEventPublisher) Notify(evt mevent.Event) {
	switch actual := evt.(type) {
	case user.CreatedEvent:
		s.create(actual.Context(), actual)
	case user.DeleteEvent:
		s.delete(actual.Context(), actual)
	}
}

func (s *UserEventPublisher) create(ctx context.Context, evt user.CreatedEvent) {
	bytes, err := s.translator.Created(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), evt.PlayerID(), mevrabbit.UserCreated); err != nil {
		return
	}
}

func (s *UserEventPublisher) delete(ctx context.Context, evt user.DeleteEvent) {
	bytes, err := s.translator.Deleted(evt)
	if err != nil {
		return
	}
	if err := s.publisher.Publish(ctx, bytes, evt.UserID(), uuid.Nil, mevrabbit.UserDeleted); err != nil {
		return
	}
}

func (s *UserEventPublisher) Close() {
	s.publisher.Close()
}
