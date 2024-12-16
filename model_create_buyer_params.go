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

// checks if the CreateBuyerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBuyerParams{}

// CreateBuyerParams struct for CreateBuyerParams
type CreateBuyerParams struct {
	// Adyen customerId of this buyer. If not provided but Partner is ADYEN, will create a new customer on Adyen.
	AdyenCustomerId *string `json:"adyenCustomerId,omitempty"`
	// Optional. CompanyInfo of the buyer.
	CompanyInfo *CompanyInfo `json:"companyInfo,omitempty"`
	// The customer ID to recognize the cloud marketplace buyer in your internal system.
	CustomerId *string `json:"customerId,omitempty"`
	// The description of the buyer.
	Description *string `json:"description,omitempty"`
	// Optional. The Lago Customer ID of the buyer.
	LagoCustomerId *string `json:"lagoCustomerId,omitempty"`
	// Optional. The Metronome Customer ID of the buyer.
	MetronomeCustomerId *string `json:"metronomeCustomerId,omitempty"`
	// The name of the buyer.
	Name *string `json:"name,omitempty"`
	// Optional. The Orb Customer ID of the buyer.
	OrbCustomerId *string `json:"orbCustomerId,omitempty"`
	// The channel partner where this buyer is billed. Only STRIPE & ADYEN are supported at the moment.
	Partner *Partner `json:"partner,omitempty"`
	// Payment config for billing.
	PaymentConfig *PaymentConfig `json:"paymentConfig,omitempty"`
	// Stripe customerId of this buyer. If not provided but Partner is STRIPE, will create a new customer on stripe.
	StripeCustomerId *string `json:"stripeCustomerId,omitempty"`
}

// NewCreateBuyerParams instantiates a new CreateBuyerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBuyerParams() *CreateBuyerParams {
	this := CreateBuyerParams{}
	return &this
}

// NewCreateBuyerParamsWithDefaults instantiates a new CreateBuyerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBuyerParamsWithDefaults() *CreateBuyerParams {
	this := CreateBuyerParams{}
	return &this
}

// GetAdyenCustomerId returns the AdyenCustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetAdyenCustomerId() string {
	if o == nil || IsNil(o.AdyenCustomerId) {
		var ret string
		return ret
	}
	return *o.AdyenCustomerId
}

// GetAdyenCustomerIdOk returns a tuple with the AdyenCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetAdyenCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdyenCustomerId) {
		return nil, false
	}
	return o.AdyenCustomerId, true
}

// HasAdyenCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasAdyenCustomerId() bool {
	if o != nil && !IsNil(o.AdyenCustomerId) {
		return true
	}

	return false
}

// SetAdyenCustomerId gets a reference to the given string and assigns it to the AdyenCustomerId field.
func (o *CreateBuyerParams) SetAdyenCustomerId(v string) {
	o.AdyenCustomerId = &v
}

// GetCompanyInfo returns the CompanyInfo field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetCompanyInfo() CompanyInfo {
	if o == nil || IsNil(o.CompanyInfo) {
		var ret CompanyInfo
		return ret
	}
	return *o.CompanyInfo
}

// GetCompanyInfoOk returns a tuple with the CompanyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetCompanyInfoOk() (*CompanyInfo, bool) {
	if o == nil || IsNil(o.CompanyInfo) {
		return nil, false
	}
	return o.CompanyInfo, true
}

// HasCompanyInfo returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasCompanyInfo() bool {
	if o != nil && !IsNil(o.CompanyInfo) {
		return true
	}

	return false
}

// SetCompanyInfo gets a reference to the given CompanyInfo and assigns it to the CompanyInfo field.
func (o *CreateBuyerParams) SetCompanyInfo(v CompanyInfo) {
	o.CompanyInfo = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CreateBuyerParams) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateBuyerParams) SetDescription(v string) {
	o.Description = &v
}

// GetLagoCustomerId returns the LagoCustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetLagoCustomerId() string {
	if o == nil || IsNil(o.LagoCustomerId) {
		var ret string
		return ret
	}
	return *o.LagoCustomerId
}

// GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetLagoCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.LagoCustomerId) {
		return nil, false
	}
	return o.LagoCustomerId, true
}

// HasLagoCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasLagoCustomerId() bool {
	if o != nil && !IsNil(o.LagoCustomerId) {
		return true
	}

	return false
}

