package relic

import (
	"context"
	"github.com/justjack1521/mevrabbit"
	"github.com/newrelic/go-agent/v3/newrelic"
	"mevway/internal/core/application"
)

type Tracer struct {
	app *newrelic.Application
}

func NewRelicTracer(app *newrelic.Application) *Tracer {
	return &Tracer{app: app}
}

func (t *Tracer) Start(ctx context.Context, name string) (context.Context, application.Transaction) {
	txn := t.app.StartTransaction(name)
	return newrelic.NewContext(ctx, txn), txn
}

func (t *Tracer) NewRabbitMQTransaction(ctx context.Context, name string) (context.Context, mevrabbit.Transaction) {
	txn := t.app.StartTransaction(name)
	return newrelic.NewContext(ctx, txn), txn
}

func (t *Tracer) NewRabbitMQSegment(ctx context.Context, exchange mevrabbit.Exchange) mevrabbit.Segment {
	var txn = newrelic.FromContext(ctx)
	if txn == nil {
		return nil
	}
	segment := newrelic.MessageProducerSegment{
		StartTime:            txn.StartSegmentNow(),
		Library:              "RabbitMQ",
		DestinationType:      newrelic.MessageExchange,
		DestinationName:      string(exchange),
		DestinationTemporary: false,
	}
	return &segment
}
