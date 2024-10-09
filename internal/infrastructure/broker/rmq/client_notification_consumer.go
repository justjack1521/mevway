package rmq

import (
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
	"mevway/internal/core/application"
	"mevway/internal/core/port"
)

type ClientNotificationConsumer struct {
	consumer   *mevrabbit.StandardConsumer
	svc        port.SocketServer
	translator application.NotificationTranslator
}

func NewClientNotificationConsumer(app *ApplicationConnection, svc port.SocketServer, translator application.NotificationTranslator) *ClientNotificationConsumer {
	service := &ClientNotificationConsumer{
		svc:        svc,
		translator: translator,
	}

	consumer, err := mevrabbit.NewStandardConsumer(app.conn, mevrabbit.ClientUpdate, mevrabbit.ClientNotification, mevrabbit.Client, service.consume)

	if err != nil {
		panic(err)
	}

	service.consumer = consumer.WithSlogging(app.slogger).WithTracing(app.tracer)
	return service

}

func (s *ClientNotificationConsumer) consume(ctx *mevrabbit.ConsumerContext) (action rabbitmq.Action, err error) {

	notification, err := s.translator.Notification(ctx.Delivery.Body)
	if err != nil {
		return rabbitmq.NackDiscard, err
	}
	notification.UserID = ctx.UserID()

	s.svc.Notify(ctx, notification)

	return rabbitmq.Ack, nil
}

func (s *ClientNotificationConsumer) Close() {
	s.consumer.Close()
}
