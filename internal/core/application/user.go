package application

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
	"mevway/internal/core/port"
)

type UserService struct {
	publisher *mevent.Publisher
	users     port.UserRepository
}

func (s *UserService) List(ctx context.Context, count, offset int) ([]user.Identity, error) {
	return s.users.GetAllUsers(ctx, count, offset)
}

func NewUserService(publisher *mevent.Publisher, users port.UserRepository) *UserService {
	return &UserService{publisher: publisher, users: users}
}

func (s *UserService) Register(ctx context.Context, username, password, confirm string) (*user.User, error) {

	if password != confirm {
		return nil, errPasswordConfirmMismatch
	}

	target, err := user.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	if err := s.users.CreateUser(ctx, target); err != nil {
		return nil, err
	}

	s.publisher.Notify(user.NewCreatedEvent(ctx, target.ID, target.PlayerID, target.CustomerID))

	return target, nil

}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {

	if err := s.users.DeleteUser(ctx, user.Identity{ID: id}); err != nil {
		return err
	}

	s.publisher.Notify(user.NewDeleteEvent(ctx, id, id))

	return nil

}

type UserEventTranslator interface {
	Created(event user.CreatedEvent) ([]byte, error)
	Deleted(event user.DeleteEvent) ([]byte, error)
}
