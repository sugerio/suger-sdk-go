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

// checks if the UpdateBuyerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBuyerParams{}

// UpdateBuyerParams struct for UpdateBuyerParams
type UpdateBuyerParams struct {
	// Optional. CompanyInfo of the buyer.
	CompanyInfo *CompanyInfo `json:"companyInfo,omitempty"`
	// The customer ID to recognize the cloud marketplace buyer in your internal system. This may be used for uploading CSV files for Batch Metering Usage
	CustomerId *string `json:"customerId,omitempty"`
	// The description of the buyer. If not provided, the description will not be updated.
	Description *string `json:"description,omitempty"`
	// The Lago Customer ID of the buyer. If not provided, the Lago Customer ID will not be updated.
	LagoCustomerId *string `json:"lagoCustomerId,omitempty"`
	// The Metronome Customer ID of the buyer. If not provided, the Metronome Customer ID will not be updated.
	MetronomeCustomerId *string `json:"metronomeCustomerId,omitempty"`
	// The name of the buyer. If not provided, the name will not be updated.
	Name *string `json:"name,omitempty"`
	// The Orb Customer ID of the buyer. If not provided, the Orb Customer ID will not be updated.
	OrbCustomerId *string `json:"orbCustomerId,omitempty"`
	// Optional. PaymentConfig of the buyer. The currency can't be updated.
	PaymentConfig *PaymentConfig `json:"paymentConfig,omitempty"`
	// The Stripe Customer ID of the buyer. If not provided, the Stripe Customer ID will not be updated.
	StripeCustomerId *string `json:"stripeCustomerId,omitempty"`
}

// NewUpdateBuyerParams instantiates a new UpdateBuyerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBuyerParams() *UpdateBuyerParams {
	this := UpdateBuyerParams{}
	return &this
}

// NewUpdateBuyerParamsWithDefaults instantiates a new UpdateBuyerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBuyerParamsWithDefaults() *UpdateBuyerParams {
	this := UpdateBuyerParams{}
	return &this
}

// GetCompanyInfo returns the CompanyInfo field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetCompanyInfo() CompanyInfo {
	if o == nil || IsNil(o.CompanyInfo) {
		var ret CompanyInfo
		return ret
	}
	return *o.CompanyInfo
}

// GetCompanyInfoOk returns a tuple with the CompanyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetCompanyInfoOk() (*CompanyInfo, bool) {
	if o == nil || IsNil(o.CompanyInfo) {
		return nil, false
	}
	return o.CompanyInfo, true
}

// HasCompanyInfo returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasCompanyInfo() bool {
	if o != nil && !IsNil(o.CompanyInfo) {
		return true
	}

	return false
}

// SetCompanyInfo gets a reference to the given CompanyInfo and assigns it to the CompanyInfo field.
func (o *UpdateBuyerParams) SetCompanyInfo(v CompanyInfo) {
	o.CompanyInfo = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *UpdateBuyerParams) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateBuyerParams) SetDescription(v string) {
	o.Description = &v
}

// GetLagoCustomerId returns the LagoCustomerId field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetLagoCustomerId() string {
	if o == nil || IsNil(o.LagoCustomerId) {
		var ret string
		return ret
	}
	return *o.LagoCustomerId
}

// GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetLagoCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoCustomerId) {
		return nil, false
	}
	return o.LagoCustomerId, true
}

// HasLagoCustomerId returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasLagoCustomerId() bool {
	if o != nil && !IsNil(o.LagoCustomerId) {
		return true
	}

	return false
}

// SetLagoCustomerId gets a reference to the given string and assigns it to the LagoCustomerId field.
func (o *UpdateBuyerParams) SetLagoCustomerId(v string) {
	o.LagoCustomerId = &v
}

// GetMetronomeCustomerId returns the MetronomeCustomerId field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetMetronomeCustomerId() string {
	if o == nil || IsNil(o.MetronomeCustomerId) {
		var ret string
		return ret
	}
	return *o.MetronomeCustomerId
}

// GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetMetronomeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetronomeCustomerId) {
		return nil, false
	}
	return o.MetronomeCustomerId, true
}

// HasMetronomeCustomerId returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasMetronomeCustomerId() bool {
	if o != nil && !IsNil(o.MetronomeCustomerId) {
		return true
	}

	return false
}

