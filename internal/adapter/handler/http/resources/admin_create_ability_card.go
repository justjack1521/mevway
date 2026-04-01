package resources

type CreateAbilityCardRequest struct {
	SysID             string `json:"SysID"`
	Active            bool   `json:"Active"`
	CardNumber        int    `json:"CardNumber"`
	BaseCardID        string `json:"BaseCardID"`
	InShop            bool   `json:"InShop"`
	AugmentConfigID   string `json:"AugmentConfigID"`
	OverrideAbilityID string `json:"OverrideAbilityID"`
	FusionEXPOverride int    `json:"FusionEXPOverride"`
	SaleGoldOverride  int    `json:"SaleGoldOverride"`
}

type CreateAbilityCardResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
