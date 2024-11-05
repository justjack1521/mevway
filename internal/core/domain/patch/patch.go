package patch

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Patch struct {
	SysID       uuid.UUID
	Application string
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

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) ToString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
