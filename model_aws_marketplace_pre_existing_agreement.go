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

// checks if the AwsMarketplacePreExistingAgreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplacePreExistingAgreement{}

// AwsMarketplacePreExistingAgreement struct for AwsMarketplacePreExistingAgreement
type AwsMarketplacePreExistingAgreement struct {
	// Indicates if the existing agreement was signed outside AWS Marketplace or within AWS Marketplace. one of values [\"External\", \"AwsMarketplace\"]
	AcquisitionChannel *AwsRenewalOfferType `json:"AcquisitionChannel,omitempty"`
	// Indicates which pricing model the existing agreement uses.
	PricingModel *AwsMarketplaceCatalogPricingModel `json:"PricingModel,omitempty"`
}

// NewAwsMarketplacePreExistingAgreement instantiates a new AwsMarketplacePreExistingAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplacePreExistingAgreement() *AwsMarketplacePreExistingAgreement {
	this := AwsMarketplacePreExistingAgreement{}
	return &this
}

// NewAwsMarketplacePreExistingAgreementWithDefaults instantiates a new AwsMarketplacePreExistingAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplacePreExistingAgreementWithDefaults() *AwsMarketplacePreExistingAgreement {
	this := AwsMarketplacePreExistingAgreement{}
	return &this
}

// GetAcquisitionChannel returns the AcquisitionChannel field value if set, zero value otherwise.
func (o *AwsMarketplacePreExistingAgreement) GetAcquisitionChannel() AwsRenewalOfferType {
	if o == nil || IsNil(o.AcquisitionChannel) {
		var ret AwsRenewalOfferType
		return ret
	}
	return *o.AcquisitionChannel
}

// GetAcquisitionChannelOk returns a tuple with the AcquisitionChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplacePreExistingAgreement) GetAcquisitionChannelOk() (*AwsRenewalOfferType, bool) {
	if o == nil || IsNil(o.AcquisitionChannel) {
		return nil, false
	}
	return o.AcquisitionChannel, true
}

// HasAcquisitionChannel returns a boolean if a field has been set.
func (o *AwsMarketplacePreExistingAgreement) HasAcquisitionChannel() bool {
	if o != nil && !IsNil(o.AcquisitionChannel) {
		return true
	}

	return false
}

// SetAcquisitionChannel gets a reference to the given AwsRenewalOfferType and assigns it to the AcquisitionChannel field.
func (o *AwsMarketplacePreExistingAgreement) SetAcquisitionChannel(v AwsRenewalOfferType) {
	o.AcquisitionChannel = &v
}

// GetPricingModel returns the PricingModel field value if set, zero value otherwise.
func (o *AwsMarketplacePreExistingAgreement) GetPricingModel() AwsMarketplaceCatalogPricingModel {
	if o == nil || IsNil(o.PricingModel) {
		var ret AwsMarketplaceCatalogPricingModel
		return ret
	}
	return *o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplacePreExistingAgreement) GetPricingModelOk() (*AwsMarketplaceCatalogPricingModel, bool) {
	if o == nil || IsNil(o.PricingModel) {
		return nil, false
	}
	return o.PricingModel, true
}

// HasPricingModel returns a boolean if a field has been set.
func (o *AwsMarketplacePreExistingAgreement) HasPricingModel() bool {
	if o != nil && !IsNil(o.PricingModel) {
		return true
	}

	return false
}

// SetPricingModel gets a reference to the given AwsMarketplaceCatalogPricingModel and assigns it to the PricingModel field.
func (o *AwsMarketplacePreExistingAgreement) SetPricingModel(v AwsMarketplaceCatalogPricingModel) {
	o.PricingModel = &v
}

func (o AwsMarketplacePreExistingAgreement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplacePreExistingAgreement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcquisitionChannel) {
		toSerialize["AcquisitionChannel"] = o.AcquisitionChannel
	}
	if !IsNil(o.PricingModel) {
		toSerialize["PricingModel"] = o.PricingModel
	}
	return toSerialize, nil
}

type NullableAwsMarketplacePreExistingAgreement struct {
	value *AwsMarketplacePreExistingAgreement
	isSet bool
}

func (v NullableAwsMarketplacePreExistingAgreement) Get() *AwsMarketplacePreExistingAgreement {
	return v.value
}

func (v *NullableAwsMarketplacePreExistingAgreement) Set(val *AwsMarketplacePreExistingAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplacePreExistingAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplacePreExistingAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplacePreExistingAgreement(val *AwsMarketplacePreExistingAgreement) *NullableAwsMarketplacePreExistingAgreement {
	return &NullableAwsMarketplacePreExistingAgreement{value: val, isSet: true}
}

func (v NullableAwsMarketplacePreExistingAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplacePreExistingAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
