# ControllerApi

All URIs are relative to *http://localhost:8080/api/v3*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addController**](ControllerApi.md#addController) | **POST** /controller/ | Adds a new controller |
| [**getControllers**](ControllerApi.md#getControllers) | **GET** /controller/ | Returns a list of controllers |


<a name="addController"></a>
# **addController**
> addController(username, password, Controller)

Adds a new controller

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**|  | [default to null] |
| **password** | **String**|  | [default to null] |
| **Controller** | [**Controller**](../Models/Controller.md)|  | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="getControllers"></a>
# **getControllers**
> List getControllers()

Returns a list of controllers

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/Controller.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

