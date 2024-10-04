package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/domain/patch"
	"time"
)

type PatchListResponse struct {
	Patches []Patch `json:"Patches"`
}

func NewPatchListResponse(p []patch.Patch) PatchListResponse {
	var response = PatchListResponse{Patches: make([]Patch, len(p))}
	for index, value := range p {
		response.Patches[index] = NewPatchResponse(value)
	}
	return response
}

type Patch struct {
	SysID       uuid.UUID      `json:"SysID"`
	ReleaseDate time.Time      `json:"ReleaseDate"`
	Description string         `json:"Description"`
	Features    []PatchFeature `json:"Features"`
	Fixes       []PatchFix     `json:"Fixes"`
}

func NewPatchResponse(p patch.Patch) Patch {
	var response = Patch{
		SysID:       p.SysID,
		ReleaseDate: p.ReleaseDate,
		Description: p.Description,
		Features:    make([]PatchFeature, len(p.Features)),
		Fixes:       make([]PatchFix, len(p.Fixes)),
	}
	for index, value := range p.Features {
		response.Features[index] = NewPatchFeature(value)
	}
	for index, value := range p.Fixes {
		response.Fixes[index] = NewPatchFix(value)
	}
	return response
}

type PatchFeature struct {
	Text  string `json:"Text"`
	Order int    `json:"Order"`
}

func NewPatchFeature(p patch.Feature) PatchFeature {
	return PatchFeature{
		Text:  p.Text,
		Order: p.Order,
	}
}

type PatchFix struct {
	Text  string `json:"Text"`
	Order int    `json:"Order"`
}

func NewPatchFix(p patch.Fix) PatchFix {
	return PatchFix{
		Text:  p.Text,
		Order: p.Order,
	}
}
