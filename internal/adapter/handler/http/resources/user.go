package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
)

type UserIdentityListResponse struct {
	Count int                    `json:"Count"`
	List  []UserIdentityResponse `json:"List"`
}

type UserIdentityResponse struct {
	UserID     uuid.UUID `json:"UserID"`
	PlayerID   uuid.UUID `json:"PlayerID"`
	CustomerID string    `json:"CustomerID"`
}

func NewUserIdentityListResponse(identities []user.Identity) UserIdentityListResponse {

	var results = UserIdentityListResponse{
		Count: len(identities),
		List:  make([]UserIdentityResponse, len(identities)),
	}

	for index, value := range identities {
		results.List[index] = NewUserIdentityResponse(value)
	}
	return results
}

func NewUserIdentityResponse(identity user.Identity) UserIdentityResponse {
	return UserIdentityResponse{
		UserID:     identity.ID,
		PlayerID:   identity.PlayerID,
		CustomerID: identity.CustomerID,
	}
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"CurrentPassword"`
	NewPassword     string `json:"NewPassword"`
	ConfirmPassword string `json:"ConfirmPassword"`
}
