# UpdateBuyerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The customer ID to recognize the cloud marketplace buyer in your internal system. This may be used for uploading CSV files for Batch Metering Usage | [optional] 
**Description** | Pointer to **string** | The description of the buyer. If not provided, the description will not be updated. | [optional] 
**MetronomeCustomerId** | Pointer to **string** | The Metronome Customer ID of the buyer. If not provided, the Metronome Customer ID will not be updated. | [optional] 
**Name** | Pointer to **string** | The name of the buyer. If not provided, the name will not be updated. | [optional] 
**OrbCustomerId** | Pointer to **string** | The Orb Customer ID of the buyer. If not provided, the Orb Customer ID will not be updated. | [optional] 

## Methods

### NewUpdateBuyerParams

`func NewUpdateBuyerParams() *UpdateBuyerParams`

NewUpdateBuyerParams instantiates a new UpdateBuyerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBuyerParamsWithDefaults

`func NewUpdateBuyerParamsWithDefaults() *UpdateBuyerParams`

NewUpdateBuyerParamsWithDefaults instantiates a new UpdateBuyerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *UpdateBuyerParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpdateBuyerParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpdateBuyerParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UpdateBuyerParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBuyerParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBuyerParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBuyerParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBuyerParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetronomeCustomerId

`func (o *UpdateBuyerParams) GetMetronomeCustomerId() string`

GetMetronomeCustomerId returns the MetronomeCustomerId field if non-nil, zero value otherwise.

### GetMetronomeCustomerIdOk

`func (o *UpdateBuyerParams) GetMetronomeCustomerIdOk() (*string, bool)`

GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeCustomerId

`func (o *UpdateBuyerParams) SetMetronomeCustomerId(v string)`

SetMetronomeCustomerId sets MetronomeCustomerId field to given value.

### HasMetronomeCustomerId

`func (o *UpdateBuyerParams) HasMetronomeCustomerId() bool`

HasMetronomeCustomerId returns a boolean if a field has been set.

### GetName

`func (o *UpdateBuyerParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBuyerParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBuyerParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBuyerParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrbCustomerId

`func (o *UpdateBuyerParams) GetOrbCustomerId() string`

GetOrbCustomerId returns the OrbCustomerId field if non-nil, zero value otherwise.

### GetOrbCustomerIdOk

`func (o *UpdateBuyerParams) GetOrbCustomerIdOk() (*string, bool)`

GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbCustomerId

`func (o *UpdateBuyerParams) SetOrbCustomerId(v string)`

SetOrbCustomerId sets OrbCustomerId field to given value.

### HasOrbCustomerId

`func (o *UpdateBuyerParams) HasOrbCustomerId() bool`

HasOrbCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


