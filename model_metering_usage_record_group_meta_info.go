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
	"time"
)

// checks if the MeteringUsageRecordGroupMetaInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeteringUsageRecordGroupMetaInfo{}

// MeteringUsageRecordGroupMetaInfo struct for MeteringUsageRecordGroupMetaInfo
type MeteringUsageRecordGroupMetaInfo struct {
	// If it is true, the validation of the usage record group is skipped.
	SkipValidation *bool `json:"SkipValidation,omitempty"`
	// for usage metering API v2
	BillableRecords []MeteringUsageRecord `json:"billableRecords,omitempty"`
	// The lago amount (in dollars) of the customer. This field keeps the largest of the monthly amount. So it can only be updated when the invoice month increases.
	LagoAmount *float32 `json:"lagoAmount,omitempty"`
	// The lago subscription ID of the customer.
	LagoSubscriptionID *string `json:"lagoSubscriptionID,omitempty"`
	// The lago usage start time of the customer usage.
	LagoUsageStartTime *time.Time `json:"lagoUsageStartTime,omitempty"`
	// The metronome daily cost amount (in dollars) of the customer.
	MetronomeDailyCostAmount *float32 `json:"metronomeDailyCostAmount,omitempty"`
	// The metronome invoice ID of the customer.
	MetronomeInvoiceID *string `json:"metronomeInvoiceID,omitempty"`
	// The metronome monthly invoice amount (in dollars) of the customer. This field keeps the largest amount of the invoice month. So it can only be updated when the invoice month increases.
	MetronomeMonthlyInvoiceAmount *float32 `json:"metronomeMonthlyInvoiceAmount,omitempty"`
	// The metronome monthly invoice amount (in dollars) of the customer, which is adjusted by the seller. This field is populated only when the invoice amount is decreased by the seller via credit granting.
	MetronomeMonthlyInvoiceAmountAdjusted *float32 `json:"metronomeMonthlyInvoiceAmountAdjusted,omitempty"`
	// The original records reported by the customer before convertion. If no dimension mapping is applied, this field is the same as the records field.
	OriginRecords *map[string]float32 `json:"originRecords,omitempty"`
	// The source of the usage record group. Can be from Suger API or other third party services, such as Metronome.
	Source *UsageRecordGroupSource `json:"source,omitempty"`
	// The timestamp (UTC)) of when the usage records were generated. Optional, if not provided, the current report timestamp will be used.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewMeteringUsageRecordGroupMetaInfo instantiates a new MeteringUsageRecordGroupMetaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeteringUsageRecordGroupMetaInfo() *MeteringUsageRecordGroupMetaInfo {
	this := MeteringUsageRecordGroupMetaInfo{}
	return &this
}

// NewMeteringUsageRecordGroupMetaInfoWithDefaults instantiates a new MeteringUsageRecordGroupMetaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeteringUsageRecordGroupMetaInfoWithDefaults() *MeteringUsageRecordGroupMetaInfo {
	this := MeteringUsageRecordGroupMetaInfo{}
	return &this
}

// GetSkipValidation returns the SkipValidation field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetSkipValidation() bool {
	if o == nil || IsNil(o.SkipValidation) {
		var ret bool
		return ret
	}
	return *o.SkipValidation
}

// GetSkipValidationOk returns a tuple with the SkipValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetSkipValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipValidation) {
		return nil, false
	}
	return o.SkipValidation, true
}

// HasSkipValidation returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasSkipValidation() bool {
	if o != nil && !IsNil(o.SkipValidation) {
		return true
	}

	return false
}

// SetSkipValidation gets a reference to the given bool and assigns it to the SkipValidation field.
func (o *MeteringUsageRecordGroupMetaInfo) SetSkipValidation(v bool) {
	o.SkipValidation = &v
}

// GetBillableRecords returns the BillableRecords field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetBillableRecords() []MeteringUsageRecord {
	if o == nil || IsNil(o.BillableRecords) {
		var ret []MeteringUsageRecord
		return ret
	}
	return o.BillableRecords
}

// GetBillableRecordsOk returns a tuple with the BillableRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetBillableRecordsOk() ([]MeteringUsageRecord, bool) {
	if o == nil || IsNil(o.BillableRecords) {
		return nil, false
	}
	return o.BillableRecords, true
}

// HasBillableRecords returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasBillableRecords() bool {
	if o != nil && !IsNil(o.BillableRecords) {
		return true
	}

	return false
}

// SetBillableRecords gets a reference to the given []MeteringUsageRecord and assigns it to the BillableRecords field.
func (o *MeteringUsageRecordGroupMetaInfo) SetBillableRecords(v []MeteringUsageRecord) {
	o.BillableRecords = v
}

// GetLagoAmount returns the LagoAmount field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoAmount() float32 {
	if o == nil || IsNil(o.LagoAmount) {
		var ret float32
		return ret
	}
	return *o.LagoAmount
}

// GetLagoAmountOk returns a tuple with the LagoAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LagoAmount) {
		return nil, false
	}
	return o.LagoAmount, true
}

