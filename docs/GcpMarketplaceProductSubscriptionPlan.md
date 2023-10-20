# GcpMarketplaceProductSubscriptionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** | such as \&quot;ONE_YEAR\&quot;, \&quot;TWO_YEAR\&quot;, \&quot;THREE_YEAR\&quot; | [optional] 
**Price** | Pointer to [**GcpPriceValue**](GcpPriceValue.md) |  | [optional] 

## Methods

### NewGcpMarketplaceProductSubscriptionPlan

`func NewGcpMarketplaceProductSubscriptionPlan() *GcpMarketplaceProductSubscriptionPlan`

NewGcpMarketplaceProductSubscriptionPlan instantiates a new GcpMarketplaceProductSubscriptionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductSubscriptionPlanWithDefaults

`func NewGcpMarketplaceProductSubscriptionPlanWithDefaults() *GcpMarketplaceProductSubscriptionPlan`

NewGcpMarketplaceProductSubscriptionPlanWithDefaults instantiates a new GcpMarketplaceProductSubscriptionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *GcpMarketplaceProductSubscriptionPlan) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GcpMarketplaceProductSubscriptionPlan) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GcpMarketplaceProductSubscriptionPlan) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GcpMarketplaceProductSubscriptionPlan) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrice

`func (o *GcpMarketplaceProductSubscriptionPlan) GetPrice() GcpPriceValue`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GcpMarketplaceProductSubscriptionPlan) GetPriceOk() (*GcpPriceValue, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GcpMarketplaceProductSubscriptionPlan) SetPrice(v GcpPriceValue)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GcpMarketplaceProductSubscriptionPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


