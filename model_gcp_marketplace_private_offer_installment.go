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

// checks if the GcpMarketplacePrivateOfferInstallment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplacePrivateOfferInstallment{}

// GcpMarketplacePrivateOfferInstallment struct for GcpMarketplacePrivateOfferInstallment
type GcpMarketplacePrivateOfferInstallment struct {
	PriceModel *GcpMarketplacePrivateOfferPriceModel `json:"priceModel,omitempty"`
	StartTime  *time.Time                            `json:"startTime,omitempty"`
}

// NewGcpMarketplacePrivateOfferInstallment instantiates a new GcpMarketplacePrivateOfferInstallment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplacePrivateOfferInstallment() *GcpMarketplacePrivateOfferInstallment {
	this := GcpMarketplacePrivateOfferInstallment{}
	return &this
}

// NewGcpMarketplacePrivateOfferInstallmentWithDefaults instantiates a new GcpMarketplacePrivateOfferInstallment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplacePrivateOfferInstallmentWithDefaults() *GcpMarketplacePrivateOfferInstallment {
	this := GcpMarketplacePrivateOfferInstallment{}
	return &this
}

// GetPriceModel returns the PriceModel field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferInstallment) GetPriceModel() GcpMarketplacePrivateOfferPriceModel {
	if o == nil || IsNil(o.PriceModel) {
		var ret GcpMarketplacePrivateOfferPriceModel
		return ret
	}
	return *o.PriceModel
}

// GetPriceModelOk returns a tuple with the PriceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferInstallment) GetPriceModelOk() (*GcpMarketplacePrivateOfferPriceModel, bool) {
	if o == nil || IsNil(o.PriceModel) {
		return nil, false
	}
	return o.PriceModel, true
}

// HasPriceModel returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferInstallment) HasPriceModel() bool {
	if o != nil && !IsNil(o.PriceModel) {
		return true
	}

	return false
}

// SetPriceModel gets a reference to the given GcpMarketplacePrivateOfferPriceModel and assigns it to the PriceModel field.
func (o *GcpMarketplacePrivateOfferInstallment) SetPriceModel(v GcpMarketplacePrivateOfferPriceModel) {
	o.PriceModel = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferInstallment) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferInstallment) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferInstallment) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GcpMarketplacePrivateOfferInstallment) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o GcpMarketplacePrivateOfferInstallment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplacePrivateOfferInstallment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceModel) {
		toSerialize["priceModel"] = o.PriceModel
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableGcpMarketplacePrivateOfferInstallment struct {
	value *GcpMarketplacePrivateOfferInstallment
	isSet bool
}

func (v NullableGcpMarketplacePrivateOfferInstallment) Get() *GcpMarketplacePrivateOfferInstallment {
	return v.value
}

func (v *NullableGcpMarketplacePrivateOfferInstallment) Set(val *GcpMarketplacePrivateOfferInstallment) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplacePrivateOfferInstallment) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplacePrivateOfferInstallment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplacePrivateOfferInstallment(val *GcpMarketplacePrivateOfferInstallment) *NullableGcpMarketplacePrivateOfferInstallment {
	return &NullableGcpMarketplacePrivateOfferInstallment{value: val, isSet: true}
}

func (v NullableGcpMarketplacePrivateOfferInstallment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplacePrivateOfferInstallment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
