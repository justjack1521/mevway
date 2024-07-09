package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

func (a *PublicAPIRouter) HandleRememberUser(ctx *gin.Context) {

	var request = &resources.RememberUserRequest{}
	if err := ctx.BindJSON(request); err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.RememberUserHandler.Handle(ctx, handler.RememberUser{
		Token:    request.Token,
		DeviceID: a.device(ctx),
	})

}
