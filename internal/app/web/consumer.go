package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/mevent"
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
)

type ServerUpdateConsumer struct {
	closed   bool
	server   *Server
	consumer *mevrabbit.StandardConsumer
}

func NewServerUpdateConsumer(server *Server, conn *rabbitmq.Conn) (*ServerUpdateConsumer, error) {
	service := &ServerUpdateConsumer{
		server: server,
	}
	consumer, err := mevrabbit.NewStandardConsumer(conn, mevrabbit.ClientUpdate, mevrabbit.ClientNotification, mevrabbit.Client, service.consume)
	if err != nil {
		panic(err)
	}
	service.consumer = consumer
	return service, nil
}

func (s *ServerUpdateConsumer) consume(ctx *mevrabbit.ConsumerContext) (action rabbitmq.Action, err error) {
	notification, err := protocommon.NewNotification(ctx.Delivery.Body)
	if err != nil {
		return rabbitmq.NackDiscard, err
	}
	s.server.NotifyClient <- &ClientNotification{ClientID: ctx.UserID(), Notification: notification}
	return rabbitmq.Ack, nil
}

func (s *ServerUpdateConsumer) Notify(evt mevent.Event) {
	switch evt.(type) {
	case mevent.ApplicationShutdownEvent:
		s.consumer.Close()
	}
}
