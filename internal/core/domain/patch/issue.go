package patch

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type IssueState int

const (
	IssueStateNew IssueState = iota
	IssueStateAccepted
	IssueStateInProgress
	IssueStateReady
	IssueStateResolved
	IssueStateCancelled
)

var (
	ClosedStates = []IssueState{IssueStateResolved, IssueStateCancelled}
)

type IssueCategory int

const (
	IssueCategoryGame IssueCategory = iota
	IssueCategoryLauncher
	IssueCategoryAccount
)

type Issue struct {
	Number      int
	SysID       uuid.UUID
	Description string
	State       int
	Category    int
	ParentIssue uuid.UUID
	CreatedAt   time.Time
}

type IssueWorkaround struct {
	SysID       uuid.UUID
	Description string
}

type IssueSummary struct {
	Number        int
	SysID         uuid.UUID
	Description   string
	State         int
	Category      int
	CreatedAt     time.Time
	HasWorkaround bool
}
