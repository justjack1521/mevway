package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/auth"
)

type PlayerIdentityResponse struct {
	PlayerID   uuid.UUID `json:"PlayerID"`
	Username   string    `json:"Username"`
	CustomerID string    `json:"CustomerID"`
}

func NewPlayerIdentityResponse(claims auth.IdentityClaims) PlayerIdentityResponse {
	return PlayerIdentityResponse{
		PlayerID:   claims.PlayerID,
		Username:   claims.Username,
		CustomerID: claims.CustomerID,
	}
}
