package web

import (
	"context"
	"fmt"
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
	ClientID() uuid.UUID
	RemoteAddress() net.Addr
}

type ClientConnectedEvent struct {
	clientID   uuid.UUID
	remoteAddr net.Addr
	ctx        context.Context
}

func (e ClientConnectedEvent) Context() context.Context {
	return e.ctx
}

func (e ClientConnectedEvent) ClientID() uuid.UUID {
	return e.clientID
}

func (e ClientConnectedEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientConnectedEvent) Name() string {
	return "event.client.connected"
}

func (e ClientConnectedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event.name": e.Name(), "client.id": e.clientID.String(), "remote.address": e.remoteAddr.String()}
}

type ClientHeartbeatEvent struct {
	clientID   uuid.UUID
	remoteAddr net.Addr
	ctx        context.Context
}

func (e ClientHeartbeatEvent) Context() context.Context {
	return e.ctx
}

func (e ClientHeartbeatEvent) ClientID() uuid.UUID {
	return e.clientID
}

func (e ClientHeartbeatEvent) RemoteAddress() net.Addr {
	return e.remoteAddr
}

func (e ClientHeartbeatEvent) Name() string {
	return "event.client.heartbeat"
}

func (e ClientHeartbeatEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event.name": e.Name(), "client.id": e.clientID.String(), "remote.address": e.remoteAddr.String()}
}

type ClientDisconnectedEvent struct {
	ctx        context.Context
	clientID   uuid.UUID
	remoteAddr net.Addr
	source     string
}

func NewClientDisconnectedEvent(ctx context.Context, client uuid.UUID, addr net.Addr, source string) ClientDisconnectedEvent {
	return ClientDisconnectedEvent{ctx: ctx, clientID: client, remoteAddr: addr, source: source}
}

func (e ClientDisconnectedEvent) Context() context.Context {
	return e.ctx
}

func (e ClientDisconnectedEvent) ClientID() uuid.UUID {
	return e.clientID
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
		"client.id":      e.clientID.String(),
		"remote.address": e.remoteAddr.String(),
		"source":         e.source,
	}
}

type ClientMessageEvent interface {
	ClientEvent
	Error() error
}

type ClientMessageErrorEvent struct {
	clientID   uuid.UUID
	remoteAddr net.Addr
	err        error
}

func NewClientMessageErrorEvent(client uuid.UUID, addr net.Addr, err error) *ClientMessageErrorEvent {
	fmt.Println(err)
	return &ClientMessageErrorEvent{clientID: client, remoteAddr: addr, err: err}
}

func (e ClientMessageErrorEvent) ClientID() uuid.UUID {
	return e.clientID
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
	return logrus.Fields{"event.name": e.Name(), "client.id": e.clientID.String(), "remote.address": e.remoteAddr.String(), "error": e.err.Error()}
}
