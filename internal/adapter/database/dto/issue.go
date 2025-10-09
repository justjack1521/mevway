package dto

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"time"
)

type IssueGorm struct {
	Number      int       `gorm:"column:number"`
	SysID       uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Description string    `gorm:"column:description"`
	Category    int       `gorm:"column:category"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ParentIssue uuid.UUID `gorm:"column:parent_issue"`
	State       int       `gorm:"column:state"`
}

func (IssueGorm) TableName() string {
	return "system.issue"
}

func (x *IssueGorm) ToEntity() patch.Issue {
	return patch.Issue{
		Number:      x.Number,
		SysID:       x.SysID,
		Description: x.Description,
		State:       x.State,
		Category:    x.Category,
		ParentIssue: x.ParentIssue,
		CreatedAt:   x.CreatedAt,
	}
}