// HasLagoAmount returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasLagoAmount() bool {
	if o != nil && !IsNil(o.LagoAmount) {
		return true
	}

	return false
}

// SetLagoAmount gets a reference to the given float32 and assigns it to the LagoAmount field.
func (o *MeteringUsageRecordGroupMetaInfo) SetLagoAmount(v float32) {
	o.LagoAmount = &v
}

// GetLagoSubscriptionID returns the LagoSubscriptionID field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoSubscriptionID() string {
	if o == nil || IsNil(o.LagoSubscriptionID) {
		var ret string
		return ret
	}
	return *o.LagoSubscriptionID
}

// GetLagoSubscriptionIDOk returns a tuple with the LagoSubscriptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoSubscriptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.LagoSubscriptionID) {
		return nil, false
	}
	return o.LagoSubscriptionID, true
}

// HasLagoSubscriptionID returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasLagoSubscriptionID() bool {
	if o != nil && !IsNil(o.LagoSubscriptionID) {
		return true
	}

	return false
}

// SetLagoSubscriptionID gets a reference to the given string and assigns it to the LagoSubscriptionID field.
func (o *MeteringUsageRecordGroupMetaInfo) SetLagoSubscriptionID(v string) {
	o.LagoSubscriptionID = &v
}

// GetLagoUsageStartTime returns the LagoUsageStartTime field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoUsageStartTime() time.Time {
	if o == nil || IsNil(o.LagoUsageStartTime) {
		var ret time.Time
		return ret
	}
	return *o.LagoUsageStartTime
}

// GetLagoUsageStartTimeOk returns a tuple with the LagoUsageStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetLagoUsageStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LagoUsageStartTime) {
		return nil, false
	}
	return o.LagoUsageStartTime, true
}

// HasLagoUsageStartTime returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasLagoUsageStartTime() bool {
	if o != nil && !IsNil(o.LagoUsageStartTime) {
		return true
	}

	return false
}

// SetLagoUsageStartTime gets a reference to the given time.Time and assigns it to the LagoUsageStartTime field.
func (o *MeteringUsageRecordGroupMetaInfo) SetLagoUsageStartTime(v time.Time) {
	o.LagoUsageStartTime = &v
}

// GetMetronomeDailyCostAmount returns the MetronomeDailyCostAmount field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeDailyCostAmount() float32 {
	if o == nil || IsNil(o.MetronomeDailyCostAmount) {
		var ret float32
		return ret
	}
	return *o.MetronomeDailyCostAmount
}

// GetMetronomeDailyCostAmountOk returns a tuple with the MetronomeDailyCostAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeDailyCostAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MetronomeDailyCostAmount) {
		return nil, false
	}
	return o.MetronomeDailyCostAmount, true
}

// HasMetronomeDailyCostAmount returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeDailyCostAmount() bool {
	if o != nil && !IsNil(o.MetronomeDailyCostAmount) {
		return true
	}

	return false
}

// SetMetronomeDailyCostAmount gets a reference to the given float32 and assigns it to the MetronomeDailyCostAmount field.
func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeDailyCostAmount(v float32) {
	o.MetronomeDailyCostAmount = &v
}

// GetMetronomeInvoiceID returns the MetronomeInvoiceID field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeInvoiceID() string {
	if o == nil || IsNil(o.MetronomeInvoiceID) {
		var ret string
		return ret
	}
	return *o.MetronomeInvoiceID
}

// GetMetronomeInvoiceIDOk returns a tuple with the MetronomeInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeInvoiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.MetronomeInvoiceID) {
		return nil, false
	}
	return o.MetronomeInvoiceID, true
}

// HasMetronomeInvoiceID returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeInvoiceID() bool {
	if o != nil && !IsNil(o.MetronomeInvoiceID) {
		return true
	}

	return false
}

// SetMetronomeInvoiceID gets a reference to the given string and assigns it to the MetronomeInvoiceID field.
func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeInvoiceID(v string) {
	o.MetronomeInvoiceID = &v
}

// GetMetronomeMonthlyInvoiceAmount returns the MetronomeMonthlyInvoiceAmount field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmount() float32 {
	if o == nil || IsNil(o.MetronomeMonthlyInvoiceAmount) {
		var ret float32
		return ret
	}
	return *o.MetronomeMonthlyInvoiceAmount
}

// GetMetronomeMonthlyInvoiceAmountOk returns a tuple with the MetronomeMonthlyInvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MetronomeMonthlyInvoiceAmount) {
		return nil, false
	}
	return o.MetronomeMonthlyInvoiceAmount, true
}

// HasMetronomeMonthlyInvoiceAmount returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeMonthlyInvoiceAmount() bool {
	if o != nil && !IsNil(o.MetronomeMonthlyInvoiceAmount) {
		return true
	}

	return false
}

// SetMetronomeMonthlyInvoiceAmount gets a reference to the given float32 and assigns it to the MetronomeMonthlyInvoiceAmount field.
func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeMonthlyInvoiceAmount(v float32) {
	o.MetronomeMonthlyInvoiceAmount = &v
}

