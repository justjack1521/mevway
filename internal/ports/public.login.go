package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

func (a *PublicAPIRouter) HandleLoginUser(ctx *gin.Context) {

	var request = &resources.UserLoginRequest{}

	if err := ctx.BindJSON(request); err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.LoginUserHandle.Handle(ctx, handler.LoginUser{
		Username: request.Username,
		Password: request.Password,
		DeviceID: a.device(ctx),
	})

}
