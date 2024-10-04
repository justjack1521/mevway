package keycloak

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/domain/user"
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

type UserClient struct {
	client *gocloak.GoCloak
	config mevconn.KeyCloakConfig
}

func NewUserClient(client *gocloak.GoCloak, config mevconn.KeyCloakConfig) *UserClient {
	return &UserClient{
		client: client,
		config: config,
	}
}

var (
	errFailedToRegisterUser = func(err error) error {
		return fmt.Errorf("failed to register user: %w", err)
	}
)

func (c *UserClient) CreateUser(ctx context.Context, target user.User) error {

	token, err := c.LoginAdmin(ctx)
	if err != nil {
		return errFailedToRegisterUser(err)
	}

	var credentials = []gocloak.CredentialRepresentation{
		{
			Type:  gocloak.StringP("password"),
			Value: gocloak.StringP(target.Password),
		},
	}

	var roles = []string{
		"default-roles-mevius",
	}

	var attributes = map[string][]string{
		"profile":  {target.PlayerID.String()},
		"customer": {target.CustomerID},
	}

	_, err = c.client.CreateUser(ctx, token, c.config.Realm(), gocloak.User{
		ID:          gocloak.StringP(target.UserID.String()),
		Username:    gocloak.StringP(target.Username),
		Enabled:     gocloak.BoolP(true),
		Credentials: &credentials,
		RealmRoles:  &roles,
		Attributes:  &attributes,
	})

	if err != nil {
		return errFailedToRegisterUser(err)
	}

	return nil

}

var (
	errFailedAdminLogin = func(err error) error {
		return fmt.Errorf("failed to log in as admin: %w", err)
	}
)

func (c *UserClient) LoginAdmin(ctx context.Context) (string, error) {

	username, password := c.config.AdminCredentials()

	tkn, err := c.client.LoginAdmin(ctx, username, password, "master")
	if err != nil {
		return "", errFailedAdminLogin(err)
	}

	return tkn.AccessToken, nil

}

func (c *UserClient) GetPlayerIDFromCustomerID(ctx context.Context, id string) (uuid.UUID, error) {

	token, err := c.LoginAdmin(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	users, err := c.client.GetUsers(ctx, token, c.config.Realm(), gocloak.GetUsersParams{
		Q: gocloak.StringP("customer:f03c-33ce-41ed"),
	})

	if err != nil {
		return uuid.Nil, err
	}

	var attrs = *users[0].Attributes
	profile, _ := attrs["profile"]
	return uuid.FromStringOrNil(profile[0]), nil

}
