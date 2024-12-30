package port

import (
	"context"
	"mevway/internal/core/domain/socket"
	"time"
)

type Client interface {
	Read()
	Write()
	Notify(data []byte)
	Close()
	LastMessage() time.Time
}

type ClientConnectionRepository interface {
	Add(ctx context.Context, client socket.Client) error
	Remove(ctx context.Context, client socket.Client) error
	RemoveAll(ctx context.Context) error
	List(ctx context.Context) ([]socket.Client, error)
}

type SocketServer interface {
	Register(client socket.Client, notifier Client) error
	Unregister(client socket.Client)
	Notify(ctx context.Context, message socket.Message)
}

type SocketMessageRouter interface {
	Route(ctx context.Context, message socket.Message) (socket.Response, error)
}
