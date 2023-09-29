# \DeploymentAPI

All URIs are relative to *http://localhost:8080/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeployment**](DeploymentAPI.md#CreateDeployment) | **Post** /deployment/ | Creates a new deployment
[**DeleteDeploymentById**](DeploymentAPI.md#DeleteDeploymentById) | **Delete** /deployment/{deploymentId} | Deletes a deployment
[**GetDeploymentById**](DeploymentAPI.md#GetDeploymentById) | **Get** /deployment/{deploymentId} | Find deployment by ID
[**GetDeployments**](DeploymentAPI.md#GetDeployments) | **Get** /deployment/ | Returns a list of deployments
[**UpdateDeployment**](DeploymentAPI.md#UpdateDeployment) | **Put** /deployment/{deploymentId} | Updates a deployment



## CreateDeployment

> CreateDeployment(ctx).Body(body).Execute()

Creates a new deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.CreateDeployment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.CreateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentById

> DeleteDeploymentById(ctx, deploymentId).Execute()

Deletes a deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    deploymentId := int64(789) // int64 | ID of deployment that needs to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DeleteDeploymentById(context.Background(), deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DeleteDeploymentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | ID of deployment that needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentById

> Deployment GetDeploymentById(ctx, deploymentId).Execute()

Find deployment by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    deploymentId := int64(789) // int64 | ID of deployment to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetDeploymentById(context.Background(), deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeploymentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentById`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeploymentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | ID of deployment to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deployment**](Deployment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployments

> []Deployment GetDeployments(ctx).Execute()

Returns a list of deployments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetDeployments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployments`: []Deployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsRequest struct via the builder pattern


### Return type

[**[]Deployment**](Deployment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeployment

> UpdateDeployment(ctx, deploymentId).Body(body).Execute()

Updates a deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    deploymentId := int64(789) // int64 | ID of deployment that needs to be updated
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.UpdateDeployment(context.Background(), deploymentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.UpdateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | ID of deployment that needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

