# AzureMarketplacePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | ISO 4217 currency code | [optional] 
**Markets** | Pointer to **[]string** |  | [optional] 
**Price** | Pointer to **float32** | default 0 | [optional] 
**Prices** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAzureMarketplacePrice

`func NewAzureMarketplacePrice() *AzureMarketplacePrice`

NewAzureMarketplacePrice instantiates a new AzureMarketplacePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceWithDefaults

`func NewAzureMarketplacePriceWithDefaults() *AzureMarketplacePrice`

NewAzureMarketplacePriceWithDefaults instantiates a new AzureMarketplacePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *AzureMarketplacePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AzureMarketplacePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AzureMarketplacePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AzureMarketplacePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMarkets

`func (o *AzureMarketplacePrice) GetMarkets() []string`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *AzureMarketplacePrice) GetMarketsOk() (*[]string, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *AzureMarketplacePrice) SetMarkets(v []string)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *AzureMarketplacePrice) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetPrice

`func (o *AzureMarketplacePrice) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AzureMarketplacePrice) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AzureMarketplacePrice) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AzureMarketplacePrice) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPrices

`func (o *AzureMarketplacePrice) GetPrices() map[string]interface{}`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AzureMarketplacePrice) GetPricesOk() (*map[string]interface{}, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AzureMarketplacePrice) SetPrices(v map[string]interface{})`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AzureMarketplacePrice) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