// SetLagoCustomerId gets a reference to the given string and assigns it to the LagoCustomerId field.
func (o *CreateBuyerParams) SetLagoCustomerId(v string) {
	o.LagoCustomerId = &v
}

// GetMetronomeCustomerId returns the MetronomeCustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetMetronomeCustomerId() string {
	if o == nil || IsNil(o.MetronomeCustomerId) {
		var ret string
		return ret
	}
	return *o.MetronomeCustomerId
}

// GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetMetronomeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetronomeCustomerId) {
		return nil, false
	}
	return o.MetronomeCustomerId, true
}

// HasMetronomeCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasMetronomeCustomerId() bool {
	if o != nil && !IsNil(o.MetronomeCustomerId) {
		return true
	}

	return false
}

// SetMetronomeCustomerId gets a reference to the given string and assigns it to the MetronomeCustomerId field.
func (o *CreateBuyerParams) SetMetronomeCustomerId(v string) {
	o.MetronomeCustomerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateBuyerParams) SetName(v string) {
	o.Name = &v
}

// GetOrbCustomerId returns the OrbCustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetOrbCustomerId() string {
	if o == nil || IsNil(o.OrbCustomerId) {
		var ret string
		return ret
	}
	return *o.OrbCustomerId
}

// GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetOrbCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrbCustomerId) {
		return nil, false
	}
	return o.OrbCustomerId, true
}

// HasOrbCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasOrbCustomerId() bool {
	if o != nil && !IsNil(o.OrbCustomerId) {
		return true
	}

	return false
}

// SetOrbCustomerId gets a reference to the given string and assigns it to the OrbCustomerId field.
func (o *CreateBuyerParams) SetOrbCustomerId(v string) {
	o.OrbCustomerId = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *CreateBuyerParams) SetPartner(v Partner) {
	o.Partner = &v
}

// GetPaymentConfig returns the PaymentConfig field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetPaymentConfig() PaymentConfig {
	if o == nil || IsNil(o.PaymentConfig) {
		var ret PaymentConfig
		return ret
	}
	return *o.PaymentConfig
}

// GetPaymentConfigOk returns a tuple with the PaymentConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetPaymentConfigOk() (*PaymentConfig, bool) {
	if o == nil || IsNil(o.PaymentConfig) {
		return nil, false
	}
	return o.PaymentConfig, true
}

// HasPaymentConfig returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasPaymentConfig() bool {
	if o != nil && !IsNil(o.PaymentConfig) {
		return true
	}

	return false
}

// SetPaymentConfig gets a reference to the given PaymentConfig and assigns it to the PaymentConfig field.
func (o *CreateBuyerParams) SetPaymentConfig(v PaymentConfig) {
	o.PaymentConfig = &v
}

// GetStripeCustomerId returns the StripeCustomerId field value if set, zero value otherwise.
func (o *CreateBuyerParams) GetStripeCustomerId() string {
	if o == nil || IsNil(o.StripeCustomerId) {
		var ret string
		return ret
	}
	return *o.StripeCustomerId
}

// GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBuyerParams) GetStripeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeCustomerId) {
		return nil, false
	}
	return o.StripeCustomerId, true
}

// HasStripeCustomerId returns a boolean if a field has been set.
func (o *CreateBuyerParams) HasStripeCustomerId() bool {
	if o != nil && !IsNil(o.StripeCustomerId) {
		return true
	}

	return false
}

// SetStripeCustomerId gets a reference to the given string and assigns it to the StripeCustomerId field.
func (o *CreateBuyerParams) SetStripeCustomerId(v string) {
	o.StripeCustomerId = &v
}

func (o CreateBuyerParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBuyerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdyenCustomerId) {
		toSerialize["adyenCustomerId"] = o.AdyenCustomerId
	}
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
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.PaymentConfig) {
		toSerialize["paymentConfig"] = o.PaymentConfig
	}
	if !IsNil(o.StripeCustomerId) {
		toSerialize["stripeCustomerId"] = o.StripeCustomerId
	}
	return toSerialize, nil
}

type NullableCreateBuyerParams struct {
	value *CreateBuyerParams
	isSet bool
}

func (v NullableCreateBuyerParams) Get() *CreateBuyerParams {
	return v.value
}

func (v *NullableCreateBuyerParams) Set(val *CreateBuyerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBuyerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBuyerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBuyerParams(val *CreateBuyerParams) *NullableCreateBuyerParams {
	return &NullableCreateBuyerParams{value: val, isSet: true}
}

func (v NullableCreateBuyerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBuyerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
