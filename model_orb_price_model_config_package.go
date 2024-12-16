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

// checks if the OrbPriceModelConfigPACKAGE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPriceModelConfigPACKAGE{}

// OrbPriceModelConfigPACKAGE struct for OrbPriceModelConfigPACKAGE
type OrbPriceModelConfigPACKAGE struct {
	PackageAmount *string  `json:"package_amount,omitempty"`
	PackageSize   *float32 `json:"package_size,omitempty"`
}

// NewOrbPriceModelConfigPACKAGE instantiates a new OrbPriceModelConfigPACKAGE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPriceModelConfigPACKAGE() *OrbPriceModelConfigPACKAGE {
	this := OrbPriceModelConfigPACKAGE{}
	return &this
}

// NewOrbPriceModelConfigPACKAGEWithDefaults instantiates a new OrbPriceModelConfigPACKAGE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPriceModelConfigPACKAGEWithDefaults() *OrbPriceModelConfigPACKAGE {
	this := OrbPriceModelConfigPACKAGE{}
	return &this
}

// GetPackageAmount returns the PackageAmount field value if set, zero value otherwise.
func (o *OrbPriceModelConfigPACKAGE) GetPackageAmount() string {
	if o == nil || IsNil(o.PackageAmount) {
		var ret string
		return ret
	}
	return *o.PackageAmount
}

// GetPackageAmountOk returns a tuple with the PackageAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigPACKAGE) GetPackageAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PackageAmount) {
		return nil, false
	}
	return o.PackageAmount, true
}

// HasPackageAmount returns a boolean if a field has been set.
func (o *OrbPriceModelConfigPACKAGE) HasPackageAmount() bool {
	if o != nil && !IsNil(o.PackageAmount) {
		return true
	}

	return false
}

// SetPackageAmount gets a reference to the given string and assigns it to the PackageAmount field.
func (o *OrbPriceModelConfigPACKAGE) SetPackageAmount(v string) {
	o.PackageAmount = &v
}

// GetPackageSize returns the PackageSize field value if set, zero value otherwise.
func (o *OrbPriceModelConfigPACKAGE) GetPackageSize() float32 {
	if o == nil || IsNil(o.PackageSize) {
		var ret float32
		return ret
	}
	return *o.PackageSize
}

// GetPackageSizeOk returns a tuple with the PackageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigPACKAGE) GetPackageSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.PackageSize) {
		return nil, false
	}
	return o.PackageSize, true
}

// HasPackageSize returns a boolean if a field has been set.
func (o *OrbPriceModelConfigPACKAGE) HasPackageSize() bool {
	if o != nil && !IsNil(o.PackageSize) {
		return true
	}

	return false
}

// SetPackageSize gets a reference to the given float32 and assigns it to the PackageSize field.
func (o *OrbPriceModelConfigPACKAGE) SetPackageSize(v float32) {
	o.PackageSize = &v
}

func (o OrbPriceModelConfigPACKAGE) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPriceModelConfigPACKAGE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageAmount) {
		toSerialize["package_amount"] = o.PackageAmount
	}
	if !IsNil(o.PackageSize) {
		toSerialize["package_size"] = o.PackageSize
	}
	return toSerialize, nil
}

type NullableOrbPriceModelConfigPACKAGE struct {
	value *OrbPriceModelConfigPACKAGE
	isSet bool
}

func (v NullableOrbPriceModelConfigPACKAGE) Get() *OrbPriceModelConfigPACKAGE {
	return v.value
}

func (v *NullableOrbPriceModelConfigPACKAGE) Set(val *OrbPriceModelConfigPACKAGE) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPriceModelConfigPACKAGE) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPriceModelConfigPACKAGE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPriceModelConfigPACKAGE(val *OrbPriceModelConfigPACKAGE) *NullableOrbPriceModelConfigPACKAGE {
	return &NullableOrbPriceModelConfigPACKAGE{value: val, isSet: true}
}

func (v NullableOrbPriceModelConfigPACKAGE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPriceModelConfigPACKAGE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
