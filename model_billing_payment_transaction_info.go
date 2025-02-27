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

// checks if the BillingPaymentTransactionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPaymentTransactionInfo{}

// BillingPaymentTransactionInfo struct for BillingPaymentTransactionInfo
type BillingPaymentTransactionInfo struct {
	// The invoice issue date.
	InvoiceDate *time.Time `json:"invoiceDate,omitempty"`
	// Refund flag marks whether the transaction has any refund records.
	RefundExists *bool `json:"refundExists,omitempty"`
	// Balance transaction that describes the impact of this charge on your account balance.
	StripeBalanceTransaction *StripeBalanceTransaction `json:"stripeBalanceTransaction,omitempty"`
	// Stripe dispute result, got by Dispute API, there may be multiple disputes.
	StripeDisputes []StripeDispute `json:"stripeDisputes,omitempty"`
	// Error of stripe API call
	StripeError *StripeError `json:"stripeError,omitempty"`
	// Stripe payment intent result, got by PaymentIntent API
	StripePaymentIntent *StripePaymentIntent `json:"stripePaymentIntent,omitempty"`
	// Stripe refund result, got by Refund API
	StripeRefund *StripeRefund `json:"stripeRefund,omitempty"`
}

// NewBillingPaymentTransactionInfo instantiates a new BillingPaymentTransactionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPaymentTransactionInfo() *BillingPaymentTransactionInfo {
	this := BillingPaymentTransactionInfo{}
	return &this
}

// NewBillingPaymentTransactionInfoWithDefaults instantiates a new BillingPaymentTransactionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPaymentTransactionInfoWithDefaults() *BillingPaymentTransactionInfo {
	this := BillingPaymentTransactionInfo{}
	return &this
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetInvoiceDate() time.Time {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret time.Time
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetInvoiceDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given time.Time and assigns it to the InvoiceDate field.
func (o *BillingPaymentTransactionInfo) SetInvoiceDate(v time.Time) {
	o.InvoiceDate = &v
}

// GetRefundExists returns the RefundExists field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetRefundExists() bool {
	if o == nil || IsNil(o.RefundExists) {
		var ret bool
		return ret
	}
	return *o.RefundExists
}

// GetRefundExistsOk returns a tuple with the RefundExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetRefundExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.RefundExists) {
		return nil, false
	}
	return o.RefundExists, true
}

// HasRefundExists returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasRefundExists() bool {
	if o != nil && !IsNil(o.RefundExists) {
		return true
	}

	return false
}

// SetRefundExists gets a reference to the given bool and assigns it to the RefundExists field.
func (o *BillingPaymentTransactionInfo) SetRefundExists(v bool) {
	o.RefundExists = &v
}

// GetStripeBalanceTransaction returns the StripeBalanceTransaction field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetStripeBalanceTransaction() StripeBalanceTransaction {
	if o == nil || IsNil(o.StripeBalanceTransaction) {
		var ret StripeBalanceTransaction
		return ret
	}
	return *o.StripeBalanceTransaction
}

// GetStripeBalanceTransactionOk returns a tuple with the StripeBalanceTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetStripeBalanceTransactionOk() (*StripeBalanceTransaction, bool) {
	if o == nil || IsNil(o.StripeBalanceTransaction) {
		return nil, false
	}
	return o.StripeBalanceTransaction, true
}

// HasStripeBalanceTransaction returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasStripeBalanceTransaction() bool {
	if o != nil && !IsNil(o.StripeBalanceTransaction) {
		return true
	}

	return false
}

// SetStripeBalanceTransaction gets a reference to the given StripeBalanceTransaction and assigns it to the StripeBalanceTransaction field.
func (o *BillingPaymentTransactionInfo) SetStripeBalanceTransaction(v StripeBalanceTransaction) {
	o.StripeBalanceTransaction = &v
}

// GetStripeDisputes returns the StripeDisputes field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetStripeDisputes() []StripeDispute {
	if o == nil || IsNil(o.StripeDisputes) {
		var ret []StripeDispute
		return ret
	}
	return o.StripeDisputes
}

// GetStripeDisputesOk returns a tuple with the StripeDisputes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetStripeDisputesOk() ([]StripeDispute, bool) {
	if o == nil || IsNil(o.StripeDisputes) {
		return nil, false
	}
	return o.StripeDisputes, true
}

// HasStripeDisputes returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasStripeDisputes() bool {
	if o != nil && !IsNil(o.StripeDisputes) {
		return true
	}

	return false
}

