package user

import (
	"crypto/rand"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

var (
	errUserNameIsEmpty = errors.New("username is empty")
	errPasswordIsEmpty = errors.New("password is empty")
)

type User struct {
	UserID     uuid.UUID
	PlayerID   uuid.UUID
	CustomerID string
	Username   string
	Password   string
}

func NewUser(username string, password string) (User, error) {

	if username == "" {
		return User{}, errUserNameIsEmpty
	}

	if password == "" {
		return User{}, errPasswordIsEmpty
	}

	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return User{}, err
	}

	return User{
		UserID:     uuid.NewV4(),
		PlayerID:   uuid.NewV4(),
		CustomerID: fmt.Sprintf("%x-%x-%x", b[0:2], b[2:4], b[4:]),
		Username:   username,
		Password:   password,
	}, nil
}

func (u User) HasValidLoginCredentials() error {
	if u.Username == "" {
		return errUserNameIsEmpty
	}
	if u.Password == "" {
		return errPasswordIsEmpty
	}
	return nil
}
