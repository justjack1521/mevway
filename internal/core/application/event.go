package application

import (
	"context"
)

type StartEvent struct {
	ctx context.Context
}

func NewStartEvent(ctx context.Context) StartEvent {
	return StartEvent{ctx: ctx}
}

func (e StartEvent) Context() context.Context {
	return e.ctx
}

func (e StartEvent) Name() string {
	return "application.start"
}

type ShutdownEvent struct {
	ctx context.Context
}

func NewShutdownEvent(ctx context.Context) ShutdownEvent {
	return ShutdownEvent{ctx: ctx}
}

func (e ShutdownEvent) Name() string {
	return "application.shutdown"
}

func (e ShutdownEvent) Context() context.Context {
	return e.ctx
}
