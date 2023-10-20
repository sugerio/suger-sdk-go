# AzurePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | ISO currency code, Three characters | [optional] 
**OpenPrice** | Pointer to **float32** |  | [optional] 
**PriceTierID** | Pointer to **string** |  | [optional] 

## Methods

### NewAzurePrice

`func NewAzurePrice() *AzurePrice`

NewAzurePrice instantiates a new AzurePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePriceWithDefaults

`func NewAzurePriceWithDefaults() *AzurePrice`

NewAzurePriceWithDefaults instantiates a new AzurePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *AzurePrice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AzurePrice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AzurePrice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AzurePrice) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetOpenPrice

`func (o *AzurePrice) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *AzurePrice) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *AzurePrice) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *AzurePrice) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetPriceTierID

`func (o *AzurePrice) GetPriceTierID() string`

GetPriceTierID returns the PriceTierID field if non-nil, zero value otherwise.

### GetPriceTierIDOk

`func (o *AzurePrice) GetPriceTierIDOk() (*string, bool)`

GetPriceTierIDOk returns a tuple with the PriceTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTierID

`func (o *AzurePrice) SetPriceTierID(v string)`

SetPriceTierID sets PriceTierID field to given value.

### HasPriceTierID

`func (o *AzurePrice) HasPriceTierID() bool`

HasPriceTierID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


