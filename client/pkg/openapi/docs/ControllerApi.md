# \ControllerApi

All URIs are relative to *http://localhost:8080/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddController**](ControllerApi.md#AddController) | **Post** /controller/ | Adds a new controller
[**GetControllers**](ControllerApi.md#GetControllers) | **Get** /controller/ | Returns a list of controllers



## AddController

> AddController(ctx).Username(username).Password(password).Controller(controller).Execute()

Adds a new controller

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
    username := "username_example" // string | 
    password := "password_example" // string | 
    controller := *openapiclient.NewController("controller_1", "192.168.100.1") // Controller | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControllerApi.AddController(context.Background()).Username(username).Password(password).Controller(controller).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.AddController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **controller** | [**Controller**](Controller.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControllers

> []Controller GetControllers(ctx).Execute()

Returns a list of controllers

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
    resp, r, err := apiClient.ControllerApi.GetControllers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControllerApi.GetControllers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllers`: []Controller
    fmt.Fprintf(os.Stdout, "Response from `ControllerApi.GetControllers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllersRequest struct via the builder pattern


### Return type

[**[]Controller**](Controller.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

