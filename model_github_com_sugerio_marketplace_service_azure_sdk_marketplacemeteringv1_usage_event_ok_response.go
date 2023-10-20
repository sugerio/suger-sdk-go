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

// checks if the GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse{}

// GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse struct for GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse
type GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse struct {
	// Dimension identifier
	Dimension *string `json:"dimension,omitempty"`
	// Time in UTC when the usage event occurred
	EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
	// Time this message was created in UTC
	MessageTime *string `json:"messageTime,omitempty"`
	// Plan associated with the purchased offer
	PlanId *string `json:"planId,omitempty"`
	// Number of units consumed
	Quantity *float32 `json:"quantity,omitempty"`
	// Identifier of the resource against which usage is emitted
	ResourceId *string `json:"resourceId,omitempty"`
	// Identifier of the managed app resource against which usage is emitted
	ResourceUri *string `json:"resourceUri,omitempty"`
	Status *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum `json:"status,omitempty"`
	// Unique identifier associated with the usage event
	UsageEventId *string `json:"usageEventId,omitempty"`
}

// NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse instantiates a new GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse() *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse {
	this := GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse{}
	return &this
}

// NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponseWithDefaults instantiates a new GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponseWithDefaults() *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse {
	this := GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetDimension(v string) {
	o.Dimension = &v
}

// GetEffectiveStartTime returns the EffectiveStartTime field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetEffectiveStartTime() string {
	if o == nil || IsNil(o.EffectiveStartTime) {
		var ret string
		return ret
	}
	return *o.EffectiveStartTime
}

// GetEffectiveStartTimeOk returns a tuple with the EffectiveStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetEffectiveStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartTime) {
		return nil, false
	}
	return o.EffectiveStartTime, true
}

// HasEffectiveStartTime returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasEffectiveStartTime() bool {
	if o != nil && !IsNil(o.EffectiveStartTime) {
		return true
	}

	return false
}

// SetEffectiveStartTime gets a reference to the given string and assigns it to the EffectiveStartTime field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetEffectiveStartTime(v string) {
	o.EffectiveStartTime = &v
}

// GetMessageTime returns the MessageTime field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetMessageTime() string {
	if o == nil || IsNil(o.MessageTime) {
		var ret string
		return ret
	}
	return *o.MessageTime
}

// GetMessageTimeOk returns a tuple with the MessageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetMessageTimeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageTime) {
		return nil, false
	}
	return o.MessageTime, true
}

// HasMessageTime returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasMessageTime() bool {
	if o != nil && !IsNil(o.MessageTime) {
		return true
	}

	return false
}

// SetMessageTime gets a reference to the given string and assigns it to the MessageTime field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetMessageTime(v string) {
	o.MessageTime = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceUri returns the ResourceUri field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceUri() string {
	if o == nil || IsNil(o.ResourceUri) {
		var ret string
		return ret
	}
	return *o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceUriOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceUri) {
		return nil, false
	}
	return o.ResourceUri, true
}

// HasResourceUri returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasResourceUri() bool {
	if o != nil && !IsNil(o.ResourceUri) {
		return true
	}

	return false
}

// SetResourceUri gets a reference to the given string and assigns it to the ResourceUri field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetResourceUri(v string) {
	o.ResourceUri = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetStatus() GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetStatusOk() (*GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum and assigns it to the Status field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetStatus(v GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum) {
	o.Status = &v
}

// GetUsageEventId returns the UsageEventId field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetUsageEventId() string {
	if o == nil || IsNil(o.UsageEventId) {
		var ret string
		return ret
	}
	return *o.UsageEventId
}

// GetUsageEventIdOk returns a tuple with the UsageEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetUsageEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.UsageEventId) {
		return nil, false
	}
	return o.UsageEventId, true
}

// HasUsageEventId returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasUsageEventId() bool {
	if o != nil && !IsNil(o.UsageEventId) {
		return true
	}

	return false
}

// SetUsageEventId gets a reference to the given string and assigns it to the UsageEventId field.
func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetUsageEventId(v string) {
	o.UsageEventId = &v
}

func (o GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.EffectiveStartTime) {
		toSerialize["effectiveStartTime"] = o.EffectiveStartTime
	}
	if !IsNil(o.MessageTime) {
		toSerialize["messageTime"] = o.MessageTime
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceUri) {
		toSerialize["resourceUri"] = o.ResourceUri
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UsageEventId) {
		toSerialize["usageEventId"] = o.UsageEventId
	}
	return toSerialize, nil
}

type NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse struct {
	value *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse
	isSet bool
}

func (v NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) Get() *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse {
	return v.value
}

func (v *NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) Set(val *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse(val *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) *NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse {
	return &NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse{value: val, isSet: true}
}

func (v NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


