package resources

type CreateAugmentMaterialsRequest struct {
	AbilityCardID string            `json:"AbilityCardID" binding:"required"`
	Materials     []AugmentMaterial `json:"Materials"`
}

type AugmentMaterial struct {
	SysID    string `json:"SysID"`
	Quantity int    `json:"Quantity"`
}

type CreateAugmentMaterialsResponse struct {
	Error        bool   `json:"Error"`
	ErrorMessage string `json:"ErrorMessage"`
}
