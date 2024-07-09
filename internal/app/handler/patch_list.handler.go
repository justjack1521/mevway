package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
	"mevway/internal/domain/patch"
	"mevway/internal/resources"
)

type PatchList struct {
	Environment uuid.UUID
	Limit       int
	Offset      int
}

type PatchListHandler decorator.APIRouterHandler[PatchList]

type patchListHandler struct {
	repository patch.ReadRepository
}

func NewPatchListHandler(repository patch.ReadRepository) PatchListHandler {
	return patchListHandler{repository: repository}
}

func (h patchListHandler) Handle(ctx *gin.Context, query PatchList) {

	patches, err := h.repository.Get(ctx, query.Environment, query.Limit)

	if err != nil {
		httperr.InternalError(err, err.Error(), ctx)
		return
	}

	var response = resources.PatchListResponse{Patches: make([]resources.Patch, len(patches))}

	for index, value := range patches {

		var item = resources.Patch{
			SysID:       value.SysID,
			ReleaseDate: value.ReleaseDate,
			Description: value.Description,
			Features:    make([]resources.PatchFeature, len(value.Features)),
			Fixes:       make([]resources.PatchFix, len(value.Fixes)),
		}

		for j, feature := range value.Features {
			item.Features[j] = resources.PatchFeature{
				Text:  feature.Text,
				Order: feature.Order,
			}
		}

		for k, fix := range value.Fixes {
			item.Fixes[k] = resources.PatchFix{
				Text:  fix.Text,
				Order: fix.Order,
			}
		}

		response.Patches[index] = item

	}

	ctx.JSON(200, response)

}
