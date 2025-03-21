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

// checks if the GcpMarketplaceAgreementDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceAgreementDocument{}

// GcpMarketplaceAgreementDocument struct for GcpMarketplaceAgreementDocument
type GcpMarketplaceAgreementDocument struct {
	EulaAgreementDocument *GcpMarketplaceDocument `json:"eulaAgreementDocument,omitempty"`
}

// NewGcpMarketplaceAgreementDocument instantiates a new GcpMarketplaceAgreementDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceAgreementDocument() *GcpMarketplaceAgreementDocument {
	this := GcpMarketplaceAgreementDocument{}
	return &this
}

// NewGcpMarketplaceAgreementDocumentWithDefaults instantiates a new GcpMarketplaceAgreementDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceAgreementDocumentWithDefaults() *GcpMarketplaceAgreementDocument {
	this := GcpMarketplaceAgreementDocument{}
	return &this
}

// GetEulaAgreementDocument returns the EulaAgreementDocument field value if set, zero value otherwise.
func (o *GcpMarketplaceAgreementDocument) GetEulaAgreementDocument() GcpMarketplaceDocument {
	if o == nil || IsNil(o.EulaAgreementDocument) {
		var ret GcpMarketplaceDocument
		return ret
	}
	return *o.EulaAgreementDocument
}

// GetEulaAgreementDocumentOk returns a tuple with the EulaAgreementDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceAgreementDocument) GetEulaAgreementDocumentOk() (*GcpMarketplaceDocument, bool) {
	if o == nil || IsNil(o.EulaAgreementDocument) {
		return nil, false
	}
	return o.EulaAgreementDocument, true
}

// HasEulaAgreementDocument returns a boolean if a field has been set.
func (o *GcpMarketplaceAgreementDocument) HasEulaAgreementDocument() bool {
	if o != nil && !IsNil(o.EulaAgreementDocument) {
		return true
	}

	return false
}

// SetEulaAgreementDocument gets a reference to the given GcpMarketplaceDocument and assigns it to the EulaAgreementDocument field.
func (o *GcpMarketplaceAgreementDocument) SetEulaAgreementDocument(v GcpMarketplaceDocument) {
	o.EulaAgreementDocument = &v
}

func (o GcpMarketplaceAgreementDocument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceAgreementDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EulaAgreementDocument) {
		toSerialize["eulaAgreementDocument"] = o.EulaAgreementDocument
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceAgreementDocument struct {
	value *GcpMarketplaceAgreementDocument
	isSet bool
}

func (v NullableGcpMarketplaceAgreementDocument) Get() *GcpMarketplaceAgreementDocument {
	return v.value
}

func (v *NullableGcpMarketplaceAgreementDocument) Set(val *GcpMarketplaceAgreementDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceAgreementDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceAgreementDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceAgreementDocument(val *GcpMarketplaceAgreementDocument) *NullableGcpMarketplaceAgreementDocument {
	return &NullableGcpMarketplaceAgreementDocument{value: val, isSet: true}
}

func (v NullableGcpMarketplaceAgreementDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceAgreementDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
