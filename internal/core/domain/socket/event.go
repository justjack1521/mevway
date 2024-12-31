package socket

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"log/slog"
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

func (e ClientConnectedEvent) ToSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.String("session.id", e.sessionID.String()),
		slog.String("user.id", e.userID.String()),
		slog.String("player.id", e.playerID.String()),
	}
}

type ClientDisconnectedEvent struct {
	ctx       context.Context
	sessionID uuid.UUID
	userID    uuid.UUID
	playerID  uuid.UUID
	reason    ClosureReason
}

func NewClientDisconnectedEvent(ctx context.Context, session, user, player uuid.UUID, reason ClosureReason) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{ctx: ctx, sessionID: session, userID: user, playerID: player, reason: reason}
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

func (e ClientDisconnectedEvent) ToSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.String("session.id", e.sessionID.String()),
		slog.String("user.id", e.userID.String()),
		slog.String("player.id", e.playerID.String()),
		slog.Int("closure.reason", int(e.reason)),
	}
}

type ServerReapEvent struct {
	ctx   context.Context
	count int
}

func NewServerReapEvent(count int) ServerReapEvent {
	return ServerReapEvent{count: count}
}

func (e ServerReapEvent) Name() string {
	return "event.server.reap"
}

func (e ServerReapEvent) ToSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.Int("count", e.count),
	}
}
