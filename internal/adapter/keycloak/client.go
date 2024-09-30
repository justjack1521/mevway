package keycloak

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/domain/auth"
)

var (
	errTokenAuthoriseFailed = func(err error) error {
		return fmt.Errorf("token authorisation failed: %w", err)
	}
	errTokenExtractionFailed = func(err error) error {
		return fmt.Errorf("token extraction failed: %w", err)
	}
	errTokenInactive = errors.New("token no longer active")
)

type Client struct {
	gocloak      *gocloak.GoCloak
	clientID     string
	clientSecret string
	realm        string
}

func NewClient(config mevconn.KeyCloakConfig) *Client {
	return &Client{
		gocloak:      gocloak.NewClient(config.Hostname()),
		clientID:     config.ClientID(),
		clientSecret: config.ClientSecret(),
		realm:        config.Realm(),
	}
}

func (c *Client) Login(ctx context.Context, request auth.LoginRequest) (auth.LoginResponse, error) {

	jwt, err := c.gocloak.Login(ctx, c.clientID, c.clientSecret, c.realm, request.Username, request.Password)
	if err != nil {
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}, nil

}

var (
	errFailedToRegisterUser = func(err error) error {
		return fmt.Errorf("failed to register user: %w", err)
	}
)

func (c *Client) Register(ctx context.Context, username string, password string) (uuid.UUID, error) {

	jwt, err := c.gocloak.LoginAdmin(ctx, "admin", "admin", "master")
	if err != nil {
		return uuid.Nil, errFailedToRegisterUser(err)
	}

	id, err := c.gocloak.CreateUser(ctx, jwt.AccessToken, c.realm, gocloak.User{
		Username: gocloak.StringP(username),
		Enabled:  gocloak.BoolP(true),
	})
	if err != nil {
		return uuid.Nil, errFailedToRegisterUser(err)
	}

	if err := c.gocloak.SetPassword(ctx, jwt.AccessToken, id, c.realm, password, false); err != nil {
		return uuid.Nil, errFailedToRegisterUser(err)
	}

	return uuid.FromStringOrNil(id), nil

}

func (c *Client) ExtractToken(ctx context.Context, token string) (auth.TokenClaims, error) {

	_, claims, err := c.gocloak.DecodeAccessToken(ctx, token, c.realm)

	if err != nil {
		return auth.TokenClaims{}, errTokenExtractionFailed(err)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return auth.TokenClaims{}, errTokenExtractionFailed(err)
	}
	usr, err := uuid.FromString(sub)
	if err != nil {
		return auth.TokenClaims{}, errTokenExtractionFailed(err)
	}

	return auth.TokenClaims{
		UserID:      usr,
		PlayerID:    uuid.UUID{},
		Environment: "",
	}, nil

}

func (c *Client) AuthoriseToken(ctx context.Context, token string) error {

	result, err := c.gocloak.RetrospectToken(ctx, token, c.clientID, c.clientSecret, c.realm)
	if err != nil {
		return errTokenAuthoriseFailed(err)
	}

	if *result.Active == false {
		return errTokenAuthoriseFailed(errTokenInactive)
	}

	return nil

}
