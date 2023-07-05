package main

import (
	"context"
	"log"

	openapi "shellclient/openapi"
)

var client *openapi.APIClient

func main() {
	cfg := openapi.NewConfiguration()
	client = openapi.NewAPIClient(cfg)

	httpRes, err := client.DefaultApi.GetHealthcheck(context.Background()).Execute()
	if err != nil {
		log.Fatal("Error while connecting to server")
		log.Fatal(err)
	}
	if httpRes.StatusCode != 200 {
		log.Printf("Wrong status")
	} else {
		log.Printf("OK!")
	}
}
