# AzureVariantResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**VariantID** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureVariantResource

`func NewAzureVariantResource() *AzureVariantResource`

NewAzureVariantResource instantiates a new AzureVariantResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVariantResourceWithDefaults

`func NewAzureVariantResourceWithDefaults() *AzureVariantResource`

NewAzureVariantResourceWithDefaults instantiates a new AzureVariantResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *AzureVariantResource) GetResources() []AzureTypeValue`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AzureVariantResource) GetResourcesOk() (*[]AzureTypeValue, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AzureVariantResource) SetResources(v []AzureTypeValue)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AzureVariantResource) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetVariantID

`func (o *AzureVariantResource) GetVariantID() string`

GetVariantID returns the VariantID field if non-nil, zero value otherwise.

### GetVariantIDOk

`func (o *AzureVariantResource) GetVariantIDOk() (*string, bool)`

GetVariantIDOk returns a tuple with the VariantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantID

`func (o *AzureVariantResource) SetVariantID(v string)`

SetVariantID sets VariantID field to given value.

### HasVariantID

`func (o *AzureVariantResource) HasVariantID() bool`

HasVariantID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


