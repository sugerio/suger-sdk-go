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

// checks if the LastModifiedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastModifiedBy{}

// LastModifiedBy struct for LastModifiedBy
type LastModifiedBy struct {
	// The email of the creator.
	Email *string `json:"email,omitempty"`
	// The ID of the creator.
	EntityId *string `json:"entityId,omitempty"`
	// The Entity type of the creator, either USER or API_CLIENT.
	EntityType *EntityType `json:"entityType,omitempty"`
	// The name of the creator.
	Name *string `json:"name,omitempty"`
}

// NewLastModifiedBy instantiates a new LastModifiedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastModifiedBy() *LastModifiedBy {
	this := LastModifiedBy{}
	return &this
}

// NewLastModifiedByWithDefaults instantiates a new LastModifiedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastModifiedByWithDefaults() *LastModifiedBy {
	this := LastModifiedBy{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LastModifiedBy) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedBy) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LastModifiedBy) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LastModifiedBy) SetEmail(v string) {
	o.Email = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *LastModifiedBy) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedBy) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *LastModifiedBy) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *LastModifiedBy) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *LastModifiedBy) GetEntityType() EntityType {
	if o == nil || IsNil(o.EntityType) {
		var ret EntityType
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedBy) GetEntityTypeOk() (*EntityType, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *LastModifiedBy) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given EntityType and assigns it to the EntityType field.
func (o *LastModifiedBy) SetEntityType(v EntityType) {
	o.EntityType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LastModifiedBy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastModifiedBy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LastModifiedBy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LastModifiedBy) SetName(v string) {
	o.Name = &v
}

func (o LastModifiedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastModifiedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLastModifiedBy struct {
	value *LastModifiedBy
	isSet bool
}

func (v NullableLastModifiedBy) Get() *LastModifiedBy {
	return v.value
}

func (v *NullableLastModifiedBy) Set(val *LastModifiedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableLastModifiedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableLastModifiedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastModifiedBy(val *LastModifiedBy) *NullableLastModifiedBy {
	return &NullableLastModifiedBy{value: val, isSet: true}
}

func (v NullableLastModifiedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastModifiedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}