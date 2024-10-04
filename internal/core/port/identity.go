package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type IdentityRepository interface {
	GetPlayerIDFromCustomerID(ctx context.Context, customer string) (uuid.UUID, error)
}
