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

// checks if the AzureMarketplacePrivateOfferBeneficiaryRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePrivateOfferBeneficiaryRecipient{}

// AzureMarketplacePrivateOfferBeneficiaryRecipient struct for AzureMarketplacePrivateOfferBeneficiaryRecipient
type AzureMarketplacePrivateOfferBeneficiaryRecipient struct {
	Id *string `json:"id,omitempty"`
	RecipientType *string `json:"recipientType,omitempty"`
}

// NewAzureMarketplacePrivateOfferBeneficiaryRecipient instantiates a new AzureMarketplacePrivateOfferBeneficiaryRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePrivateOfferBeneficiaryRecipient() *AzureMarketplacePrivateOfferBeneficiaryRecipient {
	this := AzureMarketplacePrivateOfferBeneficiaryRecipient{}
	return &this
}

// NewAzureMarketplacePrivateOfferBeneficiaryRecipientWithDefaults instantiates a new AzureMarketplacePrivateOfferBeneficiaryRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePrivateOfferBeneficiaryRecipientWithDefaults() *AzureMarketplacePrivateOfferBeneficiaryRecipient {
	this := AzureMarketplacePrivateOfferBeneficiaryRecipient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) SetId(v string) {
	o.Id = &v
}

// GetRecipientType returns the RecipientType field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) GetRecipientType() string {
	if o == nil || IsNil(o.RecipientType) {
		var ret string
		return ret
	}
	return *o.RecipientType
}

// GetRecipientTypeOk returns a tuple with the RecipientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) GetRecipientTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientType) {
		return nil, false
	}
	return o.RecipientType, true
}

// HasRecipientType returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) HasRecipientType() bool {
	if o != nil && !IsNil(o.RecipientType) {
		return true
	}

	return false
}

// SetRecipientType gets a reference to the given string and assigns it to the RecipientType field.
func (o *AzureMarketplacePrivateOfferBeneficiaryRecipient) SetRecipientType(v string) {
	o.RecipientType = &v
}

func (o AzureMarketplacePrivateOfferBeneficiaryRecipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePrivateOfferBeneficiaryRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecipientType) {
		toSerialize["recipientType"] = o.RecipientType
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePrivateOfferBeneficiaryRecipient struct {
	value *AzureMarketplacePrivateOfferBeneficiaryRecipient
	isSet bool
}

func (v NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) Get() *AzureMarketplacePrivateOfferBeneficiaryRecipient {
	return v.value
}

func (v *NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) Set(val *AzureMarketplacePrivateOfferBeneficiaryRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePrivateOfferBeneficiaryRecipient(val *AzureMarketplacePrivateOfferBeneficiaryRecipient) *NullableAzureMarketplacePrivateOfferBeneficiaryRecipient {
	return &NullableAzureMarketplacePrivateOfferBeneficiaryRecipient{value: val, isSet: true}
}

func (v NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePrivateOfferBeneficiaryRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


