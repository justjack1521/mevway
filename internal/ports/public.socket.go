package ports

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/app/handler"
)

func (a *PublicAPIRouter) HandleSocket(ctx *gin.Context) {
	a.WebsocketHandle.Handle(ctx, handler.WebSocketQuery{
		ClientID: uuid.NewV4().String(), //a.client(ctx),
	})
}
