package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
)

type UserRepository interface {
	GetUser(ctx context.Context, id uuid.UUID) (user.Identity, error)
	GetAllUsers(ctx context.Context, count, offset int) ([]user.Identity, error)
	CreateUser(ctx context.Context, target *user.User) error
	DeleteUser(ctx context.Context, target user.Identity) error
	ChangePassword(ctx context.Context, target *user.User) error
}

type UserService interface {
	Get(ctx context.Context, id uuid.UUID) (user.Identity, error)
	List(ctx context.Context, count, offset int) ([]user.Identity, error)
	Register(ctx context.Context, username, password, confirm string) (*user.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ChangePassword(ctx context.Context, target user.Identity, password, confirm string) error
}

type IdentityRepository interface {
	IdentityFromCustomerID(ctx context.Context, customer string) (user.Identity, error)
}
