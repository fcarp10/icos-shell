package cli

import (
	"context"
	"fmt"
	"log"
	"os"
	openapi "shellclient/pkg/openapi"

	"gopkg.in/yaml.v3"
)

func CreateDeployment(yamlFile []byte) (result string) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	// token := viper.GetString("ICOS_TOKEN") ## TBD
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
