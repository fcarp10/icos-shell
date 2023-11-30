package shellbackend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/spf13/viper"
)

// receiveAndValidateAccessToken - Check if refresh token is valid, if so, generate a new access token
//   - input: refresh token (apiKey)
//   - returns: access token
func receiveAndValidateAccessToken(ctx context.Context, apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("token is empty")
	}
	if string(apiKey[0]) == "\"" {
		apiKey = apiKey[1 : len(apiKey)-1] // remove special characters from token
	}
	client := gocloak.NewClient(viper.GetString("keycloak.server"))
	rptResult, err := client.RetrospectToken(ctx, apiKey, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
	if err != nil {
		return "", errors.New("inspection of token failed")
	}
	if !*rptResult.Active {
		return "", errors.New("token is not active")
	} else {
		access_token, err := client.RefreshToken(ctx, apiKey, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
		if err != nil {
			return "", errors.New("wrong or expired refresh token")
		} else {
			access_token_rptResult, err := client.RetrospectToken(ctx, access_token.AccessToken, viper.GetString("keycloak.client_id"), viper.GetString("keycloak.client_secret"), viper.GetString("keycloak.realm"))
			if err != nil {
				return "", errors.New("inspection of access token failed")
			}
			if !*access_token_rptResult.Active {
				return "", errors.New("access token is not active")
			} else {
				return access_token.AccessToken, nil
			}
		}
	}
}

func addBearerToToken(ctx context.Context, apiKey string, req *http.Request) *http.Request {
	accessToken, _ := receiveAndValidateAccessToken(ctx, apiKey)
	accessToken = "Bearer " + accessToken
	req.Header.Add("Authorization", accessToken)
	return req
}

func errorConnect(resp *http.Response, err error) (ImplResponse, error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	if resp != nil {
		return Response(500, resp.Body), nil
	}
	return Response(500, "Error while connecting to the component"), nil
}

func unexpectedCode(statusCode int) (ImplResponse, error) {
	return Response(statusCode, "Unexpected status code received"), nil
}

func readResponse(response *http.Response) []byte {
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil
	} else {
		// log.Println(response.StatusCode, string(body))
		return body
	}
}

func unmarshalArrayResponse(response *http.Response) []map[string]interface{} {
	var body = readResponse(response)
	var data []map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil
	}
	return data
}

func unmarshalResponse(response *http.Response) map[string]interface{} {
	var body = readResponse(response)
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil
	}
	return data
}
