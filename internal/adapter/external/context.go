package external

import (
	"context"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/core/application"
)

func OutgoingContext(ctx context.Context) context.Context {
	var md = application.MetadataFromContext(ctx)
	return mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID)
}
