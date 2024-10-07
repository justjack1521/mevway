package web

import (
	"context"
	"github.com/gorilla/websocket"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type ClientFactory struct {
	server       port.SocketServer
	router       port.SocketMessageRouter
	instrumenter application.TransactionInstrumenter
	translator   application.MessageTranslator
}

func NewClientFactory(server port.SocketServer, router port.SocketMessageRouter, instrumenter application.TransactionInstrumenter, translator application.MessageTranslator) *ClientFactory {
	return &ClientFactory{router: router, instrumenter: instrumenter, translator: translator}
}

func (f *ClientFactory) Create(ctx context.Context, client socket.Client, connection *websocket.Conn) (port.Client, error) {
	return NewClient(client, connection, f.server, f.router, f.instrumenter, f.translator), nil
}
