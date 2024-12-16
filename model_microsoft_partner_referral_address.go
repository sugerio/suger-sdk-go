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

// checks if the MicrosoftPartnerReferralAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftPartnerReferralAddress{}

// MicrosoftPartnerReferralAddress struct for MicrosoftPartnerReferralAddress
type MicrosoftPartnerReferralAddress struct {
	AddressLine1 *string                `json:"addressLine1,omitempty"`
	AddressLine2 *string                `json:"addressLine2,omitempty"`
	City         *string                `json:"city,omitempty"`
	Country      *string                `json:"country,omitempty"`
	PostalCode   map[string]interface{} `json:"postalCode,omitempty"`
	Region       *string                `json:"region,omitempty"`
	State        *string                `json:"state,omitempty"`
}

// NewMicrosoftPartnerReferralAddress instantiates a new MicrosoftPartnerReferralAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftPartnerReferralAddress() *MicrosoftPartnerReferralAddress {
	this := MicrosoftPartnerReferralAddress{}
	return &this
}

// NewMicrosoftPartnerReferralAddressWithDefaults instantiates a new MicrosoftPartnerReferralAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftPartnerReferralAddressWithDefaults() *MicrosoftPartnerReferralAddress {
	this := MicrosoftPartnerReferralAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *MicrosoftPartnerReferralAddress) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *MicrosoftPartnerReferralAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MicrosoftPartnerReferralAddress) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *MicrosoftPartnerReferralAddress) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetPostalCode() map[string]interface{} {
	if o == nil || IsNil(o.PostalCode) {
		var ret map[string]interface{}
		return ret
	}
	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetPostalCodeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return map[string]interface{}{}, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given map[string]interface{} and assigns it to the PostalCode field.
func (o *MicrosoftPartnerReferralAddress) SetPostalCode(v map[string]interface{}) {
	o.PostalCode = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MicrosoftPartnerReferralAddress) SetRegion(v string) {
	o.Region = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MicrosoftPartnerReferralAddress) SetState(v string) {
	o.State = &v
}

func (o MicrosoftPartnerReferralAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrosoftPartnerReferralAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableMicrosoftPartnerReferralAddress struct {
	value *MicrosoftPartnerReferralAddress
	isSet bool
}

func (v NullableMicrosoftPartnerReferralAddress) Get() *MicrosoftPartnerReferralAddress {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralAddress) Set(val *MicrosoftPartnerReferralAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralAddress(val *MicrosoftPartnerReferralAddress) *NullableMicrosoftPartnerReferralAddress {
	return &NullableMicrosoftPartnerReferralAddress{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
