package resources

type CreateBaseCardRequest struct {
	SysID               string  `json:"SysID"`
	Active              bool    `json:"Active"`
	Name                string  `json:"Name"`
	AbilityID           string  `json:"AbilityID"`
	FiendCard           bool    `json:"FiendCard"`
	SkillSeedOne        string  `json:"SkillSeedOne"`
	SkillSeedTwo        string  `json:"SkillSeedTwo"`
	SkillSeedSplit      string  `json:"SkillSeedSplit"`
	FastLearner         bool    `json:"FastLearner"`
	SeedFusionBoost     float32 `json:"SeedFusionBoost"`
	EXPFusionMultiplier float32 `json:"EXPFusionMultiplier"`
}

type CreateBaseCardResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
