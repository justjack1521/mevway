package game

import uuid "github.com/satori/go.uuid"

type SkillPanel struct {
	DefinitionType string
	Index          int
	ReferenceID    uuid.UUID
	Value          int
	CostItems      []CostItem
}

type CostItem struct {
	ItemID uuid.UUID
	Value  int
}
