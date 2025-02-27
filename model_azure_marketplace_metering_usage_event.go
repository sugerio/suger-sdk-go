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

// checks if the AzureMarketplaceMeteringUsageEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceMeteringUsageEvent{}

// AzureMarketplaceMeteringUsageEvent struct for AzureMarketplaceMeteringUsageEvent
type AzureMarketplaceMeteringUsageEvent struct {
	// Dimension identifier
	Dimension *string `json:"dimension,omitempty"`
	// Time in UTC when the usage event occurred
	EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
	// Plan associated with the purchased offer
	PlanId *string `json:"planId,omitempty"`
	// Number of units consumed
	Quantity *float32 `json:"quantity,omitempty"`
	// subscriptionId property value for SaaS offer subscriptions; resourceUsageId property on the managed application resource for managed application offers. For managed applications, only use one of resourceId or resourceUri.
	ResourceId *string `json:"resourceId,omitempty"`
	// Resource URI for the managed app. Used with managed applications. Only use resourceUri or resourceId, but never both.
	ResourceUri *string `json:"resourceUri,omitempty"`
}

// NewAzureMarketplaceMeteringUsageEvent instantiates a new AzureMarketplaceMeteringUsageEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceMeteringUsageEvent() *AzureMarketplaceMeteringUsageEvent {
	this := AzureMarketplaceMeteringUsageEvent{}
	return &this
}

// NewAzureMarketplaceMeteringUsageEventWithDefaults instantiates a new AzureMarketplaceMeteringUsageEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceMeteringUsageEventWithDefaults() *AzureMarketplaceMeteringUsageEvent {
	this := AzureMarketplaceMeteringUsageEvent{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *AzureMarketplaceMeteringUsageEvent) SetDimension(v string) {
	o.Dimension = &v
}

// GetEffectiveStartTime returns the EffectiveStartTime field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetEffectiveStartTime() string {
	if o == nil || IsNil(o.EffectiveStartTime) {
		var ret string
		return ret
	}
	return *o.EffectiveStartTime
}

// GetEffectiveStartTimeOk returns a tuple with the EffectiveStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetEffectiveStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartTime) {
		return nil, false
	}
	return o.EffectiveStartTime, true
}

// HasEffectiveStartTime returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasEffectiveStartTime() bool {
	if o != nil && !IsNil(o.EffectiveStartTime) {
		return true
	}

	return false
}

// SetEffectiveStartTime gets a reference to the given string and assigns it to the EffectiveStartTime field.
func (o *AzureMarketplaceMeteringUsageEvent) SetEffectiveStartTime(v string) {
	o.EffectiveStartTime = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *AzureMarketplaceMeteringUsageEvent) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *AzureMarketplaceMeteringUsageEvent) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *AzureMarketplaceMeteringUsageEvent) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceUri returns the ResourceUri field value if set, zero value otherwise.
func (o *AzureMarketplaceMeteringUsageEvent) GetResourceUri() string {
	if o == nil || IsNil(o.ResourceUri) {
		var ret string
		return ret
	}
	return *o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceMeteringUsageEvent) GetResourceUriOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceUri) {
		return nil, false
	}
	return o.ResourceUri, true
}

// HasResourceUri returns a boolean if a field has been set.
func (o *AzureMarketplaceMeteringUsageEvent) HasResourceUri() bool {
	if o != nil && !IsNil(o.ResourceUri) {
		return true
	}

	return false
}

// SetResourceUri gets a reference to the given string and assigns it to the ResourceUri field.
func (o *AzureMarketplaceMeteringUsageEvent) SetResourceUri(v string) {
	o.ResourceUri = &v
}

func (o AzureMarketplaceMeteringUsageEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceMeteringUsageEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimension) {
		toSerialize["dimension"] = o.Dimension
	}
	if !IsNil(o.EffectiveStartTime) {
		toSerialize["effectiveStartTime"] = o.EffectiveStartTime
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
	return toSerialize, nil
}

type NullableAzureMarketplaceMeteringUsageEvent struct {
	value *AzureMarketplaceMeteringUsageEvent
	isSet bool
}

func (v NullableAzureMarketplaceMeteringUsageEvent) Get() *AzureMarketplaceMeteringUsageEvent {
	return v.value
}

func (v *NullableAzureMarketplaceMeteringUsageEvent) Set(val *AzureMarketplaceMeteringUsageEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceMeteringUsageEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceMeteringUsageEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceMeteringUsageEvent(val *AzureMarketplaceMeteringUsageEvent) *NullableAzureMarketplaceMeteringUsageEvent {
	return &NullableAzureMarketplaceMeteringUsageEvent{value: val, isSet: true}
}

func (v NullableAzureMarketplaceMeteringUsageEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceMeteringUsageEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
