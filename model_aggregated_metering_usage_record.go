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

// checks if the AggregatedMeteringUsageRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregatedMeteringUsageRecord{}

// AggregatedMeteringUsageRecord struct for AggregatedMeteringUsageRecord
type AggregatedMeteringUsageRecord struct {
	// Amount calculated by billable dimension's price model, this is only used for report billable usage records to marketplace.
	Amount                        *float32                       `json:"amount,omitempty"`
	BillableMetricAggregationType *BillableMetricAggregationType `json:"billableMetricAggregationType,omitempty"`
	BillableMetricInfo            *BillableMetricInfo            `json:"billableMetricInfo,omitempty"`
	// GroupBysExpression is the string expression of array of group bys.
	GroupBysExpression *string `json:"groupBysExpression,omitempty"`
	// Key is the unique identifier of a billable metric.
	Key *string `json:"key,omitempty"`
	// Name is the name of a billable metric. Optional, it is only for display purpose.
	Name *string `json:"name,omitempty"`
	// Value is the value of a billable metric.
	Quantity *float32 `json:"quantity,omitempty"`
	// Unique count metric aggregate result.
	UniqueCountAggregationResult *UniqueCountAggregationResult `json:"uniqueCountAggregationResult,omitempty"`
}

// NewAggregatedMeteringUsageRecord instantiates a new AggregatedMeteringUsageRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedMeteringUsageRecord() *AggregatedMeteringUsageRecord {
	this := AggregatedMeteringUsageRecord{}
	return &this
}

// NewAggregatedMeteringUsageRecordWithDefaults instantiates a new AggregatedMeteringUsageRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedMeteringUsageRecordWithDefaults() *AggregatedMeteringUsageRecord {
	this := AggregatedMeteringUsageRecord{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *AggregatedMeteringUsageRecord) SetAmount(v float32) {
	o.Amount = &v
}

// GetBillableMetricAggregationType returns the BillableMetricAggregationType field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetBillableMetricAggregationType() BillableMetricAggregationType {
	if o == nil || IsNil(o.BillableMetricAggregationType) {
		var ret BillableMetricAggregationType
		return ret
	}
	return *o.BillableMetricAggregationType
}

// GetBillableMetricAggregationTypeOk returns a tuple with the BillableMetricAggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetBillableMetricAggregationTypeOk() (*BillableMetricAggregationType, bool) {
	if o == nil || IsNil(o.BillableMetricAggregationType) {
		return nil, false
	}
	return o.BillableMetricAggregationType, true
}

// HasBillableMetricAggregationType returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasBillableMetricAggregationType() bool {
	if o != nil && !IsNil(o.BillableMetricAggregationType) {
		return true
	}

	return false
}

// SetBillableMetricAggregationType gets a reference to the given BillableMetricAggregationType and assigns it to the BillableMetricAggregationType field.
func (o *AggregatedMeteringUsageRecord) SetBillableMetricAggregationType(v BillableMetricAggregationType) {
	o.BillableMetricAggregationType = &v
}

// GetBillableMetricInfo returns the BillableMetricInfo field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetBillableMetricInfo() BillableMetricInfo {
	if o == nil || IsNil(o.BillableMetricInfo) {
		var ret BillableMetricInfo
		return ret
	}
	return *o.BillableMetricInfo
}

// GetBillableMetricInfoOk returns a tuple with the BillableMetricInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetBillableMetricInfoOk() (*BillableMetricInfo, bool) {
	if o == nil || IsNil(o.BillableMetricInfo) {
		return nil, false
	}
	return o.BillableMetricInfo, true
}

// HasBillableMetricInfo returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasBillableMetricInfo() bool {
	if o != nil && !IsNil(o.BillableMetricInfo) {
		return true
	}

	return false
}

// SetBillableMetricInfo gets a reference to the given BillableMetricInfo and assigns it to the BillableMetricInfo field.
func (o *AggregatedMeteringUsageRecord) SetBillableMetricInfo(v BillableMetricInfo) {
	o.BillableMetricInfo = &v
}

// GetGroupBysExpression returns the GroupBysExpression field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetGroupBysExpression() string {
	if o == nil || IsNil(o.GroupBysExpression) {
		var ret string
		return ret
	}
	return *o.GroupBysExpression
}

// GetGroupBysExpressionOk returns a tuple with the GroupBysExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetGroupBysExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBysExpression) {
		return nil, false
	}
	return o.GroupBysExpression, true
}

// HasGroupBysExpression returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasGroupBysExpression() bool {
	if o != nil && !IsNil(o.GroupBysExpression) {
		return true
	}

	return false
}

// SetGroupBysExpression gets a reference to the given string and assigns it to the GroupBysExpression field.
func (o *AggregatedMeteringUsageRecord) SetGroupBysExpression(v string) {
	o.GroupBysExpression = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AggregatedMeteringUsageRecord) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AggregatedMeteringUsageRecord) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *AggregatedMeteringUsageRecord) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUniqueCountAggregationResult returns the UniqueCountAggregationResult field value if set, zero value otherwise.
func (o *AggregatedMeteringUsageRecord) GetUniqueCountAggregationResult() UniqueCountAggregationResult {
	if o == nil || IsNil(o.UniqueCountAggregationResult) {
		var ret UniqueCountAggregationResult
		return ret
	}
	return *o.UniqueCountAggregationResult
}

// GetUniqueCountAggregationResultOk returns a tuple with the UniqueCountAggregationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedMeteringUsageRecord) GetUniqueCountAggregationResultOk() (*UniqueCountAggregationResult, bool) {
	if o == nil || IsNil(o.UniqueCountAggregationResult) {
		return nil, false
	}
	return o.UniqueCountAggregationResult, true
}

// HasUniqueCountAggregationResult returns a boolean if a field has been set.
func (o *AggregatedMeteringUsageRecord) HasUniqueCountAggregationResult() bool {
	if o != nil && !IsNil(o.UniqueCountAggregationResult) {
		return true
	}

	return false
}

// SetUniqueCountAggregationResult gets a reference to the given UniqueCountAggregationResult and assigns it to the UniqueCountAggregationResult field.
func (o *AggregatedMeteringUsageRecord) SetUniqueCountAggregationResult(v UniqueCountAggregationResult) {
	o.UniqueCountAggregationResult = &v
}

func (o AggregatedMeteringUsageRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregatedMeteringUsageRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BillableMetricAggregationType) {
		toSerialize["billableMetricAggregationType"] = o.BillableMetricAggregationType
	}
	if !IsNil(o.BillableMetricInfo) {
		toSerialize["billableMetricInfo"] = o.BillableMetricInfo
	}
	if !IsNil(o.GroupBysExpression) {
		toSerialize["groupBysExpression"] = o.GroupBysExpression
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UniqueCountAggregationResult) {
		toSerialize["uniqueCountAggregationResult"] = o.UniqueCountAggregationResult
	}
	return toSerialize, nil
}

type NullableAggregatedMeteringUsageRecord struct {
	value *AggregatedMeteringUsageRecord
	isSet bool
}

func (v NullableAggregatedMeteringUsageRecord) Get() *AggregatedMeteringUsageRecord {
	return v.value
}

func (v *NullableAggregatedMeteringUsageRecord) Set(val *AggregatedMeteringUsageRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedMeteringUsageRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedMeteringUsageRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedMeteringUsageRecord(val *AggregatedMeteringUsageRecord) *NullableAggregatedMeteringUsageRecord {
	return &NullableAggregatedMeteringUsageRecord{value: val, isSet: true}
}

func (v NullableAggregatedMeteringUsageRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedMeteringUsageRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
