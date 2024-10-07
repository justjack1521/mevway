package subscriber

import (
	"fmt"
	"github.com/justjack1521/mevium/pkg/mevent"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type ClientPersistenceSubscriber struct {
	repository port.ClientConnectionRepository
}

func NewClientPersistenceConsumer(publisher *mevent.Publisher, repository port.ClientConnectionRepository) *ClientPersistenceSubscriber {
	var consumer = &ClientPersistenceSubscriber{repository: repository}
	publisher.Subscribe(consumer, socket.ClientConnectedEvent{}, socket.ClientDisconnectedEvent{}, application.ShutdownEvent{})
	return consumer
}

func (c *ClientPersistenceSubscriber) Notify(event mevent.Event) {
	switch actual := event.(type) {
	case socket.ClientConnectedEvent:
		c.handleClientConnect(actual)
	case socket.ClientDisconnectedEvent:
		c.handleClientDisconnect(actual)
	case application.ShutdownEvent:
		c.handleApplicationShutdown(actual)
	}
}

func (c *ClientPersistenceSubscriber) handleApplicationShutdown(evt application.ShutdownEvent) {
	if err := c.repository.RemoveAll(evt.Context()); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceSubscriber) handleClientConnect(evt socket.ClientConnectedEvent) {
	if err := c.repository.Add(evt.Context(), socket.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceSubscriber) handleClientDisconnect(evt socket.ClientDisconnectedEvent) {
	if err := c.repository.Remove(evt.Context(), socket.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}
