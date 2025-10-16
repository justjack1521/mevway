package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
)

func RespondWithError(c *gin.Context, err resources.ErrorResponse) {
	c.Error(err)
	c.JSON(err.Code, err)
	c.Abort()
}
