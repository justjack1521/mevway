package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type KnownIssueGorm struct {
	SysID   uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Text    string    `gorm:"column:text"`
	FixedBy uuid.UUID `gorm:"column:fixed_by"`
}

func (KnownIssueGorm) TableName() string {
	return "system.known_issues"
}

func (x *KnownIssueGorm) ToEntity() patch.KnownIssue {
	return patch.KnownIssue{
		SysID: x.SysID,
		Text:  x.Text,
	}
}

type GameFeatureGorm struct {
	SysID      uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Text       string    `gorm:"column:text"`
	ReleasedBy uuid.UUID `gorm:"column:released_by"`
}

func (GameFeatureGorm) TableName() string {
	return "system.game_feature"
}

func (x *GameFeatureGorm) ToEntity() patch.GameFeature {
	return patch.GameFeature{
		SysID: x.SysID,
		Text:  x.Text,
	}
}

type PatchGorm struct {
	SysID       uuid.UUID          `gorm:"primaryKey;column:sys_id"`
	ReleaseDate time.Time          `gorm:"column:release_date"`
	Application string             `gorm:"column:application"`
	Version     string             `gorm:"column:version"`
	Description string             `gorm:"column:description"`
	Released    bool               `gorm:"column:released"`
	Features    []*GameFeatureGorm `gorm:"foreignKey:ReleasedBy"`
	Fixes       []*KnownIssueGorm  `gorm:"foreignKey:FixedBy"`
}

func (PatchGorm) TableName() string {
	return "system.patch"
}

func (x *PatchGorm) ToEntity() patch.Patch {
	var result = patch.Patch{
		SysID:       x.SysID,
		Application: x.Application,
		Version:     x.Version,
		ReleaseDate: x.ReleaseDate,
		Description: x.Description,
		Released:    x.Released,
	}

	if x.Features != nil {
		result.Features = make([]patch.GameFeature, len(x.Features))
		for index, feature := range x.Features {
			result.Features[index] = feature.ToEntity()
		}
	}

	if x.Fixes != nil {
		result.Fixes = make([]patch.KnownIssue, len(x.Fixes))
		for index, fix := range x.Fixes {
			result.Fixes[index] = fix.ToEntity()
		}
	}

	return result
}
