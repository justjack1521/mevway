package resources

type ValidateBaseItemRequest struct {
	BaseItem BaseItem `json:"Item"`
}

type BaseItem struct {
	SysID          string `json:"SysID"`
	Active         bool   `json:"Active"`
	Name           string `json:"Name"`
	Maximum        int    `json:"Maximum"`
	MonthlyMaximum int    `json:"MonthlyMaximum"`
}
