# StripePaymentMethodBACSDebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same. | [optional] 
**Last4** | Pointer to **string** | Last four digits of the bank account number. | [optional] 
**SortCode** | Pointer to **string** | Sort code of the bank account. (e.g., &#x60;10-20-30&#x60;) | [optional] 

## Methods

### NewStripePaymentMethodBACSDebit

`func NewStripePaymentMethodBACSDebit() *StripePaymentMethodBACSDebit`

NewStripePaymentMethodBACSDebit instantiates a new StripePaymentMethodBACSDebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentMethodBACSDebitWithDefaults

`func NewStripePaymentMethodBACSDebitWithDefaults() *StripePaymentMethodBACSDebit`

NewStripePaymentMethodBACSDebitWithDefaults instantiates a new StripePaymentMethodBACSDebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *StripePaymentMethodBACSDebit) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *StripePaymentMethodBACSDebit) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *StripePaymentMethodBACSDebit) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *StripePaymentMethodBACSDebit) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLast4

`func (o *StripePaymentMethodBACSDebit) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StripePaymentMethodBACSDebit) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StripePaymentMethodBACSDebit) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StripePaymentMethodBACSDebit) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetSortCode

`func (o *StripePaymentMethodBACSDebit) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *StripePaymentMethodBACSDebit) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *StripePaymentMethodBACSDebit) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.

### HasSortCode

`func (o *StripePaymentMethodBACSDebit) HasSortCode() bool`

HasSortCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


