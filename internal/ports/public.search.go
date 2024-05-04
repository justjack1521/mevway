package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

func (a *PublicAPIRouter) HandlePlayerSearch(ctx *gin.Context) {

	var request = resources.PlayerSearchRequest{
		CustomerID: ctx.Param("customer_id"),
	}

	user, err := a.user(ctx)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	a.PlayerSearchHandle.Handle(ctx, handler.PlayerSearch{
		UserID:     user,
		CustomerID: request.CustomerID,
	})

}
