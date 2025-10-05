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

func NewClientPersistenceSubscriber(publisher *mevent.Publisher, repository port.ClientConnectionRepository) *ClientPersistenceSubscriber {
	var consumer = &ClientPersistenceSubscriber{repository: repository}
	publisher.Subscribe(consumer, socket.ClientConnectedEvent{}, socket.ClientDisconnectedEvent{}, application.StartEvent{})
	return consumer
}

func (c *ClientPersistenceSubscriber) Notify(event mevent.Event) {
	switch actual := event.(type) {
	case socket.ClientConnectedEvent:
		c.handleClientConnect(actual)
	case socket.ClientDisconnectedEvent:
		c.handleClientDisconnect(actual)
	case application.StartEvent:
		c.handleApplicationRestart(actual)
	}
}

func (c *ClientPersistenceSubscriber) handleApplicationRestart(evt application.StartEvent) {
	if err := c.repository.RemoveAll(evt.Context()); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceSubscriber) handleClientConnect(evt socket.ClientConnectedEvent) {

	var client = socket.NewClient(evt.SessionID(), evt.UserID(), evt.PlayerID())
	client.PatchID = evt.PatchID()

	if err := c.repository.Add(evt.Context(), client); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceSubscriber) handleClientDisconnect(evt socket.ClientDisconnectedEvent) {
	var client = socket.NewClient(evt.SessionID(), evt.UserID(), evt.PlayerID())
	if err := c.repository.Remove(evt.Context(), client); err != nil {
		fmt.Println(err)
	}

}
