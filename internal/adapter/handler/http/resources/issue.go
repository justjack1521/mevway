package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
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
	SysID uuid.UUID `json:"SysID"`
}

func NewIssue(p patch.Issue) Issue {
	return Issue{
		SysID: p.SysID,
	}
}
