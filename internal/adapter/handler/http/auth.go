package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"mevway/internal/domain/user"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationType      = "bearer"
)

const (
	UserIDContextKey   string = "UserIDContextKey"
	PlayerIDContextKey string = "PlayerIDContextKey"
	UserEnvironmentKey string = "UserEnvironmentKey"
)

type AuthenticationHandler struct {
	svc        port.AuthenticationService
	repository port.TokenRepository
}

func NewAuthenticationHandler(svc port.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{svc: svc}
}

func (h *AuthenticationHandler) Login(ctx *gin.Context) {

	var request = &resources.UserLoginRequest{}

	if err := ctx.BindJSON(request); err != nil {
		return
	}

	result, err := h.svc.Login(ctx, user.User{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		return
	}

	ctx.JSON(200, resources.UserLoginResponse{
		IDToken:      result.IDToken,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	})

}

func (h *AuthenticationHandler) Register(ctx *gin.Context) {

	var request = &resources.UserRegisterRequest{}

	if err := ctx.BindJSON(request); err != nil {
		return
	}

	result, err := h.svc.Register(ctx, request.Username, request.Password, request.ConfirmPassword)
	if err != nil {
		return
	}

	ctx.JSON(200, resources.UserRegisterResponse{SysUser: result.UserID})

}

func (m *AuthenticationHandler) TokenAuthorise(ctx *gin.Context) {

	var header = ctx.GetHeader(authorizationHeaderKey)

	if len(header) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var fields = strings.Fields(header)

	if len(fields) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var method = strings.ToLower(fields[0])
	if method != authorizationType {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var token = fields[1]
	claims, err := m.repository.VerifyToken(ctx, token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	ctx.Set(UserIDContextKey, claims.UserID)
	ctx.Set(PlayerIDContextKey, claims.PlayerID)
	ctx.Set(UserEnvironmentKey, claims.Environment)

}