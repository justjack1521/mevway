package ports

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/app/handler"
)

func (a *PublicAPIRouter) HandleSocket(ctx *gin.Context) {
	a.WebsocketHandle.Handle(ctx, handler.WebSocketQuery{
		ClientID: a.client(ctx),
	})
}
