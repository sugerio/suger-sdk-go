# AzureMarketplacePriceAndAvailabilityRecurrentPriceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingTerm** | Pointer to [**AzureMarketplaceTerm**](AzureMarketplaceTerm.md) |  | [optional] 
**PaymentOption** | Pointer to [**AzureMarketplaceTerm**](AzureMarketplaceTerm.md) |  | [optional] 
**PricePerPaymentInUsd** | Pointer to **float32** |  | [optional] 
**Prices** | Pointer to [**[]AzureMarketplacePrice**](AzureMarketplacePrice.md) |  | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItem

`func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItem() *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem`

NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItem instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPriceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItemWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItemWithDefaults() *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem`

NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItemWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPriceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetBillingTerm() AzureMarketplaceTerm`

GetBillingTerm returns the BillingTerm field if non-nil, zero value otherwise.

### GetBillingTermOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetBillingTermOk() (*AzureMarketplaceTerm, bool)`

GetBillingTermOk returns a tuple with the BillingTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetBillingTerm(v AzureMarketplaceTerm)`

SetBillingTerm sets BillingTerm field to given value.

### HasBillingTerm

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasBillingTerm() bool`

HasBillingTerm returns a boolean if a field has been set.

### GetPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPaymentOption() AzureMarketplaceTerm`

GetPaymentOption returns the PaymentOption field if non-nil, zero value otherwise.

### GetPaymentOptionOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPaymentOptionOk() (*AzureMarketplaceTerm, bool)`

GetPaymentOptionOk returns a tuple with the PaymentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPaymentOption(v AzureMarketplaceTerm)`

SetPaymentOption sets PaymentOption field to given value.

### HasPaymentOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPaymentOption() bool`

HasPaymentOption returns a boolean if a field has been set.

### GetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricePerPaymentInUsd() float32`

GetPricePerPaymentInUsd returns the PricePerPaymentInUsd field if non-nil, zero value otherwise.

### GetPricePerPaymentInUsdOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricePerPaymentInUsdOk() (*float32, bool)`

GetPricePerPaymentInUsdOk returns a tuple with the PricePerPaymentInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPricePerPaymentInUsd(v float32)`

SetPricePerPaymentInUsd sets PricePerPaymentInUsd field to given value.

### HasPricePerPaymentInUsd

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPricePerPaymentInUsd() bool`

HasPricePerPaymentInUsd returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPrices() []AzureMarketplacePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricesOk() (*[]AzureMarketplacePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPrices(v []AzureMarketplacePrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


