# AzureMarketplacePriceInitialCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** |  | [optional] 
**PricePerPaymentInUsd** | Pointer to **float32** | For non-VM offers (SaaS, etc.) could be 0 for the amount | [optional] 
**Prices** | Pointer to [**[]AzureMarketplacePrice**](AzureMarketplacePrice.md) |  | [optional] 
**UnitPricePerPaymentPeriodInUsd** | Pointer to **float32** | For VM offers with flexible billing could be 0 for the amount | [optional] 

## Methods

### NewAzureMarketplacePriceInitialCharge

`func NewAzureMarketplacePriceInitialCharge() *AzureMarketplacePriceInitialCharge`

NewAzureMarketplacePriceInitialCharge instantiates a new AzureMarketplacePriceInitialCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceInitialChargeWithDefaults

`func NewAzureMarketplacePriceInitialChargeWithDefaults() *AzureMarketplacePriceInitialCharge`

NewAzureMarketplacePriceInitialChargeWithDefaults instantiates a new AzureMarketplacePriceInitialCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *AzureMarketplacePriceInitialCharge) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AzureMarketplacePriceInitialCharge) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AzureMarketplacePriceInitialCharge) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AzureMarketplacePriceInitialCharge) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceInitialCharge) GetPricePerPaymentInUsd() float32`

GetPricePerPaymentInUsd returns the PricePerPaymentInUsd field if non-nil, zero value otherwise.

### GetPricePerPaymentInUsdOk

`func (o *AzureMarketplacePriceInitialCharge) GetPricePerPaymentInUsdOk() (*float32, bool)`

GetPricePerPaymentInUsdOk returns a tuple with the PricePerPaymentInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceInitialCharge) SetPricePerPaymentInUsd(v float32)`

SetPricePerPaymentInUsd sets PricePerPaymentInUsd field to given value.

### HasPricePerPaymentInUsd

`func (o *AzureMarketplacePriceInitialCharge) HasPricePerPaymentInUsd() bool`

HasPricePerPaymentInUsd returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePriceInitialCharge) GetPrices() []AzureMarketplacePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePriceInitialCharge) GetPricesOk() (*[]AzureMarketplacePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePriceInitialCharge) SetPrices(v []AzureMarketplacePrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePriceInitialCharge) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceInitialCharge) GetUnitPricePerPaymentPeriodInUsd() float32`

GetUnitPricePerPaymentPeriodInUsd returns the UnitPricePerPaymentPeriodInUsd field if non-nil, zero value otherwise.

### GetUnitPricePerPaymentPeriodInUsdOk

`func (o *AzureMarketplacePriceInitialCharge) GetUnitPricePerPaymentPeriodInUsdOk() (*float32, bool)`

GetUnitPricePerPaymentPeriodInUsdOk returns a tuple with the UnitPricePerPaymentPeriodInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceInitialCharge) SetUnitPricePerPaymentPeriodInUsd(v float32)`

SetUnitPricePerPaymentPeriodInUsd sets UnitPricePerPaymentPeriodInUsd field to given value.

### HasUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceInitialCharge) HasUnitPricePerPaymentPeriodInUsd() bool`

HasUnitPricePerPaymentPeriodInUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


