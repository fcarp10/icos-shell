/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	// shellbackend "github.com/GIT_USER_ID/GIT_REPO_ID/go" // full path for debugging
	shellbackend "./go"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := shellbackend.NewDefaultApiService()
	DefaultApiController := shellbackend.NewDefaultApiController(DefaultApiService)

	DeploymentApiService := shellbackend.NewDeploymentApiService()
	DeploymentApiController := shellbackend.NewDeploymentApiController(DeploymentApiService)

	ResourceApiService := shellbackend.NewResourceApiService()
	ResourceApiController := shellbackend.NewResourceApiController(ResourceApiService)

	UserApiService := shellbackend.NewUserApiService()
	UserApiController := shellbackend.NewUserApiController(UserApiService)

	router := shellbackend.NewRouter(DefaultApiController, DeploymentApiController, ResourceApiController, UserApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
