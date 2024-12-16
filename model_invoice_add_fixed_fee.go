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

// checks if the InvoiceAddFixedFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceAddFixedFee{}

// InvoiceAddFixedFee struct for InvoiceAddFixedFee
type InvoiceAddFixedFee struct {
	EndDate   *time.Time `json:"endDate,omitempty"`
	Quantity  *int32     `json:"quantity,omitempty"`
	Rate      *float32   `json:"rate,omitempty"`
	Reason    *string    `json:"reason,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
}

// NewInvoiceAddFixedFee instantiates a new InvoiceAddFixedFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceAddFixedFee() *InvoiceAddFixedFee {
	this := InvoiceAddFixedFee{}
	return &this
}

// NewInvoiceAddFixedFeeWithDefaults instantiates a new InvoiceAddFixedFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceAddFixedFeeWithDefaults() *InvoiceAddFixedFee {
	this := InvoiceAddFixedFee{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InvoiceAddFixedFee) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAddFixedFee) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InvoiceAddFixedFee) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *InvoiceAddFixedFee) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *InvoiceAddFixedFee) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAddFixedFee) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *InvoiceAddFixedFee) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *InvoiceAddFixedFee) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *InvoiceAddFixedFee) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAddFixedFee) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InvoiceAddFixedFee) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *InvoiceAddFixedFee) SetRate(v float32) {
	o.Rate = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvoiceAddFixedFee) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAddFixedFee) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvoiceAddFixedFee) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvoiceAddFixedFee) SetReason(v string) {
	o.Reason = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InvoiceAddFixedFee) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAddFixedFee) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InvoiceAddFixedFee) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *InvoiceAddFixedFee) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o InvoiceAddFixedFee) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceAddFixedFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableInvoiceAddFixedFee struct {
	value *InvoiceAddFixedFee
	isSet bool
}

func (v NullableInvoiceAddFixedFee) Get() *InvoiceAddFixedFee {
	return v.value
}

func (v *NullableInvoiceAddFixedFee) Set(val *InvoiceAddFixedFee) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceAddFixedFee) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceAddFixedFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceAddFixedFee(val *InvoiceAddFixedFee) *NullableInvoiceAddFixedFee {
	return &NullableInvoiceAddFixedFee{value: val, isSet: true}
}

func (v NullableInvoiceAddFixedFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceAddFixedFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
