package ports

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

func (a *PublicAPIRouter) HandlePlayerSearch(ctx *gin.Context) {

	var request = resources.PlayerSearchRequest{
		CustomerID: ctx.Param("customer_id"),
	}

	a.PlayerSearchHandle.Handle(ctx, handler.PlayerSearch{
		UserID:     a.client(ctx),
		CustomerID: request.CustomerID,
	})

}
