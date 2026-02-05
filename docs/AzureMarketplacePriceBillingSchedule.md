# AzureMarketplacePriceBillingSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeDate** | Pointer to **string** | In format of YYYY-MM-DD. | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PricePerPaymentInUsd** | Pointer to **float32** | For non-VM offers (SaaS, etc.) | [optional] 
**Prices** | Pointer to [**[]AzureMarketplacePrice**](AzureMarketplacePrice.md) |  | [optional] 
**UnitPricePerPaymentPeriodInUsd** | Pointer to **float32** | For VM offers with flexible billing could be 0 for the amount | [optional] 

## Methods

### NewAzureMarketplacePriceBillingSchedule

`func NewAzureMarketplacePriceBillingSchedule() *AzureMarketplacePriceBillingSchedule`

NewAzureMarketplacePriceBillingSchedule instantiates a new AzureMarketplacePriceBillingSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceBillingScheduleWithDefaults

`func NewAzureMarketplacePriceBillingScheduleWithDefaults() *AzureMarketplacePriceBillingSchedule`

NewAzureMarketplacePriceBillingScheduleWithDefaults instantiates a new AzureMarketplacePriceBillingSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeDate

`func (o *AzureMarketplacePriceBillingSchedule) GetChargeDate() string`

GetChargeDate returns the ChargeDate field if non-nil, zero value otherwise.

### GetChargeDateOk

`func (o *AzureMarketplacePriceBillingSchedule) GetChargeDateOk() (*string, bool)`

GetChargeDateOk returns a tuple with the ChargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDate

`func (o *AzureMarketplacePriceBillingSchedule) SetChargeDate(v string)`

SetChargeDate sets ChargeDate field to given value.

### HasChargeDate

`func (o *AzureMarketplacePriceBillingSchedule) HasChargeDate() bool`

HasChargeDate returns a boolean if a field has been set.

### GetNote

`func (o *AzureMarketplacePriceBillingSchedule) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AzureMarketplacePriceBillingSchedule) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AzureMarketplacePriceBillingSchedule) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AzureMarketplacePriceBillingSchedule) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceBillingSchedule) GetPricePerPaymentInUsd() float32`

GetPricePerPaymentInUsd returns the PricePerPaymentInUsd field if non-nil, zero value otherwise.

### GetPricePerPaymentInUsdOk

`func (o *AzureMarketplacePriceBillingSchedule) GetPricePerPaymentInUsdOk() (*float32, bool)`

GetPricePerPaymentInUsdOk returns a tuple with the PricePerPaymentInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceBillingSchedule) SetPricePerPaymentInUsd(v float32)`

SetPricePerPaymentInUsd sets PricePerPaymentInUsd field to given value.

### HasPricePerPaymentInUsd

`func (o *AzureMarketplacePriceBillingSchedule) HasPricePerPaymentInUsd() bool`

HasPricePerPaymentInUsd returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePriceBillingSchedule) GetPrices() []AzureMarketplacePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePriceBillingSchedule) GetPricesOk() (*[]AzureMarketplacePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePriceBillingSchedule) SetPrices(v []AzureMarketplacePrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePriceBillingSchedule) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceBillingSchedule) GetUnitPricePerPaymentPeriodInUsd() float32`

GetUnitPricePerPaymentPeriodInUsd returns the UnitPricePerPaymentPeriodInUsd field if non-nil, zero value otherwise.

### GetUnitPricePerPaymentPeriodInUsdOk

`func (o *AzureMarketplacePriceBillingSchedule) GetUnitPricePerPaymentPeriodInUsdOk() (*float32, bool)`

GetUnitPricePerPaymentPeriodInUsdOk returns a tuple with the UnitPricePerPaymentPeriodInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceBillingSchedule) SetUnitPricePerPaymentPeriodInUsd(v float32)`

SetUnitPricePerPaymentPeriodInUsd sets UnitPricePerPaymentPeriodInUsd field to given value.

### HasUnitPricePerPaymentPeriodInUsd

`func (o *AzureMarketplacePriceBillingSchedule) HasUnitPricePerPaymentPeriodInUsd() bool`

HasUnitPricePerPaymentPeriodInUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


