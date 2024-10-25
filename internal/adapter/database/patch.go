package database

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/core/domain/patch"
)

type PatchRepository struct {
	database *gorm.DB
}

func NewPatchRepository(db *gorm.DB) *PatchRepository {
	return &PatchRepository{database: db}
}

func (r *PatchRepository) GetPatchList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error) {
	var cond = &dto.PatchGorm{
		Released: true,
	}
	var res []dto.PatchGorm
	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Limit(limit).Order("release_date DESC").Find(&res, cond).Error; err != nil {
		return nil, err
	}

	dest := make([]patch.Patch, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}

func (r *PatchRepository) GetLatestPatch(ctx context.Context, environment uuid.UUID) (patch.Patch, error) {

	var cond = &dto.PatchGorm{
		Released: true,
	}
	var res = &dto.PatchGorm{}

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Limit(1).Order("release_date DESC").First(res, cond).Error; err != nil {
		return patch.Patch{}, err
	}

	return res.ToEntity(), nil

}

func (r *PatchRepository) GetOpenIssuesList(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error) {

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
