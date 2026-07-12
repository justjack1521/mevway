package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"mevway/internal/core/port"
)

var (
	errUserDisabled = errors.New("user is disabled")
)

type UserStatusMiddleware struct {
	users port.UserService
}

func NewUserStatusMiddleware(users port.UserService) *UserStatusMiddleware {
	return &UserStatusMiddleware{users: users}
}

func (m *UserStatusMiddleware) Handle(ctx *gin.Context) {

	id, err := UserIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	identity, err := m.users.Get(ctx, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if identity.Enabled == false {
		ctx.AbortWithError(http.StatusForbidden, errUserDisabled)
		return
	}

	ctx.Next()

}
