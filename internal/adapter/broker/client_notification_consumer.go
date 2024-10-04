package broker

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

func NewClientNotificationConsumer(conn *rabbitmq.Conn, svc port.SocketServer, translator application.NotificationTranslator) *ClientNotificationConsumer {
	service := &ClientNotificationConsumer{
		svc:        svc,
		translator: translator,
	}
	consumer, err := mevrabbit.NewStandardConsumer(conn, mevrabbit.ClientUpdate, mevrabbit.ClientNotification, mevrabbit.Client, service.consume)
	if err != nil {
		panic(err)
	}
	service.consumer = consumer
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

func (s *ClientNotificationConsumer) Close() error {
	s.consumer.Close()
	return nil
}
