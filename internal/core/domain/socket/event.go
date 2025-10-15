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
	patchID   uuid.UUID
}

func NewClientConnectedEvent(ctx context.Context, session, user, player, patch uuid.UUID) ClientConnectedEvent {
	return ClientConnectedEvent{ctx: ctx, sessionID: session, userID: user, playerID: player, patchID: patch}
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

func (e ClientConnectedEvent) PatchID() uuid.UUID {
	return e.patchID
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
		slog.Int("reap.count", e.count),
	}
}

type ConnectionTerminateEvent struct {
	ctx  context.Context
	user uuid.UUID
}

func NewConnectionTerminateEvent(ctx context.Context, user uuid.UUID) ConnectionTerminateEvent {
	return ConnectionTerminateEvent{
		ctx:  ctx,
		user: user,
	}
}

func (e ConnectionTerminateEvent) Name() string {
	return "event.connection.terminate"
}

func (e ConnectionTerminateEvent) ToSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.String("user.id", e.user.String()),
	}
}

func (e ConnectionTerminateEvent) Context() context.Context {
	return e.ctx
}

func (e ConnectionTerminateEvent) UserID() uuid.UUID {
	return e.user
}

type SuspiciousConnectionEvent struct {
	ctx      context.Context
	user     uuid.UUID
	existing uuid.UUID
	attempt  uuid.UUID
}

func NewSuspiciousConnectionEvent(ctx context.Context, user, existing, attempt uuid.UUID) SuspiciousConnectionEvent {
	return SuspiciousConnectionEvent{ctx: ctx, user: user, existing: existing, attempt: attempt}
}

func (e SuspiciousConnectionEvent) Name() string {
	return "event.connection.suspicious"
}

func (e SuspiciousConnectionEvent) ToSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.String("user.id", e.user.String()),
		slog.String("existing.id", e.existing.String()),
		slog.String("attempted.id", e.attempt.String()),
	}
}

func (e SuspiciousConnectionEvent) Context() context.Context {
	return e.ctx
}

func (e SuspiciousConnectionEvent) UserID() uuid.UUID {
	return e.user
}

func (e SuspiciousConnectionEvent) ExistingSessionID() uuid.UUID {
	return e.existing
}

func (e SuspiciousConnectionEvent) NewSessionID() uuid.UUID {
	return e.attempt
}
