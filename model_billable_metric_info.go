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

// checks if the BillableMetricInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableMetricInfo{}

// BillableMetricInfo struct for BillableMetricInfo
type BillableMetricInfo struct {
	// FilterGroups is a list of filter groups. The filterGroups are connected by AND.
	FilterGroups []BillableMetricFilterGroup `json:"filterGroups,omitempty"`
	// GroupBys is a list of fields to group by.
	GroupBys []string `json:"groupBys,omitempty"`
	// The target property for unique count aggregate.
	PropertyUniqueOn *string `json:"propertyUniqueOn,omitempty"`
}

// NewBillableMetricInfo instantiates a new BillableMetricInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableMetricInfo() *BillableMetricInfo {
	this := BillableMetricInfo{}
	return &this
}

// NewBillableMetricInfoWithDefaults instantiates a new BillableMetricInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableMetricInfoWithDefaults() *BillableMetricInfo {
	this := BillableMetricInfo{}
	return &this
}

// GetFilterGroups returns the FilterGroups field value if set, zero value otherwise.
func (o *BillableMetricInfo) GetFilterGroups() []BillableMetricFilterGroup {
	if o == nil || IsNil(o.FilterGroups) {
		var ret []BillableMetricFilterGroup
		return ret
	}
	return o.FilterGroups
}

// GetFilterGroupsOk returns a tuple with the FilterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricInfo) GetFilterGroupsOk() ([]BillableMetricFilterGroup, bool) {
	if o == nil || IsNil(o.FilterGroups) {
		return nil, false
	}
	return o.FilterGroups, true
}

// HasFilterGroups returns a boolean if a field has been set.
func (o *BillableMetricInfo) HasFilterGroups() bool {
	if o != nil && !IsNil(o.FilterGroups) {
		return true
	}

	return false
}

// SetFilterGroups gets a reference to the given []BillableMetricFilterGroup and assigns it to the FilterGroups field.
func (o *BillableMetricInfo) SetFilterGroups(v []BillableMetricFilterGroup) {
	o.FilterGroups = v
}

// GetGroupBys returns the GroupBys field value if set, zero value otherwise.
func (o *BillableMetricInfo) GetGroupBys() []string {
	if o == nil || IsNil(o.GroupBys) {
		var ret []string
		return ret
	}
	return o.GroupBys
}

// GetGroupBysOk returns a tuple with the GroupBys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricInfo) GetGroupBysOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupBys) {
		return nil, false
	}
	return o.GroupBys, true
}

// HasGroupBys returns a boolean if a field has been set.
func (o *BillableMetricInfo) HasGroupBys() bool {
	if o != nil && !IsNil(o.GroupBys) {
		return true
	}

	return false
}

// SetGroupBys gets a reference to the given []string and assigns it to the GroupBys field.
func (o *BillableMetricInfo) SetGroupBys(v []string) {
	o.GroupBys = v
}

// GetPropertyUniqueOn returns the PropertyUniqueOn field value if set, zero value otherwise.
func (o *BillableMetricInfo) GetPropertyUniqueOn() string {
	if o == nil || IsNil(o.PropertyUniqueOn) {
		var ret string
		return ret
	}
	return *o.PropertyUniqueOn
}

// GetPropertyUniqueOnOk returns a tuple with the PropertyUniqueOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableMetricInfo) GetPropertyUniqueOnOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyUniqueOn) {
		return nil, false
	}
	return o.PropertyUniqueOn, true
}

// HasPropertyUniqueOn returns a boolean if a field has been set.
func (o *BillableMetricInfo) HasPropertyUniqueOn() bool {
	if o != nil && !IsNil(o.PropertyUniqueOn) {
		return true
	}

	return false
}

// SetPropertyUniqueOn gets a reference to the given string and assigns it to the PropertyUniqueOn field.
func (o *BillableMetricInfo) SetPropertyUniqueOn(v string) {
	o.PropertyUniqueOn = &v
}

func (o BillableMetricInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableMetricInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterGroups) {
		toSerialize["filterGroups"] = o.FilterGroups
	}
	if !IsNil(o.GroupBys) {
		toSerialize["groupBys"] = o.GroupBys
	}
	if !IsNil(o.PropertyUniqueOn) {
		toSerialize["propertyUniqueOn"] = o.PropertyUniqueOn
	}
	return toSerialize, nil
}

type NullableBillableMetricInfo struct {
	value *BillableMetricInfo
	isSet bool
}

func (v NullableBillableMetricInfo) Get() *BillableMetricInfo {
	return v.value
}

func (v *NullableBillableMetricInfo) Set(val *BillableMetricInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricInfo(val *BillableMetricInfo) *NullableBillableMetricInfo {
	return &NullableBillableMetricInfo{value: val, isSet: true}
}

func (v NullableBillableMetricInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
