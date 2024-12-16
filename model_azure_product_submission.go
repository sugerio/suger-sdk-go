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
	"time"
)

// checks if the AzureProductSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductSubmission{}

// AzureProductSubmission struct for AzureProductSubmission
type AzureProductSubmission struct {
	AreResourcesReady  *bool                   `json:"areResourcesReady,omitempty"`
	FriendlyName       *string                 `json:"friendlyName,omitempty"`
	Id                 *string                 `json:"id,omitempty"`
	PendingUpdateInfo  *AzurePendingUpdateInfo `json:"pendingUpdateInfo,omitempty"`
	PublishedTimeInUtc *time.Time              `json:"publishedTimeInUtc,omitempty"`
	ReleaseNumber      *int32                  `json:"releaseNumber,omitempty"`
	ResourceType       *string                 `json:"resourceType,omitempty"`
	Resources          []AzureTypeValue        `json:"resources,omitempty"`
	State              *string                 `json:"state,omitempty"`
	SubState           *string                 `json:"subState,omitempty"`
	Targets            []AzureTypeValue        `json:"targets,omitempty"`
	VariantResources   []AzureVariantResource  `json:"variantResources,omitempty"`
}

// NewAzureProductSubmission instantiates a new AzureProductSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductSubmission() *AzureProductSubmission {
	this := AzureProductSubmission{}
	return &this
}

// NewAzureProductSubmissionWithDefaults instantiates a new AzureProductSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductSubmissionWithDefaults() *AzureProductSubmission {
	this := AzureProductSubmission{}
	return &this
}

// GetAreResourcesReady returns the AreResourcesReady field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetAreResourcesReady() bool {
	if o == nil || IsNil(o.AreResourcesReady) {
		var ret bool
		return ret
	}
	return *o.AreResourcesReady
}

// GetAreResourcesReadyOk returns a tuple with the AreResourcesReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetAreResourcesReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.AreResourcesReady) {
		return nil, false
	}
	return o.AreResourcesReady, true
}

// HasAreResourcesReady returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasAreResourcesReady() bool {
	if o != nil && !IsNil(o.AreResourcesReady) {
		return true
	}

	return false
}

// SetAreResourcesReady gets a reference to the given bool and assigns it to the AreResourcesReady field.
func (o *AzureProductSubmission) SetAreResourcesReady(v bool) {
	o.AreResourcesReady = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *AzureProductSubmission) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureProductSubmission) SetId(v string) {
	o.Id = &v
}

// GetPendingUpdateInfo returns the PendingUpdateInfo field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetPendingUpdateInfo() AzurePendingUpdateInfo {
	if o == nil || IsNil(o.PendingUpdateInfo) {
		var ret AzurePendingUpdateInfo
		return ret
	}
	return *o.PendingUpdateInfo
}

// GetPendingUpdateInfoOk returns a tuple with the PendingUpdateInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetPendingUpdateInfoOk() (*AzurePendingUpdateInfo, bool) {
	if o == nil || IsNil(o.PendingUpdateInfo) {
		return nil, false
	}
	return o.PendingUpdateInfo, true
}

// HasPendingUpdateInfo returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasPendingUpdateInfo() bool {
	if o != nil && !IsNil(o.PendingUpdateInfo) {
		return true
	}

	return false
}

// SetPendingUpdateInfo gets a reference to the given AzurePendingUpdateInfo and assigns it to the PendingUpdateInfo field.
func (o *AzureProductSubmission) SetPendingUpdateInfo(v AzurePendingUpdateInfo) {
	o.PendingUpdateInfo = &v
}

// GetPublishedTimeInUtc returns the PublishedTimeInUtc field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetPublishedTimeInUtc() time.Time {
	if o == nil || IsNil(o.PublishedTimeInUtc) {
		var ret time.Time
		return ret
	}
	return *o.PublishedTimeInUtc
}

// GetPublishedTimeInUtcOk returns a tuple with the PublishedTimeInUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetPublishedTimeInUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedTimeInUtc) {
		return nil, false
	}
	return o.PublishedTimeInUtc, true
}

// HasPublishedTimeInUtc returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasPublishedTimeInUtc() bool {
	if o != nil && !IsNil(o.PublishedTimeInUtc) {
		return true
	}

	return false
}

// SetPublishedTimeInUtc gets a reference to the given time.Time and assigns it to the PublishedTimeInUtc field.
func (o *AzureProductSubmission) SetPublishedTimeInUtc(v time.Time) {
	o.PublishedTimeInUtc = &v
}

