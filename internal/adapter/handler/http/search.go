package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type SearchHandler struct {
	svc port.PlayerSearchService
}

func NewSearchHandler(svc port.PlayerSearchService) *SearchHandler {
	return &SearchHandler{svc: svc}
}

func (h *SearchHandler) Search(ctx *gin.Context) {

	var request = resources.PlayerSearchRequest{
		CustomerID: ctx.Param("customer_id"),
	}

	fmt.Println(request.CustomerID)

	result, err := h.svc.Search(ctx, request.CustomerID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPlayerSearchResponse(result))

}
