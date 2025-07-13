package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/progress"
)

type GameFeatureProgressGorm struct {
	SysID   uuid.UUID                        `gorm:"primaryKey;column:sys_id"`
	Active  bool                             `gorm:"column:active"`
	Title   string                           `gorm:"column:title"`
	Metrics []*GameFeatureProgressMetricGorm `gorm:"foreignKey:Parent"`
}

func (GameFeatureProgressGorm) TableName() string {
	return "system.game_feature_progress"
}

func (x *GameFeatureProgressGorm) ToEntity() progress.GameFeature {
	var result = progress.GameFeature{
		SysID: x.SysID,
		Title: x.Title,
	}
	if x.Metrics != nil {
		result.Metrics = make([]progress.GameFeatureMetric, len(x.Metrics))
		for index, value := range x.Metrics {
			result.Metrics[index] = value.ToEntity()
		}
	}
	return result
}

type GameFeatureProgressMetricGorm struct {
	SysID  uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Parent uuid.UUID `gorm:"column:parent"`
	Title  string    `gorm:"column:title"`
	Value  int       `gorm:"column:value"`
	Total  int       `gorm:"column:total"`
}

func (x *GameFeatureProgressMetricGorm) ToEntity() progress.GameFeatureMetric {
	return progress.GameFeatureMetric{
		SysID: x.SysID,
		Title: x.Title,
		Value: x.Value,
		Total: x.Total,
	}
}

func (GameFeatureProgressMetricGorm) TableName() string {
	return "system.game_feature_progress_metric"
}
