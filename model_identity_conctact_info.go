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

// checks if the IdentityConctactInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityConctactInfo{}

// IdentityConctactInfo struct for IdentityConctactInfo
type IdentityConctactInfo struct {
	CompanyLocation *string `json:"companyLocation,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NewIdentityConctactInfo instantiates a new IdentityConctactInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityConctactInfo() *IdentityConctactInfo {
	this := IdentityConctactInfo{}
	return &this
}

// NewIdentityConctactInfoWithDefaults instantiates a new IdentityConctactInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityConctactInfoWithDefaults() *IdentityConctactInfo {
	this := IdentityConctactInfo{}
	return &this
}

// GetCompanyLocation returns the CompanyLocation field value if set, zero value otherwise.
func (o *IdentityConctactInfo) GetCompanyLocation() string {
	if o == nil || IsNil(o.CompanyLocation) {
		var ret string
		return ret
	}
	return *o.CompanyLocation
}

// GetCompanyLocationOk returns a tuple with the CompanyLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityConctactInfo) GetCompanyLocationOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyLocation) {
		return nil, false
	}
	return o.CompanyLocation, true
}

// HasCompanyLocation returns a boolean if a field has been set.
func (o *IdentityConctactInfo) HasCompanyLocation() bool {
	if o != nil && !IsNil(o.CompanyLocation) {
		return true
	}

	return false
}

// SetCompanyLocation gets a reference to the given string and assigns it to the CompanyLocation field.
func (o *IdentityConctactInfo) SetCompanyLocation(v string) {
	o.CompanyLocation = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *IdentityConctactInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityConctactInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *IdentityConctactInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *IdentityConctactInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *IdentityConctactInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityConctactInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityConctactInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *IdentityConctactInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *IdentityConctactInfo) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityConctactInfo) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *IdentityConctactInfo) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *IdentityConctactInfo) SetRole(v string) {
	o.Role = &v
}

func (o IdentityConctactInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityConctactInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyLocation) {
		toSerialize["companyLocation"] = o.CompanyLocation
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableIdentityConctactInfo struct {
	value *IdentityConctactInfo
	isSet bool
}

func (v NullableIdentityConctactInfo) Get() *IdentityConctactInfo {
	return v.value
}

func (v *NullableIdentityConctactInfo) Set(val *IdentityConctactInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityConctactInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityConctactInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityConctactInfo(val *IdentityConctactInfo) *NullableIdentityConctactInfo {
	return &NullableIdentityConctactInfo{value: val, isSet: true}
}

func (v NullableIdentityConctactInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityConctactInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


