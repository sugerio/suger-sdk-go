# StripePaymentMethodSEPADebit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankCode** | Pointer to **string** | Bank code of bank associated with the bank account. | [optional] 
**BranchCode** | Pointer to **string** | Branch code of bank associated with the bank account. | [optional] 
**Country** | Pointer to **string** | Two-letter ISO code representing the country the bank account is located in. | [optional] 
**Fingerprint** | Pointer to **string** | Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same. | [optional] 
**Last4** | Pointer to **string** | Last four characters of the IBAN. | [optional] 

## Methods

### NewStripePaymentMethodSEPADebit

`func NewStripePaymentMethodSEPADebit() *StripePaymentMethodSEPADebit`

NewStripePaymentMethodSEPADebit instantiates a new StripePaymentMethodSEPADebit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentMethodSEPADebitWithDefaults

`func NewStripePaymentMethodSEPADebitWithDefaults() *StripePaymentMethodSEPADebit`

NewStripePaymentMethodSEPADebitWithDefaults instantiates a new StripePaymentMethodSEPADebit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankCode

`func (o *StripePaymentMethodSEPADebit) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *StripePaymentMethodSEPADebit) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *StripePaymentMethodSEPADebit) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *StripePaymentMethodSEPADebit) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBranchCode

`func (o *StripePaymentMethodSEPADebit) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *StripePaymentMethodSEPADebit) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *StripePaymentMethodSEPADebit) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *StripePaymentMethodSEPADebit) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetCountry

`func (o *StripePaymentMethodSEPADebit) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StripePaymentMethodSEPADebit) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StripePaymentMethodSEPADebit) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StripePaymentMethodSEPADebit) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFingerprint

`func (o *StripePaymentMethodSEPADebit) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *StripePaymentMethodSEPADebit) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *StripePaymentMethodSEPADebit) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *StripePaymentMethodSEPADebit) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLast4

`func (o *StripePaymentMethodSEPADebit) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StripePaymentMethodSEPADebit) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StripePaymentMethodSEPADebit) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StripePaymentMethodSEPADebit) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


