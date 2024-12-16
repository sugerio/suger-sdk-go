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

// checks if the AzureMarketplacePriceAndAvailabilitySoftwareReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilitySoftwareReservation{}

// AzureMarketplacePriceAndAvailabilitySoftwareReservation struct for AzureMarketplacePriceAndAvailabilitySoftwareReservation
type AzureMarketplacePriceAndAvailabilitySoftwareReservation struct {
	// default 0
	PercentageSave *float32 `json:"percentageSave,omitempty"`
	// default 0
	Term *float32 `json:"term,omitempty"`
	Type *string  `json:"type,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilitySoftwareReservation instantiates a new AzureMarketplacePriceAndAvailabilitySoftwareReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilitySoftwareReservation() *AzureMarketplacePriceAndAvailabilitySoftwareReservation {
	this := AzureMarketplacePriceAndAvailabilitySoftwareReservation{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilitySoftwareReservationWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilitySoftwareReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilitySoftwareReservationWithDefaults() *AzureMarketplacePriceAndAvailabilitySoftwareReservation {
	this := AzureMarketplacePriceAndAvailabilitySoftwareReservation{}
	return &this
}

// GetPercentageSave returns the PercentageSave field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetPercentageSave() float32 {
	if o == nil || IsNil(o.PercentageSave) {
		var ret float32
		return ret
	}
	return *o.PercentageSave
}

// GetPercentageSaveOk returns a tuple with the PercentageSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetPercentageSaveOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentageSave) {
		return nil, false
	}
	return o.PercentageSave, true
}

// HasPercentageSave returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) HasPercentageSave() bool {
	if o != nil && !IsNil(o.PercentageSave) {
		return true
	}

	return false
}

// SetPercentageSave gets a reference to the given float32 and assigns it to the PercentageSave field.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) SetPercentageSave(v float32) {
	o.PercentageSave = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetTerm() float32 {
	if o == nil || IsNil(o.Term) {
		var ret float32
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetTermOk() (*float32, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given float32 and assigns it to the Term field.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) SetTerm(v float32) {
	o.Term = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureMarketplacePriceAndAvailabilitySoftwareReservation) SetType(v string) {
	o.Type = &v
}

func (o AzureMarketplacePriceAndAvailabilitySoftwareReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilitySoftwareReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PercentageSave) {
		toSerialize["percentageSave"] = o.PercentageSave
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation struct {
	value *AzureMarketplacePriceAndAvailabilitySoftwareReservation
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) Get() *AzureMarketplacePriceAndAvailabilitySoftwareReservation {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) Set(val *AzureMarketplacePriceAndAvailabilitySoftwareReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilitySoftwareReservation(val *AzureMarketplacePriceAndAvailabilitySoftwareReservation) *NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation {
	return &NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySoftwareReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
