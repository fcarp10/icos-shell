package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	openapi "shellclient/openapi"
)

var client *openapi.APIClient

func main() {
	host := "localhost:8080"
	option := flag.String("server", host, "URL of the shell-backend")
	flag.Parse()
	if *option != "" {
		fmt.Println("Trying to connect to", *option)
		host = *option
	}
	cfg := openapi.NewConfiguration()
	cfg.Host = host
	client = openapi.NewAPIClient(cfg)

	httpRes, err := client.DefaultApi.GetHealthcheck(context.Background()).Execute()
	if err != nil {
		log.Fatal("Error connecting to server")
		log.Fatal(err)
	}
	if httpRes.StatusCode != 200 {
		log.Printf("Server returned wrong status code")
	} else {
		log.Printf("Server connection successful!")
	}
}
