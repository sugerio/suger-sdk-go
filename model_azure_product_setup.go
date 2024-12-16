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

// checks if the AzureProductSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductSetup{}

// AzureProductSetup struct for AzureProductSetup
type AzureProductSetup struct {
	CallToAction    *string          `json:"callToAction,omitempty"`
	ChannelStates   []AzureTypeValue `json:"channelStates,omitempty"`
	EnableTestDrive *bool            `json:"enableTestDrive,omitempty"`
	ResourceType    *string          `json:"resourceType,omitempty"`
	SellingOption   *string          `json:"sellingOption,omitempty"`
	TestDriveType   *string          `json:"testDriveType,omitempty"`
	TrialUri        *string          `json:"trialUri,omitempty"`
}

// NewAzureProductSetup instantiates a new AzureProductSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductSetup() *AzureProductSetup {
	this := AzureProductSetup{}
	return &this
}

// NewAzureProductSetupWithDefaults instantiates a new AzureProductSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductSetupWithDefaults() *AzureProductSetup {
	this := AzureProductSetup{}
	return &this
}

// GetCallToAction returns the CallToAction field value if set, zero value otherwise.
func (o *AzureProductSetup) GetCallToAction() string {
	if o == nil || IsNil(o.CallToAction) {
		var ret string
		return ret
	}
	return *o.CallToAction
}

// GetCallToActionOk returns a tuple with the CallToAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetCallToActionOk() (*string, bool) {
	if o == nil || IsNil(o.CallToAction) {
		return nil, false
	}
	return o.CallToAction, true
}

// HasCallToAction returns a boolean if a field has been set.
func (o *AzureProductSetup) HasCallToAction() bool {
	if o != nil && !IsNil(o.CallToAction) {
		return true
	}

	return false
}

// SetCallToAction gets a reference to the given string and assigns it to the CallToAction field.
func (o *AzureProductSetup) SetCallToAction(v string) {
	o.CallToAction = &v
}

// GetChannelStates returns the ChannelStates field value if set, zero value otherwise.
func (o *AzureProductSetup) GetChannelStates() []AzureTypeValue {
	if o == nil || IsNil(o.ChannelStates) {
		var ret []AzureTypeValue
		return ret
	}
	return o.ChannelStates
}

// GetChannelStatesOk returns a tuple with the ChannelStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetChannelStatesOk() ([]AzureTypeValue, bool) {
	if o == nil || IsNil(o.ChannelStates) {
		return nil, false
	}
	return o.ChannelStates, true
}

// HasChannelStates returns a boolean if a field has been set.
func (o *AzureProductSetup) HasChannelStates() bool {
	if o != nil && !IsNil(o.ChannelStates) {
		return true
	}

	return false
}

// SetChannelStates gets a reference to the given []AzureTypeValue and assigns it to the ChannelStates field.
func (o *AzureProductSetup) SetChannelStates(v []AzureTypeValue) {
	o.ChannelStates = v
}

// GetEnableTestDrive returns the EnableTestDrive field value if set, zero value otherwise.
func (o *AzureProductSetup) GetEnableTestDrive() bool {
	if o == nil || IsNil(o.EnableTestDrive) {
		var ret bool
		return ret
	}
	return *o.EnableTestDrive
}

// GetEnableTestDriveOk returns a tuple with the EnableTestDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetEnableTestDriveOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTestDrive) {
		return nil, false
	}
	return o.EnableTestDrive, true
}

// HasEnableTestDrive returns a boolean if a field has been set.
func (o *AzureProductSetup) HasEnableTestDrive() bool {
	if o != nil && !IsNil(o.EnableTestDrive) {
		return true
	}

	return false
}

// SetEnableTestDrive gets a reference to the given bool and assigns it to the EnableTestDrive field.
func (o *AzureProductSetup) SetEnableTestDrive(v bool) {
	o.EnableTestDrive = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AzureProductSetup) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AzureProductSetup) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AzureProductSetup) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSellingOption returns the SellingOption field value if set, zero value otherwise.
func (o *AzureProductSetup) GetSellingOption() string {
	if o == nil || IsNil(o.SellingOption) {
		var ret string
		return ret
	}
	return *o.SellingOption
}

// GetSellingOptionOk returns a tuple with the SellingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetSellingOptionOk() (*string, bool) {
	if o == nil || IsNil(o.SellingOption) {
		return nil, false
	}
	return o.SellingOption, true
}

// HasSellingOption returns a boolean if a field has been set.
func (o *AzureProductSetup) HasSellingOption() bool {
	if o != nil && !IsNil(o.SellingOption) {
		return true
	}

	return false
}

// SetSellingOption gets a reference to the given string and assigns it to the SellingOption field.
func (o *AzureProductSetup) SetSellingOption(v string) {
	o.SellingOption = &v
}

// GetTestDriveType returns the TestDriveType field value if set, zero value otherwise.
func (o *AzureProductSetup) GetTestDriveType() string {
	if o == nil || IsNil(o.TestDriveType) {
		var ret string
		return ret
	}
	return *o.TestDriveType
}

// GetTestDriveTypeOk returns a tuple with the TestDriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetTestDriveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TestDriveType) {
		return nil, false
	}
	return o.TestDriveType, true
}

// HasTestDriveType returns a boolean if a field has been set.
func (o *AzureProductSetup) HasTestDriveType() bool {
	if o != nil && !IsNil(o.TestDriveType) {
		return true
	}

	return false
}

// SetTestDriveType gets a reference to the given string and assigns it to the TestDriveType field.
func (o *AzureProductSetup) SetTestDriveType(v string) {
	o.TestDriveType = &v
}

// GetTrialUri returns the TrialUri field value if set, zero value otherwise.
func (o *AzureProductSetup) GetTrialUri() string {
	if o == nil || IsNil(o.TrialUri) {
		var ret string
		return ret
	}
	return *o.TrialUri
}

// GetTrialUriOk returns a tuple with the TrialUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductSetup) GetTrialUriOk() (*string, bool) {
	if o == nil || IsNil(o.TrialUri) {
		return nil, false
	}
	return o.TrialUri, true
}

// HasTrialUri returns a boolean if a field has been set.
func (o *AzureProductSetup) HasTrialUri() bool {
	if o != nil && !IsNil(o.TrialUri) {
		return true
	}

	return false
}

// SetTrialUri gets a reference to the given string and assigns it to the TrialUri field.
func (o *AzureProductSetup) SetTrialUri(v string) {
	o.TrialUri = &v
}

func (o AzureProductSetup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallToAction) {
		toSerialize["callToAction"] = o.CallToAction
	}
	if !IsNil(o.ChannelStates) {
		toSerialize["channelStates"] = o.ChannelStates
	}
	if !IsNil(o.EnableTestDrive) {
		toSerialize["enableTestDrive"] = o.EnableTestDrive
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.SellingOption) {
		toSerialize["sellingOption"] = o.SellingOption
	}
	if !IsNil(o.TestDriveType) {
		toSerialize["testDriveType"] = o.TestDriveType
	}
	if !IsNil(o.TrialUri) {
		toSerialize["trialUri"] = o.TrialUri
	}
	return toSerialize, nil
}

type NullableAzureProductSetup struct {
	value *AzureProductSetup
	isSet bool
}

func (v NullableAzureProductSetup) Get() *AzureProductSetup {
	return v.value
}

func (v *NullableAzureProductSetup) Set(val *AzureProductSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductSetup(val *AzureProductSetup) *NullableAzureProductSetup {
	return &NullableAzureProductSetup{value: val, isSet: true}
}

func (v NullableAzureProductSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
