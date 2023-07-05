# CreateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Requirements** | **[]string** |  | 

## Methods

### NewCreateDeploymentRequest

`func NewCreateDeploymentRequest(application string, requirements []string, ) *CreateDeploymentRequest`

NewCreateDeploymentRequest instantiates a new CreateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeploymentRequestWithDefaults

`func NewCreateDeploymentRequestWithDefaults() *CreateDeploymentRequest`

NewCreateDeploymentRequestWithDefaults instantiates a new CreateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CreateDeploymentRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CreateDeploymentRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CreateDeploymentRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetRequirements

`func (o *CreateDeploymentRequest) GetRequirements() []string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CreateDeploymentRequest) GetRequirementsOk() (*[]string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CreateDeploymentRequest) SetRequirements(v []string)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


