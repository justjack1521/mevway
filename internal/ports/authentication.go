package ports

import (
	"context"
	"mevway/internal/domain/auth"
)

type AuthenticationClient interface {
	Login(ctx context.Context, request auth.LoginRequest) (auth.LoginResponse, error)
}
