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

// checks if the GcpMarketplaceResellerPrivateOfferPlanDurationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlanDurationConfig{}

// GcpMarketplaceResellerPrivateOfferPlanDurationConfig struct for GcpMarketplaceResellerPrivateOfferPlanDurationConfig
type GcpMarketplaceResellerPrivateOfferPlanDurationConfig struct {
	EndTime   *time.Time `json:"endTime,omitempty"`
	StartTime *time.Time `json:"startTime,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlanDurationConfig instantiates a new GcpMarketplaceResellerPrivateOfferPlanDurationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlanDurationConfig() *GcpMarketplaceResellerPrivateOfferPlanDurationConfig {
	this := GcpMarketplaceResellerPrivateOfferPlanDurationConfig{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanDurationConfigWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanDurationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanDurationConfigWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanDurationConfig {
	this := GcpMarketplaceResellerPrivateOfferPlanDurationConfig{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o GcpMarketplaceResellerPrivateOfferPlanDurationConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlanDurationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig struct {
	value *GcpMarketplaceResellerPrivateOfferPlanDurationConfig
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) Get() *GcpMarketplaceResellerPrivateOfferPlanDurationConfig {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) Set(val *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig(val *GcpMarketplaceResellerPrivateOfferPlanDurationConfig) *NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig {
	return &NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanDurationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
