package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type RankHandler struct {
	svc port.RankService
}

func NewRankHandler(svc port.RankService) *RankHandler {
	return &RankHandler{svc: svc}
}

func (h *RankHandler) Top(ctx *gin.Context) {

	ctx.Set(middleware.UserIDContextKey, uuid.NewV4())
	ctx.Set(middleware.PlayerIDContextKey, uuid.NewV4())

	results, err := h.svc.ListTopRankings(ctx, ctx.Param("code"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewListRankingResponse(results))

}
