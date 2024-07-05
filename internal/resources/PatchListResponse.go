package resources

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type PatchListResponse struct {
	Patches []Patch `json:"Patches"`
}

type Patch struct {
	SysID       uuid.UUID      `json:"SysID"`
	ReleaseDate time.Time      `json:"ReleaseDate"`
	Description string         `json:"Description"`
	Features    []PatchFeature `json:"Features"`
	Fixes       []PatchFix     `json:"Fixes"`
}

type PatchFeature struct {
	Text  string `json:"Text"`
	Order int    `json:"Order"`
}

type PatchFix struct {
	Text  string `json:"Text"`
	Order int    `json:"Order"`
}
