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
	Features    []Feature
	Fixes       []KnownIssue
}

type Feature struct {
	Text  string
	Order int
}

type Fix struct {
	Text  string
	Order int
}

type KnownIssue struct {
	SysID uuid.UUID
	Text  string
}
