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

// checks if the SlackOAuthResponseIncomingWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackOAuthResponseIncomingWebhook{}

// SlackOAuthResponseIncomingWebhook struct for SlackOAuthResponseIncomingWebhook
type SlackOAuthResponseIncomingWebhook struct {
	Channel          *string `json:"channel,omitempty"`
	ChannelId        *string `json:"channel_id,omitempty"`
	ConfigurationUrl *string `json:"configuration_url,omitempty"`
	Url              *string `json:"url,omitempty"`
}

// NewSlackOAuthResponseIncomingWebhook instantiates a new SlackOAuthResponseIncomingWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackOAuthResponseIncomingWebhook() *SlackOAuthResponseIncomingWebhook {
	this := SlackOAuthResponseIncomingWebhook{}
	return &this
}

// NewSlackOAuthResponseIncomingWebhookWithDefaults instantiates a new SlackOAuthResponseIncomingWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackOAuthResponseIncomingWebhookWithDefaults() *SlackOAuthResponseIncomingWebhook {
	this := SlackOAuthResponseIncomingWebhook{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SlackOAuthResponseIncomingWebhook) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackOAuthResponseIncomingWebhook) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SlackOAuthResponseIncomingWebhook) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *SlackOAuthResponseIncomingWebhook) SetChannel(v string) {
	o.Channel = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *SlackOAuthResponseIncomingWebhook) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackOAuthResponseIncomingWebhook) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *SlackOAuthResponseIncomingWebhook) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *SlackOAuthResponseIncomingWebhook) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetConfigurationUrl returns the ConfigurationUrl field value if set, zero value otherwise.
func (o *SlackOAuthResponseIncomingWebhook) GetConfigurationUrl() string {
	if o == nil || IsNil(o.ConfigurationUrl) {
		var ret string
		return ret
	}
	return *o.ConfigurationUrl
}

// GetConfigurationUrlOk returns a tuple with the ConfigurationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackOAuthResponseIncomingWebhook) GetConfigurationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationUrl) {
		return nil, false
	}
	return o.ConfigurationUrl, true
}

// HasConfigurationUrl returns a boolean if a field has been set.
func (o *SlackOAuthResponseIncomingWebhook) HasConfigurationUrl() bool {
	if o != nil && !IsNil(o.ConfigurationUrl) {
		return true
	}

	return false
}

// SetConfigurationUrl gets a reference to the given string and assigns it to the ConfigurationUrl field.
func (o *SlackOAuthResponseIncomingWebhook) SetConfigurationUrl(v string) {
	o.ConfigurationUrl = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SlackOAuthResponseIncomingWebhook) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackOAuthResponseIncomingWebhook) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SlackOAuthResponseIncomingWebhook) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SlackOAuthResponseIncomingWebhook) SetUrl(v string) {
	o.Url = &v
}

func (o SlackOAuthResponseIncomingWebhook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackOAuthResponseIncomingWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.ConfigurationUrl) {
		toSerialize["configuration_url"] = o.ConfigurationUrl
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSlackOAuthResponseIncomingWebhook struct {
	value *SlackOAuthResponseIncomingWebhook
	isSet bool
}

func (v NullableSlackOAuthResponseIncomingWebhook) Get() *SlackOAuthResponseIncomingWebhook {
	return v.value
}

func (v *NullableSlackOAuthResponseIncomingWebhook) Set(val *SlackOAuthResponseIncomingWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackOAuthResponseIncomingWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackOAuthResponseIncomingWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackOAuthResponseIncomingWebhook(val *SlackOAuthResponseIncomingWebhook) *NullableSlackOAuthResponseIncomingWebhook {
	return &NullableSlackOAuthResponseIncomingWebhook{value: val, isSet: true}
}

func (v NullableSlackOAuthResponseIncomingWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackOAuthResponseIncomingWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
