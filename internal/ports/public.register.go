package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

func (a *PublicAPIRouter) HandleRegisterUser(ctx *gin.Context) {

	request, err := resources.Binder[resources.UserRegisterRequest](ctx, resources.UserRegisterRequest{})

	if err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.RegisterUserHandle.Handle(ctx, handler.RegisterUser{
		Username:        request.Username,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	})

}
