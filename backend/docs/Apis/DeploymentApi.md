# DeploymentApi

All URIs are relative to *http://localhost:8080/api/v3*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createDeployment**](DeploymentApi.md#createDeployment) | **POST** /deployment/ | Creates a new deployment |
| [**deleteDeploymentById**](DeploymentApi.md#deleteDeploymentById) | **DELETE** /deployment/{deploymentId} | Deletes a deployment |
| [**getDeploymentById**](DeploymentApi.md#getDeploymentById) | **GET** /deployment/{deploymentId} | Find deployment by ID |
| [**getDeployments**](DeploymentApi.md#getDeployments) | **GET** /deployment/ | Returns a list of deployments |
| [**updateDeployment**](DeploymentApi.md#updateDeployment) | **PUT** /deployment/{deploymentId} | Updates a deployment |


<a name="createDeployment"></a>
# **createDeployment**
> createDeployment(body, api\_key)

Creates a new deployment

    

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | **Object**|  | |
| **api\_key** | **String**|  | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

<a name="deleteDeploymentById"></a>
# **deleteDeploymentById**
> deleteDeploymentById(deploymentId, api\_key)

Deletes a deployment

    Deletes a deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment that needs to be deleted | [default to null] |
| **api\_key** | **String**|  | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getDeploymentById"></a>
# **getDeploymentById**
> Map getDeploymentById(deploymentId, api\_key)

Find deployment by ID

    Returns a single deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment to return | [default to null] |
| **api\_key** | **String**|  | [optional] [default to null] |

### Return type

[**Map**](../Models/AnyType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDeployments"></a>
# **getDeployments**
> Map getDeployments(api\_key)

Returns a list of deployments

    Returns the list of deployments

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

<a name="updateDeployment"></a>
# **updateDeployment**
> updateDeployment(deploymentId, body, api\_key)

Updates a deployment

    Updates a deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment that needs to be updated | [default to null] |
| **body** | **Object**|  | |
| **api\_key** | **String**|  | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

