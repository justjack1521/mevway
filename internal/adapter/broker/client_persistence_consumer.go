package broker

import (
	"fmt"
	"github.com/justjack1521/mevium/pkg/mevent"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type ClientPersistenceConsumer struct {
	repository port.ClientConnectionRepository
}

func NewClientPersistenceConsumer(publisher *mevent.Publisher, repository port.ClientConnectionRepository) *ClientPersistenceConsumer {
	var consumer = &ClientPersistenceConsumer{repository: repository}
	publisher.Subscribe(consumer, socket.ClientConnectedEvent{})
	return consumer
}

func (c *ClientPersistenceConsumer) Notify(event mevent.Event) {
	switch actual := event.(type) {
	case socket.ClientConnectedEvent:
		c.handleClientConnect(actual)
	case socket.ClientDisconnectedEvent:
		c.handleClientDisconnect(actual)
	}
}

func (c *ClientPersistenceConsumer) handleClientConnect(evt socket.ClientConnectedEvent) {
	if err := c.repository.Add(evt.Context(), socket.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceConsumer) handleClientDisconnect(evt socket.ClientDisconnectedEvent) {
	if err := c.repository.Remove(evt.Context(), socket.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}
