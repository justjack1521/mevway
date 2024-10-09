package keycloak

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
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
	errFailedDeleteUser = func(err error) error {
		return fmt.Errorf("failed to delete user: %w", err)
	}
)

func (c *UserClient) DeleteUser(ctx context.Context, target user.Identity) error {

	token, err := c.LoginAdmin(ctx)
	if err != nil {
		return errFailedDeleteUser(err)
	}

	if err := c.client.DeleteUser(ctx, token, c.config.Realm(), target.ID.String()); err != nil {
		return errFailedDeleteUser(err)
	}

	return nil

}

var (
	errFailedToCreateUser = func(err error) error {
		return fmt.Errorf("failed to register user: %w", err)
	}
)

func (c *UserClient) CreateUser(ctx context.Context, target user.User) error {

	token, err := c.LoginAdmin(ctx)
	if err != nil {
		return errFailedToCreateUser(err)
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
		ID:          gocloak.StringP(target.ID.String()),
		Username:    gocloak.StringP(target.Username),
		Enabled:     gocloak.BoolP(true),
		Credentials: &credentials,
		RealmRoles:  &roles,
		Attributes:  &attributes,
	})

	if err != nil {
		return errFailedToCreateUser(err)
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

var (
	errUserMatchingCustomerIDNotFound  = errors.New("user matching customer id not found")
	errProfileAttributeNotFound        = errors.New("profile attribute not found")
	errFailedGetIdentityFromCustomerID = func(err error) error {
		return fmt.Errorf("failed to get identity from customer id: %w", err)
	}
)

func (c *UserClient) IdentityFromCustomerID(ctx context.Context, customer string) (user.Identity, error) {

	token, err := c.LoginAdmin(ctx)
	if err != nil {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(err)
	}

	users, err := c.client.GetUsers(ctx, token, c.config.Realm(), gocloak.GetUsersParams{
		Q: gocloak.StringP(fmt.Sprintf("customer:%s", customer)),
	})
	if err != nil {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(err)
	}

	if users == nil || len(users) == 0 {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(errUserMatchingCustomerIDNotFound)
	}

	id, err := uuid.FromString(gocloak.PString(users[0].ID))
	if err != nil {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(err)
	}

	var attrs = *users[0].Attributes
	profile, ok := attrs["profile"]
	if ok == false {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(errProfileAttributeNotFound)
	}

	pid, err := uuid.FromString(profile[0])
	if err != nil {
		return user.Identity{}, errFailedGetIdentityFromCustomerID(err)
	}

	return user.Identity{
		ID:         id,
		PlayerID:   pid,
		CustomerID: customer,
	}, nil

}
