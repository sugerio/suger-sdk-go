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

// checks if the AwsSnsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsSnsSubscription{}

// AwsSnsSubscription struct for AwsSnsSubscription
type AwsSnsSubscription struct {
	Endpoint *string `json:"Endpoint,omitempty"`
	Protocol *string `json:"Protocol,omitempty"`
	Status *AwsSnsSubscriptionStatus `json:"Status,omitempty"`
	SubscriptionArn *string `json:"SubscriptionArn,omitempty"`
	TopicArn *string `json:"TopicArn,omitempty"`
}

// NewAwsSnsSubscription instantiates a new AwsSnsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSnsSubscription() *AwsSnsSubscription {
	this := AwsSnsSubscription{}
	return &this
}

// NewAwsSnsSubscriptionWithDefaults instantiates a new AwsSnsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSnsSubscriptionWithDefaults() *AwsSnsSubscription {
	this := AwsSnsSubscription{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AwsSnsSubscription) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnsSubscription) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AwsSnsSubscription) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AwsSnsSubscription) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AwsSnsSubscription) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnsSubscription) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AwsSnsSubscription) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AwsSnsSubscription) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsSnsSubscription) GetStatus() AwsSnsSubscriptionStatus {
	if o == nil || IsNil(o.Status) {
		var ret AwsSnsSubscriptionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnsSubscription) GetStatusOk() (*AwsSnsSubscriptionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsSnsSubscription) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AwsSnsSubscriptionStatus and assigns it to the Status field.
func (o *AwsSnsSubscription) SetStatus(v AwsSnsSubscriptionStatus) {
	o.Status = &v
}

// GetSubscriptionArn returns the SubscriptionArn field value if set, zero value otherwise.
func (o *AwsSnsSubscription) GetSubscriptionArn() string {
	if o == nil || IsNil(o.SubscriptionArn) {
		var ret string
		return ret
	}
	return *o.SubscriptionArn
}

// GetSubscriptionArnOk returns a tuple with the SubscriptionArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnsSubscription) GetSubscriptionArnOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionArn) {
		return nil, false
	}
	return o.SubscriptionArn, true
}

// HasSubscriptionArn returns a boolean if a field has been set.
func (o *AwsSnsSubscription) HasSubscriptionArn() bool {
	if o != nil && !IsNil(o.SubscriptionArn) {
		return true
	}

	return false
}

// SetSubscriptionArn gets a reference to the given string and assigns it to the SubscriptionArn field.
func (o *AwsSnsSubscription) SetSubscriptionArn(v string) {
	o.SubscriptionArn = &v
}

// GetTopicArn returns the TopicArn field value if set, zero value otherwise.
func (o *AwsSnsSubscription) GetTopicArn() string {
	if o == nil || IsNil(o.TopicArn) {
		var ret string
		return ret
	}
	return *o.TopicArn
}

// GetTopicArnOk returns a tuple with the TopicArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSnsSubscription) GetTopicArnOk() (*string, bool) {
	if o == nil || IsNil(o.TopicArn) {
		return nil, false
	}
	return o.TopicArn, true
}

// HasTopicArn returns a boolean if a field has been set.
func (o *AwsSnsSubscription) HasTopicArn() bool {
	if o != nil && !IsNil(o.TopicArn) {
		return true
	}

	return false
}

// SetTopicArn gets a reference to the given string and assigns it to the TopicArn field.
func (o *AwsSnsSubscription) SetTopicArn(v string) {
	o.TopicArn = &v
}

func (o AwsSnsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsSnsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["Endpoint"] = o.Endpoint
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.SubscriptionArn) {
		toSerialize["SubscriptionArn"] = o.SubscriptionArn
	}
	if !IsNil(o.TopicArn) {
		toSerialize["TopicArn"] = o.TopicArn
	}
	return toSerialize, nil
}

type NullableAwsSnsSubscription struct {
	value *AwsSnsSubscription
	isSet bool
}

func (v NullableAwsSnsSubscription) Get() *AwsSnsSubscription {
	return v.value
}

func (v *NullableAwsSnsSubscription) Set(val *AwsSnsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSnsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSnsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSnsSubscription(val *AwsSnsSubscription) *NullableAwsSnsSubscription {
	return &NullableAwsSnsSubscription{value: val, isSet: true}
}

func (v NullableAwsSnsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSnsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


