package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func AddController(name string, address string) {
	controller := *openapi.NewController(name, address)
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.ControllerAPI.AddController(context.Background()).ApiKey(token).Controller(controller).Execute()
	printResponseSimple(resp, err)
}

func GetController() string {
	controllers, resp, err := openapi.Client.ControllerAPI.GetControllers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		if resp.StatusCode == 200 {
			controllers_json, _ := json.MarshalIndent(controllers, "", "  ")
			fmt.Fprintln(os.Stdout, string(controllers_json))
			return controllers[0].Address // for now, return the address of first controller (TBD)
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No controllers found")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
	return ""
}
