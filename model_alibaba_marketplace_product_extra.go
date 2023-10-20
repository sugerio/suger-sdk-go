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

// checks if the AlibabaMarketplaceProductExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlibabaMarketplaceProductExtra{}

// AlibabaMarketplaceProductExtra struct for AlibabaMarketplaceProductExtra
type AlibabaMarketplaceProductExtra struct {
	Key *string `json:"Key,omitempty"`
	Label *string `json:"Label,omitempty"`
	Order *int32 `json:"Order,omitempty"`
	Type *string `json:"Type,omitempty"`
	Values map[string]interface{} `json:"Values,omitempty"`
}

// NewAlibabaMarketplaceProductExtra instantiates a new AlibabaMarketplaceProductExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaMarketplaceProductExtra() *AlibabaMarketplaceProductExtra {
	this := AlibabaMarketplaceProductExtra{}
	return &this
}

// NewAlibabaMarketplaceProductExtraWithDefaults instantiates a new AlibabaMarketplaceProductExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaMarketplaceProductExtraWithDefaults() *AlibabaMarketplaceProductExtra {
	this := AlibabaMarketplaceProductExtra{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductExtra) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductExtra) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductExtra) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AlibabaMarketplaceProductExtra) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductExtra) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductExtra) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductExtra) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AlibabaMarketplaceProductExtra) SetLabel(v string) {
	o.Label = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductExtra) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductExtra) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductExtra) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *AlibabaMarketplaceProductExtra) SetOrder(v int32) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductExtra) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductExtra) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductExtra) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlibabaMarketplaceProductExtra) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductExtra) GetValues() map[string]interface{} {
	if o == nil || IsNil(o.Values) {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductExtra) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Values) {
		return map[string]interface{}{}, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductExtra) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *AlibabaMarketplaceProductExtra) SetValues(v map[string]interface{}) {
	o.Values = v
}

func (o AlibabaMarketplaceProductExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaMarketplaceProductExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Order) {
		toSerialize["Order"] = o.Order
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Values) {
		toSerialize["Values"] = o.Values
	}
	return toSerialize, nil
}

type NullableAlibabaMarketplaceProductExtra struct {
	value *AlibabaMarketplaceProductExtra
	isSet bool
}

func (v NullableAlibabaMarketplaceProductExtra) Get() *AlibabaMarketplaceProductExtra {
	return v.value
}

func (v *NullableAlibabaMarketplaceProductExtra) Set(val *AlibabaMarketplaceProductExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaMarketplaceProductExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaMarketplaceProductExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaMarketplaceProductExtra(val *AlibabaMarketplaceProductExtra) *NullableAlibabaMarketplaceProductExtra {
	return &NullableAlibabaMarketplaceProductExtra{value: val, isSet: true}
}

func (v NullableAlibabaMarketplaceProductExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaMarketplaceProductExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


