package cli

import (
	"context"
	"fmt"
	"os"
	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func CreateDeployment(yamlFile []byte) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling YAML: %v", err)
	}
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.CreateDeployment(context.Background()).ApiKey(token).Body(body).Execute()
	printResponseSimple(resp, err)
}

func UpdateDeployment(id int64, yamlFile []byte) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling YAML: %v", err)
	}
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.UpdateDeployment(context.Background(), id).ApiKey(token).Body(body).Execute()
	printResponseSimple(resp, err)
}

func GetDeployment() {
	token := viper.GetString("auth_token")
	deployments, resp, err := openapi.Client.DeploymentAPI.GetDeployments(context.Background()).ApiKey(token).Execute()
	printPrettyJSON(deployments, resp, err)
}

func GetDeploymentById(id int64) {
	token := viper.GetString("auth_token")
	deployment, resp, err := openapi.Client.DeploymentAPI.GetDeploymentById(context.Background(), id).ApiKey(token).Execute()
	printPrettyJSON(deployment, resp, err)
}

func DeleteDeployment(id int64) {
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.DeleteDeploymentById(context.Background(), id).ApiKey(token).Execute()
	printResponseSimple(resp, err)
}
