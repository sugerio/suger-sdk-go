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

// checks if the UsageMeteringDailyVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageMeteringDailyVerification{}

// UsageMeteringDailyVerification struct for UsageMeteringDailyVerification
type UsageMeteringDailyVerification struct {
	// The amount of the usage metering billed by the cloud marketplace metering service.
	BilledAmount *float32 `json:"billedAmount,omitempty"`
	// The date when the usage metering is billed by the cloud marketplace metering service.
	BilledDate *time.Time `json:"billedDate,omitempty"`
	// The quantity of the usage metering billed by the cloud marketplace metering service.
	BilledQuantity *float32 `json:"billedQuantity,omitempty"`
	// Whether the amount is matched between reported & billed usage metering.
	IsAmountMatched *bool `json:"isAmountMatched,omitempty"`
	// Whether the quantity is matched between reported & billed usage metering.
	IsQuantityMatched *bool `json:"isQuantityMatched,omitempty"`
	// The dimension key of the usage metering.
	Key     *string  `json:"key,omitempty"`
	Partner *Partner `json:"partner,omitempty"`
	// The amount of the usage metering reported to the cloud marketplace.
	ReportedAmount *float32 `json:"reportedAmount,omitempty"`
	// The date when the usage metering is reported to the cloud marketplace.
	ReportedDate *time.Time `json:"reportedDate,omitempty"`
	// The quantity of the usage metering reported to the cloud marketplace.
	ReportedQuantity *float32 `json:"reportedQuantity,omitempty"`
}

// NewUsageMeteringDailyVerification instantiates a new UsageMeteringDailyVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMeteringDailyVerification() *UsageMeteringDailyVerification {
	this := UsageMeteringDailyVerification{}
	return &this
}

// NewUsageMeteringDailyVerificationWithDefaults instantiates a new UsageMeteringDailyVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMeteringDailyVerificationWithDefaults() *UsageMeteringDailyVerification {
	this := UsageMeteringDailyVerification{}
	return &this
}

// GetBilledAmount returns the BilledAmount field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetBilledAmount() float32 {
	if o == nil || IsNil(o.BilledAmount) {
		var ret float32
		return ret
	}
	return *o.BilledAmount
}

// GetBilledAmountOk returns a tuple with the BilledAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetBilledAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.BilledAmount) {
		return nil, false
	}
	return o.BilledAmount, true
}

// HasBilledAmount returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasBilledAmount() bool {
	if o != nil && !IsNil(o.BilledAmount) {
		return true
	}

	return false
}

// SetBilledAmount gets a reference to the given float32 and assigns it to the BilledAmount field.
func (o *UsageMeteringDailyVerification) SetBilledAmount(v float32) {
	o.BilledAmount = &v
}

// GetBilledDate returns the BilledDate field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetBilledDate() time.Time {
	if o == nil || IsNil(o.BilledDate) {
		var ret time.Time
		return ret
	}
	return *o.BilledDate
}

// GetBilledDateOk returns a tuple with the BilledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetBilledDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BilledDate) {
		return nil, false
	}
	return o.BilledDate, true
}

// HasBilledDate returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasBilledDate() bool {
	if o != nil && !IsNil(o.BilledDate) {
		return true
	}

	return false
}

// SetBilledDate gets a reference to the given time.Time and assigns it to the BilledDate field.
func (o *UsageMeteringDailyVerification) SetBilledDate(v time.Time) {
	o.BilledDate = &v
}

// GetBilledQuantity returns the BilledQuantity field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetBilledQuantity() float32 {
	if o == nil || IsNil(o.BilledQuantity) {
		var ret float32
		return ret
	}
	return *o.BilledQuantity
}

// GetBilledQuantityOk returns a tuple with the BilledQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetBilledQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.BilledQuantity) {
		return nil, false
	}
	return o.BilledQuantity, true
}

// HasBilledQuantity returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasBilledQuantity() bool {
	if o != nil && !IsNil(o.BilledQuantity) {
		return true
	}

	return false
}

// SetBilledQuantity gets a reference to the given float32 and assigns it to the BilledQuantity field.
func (o *UsageMeteringDailyVerification) SetBilledQuantity(v float32) {
	o.BilledQuantity = &v
}

// GetIsAmountMatched returns the IsAmountMatched field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetIsAmountMatched() bool {
	if o == nil || IsNil(o.IsAmountMatched) {
		var ret bool
		return ret
	}
	return *o.IsAmountMatched
}

// GetIsAmountMatchedOk returns a tuple with the IsAmountMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetIsAmountMatchedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountMatched) {
		return nil, false
	}
	return o.IsAmountMatched, true
}

// HasIsAmountMatched returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasIsAmountMatched() bool {
	if o != nil && !IsNil(o.IsAmountMatched) {
		return true
	}

	return false
}

// SetIsAmountMatched gets a reference to the given bool and assigns it to the IsAmountMatched field.
func (o *UsageMeteringDailyVerification) SetIsAmountMatched(v bool) {
	o.IsAmountMatched = &v
}

