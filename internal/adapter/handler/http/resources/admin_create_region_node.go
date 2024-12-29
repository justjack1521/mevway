package resources

type CreateRegionNodeRequest struct {
	SysID         string             `json:"SysID" binding:"required"`
	RegionMapID   string             `json:"RegionMapID"`
	Index         int                `json:"Index"`
	Name          string             `json:"Name"`
	StaminaCost   int                `json:"StaminaCost"`
	EnemyLevel    int                `json:"EnemyLevel"`
	DisableRevive int                `json:"DisableRevive"`
	Fiends        []RegionNodeFiends `json:"Fiends"`
}

type RegionNodeFiends struct {
	SysID string `json:"SysID"`
	Count int    `json:"Count"`
}
