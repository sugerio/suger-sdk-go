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

// checks if the UpdateIntegrationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIntegrationParams{}

// UpdateIntegrationParams struct for UpdateIntegrationParams
type UpdateIntegrationParams struct {
	Info           IntegrationInfo `json:"info"`
	OrganizationID string          `json:"organizationID"`
	Partner        Partner         `json:"partner"`
	Service        PartnerService  `json:"service"`
}

// NewUpdateIntegrationParams instantiates a new UpdateIntegrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIntegrationParams(info IntegrationInfo, organizationID string, partner Partner, service PartnerService) *UpdateIntegrationParams {
	this := UpdateIntegrationParams{}
	this.Info = info
	this.OrganizationID = organizationID
	this.Partner = partner
	this.Service = service
	return &this
}

// NewUpdateIntegrationParamsWithDefaults instantiates a new UpdateIntegrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIntegrationParamsWithDefaults() *UpdateIntegrationParams {
	this := UpdateIntegrationParams{}
	return &this
}

// GetInfo returns the Info field value
func (o *UpdateIntegrationParams) GetInfo() IntegrationInfo {
	if o == nil {
		var ret IntegrationInfo
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationParams) GetInfoOk() (*IntegrationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *UpdateIntegrationParams) SetInfo(v IntegrationInfo) {
	o.Info = v
}

// GetOrganizationID returns the OrganizationID field value
func (o *UpdateIntegrationParams) GetOrganizationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationParams) GetOrganizationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationID, true
}

// SetOrganizationID sets field value
func (o *UpdateIntegrationParams) SetOrganizationID(v string) {
	o.OrganizationID = v
}

// GetPartner returns the Partner field value
func (o *UpdateIntegrationParams) GetPartner() Partner {
	if o == nil {
		var ret Partner
		return ret
	}

	return o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationParams) GetPartnerOk() (*Partner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partner, true
}

// SetPartner sets field value
func (o *UpdateIntegrationParams) SetPartner(v Partner) {
	o.Partner = v
}

// GetService returns the Service field value
func (o *UpdateIntegrationParams) GetService() PartnerService {
	if o == nil {
		var ret PartnerService
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *UpdateIntegrationParams) GetServiceOk() (*PartnerService, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *UpdateIntegrationParams) SetService(v PartnerService) {
	o.Service = v
}

func (o UpdateIntegrationParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIntegrationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["info"] = o.Info
	toSerialize["organizationID"] = o.OrganizationID
	toSerialize["partner"] = o.Partner
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

type NullableUpdateIntegrationParams struct {
	value *UpdateIntegrationParams
	isSet bool
}

func (v NullableUpdateIntegrationParams) Get() *UpdateIntegrationParams {
	return v.value
}

func (v *NullableUpdateIntegrationParams) Set(val *UpdateIntegrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIntegrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIntegrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIntegrationParams(val *UpdateIntegrationParams) *NullableUpdateIntegrationParams {
	return &NullableUpdateIntegrationParams{value: val, isSet: true}
}

func (v NullableUpdateIntegrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIntegrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
