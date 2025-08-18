package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/content"
	"time"
)

type GameFeatureReleaseContentGorm struct {
	SysID     uuid.UUID                            `gorm:"primaryKey;column:sys_id"`
	Name      string                               `gorm:"column:name"`
	Banner    string                               `gorm:"column:banner"`
	CreatedAt time.Time                            `gorm:"column:created_at"`
	Items     []*GameFeatureReleaseContentItemGorm `gorm:"foreignKey:Parent"`
	Active    bool                                 `gorm:"column:active"`
}

func (GameFeatureReleaseContentGorm) TableName() string {
	return "system.game_feature_release_content"
}

func (x *GameFeatureReleaseContentGorm) ToEntity() content.GameFeatureRelease {
	var result = content.GameFeatureRelease{
		SysID:     x.SysID,
		Name:      x.Name,
		BannerURL: x.Banner,
	}
	if x.Items != nil {
		result.Items = make([]content.GameFeatureReleaseItem, len(x.Items))
		for index, value := range x.Items {
			result.Items[index] = value.ToEntity()
		}
	}
	return result
}

type GameFeatureReleaseContentItemGorm struct {
	Parent  uuid.UUID `gorm:"column:parent"`
	Title   string    `gorm:"column:title"`
	Link    string    `gorm:"column:link"`
	Content string    `gorm:"column:content"`
}

func (GameFeatureReleaseContentItemGorm) TableName() string {
	return "system.game_feature_release_content_item"
}

func (x *GameFeatureReleaseContentItemGorm) ToEntity() content.GameFeatureReleaseItem {
	return content.GameFeatureReleaseItem{
		Title:   x.Title,
		Link:    x.Link,
		Content: x.Content,
	}
}
