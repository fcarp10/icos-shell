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
	// token := viper.GetString("auth_token") ## TBD
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

func GetDeployment() (result string) {

	deployments, resp, err := openapi.Client.DeploymentAPI.GetDeployments(context.Background()).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving deployments"
	} else {
		if resp.StatusCode == 200 {
			for _, element := range deployments {
				fmt.Println(element)
			}
			return "Deployment list retrieved"
		} else if resp.StatusCode == 204 {
			return "No deployments found"
		} else {
			return "Wrong status code received"
		}
	}
}

func GetDeploymentById(id int64) (result string) {

	deployment, resp, err := openapi.Client.DeploymentAPI.GetDeploymentById(context.Background(), id).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving deployments"
	} else {
		if resp.StatusCode == 200 {
			fmt.Println(deployment)
			return "Deployment list retrieved"
		} else if resp.StatusCode == 204 {
			return "No deployments found"
		} else {
			return "Wrong status code received"
		}
	}
}
