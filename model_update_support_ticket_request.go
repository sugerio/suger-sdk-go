/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UpdateSupportTicketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSupportTicketRequest{}

// UpdateSupportTicketRequest struct for UpdateSupportTicketRequest
type UpdateSupportTicketRequest struct {
	Priority SupportTicketPriority `json:"priority"`
	Watchers []string              `json:"watchers"`
}

type _UpdateSupportTicketRequest UpdateSupportTicketRequest

// NewUpdateSupportTicketRequest instantiates a new UpdateSupportTicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSupportTicketRequest(priority SupportTicketPriority, watchers []string) *UpdateSupportTicketRequest {
	this := UpdateSupportTicketRequest{}
	this.Priority = priority
	this.Watchers = watchers
	return &this
}

// NewUpdateSupportTicketRequestWithDefaults instantiates a new UpdateSupportTicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSupportTicketRequestWithDefaults() *UpdateSupportTicketRequest {
	this := UpdateSupportTicketRequest{}
	return &this
}

// GetPriority returns the Priority field value
func (o *UpdateSupportTicketRequest) GetPriority() SupportTicketPriority {
	if o == nil {
		var ret SupportTicketPriority
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *UpdateSupportTicketRequest) GetPriorityOk() (*SupportTicketPriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *UpdateSupportTicketRequest) SetPriority(v SupportTicketPriority) {
	o.Priority = v
}

// GetWatchers returns the Watchers field value
func (o *UpdateSupportTicketRequest) GetWatchers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Watchers
}

// GetWatchersOk returns a tuple with the Watchers field value
// and a boolean to check if the value has been set.
func (o *UpdateSupportTicketRequest) GetWatchersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Watchers, true
}

// SetWatchers sets field value
func (o *UpdateSupportTicketRequest) SetWatchers(v []string) {
	o.Watchers = v
}

func (o UpdateSupportTicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSupportTicketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["priority"] = o.Priority
	toSerialize["watchers"] = o.Watchers
	return toSerialize, nil
}

func (o *UpdateSupportTicketRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"priority",
		"watchers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateSupportTicketRequest := _UpdateSupportTicketRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateSupportTicketRequest)

	if err != nil {
		return err
	}

	*o = UpdateSupportTicketRequest(varUpdateSupportTicketRequest)

	return err
}

type NullableUpdateSupportTicketRequest struct {
	value *UpdateSupportTicketRequest
	isSet bool
}

func (v NullableUpdateSupportTicketRequest) Get() *UpdateSupportTicketRequest {
	return v.value
}

func (v *NullableUpdateSupportTicketRequest) Set(val *UpdateSupportTicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSupportTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSupportTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSupportTicketRequest(val *UpdateSupportTicketRequest) *NullableUpdateSupportTicketRequest {
	return &NullableUpdateSupportTicketRequest{value: val, isSet: true}
}

func (v NullableUpdateSupportTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSupportTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
