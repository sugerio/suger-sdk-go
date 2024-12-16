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

// checks if the AzureMarketplaceCustomAmendmentTenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceCustomAmendmentTenant{}

// AzureMarketplaceCustomAmendmentTenant struct for AzureMarketplaceCustomAmendmentTenant
type AzureMarketplaceCustomAmendmentTenant struct {
	ManualEntries []AzureMarketplaceCustomAmendmentTenantManualEntry `json:"manualEntries,omitempty"`
}

// NewAzureMarketplaceCustomAmendmentTenant instantiates a new AzureMarketplaceCustomAmendmentTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceCustomAmendmentTenant() *AzureMarketplaceCustomAmendmentTenant {
	this := AzureMarketplaceCustomAmendmentTenant{}
	return &this
}

// NewAzureMarketplaceCustomAmendmentTenantWithDefaults instantiates a new AzureMarketplaceCustomAmendmentTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceCustomAmendmentTenantWithDefaults() *AzureMarketplaceCustomAmendmentTenant {
	this := AzureMarketplaceCustomAmendmentTenant{}
	return &this
}

// GetManualEntries returns the ManualEntries field value if set, zero value otherwise.
func (o *AzureMarketplaceCustomAmendmentTenant) GetManualEntries() []AzureMarketplaceCustomAmendmentTenantManualEntry {
	if o == nil || IsNil(o.ManualEntries) {
		var ret []AzureMarketplaceCustomAmendmentTenantManualEntry
		return ret
	}
	return o.ManualEntries
}

// GetManualEntriesOk returns a tuple with the ManualEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceCustomAmendmentTenant) GetManualEntriesOk() ([]AzureMarketplaceCustomAmendmentTenantManualEntry, bool) {
	if o == nil || IsNil(o.ManualEntries) {
		return nil, false
	}
	return o.ManualEntries, true
}

// HasManualEntries returns a boolean if a field has been set.
func (o *AzureMarketplaceCustomAmendmentTenant) HasManualEntries() bool {
	if o != nil && !IsNil(o.ManualEntries) {
		return true
	}

	return false
}

// SetManualEntries gets a reference to the given []AzureMarketplaceCustomAmendmentTenantManualEntry and assigns it to the ManualEntries field.
func (o *AzureMarketplaceCustomAmendmentTenant) SetManualEntries(v []AzureMarketplaceCustomAmendmentTenantManualEntry) {
	o.ManualEntries = v
}

func (o AzureMarketplaceCustomAmendmentTenant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceCustomAmendmentTenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManualEntries) {
		toSerialize["manualEntries"] = o.ManualEntries
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceCustomAmendmentTenant struct {
	value *AzureMarketplaceCustomAmendmentTenant
	isSet bool
}

func (v NullableAzureMarketplaceCustomAmendmentTenant) Get() *AzureMarketplaceCustomAmendmentTenant {
	return v.value
}

func (v *NullableAzureMarketplaceCustomAmendmentTenant) Set(val *AzureMarketplaceCustomAmendmentTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceCustomAmendmentTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceCustomAmendmentTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceCustomAmendmentTenant(val *AzureMarketplaceCustomAmendmentTenant) *NullableAzureMarketplaceCustomAmendmentTenant {
	return &NullableAzureMarketplaceCustomAmendmentTenant{value: val, isSet: true}
}

func (v NullableAzureMarketplaceCustomAmendmentTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceCustomAmendmentTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
