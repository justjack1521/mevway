package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	SessionIDContextKey string = "SessionIDContextKey"
	UserIDContextKey    string = "UserIDContextKey"
	PlayerIDContextKey  string = "PlayerIDContextKey"
	UserEnvironmentKey  string = "UserEnvironmentKey"
	UserRoleContextKey  string = "UserRoleContextKey"
)

var (
	errContextMissingRoles = errors.New("context missing roles")
	errContextMissingRole  = func(role string) error {
		return fmt.Errorf("context missing role %s", role)
	}
)

func RoleFromContext(ctx *gin.Context, role string) error {
	var values = ctx.GetStringSlice(UserRoleContextKey)

	if len(values) == 0 {
		return errContextMissingRoles
	}

	for _, value := range values {
		if value == role {
			return nil
		}
	}
	return errContextMissingRole(role)
}

var (
	errContextMissingSessionID   = errors.New("context missing session id")
	errContextSessionIDMalformed = errors.New("context session id malformed")
)

func SessionIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(SessionIDContextKey)
	if value == "" {
		return uuid.Nil, errContextMissingSessionID
	}
	result, err := uuid.FromString(value)
	if err != nil || result == uuid.Nil {
		return uuid.Nil, errContextSessionIDMalformed
	}
	return result, nil
}

var (
	errContextMissingUserID   = errors.New("context missing user id")
	errContextUserIDMalformed = errors.New("context user id malformed")
)

func UserIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(UserIDContextKey)
	if value == "" {
		return uuid.Nil, errContextMissingUserID
	}
	result, err := uuid.FromString(value)
	if err != nil || result == uuid.Nil {
		return uuid.Nil, errContextUserIDMalformed
	}
	return result, nil
}

var (
	errContextMissingPlayerID   = errors.New("context missing player id")
	errContextPlayerIDMalformed = errors.New("context player id malformed")
)

func PlayerIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(PlayerIDContextKey)
	if value == "" {
		return uuid.Nil, errContextMissingPlayerID
	}
	result, err := uuid.FromString(value)
	if err != nil || result == uuid.Nil {
		return uuid.Nil, errContextPlayerIDMalformed
	}
	return result, nil
}
