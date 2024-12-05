# StripeBalanceTransactionFeeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount of the fee, in cents. | [optional] 
**Application** | Pointer to **string** | ID of the Connect application that earned the fee. | [optional] 
**Description** | Pointer to **string** | An arbitrary string attached to the object. Often useful for displaying to users. | [optional] 
**Type** | Pointer to **string** | Type of the fee, one of: &#x60;application_fee&#x60;, &#x60;payment_method_passthrough_fee&#x60;, &#x60;stripe_fee&#x60; or &#x60;tax&#x60;. | [optional] 

## Methods

### NewStripeBalanceTransactionFeeDetail

`func NewStripeBalanceTransactionFeeDetail() *StripeBalanceTransactionFeeDetail`

NewStripeBalanceTransactionFeeDetail instantiates a new StripeBalanceTransactionFeeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeBalanceTransactionFeeDetailWithDefaults

`func NewStripeBalanceTransactionFeeDetailWithDefaults() *StripeBalanceTransactionFeeDetail`

NewStripeBalanceTransactionFeeDetailWithDefaults instantiates a new StripeBalanceTransactionFeeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StripeBalanceTransactionFeeDetail) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StripeBalanceTransactionFeeDetail) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StripeBalanceTransactionFeeDetail) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StripeBalanceTransactionFeeDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApplication

`func (o *StripeBalanceTransactionFeeDetail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *StripeBalanceTransactionFeeDetail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *StripeBalanceTransactionFeeDetail) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *StripeBalanceTransactionFeeDetail) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDescription

`func (o *StripeBalanceTransactionFeeDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StripeBalanceTransactionFeeDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StripeBalanceTransactionFeeDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StripeBalanceTransactionFeeDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *StripeBalanceTransactionFeeDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeBalanceTransactionFeeDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeBalanceTransactionFeeDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeBalanceTransactionFeeDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


