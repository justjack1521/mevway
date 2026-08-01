package http

import (
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type RankHandler struct {
	svc port.RankService
}

func NewRankHandler(svc port.RankService) *RankHandler {
	return &RankHandler{svc: svc}
}

func (h *RankHandler) Top(ctx *gin.Context) {

	var md = application.ContextMetadata{
		UserID:   uuid.NewV4(),
		PlayerID: uuid.NewV4(),
	}

	results, err := h.svc.ListTopRankings(application.NewApplicationContext(ctx, md), ctx.Param("code"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewListRankingResponse(results))

}

func (h *RankHandler) Player(ctx *gin.Context) {

	p, err := uuid.FromString(ctx.Param("player"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var md = application.ContextMetadata{
		UserID:   uuid.NewV4(),
		PlayerID: p,
	}

	result, err := h.svc.ListPlayerRanking(application.NewApplicationContext(ctx, md), ctx.Param("code"), md.PlayerID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPlayerRankingResponse(result))

}

func (h *RankHandler) Range(ctx *gin.Context) {

	p, err := uuid.FromString(ctx.Param("player"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	start, err := strconv.Atoi(ctx.Param("start"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stop, err := strconv.Atoi(ctx.Param("stop"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var md = application.ContextMetadata{
		UserID:   uuid.NewV4(),
		PlayerID: p,
	}

	result, err := h.svc.ListRankingRange(application.NewApplicationContext(ctx, md), ctx.Param("code"), md.PlayerID, start, stop)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewListRankingResponse(result))

}
