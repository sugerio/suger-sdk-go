# AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingTerm** | Pointer to [**AzureMarketplaceTerm**](AzureMarketplaceTerm.md) |  | [optional] 
**IncludedQuantities** | Pointer to [**[]AzureMarketplacePriceAndAvailabilityCustomMeterPriceIncludedQuantityItem**](AzureMarketplacePriceAndAvailabilityCustomMeterPriceIncludedQuantityItem.md) |  | [optional] 
**PaymentOption** | Pointer to [**AzureMarketplaceTerm**](AzureMarketplaceTerm.md) |  | [optional] 
**PricePerPaymentInUsd** | Pointer to **float32** |  | [optional] 
**Prices** | Pointer to [**[]AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemPriceItem**](AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemPriceItem.md) |  | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem

`func NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem() *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem`

NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem instantiates a new AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemWithDefaults() *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem`

NewAzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetBillingTerm() AzureMarketplaceTerm`

GetBillingTerm returns the BillingTerm field if non-nil, zero value otherwise.

### GetBillingTermOk

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetBillingTermOk() (*AzureMarketplaceTerm, bool)`

GetBillingTermOk returns a tuple with the BillingTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) SetBillingTerm(v AzureMarketplaceTerm)`

SetBillingTerm sets BillingTerm field to given value.

### HasBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) HasBillingTerm() bool`

HasBillingTerm returns a boolean if a field has been set.

### GetIncludedQuantities

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetIncludedQuantities() []AzureMarketplacePriceAndAvailabilityCustomMeterPriceIncludedQuantityItem`

GetIncludedQuantities returns the IncludedQuantities field if non-nil, zero value otherwise.

### GetIncludedQuantitiesOk

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetIncludedQuantitiesOk() (*[]AzureMarketplacePriceAndAvailabilityCustomMeterPriceIncludedQuantityItem, bool)`

GetIncludedQuantitiesOk returns a tuple with the IncludedQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedQuantities

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) SetIncludedQuantities(v []AzureMarketplacePriceAndAvailabilityCustomMeterPriceIncludedQuantityItem)`

SetIncludedQuantities sets IncludedQuantities field to given value.

### HasIncludedQuantities

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) HasIncludedQuantities() bool`

HasIncludedQuantities returns a boolean if a field has been set.

### GetPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPaymentOption() AzureMarketplaceTerm`

GetPaymentOption returns the PaymentOption field if non-nil, zero value otherwise.

### GetPaymentOptionOk

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPaymentOptionOk() (*AzureMarketplaceTerm, bool)`

GetPaymentOptionOk returns a tuple with the PaymentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) SetPaymentOption(v AzureMarketplaceTerm)`

SetPaymentOption sets PaymentOption field to given value.

### HasPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) HasPaymentOption() bool`

HasPaymentOption returns a boolean if a field has been set.

### GetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPricePerPaymentInUsd() float32`

GetPricePerPaymentInUsd returns the PricePerPaymentInUsd field if non-nil, zero value otherwise.

### GetPricePerPaymentInUsdOk

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPricePerPaymentInUsdOk() (*float32, bool)`

GetPricePerPaymentInUsdOk returns a tuple with the PricePerPaymentInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) SetPricePerPaymentInUsd(v float32)`

SetPricePerPaymentInUsd sets PricePerPaymentInUsd field to given value.

### HasPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) HasPricePerPaymentInUsd() bool`

HasPricePerPaymentInUsd returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPrices() []AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemPriceItem`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) GetPricesOk() (*[]AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemPriceItem, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) SetPrices(v []AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItemPriceItem)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePriceAndAvailabilityCustomMeterPriceMeterItem) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


