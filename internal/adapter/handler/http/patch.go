package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type PatchHandler struct {
	svc port.PatchService
}

func NewPatchHandler(svc port.PatchService) *PatchHandler {
	return &PatchHandler{svc: svc}
}

func (h *PatchHandler) Recent(ctx *gin.Context) {

	var application = ctx.Query("application")

	if application == "" {
		application = "game"
	}

	current, err := h.svc.GetCurrentPatch(ctx, application, uuid.Nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPatchResponse(current))

}

func (h *PatchHandler) Allow(ctx *gin.Context) {

	target, err := uuid.FromString(ctx.Query("current"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var application = ctx.Query("application")

	if application == "" {
		application = "game"
	}

	allowed, err := h.svc.ListAllowPatches(ctx, application, uuid.Nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, value := range allowed {
		if uuid.Equal(value.SysID, target) {
			ctx.Status(200)
			return
		}
	}

	ctx.AbortWithStatus(http.StatusNotFound)

}

func (h *PatchHandler) List(ctx *gin.Context) {

	list, err := h.svc.ListPatches(ctx, uuid.Nil, 5)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPatchListResponse(list))

}

func (h *PatchHandler) Issues(ctx *gin.Context) {

	list, err := h.svc.ListOpenIssues(ctx, uuid.Nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewKnowLIssueListResponse(list))

}

func (h *PatchHandler) Top(ctx *gin.Context) {
	list, err := h.svc.ListTopIssues(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewIssueListResponse(list))
}

func (h *PatchHandler) Get(ctx *gin.Context) {
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
