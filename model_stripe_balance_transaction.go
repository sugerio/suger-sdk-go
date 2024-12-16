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

// checks if the StripeBalanceTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeBalanceTransaction{}

// StripeBalanceTransaction struct for StripeBalanceTransaction
type StripeBalanceTransaction struct {
	// Gross amount of this transaction (in cents (or local equivalent)). A positive value represents funds charged to another party, and a negative value represents funds sent to another party.
	Amount *int32 `json:"amount,omitempty"`
	// The date that the transaction's net funds become available in the Stripe balance.
	AvailableOn *int32 `json:"available_on,omitempty"`
	// ID of the charge which the balance transaction comes from.
	ChargeId *string `json:"chargeId,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created *int32 `json:"created,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `json:"description,omitempty"`
	// If applicable, this transaction uses an exchange rate. If money converts from currency A to currency B, then the `amount` in currency A, multipled by the `exchange_rate`, equals the `amount` in currency B. For example, if you charge a customer 10.00 EUR, the PaymentIntent's `amount` is `1000` and `currency` is `eur`. If this converts to 12.34 USD in your Stripe account, the BalanceTransaction's `amount` is `1234`, its `currency` is `usd`, and the `exchange_rate` is `1.234`.
	ExchangeRate *float32 `json:"exchange_rate,omitempty"`
	// Fees (in cents (or local equivalent)) paid for this transaction. Represented as a positive integer when assessed.
	Fee *int32 `json:"fee,omitempty"`
	// Detailed breakdown of fees (in cents (or local equivalent)) paid for this transaction.
	FeeDetails []StripeBalanceTransactionFeeDetail `json:"fee_details,omitempty"`
	// Unique identifier for the object.
	Id *string `json:"id,omitempty"`
	// Net impact to a Stripe balance (in cents (or local equivalent)). A positive value represents incrementing a Stripe balance, and a negative value decrementing a Stripe balance. You can calculate the net impact of a transaction on a balance by `amount` - `fee`
	Net *int32 `json:"net,omitempty"`
	// The transaction's net funds status in the Stripe balance, which are either `available` or `pending`.
	Status *string `json:"status,omitempty"`
	// Transaction type: `adjustment`, `advance`, `advance_funding`, `anticipation_repayment`, `application_fee`, `application_fee_refund`, `charge`, `climate_order_purchase`, `climate_order_refund`, `connect_collection_transfer`, `contribution`, `issuing_authorization_hold`, `issuing_authorization_release`, `issuing_dispute`, `issuing_transaction`, `obligation_outbound`, `obligation_reversal_inbound`, `payment`, `payment_failure_refund`, `payment_network_reserve_hold`, `payment_network_reserve_release`, `payment_refund`, `payment_reversal`, `payment_unreconciled`, `payout`, `payout_cancel`, `payout_failure`, `refund`, `refund_failure`, `reserve_transaction`, `reserved_funds`, `stripe_fee`, `stripe_fx_fee`, `tax_fee`, `topup`, `topup_reversal`, `transfer`, `transfer_cancel`, `transfer_failure`, or `transfer_refund`. Learn more about [balance transaction types and what they represent](https://stripe.com/docs/reports/balance-transaction-types). To classify transactions for accounting purposes, consider `reporting_category` instead.
	Type *string `json:"type,omitempty"`
}

// NewStripeBalanceTransaction instantiates a new StripeBalanceTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeBalanceTransaction() *StripeBalanceTransaction {
	this := StripeBalanceTransaction{}
	return &this
}

// NewStripeBalanceTransactionWithDefaults instantiates a new StripeBalanceTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeBalanceTransactionWithDefaults() *StripeBalanceTransaction {
	this := StripeBalanceTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *StripeBalanceTransaction) SetAmount(v int32) {
	o.Amount = &v
}

// GetAvailableOn returns the AvailableOn field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetAvailableOn() int32 {
	if o == nil || IsNil(o.AvailableOn) {
		var ret int32
		return ret
	}
	return *o.AvailableOn
}

// GetAvailableOnOk returns a tuple with the AvailableOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetAvailableOnOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableOn) {
		return nil, false
	}
	return o.AvailableOn, true
}

