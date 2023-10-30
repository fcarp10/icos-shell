package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	openapi "shellclient/pkg/openapi"
	"strconv"

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

func UpdateDeployment(id int64, yamlFile []byte) (result string) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling YAML: %v", err)
	}
	// token := viper.GetString("auth_token") ## TBD
	resp, err := openapi.Client.DeploymentAPI.UpdateDeployment(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		//		fmt.Fprintf(os.Stderr, "Received response: %v\n", resp.Body)
		return "Error when updating deployment"
	} else {
		if resp.StatusCode == 201 {
			return "Deployment successfully updated!"
		} else {
			return "Unexpected status code received: " + strconv.Itoa(resp.StatusCode)
		}
	}
}

func GetDeployment() (result string) {

	deployments, resp, err := openapi.Client.DeploymentAPI.GetDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving deployments: " + strconv.Itoa(resp.StatusCode)
	} else {
		if resp.StatusCode == 200 {
			deployments_json, _ := json.Marshal(deployments)
			fmt.Fprintln(os.Stdout, string(deployments_json))
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No deployments found")
		} else {
			return "Unexpected status code received: " + strconv.Itoa(resp.StatusCode)
		}
	}
	return
}

func GetDeploymentById(id int64) (result string) {

	deployment, resp, err := openapi.Client.DeploymentAPI.GetDeploymentById(context.Background(), id).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving deployment: " + strconv.Itoa(resp.StatusCode)
	} else {
		if resp.StatusCode == 200 {
			fmt.Fprintln(os.Stdout, deployment)
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "Deployment not found")
		} else {
			return "Unexpected status code received: " + strconv.Itoa(resp.StatusCode)
		}
	}
	return
}

func DeleteDeployment(id int64) (result string) {

	resp, err := openapi.Client.DeploymentAPI.DeleteDeploymentById(context.Background(), id).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when deleting deployments"
	} else {
		if resp.StatusCode == 200 {
			return "Deployment successfully deleted"
		} else if resp.StatusCode == 204 {
			return "No matching deployments found"
		} else {
			return "Unexpected status code received: " + strconv.Itoa(resp.StatusCode)
		}
	}
}
