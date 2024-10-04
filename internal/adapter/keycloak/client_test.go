package keycloak_test

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/justjack1521/mevconn"
)

type config struct {
	hostname      string
	clientID      string
	clientSecret  string
	realm         string
	adminUsername string
	adminPassword string
}

func (c config) Hostname() string {
	return c.hostname
}

func (c config) ClientID() string {
	return c.clientID
}

func (c config) ClientSecret() string {
	return c.clientSecret
}

func (c config) Realm() string {
	return c.realm
}

func (c config) AdminCredentials() (username string, password string) {
	return c.adminUsername, c.adminPassword
}

func NewConfig() mevconn.KeyCloakConfig {
	return config{
		hostname:      "https://mevkey-hqezz.ondigitalocean.app",
		clientID:      "mevius-client",
		clientSecret:  "gLiXvjPLwN4d1ToKrxm8JioA6C9GNF5f",
		realm:         "mevius",
		adminUsername: "admin",
		adminPassword: "admin",
	}
}

func NewClient(config mevconn.KeyCloakConfig) *gocloak.GoCloak {
	return gocloak.NewClient(config.Hostname())
}
