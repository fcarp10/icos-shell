package cli

import (
	"context"
	"fmt"
	"os"
	openapi "shellclient/pkg/openapi"
)

func GetHealthcheck() {
	resp, err := openapi.Client.DefaultAPI.GetHealthcheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		if resp.StatusCode == 200 {
			fmt.Fprintln(os.Stderr, "healthcheck was successful")
		} else {
			fmt.Fprintln(os.Stderr, "healthcheck failed")
		}
	}
}
