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
	gocloak *gocloak.GoCloak
	config  mevconn.KeyCloakConfig
}

func NewClient(config mevconn.KeyCloakConfig) *Client {
	return &Client{
		gocloak: gocloak.NewClient(config.Hostname()),
		config:  config,
	}
}

func (c *Client) Login(ctx context.Context, request auth.LoginRequest) (auth.LoginResponse, error) {

	jwt, err := c.gocloak.Login(ctx, c.config.ClientID(), c.config.ClientSecret(), c.config.Realm(), request.Username, request.Password)
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

	admin, creds := c.config.AdminCredentials()

	jwt, err := c.gocloak.LoginAdmin(ctx, admin, creds, "master")
	if err != nil {
		return uuid.Nil, errFailedToRegisterUser(err)
	}

	var credentials = []gocloak.CredentialRepresentation{
		{
			Type:  gocloak.StringP("password"),
			Value: gocloak.StringP(password),
		},
	}

	var roles = []string{
		"default-roles-mevius",
	}

	id, err := c.gocloak.CreateUser(ctx, jwt.AccessToken, c.config.Realm(), gocloak.User{
		Username:    gocloak.StringP(username),
		Enabled:     gocloak.BoolP(true),
		Credentials: &credentials,
		RealmRoles:  &roles,
	})
	if err != nil {
		return uuid.Nil, errFailedToRegisterUser(err)
	}

	return uuid.FromStringOrNil(id), nil

}

func (c *Client) ExtractToken(ctx context.Context, token string) (auth.TokenClaims, error) {

	_, claims, err := c.gocloak.DecodeAccessToken(ctx, token, c.config.ClientID())

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

	result, err := c.gocloak.RetrospectToken(ctx, token, c.config.ClientID(), c.config.ClientSecret(), c.config.Realm())
	if err != nil {
		return errTokenAuthoriseFailed(err)
	}

	if *result.Active == false {
		return errTokenAuthoriseFailed(errTokenInactive)
	}

	return nil

}
