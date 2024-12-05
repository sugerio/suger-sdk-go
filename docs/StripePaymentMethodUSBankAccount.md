# StripePaymentMethodUSBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderType** | Pointer to **string** | Account holder type: individual or company. | [optional] 
**AccountType** | Pointer to **string** | Account type: checkings or savings. Defaults to checking if omitted. | [optional] 
**BankName** | Pointer to **string** | The name of the bank. | [optional] 
**Fingerprint** | Pointer to **string** | Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same. | [optional] 
**Last4** | Pointer to **string** | Last four digits of the bank account number. | [optional] 
**RoutingNumber** | Pointer to **string** | Routing number of the bank account. | [optional] 

## Methods

### NewStripePaymentMethodUSBankAccount

`func NewStripePaymentMethodUSBankAccount() *StripePaymentMethodUSBankAccount`

NewStripePaymentMethodUSBankAccount instantiates a new StripePaymentMethodUSBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentMethodUSBankAccountWithDefaults

`func NewStripePaymentMethodUSBankAccountWithDefaults() *StripePaymentMethodUSBankAccount`

NewStripePaymentMethodUSBankAccountWithDefaults instantiates a new StripePaymentMethodUSBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderType

`func (o *StripePaymentMethodUSBankAccount) GetAccountHolderType() string`

GetAccountHolderType returns the AccountHolderType field if non-nil, zero value otherwise.

### GetAccountHolderTypeOk

`func (o *StripePaymentMethodUSBankAccount) GetAccountHolderTypeOk() (*string, bool)`

GetAccountHolderTypeOk returns a tuple with the AccountHolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderType

`func (o *StripePaymentMethodUSBankAccount) SetAccountHolderType(v string)`

SetAccountHolderType sets AccountHolderType field to given value.

### HasAccountHolderType

`func (o *StripePaymentMethodUSBankAccount) HasAccountHolderType() bool`

HasAccountHolderType returns a boolean if a field has been set.

### GetAccountType

`func (o *StripePaymentMethodUSBankAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *StripePaymentMethodUSBankAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *StripePaymentMethodUSBankAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *StripePaymentMethodUSBankAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBankName

`func (o *StripePaymentMethodUSBankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *StripePaymentMethodUSBankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *StripePaymentMethodUSBankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *StripePaymentMethodUSBankAccount) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetFingerprint

`func (o *StripePaymentMethodUSBankAccount) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *StripePaymentMethodUSBankAccount) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *StripePaymentMethodUSBankAccount) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *StripePaymentMethodUSBankAccount) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLast4

`func (o *StripePaymentMethodUSBankAccount) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StripePaymentMethodUSBankAccount) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StripePaymentMethodUSBankAccount) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StripePaymentMethodUSBankAccount) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *StripePaymentMethodUSBankAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *StripePaymentMethodUSBankAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *StripePaymentMethodUSBankAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *StripePaymentMethodUSBankAccount) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


