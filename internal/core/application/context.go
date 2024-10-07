package application

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type contextKey struct{}

type ContextMetadata struct {
	UserID   uuid.UUID
	PlayerID uuid.UUID
}

func NewApplicationContext(ctx context.Context, md ContextMetadata) context.Context {
	fmt.Println(md.UserID)
	fmt.Println(md.PlayerID)
	return context.WithValue(ctx, contextKey{}, md)
}

func MetadataFromContext(ctx context.Context) ContextMetadata {
	md, ok := ctx.Value(contextKey{}).(ContextMetadata)
	if ok == false {
		return ContextMetadata{}
	}
	return md
}
