package database

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/core/domain/patch"
)

type IssueRepository struct {
	database *gorm.DB
}

func NewIssueRepository(database *gorm.DB) *IssueRepository {
	return &IssueRepository{database: database}
}

func (r *IssueRepository) GetOpenIssuesList(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error) {

	var results []dto.KnownIssueGorm
	if err := r.database.WithContext(ctx).Table("system.known_issues AS issue").
		Select("issue.*").
		Joins("LEFT JOIN system.patch AS patch ON issue.fixed_by = patch.sys_id").
		Where("issue.fixed_by IS NULL OR patch.released = false").
		Find(&results).Error; err != nil {
		return nil, err
	}

	var dest = make([]patch.KnownIssue, len(results))
	for index, value := range results {
		dest[index] = value.ToEntity()
	}

	return dest, nil

}

func (r *IssueRepository) GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error) {

	var cond = &dto.IssueGorm{SysID: id}
	var res = &dto.IssueGorm{}

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).First(res, cond).Error; err != nil {
		return patch.Issue{}, err
	}

	return res.ToEntity(), nil

}

func (r *IssueRepository) GetTopLevelIssueList(ctx context.Context) ([]patch.Issue, error) {

	var res []dto.IssueGorm

	if err := r.database.WithContext(ctx).Model(&dto.IssueGorm{}).Not("state IN ?", patch.ClosedStates).Where("category = ?", int(patch.IssueCategoryGame)).Order("number DESC").Find(&res).Error; err != nil {
		return nil, err
	}

	var results = make([]patch.Issue, len(res))
	for index, value := range res {
		results[index] = value.ToEntity()
	}

	return results, nil

}

func (r IssueRepository) IssueHasWorkaround(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64 = 0
	if err := r.database.WithContext(ctx).Model(&dto.IssueWorkaroundGorm{}).Where("issue_id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
