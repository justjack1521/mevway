package port

import (
	"context"
	"mevway/internal/core/domain/auth"
	"mevway/internal/core/domain/user"
)

type AuthenticationService interface {
	Login(ctx context.Context, target user.User) (auth.LoginResult, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, target user.User) (auth.LoginResult, error)
	VerifyAccessToken(ctx context.Context, token string) (auth.AccessClaims, error)
	VerifyIdentityToken(ctx context.Context, token string) (auth.IdentityClaims, error)
}
