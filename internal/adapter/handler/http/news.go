package http

import (
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type NewsHandler struct {
	svc port.NewsService
}

func NewNewsHandler(svc port.NewsService) *NewsHandler {
	return &NewsHandler{svc: svc}
}

func (h *NewsHandler) Get(ctx *gin.Context) {

	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	list, err := h.svc.GetNewsArticle(ctx, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result, err := resources.NewNewsArticle(list)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, result)
}
