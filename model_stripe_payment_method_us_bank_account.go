/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suger

import (
	"encoding/json"
)

// checks if the StripePaymentMethodUSBankAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripePaymentMethodUSBankAccount{}

// StripePaymentMethodUSBankAccount struct for StripePaymentMethodUSBankAccount
type StripePaymentMethodUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType *string `json:"account_holder_type,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `json:"account_type,omitempty"`
	// The name of the bank.
	BankName *string `json:"bank_name,omitempty"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Last four digits of the bank account number.
	Last4 *string `json:"last4,omitempty"`
	// Routing number of the bank account.
	RoutingNumber *string `json:"routing_number,omitempty"`
}

// NewStripePaymentMethodUSBankAccount instantiates a new StripePaymentMethodUSBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentMethodUSBankAccount() *StripePaymentMethodUSBankAccount {
	this := StripePaymentMethodUSBankAccount{}
	return &this
}

// NewStripePaymentMethodUSBankAccountWithDefaults instantiates a new StripePaymentMethodUSBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentMethodUSBankAccountWithDefaults() *StripePaymentMethodUSBankAccount {
	this := StripePaymentMethodUSBankAccount{}
	return &this
}

// GetAccountHolderType returns the AccountHolderType field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetAccountHolderType() string {
	if o == nil || IsNil(o.AccountHolderType) {
		var ret string
		return ret
	}
	return *o.AccountHolderType
}

// GetAccountHolderTypeOk returns a tuple with the AccountHolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetAccountHolderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountHolderType) {
		return nil, false
	}
	return o.AccountHolderType, true
}

// HasAccountHolderType returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasAccountHolderType() bool {
	if o != nil && !IsNil(o.AccountHolderType) {
		return true
	}

	return false
}

// SetAccountHolderType gets a reference to the given string and assigns it to the AccountHolderType field.
func (o *StripePaymentMethodUSBankAccount) SetAccountHolderType(v string) {
	o.AccountHolderType = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *StripePaymentMethodUSBankAccount) SetAccountType(v string) {
	o.AccountType = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *StripePaymentMethodUSBankAccount) SetBankName(v string) {
	o.BankName = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *StripePaymentMethodUSBankAccount) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetLast4() string {
	if o == nil || IsNil(o.Last4) {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetLast4Ok() (*string, bool) {
	if o == nil || IsNil(o.Last4) {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasLast4() bool {
	if o != nil && !IsNil(o.Last4) {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *StripePaymentMethodUSBankAccount) SetLast4(v string) {
	o.Last4 = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *StripePaymentMethodUSBankAccount) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethodUSBankAccount) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *StripePaymentMethodUSBankAccount) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *StripePaymentMethodUSBankAccount) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

func (o StripePaymentMethodUSBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripePaymentMethodUSBankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountHolderType) {
		toSerialize["account_holder_type"] = o.AccountHolderType
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.Last4) {
		toSerialize["last4"] = o.Last4
	}
	if !IsNil(o.RoutingNumber) {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	return toSerialize, nil
}

type NullableStripePaymentMethodUSBankAccount struct {
	value *StripePaymentMethodUSBankAccount
	isSet bool
}

func (v NullableStripePaymentMethodUSBankAccount) Get() *StripePaymentMethodUSBankAccount {
	return v.value
}

func (v *NullableStripePaymentMethodUSBankAccount) Set(val *StripePaymentMethodUSBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentMethodUSBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentMethodUSBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentMethodUSBankAccount(val *StripePaymentMethodUSBankAccount) *NullableStripePaymentMethodUSBankAccount {
	return &NullableStripePaymentMethodUSBankAccount{value: val, isSet: true}
}

func (v NullableStripePaymentMethodUSBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentMethodUSBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
