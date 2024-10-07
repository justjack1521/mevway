package application

import (
	"context"
	"github.com/gorilla/websocket"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
)

type SocketClientFactory interface {
	Create(ctx context.Context, client socket.Client, connection *websocket.Conn) (port.Client, error)
}

type MessageTranslator interface {
	Translate(client socket.Client, message []byte) (socket.Message, error)
}

type ResponseTranslator interface {
	Response(message socket.Message, response []byte) (socket.Response, error)
	Error(message socket.Message, err error) (socket.Response, error)
}

type NotificationTranslator interface {
	Notification(data []byte) (socket.Message, error)
}

type SocketEventTranslator interface {
	Connected(event socket.ClientConnectedEvent) ([]byte, error)
	Disconnected(event socket.ClientDisconnectedEvent) ([]byte, error)
}
