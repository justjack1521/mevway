package socket

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ClientConnectedEvent struct {
	ctx       context.Context
	sessionID uuid.UUID
	userID    uuid.UUID
	playerID  uuid.UUID
}

func NewClientConnectedEvent(ctx context.Context, session, user, player uuid.UUID) ClientConnectedEvent {
	return ClientConnectedEvent{ctx: ctx, sessionID: session, userID: user, playerID: player}
}

func (e ClientConnectedEvent) Context() context.Context {
	return e.ctx
}

func (e ClientConnectedEvent) SessionID() uuid.UUID {
	return e.sessionID
}

func (e ClientConnectedEvent) UserID() uuid.UUID {
	return e.userID
}

func (e ClientConnectedEvent) PlayerID() uuid.UUID {
	return e.playerID
}

func (e ClientConnectedEvent) Name() string {
	return "event.client.connected"
}

func (e ClientConnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name": e.Name(),
		"user.id":    e.userID.String(),
		"player.id":  e.playerID.String(),
	}
}

type ClientDisconnectedEvent struct {
	ctx       context.Context
	userID    uuid.UUID
	playerID  uuid.UUID
	sessionID uuid.UUID
}

func NewClientDisconnectedEvent(ctx context.Context, session, user, player uuid.UUID) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{ctx: ctx, userID: user, playerID: player, sessionID: session}
}

func (e ClientDisconnectedEvent) Context() context.Context {
	return e.ctx
}

func (e ClientDisconnectedEvent) UserID() uuid.UUID {
	return e.userID
}

func (e ClientDisconnectedEvent) PlayerID() uuid.UUID {
	return e.playerID
}

func (e ClientDisconnectedEvent) SessionID() uuid.UUID {
	return e.sessionID
}

func (e ClientDisconnectedEvent) Name() string {
	return "event.client.disconnected"
}

func (e ClientDisconnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name": e.Name(),
		"user.id":    e.userID.String(),
		"player.id":  e.playerID.String(),
	}
}
