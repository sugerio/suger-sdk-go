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

// checks if the ClientPushMeteringDataResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientPushMeteringDataResponseBody{}

// ClientPushMeteringDataResponseBody struct for ClientPushMeteringDataResponseBody
type ClientPushMeteringDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty"`
}

// NewClientPushMeteringDataResponseBody instantiates a new ClientPushMeteringDataResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPushMeteringDataResponseBody() *ClientPushMeteringDataResponseBody {
	this := ClientPushMeteringDataResponseBody{}
	return &this
}

// NewClientPushMeteringDataResponseBodyWithDefaults instantiates a new ClientPushMeteringDataResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPushMeteringDataResponseBodyWithDefaults() *ClientPushMeteringDataResponseBody {
	this := ClientPushMeteringDataResponseBody{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ClientPushMeteringDataResponseBody) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPushMeteringDataResponseBody) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ClientPushMeteringDataResponseBody) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ClientPushMeteringDataResponseBody) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ClientPushMeteringDataResponseBody) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPushMeteringDataResponseBody) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ClientPushMeteringDataResponseBody) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ClientPushMeteringDataResponseBody) SetSuccess(v bool) {
	o.Success = &v
}

func (o ClientPushMeteringDataResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientPushMeteringDataResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	if !IsNil(o.Success) {
		toSerialize["Success"] = o.Success
	}
	return toSerialize, nil
}

type NullableClientPushMeteringDataResponseBody struct {
	value *ClientPushMeteringDataResponseBody
	isSet bool
}

func (v NullableClientPushMeteringDataResponseBody) Get() *ClientPushMeteringDataResponseBody {
	return v.value
}

func (v *NullableClientPushMeteringDataResponseBody) Set(val *ClientPushMeteringDataResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPushMeteringDataResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPushMeteringDataResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPushMeteringDataResponseBody(val *ClientPushMeteringDataResponseBody) *NullableClientPushMeteringDataResponseBody {
	return &NullableClientPushMeteringDataResponseBody{value: val, isSet: true}
}

func (v NullableClientPushMeteringDataResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPushMeteringDataResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
