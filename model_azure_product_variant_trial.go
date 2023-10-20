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

// checks if the AzureProductVariantTrial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductVariantTrial{}

// AzureProductVariantTrial struct for AzureProductVariantTrial
type AzureProductVariantTrial struct {
	DateTimeRange *AzureLocalizedTimeRange `json:"dateTimeRange,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	DurationType *string `json:"durationType,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAzureProductVariantTrial instantiates a new AzureProductVariantTrial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductVariantTrial() *AzureProductVariantTrial {
	this := AzureProductVariantTrial{}
	return &this
}

// NewAzureProductVariantTrialWithDefaults instantiates a new AzureProductVariantTrial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductVariantTrialWithDefaults() *AzureProductVariantTrial {
	this := AzureProductVariantTrial{}
	return &this
}

// GetDateTimeRange returns the DateTimeRange field value if set, zero value otherwise.
func (o *AzureProductVariantTrial) GetDateTimeRange() AzureLocalizedTimeRange {
	if o == nil || IsNil(o.DateTimeRange) {
		var ret AzureLocalizedTimeRange
		return ret
	}
	return *o.DateTimeRange
}

// GetDateTimeRangeOk returns a tuple with the DateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantTrial) GetDateTimeRangeOk() (*AzureLocalizedTimeRange, bool) {
	if o == nil || IsNil(o.DateTimeRange) {
		return nil, false
	}
	return o.DateTimeRange, true
}

// HasDateTimeRange returns a boolean if a field has been set.
func (o *AzureProductVariantTrial) HasDateTimeRange() bool {
	if o != nil && !IsNil(o.DateTimeRange) {
		return true
	}

	return false
}

// SetDateTimeRange gets a reference to the given AzureLocalizedTimeRange and assigns it to the DateTimeRange field.
func (o *AzureProductVariantTrial) SetDateTimeRange(v AzureLocalizedTimeRange) {
	o.DateTimeRange = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AzureProductVariantTrial) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantTrial) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AzureProductVariantTrial) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AzureProductVariantTrial) SetDuration(v int32) {
	o.Duration = &v
}

// GetDurationType returns the DurationType field value if set, zero value otherwise.
func (o *AzureProductVariantTrial) GetDurationType() string {
	if o == nil || IsNil(o.DurationType) {
		var ret string
		return ret
	}
	return *o.DurationType
}

// GetDurationTypeOk returns a tuple with the DurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantTrial) GetDurationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DurationType) {
		return nil, false
	}
	return o.DurationType, true
}

// HasDurationType returns a boolean if a field has been set.
func (o *AzureProductVariantTrial) HasDurationType() bool {
	if o != nil && !IsNil(o.DurationType) {
		return true
	}

	return false
}

// SetDurationType gets a reference to the given string and assigns it to the DurationType field.
func (o *AzureProductVariantTrial) SetDurationType(v string) {
	o.DurationType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureProductVariantTrial) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantTrial) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureProductVariantTrial) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureProductVariantTrial) SetType(v string) {
	o.Type = &v
}

func (o AzureProductVariantTrial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductVariantTrial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateTimeRange) {
		toSerialize["dateTimeRange"] = o.DateTimeRange
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.DurationType) {
		toSerialize["durationType"] = o.DurationType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAzureProductVariantTrial struct {
	value *AzureProductVariantTrial
	isSet bool
}

func (v NullableAzureProductVariantTrial) Get() *AzureProductVariantTrial {
	return v.value
}

func (v *NullableAzureProductVariantTrial) Set(val *AzureProductVariantTrial) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductVariantTrial) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductVariantTrial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductVariantTrial(val *AzureProductVariantTrial) *NullableAzureProductVariantTrial {
	return &NullableAzureProductVariantTrial{value: val, isSet: true}
}

func (v NullableAzureProductVariantTrial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductVariantTrial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


