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

// checks if the AzureListingUri type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureListingUri{}

// AzureListingUri struct for AzureListingUri
type AzureListingUri struct {
	DisplayText *string `json:"displayText,omitempty"`
	Subtype     *string `json:"subtype,omitempty"`
	Type        *string `json:"type,omitempty"`
	Uri         *string `json:"uri,omitempty"`
}

// NewAzureListingUri instantiates a new AzureListingUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureListingUri() *AzureListingUri {
	this := AzureListingUri{}
	return &this
}

// NewAzureListingUriWithDefaults instantiates a new AzureListingUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureListingUriWithDefaults() *AzureListingUri {
	this := AzureListingUri{}
	return &this
}

// GetDisplayText returns the DisplayText field value if set, zero value otherwise.
func (o *AzureListingUri) GetDisplayText() string {
	if o == nil || IsNil(o.DisplayText) {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureListingUri) GetDisplayTextOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayText) {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *AzureListingUri) HasDisplayText() bool {
	if o != nil && !IsNil(o.DisplayText) {
		return true
	}

	return false
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *AzureListingUri) SetDisplayText(v string) {
	o.DisplayText = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *AzureListingUri) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureListingUri) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *AzureListingUri) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *AzureListingUri) SetSubtype(v string) {
	o.Subtype = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureListingUri) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureListingUri) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureListingUri) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureListingUri) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AzureListingUri) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureListingUri) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AzureListingUri) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AzureListingUri) SetUri(v string) {
	o.Uri = &v
}

func (o AzureListingUri) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureListingUri) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayText) {
		toSerialize["displayText"] = o.DisplayText
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableAzureListingUri struct {
	value *AzureListingUri
	isSet bool
}

func (v NullableAzureListingUri) Get() *AzureListingUri {
	return v.value
}

func (v *NullableAzureListingUri) Set(val *AzureListingUri) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureListingUri) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureListingUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureListingUri(val *AzureListingUri) *NullableAzureListingUri {
	return &NullableAzureListingUri{value: val, isSet: true}
}

func (v NullableAzureListingUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureListingUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
