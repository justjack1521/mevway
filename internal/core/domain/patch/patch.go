package patch

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Patch struct {
	SysID       uuid.UUID
	Application uuid.UUID
	ReleaseDate time.Time
	Version     string
	Description string
	Released    bool
	Features    []GameFeature
	Fixes       []KnownIssue
}

type KnownIssue struct {
	SysID uuid.UUID
	Text  string
}

type GameFeature struct {
	SysID uuid.UUID
	Text  string
}
