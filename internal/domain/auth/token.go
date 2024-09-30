package auth

import uuid "github.com/satori/go.uuid"

type TokenAuthoriseRequest struct {
	Token string
}

type TokenClaims struct {
	UserID      uuid.UUID
	PlayerID    uuid.UUID
	Environment string
}