// SetMetronomeCustomerId gets a reference to the given string and assigns it to the MetronomeCustomerId field.
func (o *UpdateBuyerParams) SetMetronomeCustomerId(v string) {
	o.MetronomeCustomerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateBuyerParams) SetName(v string) {
	o.Name = &v
}

// GetOrbCustomerId returns the OrbCustomerId field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetOrbCustomerId() string {
	if o == nil || IsNil(o.OrbCustomerId) {
		var ret string
		return ret
	}
	return *o.OrbCustomerId
}

// GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetOrbCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrbCustomerId) {
		return nil, false
	}
	return o.OrbCustomerId, true
}

// HasOrbCustomerId returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasOrbCustomerId() bool {
	if o != nil && !IsNil(o.OrbCustomerId) {
		return true
	}

	return false
}

// SetOrbCustomerId gets a reference to the given string and assigns it to the OrbCustomerId field.
func (o *UpdateBuyerParams) SetOrbCustomerId(v string) {
	o.OrbCustomerId = &v
}

// GetPaymentConfig returns the PaymentConfig field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetPaymentConfig() PaymentConfig {
	if o == nil || IsNil(o.PaymentConfig) {
		var ret PaymentConfig
		return ret
	}
	return *o.PaymentConfig
}

// GetPaymentConfigOk returns a tuple with the PaymentConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetPaymentConfigOk() (*PaymentConfig, bool) {
	if o == nil || IsNil(o.PaymentConfig) {
		return nil, false
	}
	return o.PaymentConfig, true
}

// HasPaymentConfig returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasPaymentConfig() bool {
	if o != nil && !IsNil(o.PaymentConfig) {
		return true
	}

	return false
}

// SetPaymentConfig gets a reference to the given PaymentConfig and assigns it to the PaymentConfig field.
func (o *UpdateBuyerParams) SetPaymentConfig(v PaymentConfig) {
	o.PaymentConfig = &v
}

// GetStripeCustomerId returns the StripeCustomerId field value if set, zero value otherwise.
func (o *UpdateBuyerParams) GetStripeCustomerId() string {
	if o == nil || IsNil(o.StripeCustomerId) {
		var ret string
		return ret
	}
	return *o.StripeCustomerId
}

// GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBuyerParams) GetStripeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeCustomerId) {
		return nil, false
	}
	return o.StripeCustomerId, true
}

// HasStripeCustomerId returns a boolean if a field has been set.
func (o *UpdateBuyerParams) HasStripeCustomerId() bool {
	if o != nil && !IsNil(o.StripeCustomerId) {
		return true
	}

	return false
}

// SetStripeCustomerId gets a reference to the given string and assigns it to the StripeCustomerId field.
func (o *UpdateBuyerParams) SetStripeCustomerId(v string) {
	o.StripeCustomerId = &v
}

func (o UpdateBuyerParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBuyerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyInfo) {
		toSerialize["companyInfo"] = o.CompanyInfo
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LagoCustomerId) {
		toSerialize["lagoCustomerId"] = o.LagoCustomerId
	}
	if !IsNil(o.MetronomeCustomerId) {
		toSerialize["metronomeCustomerId"] = o.MetronomeCustomerId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrbCustomerId) {
		toSerialize["orbCustomerId"] = o.OrbCustomerId
	}
	if !IsNil(o.PaymentConfig) {
		toSerialize["paymentConfig"] = o.PaymentConfig
	}
	if !IsNil(o.StripeCustomerId) {
		toSerialize["stripeCustomerId"] = o.StripeCustomerId
	}
	return toSerialize, nil
}

type NullableUpdateBuyerParams struct {
	value *UpdateBuyerParams
	isSet bool
}

func (v NullableUpdateBuyerParams) Get() *UpdateBuyerParams {
	return v.value
}

func (v *NullableUpdateBuyerParams) Set(val *UpdateBuyerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBuyerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBuyerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBuyerParams(val *UpdateBuyerParams) *NullableUpdateBuyerParams {
	return &NullableUpdateBuyerParams{value: val, isSet: true}
}

func (v NullableUpdateBuyerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBuyerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
