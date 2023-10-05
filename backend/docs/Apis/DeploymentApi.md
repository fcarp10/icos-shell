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
> deleteDeploymentById(deploymentId)

Deletes a deployment

    Deletes a deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment that needs to be deleted | [default to null] |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="getDeploymentById"></a>
# **getDeploymentById**
> Deployment getDeploymentById(deploymentId)

Find deployment by ID

    Returns a single deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment to return | [default to null] |

### Return type

[**Deployment**](../Models/Deployment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDeployments"></a>
# **getDeployments**
> List getDeployments()

Returns a list of deployments

    Returns the list of deployments

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/Deployment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateDeployment"></a>
# **updateDeployment**
> updateDeployment(deploymentId, body)

Updates a deployment

    Updates a deployment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentId** | **Long**| ID of deployment that needs to be updated | [default to null] |
| **body** | **Object**|  | |

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

