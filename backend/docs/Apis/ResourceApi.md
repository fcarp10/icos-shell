# ResourceApi

All URIs are relative to *http://localhost:8080/api/v3*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getResourceById**](ResourceApi.md#getResourceById) | **GET** /resource/{resourceId} | Find resource by ID |
| [**getResources**](ResourceApi.md#getResources) | **GET** /resource/ | Returns a list of resources |


<a name="getResourceById"></a>
# **getResourceById**
> Resource getResourceById(resourceId)

Find resource by ID

    Returns a single resource

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **resourceId** | **Long**| ID of resource to return | [default to null] |

### Return type

[**Resource**](../Models/Resource.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getResources"></a>
# **getResources**
> List getResources()

Returns a list of resources

    Returns a list of resources

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/Resource.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