// GetReleaseNumber returns the ReleaseNumber field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetReleaseNumber() int32 {
	if o == nil || IsNil(o.ReleaseNumber) {
		var ret int32
		return ret
	}
	return *o.ReleaseNumber
}

// GetReleaseNumberOk returns a tuple with the ReleaseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetReleaseNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleaseNumber) {
		return nil, false
	}
	return o.ReleaseNumber, true
}

// HasReleaseNumber returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasReleaseNumber() bool {
	if o != nil && !IsNil(o.ReleaseNumber) {
		return true
	}

	return false
}

// SetReleaseNumber gets a reference to the given int32 and assigns it to the ReleaseNumber field.
func (o *AzureProductSubmission) SetReleaseNumber(v int32) {
	o.ReleaseNumber = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AzureProductSubmission) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetResources() []AzureTypeValue {
	if o == nil || IsNil(o.Resources) {
		var ret []AzureTypeValue
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetResourcesOk() ([]AzureTypeValue, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []AzureTypeValue and assigns it to the Resources field.
func (o *AzureProductSubmission) SetResources(v []AzureTypeValue) {
	o.Resources = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AzureProductSubmission) SetState(v string) {
	o.State = &v
}

// GetSubState returns the SubState field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetSubState() string {
	if o == nil || IsNil(o.SubState) {
		var ret string
		return ret
	}
	return *o.SubState
}

// GetSubStateOk returns a tuple with the SubState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetSubStateOk() (*string, bool) {
	if o == nil || IsNil(o.SubState) {
		return nil, false
	}
	return o.SubState, true
}

// HasSubState returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasSubState() bool {
	if o != nil && !IsNil(o.SubState) {
		return true
	}

	return false
}

// SetSubState gets a reference to the given string and assigns it to the SubState field.
func (o *AzureProductSubmission) SetSubState(v string) {
	o.SubState = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetTargets() []AzureTypeValue {
	if o == nil || IsNil(o.Targets) {
		var ret []AzureTypeValue
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetTargetsOk() ([]AzureTypeValue, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []AzureTypeValue and assigns it to the Targets field.
func (o *AzureProductSubmission) SetTargets(v []AzureTypeValue) {
	o.Targets = v
}

// GetVariantResources returns the VariantResources field value if set, zero value otherwise.
func (o *AzureProductSubmission) GetVariantResources() []AzureVariantResource {
	if o == nil || IsNil(o.VariantResources) {
		var ret []AzureVariantResource
		return ret
	}
	return o.VariantResources
}

// GetVariantResourcesOk returns a tuple with the VariantResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSubmission) GetVariantResourcesOk() ([]AzureVariantResource, bool) {
	if o == nil || IsNil(o.VariantResources) {
		return nil, false
	}
	return o.VariantResources, true
}

// HasVariantResources returns a boolean if a field has been set.
func (o *AzureProductSubmission) HasVariantResources() bool {
	if o != nil && !IsNil(o.VariantResources) {
		return true
	}

	return false
}

// SetVariantResources gets a reference to the given []AzureVariantResource and assigns it to the VariantResources field.
func (o *AzureProductSubmission) SetVariantResources(v []AzureVariantResource) {
	o.VariantResources = v
}

func (o AzureProductSubmission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreResourcesReady) {
		toSerialize["areResourcesReady"] = o.AreResourcesReady
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendlyName"] = o.FriendlyName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PendingUpdateInfo) {
		toSerialize["pendingUpdateInfo"] = o.PendingUpdateInfo
	}
	if !IsNil(o.PublishedTimeInUtc) {
		toSerialize["publishedTimeInUtc"] = o.PublishedTimeInUtc
	}
	if !IsNil(o.ReleaseNumber) {
		toSerialize["releaseNumber"] = o.ReleaseNumber
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.SubState) {
		toSerialize["subState"] = o.SubState
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !IsNil(o.VariantResources) {
		toSerialize["variantResources"] = o.VariantResources
	}
	return toSerialize, nil
}

type NullableAzureProductSubmission struct {
	value *AzureProductSubmission
	isSet bool
}

func (v NullableAzureProductSubmission) Get() *AzureProductSubmission {
	return v.value
}

func (v *NullableAzureProductSubmission) Set(val *AzureProductSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductSubmission(val *AzureProductSubmission) *NullableAzureProductSubmission {
	return &NullableAzureProductSubmission{value: val, isSet: true}
}

func (v NullableAzureProductSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
