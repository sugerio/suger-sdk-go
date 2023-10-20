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

// checks if the HubspotSyncFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubspotSyncFilter{}

// HubspotSyncFilter struct for HubspotSyncFilter
type HubspotSyncFilter struct {
	Operator *string `json:"operator,omitempty"`
	PropertyName *string `json:"propertyName,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewHubspotSyncFilter instantiates a new HubspotSyncFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubspotSyncFilter() *HubspotSyncFilter {
	this := HubspotSyncFilter{}
	return &this
}

// NewHubspotSyncFilterWithDefaults instantiates a new HubspotSyncFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubspotSyncFilterWithDefaults() *HubspotSyncFilter {
	this := HubspotSyncFilter{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *HubspotSyncFilter) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubspotSyncFilter) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *HubspotSyncFilter) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *HubspotSyncFilter) SetOperator(v string) {
	o.Operator = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *HubspotSyncFilter) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubspotSyncFilter) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *HubspotSyncFilter) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *HubspotSyncFilter) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HubspotSyncFilter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubspotSyncFilter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HubspotSyncFilter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HubspotSyncFilter) SetValue(v string) {
	o.Value = &v
}

func (o HubspotSyncFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubspotSyncFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableHubspotSyncFilter struct {
	value *HubspotSyncFilter
	isSet bool
}

func (v NullableHubspotSyncFilter) Get() *HubspotSyncFilter {
	return v.value
}

func (v *NullableHubspotSyncFilter) Set(val *HubspotSyncFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableHubspotSyncFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableHubspotSyncFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubspotSyncFilter(val *HubspotSyncFilter) *NullableHubspotSyncFilter {
	return &NullableHubspotSyncFilter{value: val, isSet: true}
}

func (v NullableHubspotSyncFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubspotSyncFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


