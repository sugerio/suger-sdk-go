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

// checks if the CreateIntegrationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIntegrationParams{}

// CreateIntegrationParams struct for CreateIntegrationParams
type CreateIntegrationParams struct {
	CreatedBy      *string         `json:"createdBy,omitempty"`
	Info           IntegrationInfo `json:"info"`
	OrganizationID string          `json:"organizationID"`
	Partner        Partner         `json:"partner"`
	Service        PartnerService  `json:"service"`
}

// NewCreateIntegrationParams instantiates a new CreateIntegrationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIntegrationParams(info IntegrationInfo, organizationID string, partner Partner, service PartnerService) *CreateIntegrationParams {
	this := CreateIntegrationParams{}
	this.Info = info
	this.OrganizationID = organizationID
	this.Partner = partner
	this.Service = service
	return &this
}

// NewCreateIntegrationParamsWithDefaults instantiates a new CreateIntegrationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIntegrationParamsWithDefaults() *CreateIntegrationParams {
	this := CreateIntegrationParams{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CreateIntegrationParams) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIntegrationParams) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CreateIntegrationParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CreateIntegrationParams) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetInfo returns the Info field value
func (o *CreateIntegrationParams) GetInfo() IntegrationInfo {
	if o == nil {
		var ret IntegrationInfo
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationParams) GetInfoOk() (*IntegrationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *CreateIntegrationParams) SetInfo(v IntegrationInfo) {
	o.Info = v
}

// GetOrganizationID returns the OrganizationID field value
func (o *CreateIntegrationParams) GetOrganizationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationParams) GetOrganizationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationID, true
}

// SetOrganizationID sets field value
func (o *CreateIntegrationParams) SetOrganizationID(v string) {
	o.OrganizationID = v
}

// GetPartner returns the Partner field value
func (o *CreateIntegrationParams) GetPartner() Partner {
	if o == nil {
		var ret Partner
		return ret
	}

	return o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationParams) GetPartnerOk() (*Partner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partner, true
}

// SetPartner sets field value
func (o *CreateIntegrationParams) SetPartner(v Partner) {
	o.Partner = v
}

// GetService returns the Service field value
func (o *CreateIntegrationParams) GetService() PartnerService {
	if o == nil {
		var ret PartnerService
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *CreateIntegrationParams) GetServiceOk() (*PartnerService, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *CreateIntegrationParams) SetService(v PartnerService) {
	o.Service = v
}

func (o CreateIntegrationParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIntegrationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	toSerialize["info"] = o.Info
	toSerialize["organizationID"] = o.OrganizationID
	toSerialize["partner"] = o.Partner
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

type NullableCreateIntegrationParams struct {
	value *CreateIntegrationParams
	isSet bool
}

func (v NullableCreateIntegrationParams) Get() *CreateIntegrationParams {
	return v.value
}

func (v *NullableCreateIntegrationParams) Set(val *CreateIntegrationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIntegrationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIntegrationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIntegrationParams(val *CreateIntegrationParams) *NullableCreateIntegrationParams {
	return &NullableCreateIntegrationParams{value: val, isSet: true}
}

func (v NullableCreateIntegrationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIntegrationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
