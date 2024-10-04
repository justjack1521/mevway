package application

import (
	"context"
	"errors"
	"mevway/internal/core/port"
	"mevway/internal/domain/auth"
	"mevway/internal/domain/user"
)

var (
	errPasswordConfirmMismatch = errors.New("passwords do not match")
)

type AuthenticationService struct {
	tokens port.TokenRepository
	users  port.UserRepository
}

func NewAuthenticationService(users port.UserRepository, tokens port.TokenRepository) *AuthenticationService {
	return &AuthenticationService{users: users, tokens: tokens}
}

func (s *AuthenticationService) Login(ctx context.Context, target user.User) (auth.LoginResult, error) {

	if err := target.HasValidLoginCredentials(); err != nil {
		return auth.LoginResult{}, err
	}

	login, err := s.tokens.CreateToken(ctx, target)
	if err != nil {
		return auth.LoginResult{}, err
	}
	return login, nil
}

func (s *AuthenticationService) Register(ctx context.Context, username, password, confirm string) (user.User, error) {

	if password != confirm {
		return user.User{}, errPasswordConfirmMismatch
	}

	target, err := user.NewUser(username, password)
	if err != nil {
		return user.User{}, err
	}

	if err := s.users.CreateUser(ctx, target); err != nil {
		return user.User{}, err
	}

	return target, nil

}
