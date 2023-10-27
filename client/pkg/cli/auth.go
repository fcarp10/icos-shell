package cli

import (
	"context"
	"fmt"
	"os"

	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func LoginUser() (result string) {
	value, resp, err := openapi.Client.UserAPI.LoginUser(context.Background()).Username(viper.GetString("keycloak.user")).Password(viper.GetString("keycloak.pass")).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return ""
	} else {
		if resp.StatusCode == 200 {
			fmt.Fprintln(os.Stderr, "Token received, user successfully logged in!")
			return value
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
			return ""
		}
	}
}
