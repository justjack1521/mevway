package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/content"
	"time"
)

type GameFeatureProgressGorm struct {
	SysID     uuid.UUID                        `gorm:"primaryKey;column:sys_id"`
	Active    bool                             `gorm:"column:active"`
	Title     string                           `gorm:"column:title"`
	CreatedAt time.Time                        `gorm:"column:created_at"`
	UpdatedAt time.Time                        `gorm:"column:updated_at"`
	Metrics   []*GameFeatureProgressMetricGorm `gorm:"foreignKey:Parent"`
}

func (GameFeatureProgressGorm) TableName() string {
	return "system.game_feature_progress"
}

func (x *GameFeatureProgressGorm) ToEntity() content.GameFeature {
	var result = content.GameFeature{
		SysID:     x.SysID,
		Title:     x.Title,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
	if x.Metrics != nil {
		result.Metrics = make([]content.GameFeatureMetric, len(x.Metrics))
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

func (x *GameFeatureProgressMetricGorm) ToEntity() content.GameFeatureMetric {
	return content.GameFeatureMetric{
		SysID: x.SysID,
		Title: x.Title,
		Value: x.Value,
		Total: x.Total,
	}
}

func (GameFeatureProgressMetricGorm) TableName() string {
	return "system.game_feature_progress_metric"
}
