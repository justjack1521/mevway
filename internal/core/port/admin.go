package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type GameAdminService interface {
	GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error
}