// GetIsQuantityMatched returns the IsQuantityMatched field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetIsQuantityMatched() bool {
	if o == nil || IsNil(o.IsQuantityMatched) {
		var ret bool
		return ret
	}
	return *o.IsQuantityMatched
}

// GetIsQuantityMatchedOk returns a tuple with the IsQuantityMatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetIsQuantityMatchedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuantityMatched) {
		return nil, false
	}
	return o.IsQuantityMatched, true
}

// HasIsQuantityMatched returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasIsQuantityMatched() bool {
	if o != nil && !IsNil(o.IsQuantityMatched) {
		return true
	}

	return false
}

// SetIsQuantityMatched gets a reference to the given bool and assigns it to the IsQuantityMatched field.
func (o *UsageMeteringDailyVerification) SetIsQuantityMatched(v bool) {
	o.IsQuantityMatched = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UsageMeteringDailyVerification) SetKey(v string) {
	o.Key = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *UsageMeteringDailyVerification) SetPartner(v Partner) {
	o.Partner = &v
}

// GetReportedAmount returns the ReportedAmount field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetReportedAmount() float32 {
	if o == nil || IsNil(o.ReportedAmount) {
		var ret float32
		return ret
	}
	return *o.ReportedAmount
}

// GetReportedAmountOk returns a tuple with the ReportedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetReportedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ReportedAmount) {
		return nil, false
	}
	return o.ReportedAmount, true
}

// HasReportedAmount returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasReportedAmount() bool {
	if o != nil && !IsNil(o.ReportedAmount) {
		return true
	}

	return false
}

// SetReportedAmount gets a reference to the given float32 and assigns it to the ReportedAmount field.
func (o *UsageMeteringDailyVerification) SetReportedAmount(v float32) {
	o.ReportedAmount = &v
}

// GetReportedDate returns the ReportedDate field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetReportedDate() time.Time {
	if o == nil || IsNil(o.ReportedDate) {
		var ret time.Time
		return ret
	}
	return *o.ReportedDate
}

// GetReportedDateOk returns a tuple with the ReportedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetReportedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReportedDate) {
		return nil, false
	}
	return o.ReportedDate, true
}

// HasReportedDate returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasReportedDate() bool {
	if o != nil && !IsNil(o.ReportedDate) {
		return true
	}

	return false
}

// SetReportedDate gets a reference to the given time.Time and assigns it to the ReportedDate field.
func (o *UsageMeteringDailyVerification) SetReportedDate(v time.Time) {
	o.ReportedDate = &v
}

// GetReportedQuantity returns the ReportedQuantity field value if set, zero value otherwise.
func (o *UsageMeteringDailyVerification) GetReportedQuantity() float32 {
	if o == nil || IsNil(o.ReportedQuantity) {
		var ret float32
		return ret
	}
	return *o.ReportedQuantity
}

// GetReportedQuantityOk returns a tuple with the ReportedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyVerification) GetReportedQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.ReportedQuantity) {
		return nil, false
	}
	return o.ReportedQuantity, true
}

// HasReportedQuantity returns a boolean if a field has been set.
func (o *UsageMeteringDailyVerification) HasReportedQuantity() bool {
	if o != nil && !IsNil(o.ReportedQuantity) {
		return true
	}

	return false
}

// SetReportedQuantity gets a reference to the given float32 and assigns it to the ReportedQuantity field.
func (o *UsageMeteringDailyVerification) SetReportedQuantity(v float32) {
	o.ReportedQuantity = &v
}

func (o UsageMeteringDailyVerification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageMeteringDailyVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BilledAmount) {
		toSerialize["billedAmount"] = o.BilledAmount
	}
	if !IsNil(o.BilledDate) {
		toSerialize["billedDate"] = o.BilledDate
	}
	if !IsNil(o.BilledQuantity) {
		toSerialize["billedQuantity"] = o.BilledQuantity
	}
	if !IsNil(o.IsAmountMatched) {
		toSerialize["isAmountMatched"] = o.IsAmountMatched
	}
	if !IsNil(o.IsQuantityMatched) {
		toSerialize["isQuantityMatched"] = o.IsQuantityMatched
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.ReportedAmount) {
		toSerialize["reportedAmount"] = o.ReportedAmount
	}
	if !IsNil(o.ReportedDate) {
		toSerialize["reportedDate"] = o.ReportedDate
	}
	if !IsNil(o.ReportedQuantity) {
		toSerialize["reportedQuantity"] = o.ReportedQuantity
	}
	return toSerialize, nil
}

type NullableUsageMeteringDailyVerification struct {
	value *UsageMeteringDailyVerification
	isSet bool
}

func (v NullableUsageMeteringDailyVerification) Get() *UsageMeteringDailyVerification {
	return v.value
}

func (v *NullableUsageMeteringDailyVerification) Set(val *UsageMeteringDailyVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMeteringDailyVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMeteringDailyVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMeteringDailyVerification(val *UsageMeteringDailyVerification) *NullableUsageMeteringDailyVerification {
	return &NullableUsageMeteringDailyVerification{value: val, isSet: true}
}

func (v NullableUsageMeteringDailyVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMeteringDailyVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
