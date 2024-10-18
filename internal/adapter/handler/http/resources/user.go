package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
)

type UserIdentityResponse struct {
	UserID     uuid.UUID
	PlayerID   uuid.UUID
	CustomerID string
}

func NewUserIdentityListResponse(identities []user.Identity) []UserIdentityResponse {
	var result = make([]UserIdentityResponse, len(identities))
	for index, value := range identities {
		result[index] = NewUserIdentityResponse(value)
	}
	return result
}

func NewUserIdentityResponse(identity user.Identity) UserIdentityResponse {
	return UserIdentityResponse{
		UserID:     identity.ID,
		PlayerID:   identity.PlayerID,
		CustomerID: identity.CustomerID,
	}
}
