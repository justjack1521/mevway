package http

import (
	"github.com/gin-gonic/gin"
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

	results, err := h.svc.ListTopRankings(ctx, ctx.Param("code"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewListRankingResponse(results))

}
