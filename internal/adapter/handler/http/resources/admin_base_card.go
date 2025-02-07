package resources

type ValidateAbilityCard struct {
	AbilityCard AbilityCard `json:"AbilityCard"`
}

type AbilityCard struct {
	SysID    string   `json:"SysID"`
	BaseCard BaseCard `json:"BaseCard"`
}

type BaseCard struct {
	SysID     string `json:"SysID"`
	Name      string `json:"Name"`
	AbilityID string `json:"AbilityID"`
}

type ValidateBaseCardResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
