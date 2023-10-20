# GcpMarketplaceProductUsageFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayQuantity** | Pointer to **int32** | such as 1 | [optional] 
**MetricId** | Pointer to **string** | such as \&quot;Starter_storage\&quot; | [optional] 
**PriceTiers** | Pointer to [**[]GcpPriceTier**](GcpPriceTier.md) |  | [optional] 

## Methods

### NewGcpMarketplaceProductUsageFee

`func NewGcpMarketplaceProductUsageFee() *GcpMarketplaceProductUsageFee`

NewGcpMarketplaceProductUsageFee instantiates a new GcpMarketplaceProductUsageFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductUsageFeeWithDefaults

`func NewGcpMarketplaceProductUsageFeeWithDefaults() *GcpMarketplaceProductUsageFee`

NewGcpMarketplaceProductUsageFeeWithDefaults instantiates a new GcpMarketplaceProductUsageFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayQuantity

`func (o *GcpMarketplaceProductUsageFee) GetDisplayQuantity() int32`

GetDisplayQuantity returns the DisplayQuantity field if non-nil, zero value otherwise.

### GetDisplayQuantityOk

`func (o *GcpMarketplaceProductUsageFee) GetDisplayQuantityOk() (*int32, bool)`

GetDisplayQuantityOk returns a tuple with the DisplayQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayQuantity

`func (o *GcpMarketplaceProductUsageFee) SetDisplayQuantity(v int32)`

SetDisplayQuantity sets DisplayQuantity field to given value.

### HasDisplayQuantity

`func (o *GcpMarketplaceProductUsageFee) HasDisplayQuantity() bool`

HasDisplayQuantity returns a boolean if a field has been set.

### GetMetricId

`func (o *GcpMarketplaceProductUsageFee) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *GcpMarketplaceProductUsageFee) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *GcpMarketplaceProductUsageFee) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *GcpMarketplaceProductUsageFee) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetPriceTiers

`func (o *GcpMarketplaceProductUsageFee) GetPriceTiers() []GcpPriceTier`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *GcpMarketplaceProductUsageFee) GetPriceTiersOk() (*[]GcpPriceTier, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *GcpMarketplaceProductUsageFee) SetPriceTiers(v []GcpPriceTier)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *GcpMarketplaceProductUsageFee) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


