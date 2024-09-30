package user

import (
	"crypto/rand"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserID     uuid.UUID
	PlayerID   uuid.UUID
	CustomerID string
	Username   string
	Password   string
}

func NewUser(username string, password string) User {
	b := make([]byte, 6)
	rand.Read(b)
	return User{
		UserID:     uuid.NewV4(),
		PlayerID:   uuid.NewV4(),
		CustomerID: fmt.Sprintf("%x-%x-%x", b[0:2], b[2:4], b[4:]),
		Username:   username,
		Password:   password,
	}
}
