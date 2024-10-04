package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func UserIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(UserIDContextKey)
	if value == "" {
		return uuid.Nil, errors.New("context missing user id")
	}
	result, err := uuid.FromString(value)
	if err != nil {
		return uuid.Nil, err
	}
	if result == uuid.Nil {
		return uuid.Nil, errors.New("context missing user id")
	}
	return result, nil
}

func PlayerIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetString(PlayerIDContextKey)
	if value == "" {
		return uuid.Nil, errors.New("context missing player id")
	}
	result, err := uuid.FromString(value)
	if err != nil {
		return uuid.Nil, err
	}
	if result == uuid.Nil {
		return uuid.Nil, errors.New("context missing player id")
	}
	return result, nil
}

func SessionIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	var value = ctx.GetHeader("X-API-SESSION")
	if value == "" {
		return uuid.Nil, errors.New("context missing session id")
	}
	result, err := uuid.FromString(value)
	if err != nil {
		return uuid.Nil, err
	}
	if result == uuid.Nil {
		return uuid.Nil, errors.New("context missing session id")
	}
	return result, nil
}
