package resources

import uuid "github.com/satori/go.uuid"

type UserDeleteRequest struct {
	UserID uuid.UUID `json:"UserID" binding:"required"`
}

type UserDeleteResponse struct {
}
