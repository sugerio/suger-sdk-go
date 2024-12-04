# AwsMarketplaceMeteringUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIdentifier** | Pointer to **string** | The CustomerIdentifier is obtained through the ResolveCustomer operation and represents an individual buyer in your application. | [optional] 
**Dimension** | Pointer to **string** | During the process of registering a product on AWS Marketplace, dimensions are specified. These represent different units of value in your application. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of usage consumed by the customer for the given dimension and time. Defaults to 0 if not specified. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp, in UTC, for which the usage is being reported. Your application can meter usage for up to one hour in the past. Make sure the timestamp value is not before the start of the software usage. | [optional] 
**UsageAllocations** | Pointer to [**[]AwsMarketplaceMeteringUsageAllocation**](AwsMarketplaceMeteringUsageAllocation.md) | The set of UsageAllocations to submit. The sum of all UsageAllocation quantities must equal the Quantity of the UsageRecord. | [optional] 

## Methods

### NewAwsMarketplaceMeteringUsageRecord

`func NewAwsMarketplaceMeteringUsageRecord() *AwsMarketplaceMeteringUsageRecord`

NewAwsMarketplaceMeteringUsageRecord instantiates a new AwsMarketplaceMeteringUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceMeteringUsageRecordWithDefaults

`func NewAwsMarketplaceMeteringUsageRecordWithDefaults() *AwsMarketplaceMeteringUsageRecord`

NewAwsMarketplaceMeteringUsageRecordWithDefaults instantiates a new AwsMarketplaceMeteringUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIdentifier

`func (o *AwsMarketplaceMeteringUsageRecord) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *AwsMarketplaceMeteringUsageRecord) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *AwsMarketplaceMeteringUsageRecord) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *AwsMarketplaceMeteringUsageRecord) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetDimension

`func (o *AwsMarketplaceMeteringUsageRecord) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *AwsMarketplaceMeteringUsageRecord) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *AwsMarketplaceMeteringUsageRecord) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *AwsMarketplaceMeteringUsageRecord) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetQuantity

`func (o *AwsMarketplaceMeteringUsageRecord) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AwsMarketplaceMeteringUsageRecord) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AwsMarketplaceMeteringUsageRecord) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AwsMarketplaceMeteringUsageRecord) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTimestamp

`func (o *AwsMarketplaceMeteringUsageRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AwsMarketplaceMeteringUsageRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AwsMarketplaceMeteringUsageRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AwsMarketplaceMeteringUsageRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsageAllocations

`func (o *AwsMarketplaceMeteringUsageRecord) GetUsageAllocations() []AwsMarketplaceMeteringUsageAllocation`

GetUsageAllocations returns the UsageAllocations field if non-nil, zero value otherwise.

### GetUsageAllocationsOk

`func (o *AwsMarketplaceMeteringUsageRecord) GetUsageAllocationsOk() (*[]AwsMarketplaceMeteringUsageAllocation, bool)`

GetUsageAllocationsOk returns a tuple with the UsageAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAllocations

`func (o *AwsMarketplaceMeteringUsageRecord) SetUsageAllocations(v []AwsMarketplaceMeteringUsageAllocation)`

SetUsageAllocations sets UsageAllocations field to given value.

### HasUsageAllocations

`func (o *AwsMarketplaceMeteringUsageRecord) HasUsageAllocations() bool`

HasUsageAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


