package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	"github.com/justjack1521/mevconn"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/domain/auth"
	"mevway/internal/domain/user"
)

type TokenClient struct {
	client *gocloak.GoCloak
	config mevconn.KeyCloakConfig
}

func NewTokenClient(client *gocloak.GoCloak, config mevconn.KeyCloakConfig) *TokenClient {
	return &TokenClient{client: client, config: config}
}

func (c *TokenClient) CreateToken(ctx context.Context, target user.User) (auth.LoginResult, error) {

	tkn, err := c.client.Login(ctx, c.config.ClientID(), c.config.ClientSecret(), c.config.Realm(), target.Username, target.Password)
	if err != nil {
		return auth.LoginResult{}, err
	}

	return auth.LoginResult{
		IDToken:      tkn.IDToken,
		AccessToken:  tkn.AccessToken,
		RefreshToken: tkn.RefreshToken,
		ExpiresIn:    tkn.ExpiresIn,
	}, nil

}

func (c *TokenClient) VerifyToken(ctx context.Context, token string) (auth.TokenClaims, error) {

	claims := jwt.MapClaims{}

	_, err := c.client.DecodeAccessTokenCustomClaims(ctx, token, c.config.Realm(), claims)

	if err != nil {
		return auth.TokenClaims{}, errTokenExtractionFailed(err)
	}

	aud, ok := claims["sid"]
	if ok == false {
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

	profile, ok := claims["profile"]
	if ok == false {
		return auth.TokenClaims{}, errTokenExtractionFailed(err)
	}

	var roles = make([]string, 0)

	scope := claims["roles"]
	if scope != nil {
		slc, ok := scope.([]any)
		if ok {
			for _, s := range slc {
				str, ok := s.(string)
				if !ok {
					continue
				}
				roles = append(roles, str)
			}
		}
	}

	return auth.TokenClaims{
		SessionID:   uuid.FromStringOrNil(fmt.Sprintf("%v", aud)),
		UserID:      usr,
		PlayerID:    uuid.FromStringOrNil(fmt.Sprintf("%v", profile)),
		Environment: "development",
		Roles:       roles,
	}, nil

}

func (c *TokenClient) AuthoriseToken(ctx context.Context, token string) error {

	result, err := c.client.RetrospectToken(ctx, token, c.config.ClientID(), c.config.ClientSecret(), c.config.Realm())
	if err != nil {
		return errTokenAuthoriseFailed(err)
	}

	if *result.Active == false {
		return errTokenAuthoriseFailed(errTokenInactive)
	}

	return nil

}
