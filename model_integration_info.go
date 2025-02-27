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

// checks if the IntegrationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationInfo{}

// IntegrationInfo struct for IntegrationInfo
type IntegrationInfo struct {
	AlibabaIntegration       *AlibabaMarketplaceIntegration `json:"alibabaIntegration,omitempty"`
	AwsAceIntegration        *AwsAceIntegration             `json:"awsAceIntegration,omitempty"`
	AwsIntegration           *AwsMarketplaceIntegration     `json:"awsIntegration,omitempty"`
	AzureIntegration         *AzureIntegration              `json:"azureIntegration,omitempty"`
	GcpIntegration           *GcpIntegration                `json:"gcpIntegration,omitempty"`
	HubspotCrmIntegration    *HubspotCrmIntegration         `json:"hubspotCrmIntegration,omitempty"`
	MetronomeIntegration     *MetronomeIntegration          `json:"metronomeIntegration,omitempty"`
	OrbIntegration           *OrbIntegration                `json:"orbIntegration,omitempty"`
	SalesforceCrmIntegration *SalesforceCrmIntegration      `json:"salesforceCrmIntegration,omitempty"`
	SlackIntegration         *SlackIntegration              `json:"slackIntegration,omitempty"`
}

// NewIntegrationInfo instantiates a new IntegrationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationInfo() *IntegrationInfo {
	this := IntegrationInfo{}
	return &this
}

// NewIntegrationInfoWithDefaults instantiates a new IntegrationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationInfoWithDefaults() *IntegrationInfo {
	this := IntegrationInfo{}
	return &this
}

// GetAlibabaIntegration returns the AlibabaIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetAlibabaIntegration() AlibabaMarketplaceIntegration {
	if o == nil || IsNil(o.AlibabaIntegration) {
		var ret AlibabaMarketplaceIntegration
		return ret
	}
	return *o.AlibabaIntegration
}

// GetAlibabaIntegrationOk returns a tuple with the AlibabaIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetAlibabaIntegrationOk() (*AlibabaMarketplaceIntegration, bool) {
	if o == nil || IsNil(o.AlibabaIntegration) {
		return nil, false
	}
	return o.AlibabaIntegration, true
}

// HasAlibabaIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasAlibabaIntegration() bool {
	if o != nil && !IsNil(o.AlibabaIntegration) {
		return true
	}

	return false
}

// SetAlibabaIntegration gets a reference to the given AlibabaMarketplaceIntegration and assigns it to the AlibabaIntegration field.
func (o *IntegrationInfo) SetAlibabaIntegration(v AlibabaMarketplaceIntegration) {
	o.AlibabaIntegration = &v
}

// GetAwsAceIntegration returns the AwsAceIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetAwsAceIntegration() AwsAceIntegration {
	if o == nil || IsNil(o.AwsAceIntegration) {
		var ret AwsAceIntegration
		return ret
	}
	return *o.AwsAceIntegration
}

// GetAwsAceIntegrationOk returns a tuple with the AwsAceIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetAwsAceIntegrationOk() (*AwsAceIntegration, bool) {
	if o == nil || IsNil(o.AwsAceIntegration) {
		return nil, false
	}
	return o.AwsAceIntegration, true
}

// HasAwsAceIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasAwsAceIntegration() bool {
	if o != nil && !IsNil(o.AwsAceIntegration) {
		return true
	}

	return false
}

// SetAwsAceIntegration gets a reference to the given AwsAceIntegration and assigns it to the AwsAceIntegration field.
func (o *IntegrationInfo) SetAwsAceIntegration(v AwsAceIntegration) {
	o.AwsAceIntegration = &v
}

// GetAwsIntegration returns the AwsIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetAwsIntegration() AwsMarketplaceIntegration {
	if o == nil || IsNil(o.AwsIntegration) {
		var ret AwsMarketplaceIntegration
		return ret
	}
	return *o.AwsIntegration
}

// GetAwsIntegrationOk returns a tuple with the AwsIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetAwsIntegrationOk() (*AwsMarketplaceIntegration, bool) {
	if o == nil || IsNil(o.AwsIntegration) {
		return nil, false
	}
	return o.AwsIntegration, true
}

// HasAwsIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasAwsIntegration() bool {
	if o != nil && !IsNil(o.AwsIntegration) {
		return true
	}

	return false
}

