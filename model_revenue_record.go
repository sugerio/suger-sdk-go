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

// checks if the RevenueRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevenueRecord{}

// RevenueRecord struct for RevenueRecord
type RevenueRecord struct {
	// The revenue amount for the revenue report
	Amount  *float32 `json:"amount,omitempty"`
	BuyerID *string  `json:"buyerID,omitempty"`
	// The revenue amount that the seller/ISV can collect.
	CollectableAmount *float32 `json:"collectableAmount,omitempty"`
	// The currency of the revenue in ISO 4217 format, such as \"USD\".
	Currency *string `json:"currency,omitempty"`
	// The date for the revenue report
	Date           *time.Time         `json:"date,omitempty"`
	DisburseAmount *float32           `json:"disburseAmount,omitempty"`
	DisburseDate   *time.Time         `json:"disburseDate,omitempty"`
	EntitlementID  *string            `json:"entitlementID,omitempty"`
	Id             *string            `json:"id,omitempty"`
	Info           *RevenueRecordInfo `json:"info,omitempty"`
	InvoiceAmount  *float32           `json:"invoiceAmount,omitempty"`
	InvoiceDate    *time.Time         `json:"invoiceDate,omitempty"`
	OrganizationID *string            `json:"organizationID,omitempty"`
	// The value is from type Partner string
	Partner              *string    `json:"partner,omitempty"`
	PaymentDueDate       *time.Time `json:"paymentDueDate,omitempty"`
	ProductID            *string    `json:"productID,omitempty"`
	RefundDisburseAmount *float32   `json:"refundDisburseAmount,omitempty"`
	RefundDisburseDate   *time.Time `json:"refundDisburseDate,omitempty"`
	RefundInvoiceAmount  *float32   `json:"refundInvoiceAmount,omitempty"`
	RefundInvoiceDate    *time.Time `json:"refundInvoiceDate,omitempty"`
	TaxAmount            *float32   `json:"taxAmount,omitempty"`
}

// NewRevenueRecord instantiates a new RevenueRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevenueRecord() *RevenueRecord {
	this := RevenueRecord{}
	return &this
}

// NewRevenueRecordWithDefaults instantiates a new RevenueRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevenueRecordWithDefaults() *RevenueRecord {
	this := RevenueRecord{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RevenueRecord) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *RevenueRecord) SetAmount(v float32) {
	o.Amount = &v
}

// GetBuyerID returns the BuyerID field value if set, zero value otherwise.
func (o *RevenueRecord) GetBuyerID() string {
	if o == nil || IsNil(o.BuyerID) {
		var ret string
		return ret
	}
	return *o.BuyerID
}

// GetBuyerIDOk returns a tuple with the BuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerID) {
		return nil, false
	}
	return o.BuyerID, true
}

// HasBuyerID returns a boolean if a field has been set.
func (o *RevenueRecord) HasBuyerID() bool {
	if o != nil && !IsNil(o.BuyerID) {
		return true
	}

	return false
}

// SetBuyerID gets a reference to the given string and assigns it to the BuyerID field.
func (o *RevenueRecord) SetBuyerID(v string) {
	o.BuyerID = &v
}

// GetCollectableAmount returns the CollectableAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetCollectableAmount() float32 {
	if o == nil || IsNil(o.CollectableAmount) {
		var ret float32
		return ret
	}
	return *o.CollectableAmount
}

// GetCollectableAmountOk returns a tuple with the CollectableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetCollectableAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CollectableAmount) {
		return nil, false
	}
	return o.CollectableAmount, true
}

// HasCollectableAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasCollectableAmount() bool {
	if o != nil && !IsNil(o.CollectableAmount) {
		return true
	}

	return false
}

// SetCollectableAmount gets a reference to the given float32 and assigns it to the CollectableAmount field.
func (o *RevenueRecord) SetCollectableAmount(v float32) {
	o.CollectableAmount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *RevenueRecord) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *RevenueRecord) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *RevenueRecord) SetCurrency(v string) {
	o.Currency = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RevenueRecord) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *RevenueRecord) SetDate(v time.Time) {
	o.Date = &v
}

// GetDisburseAmount returns the DisburseAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetDisburseAmount() float32 {
	if o == nil || IsNil(o.DisburseAmount) {
		var ret float32
		return ret
	}
	return *o.DisburseAmount
}

// GetDisburseAmountOk returns a tuple with the DisburseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetDisburseAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DisburseAmount) {
		return nil, false
	}
	return o.DisburseAmount, true
}

// HasDisburseAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasDisburseAmount() bool {
	if o != nil && !IsNil(o.DisburseAmount) {
		return true
	}

	return false
}

// SetDisburseAmount gets a reference to the given float32 and assigns it to the DisburseAmount field.
func (o *RevenueRecord) SetDisburseAmount(v float32) {
	o.DisburseAmount = &v
}

// GetDisburseDate returns the DisburseDate field value if set, zero value otherwise.
func (o *RevenueRecord) GetDisburseDate() time.Time {
	if o == nil || IsNil(o.DisburseDate) {
		var ret time.Time
		return ret
	}
	return *o.DisburseDate
}

