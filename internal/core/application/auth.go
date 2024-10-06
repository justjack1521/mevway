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
	tokens    port.TokenRepository
	users     port.UserRepository
	publisher *mevent.Publisher
}

func NewAuthenticationService(tokens port.TokenRepository, users port.UserRepository, publisher *mevent.Publisher) *AuthenticationService {
	return &AuthenticationService{tokens: tokens, users: users, publisher: publisher}
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

	s.publisher.Notify(user.NewCreatedEvent(ctx, target.ID, target.PlayerID, target.CustomerID))

	return target, nil

}
