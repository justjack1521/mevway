package resources

import (
	"github.com/gin-gonic/gin"
)

func Binder[R any](ctx *gin.Context, input R) (R, error) {
	if err := ctx.Bind(&input); err != nil {
		return input, err
	}
	return input, nil
}
