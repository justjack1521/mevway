package broker

import (
	"fmt"
	"github.com/justjack1521/mevium/pkg/mevent"
	socket2 "mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type ClientPersistenceConsumer struct {
	repository port.ClientRepository
}

func NewClientPersistenceConsumer(publisher *mevent.Publisher, repository port.ClientRepository) *ClientPersistenceConsumer {
	var consumer = &ClientPersistenceConsumer{repository: repository}
	publisher.Subscribe(consumer, socket2.ClientConnectedEvent{})
	return consumer
}

func (c *ClientPersistenceConsumer) Notify(event mevent.Event) {
	switch actual := event.(type) {
	case socket2.ClientConnectedEvent:
		c.handleClientConnect(actual)
	case socket2.ClientDisconnectedEvent:
		c.handleClientDisconnect(actual)
	}
}

func (c *ClientPersistenceConsumer) handleClientConnect(evt socket2.ClientConnectedEvent) {
	if err := c.repository.Add(evt.Context(), socket2.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}

func (c *ClientPersistenceConsumer) handleClientDisconnect(evt socket2.ClientDisconnectedEvent) {
	if err := c.repository.Remove(evt.Context(), socket2.Client{
		Session:  evt.SessionID(),
		UserID:   evt.UserID(),
		PlayerID: evt.PlayerID(),
	}); err != nil {
		fmt.Println(err)
	}
}
