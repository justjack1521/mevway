package socket

import uuid "github.com/satori/go.uuid"

type Client struct {
	Session  uuid.UUID
	UserID   uuid.UUID
	PlayerID uuid.UUID
}

func NewClient(session uuid.UUID, user uuid.UUID, player uuid.UUID) Client {
	return Client{Session: session, UserID: user, PlayerID: player}
}
