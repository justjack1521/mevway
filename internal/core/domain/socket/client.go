package socket

import uuid "github.com/satori/go.uuid"

type ClosureReason int

const (
	ClosureReasonDefault    = iota
	ClosureReasonReadStop   = 1
	ClosureReasonWriteStop  = 2
	ClosureReasonServerStop = 3
)

type Client struct {
	Session  uuid.UUID
	UserID   uuid.UUID
	PlayerID uuid.UUID
}

func NewClient(session uuid.UUID, user uuid.UUID, player uuid.UUID) Client {
	return Client{Session: session, UserID: user, PlayerID: player}
}
