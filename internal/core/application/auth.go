package application

import (
	"context"
	"errors"
	"github.com/justjack1521/mevium/pkg/mevent"
	"mevway/internal/core/domain/auth"
	"mevway/internal/core/domain/user"
	"mevway/internal/core/port"
)

var (
	errPasswordConfirmMismatch = errors.New("passwords do not match")
)

type AuthenticationService struct {
	publisher *mevent.Publisher
	tokens    port.TokenRepository
}

func NewAuthenticationService(publisher *mevent.Publisher, tokens port.TokenRepository) *AuthenticationService {
	return &AuthenticationService{publisher: publisher, tokens: tokens}
}

func (s *AuthenticationService) Login(ctx context.Context, target user.User) (auth.LoginResult, error) {

	if target.Username != "kollaps" {
		return auth.LoginResult{}, errors.New("login failed")
	}

	if err := target.HasValidLoginCredentials(); err != nil {
		return auth.LoginResult{}, err
	}

	login, err := s.tokens.CreateToken(ctx, target)
	if err != nil {
		return auth.LoginResult{}, err
	}
	return login, nil
}