// SetStripeDisputes gets a reference to the given []StripeDispute and assigns it to the StripeDisputes field.
func (o *BillingPaymentTransactionInfo) SetStripeDisputes(v []StripeDispute) {
	o.StripeDisputes = v
}

// GetStripeError returns the StripeError field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetStripeError() StripeError {
	if o == nil || IsNil(o.StripeError) {
		var ret StripeError
		return ret
	}
	return *o.StripeError
}

// GetStripeErrorOk returns a tuple with the StripeError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetStripeErrorOk() (*StripeError, bool) {
	if o == nil || IsNil(o.StripeError) {
		return nil, false
	}
	return o.StripeError, true
}

// HasStripeError returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasStripeError() bool {
	if o != nil && !IsNil(o.StripeError) {
		return true
	}

	return false
}

// SetStripeError gets a reference to the given StripeError and assigns it to the StripeError field.
func (o *BillingPaymentTransactionInfo) SetStripeError(v StripeError) {
	o.StripeError = &v
}

// GetStripePaymentIntent returns the StripePaymentIntent field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetStripePaymentIntent() StripePaymentIntent {
	if o == nil || IsNil(o.StripePaymentIntent) {
		var ret StripePaymentIntent
		return ret
	}
	return *o.StripePaymentIntent
}

// GetStripePaymentIntentOk returns a tuple with the StripePaymentIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetStripePaymentIntentOk() (*StripePaymentIntent, bool) {
	if o == nil || IsNil(o.StripePaymentIntent) {
		return nil, false
	}
	return o.StripePaymentIntent, true
}

// HasStripePaymentIntent returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasStripePaymentIntent() bool {
	if o != nil && !IsNil(o.StripePaymentIntent) {
		return true
	}

	return false
}

// SetStripePaymentIntent gets a reference to the given StripePaymentIntent and assigns it to the StripePaymentIntent field.
func (o *BillingPaymentTransactionInfo) SetStripePaymentIntent(v StripePaymentIntent) {
	o.StripePaymentIntent = &v
}

// GetStripeRefund returns the StripeRefund field value if set, zero value otherwise.
func (o *BillingPaymentTransactionInfo) GetStripeRefund() StripeRefund {
	if o == nil || IsNil(o.StripeRefund) {
		var ret StripeRefund
		return ret
	}
	return *o.StripeRefund
}

// GetStripeRefundOk returns a tuple with the StripeRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransactionInfo) GetStripeRefundOk() (*StripeRefund, bool) {
	if o == nil || IsNil(o.StripeRefund) {
		return nil, false
	}
	return o.StripeRefund, true
}

// HasStripeRefund returns a boolean if a field has been set.
func (o *BillingPaymentTransactionInfo) HasStripeRefund() bool {
	if o != nil && !IsNil(o.StripeRefund) {
		return true
	}

	return false
}

// SetStripeRefund gets a reference to the given StripeRefund and assigns it to the StripeRefund field.
func (o *BillingPaymentTransactionInfo) SetStripeRefund(v StripeRefund) {
	o.StripeRefund = &v
}

func (o BillingPaymentTransactionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPaymentTransactionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	if !IsNil(o.RefundExists) {
		toSerialize["refundExists"] = o.RefundExists
	}
	if !IsNil(o.StripeBalanceTransaction) {
		toSerialize["stripeBalanceTransaction"] = o.StripeBalanceTransaction
	}
	if !IsNil(o.StripeDisputes) {
		toSerialize["stripeDisputes"] = o.StripeDisputes
	}
	if !IsNil(o.StripeError) {
		toSerialize["stripeError"] = o.StripeError
	}
	if !IsNil(o.StripePaymentIntent) {
		toSerialize["stripePaymentIntent"] = o.StripePaymentIntent
	}
	if !IsNil(o.StripeRefund) {
		toSerialize["stripeRefund"] = o.StripeRefund
	}
	return toSerialize, nil
}

type NullableBillingPaymentTransactionInfo struct {
	value *BillingPaymentTransactionInfo
	isSet bool
}

func (v NullableBillingPaymentTransactionInfo) Get() *BillingPaymentTransactionInfo {
	return v.value
}

func (v *NullableBillingPaymentTransactionInfo) Set(val *BillingPaymentTransactionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentTransactionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentTransactionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentTransactionInfo(val *BillingPaymentTransactionInfo) *NullableBillingPaymentTransactionInfo {
	return &NullableBillingPaymentTransactionInfo{value: val, isSet: true}
}

func (v NullableBillingPaymentTransactionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentTransactionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
