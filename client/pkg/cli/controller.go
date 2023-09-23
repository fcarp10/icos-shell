package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/tubskns/icos-shell/client/pkg/openapi"
)

func AddController(name string, address string) (result string) {
	controller := *openapi.NewController(name, address)
	r, err := openapi.Client.ControllerApi.AddController(context.Background()).Username(viper.GetString("username")).Password(viper.GetString("password")).Controller(controller).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", r.Body)
		return "Error when adding a controller"
	} else {
		return "Controller added successfully!"
	}
}

func GetController() (result string) {
	resp, _, err := openapi.Client.ControllerApi.GetControllers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving controllers"
	} else {
		for _, element := range resp {
			fmt.Println(element)
		}
		return "Controller list retrieved"
	}
}
