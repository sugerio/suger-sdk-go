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

// checks if the AzureMarketplacePrivateOfferPartner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePrivateOfferPartner{}

// AzureMarketplacePrivateOfferPartner struct for AzureMarketplacePrivateOfferPartner
type AzureMarketplacePrivateOfferPartner struct {
	Id          *string `json:"id,omitempty"`
	Location    *string `json:"location,omitempty"`
	PartnerName *string `json:"partnerName,omitempty"`
}

// NewAzureMarketplacePrivateOfferPartner instantiates a new AzureMarketplacePrivateOfferPartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePrivateOfferPartner() *AzureMarketplacePrivateOfferPartner {
	this := AzureMarketplacePrivateOfferPartner{}
	return &this
}

// NewAzureMarketplacePrivateOfferPartnerWithDefaults instantiates a new AzureMarketplacePrivateOfferPartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePrivateOfferPartnerWithDefaults() *AzureMarketplacePrivateOfferPartner {
	this := AzureMarketplacePrivateOfferPartner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOfferPartner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOfferPartner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOfferPartner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplacePrivateOfferPartner) SetId(v string) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOfferPartner) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOfferPartner) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOfferPartner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AzureMarketplacePrivateOfferPartner) SetLocation(v string) {
	o.Location = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOfferPartner) GetPartnerName() string {
	if o == nil || IsNil(o.PartnerName) {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOfferPartner) GetPartnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerName) {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOfferPartner) HasPartnerName() bool {
	if o != nil && !IsNil(o.PartnerName) {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *AzureMarketplacePrivateOfferPartner) SetPartnerName(v string) {
	o.PartnerName = &v
}

func (o AzureMarketplacePrivateOfferPartner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePrivateOfferPartner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.PartnerName) {
		toSerialize["partnerName"] = o.PartnerName
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePrivateOfferPartner struct {
	value *AzureMarketplacePrivateOfferPartner
	isSet bool
}

func (v NullableAzureMarketplacePrivateOfferPartner) Get() *AzureMarketplacePrivateOfferPartner {
	return v.value
}

func (v *NullableAzureMarketplacePrivateOfferPartner) Set(val *AzureMarketplacePrivateOfferPartner) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePrivateOfferPartner) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePrivateOfferPartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePrivateOfferPartner(val *AzureMarketplacePrivateOfferPartner) *NullableAzureMarketplacePrivateOfferPartner {
	return &NullableAzureMarketplacePrivateOfferPartner{value: val, isSet: true}
}

func (v NullableAzureMarketplacePrivateOfferPartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePrivateOfferPartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