// GetDisburseDateOk returns a tuple with the DisburseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetDisburseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DisburseDate) {
		return nil, false
	}
	return o.DisburseDate, true
}

// HasDisburseDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasDisburseDate() bool {
	if o != nil && !IsNil(o.DisburseDate) {
		return true
	}

	return false
}

// SetDisburseDate gets a reference to the given time.Time and assigns it to the DisburseDate field.
func (o *RevenueRecord) SetDisburseDate(v time.Time) {
	o.DisburseDate = &v
}

// GetEntitlementID returns the EntitlementID field value if set, zero value otherwise.
func (o *RevenueRecord) GetEntitlementID() string {
	if o == nil || IsNil(o.EntitlementID) {
		var ret string
		return ret
	}
	return *o.EntitlementID
}

// GetEntitlementIDOk returns a tuple with the EntitlementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetEntitlementIDOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementID) {
		return nil, false
	}
	return o.EntitlementID, true
}

// HasEntitlementID returns a boolean if a field has been set.
func (o *RevenueRecord) HasEntitlementID() bool {
	if o != nil && !IsNil(o.EntitlementID) {
		return true
	}

	return false
}

// SetEntitlementID gets a reference to the given string and assigns it to the EntitlementID field.
func (o *RevenueRecord) SetEntitlementID(v string) {
	o.EntitlementID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RevenueRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RevenueRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RevenueRecord) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *RevenueRecord) GetInfo() RevenueRecordInfo {
	if o == nil || IsNil(o.Info) {
		var ret RevenueRecordInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetInfoOk() (*RevenueRecordInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *RevenueRecord) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given RevenueRecordInfo and assigns it to the Info field.
func (o *RevenueRecord) SetInfo(v RevenueRecordInfo) {
	o.Info = &v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetInvoiceAmount() float32 {
	if o == nil || IsNil(o.InvoiceAmount) {
		var ret float32
		return ret
	}
	return *o.InvoiceAmount
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetInvoiceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.InvoiceAmount) {
		return nil, false
	}
	return o.InvoiceAmount, true
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasInvoiceAmount() bool {
	if o != nil && !IsNil(o.InvoiceAmount) {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given float32 and assigns it to the InvoiceAmount field.
func (o *RevenueRecord) SetInvoiceAmount(v float32) {
	o.InvoiceAmount = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *RevenueRecord) GetInvoiceDate() time.Time {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret time.Time
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetInvoiceDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given time.Time and assigns it to the InvoiceDate field.
func (o *RevenueRecord) SetInvoiceDate(v time.Time) {
	o.InvoiceDate = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *RevenueRecord) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *RevenueRecord) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *RevenueRecord) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *RevenueRecord) GetPartner() string {
	if o == nil || IsNil(o.Partner) {
		var ret string
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetPartnerOk() (*string, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *RevenueRecord) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given string and assigns it to the Partner field.
func (o *RevenueRecord) SetPartner(v string) {
	o.Partner = &v
}

// GetPaymentDueDate returns the PaymentDueDate field value if set, zero value otherwise.
func (o *RevenueRecord) GetPaymentDueDate() time.Time {
	if o == nil || IsNil(o.PaymentDueDate) {
		var ret time.Time
		return ret
	}
	return *o.PaymentDueDate
}

// GetPaymentDueDateOk returns a tuple with the PaymentDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetPaymentDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PaymentDueDate) {
		return nil, false
	}
	return o.PaymentDueDate, true
}

// HasPaymentDueDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasPaymentDueDate() bool {
	if o != nil && !IsNil(o.PaymentDueDate) {
		return true
	}

	return false
}

// SetPaymentDueDate gets a reference to the given time.Time and assigns it to the PaymentDueDate field.
func (o *RevenueRecord) SetPaymentDueDate(v time.Time) {
	o.PaymentDueDate = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *RevenueRecord) GetProductID() string {
	if o == nil || IsNil(o.ProductID) {
		var ret string
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetProductIDOk() (*string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *RevenueRecord) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given string and assigns it to the ProductID field.
func (o *RevenueRecord) SetProductID(v string) {
	o.ProductID = &v
}

// GetRefundDisburseAmount returns the RefundDisburseAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetRefundDisburseAmount() float32 {
	if o == nil || IsNil(o.RefundDisburseAmount) {
		var ret float32
		return ret
	}
	return *o.RefundDisburseAmount
}

// GetRefundDisburseAmountOk returns a tuple with the RefundDisburseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetRefundDisburseAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundDisburseAmount) {
		return nil, false
	}
	return o.RefundDisburseAmount, true
}

// HasRefundDisburseAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasRefundDisburseAmount() bool {
	if o != nil && !IsNil(o.RefundDisburseAmount) {
		return true
	}

	return false
}

// SetRefundDisburseAmount gets a reference to the given float32 and assigns it to the RefundDisburseAmount field.
func (o *RevenueRecord) SetRefundDisburseAmount(v float32) {
	o.RefundDisburseAmount = &v
}

