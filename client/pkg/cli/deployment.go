package cli

import (
	"context"
	"fmt"
	"os"
	openapi "shellclient/pkg/openapi"
)

func CreateDeployment(file string) (result string) {
	body := make(map[string]interface{})
	resp, err := openapi.Client.DeploymentAPI.CreateDeployment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", resp.Body)
		return "Error when creating deployment"
	} else {
		if resp.StatusCode == 201 {
			return "Deployment successfully created!"
		} else if resp.StatusCode == 202 {
			return "Deployment already exists"
		} else {
			return "Wrong status code received"
		}
	}
}
