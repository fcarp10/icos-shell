# ResourceApi

All URIs are relative to *http://localhost:8080/api/v3*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getResourceById**](ResourceApi.md#getResourceById) | **GET** /resource/{resourceId} | Find resource by ID |
| [**getResources**](ResourceApi.md#getResources) | **GET** /resource/ | Returns a list of resources |


<a name="getResourceById"></a>
# **getResourceById**
> Map getResourceById(resourceId)

Find resource by ID

    Returns a single resource

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **resourceId** | **Long**| ID of resource to return | [default to null] |

### Return type

[**Map**](../Models/AnyType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getResources"></a>
# **getResources**
> Map getResources(api\_key)

Returns a list of resources

    Returns a list of resources

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **api\_key** | **String**|  | [optional] [default to null] |

### Return type

[**Map**](../Models/AnyType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

