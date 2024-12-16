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

// checks if the EntitlementTermInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementTermInfo{}

// EntitlementTermInfo struct for EntitlementTermInfo
type EntitlementTermInfo struct {
	// The decimal part of the dimension quantity, in format of <dimensionKey, decimalPart> It is used to save the decimal part of the dimension quantity for AWS Marketplace only. Because AWS Marketplace only accepts integer for dimension quantity. If the dimension quantity is a decimal number, we need to save the decimal part for future use.
	DimensionQuantityDecimalParts *map[string]float32 `json:"dimensionQuantityDecimalParts,omitempty"`
	// Whether the commit is divided into multiple sub entitlement terms. If true, it has subEntitlementTermIds.
	IsCommitDivided *bool `json:"isCommitDivided,omitempty"`
	// The partner's entitlement term ID. It stands for the partner's entitlement term. Applicable to the sub entitlement term only.
	ParentEntitlementTermId *string `json:"parentEntitlementTermId,omitempty"`
	// All sub entitlement terms id of this entitlement term if it is commit divided.
	SubEntitlementTermIds []string             `json:"subEntitlementTermIds,omitempty"`
	Type                  *EntitlementTermType `json:"type,omitempty"`
}

// NewEntitlementTermInfo instantiates a new EntitlementTermInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementTermInfo() *EntitlementTermInfo {
	this := EntitlementTermInfo{}
	return &this
}

// NewEntitlementTermInfoWithDefaults instantiates a new EntitlementTermInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementTermInfoWithDefaults() *EntitlementTermInfo {
	this := EntitlementTermInfo{}
	return &this
}

// GetDimensionQuantityDecimalParts returns the DimensionQuantityDecimalParts field value if set, zero value otherwise.
func (o *EntitlementTermInfo) GetDimensionQuantityDecimalParts() map[string]float32 {
	if o == nil || IsNil(o.DimensionQuantityDecimalParts) {
		var ret map[string]float32
		return ret
	}
	return *o.DimensionQuantityDecimalParts
}

// GetDimensionQuantityDecimalPartsOk returns a tuple with the DimensionQuantityDecimalParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTermInfo) GetDimensionQuantityDecimalPartsOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.DimensionQuantityDecimalParts) {
		return nil, false
	}
	return o.DimensionQuantityDecimalParts, true
}

// HasDimensionQuantityDecimalParts returns a boolean if a field has been set.
func (o *EntitlementTermInfo) HasDimensionQuantityDecimalParts() bool {
	if o != nil && !IsNil(o.DimensionQuantityDecimalParts) {
		return true
	}

	return false
}

// SetDimensionQuantityDecimalParts gets a reference to the given map[string]float32 and assigns it to the DimensionQuantityDecimalParts field.
func (o *EntitlementTermInfo) SetDimensionQuantityDecimalParts(v map[string]float32) {
	o.DimensionQuantityDecimalParts = &v
}

// GetIsCommitDivided returns the IsCommitDivided field value if set, zero value otherwise.
func (o *EntitlementTermInfo) GetIsCommitDivided() bool {
	if o == nil || IsNil(o.IsCommitDivided) {
		var ret bool
		return ret
	}
	return *o.IsCommitDivided
}

// GetIsCommitDividedOk returns a tuple with the IsCommitDivided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTermInfo) GetIsCommitDividedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCommitDivided) {
		return nil, false
	}
	return o.IsCommitDivided, true
}

// HasIsCommitDivided returns a boolean if a field has been set.
func (o *EntitlementTermInfo) HasIsCommitDivided() bool {
	if o != nil && !IsNil(o.IsCommitDivided) {
		return true
	}

	return false
}

// SetIsCommitDivided gets a reference to the given bool and assigns it to the IsCommitDivided field.
func (o *EntitlementTermInfo) SetIsCommitDivided(v bool) {
	o.IsCommitDivided = &v
}

// GetParentEntitlementTermId returns the ParentEntitlementTermId field value if set, zero value otherwise.
func (o *EntitlementTermInfo) GetParentEntitlementTermId() string {
	if o == nil || IsNil(o.ParentEntitlementTermId) {
		var ret string
		return ret
	}
	return *o.ParentEntitlementTermId
}

// GetParentEntitlementTermIdOk returns a tuple with the ParentEntitlementTermId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTermInfo) GetParentEntitlementTermIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentEntitlementTermId) {
		return nil, false
	}
	return o.ParentEntitlementTermId, true
}

// HasParentEntitlementTermId returns a boolean if a field has been set.
func (o *EntitlementTermInfo) HasParentEntitlementTermId() bool {
	if o != nil && !IsNil(o.ParentEntitlementTermId) {
		return true
	}

	return false
}

// SetParentEntitlementTermId gets a reference to the given string and assigns it to the ParentEntitlementTermId field.
func (o *EntitlementTermInfo) SetParentEntitlementTermId(v string) {
	o.ParentEntitlementTermId = &v
}

// GetSubEntitlementTermIds returns the SubEntitlementTermIds field value if set, zero value otherwise.
func (o *EntitlementTermInfo) GetSubEntitlementTermIds() []string {
	if o == nil || IsNil(o.SubEntitlementTermIds) {
		var ret []string
		return ret
	}
	return o.SubEntitlementTermIds
}

// GetSubEntitlementTermIdsOk returns a tuple with the SubEntitlementTermIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTermInfo) GetSubEntitlementTermIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubEntitlementTermIds) {
		return nil, false
	}
	return o.SubEntitlementTermIds, true
}

// HasSubEntitlementTermIds returns a boolean if a field has been set.
func (o *EntitlementTermInfo) HasSubEntitlementTermIds() bool {
	if o != nil && !IsNil(o.SubEntitlementTermIds) {
		return true
	}

	return false
}

// SetSubEntitlementTermIds gets a reference to the given []string and assigns it to the SubEntitlementTermIds field.
func (o *EntitlementTermInfo) SetSubEntitlementTermIds(v []string) {
	o.SubEntitlementTermIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitlementTermInfo) GetType() EntitlementTermType {
	if o == nil || IsNil(o.Type) {
		var ret EntitlementTermType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTermInfo) GetTypeOk() (*EntitlementTermType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitlementTermInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EntitlementTermType and assigns it to the Type field.
func (o *EntitlementTermInfo) SetType(v EntitlementTermType) {
	o.Type = &v
}

func (o EntitlementTermInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementTermInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DimensionQuantityDecimalParts) {
		toSerialize["dimensionQuantityDecimalParts"] = o.DimensionQuantityDecimalParts
	}
	if !IsNil(o.IsCommitDivided) {
		toSerialize["isCommitDivided"] = o.IsCommitDivided
	}
	if !IsNil(o.ParentEntitlementTermId) {
		toSerialize["parentEntitlementTermId"] = o.ParentEntitlementTermId
	}
	if !IsNil(o.SubEntitlementTermIds) {
		toSerialize["subEntitlementTermIds"] = o.SubEntitlementTermIds
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEntitlementTermInfo struct {
	value *EntitlementTermInfo
	isSet bool
}

func (v NullableEntitlementTermInfo) Get() *EntitlementTermInfo {
	return v.value
}

func (v *NullableEntitlementTermInfo) Set(val *EntitlementTermInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementTermInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementTermInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementTermInfo(val *EntitlementTermInfo) *NullableEntitlementTermInfo {
	return &NullableEntitlementTermInfo{value: val, isSet: true}
}

func (v NullableEntitlementTermInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementTermInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
