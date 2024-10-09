package user

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type CreatedEvent struct {
	ctx        context.Context
	userID     uuid.UUID
	playerID   uuid.UUID
	customerID string
}

func NewCreatedEvent(ctx context.Context, usr uuid.UUID, player uuid.UUID, customer string) CreatedEvent {
	return CreatedEvent{ctx: ctx, userID: usr, playerID: player, customerID: customer}
}

func (e CreatedEvent) Name() string {
	return "event.user.created"
}

func (e CreatedEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name":  e.Name(),
		"user.id":     e.userID.String(),
		"player.id":   e.playerID.String(),
		"customer.id": e.customerID,
	}
}

func (e CreatedEvent) Context() context.Context {
	return e.ctx
}

func (e CreatedEvent) UserID() uuid.UUID {
	return e.userID
}

func (e CreatedEvent) PlayerID() uuid.UUID {
	return e.playerID
}

func (e CreatedEvent) CustomerID() string {
	return e.customerID
}

type DeleteEvent struct {
	ctx    context.Context
	userID uuid.UUID
}

func NewDeleteEvent(ctx context.Context, id uuid.UUID) DeleteEvent {
	return DeleteEvent{ctx: ctx, userID: id}
}

func (e DeleteEvent) Name() string {
	return "event.user.deleted"
}

func (e DeleteEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{
		"event.name": e.Name(),
		"user.id":    e.userID.String(),
	}
}

func (e DeleteEvent) Context() context.Context {
	return e.ctx
}

func (e DeleteEvent) UserID() uuid.UUID {
	return e.userID
}
