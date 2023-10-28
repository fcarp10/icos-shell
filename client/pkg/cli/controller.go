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
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		if resp.StatusCode == 201 {
			fmt.Fprintln(os.Stderr, "Controller successfully added!")
		} else if resp.StatusCode == 202 {
			fmt.Fprintln(os.Stderr, "Controller already exists")
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}

func GetController() {
	controllers, resp, err := openapi.Client.ControllerAPI.GetControllers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		if resp.StatusCode == 200 {
			controllers_json, _ := json.Marshal(controllers)
			fmt.Fprintln(os.Stdout, string(controllers_json))
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No controllers found")
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}
