package patch

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Patch struct {
	SysID       uuid.UUID
	ReleaseDate time.Time
	Description string
	Environment uuid.UUID
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
