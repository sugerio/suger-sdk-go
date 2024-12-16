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

// checks if the ListRevenueRecordDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRevenueRecordDetailsResponse{}

// ListRevenueRecordDetailsResponse struct for ListRevenueRecordDetailsResponse
type ListRevenueRecordDetailsResponse struct {
	NextOffset           *int32                `json:"nextOffset,omitempty"`
	RevenueRecordDetails []RevenueRecordDetail `json:"revenueRecordDetails,omitempty"`
}

// NewListRevenueRecordDetailsResponse instantiates a new ListRevenueRecordDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRevenueRecordDetailsResponse() *ListRevenueRecordDetailsResponse {
	this := ListRevenueRecordDetailsResponse{}
	return &this
}

// NewListRevenueRecordDetailsResponseWithDefaults instantiates a new ListRevenueRecordDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRevenueRecordDetailsResponseWithDefaults() *ListRevenueRecordDetailsResponse {
	this := ListRevenueRecordDetailsResponse{}
	return &this
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *ListRevenueRecordDetailsResponse) GetNextOffset() int32 {
	if o == nil || IsNil(o.NextOffset) {
		var ret int32
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRevenueRecordDetailsResponse) GetNextOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.NextOffset) {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *ListRevenueRecordDetailsResponse) HasNextOffset() bool {
	if o != nil && !IsNil(o.NextOffset) {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given int32 and assigns it to the NextOffset field.
func (o *ListRevenueRecordDetailsResponse) SetNextOffset(v int32) {
	o.NextOffset = &v
}

// GetRevenueRecordDetails returns the RevenueRecordDetails field value if set, zero value otherwise.
func (o *ListRevenueRecordDetailsResponse) GetRevenueRecordDetails() []RevenueRecordDetail {
	if o == nil || IsNil(o.RevenueRecordDetails) {
		var ret []RevenueRecordDetail
		return ret
	}
	return o.RevenueRecordDetails
}

// GetRevenueRecordDetailsOk returns a tuple with the RevenueRecordDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRevenueRecordDetailsResponse) GetRevenueRecordDetailsOk() ([]RevenueRecordDetail, bool) {
	if o == nil || IsNil(o.RevenueRecordDetails) {
		return nil, false
	}
	return o.RevenueRecordDetails, true
}

// HasRevenueRecordDetails returns a boolean if a field has been set.
func (o *ListRevenueRecordDetailsResponse) HasRevenueRecordDetails() bool {
	if o != nil && !IsNil(o.RevenueRecordDetails) {
		return true
	}

	return false
}

// SetRevenueRecordDetails gets a reference to the given []RevenueRecordDetail and assigns it to the RevenueRecordDetails field.
func (o *ListRevenueRecordDetailsResponse) SetRevenueRecordDetails(v []RevenueRecordDetail) {
	o.RevenueRecordDetails = v
}

func (o ListRevenueRecordDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRevenueRecordDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextOffset) {
		toSerialize["nextOffset"] = o.NextOffset
	}
	if !IsNil(o.RevenueRecordDetails) {
		toSerialize["revenueRecordDetails"] = o.RevenueRecordDetails
	}
	return toSerialize, nil
}

type NullableListRevenueRecordDetailsResponse struct {
	value *ListRevenueRecordDetailsResponse
	isSet bool
}

func (v NullableListRevenueRecordDetailsResponse) Get() *ListRevenueRecordDetailsResponse {
	return v.value
}

func (v *NullableListRevenueRecordDetailsResponse) Set(val *ListRevenueRecordDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRevenueRecordDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRevenueRecordDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRevenueRecordDetailsResponse(val *ListRevenueRecordDetailsResponse) *NullableListRevenueRecordDetailsResponse {
	return &NullableListRevenueRecordDetailsResponse{value: val, isSet: true}
}

func (v NullableListRevenueRecordDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRevenueRecordDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
