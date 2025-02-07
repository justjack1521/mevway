package resources

type ValidateAbilityCard struct {
	AbilityCard AbilityCard `json:"AbilityCard"`
}

type AbilityCard struct {
	SysID             string   `json:"SysID"`
	BaseCard          BaseCard `json:"BaseCard"`
	OverrideAbilityID string   `json:"OverrideAbilityID"`
	FusionEXPOverride int      `json:"FusionEXPOverride"`
	SaleGoldOverride  int      `json:"SaleGoldOverride"`
}

type BaseCard struct {
	SysID     string `json:"SysID"`
	Name      string `json:"Name"`
	AbilityID string `json:"AbilityID"`
}

type ValidateModelResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
