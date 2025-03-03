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

// checks if the DatabaseSqlNullBool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseSqlNullBool{}

// DatabaseSqlNullBool struct for DatabaseSqlNullBool
type DatabaseSqlNullBool struct {
	Bool *bool `json:"bool,omitempty"`
	// Valid is true if Bool is not NULL
	Valid *bool `json:"valid,omitempty"`
}

// NewDatabaseSqlNullBool instantiates a new DatabaseSqlNullBool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseSqlNullBool() *DatabaseSqlNullBool {
	this := DatabaseSqlNullBool{}
	return &this
}

// NewDatabaseSqlNullBoolWithDefaults instantiates a new DatabaseSqlNullBool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseSqlNullBoolWithDefaults() *DatabaseSqlNullBool {
	this := DatabaseSqlNullBool{}
	return &this
}

// GetBool returns the Bool field value if set, zero value otherwise.
func (o *DatabaseSqlNullBool) GetBool() bool {
	if o == nil || IsNil(o.Bool) {
		var ret bool
		return ret
	}
	return *o.Bool
}

// GetBoolOk returns a tuple with the Bool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSqlNullBool) GetBoolOk() (*bool, bool) {
	if o == nil || IsNil(o.Bool) {
		return nil, false
	}
	return o.Bool, true
}

// HasBool returns a boolean if a field has been set.
func (o *DatabaseSqlNullBool) HasBool() bool {
	if o != nil && !IsNil(o.Bool) {
		return true
	}

	return false
}

// SetBool gets a reference to the given bool and assigns it to the Bool field.
func (o *DatabaseSqlNullBool) SetBool(v bool) {
	o.Bool = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *DatabaseSqlNullBool) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSqlNullBool) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *DatabaseSqlNullBool) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *DatabaseSqlNullBool) SetValid(v bool) {
	o.Valid = &v
}

func (o DatabaseSqlNullBool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseSqlNullBool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bool) {
		toSerialize["bool"] = o.Bool
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableDatabaseSqlNullBool struct {
	value *DatabaseSqlNullBool
	isSet bool
}

func (v NullableDatabaseSqlNullBool) Get() *DatabaseSqlNullBool {
	return v.value
}

func (v *NullableDatabaseSqlNullBool) Set(val *DatabaseSqlNullBool) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseSqlNullBool) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseSqlNullBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseSqlNullBool(val *DatabaseSqlNullBool) *NullableDatabaseSqlNullBool {
	return &NullableDatabaseSqlNullBool{value: val, isSet: true}
}

func (v NullableDatabaseSqlNullBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseSqlNullBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
