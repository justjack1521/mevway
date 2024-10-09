package application

import "context"

type TransactionTracer interface {
	Start(ctx context.Context, name string) (context.Context, Transaction)
}

type Transaction interface {
	AddAttribute(key string, val any)
	NoticeError(err error)
	End()
}
