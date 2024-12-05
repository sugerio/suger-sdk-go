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

// checks if the GcpMarketplaceResellerPrivateOfferPlanMargin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlanMargin{}

// GcpMarketplaceResellerPrivateOfferPlanMargin struct for GcpMarketplaceResellerPrivateOfferPlanMargin
type GcpMarketplaceResellerPrivateOfferPlanMargin struct {
	MarginPercentage *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage `json:"marginPercentage,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlanMargin instantiates a new GcpMarketplaceResellerPrivateOfferPlanMargin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlanMargin() *GcpMarketplaceResellerPrivateOfferPlanMargin {
	this := GcpMarketplaceResellerPrivateOfferPlanMargin{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanMarginWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanMargin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanMarginWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanMargin {
	this := GcpMarketplaceResellerPrivateOfferPlanMargin{}
	return &this
}

// GetMarginPercentage returns the MarginPercentage field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanMargin) GetMarginPercentage() GcpMarketplaceResellerPrivateOfferPlanMarginPercentage {
	if o == nil || IsNil(o.MarginPercentage) {
		var ret GcpMarketplaceResellerPrivateOfferPlanMarginPercentage
		return ret
	}
	return *o.MarginPercentage
}

// GetMarginPercentageOk returns a tuple with the MarginPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanMargin) GetMarginPercentageOk() (*GcpMarketplaceResellerPrivateOfferPlanMarginPercentage, bool) {
	if o == nil || IsNil(o.MarginPercentage) {
		return nil, false
	}
	return o.MarginPercentage, true
}

// HasMarginPercentage returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanMargin) HasMarginPercentage() bool {
	if o != nil && !IsNil(o.MarginPercentage) {
		return true
	}

	return false
}

// SetMarginPercentage gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanMarginPercentage and assigns it to the MarginPercentage field.
func (o *GcpMarketplaceResellerPrivateOfferPlanMargin) SetMarginPercentage(v GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) {
	o.MarginPercentage = &v
}

func (o GcpMarketplaceResellerPrivateOfferPlanMargin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlanMargin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarginPercentage) {
		toSerialize["marginPercentage"] = o.MarginPercentage
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlanMargin struct {
	value *GcpMarketplaceResellerPrivateOfferPlanMargin
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMargin) Get() *GcpMarketplaceResellerPrivateOfferPlanMargin {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMargin) Set(val *GcpMarketplaceResellerPrivateOfferPlanMargin) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMargin) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlanMargin(val *GcpMarketplaceResellerPrivateOfferPlanMargin) *NullableGcpMarketplaceResellerPrivateOfferPlanMargin {
	return &NullableGcpMarketplaceResellerPrivateOfferPlanMargin{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
