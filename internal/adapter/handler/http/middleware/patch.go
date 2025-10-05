package middleware

import "github.com/gin-gonic/gin"

const patchHeaderKey = "X-API-PATCH"

type patchMiddleware struct {
}

func NewPatchMiddleware() *patchMiddleware {
	return &patchMiddleware{}
}

func (m *patchMiddleware) Handle(ctx *gin.Context) {
	var id = ctx.GetHeader(patchHeaderKey)
	ctx.Set(PatchVersionContextKey, id)
	ctx.Next()
}
