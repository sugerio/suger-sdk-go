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
	"time"
)

// checks if the BillingWalletInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingWalletInfo{}

// BillingWalletInfo struct for BillingWalletInfo
type BillingWalletInfo struct {
	// The close date of the wallet if applicable.
	CloseDate           *time.Time           `json:"closeDate,omitempty"`
	StripePaymentMethod *StripePaymentMethod `json:"stripePaymentMethod,omitempty"`
	StripeSetupIntentId *string              `json:"stripeSetupIntentId,omitempty"`
}

// NewBillingWalletInfo instantiates a new BillingWalletInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingWalletInfo() *BillingWalletInfo {
	this := BillingWalletInfo{}
	return &this
}

// NewBillingWalletInfoWithDefaults instantiates a new BillingWalletInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingWalletInfoWithDefaults() *BillingWalletInfo {
	this := BillingWalletInfo{}
	return &this
}

// GetCloseDate returns the CloseDate field value if set, zero value otherwise.
func (o *BillingWalletInfo) GetCloseDate() time.Time {
	if o == nil || IsNil(o.CloseDate) {
		var ret time.Time
		return ret
	}
	return *o.CloseDate
}

// GetCloseDateOk returns a tuple with the CloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWalletInfo) GetCloseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CloseDate) {
		return nil, false
	}
	return o.CloseDate, true
}

// HasCloseDate returns a boolean if a field has been set.
func (o *BillingWalletInfo) HasCloseDate() bool {
	if o != nil && !IsNil(o.CloseDate) {
		return true
	}

	return false
}

// SetCloseDate gets a reference to the given time.Time and assigns it to the CloseDate field.
func (o *BillingWalletInfo) SetCloseDate(v time.Time) {
	o.CloseDate = &v
}

// GetStripePaymentMethod returns the StripePaymentMethod field value if set, zero value otherwise.
func (o *BillingWalletInfo) GetStripePaymentMethod() StripePaymentMethod {
	if o == nil || IsNil(o.StripePaymentMethod) {
		var ret StripePaymentMethod
		return ret
	}
	return *o.StripePaymentMethod
}

// GetStripePaymentMethodOk returns a tuple with the StripePaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWalletInfo) GetStripePaymentMethodOk() (*StripePaymentMethod, bool) {
	if o == nil || IsNil(o.StripePaymentMethod) {
		return nil, false
	}
	return o.StripePaymentMethod, true
}

// HasStripePaymentMethod returns a boolean if a field has been set.
func (o *BillingWalletInfo) HasStripePaymentMethod() bool {
	if o != nil && !IsNil(o.StripePaymentMethod) {
		return true
	}

	return false
}

// SetStripePaymentMethod gets a reference to the given StripePaymentMethod and assigns it to the StripePaymentMethod field.
func (o *BillingWalletInfo) SetStripePaymentMethod(v StripePaymentMethod) {
	o.StripePaymentMethod = &v
}

// GetStripeSetupIntentId returns the StripeSetupIntentId field value if set, zero value otherwise.
func (o *BillingWalletInfo) GetStripeSetupIntentId() string {
	if o == nil || IsNil(o.StripeSetupIntentId) {
		var ret string
		return ret
	}
	return *o.StripeSetupIntentId
}

// GetStripeSetupIntentIdOk returns a tuple with the StripeSetupIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWalletInfo) GetStripeSetupIntentIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeSetupIntentId) {
		return nil, false
	}
	return o.StripeSetupIntentId, true
}

// HasStripeSetupIntentId returns a boolean if a field has been set.
func (o *BillingWalletInfo) HasStripeSetupIntentId() bool {
	if o != nil && !IsNil(o.StripeSetupIntentId) {
		return true
	}

	return false
}

// SetStripeSetupIntentId gets a reference to the given string and assigns it to the StripeSetupIntentId field.
func (o *BillingWalletInfo) SetStripeSetupIntentId(v string) {
	o.StripeSetupIntentId = &v
}

func (o BillingWalletInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingWalletInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloseDate) {
		toSerialize["closeDate"] = o.CloseDate
	}
	if !IsNil(o.StripePaymentMethod) {
		toSerialize["stripePaymentMethod"] = o.StripePaymentMethod
	}
	if !IsNil(o.StripeSetupIntentId) {
		toSerialize["stripeSetupIntentId"] = o.StripeSetupIntentId
	}
	return toSerialize, nil
}

type NullableBillingWalletInfo struct {
	value *BillingWalletInfo
	isSet bool
}

func (v NullableBillingWalletInfo) Get() *BillingWalletInfo {
	return v.value
}

func (v *NullableBillingWalletInfo) Set(val *BillingWalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingWalletInfo(val *BillingWalletInfo) *NullableBillingWalletInfo {
	return &NullableBillingWalletInfo{value: val, isSet: true}
}

func (v NullableBillingWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
