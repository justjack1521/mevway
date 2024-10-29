package resources

type CreateSkillPanelRequest struct {
	BaseJobID string     `json:"BaseJobID" binding:"required"`
	PageIndex int        `json:"PageIndex"`
	Panel     SkillPanel `json:"Panel"`
}

type SkillPanel struct {
	DefinitionType string     `json:"DefinitionType" binding:"required"`
	Index          int        `json:"Index"`
	ReferenceID    string     `json:"ReferenceID"`
	Value          int        `json:"Value"`
	CostItems      []CostItem `json:"CostItems"`
}

type CostItem struct {
	ItemID string `json:"ItemID" binding:"required"`
	Value  int    `json:"Value"`
}
