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

// checks if the MicrosoftPartnerReferralTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftPartnerReferralTag{}

// MicrosoftPartnerReferralTag struct for MicrosoftPartnerReferralTag
type MicrosoftPartnerReferralTag struct {
	Id *string `json:"id,omitempty"`
}

// NewMicrosoftPartnerReferralTag instantiates a new MicrosoftPartnerReferralTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftPartnerReferralTag() *MicrosoftPartnerReferralTag {
	this := MicrosoftPartnerReferralTag{}
	return &this
}

// NewMicrosoftPartnerReferralTagWithDefaults instantiates a new MicrosoftPartnerReferralTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftPartnerReferralTagWithDefaults() *MicrosoftPartnerReferralTag {
	this := MicrosoftPartnerReferralTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftPartnerReferralTag) SetId(v string) {
	o.Id = &v
}

func (o MicrosoftPartnerReferralTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrosoftPartnerReferralTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableMicrosoftPartnerReferralTag struct {
	value *MicrosoftPartnerReferralTag
	isSet bool
}

func (v NullableMicrosoftPartnerReferralTag) Get() *MicrosoftPartnerReferralTag {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralTag) Set(val *MicrosoftPartnerReferralTag) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralTag) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralTag(val *MicrosoftPartnerReferralTag) *NullableMicrosoftPartnerReferralTag {
	return &NullableMicrosoftPartnerReferralTag{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
