package resources

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type PatchListResponse struct {
	Patches []Patch `json:"patches"`
}

type Patch struct {
	SysID       uuid.UUID      `json:"sys_id"`
	ReleaseDate time.Time      `json:"release_date"`
	Description string         `json:"description"`
	Features    []PatchFeature `json:"features"`
	Fixes       []PatchFix     `json:"fixes"`
}

type PatchFeature struct {
	Text  string `json:"text"`
	Order int    `json:"order"`
}

type PatchFix struct {
	Text  string `json:"text"`
	Order int    `json:"order"`
}
