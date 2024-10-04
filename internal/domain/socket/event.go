package socket

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net"
)

type ClientConnectedEvent struct {
	ctx        context.Context
	userID     uuid.UUID
	playerID   uuid.UUID
	remoteAddr net.Addr
}

func NewClientConnectedEvent(ctx context.Context, user uuid.UUID, player uuid.UUID, addr net.Addr) ClientConnectedEvent {
	return ClientConnectedEvent{ctx: ctx, userID: user, playerID: player, remoteAddr: addr}
}

func (e ClientConnectedEvent) Context() context.Context {
	return e.ctx
}

func (e ClientConnectedEvent) UserID() uuid.UUID {
	return e.userID
}

func (e ClientConnectedEvent) PlayerID() uuid.UUID {
	return e.playerID
}

func (e ClientConnectedEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientConnectedEvent) Name() string {
	return "event.client.connected"
}

func (e ClientConnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name":     e.Name(),
		"client.id":      e.userID.String(),
		"remote.address": e.remoteAddr.String(),
	}
}

type ClientDisconnectedEvent struct {
	ctx        context.Context
	userID     uuid.UUID
	playerID   uuid.UUID
	sessionID  uuid.UUID
	remoteAddr net.Addr
}

func NewClientDisconnectedEvent(ctx context.Context, user, player, session uuid.UUID, addr net.Addr) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{ctx: ctx, userID: user, playerID: player, sessionID: session, remoteAddr: addr}
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

func (e ClientDisconnectedEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientDisconnectedEvent) Name() string {
	return "event.client.disconnected"
}

func (e ClientDisconnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name":     e.Name(),
		"client.id":      e.userID.String(),
		"remote.address": e.remoteAddr.String(),
	}
}
