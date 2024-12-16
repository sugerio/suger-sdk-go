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

// checks if the TypesUsageAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypesUsageAllocation{}

// TypesUsageAllocation struct for TypesUsageAllocation
type TypesUsageAllocation struct {
	// The total quantity allocated to this bucket of usage.  This member is required.
	AllocatedUsageQuantity *int32 `json:"allocatedUsageQuantity,omitempty"`
	// The set of tags that define the bucket of usage. For the bucket of items with no tags, this parameter can be left out.
	Tags []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag `json:"tags,omitempty"`
}

// NewTypesUsageAllocation instantiates a new TypesUsageAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesUsageAllocation() *TypesUsageAllocation {
	this := TypesUsageAllocation{}
	return &this
}

// NewTypesUsageAllocationWithDefaults instantiates a new TypesUsageAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesUsageAllocationWithDefaults() *TypesUsageAllocation {
	this := TypesUsageAllocation{}
	return &this
}

// GetAllocatedUsageQuantity returns the AllocatedUsageQuantity field value if set, zero value otherwise.
func (o *TypesUsageAllocation) GetAllocatedUsageQuantity() int32 {
	if o == nil || IsNil(o.AllocatedUsageQuantity) {
		var ret int32
		return ret
	}
	return *o.AllocatedUsageQuantity
}

// GetAllocatedUsageQuantityOk returns a tuple with the AllocatedUsageQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUsageAllocation) GetAllocatedUsageQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.AllocatedUsageQuantity) {
		return nil, false
	}
	return o.AllocatedUsageQuantity, true
}

// HasAllocatedUsageQuantity returns a boolean if a field has been set.
func (o *TypesUsageAllocation) HasAllocatedUsageQuantity() bool {
	if o != nil && !IsNil(o.AllocatedUsageQuantity) {
		return true
	}

	return false
}

// SetAllocatedUsageQuantity gets a reference to the given int32 and assigns it to the AllocatedUsageQuantity field.
func (o *TypesUsageAllocation) SetAllocatedUsageQuantity(v int32) {
	o.AllocatedUsageQuantity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TypesUsageAllocation) GetTags() []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag {
	if o == nil || IsNil(o.Tags) {
		var ret []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesUsageAllocation) GetTagsOk() ([]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TypesUsageAllocation) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag and assigns it to the Tags field.
func (o *TypesUsageAllocation) SetTags(v []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesTag) {
	o.Tags = v
}

func (o TypesUsageAllocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypesUsageAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllocatedUsageQuantity) {
		toSerialize["allocatedUsageQuantity"] = o.AllocatedUsageQuantity
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableTypesUsageAllocation struct {
	value *TypesUsageAllocation
	isSet bool
}

func (v NullableTypesUsageAllocation) Get() *TypesUsageAllocation {
	return v.value
}

func (v *NullableTypesUsageAllocation) Set(val *TypesUsageAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUsageAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUsageAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUsageAllocation(val *TypesUsageAllocation) *NullableTypesUsageAllocation {
	return &NullableTypesUsageAllocation{value: val, isSet: true}
}

func (v NullableTypesUsageAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUsageAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
