package auth

import uuid "github.com/satori/go.uuid"

type TokenAuthoriseRequest struct {
	Token string
}

type IdentityClaims struct {
	PlayerID   uuid.UUID
	Username   string
	CustomerID string
}

type AccessClaims struct {
	SessionID   uuid.UUID
	UserID      uuid.UUID
	PlayerID    uuid.UUID
	Environment string
	Roles       []string
}
