package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	adminRoleName = "mevius-admin"
)

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := RoleFromContext(c, adminRoleName); err != nil {
			c.AbortWithError(403, err)
		}
	}
}
