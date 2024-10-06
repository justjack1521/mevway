package port

import (
	"context"
	"mevway/internal/core/domain/user"
)

type IdentityRepository interface {
	IdentityFromCustomerID(ctx context.Context, customer string) (user.Identity, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, target user.User) error
}
