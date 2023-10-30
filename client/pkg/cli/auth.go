package cli

import (
	"context"
	"fmt"
	"os"

	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func LoginUser() {
	token, resp, err := openapi.Client.UserAPI.LoginUser(context.Background()).Username(viper.GetString("keycloak.user")).Password(viper.GetString("keycloak.pass")).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
			if token != "" {
				fmt.Printf("%s%s%s", "ICOS_AUTH_TOKEN='", token, "'")
			} else {
				fmt.Fprintln(os.Stderr, "The token received is empty")
			}
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}
