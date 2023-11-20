package cli

import (
	"context"
	"encoding/json"
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
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 201 {
			fmt.Fprintln(os.Stderr, "Deployment successfully added!")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else if resp.StatusCode == 202 {
			fmt.Fprintln(os.Stderr, "Deployment already exists")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}

func UpdateDeployment(id int64, yamlFile []byte) {
	body := make(map[string]interface{})
	err := yaml.Unmarshal(yamlFile, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling YAML: %v", err)
	}
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.UpdateDeployment(context.Background(), id).ApiKey(token).Body(body).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 201 {
			fmt.Fprintln(os.Stderr, "Deployment successfully updated!")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}

func GetDeployment() {
	token := viper.GetString("auth_token")
	deployments, resp, err := openapi.Client.DeploymentAPI.GetDeployments(context.Background()).ApiKey(token).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
			for _, deployment := range deployments {
				prettyJSON, err := json.MarshalIndent(deployments, "", "  ")
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
					return
				} else {
					fmt.Fprintln(os.Stdout, string(prettyJSON))
				}
				fmt.Fprintln(os.Stdout, deployment)
			}
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No deployments found")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}

func GetDeploymentById(id int64) {
	token := viper.GetString("auth_token")
	deployment, resp, err := openapi.Client.DeploymentAPI.GetDeploymentById(context.Background(), id).ApiKey(token).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
			deployments_json, _ := json.Marshal(deployment)
			fmt.Fprintln(os.Stdout, string(deployments_json))
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "Deployment not found")
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}

func DeleteDeployment(id int64) {
	token := viper.GetString("auth_token")
	resp, err := openapi.Client.DeploymentAPI.DeleteDeploymentById(context.Background(), id).ApiKey(token).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
			fmt.Fprintln(os.Stderr, "Deployment successfully deleted")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else if resp.StatusCode == 204 {
			fmt.Fprintln(os.Stderr, "No matching deployments found")
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected status code received: ", resp.StatusCode)
		}
	}
}
