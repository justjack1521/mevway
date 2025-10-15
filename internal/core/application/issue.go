package application

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"mevway/internal/core/port"
)

type IssueService struct {
	repository port.IssueRepository
}

func NewIssueService(repository port.IssueRepository) *IssueService {
	return &IssueService{repository: repository}
}

func (s *IssueService) ListOpenIssues(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error) {
	return s.repository.GetOpenIssuesList(ctx, environment)
}

func (s *IssueService) ListTopIssues(ctx context.Context) ([]patch.Issue, error) {
	return s.repository.GetTopLevelIssueList(ctx)
}

func (s *IssueService) GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error) {
	return s.repository.GetIssue(ctx, id)
}
