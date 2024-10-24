package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type KnownIssueGorm struct {
	SysID uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Text  string    `gorm:"column:text"`
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

type PatchGorm struct {
	SysID       uuid.UUID           `gorm:"primaryKey;column:sys_id"`
	ReleaseDate time.Time           `gorm:"column:release_date"`
	Description string              `gorm:"column:description"`
	Released    bool                `gorm:"column:released"`
	Environment uuid.UUID           `gorm:"column:environment"`
	Features    []*PatchFeatureGorm `gorm:"foreignKey:Patch"`
	Fixes       []*PatchFixGorm     `gorm:"foreignKey:Patch"`
}

func (PatchGorm) TableName() string {
	return "system.patch"
}

func (x *PatchGorm) ToEntity() patch.Patch {
	var result = patch.Patch{
		SysID:       x.SysID,
		ReleaseDate: x.ReleaseDate,
		Description: x.Description,
		Released:    x.Released,
		Environment: x.Environment,
	}

	if x.Features != nil {
		result.Features = make([]patch.Feature, len(x.Features))
		for index, feature := range x.Features {
			result.Features[index] = feature.ToEntity()
		}
	}

	if x.Fixes != nil {
		result.Fixes = make([]patch.Fix, len(x.Fixes))
		for index, fix := range x.Fixes {
			result.Fixes[index] = fix.ToEntity()
		}
	}

	return result
}

type PatchFeatureGorm struct {
	Patch uuid.UUID `gorm:"column:patch"`
	Text  string    `gorm:"column:text"`
	Order int       `gorm:"column:order"`
}

func (PatchFeatureGorm) TableName() string {
	return "system.patch_feature"
}

func (x *PatchFeatureGorm) ToEntity() patch.Feature {
	return patch.Feature{
		Text:  x.Text,
		Order: x.Order,
	}
}

type PatchFixGorm struct {
	Patch uuid.UUID `gorm:"column:patch"`
	Text  string    `gorm:"column:text"`
	Order int       `gorm:"column:order"`
}

func (PatchFixGorm) TableName() string {
	return "system.patch_fix"
}

func (x *PatchFixGorm) ToEntity() patch.Fix {
	return patch.Fix{
		Text:  x.Text,
		Order: x.Order,
	}
}
