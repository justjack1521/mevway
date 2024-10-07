package application

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type contextKey struct{}

type ContextMetadata struct {
	UserID   uuid.UUID
	PlayerID uuid.UUID
}

func NewApplicationContext(ctx context.Context, md ContextMetadata) context.Context {
	return context.WithValue(ctx, contextKey{}, md)
}

func MetadataFromContext(ctx context.Context) ContextMetadata {
	md, ok := ctx.Value(contextKey{}).(ContextMetadata)
	if ok == false {
		return ContextMetadata{}
	}
	return md
}
