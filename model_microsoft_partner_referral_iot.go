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

// checks if the MicrosoftPartnerReferralIot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftPartnerReferralIot{}

// MicrosoftPartnerReferralIot struct for MicrosoftPartnerReferralIot
type MicrosoftPartnerReferralIot struct {
	AttachServices                 *bool                  `json:"attachServices,omitempty"`
	AzureCertifiedDevice           *bool                  `json:"azureCertifiedDevice,omitempty"`
	CustomerLicenseAgreementNumber map[string]interface{} `json:"customerLicenseAgreementNumber,omitempty"`
	DeviceCategory                 map[string]interface{} `json:"deviceCategory,omitempty"`
	SiliconType                    map[string]interface{} `json:"siliconType,omitempty"`
}

// NewMicrosoftPartnerReferralIot instantiates a new MicrosoftPartnerReferralIot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftPartnerReferralIot() *MicrosoftPartnerReferralIot {
	this := MicrosoftPartnerReferralIot{}
	return &this
}

// NewMicrosoftPartnerReferralIotWithDefaults instantiates a new MicrosoftPartnerReferralIot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftPartnerReferralIotWithDefaults() *MicrosoftPartnerReferralIot {
	this := MicrosoftPartnerReferralIot{}
	return &this
}

// GetAttachServices returns the AttachServices field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralIot) GetAttachServices() bool {
	if o == nil || IsNil(o.AttachServices) {
		var ret bool
		return ret
	}
	return *o.AttachServices
}

// GetAttachServicesOk returns a tuple with the AttachServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralIot) GetAttachServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.AttachServices) {
		return nil, false
	}
	return o.AttachServices, true
}

// HasAttachServices returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralIot) HasAttachServices() bool {
	if o != nil && !IsNil(o.AttachServices) {
		return true
	}

	return false
}

// SetAttachServices gets a reference to the given bool and assigns it to the AttachServices field.
func (o *MicrosoftPartnerReferralIot) SetAttachServices(v bool) {
	o.AttachServices = &v
}

// GetAzureCertifiedDevice returns the AzureCertifiedDevice field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralIot) GetAzureCertifiedDevice() bool {
	if o == nil || IsNil(o.AzureCertifiedDevice) {
		var ret bool
		return ret
	}
	return *o.AzureCertifiedDevice
}

// GetAzureCertifiedDeviceOk returns a tuple with the AzureCertifiedDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralIot) GetAzureCertifiedDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.AzureCertifiedDevice) {
		return nil, false
	}
	return o.AzureCertifiedDevice, true
}

// HasAzureCertifiedDevice returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralIot) HasAzureCertifiedDevice() bool {
	if o != nil && !IsNil(o.AzureCertifiedDevice) {
		return true
	}

	return false
}

// SetAzureCertifiedDevice gets a reference to the given bool and assigns it to the AzureCertifiedDevice field.
func (o *MicrosoftPartnerReferralIot) SetAzureCertifiedDevice(v bool) {
	o.AzureCertifiedDevice = &v
}

// GetCustomerLicenseAgreementNumber returns the CustomerLicenseAgreementNumber field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralIot) GetCustomerLicenseAgreementNumber() map[string]interface{} {
	if o == nil || IsNil(o.CustomerLicenseAgreementNumber) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomerLicenseAgreementNumber
}

// GetCustomerLicenseAgreementNumberOk returns a tuple with the CustomerLicenseAgreementNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralIot) GetCustomerLicenseAgreementNumberOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomerLicenseAgreementNumber) {
		return map[string]interface{}{}, false
	}
	return o.CustomerLicenseAgreementNumber, true
}

// HasCustomerLicenseAgreementNumber returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralIot) HasCustomerLicenseAgreementNumber() bool {
	if o != nil && !IsNil(o.CustomerLicenseAgreementNumber) {
		return true
	}

	return false
}

// SetCustomerLicenseAgreementNumber gets a reference to the given map[string]interface{} and assigns it to the CustomerLicenseAgreementNumber field.
func (o *MicrosoftPartnerReferralIot) SetCustomerLicenseAgreementNumber(v map[string]interface{}) {
	o.CustomerLicenseAgreementNumber = v
}

// GetDeviceCategory returns the DeviceCategory field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralIot) GetDeviceCategory() map[string]interface{} {
	if o == nil || IsNil(o.DeviceCategory) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeviceCategory
}

// GetDeviceCategoryOk returns a tuple with the DeviceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralIot) GetDeviceCategoryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeviceCategory) {
		return map[string]interface{}{}, false
	}
	return o.DeviceCategory, true
}

// HasDeviceCategory returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralIot) HasDeviceCategory() bool {
	if o != nil && !IsNil(o.DeviceCategory) {
		return true
	}

	return false
}

// SetDeviceCategory gets a reference to the given map[string]interface{} and assigns it to the DeviceCategory field.
func (o *MicrosoftPartnerReferralIot) SetDeviceCategory(v map[string]interface{}) {
	o.DeviceCategory = v
}

// GetSiliconType returns the SiliconType field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralIot) GetSiliconType() map[string]interface{} {
	if o == nil || IsNil(o.SiliconType) {
		var ret map[string]interface{}
		return ret
	}
	return o.SiliconType
}

// GetSiliconTypeOk returns a tuple with the SiliconType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralIot) GetSiliconTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SiliconType) {
		return map[string]interface{}{}, false
	}
	return o.SiliconType, true
}

// HasSiliconType returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralIot) HasSiliconType() bool {
	if o != nil && !IsNil(o.SiliconType) {
		return true
	}

	return false
}

// SetSiliconType gets a reference to the given map[string]interface{} and assigns it to the SiliconType field.
func (o *MicrosoftPartnerReferralIot) SetSiliconType(v map[string]interface{}) {
	o.SiliconType = v
}

func (o MicrosoftPartnerReferralIot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrosoftPartnerReferralIot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachServices) {
		toSerialize["attachServices"] = o.AttachServices
	}
	if !IsNil(o.AzureCertifiedDevice) {
		toSerialize["azureCertifiedDevice"] = o.AzureCertifiedDevice
	}
	if !IsNil(o.CustomerLicenseAgreementNumber) {
		toSerialize["customerLicenseAgreementNumber"] = o.CustomerLicenseAgreementNumber
	}
	if !IsNil(o.DeviceCategory) {
		toSerialize["deviceCategory"] = o.DeviceCategory
	}
	if !IsNil(o.SiliconType) {
		toSerialize["siliconType"] = o.SiliconType
	}
	return toSerialize, nil
}

type NullableMicrosoftPartnerReferralIot struct {
	value *MicrosoftPartnerReferralIot
	isSet bool
}

func (v NullableMicrosoftPartnerReferralIot) Get() *MicrosoftPartnerReferralIot {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralIot) Set(val *MicrosoftPartnerReferralIot) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralIot) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralIot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralIot(val *MicrosoftPartnerReferralIot) *NullableMicrosoftPartnerReferralIot {
	return &NullableMicrosoftPartnerReferralIot{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralIot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralIot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
