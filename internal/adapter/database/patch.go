package database

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/domain/patch"
)

type PatchRepository struct {
	database *gorm.DB
}

func NewPatchRepository(db *gorm.DB) *PatchRepository {
	return &PatchRepository{database: db}
}

func (r *PatchRepository) Get(ctx context.Context, limit int) ([]patch.Patch, error) {
	var cond dto.PatchGorm
	var res []dto.PatchGorm
	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Limit(limit).Order("release_date").Find(&res, cond).Error; err != nil {
		return nil, err
	}

	dest := make([]patch.Patch, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}
