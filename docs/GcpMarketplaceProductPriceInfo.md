# GcpMarketplaceProductPriceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**PriceModel** | Pointer to [**GcpMarketplacePriceModel**](GcpMarketplacePriceModel.md) |  | [optional] 
**SubscriptionPlans** | Pointer to [**[]GcpMarketplaceProductSubscriptionPlan**](GcpMarketplaceProductSubscriptionPlan.md) | Subscription Plan (Flat Commitment) | [optional] 
**UsageFees** | Pointer to [**[]GcpMarketplaceProductUsageFee**](GcpMarketplaceProductUsageFee.md) | Usage Metering Dimension/Metric if available | [optional] 

## Methods

### NewGcpMarketplaceProductPriceInfo

`func NewGcpMarketplaceProductPriceInfo() *GcpMarketplaceProductPriceInfo`

NewGcpMarketplaceProductPriceInfo instantiates a new GcpMarketplaceProductPriceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductPriceInfoWithDefaults

`func NewGcpMarketplaceProductPriceInfoWithDefaults() *GcpMarketplaceProductPriceInfo`

NewGcpMarketplaceProductPriceInfoWithDefaults instantiates a new GcpMarketplaceProductPriceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GcpMarketplaceProductPriceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GcpMarketplaceProductPriceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GcpMarketplaceProductPriceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GcpMarketplaceProductPriceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriceModel

`func (o *GcpMarketplaceProductPriceInfo) GetPriceModel() GcpMarketplacePriceModel`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *GcpMarketplaceProductPriceInfo) GetPriceModelOk() (*GcpMarketplacePriceModel, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *GcpMarketplaceProductPriceInfo) SetPriceModel(v GcpMarketplacePriceModel)`

SetPriceModel sets PriceModel field to given value.

### HasPriceModel

`func (o *GcpMarketplaceProductPriceInfo) HasPriceModel() bool`

HasPriceModel returns a boolean if a field has been set.

### GetSubscriptionPlans

`func (o *GcpMarketplaceProductPriceInfo) GetSubscriptionPlans() []GcpMarketplaceProductSubscriptionPlan`

GetSubscriptionPlans returns the SubscriptionPlans field if non-nil, zero value otherwise.

### GetSubscriptionPlansOk

`func (o *GcpMarketplaceProductPriceInfo) GetSubscriptionPlansOk() (*[]GcpMarketplaceProductSubscriptionPlan, bool)`

GetSubscriptionPlansOk returns a tuple with the SubscriptionPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlans

`func (o *GcpMarketplaceProductPriceInfo) SetSubscriptionPlans(v []GcpMarketplaceProductSubscriptionPlan)`

SetSubscriptionPlans sets SubscriptionPlans field to given value.

### HasSubscriptionPlans

`func (o *GcpMarketplaceProductPriceInfo) HasSubscriptionPlans() bool`

HasSubscriptionPlans returns a boolean if a field has been set.

### GetUsageFees

`func (o *GcpMarketplaceProductPriceInfo) GetUsageFees() []GcpMarketplaceProductUsageFee`

GetUsageFees returns the UsageFees field if non-nil, zero value otherwise.

### GetUsageFeesOk

`func (o *GcpMarketplaceProductPriceInfo) GetUsageFeesOk() (*[]GcpMarketplaceProductUsageFee, bool)`

GetUsageFeesOk returns a tuple with the UsageFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageFees

`func (o *GcpMarketplaceProductPriceInfo) SetUsageFees(v []GcpMarketplaceProductUsageFee)`

SetUsageFees sets UsageFees field to given value.

### HasUsageFees

`func (o *GcpMarketplaceProductPriceInfo) HasUsageFees() bool`

HasUsageFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


