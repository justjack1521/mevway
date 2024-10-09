package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, target *user.User) error
	DeleteUser(ctx context.Context, target user.Identity) error
}

type UserService interface {
	Register(ctx context.Context, username, password, confirm string) (*user.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type IdentityRepository interface {
	IdentityFromCustomerID(ctx context.Context, customer string) (user.Identity, error)
}
