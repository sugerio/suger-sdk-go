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

// checks if the GcpMarketplaceProductInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductInfo{}

// GcpMarketplaceProductInfo struct for GcpMarketplaceProductInfo
type GcpMarketplaceProductInfo struct {
	// The service level (price plan) with time duration suffix, such as \"basic-plan-P1Y\"
	FlavorExternalName  *string `json:"flavorExternalName,omitempty"`
	ProductExternalName *string `json:"productExternalName,omitempty"`
	ProviderId          *string `json:"providerId,omitempty"`
	// The price plan, such as \"basic-plan\"
	ServiceLevel *string `json:"serviceLevel,omitempty"`
	ServiceName  *string `json:"serviceName,omitempty"`
}

// NewGcpMarketplaceProductInfo instantiates a new GcpMarketplaceProductInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductInfo() *GcpMarketplaceProductInfo {
	this := GcpMarketplaceProductInfo{}
	return &this
}

// NewGcpMarketplaceProductInfoWithDefaults instantiates a new GcpMarketplaceProductInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductInfoWithDefaults() *GcpMarketplaceProductInfo {
	this := GcpMarketplaceProductInfo{}
	return &this
}

// GetFlavorExternalName returns the FlavorExternalName field value if set, zero value otherwise.
func (o *GcpMarketplaceProductInfo) GetFlavorExternalName() string {
	if o == nil || IsNil(o.FlavorExternalName) {
		var ret string
		return ret
	}
	return *o.FlavorExternalName
}

// GetFlavorExternalNameOk returns a tuple with the FlavorExternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductInfo) GetFlavorExternalNameOk() (*string, bool) {
	if o == nil || IsNil(o.FlavorExternalName) {
		return nil, false
	}
	return o.FlavorExternalName, true
}

// HasFlavorExternalName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductInfo) HasFlavorExternalName() bool {
	if o != nil && !IsNil(o.FlavorExternalName) {
		return true
	}

	return false
}

// SetFlavorExternalName gets a reference to the given string and assigns it to the FlavorExternalName field.
func (o *GcpMarketplaceProductInfo) SetFlavorExternalName(v string) {
	o.FlavorExternalName = &v
}

// GetProductExternalName returns the ProductExternalName field value if set, zero value otherwise.
func (o *GcpMarketplaceProductInfo) GetProductExternalName() string {
	if o == nil || IsNil(o.ProductExternalName) {
		var ret string
		return ret
	}
	return *o.ProductExternalName
}

// GetProductExternalNameOk returns a tuple with the ProductExternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductInfo) GetProductExternalNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductExternalName) {
		return nil, false
	}
	return o.ProductExternalName, true
}

// HasProductExternalName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductInfo) HasProductExternalName() bool {
	if o != nil && !IsNil(o.ProductExternalName) {
		return true
	}

	return false
}

// SetProductExternalName gets a reference to the given string and assigns it to the ProductExternalName field.
func (o *GcpMarketplaceProductInfo) SetProductExternalName(v string) {
	o.ProductExternalName = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *GcpMarketplaceProductInfo) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductInfo) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *GcpMarketplaceProductInfo) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *GcpMarketplaceProductInfo) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *GcpMarketplaceProductInfo) GetServiceLevel() string {
	if o == nil || IsNil(o.ServiceLevel) {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductInfo) GetServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *GcpMarketplaceProductInfo) HasServiceLevel() bool {
	if o != nil && !IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *GcpMarketplaceProductInfo) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *GcpMarketplaceProductInfo) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductInfo) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductInfo) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *GcpMarketplaceProductInfo) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o GcpMarketplaceProductInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlavorExternalName) {
		toSerialize["flavorExternalName"] = o.FlavorExternalName
	}
	if !IsNil(o.ProductExternalName) {
		toSerialize["productExternalName"] = o.ProductExternalName
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.ServiceLevel) {
		toSerialize["serviceLevel"] = o.ServiceLevel
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductInfo struct {
	value *GcpMarketplaceProductInfo
	isSet bool
}

func (v NullableGcpMarketplaceProductInfo) Get() *GcpMarketplaceProductInfo {
	return v.value
}

func (v *NullableGcpMarketplaceProductInfo) Set(val *GcpMarketplaceProductInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductInfo(val *GcpMarketplaceProductInfo) *NullableGcpMarketplaceProductInfo {
	return &NullableGcpMarketplaceProductInfo{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
