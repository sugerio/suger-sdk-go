/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StripePaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripePaymentMethod{}

// StripePaymentMethod struct for StripePaymentMethod
type StripePaymentMethod struct {
	BacsDebit *StripePaymentMethodBACSDebit `json:"bacs_debit,omitempty"`
	Card      *StripePaymentMethodCard      `json:"card,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created *int32 `json:"created,omitempty"`
	// Unique identifier for the payment method on stripe.
	Id *string `json:"id,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode *bool `json:"livemode,omitempty"`
	// String representing the object’s type. Always has the value `payment_method`.
	Object        *string                           `json:"object,omitempty"`
	SepaDebit     *StripePaymentMethodSEPADebit     `json:"sepa_debit,omitempty"`
	UsBankAccount *StripePaymentMethodUSBankAccount `json:"us_bank_account,omitempty"`
}

// NewStripePaymentMethod instantiates a new StripePaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentMethod() *StripePaymentMethod {
	this := StripePaymentMethod{}
	return &this
}

// NewStripePaymentMethodWithDefaults instantiates a new StripePaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentMethodWithDefaults() *StripePaymentMethod {
	this := StripePaymentMethod{}
	return &this
}

// GetBacsDebit returns the BacsDebit field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetBacsDebit() StripePaymentMethodBACSDebit {
	if o == nil || IsNil(o.BacsDebit) {
		var ret StripePaymentMethodBACSDebit
		return ret
	}
	return *o.BacsDebit
}

// GetBacsDebitOk returns a tuple with the BacsDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetBacsDebitOk() (*StripePaymentMethodBACSDebit, bool) {
	if o == nil || IsNil(o.BacsDebit) {
		return nil, false
	}
	return o.BacsDebit, true
}

// HasBacsDebit returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasBacsDebit() bool {
	if o != nil && !IsNil(o.BacsDebit) {
		return true
	}

	return false
}

// SetBacsDebit gets a reference to the given StripePaymentMethodBACSDebit and assigns it to the BacsDebit field.
func (o *StripePaymentMethod) SetBacsDebit(v StripePaymentMethodBACSDebit) {
	o.BacsDebit = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetCard() StripePaymentMethodCard {
	if o == nil || IsNil(o.Card) {
		var ret StripePaymentMethodCard
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetCardOk() (*StripePaymentMethodCard, bool) {
	if o == nil || IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasCard() bool {
	if o != nil && !IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given StripePaymentMethodCard and assigns it to the Card field.
func (o *StripePaymentMethod) SetCard(v StripePaymentMethodCard) {
	o.Card = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetCreated() int32 {
	if o == nil || IsNil(o.Created) {
		var ret int32
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int32 and assigns it to the Created field.
func (o *StripePaymentMethod) SetCreated(v int32) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StripePaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *StripePaymentMethod) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *StripePaymentMethod) SetObject(v string) {
	o.Object = &v
}

// GetSepaDebit returns the SepaDebit field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetSepaDebit() StripePaymentMethodSEPADebit {
	if o == nil || IsNil(o.SepaDebit) {
		var ret StripePaymentMethodSEPADebit
		return ret
	}
	return *o.SepaDebit
}

// GetSepaDebitOk returns a tuple with the SepaDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetSepaDebitOk() (*StripePaymentMethodSEPADebit, bool) {
	if o == nil || IsNil(o.SepaDebit) {
		return nil, false
	}
	return o.SepaDebit, true
}

// HasSepaDebit returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasSepaDebit() bool {
	if o != nil && !IsNil(o.SepaDebit) {
		return true
	}

	return false
}

// SetSepaDebit gets a reference to the given StripePaymentMethodSEPADebit and assigns it to the SepaDebit field.
func (o *StripePaymentMethod) SetSepaDebit(v StripePaymentMethodSEPADebit) {
	o.SepaDebit = &v
}

// GetUsBankAccount returns the UsBankAccount field value if set, zero value otherwise.
func (o *StripePaymentMethod) GetUsBankAccount() StripePaymentMethodUSBankAccount {
	if o == nil || IsNil(o.UsBankAccount) {
		var ret StripePaymentMethodUSBankAccount
		return ret
	}
	return *o.UsBankAccount
}

// GetUsBankAccountOk returns a tuple with the UsBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripePaymentMethod) GetUsBankAccountOk() (*StripePaymentMethodUSBankAccount, bool) {
	if o == nil || IsNil(o.UsBankAccount) {
		return nil, false
	}
	return o.UsBankAccount, true
}

// HasUsBankAccount returns a boolean if a field has been set.
func (o *StripePaymentMethod) HasUsBankAccount() bool {
	if o != nil && !IsNil(o.UsBankAccount) {
		return true
	}

	return false
}

// SetUsBankAccount gets a reference to the given StripePaymentMethodUSBankAccount and assigns it to the UsBankAccount field.
func (o *StripePaymentMethod) SetUsBankAccount(v StripePaymentMethodUSBankAccount) {
	o.UsBankAccount = &v
}

func (o StripePaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripePaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BacsDebit) {
		toSerialize["bacs_debit"] = o.BacsDebit
	}
	if !IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.SepaDebit) {
		toSerialize["sepa_debit"] = o.SepaDebit
	}
	if !IsNil(o.UsBankAccount) {
		toSerialize["us_bank_account"] = o.UsBankAccount
	}
	return toSerialize, nil
}

type NullableStripePaymentMethod struct {
	value *StripePaymentMethod
	isSet bool
}

func (v NullableStripePaymentMethod) Get() *StripePaymentMethod {
	return v.value
}

func (v *NullableStripePaymentMethod) Set(val *StripePaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentMethod(val *StripePaymentMethod) *NullableStripePaymentMethod {
	return &NullableStripePaymentMethod{value: val, isSet: true}
}

func (v NullableStripePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}