package application

import "context"

type TransactionInstrumenter interface {
	Start(ctx context.Context, name string) (context.Context, Transaction)
}

type Transaction interface {
	NoticeError(err error)
	End()
}
