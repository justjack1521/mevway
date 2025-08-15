package database

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mevway/internal/adapter/database/dto"
	"mevway/internal/core/domain/content"
)

type ProgressRepository struct {
	database *gorm.DB
}

func NewProgressRepository(database *gorm.DB) *ProgressRepository {
	return &ProgressRepository{database: database}
}

func (r *ProgressRepository) GetReleaseList(ctx context.Context) ([]content.GameFeatureRelease, error) {

	var cond = &dto.GameFeatureReleaseContentGorm{}

	var res []dto.GameFeatureReleaseContentGorm

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Find(&res, cond).Error; err != nil {
		return nil, err
	}

	dest := make([]content.GameFeatureRelease, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}

func (r *ProgressRepository) GetProgressList(ctx context.Context) ([]content.GameFeature, error) {

	var cond = &dto.GameFeatureProgressGorm{
		Active: true,
	}

	var res []dto.GameFeatureProgressGorm

	if err := r.database.WithContext(ctx).Model(cond).Preload(clause.Associations).Find(&res, cond).Error; err != nil {
		return nil, err
	}

	dest := make([]content.GameFeature, len(res))
	for i, v := range res {
		dest[i] = v.ToEntity()
	}
	return dest, nil

}
