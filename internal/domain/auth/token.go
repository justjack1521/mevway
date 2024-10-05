package auth

import uuid "github.com/satori/go.uuid"

type TokenAuthoriseRequest struct {
	Token string
}

type TokenClaims struct {
	SessionID   uuid.UUID
	UserID      uuid.UUID
	PlayerID    uuid.UUID
	Environment string
	Roles       []string
}
