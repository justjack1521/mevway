package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/app/handler"
)

func (a *PublicAPIRouter) HandleSocket(ctx *gin.Context) {

	client, err := uuid.FromString(a.client(ctx))
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	a.WebsocketHandle.Handle(ctx, handler.WebSocketQuery{
		ClientID: client,
	})
}
