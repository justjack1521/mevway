package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/socket"
	"time"
)

type Client interface {
	Read()
	Write()
	Notify(data []byte)
	Close(reason socket.ClosureReason)
	Terminate(reason socket.ClosureReason)
	ClosureReason() socket.ClosureReason
	LastMessage() time.Time
}

type ClientConnectionRepository interface {
	Add(ctx context.Context, client socket.Client) error
	Remove(ctx context.Context, client socket.Client) error
	RemoveAll(ctx context.Context) error
	List(ctx context.Context) ([]socket.Client, error)
}

type ClientRequestRepository interface {
	Create(ctx context.Context, player uuid.UUID, data []byte) error
}

type SocketServer interface {
	Register(client socket.Client, notifier Client) error
	Unregister(client socket.Client)
	Notify(ctx context.Context, message socket.Message)
}

type SocketMessageRouter interface {
	Route(ctx context.Context, message socket.Message) (socket.Response, error)
}
