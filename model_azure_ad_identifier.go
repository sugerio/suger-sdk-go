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

// checks if the AzureADIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureADIdentifier{}

// AzureADIdentifier struct for AzureADIdentifier
type AzureADIdentifier struct {
	// Azure Billing Account ID
	BillingAccountId *string `json:"billingAccountId,omitempty"`
	CustomerId       *string `json:"customerId,omitempty"`
	// Email address
	EmailId   *string `json:"emailId,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	// Azure License Type
	LicenseType *string `json:"licenseType,omitempty"`
	ObjectId    *string `json:"objectId,omitempty"`
	// ID of the user, used as External ID of suger IdentityBuyer.
	Puid     *string `json:"puid,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}

// NewAzureADIdentifier instantiates a new AzureADIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureADIdentifier() *AzureADIdentifier {
	this := AzureADIdentifier{}
	return &this
}

// NewAzureADIdentifierWithDefaults instantiates a new AzureADIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureADIdentifierWithDefaults() *AzureADIdentifier {
	this := AzureADIdentifier{}
	return &this
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetBillingAccountId() string {
	if o == nil || IsNil(o.BillingAccountId) {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAccountId) {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasBillingAccountId() bool {
	if o != nil && !IsNil(o.BillingAccountId) {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *AzureADIdentifier) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AzureADIdentifier) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *AzureADIdentifier) SetEmailId(v string) {
	o.EmailId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AzureADIdentifier) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AzureADIdentifier) SetLastName(v string) {
	o.LastName = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetLicenseType() string {
	if o == nil || IsNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetLicenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *AzureADIdentifier) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetObjectId() string {
	if o == nil || IsNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *AzureADIdentifier) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPuid returns the Puid field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetPuid() string {
	if o == nil || IsNil(o.Puid) {
		var ret string
		return ret
	}
	return *o.Puid
}

// GetPuidOk returns a tuple with the Puid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetPuidOk() (*string, bool) {
	if o == nil || IsNil(o.Puid) {
		return nil, false
	}
	return o.Puid, true
}

// HasPuid returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasPuid() bool {
	if o != nil && !IsNil(o.Puid) {
		return true
	}

	return false
}

// SetPuid gets a reference to the given string and assigns it to the Puid field.
func (o *AzureADIdentifier) SetPuid(v string) {
	o.Puid = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AzureADIdentifier) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureADIdentifier) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AzureADIdentifier) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AzureADIdentifier) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AzureADIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureADIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAccountId) {
		toSerialize["billingAccountId"] = o.BillingAccountId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !IsNil(o.ObjectId) {
		toSerialize["objectId"] = o.ObjectId
	}
	if !IsNil(o.Puid) {
		toSerialize["puid"] = o.Puid
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	return toSerialize, nil
}

type NullableAzureADIdentifier struct {
	value *AzureADIdentifier
	isSet bool
}

func (v NullableAzureADIdentifier) Get() *AzureADIdentifier {
	return v.value
}

func (v *NullableAzureADIdentifier) Set(val *AzureADIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureADIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureADIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureADIdentifier(val *AzureADIdentifier) *NullableAzureADIdentifier {
	return &NullableAzureADIdentifier{value: val, isSet: true}
}

func (v NullableAzureADIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureADIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
