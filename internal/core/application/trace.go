package application

import "context"

type TransactionTracer interface {
	Start(ctx context.Context, name string) (context.Context, Transaction)
}

type Transaction interface {
	Segment
	AddAttribute(key string, val any)
}

type Segment interface {
	NoticeError(err error)
	End()
}
