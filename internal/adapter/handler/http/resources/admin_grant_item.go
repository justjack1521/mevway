package resources

import uuid "github.com/satori/go.uuid"

type AdminGrantItemRequest struct {
	PlayerID uuid.UUID `json:"PlayerID" binding:"required"`
	ItemID   uuid.UUID `json:"ItemID" binding:"required"`
	Quantity int       `json:"Quantity" binding:"required"`
}

type AdminGrantItemResponse struct {
}
