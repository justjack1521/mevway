package web

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net"
)

type ServerEvent interface {
	mevent.Event
}

type ServerStartEvent struct {
}

func (e ServerStartEvent) Name() string {
	return "event.server.start"
}

func (e ServerStartEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event_name": e.Name()}
}

type ClientEvent interface {
	mevent.ClientEvent
	mevent.ContextEvent
	UserID() uuid.UUID
	RemoteAddress() net.Addr
}

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
	return logrus.Fields{"event.name": e.Name(), "client.id": e.userID.String(), "remote.address": e.remoteAddr.String()}
}

type ClientHeartbeatEvent struct {
	userID     uuid.UUID
	playerID   uuid.UUID
	remoteAddr net.Addr
	ctx        context.Context
}

func (e ClientHeartbeatEvent) Context() context.Context {
	return e.ctx
}

func (e ClientHeartbeatEvent) UserID() uuid.UUID {
	return e.userID
}

func (e ClientHeartbeatEvent) PlayerID() uuid.UUID {
	return e.playerID
}

func (e ClientHeartbeatEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientHeartbeatEvent) Name() string {
	return "event.client.heartbeat"
}

func (e ClientHeartbeatEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event.name": e.Name(), "client.id": e.userID.String(), "remote.address": e.remoteAddr.String()}
}

type ClientDisconnectedEvent struct {
	ctx        context.Context
	userID     uuid.UUID
	playerID   uuid.UUID
	remoteAddr net.Addr
	source     string
}

func NewClientDisconnectedEvent(ctx context.Context, user, player uuid.UUID, addr net.Addr, source string) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{ctx: ctx, userID: user, playerID: player, remoteAddr: addr, source: source}
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

func (e ClientDisconnectedEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientDisconnectedEvent) Source() string {
	return e.source
}

func (e ClientDisconnectedEvent) Name() string {
	return "event.client.disconnected"
}

func (e ClientDisconnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name":     e.Name(),
		"client.id":      e.userID.String(),
		"remote.address": e.remoteAddr.String(),
		"source":         e.source,
	}
}

type ClientMessageEvent interface {
	ClientEvent
	Error() error
}

type ClientMessageErrorEvent struct {
	userID     uuid.UUID
	remoteAddr net.Addr
	err        error
}

func NewClientMessageErrorEvent(user uuid.UUID, addr net.Addr, err error) ClientMessageErrorEvent {
	return ClientMessageErrorEvent{userID: user, remoteAddr: addr, err: err}
}

func (e ClientMessageErrorEvent) UserID() uuid.UUID {
	return e.userID
}

func (e ClientMessageErrorEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientMessageErrorEvent) Error() error {
	return e.err
}

func (e ClientMessageErrorEvent) Name() string {
	return "event.client.error.received"
}

func (e ClientMessageErrorEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event.name": e.Name(), "client.id": e.userID.String(), "remote.address": e.remoteAddr.String(), "error": e.err.Error()}
}
