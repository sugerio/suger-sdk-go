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

// checks if the GcpPriceModelDiscountTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpPriceModelDiscountTemplate{}

// GcpPriceModelDiscountTemplate struct for GcpPriceModelDiscountTemplate
type GcpPriceModelDiscountTemplate struct {
	DiscountEconomics      *string              `json:"discountEconomics,omitempty"`
	DiscountPercentage     *GcpAmountConstraint `json:"discountPercentage,omitempty"`
	DiscountedPrice        *GcpPriceValue       `json:"discountedPrice,omitempty"`
	HideDiscountPercentage *bool                `json:"hideDiscountPercentage,omitempty"`
}

// NewGcpPriceModelDiscountTemplate instantiates a new GcpPriceModelDiscountTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpPriceModelDiscountTemplate() *GcpPriceModelDiscountTemplate {
	this := GcpPriceModelDiscountTemplate{}
	return &this
}

// NewGcpPriceModelDiscountTemplateWithDefaults instantiates a new GcpPriceModelDiscountTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpPriceModelDiscountTemplateWithDefaults() *GcpPriceModelDiscountTemplate {
	this := GcpPriceModelDiscountTemplate{}
	return &this
}

// GetDiscountEconomics returns the DiscountEconomics field value if set, zero value otherwise.
func (o *GcpPriceModelDiscountTemplate) GetDiscountEconomics() string {
	if o == nil || IsNil(o.DiscountEconomics) {
		var ret string
		return ret
	}
	return *o.DiscountEconomics
}

// GetDiscountEconomicsOk returns a tuple with the DiscountEconomics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceModelDiscountTemplate) GetDiscountEconomicsOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountEconomics) {
		return nil, false
	}
	return o.DiscountEconomics, true
}

// HasDiscountEconomics returns a boolean if a field has been set.
func (o *GcpPriceModelDiscountTemplate) HasDiscountEconomics() bool {
	if o != nil && !IsNil(o.DiscountEconomics) {
		return true
	}

	return false
}

// SetDiscountEconomics gets a reference to the given string and assigns it to the DiscountEconomics field.
func (o *GcpPriceModelDiscountTemplate) SetDiscountEconomics(v string) {
	o.DiscountEconomics = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *GcpPriceModelDiscountTemplate) GetDiscountPercentage() GcpAmountConstraint {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret GcpAmountConstraint
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceModelDiscountTemplate) GetDiscountPercentageOk() (*GcpAmountConstraint, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *GcpPriceModelDiscountTemplate) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given GcpAmountConstraint and assigns it to the DiscountPercentage field.
func (o *GcpPriceModelDiscountTemplate) SetDiscountPercentage(v GcpAmountConstraint) {
	o.DiscountPercentage = &v
}

// GetDiscountedPrice returns the DiscountedPrice field value if set, zero value otherwise.
func (o *GcpPriceModelDiscountTemplate) GetDiscountedPrice() GcpPriceValue {
	if o == nil || IsNil(o.DiscountedPrice) {
		var ret GcpPriceValue
		return ret
	}
	return *o.DiscountedPrice
}

// GetDiscountedPriceOk returns a tuple with the DiscountedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceModelDiscountTemplate) GetDiscountedPriceOk() (*GcpPriceValue, bool) {
	if o == nil || IsNil(o.DiscountedPrice) {
		return nil, false
	}
	return o.DiscountedPrice, true
}

// HasDiscountedPrice returns a boolean if a field has been set.
func (o *GcpPriceModelDiscountTemplate) HasDiscountedPrice() bool {
	if o != nil && !IsNil(o.DiscountedPrice) {
		return true
	}

	return false
}

// SetDiscountedPrice gets a reference to the given GcpPriceValue and assigns it to the DiscountedPrice field.
func (o *GcpPriceModelDiscountTemplate) SetDiscountedPrice(v GcpPriceValue) {
	o.DiscountedPrice = &v
}

// GetHideDiscountPercentage returns the HideDiscountPercentage field value if set, zero value otherwise.
func (o *GcpPriceModelDiscountTemplate) GetHideDiscountPercentage() bool {
	if o == nil || IsNil(o.HideDiscountPercentage) {
		var ret bool
		return ret
	}
	return *o.HideDiscountPercentage
}

// GetHideDiscountPercentageOk returns a tuple with the HideDiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceModelDiscountTemplate) GetHideDiscountPercentageOk() (*bool, bool) {
	if o == nil || IsNil(o.HideDiscountPercentage) {
		return nil, false
	}
	return o.HideDiscountPercentage, true
}

// HasHideDiscountPercentage returns a boolean if a field has been set.
func (o *GcpPriceModelDiscountTemplate) HasHideDiscountPercentage() bool {
	if o != nil && !IsNil(o.HideDiscountPercentage) {
		return true
	}

	return false
}

// SetHideDiscountPercentage gets a reference to the given bool and assigns it to the HideDiscountPercentage field.
func (o *GcpPriceModelDiscountTemplate) SetHideDiscountPercentage(v bool) {
	o.HideDiscountPercentage = &v
}

func (o GcpPriceModelDiscountTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpPriceModelDiscountTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountEconomics) {
		toSerialize["discountEconomics"] = o.DiscountEconomics
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.DiscountedPrice) {
		toSerialize["discountedPrice"] = o.DiscountedPrice
	}
	if !IsNil(o.HideDiscountPercentage) {
		toSerialize["hideDiscountPercentage"] = o.HideDiscountPercentage
	}
	return toSerialize, nil
}

type NullableGcpPriceModelDiscountTemplate struct {
	value *GcpPriceModelDiscountTemplate
	isSet bool
}

func (v NullableGcpPriceModelDiscountTemplate) Get() *GcpPriceModelDiscountTemplate {
	return v.value
}

func (v *NullableGcpPriceModelDiscountTemplate) Set(val *GcpPriceModelDiscountTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpPriceModelDiscountTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpPriceModelDiscountTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpPriceModelDiscountTemplate(val *GcpPriceModelDiscountTemplate) *NullableGcpPriceModelDiscountTemplate {
	return &NullableGcpPriceModelDiscountTemplate{value: val, isSet: true}
}

func (v NullableGcpPriceModelDiscountTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpPriceModelDiscountTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
