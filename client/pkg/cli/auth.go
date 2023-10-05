package cli

import (
	"context"
	"fmt"
	"os"

	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func LoginUser() (result string) {
	value, resp, err := openapi.Client.UserAPI.LoginUser(context.Background()).Username(viper.GetString("username")).Password(viper.GetString("password")).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", resp.Body)
		fmt.Println("Error while trying to login")
		return ""
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Token received, user successfully logged in!")
			return value
		} else {
			fmt.Println("Wrong status code received")
			return ""
		}
	}
}
