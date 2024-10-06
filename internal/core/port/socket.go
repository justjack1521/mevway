package port

import (
	"context"
	socket2 "mevway/internal/core/domain/socket"
)

type Client interface {
	Read()
	Write()
	Notify(data []byte)
	Close()
}

type ClientRepository interface {
	Add(ctx context.Context, client socket2.Client) error
	Remove(ctx context.Context, client socket2.Client) error
	List(ctx context.Context) ([]socket2.Client, error)
}

type SocketServer interface {
	Register(client socket2.Client, notifier Client)
	Unregister(client socket2.Client)
	Notify(ctx context.Context, message socket2.Message)
}

type SocketMessageRouter interface {
	Route(ctx context.Context, message socket2.Message) (socket2.Response, error)
}
