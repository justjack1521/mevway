package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/auth"
	"mevway/internal/decorator"
)

const (
	UserIDContextKey   string = "UserIDContextKey"
	PlayerIDContextKey string = "PlayerIDContextKey"
	UserEnvironmentKey string = "UserEnvironmentKey"
)

type TokenAuthorise struct {
	SessionID uuid.UUID
	Bearer    string
	DeviceID  string
}

type TokenAuthorisationClient interface {
	AuthoriseToken(ctx context.Context, token string) error
	VerifyToken(ctx context.Context, token string) (auth.TokenClaims, error)
}

type TokenAuthoriseHandler decorator.APIRouterHandler[TokenAuthorise]

type tokenAuthoriseHandler struct {
	client TokenAuthorisationClient
}

func NewTokenHandler(clt TokenAuthorisationClient) TokenAuthoriseHandler {
	return tokenAuthoriseHandler{
		client: clt,
	}
}

func (h tokenAuthoriseHandler) Handle(ctx *gin.Context, query TokenAuthorise) {

	if err := h.client.AuthoriseToken(ctx, query.Bearer); err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	claims, err := h.client.VerifyToken(ctx, query.Bearer)
	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	fmt.Println(claims.UserID, claims.PlayerID, claims.Environment)

	ctx.Set(UserIDContextKey, claims.UserID)
	ctx.Set(PlayerIDContextKey, claims.PlayerID)
	ctx.Set(UserEnvironmentKey, claims.Environment)

}
