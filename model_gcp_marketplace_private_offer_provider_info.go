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

// checks if the GcpMarketplacePrivateOfferProviderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplacePrivateOfferProviderInfo{}

// GcpMarketplacePrivateOfferProviderInfo struct for GcpMarketplacePrivateOfferProviderInfo
type GcpMarketplacePrivateOfferProviderInfo struct {
	// The email address of who create the private offer in the provider.
	CreatorEmailAddress *string `json:"creatorEmailAddress,omitempty"`
	// The sales contact email of the provider.
	SalesContactEmail *string `json:"salesContactEmail,omitempty"`
	// The sales contact name of the provider.
	SalesContactName *string `json:"salesContactName,omitempty"`
}

// NewGcpMarketplacePrivateOfferProviderInfo instantiates a new GcpMarketplacePrivateOfferProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplacePrivateOfferProviderInfo() *GcpMarketplacePrivateOfferProviderInfo {
	this := GcpMarketplacePrivateOfferProviderInfo{}
	return &this
}

// NewGcpMarketplacePrivateOfferProviderInfoWithDefaults instantiates a new GcpMarketplacePrivateOfferProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplacePrivateOfferProviderInfoWithDefaults() *GcpMarketplacePrivateOfferProviderInfo {
	this := GcpMarketplacePrivateOfferProviderInfo{}
	return &this
}

// GetCreatorEmailAddress returns the CreatorEmailAddress field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetCreatorEmailAddress() string {
	if o == nil || IsNil(o.CreatorEmailAddress) {
		var ret string
		return ret
	}
	return *o.CreatorEmailAddress
}

// GetCreatorEmailAddressOk returns a tuple with the CreatorEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetCreatorEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorEmailAddress) {
		return nil, false
	}
	return o.CreatorEmailAddress, true
}

// HasCreatorEmailAddress returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) HasCreatorEmailAddress() bool {
	if o != nil && !IsNil(o.CreatorEmailAddress) {
		return true
	}

	return false
}

// SetCreatorEmailAddress gets a reference to the given string and assigns it to the CreatorEmailAddress field.
func (o *GcpMarketplacePrivateOfferProviderInfo) SetCreatorEmailAddress(v string) {
	o.CreatorEmailAddress = &v
}

// GetSalesContactEmail returns the SalesContactEmail field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetSalesContactEmail() string {
	if o == nil || IsNil(o.SalesContactEmail) {
		var ret string
		return ret
	}
	return *o.SalesContactEmail
}

// GetSalesContactEmailOk returns a tuple with the SalesContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetSalesContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SalesContactEmail) {
		return nil, false
	}
	return o.SalesContactEmail, true
}

// HasSalesContactEmail returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) HasSalesContactEmail() bool {
	if o != nil && !IsNil(o.SalesContactEmail) {
		return true
	}

	return false
}

// SetSalesContactEmail gets a reference to the given string and assigns it to the SalesContactEmail field.
func (o *GcpMarketplacePrivateOfferProviderInfo) SetSalesContactEmail(v string) {
	o.SalesContactEmail = &v
}

// GetSalesContactName returns the SalesContactName field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetSalesContactName() string {
	if o == nil || IsNil(o.SalesContactName) {
		var ret string
		return ret
	}
	return *o.SalesContactName
}

// GetSalesContactNameOk returns a tuple with the SalesContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) GetSalesContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.SalesContactName) {
		return nil, false
	}
	return o.SalesContactName, true
}

// HasSalesContactName returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferProviderInfo) HasSalesContactName() bool {
	if o != nil && !IsNil(o.SalesContactName) {
		return true
	}

	return false
}

// SetSalesContactName gets a reference to the given string and assigns it to the SalesContactName field.
func (o *GcpMarketplacePrivateOfferProviderInfo) SetSalesContactName(v string) {
	o.SalesContactName = &v
}

func (o GcpMarketplacePrivateOfferProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplacePrivateOfferProviderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatorEmailAddress) {
		toSerialize["creatorEmailAddress"] = o.CreatorEmailAddress
	}
	if !IsNil(o.SalesContactEmail) {
		toSerialize["salesContactEmail"] = o.SalesContactEmail
	}
	if !IsNil(o.SalesContactName) {
		toSerialize["salesContactName"] = o.SalesContactName
	}
	return toSerialize, nil
}

type NullableGcpMarketplacePrivateOfferProviderInfo struct {
	value *GcpMarketplacePrivateOfferProviderInfo
	isSet bool
}

func (v NullableGcpMarketplacePrivateOfferProviderInfo) Get() *GcpMarketplacePrivateOfferProviderInfo {
	return v.value
}

func (v *NullableGcpMarketplacePrivateOfferProviderInfo) Set(val *GcpMarketplacePrivateOfferProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplacePrivateOfferProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplacePrivateOfferProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplacePrivateOfferProviderInfo(val *GcpMarketplacePrivateOfferProviderInfo) *NullableGcpMarketplacePrivateOfferProviderInfo {
	return &NullableGcpMarketplacePrivateOfferProviderInfo{value: val, isSet: true}
}

func (v NullableGcpMarketplacePrivateOfferProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplacePrivateOfferProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
