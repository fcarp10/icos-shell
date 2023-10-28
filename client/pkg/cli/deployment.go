package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func CreateDeployment(yamlFile []byte) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.CreateDeployment(context.Background()).ApiKey(token).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		if resp.StatusCode == 201 {
			fmt.Fprintln(os.Stderr, "Deployment successfully added!")
		} else if resp.StatusCode == 202 {
			fmt.Fprintln(os.Stderr, "Deployment already exists")
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}

func GetDeployment() {
	deployments, resp, err := openapi.Client.DeploymentAPI.GetDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		if resp.StatusCode == 200 {
			deployments_json, _ := json.Marshal(deployments)
			fmt.Fprintln(os.Stdout, string(deployments_json))
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No deployments found")
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}

func GetDeploymentById(id int64) {
	deployment, resp, err := openapi.Client.DeploymentAPI.GetDeploymentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Fprintln(os.Stdout, deployment)
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "Deployment not found")
		} else {
			fmt.Fprintln(os.Stderr, "Wrong status code received")
		}
	}
}
