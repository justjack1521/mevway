package port

import (
	"context"
	"mevway/internal/core/domain/auth"
	"mevway/internal/core/domain/user"
)

type AuthenticationService interface {
	Login(ctx context.Context, target user.User) (auth.LoginResult, error)
	Register(ctx context.Context, username, password, confirm string) (user.User, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, target user.User) (auth.LoginResult, error)
	VerifyToken(ctx context.Context, token string) (auth.TokenClaims, error)
}
