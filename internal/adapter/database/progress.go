package database

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/core/domain/progress"
)

type ProgressRepository struct {
	database *gorm.DB
}

func NewProgressRepository(database *gorm.DB) *ProgressRepository {
	return &ProgressRepository{database: database}
}

func (r *ProgressRepository) GetProgressList(ctx context.Context) ([]progress.GameFeature, error) {

	var cond = &dto.GameFeatureProgressGorm{
		Active: true,
	}

	var res []dto.GameFeatureProgressGorm

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Find(&res, cond).Error; err != nil {
		return nil, err
	}

	dest := make([]progress.GameFeature, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}
