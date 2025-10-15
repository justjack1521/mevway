package socket

import uuid "github.com/satori/go.uuid"

type ClosureReason int

const (
	ClosureReasonDefault    = iota
	ClosureReasonReadStop   = 1
	ClosureReasonWriteStop  = 2
	ClosureReasonServerStop = 3
	ClosureReasonInactivity = 4
	ClosureReasonTakeover   = 5
	ClosureReasonRejected   = 5
)

type Client struct {
	Session  uuid.UUID
	UserID   uuid.UUID
	PlayerID uuid.UUID
	PatchID  uuid.UUID
}

func NewClient(session, user, player uuid.UUID) Client {
	return Client{
		Session:  session,
		UserID:   user,
		PlayerID: player,
	}
}
