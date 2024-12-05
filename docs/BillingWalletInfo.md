# BillingWalletInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseDate** | Pointer to **time.Time** | The close date of the wallet if applicable. | [optional] 
**StripePaymentMethod** | Pointer to [**StripePaymentMethod**](StripePaymentMethod.md) |  | [optional] 
**StripeSetupIntentId** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingWalletInfo

`func NewBillingWalletInfo() *BillingWalletInfo`

NewBillingWalletInfo instantiates a new BillingWalletInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWalletInfoWithDefaults

`func NewBillingWalletInfoWithDefaults() *BillingWalletInfo`

NewBillingWalletInfoWithDefaults instantiates a new BillingWalletInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *BillingWalletInfo) GetCloseDate() time.Time`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *BillingWalletInfo) GetCloseDateOk() (*time.Time, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *BillingWalletInfo) SetCloseDate(v time.Time)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *BillingWalletInfo) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetStripePaymentMethod

`func (o *BillingWalletInfo) GetStripePaymentMethod() StripePaymentMethod`

GetStripePaymentMethod returns the StripePaymentMethod field if non-nil, zero value otherwise.

### GetStripePaymentMethodOk

`func (o *BillingWalletInfo) GetStripePaymentMethodOk() (*StripePaymentMethod, bool)`

GetStripePaymentMethodOk returns a tuple with the StripePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePaymentMethod

`func (o *BillingWalletInfo) SetStripePaymentMethod(v StripePaymentMethod)`

SetStripePaymentMethod sets StripePaymentMethod field to given value.

### HasStripePaymentMethod

`func (o *BillingWalletInfo) HasStripePaymentMethod() bool`

HasStripePaymentMethod returns a boolean if a field has been set.

### GetStripeSetupIntentId

`func (o *BillingWalletInfo) GetStripeSetupIntentId() string`

GetStripeSetupIntentId returns the StripeSetupIntentId field if non-nil, zero value otherwise.

### GetStripeSetupIntentIdOk

`func (o *BillingWalletInfo) GetStripeSetupIntentIdOk() (*string, bool)`

GetStripeSetupIntentIdOk returns a tuple with the StripeSetupIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSetupIntentId

`func (o *BillingWalletInfo) SetStripeSetupIntentId(v string)`

SetStripeSetupIntentId sets StripeSetupIntentId field to given value.

### HasStripeSetupIntentId

`func (o *BillingWalletInfo) HasStripeSetupIntentId() bool`

HasStripeSetupIntentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


