/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RevenueRecordInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevenueRecordInfo{}

// RevenueRecordInfo struct for RevenueRecordInfo
type RevenueRecordInfo struct {
	// For raw revenue records in AWS Marketplace
	AwsRevenueRecords []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent `json:"awsRevenueRecords,omitempty"`
	// For raw revenue records in Azure Marketplace
	AzureRevenueRecords []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue `json:"azureRevenueRecords,omitempty"`
	// Whether the disbursement notification has been sent to the seller/ISV.
	DisbursementNotificationSent *bool `json:"disbursementNotificationSent,omitempty"`
	// For raw revenue records in GCP Marketplace
	GcpRevenueRecords []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage `json:"gcpRevenueRecords,omitempty"`
	// Resource name for the revenue record. Applicable only to GCP Marketplace.
	Resource *string `json:"resource,omitempty"`
}

// NewRevenueRecordInfo instantiates a new RevenueRecordInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevenueRecordInfo() *RevenueRecordInfo {
	this := RevenueRecordInfo{}
	return &this
}

// NewRevenueRecordInfoWithDefaults instantiates a new RevenueRecordInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevenueRecordInfoWithDefaults() *RevenueRecordInfo {
	this := RevenueRecordInfo{}
	return &this
}

// GetAwsRevenueRecords returns the AwsRevenueRecords field value if set, zero value otherwise.
func (o *RevenueRecordInfo) GetAwsRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent {
	if o == nil || IsNil(o.AwsRevenueRecords) {
		var ret []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent
		return ret
	}
	return o.AwsRevenueRecords
}

// GetAwsRevenueRecordsOk returns a tuple with the AwsRevenueRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecordInfo) GetAwsRevenueRecordsOk() ([]GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent, bool) {
	if o == nil || IsNil(o.AwsRevenueRecords) {
		return nil, false
	}
	return o.AwsRevenueRecords, true
}

// HasAwsRevenueRecords returns a boolean if a field has been set.
func (o *RevenueRecordInfo) HasAwsRevenueRecords() bool {
	if o != nil && !IsNil(o.AwsRevenueRecords) {
		return true
	}

	return false
}

// SetAwsRevenueRecords gets a reference to the given []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent and assigns it to the AwsRevenueRecords field.
func (o *RevenueRecordInfo) SetAwsRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) {
	o.AwsRevenueRecords = v
}

// GetAzureRevenueRecords returns the AzureRevenueRecords field value if set, zero value otherwise.
func (o *RevenueRecordInfo) GetAzureRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue {
	if o == nil || IsNil(o.AzureRevenueRecords) {
		var ret []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue
		return ret
	}
	return o.AzureRevenueRecords
}

// GetAzureRevenueRecordsOk returns a tuple with the AzureRevenueRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecordInfo) GetAzureRevenueRecordsOk() ([]GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue, bool) {
	if o == nil || IsNil(o.AzureRevenueRecords) {
		return nil, false
	}
	return o.AzureRevenueRecords, true
}

// HasAzureRevenueRecords returns a boolean if a field has been set.
func (o *RevenueRecordInfo) HasAzureRevenueRecords() bool {
	if o != nil && !IsNil(o.AzureRevenueRecords) {
		return true
	}

	return false
}

// SetAzureRevenueRecords gets a reference to the given []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue and assigns it to the AzureRevenueRecords field.
func (o *RevenueRecordInfo) SetAzureRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue) {
	o.AzureRevenueRecords = v
}

// GetDisbursementNotificationSent returns the DisbursementNotificationSent field value if set, zero value otherwise.
func (o *RevenueRecordInfo) GetDisbursementNotificationSent() bool {
	if o == nil || IsNil(o.DisbursementNotificationSent) {
		var ret bool
		return ret
	}
	return *o.DisbursementNotificationSent
}

// GetDisbursementNotificationSentOk returns a tuple with the DisbursementNotificationSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecordInfo) GetDisbursementNotificationSentOk() (*bool, bool) {
	if o == nil || IsNil(o.DisbursementNotificationSent) {
		return nil, false
	}
	return o.DisbursementNotificationSent, true
}

// HasDisbursementNotificationSent returns a boolean if a field has been set.
func (o *RevenueRecordInfo) HasDisbursementNotificationSent() bool {
	if o != nil && !IsNil(o.DisbursementNotificationSent) {
		return true
	}

	return false
}

// SetDisbursementNotificationSent gets a reference to the given bool and assigns it to the DisbursementNotificationSent field.
func (o *RevenueRecordInfo) SetDisbursementNotificationSent(v bool) {
	o.DisbursementNotificationSent = &v
}

// GetGcpRevenueRecords returns the GcpRevenueRecords field value if set, zero value otherwise.
func (o *RevenueRecordInfo) GetGcpRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage {
	if o == nil || IsNil(o.GcpRevenueRecords) {
		var ret []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage
		return ret
	}
	return o.GcpRevenueRecords
}

// GetGcpRevenueRecordsOk returns a tuple with the GcpRevenueRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecordInfo) GetGcpRevenueRecordsOk() ([]GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage, bool) {
	if o == nil || IsNil(o.GcpRevenueRecords) {
		return nil, false
	}
	return o.GcpRevenueRecords, true
}

// HasGcpRevenueRecords returns a boolean if a field has been set.
func (o *RevenueRecordInfo) HasGcpRevenueRecords() bool {
	if o != nil && !IsNil(o.GcpRevenueRecords) {
		return true
	}

	return false
}

// SetGcpRevenueRecords gets a reference to the given []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage and assigns it to the GcpRevenueRecords field.
func (o *RevenueRecordInfo) SetGcpRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage) {
	o.GcpRevenueRecords = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *RevenueRecordInfo) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecordInfo) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *RevenueRecordInfo) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *RevenueRecordInfo) SetResource(v string) {
	o.Resource = &v
}

func (o RevenueRecordInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevenueRecordInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsRevenueRecords) {
		toSerialize["awsRevenueRecords"] = o.AwsRevenueRecords
	}
	if !IsNil(o.AzureRevenueRecords) {
		toSerialize["azureRevenueRecords"] = o.AzureRevenueRecords
	}
	if !IsNil(o.DisbursementNotificationSent) {
		toSerialize["disbursementNotificationSent"] = o.DisbursementNotificationSent
	}
	if !IsNil(o.GcpRevenueRecords) {
		toSerialize["gcpRevenueRecords"] = o.GcpRevenueRecords
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableRevenueRecordInfo struct {
	value *RevenueRecordInfo
	isSet bool
}

func (v NullableRevenueRecordInfo) Get() *RevenueRecordInfo {
	return v.value
}

func (v *NullableRevenueRecordInfo) Set(val *RevenueRecordInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueRecordInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueRecordInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueRecordInfo(val *RevenueRecordInfo) *NullableRevenueRecordInfo {
	return &NullableRevenueRecordInfo{value: val, isSet: true}
}

func (v NullableRevenueRecordInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueRecordInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


