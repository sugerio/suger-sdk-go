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

// checks if the AzureMarketplaceCustomAmendment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceCustomAmendment{}

// AzureMarketplaceCustomAmendment struct for AzureMarketplaceCustomAmendment
type AzureMarketplaceCustomAmendment struct {
	Tenants *AzureMarketplaceCustomAmendmentTenant `json:"tenants,omitempty"`
	Terms   *string                                `json:"terms,omitempty"`
}

// NewAzureMarketplaceCustomAmendment instantiates a new AzureMarketplaceCustomAmendment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceCustomAmendment() *AzureMarketplaceCustomAmendment {
	this := AzureMarketplaceCustomAmendment{}
	return &this
}

// NewAzureMarketplaceCustomAmendmentWithDefaults instantiates a new AzureMarketplaceCustomAmendment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceCustomAmendmentWithDefaults() *AzureMarketplaceCustomAmendment {
	this := AzureMarketplaceCustomAmendment{}
	return &this
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *AzureMarketplaceCustomAmendment) GetTenants() AzureMarketplaceCustomAmendmentTenant {
	if o == nil || IsNil(o.Tenants) {
		var ret AzureMarketplaceCustomAmendmentTenant
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceCustomAmendment) GetTenantsOk() (*AzureMarketplaceCustomAmendmentTenant, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *AzureMarketplaceCustomAmendment) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given AzureMarketplaceCustomAmendmentTenant and assigns it to the Tenants field.
func (o *AzureMarketplaceCustomAmendment) SetTenants(v AzureMarketplaceCustomAmendmentTenant) {
	o.Tenants = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *AzureMarketplaceCustomAmendment) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceCustomAmendment) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *AzureMarketplaceCustomAmendment) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *AzureMarketplaceCustomAmendment) SetTerms(v string) {
	o.Terms = &v
}

func (o AzureMarketplaceCustomAmendment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceCustomAmendment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceCustomAmendment struct {
	value *AzureMarketplaceCustomAmendment
	isSet bool
}

func (v NullableAzureMarketplaceCustomAmendment) Get() *AzureMarketplaceCustomAmendment {
	return v.value
}

func (v *NullableAzureMarketplaceCustomAmendment) Set(val *AzureMarketplaceCustomAmendment) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceCustomAmendment) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceCustomAmendment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceCustomAmendment(val *AzureMarketplaceCustomAmendment) *NullableAzureMarketplaceCustomAmendment {
	return &NullableAzureMarketplaceCustomAmendment{value: val, isSet: true}
}

func (v NullableAzureMarketplaceCustomAmendment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceCustomAmendment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