// GetRefundDisburseDate returns the RefundDisburseDate field value if set, zero value otherwise.
func (o *RevenueRecord) GetRefundDisburseDate() time.Time {
	if o == nil || IsNil(o.RefundDisburseDate) {
		var ret time.Time
		return ret
	}
	return *o.RefundDisburseDate
}

// GetRefundDisburseDateOk returns a tuple with the RefundDisburseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetRefundDisburseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RefundDisburseDate) {
		return nil, false
	}
	return o.RefundDisburseDate, true
}

// HasRefundDisburseDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasRefundDisburseDate() bool {
	if o != nil && !IsNil(o.RefundDisburseDate) {
		return true
	}

	return false
}

// SetRefundDisburseDate gets a reference to the given time.Time and assigns it to the RefundDisburseDate field.
func (o *RevenueRecord) SetRefundDisburseDate(v time.Time) {
	o.RefundDisburseDate = &v
}

// GetRefundInvoiceAmount returns the RefundInvoiceAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetRefundInvoiceAmount() float32 {
	if o == nil || IsNil(o.RefundInvoiceAmount) {
		var ret float32
		return ret
	}
	return *o.RefundInvoiceAmount
}

// GetRefundInvoiceAmountOk returns a tuple with the RefundInvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetRefundInvoiceAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RefundInvoiceAmount) {
		return nil, false
	}
	return o.RefundInvoiceAmount, true
}

// HasRefundInvoiceAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasRefundInvoiceAmount() bool {
	if o != nil && !IsNil(o.RefundInvoiceAmount) {
		return true
	}

	return false
}

// SetRefundInvoiceAmount gets a reference to the given float32 and assigns it to the RefundInvoiceAmount field.
func (o *RevenueRecord) SetRefundInvoiceAmount(v float32) {
	o.RefundInvoiceAmount = &v
}

// GetRefundInvoiceDate returns the RefundInvoiceDate field value if set, zero value otherwise.
func (o *RevenueRecord) GetRefundInvoiceDate() time.Time {
	if o == nil || IsNil(o.RefundInvoiceDate) {
		var ret time.Time
		return ret
	}
	return *o.RefundInvoiceDate
}

// GetRefundInvoiceDateOk returns a tuple with the RefundInvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetRefundInvoiceDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RefundInvoiceDate) {
		return nil, false
	}
	return o.RefundInvoiceDate, true
}

// HasRefundInvoiceDate returns a boolean if a field has been set.
func (o *RevenueRecord) HasRefundInvoiceDate() bool {
	if o != nil && !IsNil(o.RefundInvoiceDate) {
		return true
	}

	return false
}

// SetRefundInvoiceDate gets a reference to the given time.Time and assigns it to the RefundInvoiceDate field.
func (o *RevenueRecord) SetRefundInvoiceDate(v time.Time) {
	o.RefundInvoiceDate = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *RevenueRecord) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevenueRecord) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *RevenueRecord) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *RevenueRecord) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

func (o RevenueRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevenueRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BuyerID) {
		toSerialize["buyerID"] = o.BuyerID
	}
	if !IsNil(o.CollectableAmount) {
		toSerialize["collectableAmount"] = o.CollectableAmount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.DisburseAmount) {
		toSerialize["disburseAmount"] = o.DisburseAmount
	}
	if !IsNil(o.DisburseDate) {
		toSerialize["disburseDate"] = o.DisburseDate
	}
	if !IsNil(o.EntitlementID) {
		toSerialize["entitlementID"] = o.EntitlementID
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.InvoiceAmount) {
		toSerialize["invoiceAmount"] = o.InvoiceAmount
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.PaymentDueDate) {
		toSerialize["paymentDueDate"] = o.PaymentDueDate
	}
	if !IsNil(o.ProductID) {
		toSerialize["productID"] = o.ProductID
	}
	if !IsNil(o.RefundDisburseAmount) {
		toSerialize["refundDisburseAmount"] = o.RefundDisburseAmount
	}
	if !IsNil(o.RefundDisburseDate) {
		toSerialize["refundDisburseDate"] = o.RefundDisburseDate
	}
	if !IsNil(o.RefundInvoiceAmount) {
		toSerialize["refundInvoiceAmount"] = o.RefundInvoiceAmount
	}
	if !IsNil(o.RefundInvoiceDate) {
		toSerialize["refundInvoiceDate"] = o.RefundInvoiceDate
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	return toSerialize, nil
}

type NullableRevenueRecord struct {
	value *RevenueRecord
	isSet bool
}

func (v NullableRevenueRecord) Get() *RevenueRecord {
	return v.value
}

func (v *NullableRevenueRecord) Set(val *RevenueRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableRevenueRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRevenueRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevenueRecord(val *RevenueRecord) *NullableRevenueRecord {
	return &NullableRevenueRecord{value: val, isSet: true}
}

func (v NullableRevenueRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevenueRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
