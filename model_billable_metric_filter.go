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

// checks if the BillableMetricFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricFilter{}

// BillableMetricFilter struct for BillableMetricFilter
type BillableMetricFilter struct {
	Name      *string                        `json:"name,omitempty"`
	Operation *BillableMetricFilterOperation `json:"operation,omitempty"`
	// The value of the filter. The type of the value depends on the valueType.
	Value     map[string]interface{}         `json:"value,omitempty"`
	ValueType *BillableMetricFilterValueType `json:"valueType,omitempty"`
}

// NewBillableMetricFilter instantiates a new BillableMetricFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricFilter() *BillableMetricFilter {
	this := BillableMetricFilter{}
	return &this
}

// NewBillableMetricFilterWithDefaults instantiates a new BillableMetricFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricFilterWithDefaults() *BillableMetricFilter {
	this := BillableMetricFilter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillableMetricFilter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricFilter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillableMetricFilter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillableMetricFilter) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *BillableMetricFilter) GetOperation() BillableMetricFilterOperation {
	if o == nil || IsNil(o.Operation) {
		var ret BillableMetricFilterOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricFilter) GetOperationOk() (*BillableMetricFilterOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *BillableMetricFilter) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given BillableMetricFilterOperation and assigns it to the Operation field.
func (o *BillableMetricFilter) SetOperation(v BillableMetricFilterOperation) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BillableMetricFilter) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricFilter) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BillableMetricFilter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *BillableMetricFilter) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *BillableMetricFilter) GetValueType() BillableMetricFilterValueType {
	if o == nil || IsNil(o.ValueType) {
		var ret BillableMetricFilterValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricFilter) GetValueTypeOk() (*BillableMetricFilterValueType, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *BillableMetricFilter) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given BillableMetricFilterValueType and assigns it to the ValueType field.
func (o *BillableMetricFilter) SetValueType(v BillableMetricFilterValueType) {
	o.ValueType = &v
}

func (o BillableMetricFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	return toSerialize, nil
}

type NullableBillableMetricFilter struct {
	value *BillableMetricFilter
	isSet bool
}

func (v NullableBillableMetricFilter) Get() *BillableMetricFilter {
	return v.value
}

func (v *NullableBillableMetricFilter) Set(val *BillableMetricFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricFilter(val *BillableMetricFilter) *NullableBillableMetricFilter {
	return &NullableBillableMetricFilter{value: val, isSet: true}
}

func (v NullableBillableMetricFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
