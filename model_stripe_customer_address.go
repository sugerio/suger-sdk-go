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

// checks if the StripeCustomerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeCustomerAddress{}

// StripeCustomerAddress struct for StripeCustomerAddress
type StripeCustomerAddress struct {
	// City, district, suburb, town, or village.
	City *string `json:"city,omitempty"`
	// Two-letter country code (ISO 3166-1 alpha-2)
	Country *string `json:"country,omitempty"`
	// Address line 1 (e.g., street, PO Box, or company name).
	Line1 *string `json:"line1,omitempty"`
	// Address line 2 (e.g., apartment, suite, unit, or building).
	Line2 *string `json:"line2,omitempty"`
	// ZIP or postal code.
	PostalCode *string `json:"postal_code,omitempty"`
	// State, county, province, or region.
	State *string `json:"state,omitempty"`
}

// NewStripeCustomerAddress instantiates a new StripeCustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeCustomerAddress() *StripeCustomerAddress {
	this := StripeCustomerAddress{}
	return &this
}

// NewStripeCustomerAddressWithDefaults instantiates a new StripeCustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeCustomerAddressWithDefaults() *StripeCustomerAddress {
	this := StripeCustomerAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *StripeCustomerAddress) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *StripeCustomerAddress) SetCountry(v string) {
	o.Country = &v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasLine1() bool {
	if o != nil && !IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *StripeCustomerAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasLine2() bool {
	if o != nil && !IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *StripeCustomerAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *StripeCustomerAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StripeCustomerAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeCustomerAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StripeCustomerAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StripeCustomerAddress) SetState(v string) {
	o.State = &v
}

func (o StripeCustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeCustomerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableStripeCustomerAddress struct {
	value *StripeCustomerAddress
	isSet bool
}

func (v NullableStripeCustomerAddress) Get() *StripeCustomerAddress {
	return v.value
}

func (v *NullableStripeCustomerAddress) Set(val *StripeCustomerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeCustomerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeCustomerAddress(val *StripeCustomerAddress) *NullableStripeCustomerAddress {
	return &NullableStripeCustomerAddress{value: val, isSet: true}
}

func (v NullableStripeCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
