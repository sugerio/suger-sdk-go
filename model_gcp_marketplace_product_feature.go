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

// checks if the GcpMarketplaceProductFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductFeature{}

// GcpMarketplaceProductFeature struct for GcpMarketplaceProductFeature
type GcpMarketplaceProductFeature struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// NewGcpMarketplaceProductFeature instantiates a new GcpMarketplaceProductFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductFeature() *GcpMarketplaceProductFeature {
	this := GcpMarketplaceProductFeature{}
	return &this
}

// NewGcpMarketplaceProductFeatureWithDefaults instantiates a new GcpMarketplaceProductFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductFeatureWithDefaults() *GcpMarketplaceProductFeature {
	this := GcpMarketplaceProductFeature{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GcpMarketplaceProductFeature) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductFeature) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductFeature) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GcpMarketplaceProductFeature) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GcpMarketplaceProductFeature) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductFeature) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductFeature) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GcpMarketplaceProductFeature) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GcpMarketplaceProductFeature) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductFeature) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GcpMarketplaceProductFeature) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GcpMarketplaceProductFeature) SetTitle(v string) {
	o.Title = &v
}

func (o GcpMarketplaceProductFeature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductFeature struct {
	value *GcpMarketplaceProductFeature
	isSet bool
}

func (v NullableGcpMarketplaceProductFeature) Get() *GcpMarketplaceProductFeature {
	return v.value
}

func (v *NullableGcpMarketplaceProductFeature) Set(val *GcpMarketplaceProductFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductFeature(val *GcpMarketplaceProductFeature) *NullableGcpMarketplaceProductFeature {
	return &NullableGcpMarketplaceProductFeature{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
