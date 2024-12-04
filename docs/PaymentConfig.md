# PaymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedWalletTypes** | Pointer to [**[]BillingWalletType**](BillingWalletType.md) | Allowed wallet types for this buyer, include payment methods from payment provider such as card, us_bank_account and credit. | [optional] 
**Currency** | Pointer to **string** | Currency used for billing. | [optional] 
**DefaultWalletId** | Pointer to **string** | Default wallet id which is a stripe payment method used to invoice. | [optional] 

## Methods

### NewPaymentConfig

`func NewPaymentConfig() *PaymentConfig`

NewPaymentConfig instantiates a new PaymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentConfigWithDefaults

`func NewPaymentConfigWithDefaults() *PaymentConfig`

NewPaymentConfigWithDefaults instantiates a new PaymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedWalletTypes

`func (o *PaymentConfig) GetAllowedWalletTypes() []BillingWalletType`

GetAllowedWalletTypes returns the AllowedWalletTypes field if non-nil, zero value otherwise.

### GetAllowedWalletTypesOk

`func (o *PaymentConfig) GetAllowedWalletTypesOk() (*[]BillingWalletType, bool)`

GetAllowedWalletTypesOk returns a tuple with the AllowedWalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedWalletTypes

`func (o *PaymentConfig) SetAllowedWalletTypes(v []BillingWalletType)`

SetAllowedWalletTypes sets AllowedWalletTypes field to given value.

### HasAllowedWalletTypes

`func (o *PaymentConfig) HasAllowedWalletTypes() bool`

HasAllowedWalletTypes returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentConfig) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentConfig) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentConfig) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentConfig) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultWalletId

`func (o *PaymentConfig) GetDefaultWalletId() string`

GetDefaultWalletId returns the DefaultWalletId field if non-nil, zero value otherwise.

### GetDefaultWalletIdOk

`func (o *PaymentConfig) GetDefaultWalletIdOk() (*string, bool)`

GetDefaultWalletIdOk returns a tuple with the DefaultWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWalletId

`func (o *PaymentConfig) SetDefaultWalletId(v string)`

SetDefaultWalletId sets DefaultWalletId field to given value.

### HasDefaultWalletId

`func (o *PaymentConfig) HasDefaultWalletId() bool`

HasDefaultWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


