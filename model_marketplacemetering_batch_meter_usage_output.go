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

// checks if the MarketplacemeteringBatchMeterUsageOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplacemeteringBatchMeterUsageOutput{}

// MarketplacemeteringBatchMeterUsageOutput struct for MarketplacemeteringBatchMeterUsageOutput
type MarketplacemeteringBatchMeterUsageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata map[string]interface{} `json:"resultMetadata,omitempty"`
	// Contains all UsageRecords processed by BatchMeterUsage . These records were either honored by AWS Marketplace Metering Service or were invalid. Invalid records should be fixed before being resubmitted.
	Results []TypesUsageRecordResult `json:"results,omitempty"`
	// Contains all UsageRecords that were not processed by BatchMeterUsage . This is a list of UsageRecords . You can retry the failed request by making another BatchMeterUsage call with this list as input in the BatchMeterUsageRequest .
	UnprocessedRecords []TypesUsageRecord `json:"unprocessedRecords,omitempty"`
}

// NewMarketplacemeteringBatchMeterUsageOutput instantiates a new MarketplacemeteringBatchMeterUsageOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplacemeteringBatchMeterUsageOutput() *MarketplacemeteringBatchMeterUsageOutput {
	this := MarketplacemeteringBatchMeterUsageOutput{}
	return &this
}

// NewMarketplacemeteringBatchMeterUsageOutputWithDefaults instantiates a new MarketplacemeteringBatchMeterUsageOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplacemeteringBatchMeterUsageOutputWithDefaults() *MarketplacemeteringBatchMeterUsageOutput {
	this := MarketplacemeteringBatchMeterUsageOutput{}
	return &this
}

// GetResultMetadata returns the ResultMetadata field value if set, zero value otherwise.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultMetadata() map[string]interface{} {
	if o == nil || IsNil(o.ResultMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResultMetadata
}

// GetResultMetadataOk returns a tuple with the ResultMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResultMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ResultMetadata, true
}

// HasResultMetadata returns a boolean if a field has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) HasResultMetadata() bool {
	if o != nil && !IsNil(o.ResultMetadata) {
		return true
	}

	return false
}

// SetResultMetadata gets a reference to the given map[string]interface{} and assigns it to the ResultMetadata field.
func (o *MarketplacemeteringBatchMeterUsageOutput) SetResultMetadata(v map[string]interface{}) {
	o.ResultMetadata = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetResults() []TypesUsageRecordResult {
	if o == nil || IsNil(o.Results) {
		var ret []TypesUsageRecordResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultsOk() ([]TypesUsageRecordResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TypesUsageRecordResult and assigns it to the Results field.
func (o *MarketplacemeteringBatchMeterUsageOutput) SetResults(v []TypesUsageRecordResult) {
	o.Results = v
}

// GetUnprocessedRecords returns the UnprocessedRecords field value if set, zero value otherwise.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecords() []TypesUsageRecord {
	if o == nil || IsNil(o.UnprocessedRecords) {
		var ret []TypesUsageRecord
		return ret
	}
	return o.UnprocessedRecords
}

// GetUnprocessedRecordsOk returns a tuple with the UnprocessedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecordsOk() ([]TypesUsageRecord, bool) {
	if o == nil || IsNil(o.UnprocessedRecords) {
		return nil, false
	}
	return o.UnprocessedRecords, true
}

// HasUnprocessedRecords returns a boolean if a field has been set.
func (o *MarketplacemeteringBatchMeterUsageOutput) HasUnprocessedRecords() bool {
	if o != nil && !IsNil(o.UnprocessedRecords) {
		return true
	}

	return false
}

// SetUnprocessedRecords gets a reference to the given []TypesUsageRecord and assigns it to the UnprocessedRecords field.
func (o *MarketplacemeteringBatchMeterUsageOutput) SetUnprocessedRecords(v []TypesUsageRecord) {
	o.UnprocessedRecords = v
}

func (o MarketplacemeteringBatchMeterUsageOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplacemeteringBatchMeterUsageOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultMetadata) {
		toSerialize["resultMetadata"] = o.ResultMetadata
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.UnprocessedRecords) {
		toSerialize["unprocessedRecords"] = o.UnprocessedRecords
	}
	return toSerialize, nil
}

type NullableMarketplacemeteringBatchMeterUsageOutput struct {
	value *MarketplacemeteringBatchMeterUsageOutput
	isSet bool
}

func (v NullableMarketplacemeteringBatchMeterUsageOutput) Get() *MarketplacemeteringBatchMeterUsageOutput {
	return v.value
}

func (v *NullableMarketplacemeteringBatchMeterUsageOutput) Set(val *MarketplacemeteringBatchMeterUsageOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplacemeteringBatchMeterUsageOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplacemeteringBatchMeterUsageOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplacemeteringBatchMeterUsageOutput(val *MarketplacemeteringBatchMeterUsageOutput) *NullableMarketplacemeteringBatchMeterUsageOutput {
	return &NullableMarketplacemeteringBatchMeterUsageOutput{value: val, isSet: true}
}

func (v NullableMarketplacemeteringBatchMeterUsageOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplacemeteringBatchMeterUsageOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
