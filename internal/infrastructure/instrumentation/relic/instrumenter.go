package relic

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"mevway/internal/core/application"
)

type Instrumenter struct {
	app *newrelic.Application
}

func NewRelicInstrumenter(app *newrelic.Application) *Instrumenter {
	return &Instrumenter{app: app}
}

func (i *Instrumenter) Start(ctx context.Context, name string) (context.Context, application.Transaction) {
	txn := i.app.StartTransaction(name)
	return newrelic.NewContext(ctx, txn), txn
}
