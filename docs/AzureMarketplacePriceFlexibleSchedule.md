# AzureMarketplacePriceFlexibleSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingSchedule** | Pointer to [**[]AzureMarketplacePriceBillingSchedule**](AzureMarketplacePriceBillingSchedule.md) |  | [optional] 
**InitialCharge** | Pointer to [**AzureMarketplacePriceInitialCharge**](AzureMarketplacePriceInitialCharge.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 

## Methods

### NewAzureMarketplacePriceFlexibleSchedule

`func NewAzureMarketplacePriceFlexibleSchedule() *AzureMarketplacePriceFlexibleSchedule`

NewAzureMarketplacePriceFlexibleSchedule instantiates a new AzureMarketplacePriceFlexibleSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceFlexibleScheduleWithDefaults

`func NewAzureMarketplacePriceFlexibleScheduleWithDefaults() *AzureMarketplacePriceFlexibleSchedule`

NewAzureMarketplacePriceFlexibleScheduleWithDefaults instantiates a new AzureMarketplacePriceFlexibleSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingSchedule

`func (o *AzureMarketplacePriceFlexibleSchedule) GetBillingSchedule() []AzureMarketplacePriceBillingSchedule`

GetBillingSchedule returns the BillingSchedule field if non-nil, zero value otherwise.

### GetBillingScheduleOk

`func (o *AzureMarketplacePriceFlexibleSchedule) GetBillingScheduleOk() (*[]AzureMarketplacePriceBillingSchedule, bool)`

GetBillingScheduleOk returns a tuple with the BillingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSchedule

`func (o *AzureMarketplacePriceFlexibleSchedule) SetBillingSchedule(v []AzureMarketplacePriceBillingSchedule)`

SetBillingSchedule sets BillingSchedule field to given value.

### HasBillingSchedule

`func (o *AzureMarketplacePriceFlexibleSchedule) HasBillingSchedule() bool`

HasBillingSchedule returns a boolean if a field has been set.

### GetInitialCharge

`func (o *AzureMarketplacePriceFlexibleSchedule) GetInitialCharge() AzureMarketplacePriceInitialCharge`

GetInitialCharge returns the InitialCharge field if non-nil, zero value otherwise.

### GetInitialChargeOk

`func (o *AzureMarketplacePriceFlexibleSchedule) GetInitialChargeOk() (*AzureMarketplacePriceInitialCharge, bool)`

GetInitialChargeOk returns a tuple with the InitialCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCharge

`func (o *AzureMarketplacePriceFlexibleSchedule) SetInitialCharge(v AzureMarketplacePriceInitialCharge)`

SetInitialCharge sets InitialCharge field to given value.

### HasInitialCharge

`func (o *AzureMarketplacePriceFlexibleSchedule) HasInitialCharge() bool`

HasInitialCharge returns a boolean if a field has been set.

### GetPrice

`func (o *AzureMarketplacePriceFlexibleSchedule) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AzureMarketplacePriceFlexibleSchedule) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AzureMarketplacePriceFlexibleSchedule) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AzureMarketplacePriceFlexibleSchedule) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


