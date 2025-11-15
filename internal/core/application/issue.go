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

func (s *IssueService) ListTopIssues(ctx context.Context) ([]patch.IssueSummary, error) {

	results, err := s.repository.GetTopLevelIssueList(ctx)
	if err != nil {
		return nil, err
	}

	var issues = make([]patch.IssueSummary, len(results))

	for index, value := range results {

		work, err := s.repository.IssueHasWorkaround(ctx, value.SysID)
		if err != nil {
			return nil, err
		}

		var summary = patch.IssueSummary{
			Number:        value.Number,
			SysID:         value.SysID,
			Description:   value.Description,
			State:         value.State,
			Category:      value.Category,
			CreatedAt:     value.CreatedAt,
			HasWorkaround: work,
		}

		issues[index] = summary

	}

	return issues, nil
}

func (s *IssueService) GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error) {
	return s.repository.GetIssue(ctx, id)
}
