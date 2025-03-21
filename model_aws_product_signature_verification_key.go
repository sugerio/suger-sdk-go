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

// checks if the AwsProductSignatureVerificationKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProductSignatureVerificationKey{}

// AwsProductSignatureVerificationKey struct for AwsProductSignatureVerificationKey
type AwsProductSignatureVerificationKey struct {
	PublicKey        *string `json:"PublicKey,omitempty"`
	PublicKeyVersion *int32  `json:"PublicKeyVersion,omitempty"`
	Status           *string `json:"Status,omitempty"`
}

// NewAwsProductSignatureVerificationKey instantiates a new AwsProductSignatureVerificationKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProductSignatureVerificationKey() *AwsProductSignatureVerificationKey {
	this := AwsProductSignatureVerificationKey{}
	return &this
}

// NewAwsProductSignatureVerificationKeyWithDefaults instantiates a new AwsProductSignatureVerificationKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProductSignatureVerificationKeyWithDefaults() *AwsProductSignatureVerificationKey {
	this := AwsProductSignatureVerificationKey{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *AwsProductSignatureVerificationKey) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductSignatureVerificationKey) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *AwsProductSignatureVerificationKey) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *AwsProductSignatureVerificationKey) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPublicKeyVersion returns the PublicKeyVersion field value if set, zero value otherwise.
func (o *AwsProductSignatureVerificationKey) GetPublicKeyVersion() int32 {
	if o == nil || IsNil(o.PublicKeyVersion) {
		var ret int32
		return ret
	}
	return *o.PublicKeyVersion
}

// GetPublicKeyVersionOk returns a tuple with the PublicKeyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductSignatureVerificationKey) GetPublicKeyVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicKeyVersion) {
		return nil, false
	}
	return o.PublicKeyVersion, true
}

// HasPublicKeyVersion returns a boolean if a field has been set.
func (o *AwsProductSignatureVerificationKey) HasPublicKeyVersion() bool {
	if o != nil && !IsNil(o.PublicKeyVersion) {
		return true
	}

	return false
}

// SetPublicKeyVersion gets a reference to the given int32 and assigns it to the PublicKeyVersion field.
func (o *AwsProductSignatureVerificationKey) SetPublicKeyVersion(v int32) {
	o.PublicKeyVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsProductSignatureVerificationKey) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductSignatureVerificationKey) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsProductSignatureVerificationKey) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AwsProductSignatureVerificationKey) SetStatus(v string) {
	o.Status = &v
}

func (o AwsProductSignatureVerificationKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProductSignatureVerificationKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicKey) {
		toSerialize["PublicKey"] = o.PublicKey
	}
	if !IsNil(o.PublicKeyVersion) {
		toSerialize["PublicKeyVersion"] = o.PublicKeyVersion
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAwsProductSignatureVerificationKey struct {
	value *AwsProductSignatureVerificationKey
	isSet bool
}

func (v NullableAwsProductSignatureVerificationKey) Get() *AwsProductSignatureVerificationKey {
	return v.value
}

func (v *NullableAwsProductSignatureVerificationKey) Set(val *AwsProductSignatureVerificationKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProductSignatureVerificationKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProductSignatureVerificationKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProductSignatureVerificationKey(val *AwsProductSignatureVerificationKey) *NullableAwsProductSignatureVerificationKey {
	return &NullableAwsProductSignatureVerificationKey{value: val, isSet: true}
}

func (v NullableAwsProductSignatureVerificationKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProductSignatureVerificationKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
