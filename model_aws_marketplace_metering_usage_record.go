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

// checks if the AwsMarketplaceMeteringUsageRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceMeteringUsageRecord{}

// AwsMarketplaceMeteringUsageRecord struct for AwsMarketplaceMeteringUsageRecord
type AwsMarketplaceMeteringUsageRecord struct {
	// The CustomerIdentifier is obtained through the ResolveCustomer operation and represents an individual buyer in your application.
	CustomerIdentifier *string `json:"CustomerIdentifier,omitempty"`
	// During the process of registering a product on AWS Marketplace, dimensions are specified. These represent different units of value in your application.
	Dimension *string `json:"Dimension,omitempty"`
	// The quantity of usage consumed by the customer for the given dimension and time. Defaults to 0 if not specified.
	Quantity *int32 `json:"Quantity,omitempty"`
	// Timestamp, in UTC, for which the usage is being reported. Your application can meter usage for up to one hour in the past. Make sure the timestamp value is not before the start of the software usage.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The set of UsageAllocations to submit. The sum of all UsageAllocation quantities must equal the Quantity of the UsageRecord.
	UsageAllocations []AwsMarketplaceMeteringUsageAllocation `json:"UsageAllocations,omitempty"`
}

// NewAwsMarketplaceMeteringUsageRecord instantiates a new AwsMarketplaceMeteringUsageRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceMeteringUsageRecord() *AwsMarketplaceMeteringUsageRecord {
	this := AwsMarketplaceMeteringUsageRecord{}
	return &this
}

// NewAwsMarketplaceMeteringUsageRecordWithDefaults instantiates a new AwsMarketplaceMeteringUsageRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceMeteringUsageRecordWithDefaults() *AwsMarketplaceMeteringUsageRecord {
	this := AwsMarketplaceMeteringUsageRecord{}
	return &this
}

// GetCustomerIdentifier returns the CustomerIdentifier field value if set, zero value otherwise.
func (o *AwsMarketplaceMeteringUsageRecord) GetCustomerIdentifier() string {
	if o == nil || IsNil(o.CustomerIdentifier) {
		var ret string
		return ret
	}
	return *o.CustomerIdentifier
}

// GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceMeteringUsageRecord) GetCustomerIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIdentifier) {
		return nil, false
	}
	return o.CustomerIdentifier, true
}

// HasCustomerIdentifier returns a boolean if a field has been set.
func (o *AwsMarketplaceMeteringUsageRecord) HasCustomerIdentifier() bool {
	if o != nil && !IsNil(o.CustomerIdentifier) {
		return true
	}

	return false
}

// SetCustomerIdentifier gets a reference to the given string and assigns it to the CustomerIdentifier field.
func (o *AwsMarketplaceMeteringUsageRecord) SetCustomerIdentifier(v string) {
	o.CustomerIdentifier = &v
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *AwsMarketplaceMeteringUsageRecord) GetDimension() string {
	if o == nil || IsNil(o.Dimension) {
		var ret string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceMeteringUsageRecord) GetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.Dimension) {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *AwsMarketplaceMeteringUsageRecord) HasDimension() bool {
	if o != nil && !IsNil(o.Dimension) {
		return true
	}

	return false
}

// SetDimension gets a reference to the given string and assigns it to the Dimension field.
func (o *AwsMarketplaceMeteringUsageRecord) SetDimension(v string) {
	o.Dimension = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AwsMarketplaceMeteringUsageRecord) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceMeteringUsageRecord) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AwsMarketplaceMeteringUsageRecord) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *AwsMarketplaceMeteringUsageRecord) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AwsMarketplaceMeteringUsageRecord) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceMeteringUsageRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AwsMarketplaceMeteringUsageRecord) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AwsMarketplaceMeteringUsageRecord) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUsageAllocations returns the UsageAllocations field value if set, zero value otherwise.
func (o *AwsMarketplaceMeteringUsageRecord) GetUsageAllocations() []AwsMarketplaceMeteringUsageAllocation {
	if o == nil || IsNil(o.UsageAllocations) {
		var ret []AwsMarketplaceMeteringUsageAllocation
		return ret
	}
	return o.UsageAllocations
}

// GetUsageAllocationsOk returns a tuple with the UsageAllocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceMeteringUsageRecord) GetUsageAllocationsOk() ([]AwsMarketplaceMeteringUsageAllocation, bool) {
	if o == nil || IsNil(o.UsageAllocations) {
		return nil, false
	}
	return o.UsageAllocations, true
}

// HasUsageAllocations returns a boolean if a field has been set.
func (o *AwsMarketplaceMeteringUsageRecord) HasUsageAllocations() bool {
	if o != nil && !IsNil(o.UsageAllocations) {
		return true
	}

	return false
}

// SetUsageAllocations gets a reference to the given []AwsMarketplaceMeteringUsageAllocation and assigns it to the UsageAllocations field.
func (o *AwsMarketplaceMeteringUsageRecord) SetUsageAllocations(v []AwsMarketplaceMeteringUsageAllocation) {
	o.UsageAllocations = v
}

func (o AwsMarketplaceMeteringUsageRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceMeteringUsageRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIdentifier) {
		toSerialize["CustomerIdentifier"] = o.CustomerIdentifier
	}
	if !IsNil(o.Dimension) {
		toSerialize["Dimension"] = o.Dimension
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if !IsNil(o.UsageAllocations) {
		toSerialize["UsageAllocations"] = o.UsageAllocations
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceMeteringUsageRecord struct {
	value *AwsMarketplaceMeteringUsageRecord
	isSet bool
}

func (v NullableAwsMarketplaceMeteringUsageRecord) Get() *AwsMarketplaceMeteringUsageRecord {
	return v.value
}

func (v *NullableAwsMarketplaceMeteringUsageRecord) Set(val *AwsMarketplaceMeteringUsageRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceMeteringUsageRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceMeteringUsageRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceMeteringUsageRecord(val *AwsMarketplaceMeteringUsageRecord) *NullableAwsMarketplaceMeteringUsageRecord {
	return &NullableAwsMarketplaceMeteringUsageRecord{value: val, isSet: true}
}

func (v NullableAwsMarketplaceMeteringUsageRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceMeteringUsageRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
