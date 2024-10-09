package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/domain/user"
	"mevway/internal/core/port"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationType      = "bearer"
)

type AuthenticationHandler struct {
	svc    port.AuthenticationService
	tokens port.TokenRepository
}

func NewAuthenticationHandler(svc port.AuthenticationService, repository port.TokenRepository) *AuthenticationHandler {
	return &AuthenticationHandler{svc: svc, tokens: repository}
}

func (h *AuthenticationHandler) Login(ctx *gin.Context) {

	var request = &resources.UserLoginRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.svc.Login(ctx, user.User{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.UserLoginResponse{
		IDToken:      result.IDToken,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})

}

func (h *AuthenticationHandler) Identity(ctx *gin.Context) {

	token, err := h.getAuthorisationToken(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	claims, err := h.tokens.VerifyIdentityToken(ctx, token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(200, resources.NewPlayerIdentityResponse(claims))

}

func (h *AuthenticationHandler) AccessTokenAuthorise(ctx *gin.Context) {

	token, err := h.getAuthorisationToken(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	claims, err := h.tokens.VerifyAccessToken(ctx, token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.Set(middleware.SessionIDContextKey, claims.SessionID.String())
	ctx.Set(middleware.UserIDContextKey, claims.UserID.String())
	ctx.Set(middleware.PlayerIDContextKey, claims.PlayerID.String())
	ctx.Set(middleware.UserEnvironmentKey, claims.Environment)
	ctx.Set(middleware.UserRoleContextKey, claims.Roles)

}

var (
	errNoAuthorisationHeader          = errors.New("authorisation header is missing")
	errMalformedAuthorisationHeader   = errors.New("malformed authorisation header")
	errUnsupportedAuthorisationMethod = errors.New("unsupported authorisation method")
)

func (h *AuthenticationHandler) getAuthorisationToken(ctx *gin.Context) (string, error) {
	var header = ctx.GetHeader(authorizationHeaderKey)

	if len(header) == 0 {
		return "", errNoAuthorisationHeader
	}

	var fields = strings.Fields(header)

	if len(fields) != 2 {
		return "", errMalformedAuthorisationHeader
	}

	var method = strings.ToLower(fields[0])
	if method != authorizationType {
		return "", errUnsupportedAuthorisationMethod
	}

	return fields[1], nil
}