// GetMetronomeMonthlyInvoiceAmountAdjusted returns the MetronomeMonthlyInvoiceAmountAdjusted field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountAdjusted() float32 {
	if o == nil || IsNil(o.MetronomeMonthlyInvoiceAmountAdjusted) {
		var ret float32
		return ret
	}
	return *o.MetronomeMonthlyInvoiceAmountAdjusted
}

// GetMetronomeMonthlyInvoiceAmountAdjustedOk returns a tuple with the MetronomeMonthlyInvoiceAmountAdjusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountAdjustedOk() (*float32, bool) {
	if o == nil || IsNil(o.MetronomeMonthlyInvoiceAmountAdjusted) {
		return nil, false
	}
	return o.MetronomeMonthlyInvoiceAmountAdjusted, true
}

// HasMetronomeMonthlyInvoiceAmountAdjusted returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeMonthlyInvoiceAmountAdjusted() bool {
	if o != nil && !IsNil(o.MetronomeMonthlyInvoiceAmountAdjusted) {
		return true
	}

	return false
}

// SetMetronomeMonthlyInvoiceAmountAdjusted gets a reference to the given float32 and assigns it to the MetronomeMonthlyInvoiceAmountAdjusted field.
func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeMonthlyInvoiceAmountAdjusted(v float32) {
	o.MetronomeMonthlyInvoiceAmountAdjusted = &v
}

// GetOriginRecords returns the OriginRecords field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecords() map[string]float32 {
	if o == nil || IsNil(o.OriginRecords) {
		var ret map[string]float32
		return ret
	}
	return *o.OriginRecords
}

// GetOriginRecordsOk returns a tuple with the OriginRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecordsOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.OriginRecords) {
		return nil, false
	}
	return o.OriginRecords, true
}

// HasOriginRecords returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasOriginRecords() bool {
	if o != nil && !IsNil(o.OriginRecords) {
		return true
	}

	return false
}

// SetOriginRecords gets a reference to the given map[string]float32 and assigns it to the OriginRecords field.
func (o *MeteringUsageRecordGroupMetaInfo) SetOriginRecords(v map[string]float32) {
	o.OriginRecords = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetSource() UsageRecordGroupSource {
	if o == nil || IsNil(o.Source) {
		var ret UsageRecordGroupSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetSourceOk() (*UsageRecordGroupSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given UsageRecordGroupSource and assigns it to the Source field.
func (o *MeteringUsageRecordGroupMetaInfo) SetSource(v UsageRecordGroupSource) {
	o.Source = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MeteringUsageRecordGroupMetaInfo) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringUsageRecordGroupMetaInfo) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MeteringUsageRecordGroupMetaInfo) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MeteringUsageRecordGroupMetaInfo) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o MeteringUsageRecordGroupMetaInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeteringUsageRecordGroupMetaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipValidation) {
		toSerialize["SkipValidation"] = o.SkipValidation
	}
	if !IsNil(o.BillableRecords) {
		toSerialize["billableRecords"] = o.BillableRecords
	}
	if !IsNil(o.LagoAmount) {
		toSerialize["lagoAmount"] = o.LagoAmount
	}
	if !IsNil(o.LagoSubscriptionID) {
		toSerialize["lagoSubscriptionID"] = o.LagoSubscriptionID
	}
	if !IsNil(o.LagoUsageStartTime) {
		toSerialize["lagoUsageStartTime"] = o.LagoUsageStartTime
	}
	if !IsNil(o.MetronomeDailyCostAmount) {
		toSerialize["metronomeDailyCostAmount"] = o.MetronomeDailyCostAmount
	}
	if !IsNil(o.MetronomeInvoiceID) {
		toSerialize["metronomeInvoiceID"] = o.MetronomeInvoiceID
	}
	if !IsNil(o.MetronomeMonthlyInvoiceAmount) {
		toSerialize["metronomeMonthlyInvoiceAmount"] = o.MetronomeMonthlyInvoiceAmount
	}
	if !IsNil(o.MetronomeMonthlyInvoiceAmountAdjusted) {
		toSerialize["metronomeMonthlyInvoiceAmountAdjusted"] = o.MetronomeMonthlyInvoiceAmountAdjusted
	}
	if !IsNil(o.OriginRecords) {
		toSerialize["originRecords"] = o.OriginRecords
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableMeteringUsageRecordGroupMetaInfo struct {
	value *MeteringUsageRecordGroupMetaInfo
	isSet bool
}

func (v NullableMeteringUsageRecordGroupMetaInfo) Get() *MeteringUsageRecordGroupMetaInfo {
	return v.value
}

func (v *NullableMeteringUsageRecordGroupMetaInfo) Set(val *MeteringUsageRecordGroupMetaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMeteringUsageRecordGroupMetaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMeteringUsageRecordGroupMetaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeteringUsageRecordGroupMetaInfo(val *MeteringUsageRecordGroupMetaInfo) *NullableMeteringUsageRecordGroupMetaInfo {
	return &NullableMeteringUsageRecordGroupMetaInfo{value: val, isSet: true}
}

func (v NullableMeteringUsageRecordGroupMetaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeteringUsageRecordGroupMetaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
