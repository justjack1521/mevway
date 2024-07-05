package patch

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Patch struct {
	SysID       uuid.UUID
	ReleaseDate time.Time
	Description string
	Features    []Feature
	Fixes       []Fix
}

type Feature struct {
	Text  string
	Order int
}

type Fix struct {
	Text  string
	Order int
}
