# TypesUsageAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedUsageQuantity** | Pointer to **int32** | The total quantity allocated to this bucket of usage.  This member is required. | [optional] 
**Tags** | Pointer to [**[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag.md) | The set of tags that define the bucket of usage. For the bucket of items with no tags, this parameter can be left out. | [optional] 

## Methods

### NewTypesUsageAllocation

`func NewTypesUsageAllocation() *TypesUsageAllocation`

NewTypesUsageAllocation instantiates a new TypesUsageAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesUsageAllocationWithDefaults

`func NewTypesUsageAllocationWithDefaults() *TypesUsageAllocation`

NewTypesUsageAllocationWithDefaults instantiates a new TypesUsageAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedUsageQuantity

`func (o *TypesUsageAllocation) GetAllocatedUsageQuantity() int32`

GetAllocatedUsageQuantity returns the AllocatedUsageQuantity field if non-nil, zero value otherwise.

### GetAllocatedUsageQuantityOk

`func (o *TypesUsageAllocation) GetAllocatedUsageQuantityOk() (*int32, bool)`

GetAllocatedUsageQuantityOk returns a tuple with the AllocatedUsageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedUsageQuantity

`func (o *TypesUsageAllocation) SetAllocatedUsageQuantity(v int32)`

SetAllocatedUsageQuantity sets AllocatedUsageQuantity field to given value.

### HasAllocatedUsageQuantity

`func (o *TypesUsageAllocation) HasAllocatedUsageQuantity() bool`

HasAllocatedUsageQuantity returns a boolean if a field has been set.

### GetTags

`func (o *TypesUsageAllocation) GetTags() []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TypesUsageAllocation) GetTagsOk() (*[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TypesUsageAllocation) SetTags(v []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TypesUsageAllocation) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


