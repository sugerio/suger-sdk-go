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

// checks if the GcpMarketplaceResellerPrivateOfferPlanNewState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlanNewState{}

// GcpMarketplaceResellerPrivateOfferPlanNewState struct for GcpMarketplaceResellerPrivateOfferPlanNewState
type GcpMarketplaceResellerPrivateOfferPlanNewState struct {
	StateType *GcpMarketplaceResellerPrivateOfferPlanStateType `json:"stateType,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlanNewState instantiates a new GcpMarketplaceResellerPrivateOfferPlanNewState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlanNewState() *GcpMarketplaceResellerPrivateOfferPlanNewState {
	this := GcpMarketplaceResellerPrivateOfferPlanNewState{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanNewStateWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanNewState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanNewStateWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanNewState {
	this := GcpMarketplaceResellerPrivateOfferPlanNewState{}
	return &this
}

// GetStateType returns the StateType field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanNewState) GetStateType() GcpMarketplaceResellerPrivateOfferPlanStateType {
	if o == nil || IsNil(o.StateType) {
		var ret GcpMarketplaceResellerPrivateOfferPlanStateType
		return ret
	}
	return *o.StateType
}

// GetStateTypeOk returns a tuple with the StateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanNewState) GetStateTypeOk() (*GcpMarketplaceResellerPrivateOfferPlanStateType, bool) {
	if o == nil || IsNil(o.StateType) {
		return nil, false
	}
	return o.StateType, true
}

// HasStateType returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanNewState) HasStateType() bool {
	if o != nil && !IsNil(o.StateType) {
		return true
	}

	return false
}

// SetStateType gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanStateType and assigns it to the StateType field.
func (o *GcpMarketplaceResellerPrivateOfferPlanNewState) SetStateType(v GcpMarketplaceResellerPrivateOfferPlanStateType) {
	o.StateType = &v
}

func (o GcpMarketplaceResellerPrivateOfferPlanNewState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlanNewState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StateType) {
		toSerialize["stateType"] = o.StateType
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlanNewState struct {
	value *GcpMarketplaceResellerPrivateOfferPlanNewState
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanNewState) Get() *GcpMarketplaceResellerPrivateOfferPlanNewState {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanNewState) Set(val *GcpMarketplaceResellerPrivateOfferPlanNewState) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanNewState) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanNewState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlanNewState(val *GcpMarketplaceResellerPrivateOfferPlanNewState) *NullableGcpMarketplaceResellerPrivateOfferPlanNewState {
	return &NullableGcpMarketplaceResellerPrivateOfferPlanNewState{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanNewState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanNewState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
