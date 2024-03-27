package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevium/pkg/rabbitmv"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type ServerUpdateConsumer struct {
	closed   bool
	server   *Server
	consumer *rabbitmv.StandardConsumer
}

func NewServerUpdateConsumer(server *Server, connection *rabbitmq.Conn) (*ServerUpdateConsumer, error) {

	service := &ServerUpdateConsumer{
		server: server,
	}

	consumer := rabbitmv.NewStandardConsumer(
		connection,
		rabbitmv.ClientUpdate,
		rabbitmv.ClientNotification,
		rabbitmv.Client,
		rabbitmv.ConsumeLoggerMiddleWare(server.logger, service.consume),
	).WithNewRelic(server.relic)

	service.consumer = consumer

	return service, nil

}

func (s *ServerUpdateConsumer) consume(ctx *rabbitmv.ConsumerContext) (action rabbitmq.Action, err error) {
	notification, err := protocommon.NewNotification(ctx.Delivery.Body)
	if err != nil {
		return rabbitmq.NackDiscard, err
	}
	s.server.logger.WithFields(logrus.Fields{
		"client.id":            ctx.ClientID.String(),
		"notification.service": notification.Service,
		"notification.type":    notification.Type,
		"notification.length":  len(notification.Data),
	}).Info("client Notification Received")
	s.server.NotifyClient <- &ClientNotification{ClientID: ctx.ClientID, Notification: notification}
	return rabbitmq.Ack, nil
}

func (s *ServerUpdateConsumer) Notify(evt mevent.Event) {
	switch evt.(type) {
	case mevent.ApplicationShutdownEvent:
		s.consumer.Close()
	}
}
