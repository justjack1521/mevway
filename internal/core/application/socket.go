package application

import (
	"context"
	"github.com/gorilla/websocket"
	socket2 "mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type SocketClientFactory interface {
	Create(ctx context.Context, client socket2.Client, connection *websocket.Conn) (port.Client, error)
}

type MessageTranslator interface {
	Translate(client socket2.Client, message []byte) (socket2.Message, error)
}

type ResponseTranslator interface {
	Response(message socket2.Message, response []byte) (socket2.Response, error)
	Error(message socket2.Message, err error) (socket2.Response, error)
}

type NotificationTranslator interface {
	Notification(data []byte) (socket2.Message, error)
}

type SocketEventTranslator interface {
	Connected(event socket2.ClientConnectedEvent) ([]byte, error)
	Disconnected(event socket2.ClientDisconnectedEvent) ([]byte, error)
}
