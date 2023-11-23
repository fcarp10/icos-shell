package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func GetResource() {
	token := viper.GetString("auth_token")
	resources, resp, err := openapi.Client.ResourceAPI.GetResources(context.Background()).ApiKey(token).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
			prettyResources, err := json.MarshalIndent(resources, "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
				return
			} else {
				fmt.Fprintln(os.Stdout, string(prettyResources))
			}
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No resources found")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}
