package resources

type ValidateAbilityCardRequest struct {
	AbilityCard AbilityCard `json:"AbilityCard"`
}

type AbilityCard struct {
	SysID             string   `json:"SysID"`
	Active            bool     `json:"Active"`
	BaseCard          BaseCard `json:"BaseCard"`
	OverrideAbilityID string   `json:"OverrideAbilityID"`
	FusionEXPOverride int      `json:"FusionEXPOverride"`
	SaleGoldOverride  int      `json:"SaleGoldOverride"`
}

type BaseCard struct {
	SysID           string `json:"SysID"`
	Active          bool   `json:"Active"`
	Name            string `json:"Name"`
	AbilityID       string `json:"AbilityID"`
	SkillSeedOne    string `json:"SkillSeedOne"`
	SkillSeedTwo    string `json:"SkillSeedTwo"`
	SkillSeedSplit  string `json:"SkillSeedSplit"`
	SeedFusionBoost int    `json:"SeedFusionBoost"`
	Category        string `json:"Category"`
	FastLearner     bool   `json:"FastLearner"`
}

type ValidateModelResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
