# UsageAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** |  | [optional] 
**Tags** | Pointer to [**[]UsageAllocationTag**](UsageAllocationTag.md) |  | [optional] 

## Methods

### NewUsageAllocation

`func NewUsageAllocation() *UsageAllocation`

NewUsageAllocation instantiates a new UsageAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAllocationWithDefaults

`func NewUsageAllocationWithDefaults() *UsageAllocation`

NewUsageAllocationWithDefaults instantiates a new UsageAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UsageAllocation) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageAllocation) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageAllocation) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageAllocation) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTags

`func (o *UsageAllocation) GetTags() []UsageAllocationTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UsageAllocation) GetTagsOk() (*[]UsageAllocationTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UsageAllocation) SetTags(v []UsageAllocationTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UsageAllocation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