// SetAwsIntegration gets a reference to the given AwsMarketplaceIntegration and assigns it to the AwsIntegration field.
func (o *IntegrationInfo) SetAwsIntegration(v AwsMarketplaceIntegration) {
	o.AwsIntegration = &v
}

// GetAzureIntegration returns the AzureIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetAzureIntegration() AzureIntegration {
	if o == nil || IsNil(o.AzureIntegration) {
		var ret AzureIntegration
		return ret
	}
	return *o.AzureIntegration
}

// GetAzureIntegrationOk returns a tuple with the AzureIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetAzureIntegrationOk() (*AzureIntegration, bool) {
	if o == nil || IsNil(o.AzureIntegration) {
		return nil, false
	}
	return o.AzureIntegration, true
}

// HasAzureIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasAzureIntegration() bool {
	if o != nil && !IsNil(o.AzureIntegration) {
		return true
	}

	return false
}

// SetAzureIntegration gets a reference to the given AzureIntegration and assigns it to the AzureIntegration field.
func (o *IntegrationInfo) SetAzureIntegration(v AzureIntegration) {
	o.AzureIntegration = &v
}

// GetGcpIntegration returns the GcpIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetGcpIntegration() GcpIntegration {
	if o == nil || IsNil(o.GcpIntegration) {
		var ret GcpIntegration
		return ret
	}
	return *o.GcpIntegration
}

// GetGcpIntegrationOk returns a tuple with the GcpIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetGcpIntegrationOk() (*GcpIntegration, bool) {
	if o == nil || IsNil(o.GcpIntegration) {
		return nil, false
	}
	return o.GcpIntegration, true
}

// HasGcpIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasGcpIntegration() bool {
	if o != nil && !IsNil(o.GcpIntegration) {
		return true
	}

	return false
}

// SetGcpIntegration gets a reference to the given GcpIntegration and assigns it to the GcpIntegration field.
func (o *IntegrationInfo) SetGcpIntegration(v GcpIntegration) {
	o.GcpIntegration = &v
}

// GetHubspotCrmIntegration returns the HubspotCrmIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetHubspotCrmIntegration() HubspotCrmIntegration {
	if o == nil || IsNil(o.HubspotCrmIntegration) {
		var ret HubspotCrmIntegration
		return ret
	}
	return *o.HubspotCrmIntegration
}

// GetHubspotCrmIntegrationOk returns a tuple with the HubspotCrmIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetHubspotCrmIntegrationOk() (*HubspotCrmIntegration, bool) {
	if o == nil || IsNil(o.HubspotCrmIntegration) {
		return nil, false
	}
	return o.HubspotCrmIntegration, true
}

// HasHubspotCrmIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasHubspotCrmIntegration() bool {
	if o != nil && !IsNil(o.HubspotCrmIntegration) {
		return true
	}

	return false
}

// SetHubspotCrmIntegration gets a reference to the given HubspotCrmIntegration and assigns it to the HubspotCrmIntegration field.
func (o *IntegrationInfo) SetHubspotCrmIntegration(v HubspotCrmIntegration) {
	o.HubspotCrmIntegration = &v
}

// GetMetronomeIntegration returns the MetronomeIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetMetronomeIntegration() MetronomeIntegration {
	if o == nil || IsNil(o.MetronomeIntegration) {
		var ret MetronomeIntegration
		return ret
	}
	return *o.MetronomeIntegration
}

// GetMetronomeIntegrationOk returns a tuple with the MetronomeIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetMetronomeIntegrationOk() (*MetronomeIntegration, bool) {
	if o == nil || IsNil(o.MetronomeIntegration) {
		return nil, false
	}
	return o.MetronomeIntegration, true
}

// HasMetronomeIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasMetronomeIntegration() bool {
	if o != nil && !IsNil(o.MetronomeIntegration) {
		return true
	}

	return false
}

// SetMetronomeIntegration gets a reference to the given MetronomeIntegration and assigns it to the MetronomeIntegration field.
func (o *IntegrationInfo) SetMetronomeIntegration(v MetronomeIntegration) {
	o.MetronomeIntegration = &v
}

// GetOrbIntegration returns the OrbIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetOrbIntegration() OrbIntegration {
	if o == nil || IsNil(o.OrbIntegration) {
		var ret OrbIntegration
		return ret
	}
	return *o.OrbIntegration
}

// GetOrbIntegrationOk returns a tuple with the OrbIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetOrbIntegrationOk() (*OrbIntegration, bool) {
	if o == nil || IsNil(o.OrbIntegration) {
		return nil, false
	}
	return o.OrbIntegration, true
}

// HasOrbIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasOrbIntegration() bool {
	if o != nil && !IsNil(o.OrbIntegration) {
		return true
	}

	return false
}

// SetOrbIntegration gets a reference to the given OrbIntegration and assigns it to the OrbIntegration field.
func (o *IntegrationInfo) SetOrbIntegration(v OrbIntegration) {
	o.OrbIntegration = &v
}

// GetSalesforceCrmIntegration returns the SalesforceCrmIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetSalesforceCrmIntegration() SalesforceCrmIntegration {
	if o == nil || IsNil(o.SalesforceCrmIntegration) {
		var ret SalesforceCrmIntegration
		return ret
	}
	return *o.SalesforceCrmIntegration
}

// GetSalesforceCrmIntegrationOk returns a tuple with the SalesforceCrmIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetSalesforceCrmIntegrationOk() (*SalesforceCrmIntegration, bool) {
	if o == nil || IsNil(o.SalesforceCrmIntegration) {
		return nil, false
	}
	return o.SalesforceCrmIntegration, true
}

// HasSalesforceCrmIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasSalesforceCrmIntegration() bool {
	if o != nil && !IsNil(o.SalesforceCrmIntegration) {
		return true
	}

	return false
}

// SetSalesforceCrmIntegration gets a reference to the given SalesforceCrmIntegration and assigns it to the SalesforceCrmIntegration field.
func (o *IntegrationInfo) SetSalesforceCrmIntegration(v SalesforceCrmIntegration) {
	o.SalesforceCrmIntegration = &v
}

// GetSlackIntegration returns the SlackIntegration field value if set, zero value otherwise.
func (o *IntegrationInfo) GetSlackIntegration() SlackIntegration {
	if o == nil || IsNil(o.SlackIntegration) {
		var ret SlackIntegration
		return ret
	}
	return *o.SlackIntegration
}

// GetSlackIntegrationOk returns a tuple with the SlackIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationInfo) GetSlackIntegrationOk() (*SlackIntegration, bool) {
	if o == nil || IsNil(o.SlackIntegration) {
		return nil, false
	}
	return o.SlackIntegration, true
}

// HasSlackIntegration returns a boolean if a field has been set.
func (o *IntegrationInfo) HasSlackIntegration() bool {
	if o != nil && !IsNil(o.SlackIntegration) {
		return true
	}

	return false
}

// SetSlackIntegration gets a reference to the given SlackIntegration and assigns it to the SlackIntegration field.
func (o *IntegrationInfo) SetSlackIntegration(v SlackIntegration) {
	o.SlackIntegration = &v
}

func (o IntegrationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlibabaIntegration) {
		toSerialize["alibabaIntegration"] = o.AlibabaIntegration
	}
	if !IsNil(o.AwsAceIntegration) {
		toSerialize["awsAceIntegration"] = o.AwsAceIntegration
	}
	if !IsNil(o.AwsIntegration) {
		toSerialize["awsIntegration"] = o.AwsIntegration
	}
	if !IsNil(o.AzureIntegration) {
		toSerialize["azureIntegration"] = o.AzureIntegration
	}
	if !IsNil(o.GcpIntegration) {
		toSerialize["gcpIntegration"] = o.GcpIntegration
	}
	if !IsNil(o.HubspotCrmIntegration) {
		toSerialize["hubspotCrmIntegration"] = o.HubspotCrmIntegration
	}
	if !IsNil(o.MetronomeIntegration) {
		toSerialize["metronomeIntegration"] = o.MetronomeIntegration
	}
	if !IsNil(o.OrbIntegration) {
		toSerialize["orbIntegration"] = o.OrbIntegration
	}
	if !IsNil(o.SalesforceCrmIntegration) {
		toSerialize["salesforceCrmIntegration"] = o.SalesforceCrmIntegration
	}
	if !IsNil(o.SlackIntegration) {
		toSerialize["slackIntegration"] = o.SlackIntegration
	}
	return toSerialize, nil
}

type NullableIntegrationInfo struct {
	value *IntegrationInfo
	isSet bool
}

func (v NullableIntegrationInfo) Get() *IntegrationInfo {
	return v.value
}

func (v *NullableIntegrationInfo) Set(val *IntegrationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationInfo(val *IntegrationInfo) *NullableIntegrationInfo {
	return &NullableIntegrationInfo{value: val, isSet: true}
}

func (v NullableIntegrationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
