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

// checks if the SalesforceSyncFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesforceSyncFilter{}

// SalesforceSyncFilter struct for SalesforceSyncFilter
type SalesforceSyncFilter struct {
	FieldName *string                       `json:"fieldName,omitempty"`
	Operator  *SalesforceSyncFilterOperator `json:"operator,omitempty"`
	Value     map[string]interface{}        `json:"value,omitempty"`
}

// NewSalesforceSyncFilter instantiates a new SalesforceSyncFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesforceSyncFilter() *SalesforceSyncFilter {
	this := SalesforceSyncFilter{}
	return &this
}

// NewSalesforceSyncFilterWithDefaults instantiates a new SalesforceSyncFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesforceSyncFilterWithDefaults() *SalesforceSyncFilter {
	this := SalesforceSyncFilter{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *SalesforceSyncFilter) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceSyncFilter) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *SalesforceSyncFilter) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *SalesforceSyncFilter) SetFieldName(v string) {
	o.FieldName = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SalesforceSyncFilter) GetOperator() SalesforceSyncFilterOperator {
	if o == nil || IsNil(o.Operator) {
		var ret SalesforceSyncFilterOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceSyncFilter) GetOperatorOk() (*SalesforceSyncFilterOperator, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SalesforceSyncFilter) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given SalesforceSyncFilterOperator and assigns it to the Operator field.
func (o *SalesforceSyncFilter) SetOperator(v SalesforceSyncFilterOperator) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SalesforceSyncFilter) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceSyncFilter) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SalesforceSyncFilter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *SalesforceSyncFilter) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o SalesforceSyncFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesforceSyncFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["fieldName"] = o.FieldName
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSalesforceSyncFilter struct {
	value *SalesforceSyncFilter
	isSet bool
}

func (v NullableSalesforceSyncFilter) Get() *SalesforceSyncFilter {
	return v.value
}

func (v *NullableSalesforceSyncFilter) Set(val *SalesforceSyncFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesforceSyncFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesforceSyncFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesforceSyncFilter(val *SalesforceSyncFilter) *NullableSalesforceSyncFilter {
	return &NullableSalesforceSyncFilter{value: val, isSet: true}
}

func (v NullableSalesforceSyncFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesforceSyncFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
