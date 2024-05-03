package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
)

func (a *PublicAPIRouter) HandleSocket(ctx *gin.Context) {

	user, err := a.session(ctx)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	player, err := a.player(ctx)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	a.WebsocketHandle.Handle(ctx, handler.WebSocketQuery{
		UserID:   user,
		PlayerID: player,
	})
}
