package keycloak

import (
	"context"
	"github.com/Nerzal/gocloak"
	"github.com/justjack1521/mevconn"
	"mevway/internal/domain/auth"
)

type Client struct {
	gocloak      gocloak.GoCloak
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

	jwt, err := c.gocloak.Login(c.clientID, c.clientSecret, c.realm, request.Username, request.Password)
	if err != nil {
		return auth.LoginResponse{}, err
	}

	return auth.LoginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}, nil

}
