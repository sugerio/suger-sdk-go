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

// checks if the UsageMeteringDailyRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageMeteringDailyRecord{}

// UsageMeteringDailyRecord struct for UsageMeteringDailyRecord
type UsageMeteringDailyRecord struct {
	Amount *float32   `json:"amount,omitempty"`
	Date   *time.Time `json:"date,omitempty"`
	// The Entitlement ID of the usage metering daily record.
	EntitlementID *string `json:"entitlementID,omitempty"`
	// The group bys expression of the usage metering. When the same usage metering key may have multiple expressions of group bys to aggregate the usage.
	GroupBysExpression *string `json:"groupBysExpression,omitempty"`
	// The dimension key of the usage metering.
	Key *string `json:"key,omitempty"`
	// The partner where this usage metering daily record is from. Such as AWS, AZURE or GCP.
	Partner  *Partner `json:"partner,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
}

// NewUsageMeteringDailyRecord instantiates a new UsageMeteringDailyRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMeteringDailyRecord() *UsageMeteringDailyRecord {
	this := UsageMeteringDailyRecord{}
	return &this
}

// NewUsageMeteringDailyRecordWithDefaults instantiates a new UsageMeteringDailyRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMeteringDailyRecordWithDefaults() *UsageMeteringDailyRecord {
	this := UsageMeteringDailyRecord{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *UsageMeteringDailyRecord) SetAmount(v float32) {
	o.Amount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *UsageMeteringDailyRecord) SetDate(v time.Time) {
	o.Date = &v
}

// GetEntitlementID returns the EntitlementID field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetEntitlementID() string {
	if o == nil || IsNil(o.EntitlementID) {
		var ret string
		return ret
	}
	return *o.EntitlementID
}

// GetEntitlementIDOk returns a tuple with the EntitlementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetEntitlementIDOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementID) {
		return nil, false
	}
	return o.EntitlementID, true
}

// HasEntitlementID returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasEntitlementID() bool {
	if o != nil && !IsNil(o.EntitlementID) {
		return true
	}

	return false
}

// SetEntitlementID gets a reference to the given string and assigns it to the EntitlementID field.
func (o *UsageMeteringDailyRecord) SetEntitlementID(v string) {
	o.EntitlementID = &v
}

// GetGroupBysExpression returns the GroupBysExpression field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetGroupBysExpression() string {
	if o == nil || IsNil(o.GroupBysExpression) {
		var ret string
		return ret
	}
	return *o.GroupBysExpression
}

// GetGroupBysExpressionOk returns a tuple with the GroupBysExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetGroupBysExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBysExpression) {
		return nil, false
	}
	return o.GroupBysExpression, true
}

// HasGroupBysExpression returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasGroupBysExpression() bool {
	if o != nil && !IsNil(o.GroupBysExpression) {
		return true
	}

	return false
}

// SetGroupBysExpression gets a reference to the given string and assigns it to the GroupBysExpression field.
func (o *UsageMeteringDailyRecord) SetGroupBysExpression(v string) {
	o.GroupBysExpression = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UsageMeteringDailyRecord) SetKey(v string) {
	o.Key = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *UsageMeteringDailyRecord) SetPartner(v Partner) {
	o.Partner = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UsageMeteringDailyRecord) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDailyRecord) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UsageMeteringDailyRecord) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *UsageMeteringDailyRecord) SetQuantity(v float32) {
	o.Quantity = &v
}

func (o UsageMeteringDailyRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageMeteringDailyRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.EntitlementID) {
		toSerialize["entitlementID"] = o.EntitlementID
	}
	if !IsNil(o.GroupBysExpression) {
		toSerialize["groupBysExpression"] = o.GroupBysExpression
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableUsageMeteringDailyRecord struct {
	value *UsageMeteringDailyRecord
	isSet bool
}

func (v NullableUsageMeteringDailyRecord) Get() *UsageMeteringDailyRecord {
	return v.value
}

func (v *NullableUsageMeteringDailyRecord) Set(val *UsageMeteringDailyRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMeteringDailyRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMeteringDailyRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMeteringDailyRecord(val *UsageMeteringDailyRecord) *NullableUsageMeteringDailyRecord {
	return &NullableUsageMeteringDailyRecord{value: val, isSet: true}
}

func (v NullableUsageMeteringDailyRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMeteringDailyRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
