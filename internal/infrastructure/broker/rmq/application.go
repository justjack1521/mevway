package rmq

import (
	"github.com/justjack1521/mevrabbit"
	"github.com/wagslane/go-rabbitmq"
	"log/slog"
)

type ApplicationConnection struct {
	conn    *rabbitmq.Conn
	slogger *slog.Logger
	tracer  mevrabbit.TransactionTracer
}

func NewApplicationConnection(conn *rabbitmq.Conn, slogger *slog.Logger, tracer mevrabbit.TransactionTracer) *ApplicationConnection {
	return &ApplicationConnection{conn: conn, slogger: slogger, tracer: tracer}
}
