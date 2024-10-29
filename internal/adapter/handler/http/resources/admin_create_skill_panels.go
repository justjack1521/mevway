package resources

import uuid "github.com/satori/go.uuid"

type CreateSkillPanelRequest struct {
	BaseJobID uuid.UUID  `json:"BaseJobID" binding:"required"`
	PageIndex int        `json:"PageIndex"`
	Panel     SkillPanel `json:"Panel"`
}

type SkillPanel struct {
	DefinitionType string     `json:"DefinitionType" binding:"required"`
	Index          int        `json:"Index"`
	ReferenceID    uuid.UUID  `json:"ReferenceID"`
	Value          int        `json:"Value"`
	CostItems      []CostItem `json:"CostItems"`
}

type CostItem struct {
	ItemID uuid.UUID `json:"ItemID" binding:"required"`
	Value  int       `json:"Value"`
}
