package resources

type CreateBaseJobRequest struct {
	BaseJobID string `json:"BaseJobID" binding:"required"`
	Active    bool   `json:"Active"`
	Name      string `json:"Name"`
	Number    string `json:"Number"`
	TypeID    string `json:"TypeID"`
}

type CreateBaseJobResponse struct {
	Created bool `json:"Created"`
}

type CreateAbilityRequesst struct {
	SysID      string `json:"SysID" binding:"required"`
	Active     bool   `json:"Active"`
	Name       string `json:"Name"`
	ElementID  string `json:"ElementID"`
	CardTypeID string `json:"CardTypeID"`
}

type CreateAbilityResponse struct {
	Created bool `json:"Created"`
}
