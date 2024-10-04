package port

import (
	"context"
	"mevway/internal/domain/socket"
)

type Client interface {
	Read()
	Write()
	Notify(data []byte)
	Close()
}

type SocketServer interface {
	Register(client socket.Client, notifier Client)
	Unregister(client socket.Client)
	Notify(ctx context.Context, message socket.Message)
}

type SocketMessageRouter interface {
	Route(ctx context.Context, message socket.Message) (socket.Response, error)
}
