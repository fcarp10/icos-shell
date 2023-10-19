package shellbackend

import (
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"
	"github.com/spf13/viper"
)

// Functions to handle keycloak connections

// receiveAndValidateAccessToken - Implementation of Keycloak authentication with Token
//     - take the refresh token from the apiKey and check if it is valid, if yes create a new access token check its validity, output the access token
//     - apiKey: the refresh token
//     - returns: ImplResponse, error

func receiveAndValidateAccessToken(ctx context.Context, apiKey string) (ImplResponse, string, error) {
	// Implementation of Keycloak authentication with Token
	client := gocloak.NewClient(viper.GetString("keycloak.server"))

	// read refresh token from apiKey, which is just the access_token part of the token
	refresh_token := apiKey

	if refresh_token == "" {
		return Response(405, nil), "", errors.New("refresh token is empty")
	}
	// cut of the special characters from the token
	refresh_token = refresh_token[1 : len(refresh_token)-1]
	rptResult, err := client.RetrospectToken(ctx, refresh_token, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
	if err != nil {
		return Response(405, nil), "", errors.New("inspection of refresh token failed")
	}

	if !*rptResult.Active {
		return Response(405, nil), "", errors.New("refresh token is not active")
	} else {

		// New functioality to add the access token
		access_token, err := client.RefreshToken(ctx, refresh_token, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
		if err != nil {
			return Response(400, nil), "", errors.New("wrong or expired refresh token")
		} else {
			access_token_rptResult, err := client.RetrospectToken(ctx, access_token.AccessToken, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
			if err != nil {
				return Response(405, nil), "", errors.New("inspection of access token failed")
			}
			if !*access_token_rptResult.Active {
				return Response(405, nil), "", errors.New("access token is not active")
			} else {
				return Response(201, "Keycloak check successful"), access_token.AccessToken, nil
			}
		}
	}
	// Permissions could be added here
	// permissions := rptResult.Permissions
}
