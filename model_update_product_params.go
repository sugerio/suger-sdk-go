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

// checks if the UpdateProductParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductParams{}

// UpdateProductParams struct for UpdateProductParams
type UpdateProductParams struct {
	FulfillmentUrl string `json:"fulfillmentUrl"`
	Id             string `json:"id"`
	OrganizationID string `json:"organizationID"`
}

type _UpdateProductParams UpdateProductParams

// NewUpdateProductParams instantiates a new UpdateProductParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductParams(fulfillmentUrl string, id string, organizationID string) *UpdateProductParams {
	this := UpdateProductParams{}
	this.FulfillmentUrl = fulfillmentUrl
	this.Id = id
	this.OrganizationID = organizationID
	return &this
}

// NewUpdateProductParamsWithDefaults instantiates a new UpdateProductParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductParamsWithDefaults() *UpdateProductParams {
	this := UpdateProductParams{}
	return &this
}

// GetFulfillmentUrl returns the FulfillmentUrl field value
func (o *UpdateProductParams) GetFulfillmentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentUrl
}

// GetFulfillmentUrlOk returns a tuple with the FulfillmentUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateProductParams) GetFulfillmentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentUrl, true
}

// SetFulfillmentUrl sets field value
func (o *UpdateProductParams) SetFulfillmentUrl(v string) {
	o.FulfillmentUrl = v
}

// GetId returns the Id field value
func (o *UpdateProductParams) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateProductParams) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateProductParams) SetId(v string) {
	o.Id = v
}

// GetOrganizationID returns the OrganizationID field value
func (o *UpdateProductParams) GetOrganizationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value
// and a boolean to check if the value has been set.
func (o *UpdateProductParams) GetOrganizationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationID, true
}

// SetOrganizationID sets field value
func (o *UpdateProductParams) SetOrganizationID(v string) {
	o.OrganizationID = v
}

func (o UpdateProductParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fulfillmentUrl"] = o.FulfillmentUrl
	toSerialize["id"] = o.Id
	toSerialize["organizationID"] = o.OrganizationID
	return toSerialize, nil
}

func (o *UpdateProductParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fulfillmentUrl",
		"id",
		"organizationID",
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

	varUpdateProductParams := _UpdateProductParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateProductParams)

	if err != nil {
		return err
	}

	*o = UpdateProductParams(varUpdateProductParams)

	return err
}

type NullableUpdateProductParams struct {
	value *UpdateProductParams
	isSet bool
}

func (v NullableUpdateProductParams) Get() *UpdateProductParams {
	return v.value
}

func (v *NullableUpdateProductParams) Set(val *UpdateProductParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductParams(val *UpdateProductParams) *NullableUpdateProductParams {
	return &NullableUpdateProductParams{value: val, isSet: true}
}

func (v NullableUpdateProductParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
