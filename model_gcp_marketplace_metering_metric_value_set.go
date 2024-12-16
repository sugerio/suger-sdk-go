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

// checks if the GcpMarketplaceMeteringMetricValueSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceMeteringMetricValueSet{}

// GcpMarketplaceMeteringMetricValueSet struct for GcpMarketplaceMeteringMetricValueSet
type GcpMarketplaceMeteringMetricValueSet struct {
	// MetricName: The metric name defined in the service configuration.
	MetricName *string `json:"metricName,omitempty"`
	// MetricValues: The values in this metric.
	MetricValues []GcpMarketplaceMeteringMetricValue `json:"metricValues,omitempty"`
}

// NewGcpMarketplaceMeteringMetricValueSet instantiates a new GcpMarketplaceMeteringMetricValueSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceMeteringMetricValueSet() *GcpMarketplaceMeteringMetricValueSet {
	this := GcpMarketplaceMeteringMetricValueSet{}
	return &this
}

// NewGcpMarketplaceMeteringMetricValueSetWithDefaults instantiates a new GcpMarketplaceMeteringMetricValueSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceMeteringMetricValueSetWithDefaults() *GcpMarketplaceMeteringMetricValueSet {
	this := GcpMarketplaceMeteringMetricValueSet{}
	return &this
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *GcpMarketplaceMeteringMetricValueSet) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *GcpMarketplaceMeteringMetricValueSet) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMetricValues returns the MetricValues field value if set, zero value otherwise.
func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricValues() []GcpMarketplaceMeteringMetricValue {
	if o == nil || IsNil(o.MetricValues) {
		var ret []GcpMarketplaceMeteringMetricValue
		return ret
	}
	return o.MetricValues
}

// GetMetricValuesOk returns a tuple with the MetricValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricValuesOk() ([]GcpMarketplaceMeteringMetricValue, bool) {
	if o == nil || IsNil(o.MetricValues) {
		return nil, false
	}
	return o.MetricValues, true
}

// HasMetricValues returns a boolean if a field has been set.
func (o *GcpMarketplaceMeteringMetricValueSet) HasMetricValues() bool {
	if o != nil && !IsNil(o.MetricValues) {
		return true
	}

	return false
}

// SetMetricValues gets a reference to the given []GcpMarketplaceMeteringMetricValue and assigns it to the MetricValues field.
func (o *GcpMarketplaceMeteringMetricValueSet) SetMetricValues(v []GcpMarketplaceMeteringMetricValue) {
	o.MetricValues = v
}

func (o GcpMarketplaceMeteringMetricValueSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceMeteringMetricValueSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.MetricValues) {
		toSerialize["metricValues"] = o.MetricValues
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceMeteringMetricValueSet struct {
	value *GcpMarketplaceMeteringMetricValueSet
	isSet bool
}

func (v NullableGcpMarketplaceMeteringMetricValueSet) Get() *GcpMarketplaceMeteringMetricValueSet {
	return v.value
}

func (v *NullableGcpMarketplaceMeteringMetricValueSet) Set(val *GcpMarketplaceMeteringMetricValueSet) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceMeteringMetricValueSet) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceMeteringMetricValueSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceMeteringMetricValueSet(val *GcpMarketplaceMeteringMetricValueSet) *NullableGcpMarketplaceMeteringMetricValueSet {
	return &NullableGcpMarketplaceMeteringMetricValueSet{value: val, isSet: true}
}

func (v NullableGcpMarketplaceMeteringMetricValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceMeteringMetricValueSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
