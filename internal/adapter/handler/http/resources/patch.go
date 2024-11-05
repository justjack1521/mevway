package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type KnownIssueListResponse struct {
	Issues []KnownIssue `json:"Issues"`
}

func NewKnowLIssueListResponse(p []patch.KnownIssue) KnownIssueListResponse {
	var response = KnownIssueListResponse{Issues: make([]KnownIssue, len(p))}
	for index, value := range p {
		response.Issues[index] = NewKnownIssue(value)
		response.Issues[index].Order = index
	}
	return response
}

type KnownIssue struct {
	SysID uuid.UUID `json:"SysID"`
	Text  string    `json:"Text"`
	Order int       `json:"Order"`
}

func NewKnownIssue(p patch.KnownIssue) KnownIssue {
	return KnownIssue{
		SysID: p.SysID,
		Text:  p.Text,
	}
}

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
	SysID       uuid.UUID     `json:"SysID"`
	ReleaseDate time.Time     `json:"ReleaseDate"`
	Description string        `json:"Description"`
	Version     string        `json:"Version"`
	Features    []GameFeature `json:"Features"`
	Fixes       []KnownIssue  `json:"Fixes"`
}

func NewPatchResponse(p patch.Patch) Patch {
	var response = Patch{
		SysID:       p.SysID,
		ReleaseDate: p.ReleaseDate,
		Description: p.Description,
		Version:     p.Version,
		Features:    make([]GameFeature, len(p.Features)),
		Fixes:       make([]KnownIssue, len(p.Fixes)),
	}
	for index, value := range p.Features {
		response.Features[index] = NewPatchFeature(value)
		response.Features[index].Order = index
	}
	for index, value := range p.Fixes {
		response.Fixes[index] = NewKnownIssue(value)
		response.Fixes[index].Order = index
	}
	return response
}

type GameFeature struct {
	SysID uuid.UUID `json:"SysID"`
	Text  string    `json:"Text"`
	Order int       `json:"Order"`
}

func NewPatchFeature(p patch.GameFeature) GameFeature {
	return GameFeature{
		SysID: p.SysID,
		Text:  p.Text,
	}
}
