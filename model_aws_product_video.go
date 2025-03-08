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

// checks if the AwsProductVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProductVideo{}

// AwsProductVideo struct for AwsProductVideo
type AwsProductVideo struct {
	Title *string `json:"Title,omitempty"`
	Type  *string `json:"Type,omitempty"`
	Url   *string `json:"Url,omitempty"`
}

// NewAwsProductVideo instantiates a new AwsProductVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProductVideo() *AwsProductVideo {
	this := AwsProductVideo{}
	return &this
}

// NewAwsProductVideoWithDefaults instantiates a new AwsProductVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProductVideoWithDefaults() *AwsProductVideo {
	this := AwsProductVideo{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AwsProductVideo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductVideo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AwsProductVideo) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AwsProductVideo) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AwsProductVideo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductVideo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AwsProductVideo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AwsProductVideo) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AwsProductVideo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductVideo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AwsProductVideo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AwsProductVideo) SetUrl(v string) {
	o.Url = &v
}

func (o AwsProductVideo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProductVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAwsProductVideo struct {
	value *AwsProductVideo
	isSet bool
}

func (v NullableAwsProductVideo) Get() *AwsProductVideo {
	return v.value
}

func (v *NullableAwsProductVideo) Set(val *AwsProductVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProductVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProductVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProductVideo(val *AwsProductVideo) *NullableAwsProductVideo {
	return &NullableAwsProductVideo{value: val, isSet: true}
}

func (v NullableAwsProductVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProductVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
