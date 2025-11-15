package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type IssueListResponse struct {
	Issues []IssueSummary `json:"Issues"`
}

func NewIssueSummaryListResponse(p []patch.IssueSummary) IssueListResponse {
	var response = IssueListResponse{Issues: make([]IssueSummary, len(p))}
	for index, value := range p {
		response.Issues[index] = NewIssueSummary(value)
	}
	return response
}

type IssueSummary struct {
	SysID         uuid.UUID `json:"SysID"`
	Number        int       `json:"Number"`
	Description   string    `json:"Description"`
	Category      int       `json:"Category"`
	State         int       `json:"State"`
	CreatedAt     time.Time `json:"CreatedAt"`
	HasWorkaround bool      `json:"HasWorkaround"`
}

func NewIssueSummary(p patch.IssueSummary) IssueSummary {
	return IssueSummary{
		SysID:         p.SysID,
		Number:        p.Number,
		Description:   p.Description,
		Category:      p.Category,
		State:         p.State,
		CreatedAt:     p.CreatedAt,
		HasWorkaround: p.HasWorkaround,
	}
}

type Issue struct {
	SysID       uuid.UUID         `json:"SysID"`
	Number      int               `json:"Number"`
	Description string            `json:"Description"`
	Category    int               `json:"Category"`
	State       int               `json:"State"`
	CreatedAt   time.Time         `json:"CreatedAt"`
	Workarounds []IssueWorkaround `json:"Workarounds"`
}

type IssueWorkaround struct {
	SysID       uuid.UUID `json:"SysID"`
	Description string    `json:"Description"`
}

func NewIssueWorkaround(p patch.IssueWorkaround) IssueWorkaround {
	return IssueWorkaround{
		SysID:       p.SysID,
		Description: p.Description,
	}
}

func NewIssue(p patch.Issue) Issue {
	var result = Issue{
		SysID:       p.SysID,
		Number:      p.Number,
		Description: p.Description,
		Category:    p.Category,
		State:       p.State,
		CreatedAt:   p.CreatedAt,
		Workarounds: make([]IssueWorkaround, len(p.Workarounds)),
	}
	for index, value := range p.Workarounds {
		result.Workarounds[index] = NewIssueWorkaround(value)
	}
	return result
}
