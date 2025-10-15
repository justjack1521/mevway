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

func (r *PatchRepository) GetPatchListCount(ctx context.Context) (int, error) {

	var cond = &dto.PatchGorm{
		Released:    true,
		Show:        true,
		Application: "game",
	}

	var count int64

	if err := r.database.WithContext(ctx).Model(cond).Where(count).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil

}

func (r *PatchRepository) GetPatchList(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error) {

	var cond = &dto.PatchGorm{
		Released:    true,
		Show:        true,
		Application: "game",
	}

	var res []dto.PatchGorm
	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Offset(offset).Limit(limit).Order("release_date DESC").Find(&res, cond).Error; err != nil {
		return nil, err
	}

	var dest = make([]patch.Patch, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}

func (r *PatchRepository) GetLatestPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error) {

	var cond = &dto.PatchGorm{
		Application: application,
		Released:    true,
	}
	var res = &dto.PatchGorm{}

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Limit(1).Order("release_date DESC").First(res, cond).Error; err != nil {
		return patch.Patch{}, err
	}

	return res.ToEntity(), nil

}

func (r *PatchRepository) GetAllowedPatchList(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error) {
	var cond = &dto.PatchGorm{
		Application: application,
		Allowed:     true,
	}

	var res []dto.PatchGorm

	if err := r.database.WithContext(ctx).Model(cond).Order("release_date DESC").Find(&res, cond).Error; err != nil {
		return nil, err
	}

	var dest = make([]patch.Patch, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}
