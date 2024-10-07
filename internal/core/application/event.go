package application

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ShutdownEvent struct {
	ctx context.Context
}

func NewShutdownEvent(ctx context.Context) ShutdownEvent {
	return ShutdownEvent{ctx: ctx}
}

func (e ShutdownEvent) Name() string {
	return "application.shutdown"
}

func (e ShutdownEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{}
}

func (e ShutdownEvent) Context() context.Context {
	return e.ctx
}
