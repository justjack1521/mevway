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
