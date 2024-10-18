package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) List(ctx *gin.Context) {

	users, err := h.svc.List(ctx, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(200, resources.NewUserIdentityListResponse(users))

}

func (h *UserHandler) Register(ctx *gin.Context) {

	var request = &resources.UserRegisterRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.svc.Register(ctx, request.Username, request.Password, request.ConfirmPassword)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.UserRegisterResponse{SysUser: result.ID})

}

func (h *UserHandler) Delete(ctx *gin.Context) {

	var request = &resources.UserDeleteRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.svc.Delete(ctx, request.UserID); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.UserDeleteResponse{})

}
