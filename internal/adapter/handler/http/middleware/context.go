package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/application"
	"net"
	"strings"
)

const (
	SessionIDContextKey    string = "SessionIDContextKey"
	UserIDContextKey       string = "UserIDContextKey"
	PlayerIDContextKey     string = "PlayerIDContextKey"
	UserEnvironmentKey     string = "UserEnvironmentKey"
	UserRoleContextKey     string = "UserRoleContextKey"
	PatchVersionContextKey string = "PatchVersionContextKey"
)

var (
	errContextMissingRoles = errors.New("context missing roles")
	errContextMissingRole  = func(role string) error {
		return fmt.Errorf("context missing role %s", role)
	}
)

func IPFromContext(ctx *gin.Context) net.IP {
	var forward = ctx.GetHeader("X-Forwarded-For")

	var ip net.IP
	if forward != "" {
		parts := strings.Split(forward, ",")
		ip = net.ParseIP(strings.TrimSpace(parts[0]))
	} else {
		ip = net.ParseIP(ctx.ClientIP())
	}

	return ip

}

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

func ApplicationContextFromMetadata(ctx *gin.Context) (context.Context, error) {
	user, err := UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	player, err := PlayerIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var md = application.ContextMetadata{
		UserID:   user,
		PlayerID: player,
	}
	return application.NewApplicationContext(ctx, md), nil
}

var (
	errContextMissingPatchID   = errors.New("context missing patch id")
	errContextPatchIDMalformed = errors.New("context patch id malformed")
)

func PatchIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(PatchVersionContextKey)
	if value == "" {
		return uuid.Nil, errContextMissingPatchID
	}
	result, err := uuid.FromString(value)
	if err != nil || result == uuid.Nil {
		return uuid.Nil, errContextPatchIDMalformed
	}
	return result, nil
}
