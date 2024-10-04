package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	adminRoleName = "mevius-admin"
)

func RoleFromContext(ctx *gin.Context, role string) error {
	var values = ctx.GetStringSlice(role)
	for _, value := range values {
		if value == role {
			return nil
		}
	}
	return errors.New("context missing role")
}

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := RoleFromContext(c, adminRoleName); err != nil {
			c.AbortWithError(403, err)
		}
	}
}
