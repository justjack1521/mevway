package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type IssueHandler struct {
	svc port.IssueService
}

func NewIssueHandler(svc port.IssueService) *IssueHandler {
	return &IssueHandler{svc: svc}
}

func (h *IssueHandler) Issues(ctx *gin.Context) {

	list, err := h.svc.ListOpenIssues(ctx, uuid.Nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewKnowLIssueListResponse(list))

}

func (h *IssueHandler) Top(ctx *gin.Context) {
	list, err := h.svc.ListTopIssues(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewIssueListResponse(list))
}

func (h *IssueHandler) Get(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	list, err := h.svc.GetIssue(ctx, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewIssue(list))
}
