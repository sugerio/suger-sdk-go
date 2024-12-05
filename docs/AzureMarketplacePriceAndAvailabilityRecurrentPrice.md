# AzureMarketplacePriceAndAvailabilityRecurrentPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceInputOption** | Pointer to **string** | default \&quot;usd\&quot; | [optional] 
**Prices** | Pointer to [**[]AzureMarketplacePriceAndAvailabilityRecurrentPriceItem**](AzureMarketplacePriceAndAvailabilityRecurrentPriceItem.md) |  | [optional] 
**RecurrentPriceMode** | Pointer to **string** | default \&quot;flatRate\&quot; | [optional] 
**UserLimits** | Pointer to [**AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit**](AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit.md) |  | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityRecurrentPrice

`func NewAzureMarketplacePriceAndAvailabilityRecurrentPrice() *AzureMarketplacePriceAndAvailabilityRecurrentPrice`

NewAzureMarketplacePriceAndAvailabilityRecurrentPrice instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityRecurrentPriceWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceWithDefaults() *AzureMarketplacePriceAndAvailabilityRecurrentPrice`

NewAzureMarketplacePriceAndAvailabilityRecurrentPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceInputOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPriceInputOption() string`

GetPriceInputOption returns the PriceInputOption field if non-nil, zero value otherwise.

### GetPriceInputOptionOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPriceInputOptionOk() (*string, bool)`

GetPriceInputOptionOk returns a tuple with the PriceInputOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInputOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetPriceInputOption(v string)`

SetPriceInputOption sets PriceInputOption field to given value.

### HasPriceInputOption

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasPriceInputOption() bool`

HasPriceInputOption returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPrices() []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPricesOk() (*[]AzureMarketplacePriceAndAvailabilityRecurrentPriceItem, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetPrices(v []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetRecurrentPriceMode

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetRecurrentPriceMode() string`

GetRecurrentPriceMode returns the RecurrentPriceMode field if non-nil, zero value otherwise.

### GetRecurrentPriceModeOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetRecurrentPriceModeOk() (*string, bool)`

GetRecurrentPriceModeOk returns a tuple with the RecurrentPriceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrentPriceMode

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetRecurrentPriceMode(v string)`

SetRecurrentPriceMode sets RecurrentPriceMode field to given value.

### HasRecurrentPriceMode

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasRecurrentPriceMode() bool`

HasRecurrentPriceMode returns a boolean if a field has been set.

### GetUserLimits

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetUserLimits() AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit`

GetUserLimits returns the UserLimits field if non-nil, zero value otherwise.

### GetUserLimitsOk

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetUserLimitsOk() (*AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit, bool)`

GetUserLimitsOk returns a tuple with the UserLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLimits

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetUserLimits(v AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit)`

SetUserLimits sets UserLimits field to given value.

### HasUserLimits

`func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasUserLimits() bool`

HasUserLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


