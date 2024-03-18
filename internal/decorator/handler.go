package decorator

import "github.com/gin-gonic/gin"

type APIRouterHandler[Q any] interface {
	Handle(*gin.Context, Q)
}
