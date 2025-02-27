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

// checks if the GcpMarketplacePrivateOfferCustomerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplacePrivateOfferCustomerInfo{}

// GcpMarketplacePrivateOfferCustomerInfo struct for GcpMarketplacePrivateOfferCustomerInfo
type GcpMarketplacePrivateOfferCustomerInfo struct {
	// The address of the customer
	Address *string `json:"address,omitempty"`
	// The contact name of the customer
	Contact *string `json:"contact,omitempty"`
	// The email address of the customer
	Email *string `json:"email,omitempty"`
	// The company name of the customer
	Organization *string `json:"organization,omitempty"`
	// The GCP billing account ID of the customer
	UnverifiedBillingAccount *string `json:"unverifiedBillingAccount,omitempty"`
}

// NewGcpMarketplacePrivateOfferCustomerInfo instantiates a new GcpMarketplacePrivateOfferCustomerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplacePrivateOfferCustomerInfo() *GcpMarketplacePrivateOfferCustomerInfo {
	this := GcpMarketplacePrivateOfferCustomerInfo{}
	return &this
}

// NewGcpMarketplacePrivateOfferCustomerInfoWithDefaults instantiates a new GcpMarketplacePrivateOfferCustomerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplacePrivateOfferCustomerInfoWithDefaults() *GcpMarketplacePrivateOfferCustomerInfo {
	this := GcpMarketplacePrivateOfferCustomerInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GcpMarketplacePrivateOfferCustomerInfo) SetAddress(v string) {
	o.Address = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *GcpMarketplacePrivateOfferCustomerInfo) SetContact(v string) {
	o.Contact = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GcpMarketplacePrivateOfferCustomerInfo) SetEmail(v string) {
	o.Email = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *GcpMarketplacePrivateOfferCustomerInfo) SetOrganization(v string) {
	o.Organization = &v
}

// GetUnverifiedBillingAccount returns the UnverifiedBillingAccount field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetUnverifiedBillingAccount() string {
	if o == nil || IsNil(o.UnverifiedBillingAccount) {
		var ret string
		return ret
	}
	return *o.UnverifiedBillingAccount
}

// GetUnverifiedBillingAccountOk returns a tuple with the UnverifiedBillingAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) GetUnverifiedBillingAccountOk() (*string, bool) {
	if o == nil || IsNil(o.UnverifiedBillingAccount) {
		return nil, false
	}
	return o.UnverifiedBillingAccount, true
}

// HasUnverifiedBillingAccount returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferCustomerInfo) HasUnverifiedBillingAccount() bool {
	if o != nil && !IsNil(o.UnverifiedBillingAccount) {
		return true
	}

	return false
}

// SetUnverifiedBillingAccount gets a reference to the given string and assigns it to the UnverifiedBillingAccount field.
func (o *GcpMarketplacePrivateOfferCustomerInfo) SetUnverifiedBillingAccount(v string) {
	o.UnverifiedBillingAccount = &v
}

func (o GcpMarketplacePrivateOfferCustomerInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplacePrivateOfferCustomerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.UnverifiedBillingAccount) {
		toSerialize["unverifiedBillingAccount"] = o.UnverifiedBillingAccount
	}
	return toSerialize, nil
}

type NullableGcpMarketplacePrivateOfferCustomerInfo struct {
	value *GcpMarketplacePrivateOfferCustomerInfo
	isSet bool
}

func (v NullableGcpMarketplacePrivateOfferCustomerInfo) Get() *GcpMarketplacePrivateOfferCustomerInfo {
	return v.value
}

func (v *NullableGcpMarketplacePrivateOfferCustomerInfo) Set(val *GcpMarketplacePrivateOfferCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplacePrivateOfferCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplacePrivateOfferCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplacePrivateOfferCustomerInfo(val *GcpMarketplacePrivateOfferCustomerInfo) *NullableGcpMarketplacePrivateOfferCustomerInfo {
	return &NullableGcpMarketplacePrivateOfferCustomerInfo{value: val, isSet: true}
}

func (v NullableGcpMarketplacePrivateOfferCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplacePrivateOfferCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
