package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type IssueListResponse struct {
	Issues []Issue `json:"Issues"`
}

func NewIssueListResponse(p []patch.Issue) IssueListResponse {
	var response = IssueListResponse{Issues: make([]Issue, len(p))}
	for index, value := range p {
		response.Issues[index] = NewIssue(value)
	}
	return response
}

type Issue struct {
	SysID       uuid.UUID `json:"SysID"`
	Number      int       `json:"Number"`
	Description string    `json:"Description"`
	Category    int       `json:"Category"`
	State       int       `json:"State"`
	CreatedAt   time.Time `json:"CreatedAt"`
}

func NewIssue(p patch.Issue) Issue {
	return Issue{
		SysID:       p.SysID,
		Number:      p.Number,
		Description: p.Description,
		Category:    p.Category,
		State:       p.State,
		CreatedAt:   p.CreatedAt,
	}
}