// HasAvailableOn returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasAvailableOn() bool {
	if o != nil && !IsNil(o.AvailableOn) {
		return true
	}

	return false
}

// SetAvailableOn gets a reference to the given int32 and assigns it to the AvailableOn field.
func (o *StripeBalanceTransaction) SetAvailableOn(v int32) {
	o.AvailableOn = &v
}

// GetChargeId returns the ChargeId field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetChargeId() string {
	if o == nil || IsNil(o.ChargeId) {
		var ret string
		return ret
	}
	return *o.ChargeId
}

// GetChargeIdOk returns a tuple with the ChargeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetChargeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeId) {
		return nil, false
	}
	return o.ChargeId, true
}

// HasChargeId returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasChargeId() bool {
	if o != nil && !IsNil(o.ChargeId) {
		return true
	}

	return false
}

// SetChargeId gets a reference to the given string and assigns it to the ChargeId field.
func (o *StripeBalanceTransaction) SetChargeId(v string) {
	o.ChargeId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetCreated() int32 {
	if o == nil || IsNil(o.Created) {
		var ret int32
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int32 and assigns it to the Created field.
func (o *StripeBalanceTransaction) SetCreated(v int32) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StripeBalanceTransaction) SetDescription(v string) {
	o.Description = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *StripeBalanceTransaction) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetFee() int32 {
	if o == nil || IsNil(o.Fee) {
		var ret int32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetFeeOk() (*int32, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given int32 and assigns it to the Fee field.
func (o *StripeBalanceTransaction) SetFee(v int32) {
	o.Fee = &v
}

// GetFeeDetails returns the FeeDetails field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetFeeDetails() []StripeBalanceTransactionFeeDetail {
	if o == nil || IsNil(o.FeeDetails) {
		var ret []StripeBalanceTransactionFeeDetail
		return ret
	}
	return o.FeeDetails
}

// GetFeeDetailsOk returns a tuple with the FeeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetFeeDetailsOk() ([]StripeBalanceTransactionFeeDetail, bool) {
	if o == nil || IsNil(o.FeeDetails) {
		return nil, false
	}
	return o.FeeDetails, true
}

// HasFeeDetails returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasFeeDetails() bool {
	if o != nil && !IsNil(o.FeeDetails) {
		return true
	}

	return false
}

// SetFeeDetails gets a reference to the given []StripeBalanceTransactionFeeDetail and assigns it to the FeeDetails field.
func (o *StripeBalanceTransaction) SetFeeDetails(v []StripeBalanceTransactionFeeDetail) {
	o.FeeDetails = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StripeBalanceTransaction) SetId(v string) {
	o.Id = &v
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetNet() int32 {
	if o == nil || IsNil(o.Net) {
		var ret int32
		return ret
	}
	return *o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetNetOk() (*int32, bool) {
	if o == nil || IsNil(o.Net) {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasNet() bool {
	if o != nil && !IsNil(o.Net) {
		return true
	}

	return false
}

// SetNet gets a reference to the given int32 and assigns it to the Net field.
func (o *StripeBalanceTransaction) SetNet(v int32) {
	o.Net = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StripeBalanceTransaction) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StripeBalanceTransaction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeBalanceTransaction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StripeBalanceTransaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StripeBalanceTransaction) SetType(v string) {
	o.Type = &v
}

func (o StripeBalanceTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeBalanceTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AvailableOn) {
		toSerialize["available_on"] = o.AvailableOn
	}
	if !IsNil(o.ChargeId) {
		toSerialize["chargeId"] = o.ChargeId
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchange_rate"] = o.ExchangeRate
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.FeeDetails) {
		toSerialize["fee_details"] = o.FeeDetails
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Net) {
		toSerialize["net"] = o.Net
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStripeBalanceTransaction struct {
	value *StripeBalanceTransaction
	isSet bool
}

func (v NullableStripeBalanceTransaction) Get() *StripeBalanceTransaction {
	return v.value
}

func (v *NullableStripeBalanceTransaction) Set(val *StripeBalanceTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeBalanceTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeBalanceTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeBalanceTransaction(val *StripeBalanceTransaction) *NullableStripeBalanceTransaction {
	return &NullableStripeBalanceTransaction{value: val, isSet: true}
}

func (v NullableStripeBalanceTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeBalanceTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
