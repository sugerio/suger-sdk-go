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

// checks if the MicrosoftPartnerReferralProfileId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftPartnerReferralProfileId{}

// MicrosoftPartnerReferralProfileId struct for MicrosoftPartnerReferralProfileId
type MicrosoftPartnerReferralProfileId struct {
	Id          *string `json:"id,omitempty"`
	ProfileType *string `json:"profileType,omitempty"`
}

// NewMicrosoftPartnerReferralProfileId instantiates a new MicrosoftPartnerReferralProfileId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftPartnerReferralProfileId() *MicrosoftPartnerReferralProfileId {
	this := MicrosoftPartnerReferralProfileId{}
	return &this
}

// NewMicrosoftPartnerReferralProfileIdWithDefaults instantiates a new MicrosoftPartnerReferralProfileId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftPartnerReferralProfileIdWithDefaults() *MicrosoftPartnerReferralProfileId {
	this := MicrosoftPartnerReferralProfileId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralProfileId) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralProfileId) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralProfileId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftPartnerReferralProfileId) SetId(v string) {
	o.Id = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralProfileId) GetProfileType() string {
	if o == nil || IsNil(o.ProfileType) {
		var ret string
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralProfileId) GetProfileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralProfileId) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given string and assigns it to the ProfileType field.
func (o *MicrosoftPartnerReferralProfileId) SetProfileType(v string) {
	o.ProfileType = &v
}

func (o MicrosoftPartnerReferralProfileId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrosoftPartnerReferralProfileId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	return toSerialize, nil
}

type NullableMicrosoftPartnerReferralProfileId struct {
	value *MicrosoftPartnerReferralProfileId
	isSet bool
}

func (v NullableMicrosoftPartnerReferralProfileId) Get() *MicrosoftPartnerReferralProfileId {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralProfileId) Set(val *MicrosoftPartnerReferralProfileId) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralProfileId) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralProfileId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralProfileId(val *MicrosoftPartnerReferralProfileId) *NullableMicrosoftPartnerReferralProfileId {
	return &NullableMicrosoftPartnerReferralProfileId{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralProfileId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralProfileId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
