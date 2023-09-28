package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/tubskns/icos-shell/client/pkg/openapi"
)

func LoginUser() (result string) {
	value, resp, err := openapi.Client.UserApi.LoginUser(context.Background()).Username(viper.GetString("username")).Password(viper.GetString("password")).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", resp.Body)
		fmt.Println("Error while trying to login")
		return ""
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Token received: ", value)
			fmt.Println("User successfully logged in!")
			return value
		} else {
			fmt.Println("Wrong status code received")
			return ""
		}
	}
}
