package cli

import (
	"context"
	openapi "shellclient/pkg/openapi"

	"github.com/spf13/viper"
)

func GetResource() {
	token := viper.GetString("auth_token")
	resources, resp, err := openapi.Client.ResourceAPI.GetResources(context.Background()).ApiKey(token).Execute()
	printPrettyJSON(resources, resp, err)
}
