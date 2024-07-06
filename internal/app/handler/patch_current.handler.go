package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
	"mevway/internal/domain/patch"
	"mevway/internal/resources"
)

type PatchCurrent struct {
}

type PatchCurrentHandler decorator.APIRouterHandler[PatchCurrent]

type patchCurrentHandler struct {
	repository patch.ReadRepository
}

func NewPatchCurrentHandler(repository patch.ReadRepository) PatchCurrentHandler {
	return patchCurrentHandler{repository: repository}
}

func (h patchCurrentHandler) Handle(ctx *gin.Context, query PatchCurrent) {

	current, err := h.repository.Current(ctx)
	if err != nil {
		httperr.InternalError(err, err.Error(), ctx)
		return
	}

	var response = resources.Patch{
		SysID:       current.SysID,
		ReleaseDate: current.ReleaseDate,
		Description: current.Description,
	}

	ctx.JSON(200, response)

}
