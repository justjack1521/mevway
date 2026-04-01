package external

import (
	"context"
	"errors"
	"mevway/internal/core/application"

	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
)

func OutgoingContext(ctx context.Context) (context.Context, error) {
	var md = application.MetadataFromContext(ctx)
	if uuid.Equal(md.UserID, uuid.Nil) {
		return ctx, errors.New("failed to generate outgoing context, user id is nil")
	}
	if uuid.Equal(md.PlayerID, uuid.Nil) {
		return ctx, errors.New("failed to generate outgoing context, player id is nil")
	}
	return mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID), nil
}
