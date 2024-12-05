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

// checks if the DatabaseSqlNullTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseSqlNullTime{}

// DatabaseSqlNullTime struct for DatabaseSqlNullTime
type DatabaseSqlNullTime struct {
	Time *string `json:"time,omitempty"`
	// Valid is true if Time is not NULL
	Valid *bool `json:"valid,omitempty"`
}

// NewDatabaseSqlNullTime instantiates a new DatabaseSqlNullTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseSqlNullTime() *DatabaseSqlNullTime {
	this := DatabaseSqlNullTime{}
	return &this
}

// NewDatabaseSqlNullTimeWithDefaults instantiates a new DatabaseSqlNullTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseSqlNullTimeWithDefaults() *DatabaseSqlNullTime {
	this := DatabaseSqlNullTime{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DatabaseSqlNullTime) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSqlNullTime) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DatabaseSqlNullTime) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *DatabaseSqlNullTime) SetTime(v string) {
	o.Time = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *DatabaseSqlNullTime) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSqlNullTime) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *DatabaseSqlNullTime) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *DatabaseSqlNullTime) SetValid(v bool) {
	o.Valid = &v
}

func (o DatabaseSqlNullTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseSqlNullTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableDatabaseSqlNullTime struct {
	value *DatabaseSqlNullTime
	isSet bool
}

func (v NullableDatabaseSqlNullTime) Get() *DatabaseSqlNullTime {
	return v.value
}

func (v *NullableDatabaseSqlNullTime) Set(val *DatabaseSqlNullTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseSqlNullTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseSqlNullTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseSqlNullTime(val *DatabaseSqlNullTime) *NullableDatabaseSqlNullTime {
	return &NullableDatabaseSqlNullTime{value: val, isSet: true}
}

func (v NullableDatabaseSqlNullTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseSqlNullTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
