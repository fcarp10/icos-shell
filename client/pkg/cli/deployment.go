package cli

// "context"
// "fmt"
// "os"

// "github.com/tubskns/icos-shell/client/pkg/openapi"

func CreateDeployment(file string) (result string) {
	// deploymentRequest := *openapi.NewCreateDeploymentRequest(file)
	// resp, err := openapi.Client.DeploymentApi.CreateDeployment(context.Background()).CreateDeploymentRequest(deploymentRequest).Execute()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%v\n", err)
	// 	fmt.Fprintf(os.Stderr, "%v\n", resp.Body)
	// 	return "Error when creating deployment"
	// } else {
	// 	if resp.StatusCode == 201 {
	// 		return "Deployment successfully created!"
	// 	} else if resp.StatusCode == 202 {
	// 		return "Deployment already exists"
	// 	} else {
	// 		return "Wrong status code received"
	// 	}
	// }
	return ""
}
