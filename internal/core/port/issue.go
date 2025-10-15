package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
)

type IssueRepository interface {
	GetOpenIssuesList(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
	GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error)
	GetTopLevelIssueList(ctx context.Context) ([]patch.Issue, error)
}

type IssueService interface {
	ListOpenIssues(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
	ListTopIssues(ctx context.Context) ([]patch.Issue, error)
	GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error)
}
