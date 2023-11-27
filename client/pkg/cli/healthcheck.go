package cli

import (
	"context"
	openapi "shellclient/pkg/openapi"
)

func GetHealthcheck() {
	resp, err := openapi.Client.DefaultAPI.GetHealthcheck(context.Background()).Execute()
	printResponseSimple(resp, err)
}
