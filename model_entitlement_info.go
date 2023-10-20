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

// checks if the EntitlementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementInfo{}

// EntitlementInfo struct for EntitlementInfo
type EntitlementInfo struct {
	// Nullable. Alibaba Entitlements from Alibaba Marketplace.
	AlibabaEntitlements []ClientDescribeInstanceResponseBody `json:"alibabaEntitlements,omitempty"`
	// Nullable. Alibaba Orders from Alibaba Marketplace.
	AlibabaOrders []ClientDescribeOrderResponseBody `json:"alibabaOrders,omitempty"`
	// Is this Entitlement Auto Renew enabled.
	AutoRenew *bool `json:"autoRenew,omitempty"`
	// Nullable. AWS Entitlements from AWS Marketplace.
	AwsEntitlements []TypesEntitlement `json:"awsEntitlements,omitempty"`
	// Nullable. Azure Subscriptions from Azure Marketplace.
	AzureSubscriptions []AzureMarketplaceSubscription `json:"azureSubscriptions,omitempty"`
	// The amount that the seller can collect. It excludes the marketplace commision fee.
	CollectableAmount *float32 `json:"collectableAmount,omitempty"`
	// The amount that the buyer has committed to pay. It can be the sum of payment installments if applicable.
	CommitAmount *float32 `json:"commitAmount,omitempty"`
	// The dimensions for commit.
	Commits []CommitDimension `json:"commits,omitempty"`
	// The default Currency is USD.
	Currency *string `json:"currency,omitempty"`
	// The dimensions for usage-based metering.
	Dimensions []MeteringDimension `json:"dimensions,omitempty"`
	// The amount that has been disbursed to the seller account.
	DisbursedAmount *float32 `json:"disbursedAmount,omitempty"`
	EulaType *EulaType `json:"eulaType,omitempty"`
	EulaUrl *string `json:"eulaUrl,omitempty"`
	// Nullable. GCP Entitlements from GCP Marketplace.
	GcpEntitlements []GcpMarketplaceEntitlement `json:"gcpEntitlements,omitempty"`
	// Only applicable for GCP Marketplace Entitlements.
	GcpPlans []GcpMarketplaceProductPurchaseOptionSpec `json:"gcpPlans,omitempty"`
	// The amount that the buyer has got invoiced.
	InvoicedAmount *float32 `json:"invoicedAmount,omitempty"`
	// For flexible payment schedules
	PaymentInstallments []PaymentInstallment `json:"paymentInstallments,omitempty"`
	RefundCancelationPolicy *string `json:"refundCancelationPolicy,omitempty"`
	SellerNotes *string `json:"sellerNotes,omitempty"`
}

// NewEntitlementInfo instantiates a new EntitlementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementInfo() *EntitlementInfo {
	this := EntitlementInfo{}
	return &this
}

// NewEntitlementInfoWithDefaults instantiates a new EntitlementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementInfoWithDefaults() *EntitlementInfo {
	this := EntitlementInfo{}
	return &this
}

// GetAlibabaEntitlements returns the AlibabaEntitlements field value if set, zero value otherwise.
func (o *EntitlementInfo) GetAlibabaEntitlements() []ClientDescribeInstanceResponseBody {
	if o == nil || IsNil(o.AlibabaEntitlements) {
		var ret []ClientDescribeInstanceResponseBody
		return ret
	}
	return o.AlibabaEntitlements
}

// GetAlibabaEntitlementsOk returns a tuple with the AlibabaEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetAlibabaEntitlementsOk() ([]ClientDescribeInstanceResponseBody, bool) {
	if o == nil || IsNil(o.AlibabaEntitlements) {
		return nil, false
	}
	return o.AlibabaEntitlements, true
}

// HasAlibabaEntitlements returns a boolean if a field has been set.
func (o *EntitlementInfo) HasAlibabaEntitlements() bool {
	if o != nil && !IsNil(o.AlibabaEntitlements) {
		return true
	}

	return false
}

// SetAlibabaEntitlements gets a reference to the given []ClientDescribeInstanceResponseBody and assigns it to the AlibabaEntitlements field.
func (o *EntitlementInfo) SetAlibabaEntitlements(v []ClientDescribeInstanceResponseBody) {
	o.AlibabaEntitlements = v
}

// GetAlibabaOrders returns the AlibabaOrders field value if set, zero value otherwise.
func (o *EntitlementInfo) GetAlibabaOrders() []ClientDescribeOrderResponseBody {
	if o == nil || IsNil(o.AlibabaOrders) {
		var ret []ClientDescribeOrderResponseBody
		return ret
	}
	return o.AlibabaOrders
}

// GetAlibabaOrdersOk returns a tuple with the AlibabaOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetAlibabaOrdersOk() ([]ClientDescribeOrderResponseBody, bool) {
	if o == nil || IsNil(o.AlibabaOrders) {
		return nil, false
	}
	return o.AlibabaOrders, true
}

// HasAlibabaOrders returns a boolean if a field has been set.
func (o *EntitlementInfo) HasAlibabaOrders() bool {
	if o != nil && !IsNil(o.AlibabaOrders) {
		return true
	}

	return false
}

// SetAlibabaOrders gets a reference to the given []ClientDescribeOrderResponseBody and assigns it to the AlibabaOrders field.
func (o *EntitlementInfo) SetAlibabaOrders(v []ClientDescribeOrderResponseBody) {
	o.AlibabaOrders = v
}

// GetAutoRenew returns the AutoRenew field value if set, zero value otherwise.
func (o *EntitlementInfo) GetAutoRenew() bool {
	if o == nil || IsNil(o.AutoRenew) {
		var ret bool
		return ret
	}
	return *o.AutoRenew
}

// GetAutoRenewOk returns a tuple with the AutoRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetAutoRenewOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRenew) {
		return nil, false
	}
	return o.AutoRenew, true
}

// HasAutoRenew returns a boolean if a field has been set.
func (o *EntitlementInfo) HasAutoRenew() bool {
	if o != nil && !IsNil(o.AutoRenew) {
		return true
	}

	return false
}

// SetAutoRenew gets a reference to the given bool and assigns it to the AutoRenew field.
func (o *EntitlementInfo) SetAutoRenew(v bool) {
	o.AutoRenew = &v
}

// GetAwsEntitlements returns the AwsEntitlements field value if set, zero value otherwise.
func (o *EntitlementInfo) GetAwsEntitlements() []TypesEntitlement {
	if o == nil || IsNil(o.AwsEntitlements) {
		var ret []TypesEntitlement
		return ret
	}
	return o.AwsEntitlements
}

// GetAwsEntitlementsOk returns a tuple with the AwsEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetAwsEntitlementsOk() ([]TypesEntitlement, bool) {
	if o == nil || IsNil(o.AwsEntitlements) {
		return nil, false
	}
	return o.AwsEntitlements, true
}

// HasAwsEntitlements returns a boolean if a field has been set.
func (o *EntitlementInfo) HasAwsEntitlements() bool {
	if o != nil && !IsNil(o.AwsEntitlements) {
		return true
	}

	return false
}

// SetAwsEntitlements gets a reference to the given []TypesEntitlement and assigns it to the AwsEntitlements field.
func (o *EntitlementInfo) SetAwsEntitlements(v []TypesEntitlement) {
	o.AwsEntitlements = v
}

// GetAzureSubscriptions returns the AzureSubscriptions field value if set, zero value otherwise.
func (o *EntitlementInfo) GetAzureSubscriptions() []AzureMarketplaceSubscription {
	if o == nil || IsNil(o.AzureSubscriptions) {
		var ret []AzureMarketplaceSubscription
		return ret
	}
	return o.AzureSubscriptions
}

// GetAzureSubscriptionsOk returns a tuple with the AzureSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetAzureSubscriptionsOk() ([]AzureMarketplaceSubscription, bool) {
	if o == nil || IsNil(o.AzureSubscriptions) {
		return nil, false
	}
	return o.AzureSubscriptions, true
}

// HasAzureSubscriptions returns a boolean if a field has been set.
func (o *EntitlementInfo) HasAzureSubscriptions() bool {
	if o != nil && !IsNil(o.AzureSubscriptions) {
		return true
	}

	return false
}

// SetAzureSubscriptions gets a reference to the given []AzureMarketplaceSubscription and assigns it to the AzureSubscriptions field.
func (o *EntitlementInfo) SetAzureSubscriptions(v []AzureMarketplaceSubscription) {
	o.AzureSubscriptions = v
}

// GetCollectableAmount returns the CollectableAmount field value if set, zero value otherwise.
func (o *EntitlementInfo) GetCollectableAmount() float32 {
	if o == nil || IsNil(o.CollectableAmount) {
		var ret float32
		return ret
	}
	return *o.CollectableAmount
}

// GetCollectableAmountOk returns a tuple with the CollectableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetCollectableAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CollectableAmount) {
		return nil, false
	}
	return o.CollectableAmount, true
}

// HasCollectableAmount returns a boolean if a field has been set.
func (o *EntitlementInfo) HasCollectableAmount() bool {
	if o != nil && !IsNil(o.CollectableAmount) {
		return true
	}

	return false
}

// SetCollectableAmount gets a reference to the given float32 and assigns it to the CollectableAmount field.
func (o *EntitlementInfo) SetCollectableAmount(v float32) {
	o.CollectableAmount = &v
}

// GetCommitAmount returns the CommitAmount field value if set, zero value otherwise.
func (o *EntitlementInfo) GetCommitAmount() float32 {
	if o == nil || IsNil(o.CommitAmount) {
		var ret float32
		return ret
	}
	return *o.CommitAmount
}

// GetCommitAmountOk returns a tuple with the CommitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetCommitAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CommitAmount) {
		return nil, false
	}
	return o.CommitAmount, true
}

// HasCommitAmount returns a boolean if a field has been set.
func (o *EntitlementInfo) HasCommitAmount() bool {
	if o != nil && !IsNil(o.CommitAmount) {
		return true
	}

	return false
}

// SetCommitAmount gets a reference to the given float32 and assigns it to the CommitAmount field.
func (o *EntitlementInfo) SetCommitAmount(v float32) {
	o.CommitAmount = &v
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *EntitlementInfo) GetCommits() []CommitDimension {
	if o == nil || IsNil(o.Commits) {
		var ret []CommitDimension
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetCommitsOk() ([]CommitDimension, bool) {
	if o == nil || IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *EntitlementInfo) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []CommitDimension and assigns it to the Commits field.
func (o *EntitlementInfo) SetCommits(v []CommitDimension) {
	o.Commits = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *EntitlementInfo) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *EntitlementInfo) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *EntitlementInfo) SetCurrency(v string) {
	o.Currency = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *EntitlementInfo) GetDimensions() []MeteringDimension {
	if o == nil || IsNil(o.Dimensions) {
		var ret []MeteringDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetDimensionsOk() ([]MeteringDimension, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *EntitlementInfo) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []MeteringDimension and assigns it to the Dimensions field.
func (o *EntitlementInfo) SetDimensions(v []MeteringDimension) {
	o.Dimensions = v
}

// GetDisbursedAmount returns the DisbursedAmount field value if set, zero value otherwise.
func (o *EntitlementInfo) GetDisbursedAmount() float32 {
	if o == nil || IsNil(o.DisbursedAmount) {
		var ret float32
		return ret
	}
	return *o.DisbursedAmount
}

// GetDisbursedAmountOk returns a tuple with the DisbursedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetDisbursedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DisbursedAmount) {
		return nil, false
	}
	return o.DisbursedAmount, true
}

// HasDisbursedAmount returns a boolean if a field has been set.
func (o *EntitlementInfo) HasDisbursedAmount() bool {
	if o != nil && !IsNil(o.DisbursedAmount) {
		return true
	}

	return false
}

// SetDisbursedAmount gets a reference to the given float32 and assigns it to the DisbursedAmount field.
func (o *EntitlementInfo) SetDisbursedAmount(v float32) {
	o.DisbursedAmount = &v
}

// GetEulaType returns the EulaType field value if set, zero value otherwise.
func (o *EntitlementInfo) GetEulaType() EulaType {
	if o == nil || IsNil(o.EulaType) {
		var ret EulaType
		return ret
	}
	return *o.EulaType
}

// GetEulaTypeOk returns a tuple with the EulaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetEulaTypeOk() (*EulaType, bool) {
	if o == nil || IsNil(o.EulaType) {
		return nil, false
	}
	return o.EulaType, true
}

// HasEulaType returns a boolean if a field has been set.
func (o *EntitlementInfo) HasEulaType() bool {
	if o != nil && !IsNil(o.EulaType) {
		return true
	}

	return false
}

// SetEulaType gets a reference to the given EulaType and assigns it to the EulaType field.
func (o *EntitlementInfo) SetEulaType(v EulaType) {
	o.EulaType = &v
}

// GetEulaUrl returns the EulaUrl field value if set, zero value otherwise.
func (o *EntitlementInfo) GetEulaUrl() string {
	if o == nil || IsNil(o.EulaUrl) {
		var ret string
		return ret
	}
	return *o.EulaUrl
}

// GetEulaUrlOk returns a tuple with the EulaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetEulaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EulaUrl) {
		return nil, false
	}
	return o.EulaUrl, true
}

// HasEulaUrl returns a boolean if a field has been set.
func (o *EntitlementInfo) HasEulaUrl() bool {
	if o != nil && !IsNil(o.EulaUrl) {
		return true
	}

	return false
}

// SetEulaUrl gets a reference to the given string and assigns it to the EulaUrl field.
func (o *EntitlementInfo) SetEulaUrl(v string) {
	o.EulaUrl = &v
}

// GetGcpEntitlements returns the GcpEntitlements field value if set, zero value otherwise.
func (o *EntitlementInfo) GetGcpEntitlements() []GcpMarketplaceEntitlement {
	if o == nil || IsNil(o.GcpEntitlements) {
		var ret []GcpMarketplaceEntitlement
		return ret
	}
	return o.GcpEntitlements
}

// GetGcpEntitlementsOk returns a tuple with the GcpEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetGcpEntitlementsOk() ([]GcpMarketplaceEntitlement, bool) {
	if o == nil || IsNil(o.GcpEntitlements) {
		return nil, false
	}
	return o.GcpEntitlements, true
}

// HasGcpEntitlements returns a boolean if a field has been set.
func (o *EntitlementInfo) HasGcpEntitlements() bool {
	if o != nil && !IsNil(o.GcpEntitlements) {
		return true
	}

	return false
}

// SetGcpEntitlements gets a reference to the given []GcpMarketplaceEntitlement and assigns it to the GcpEntitlements field.
func (o *EntitlementInfo) SetGcpEntitlements(v []GcpMarketplaceEntitlement) {
	o.GcpEntitlements = v
}

// GetGcpPlans returns the GcpPlans field value if set, zero value otherwise.
func (o *EntitlementInfo) GetGcpPlans() []GcpMarketplaceProductPurchaseOptionSpec {
	if o == nil || IsNil(o.GcpPlans) {
		var ret []GcpMarketplaceProductPurchaseOptionSpec
		return ret
	}
	return o.GcpPlans
}

// GetGcpPlansOk returns a tuple with the GcpPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetGcpPlansOk() ([]GcpMarketplaceProductPurchaseOptionSpec, bool) {
	if o == nil || IsNil(o.GcpPlans) {
		return nil, false
	}
	return o.GcpPlans, true
}

// HasGcpPlans returns a boolean if a field has been set.
func (o *EntitlementInfo) HasGcpPlans() bool {
	if o != nil && !IsNil(o.GcpPlans) {
		return true
	}

	return false
}

// SetGcpPlans gets a reference to the given []GcpMarketplaceProductPurchaseOptionSpec and assigns it to the GcpPlans field.
func (o *EntitlementInfo) SetGcpPlans(v []GcpMarketplaceProductPurchaseOptionSpec) {
	o.GcpPlans = v
}

// GetInvoicedAmount returns the InvoicedAmount field value if set, zero value otherwise.
func (o *EntitlementInfo) GetInvoicedAmount() float32 {
	if o == nil || IsNil(o.InvoicedAmount) {
		var ret float32
		return ret
	}
	return *o.InvoicedAmount
}

// GetInvoicedAmountOk returns a tuple with the InvoicedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetInvoicedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.InvoicedAmount) {
		return nil, false
	}
	return o.InvoicedAmount, true
}

// HasInvoicedAmount returns a boolean if a field has been set.
func (o *EntitlementInfo) HasInvoicedAmount() bool {
	if o != nil && !IsNil(o.InvoicedAmount) {
		return true
	}

	return false
}

// SetInvoicedAmount gets a reference to the given float32 and assigns it to the InvoicedAmount field.
func (o *EntitlementInfo) SetInvoicedAmount(v float32) {
	o.InvoicedAmount = &v
}

// GetPaymentInstallments returns the PaymentInstallments field value if set, zero value otherwise.
func (o *EntitlementInfo) GetPaymentInstallments() []PaymentInstallment {
	if o == nil || IsNil(o.PaymentInstallments) {
		var ret []PaymentInstallment
		return ret
	}
	return o.PaymentInstallments
}

// GetPaymentInstallmentsOk returns a tuple with the PaymentInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetPaymentInstallmentsOk() ([]PaymentInstallment, bool) {
	if o == nil || IsNil(o.PaymentInstallments) {
		return nil, false
	}
	return o.PaymentInstallments, true
}

// HasPaymentInstallments returns a boolean if a field has been set.
func (o *EntitlementInfo) HasPaymentInstallments() bool {
	if o != nil && !IsNil(o.PaymentInstallments) {
		return true
	}

	return false
}

// SetPaymentInstallments gets a reference to the given []PaymentInstallment and assigns it to the PaymentInstallments field.
func (o *EntitlementInfo) SetPaymentInstallments(v []PaymentInstallment) {
	o.PaymentInstallments = v
}

// GetRefundCancelationPolicy returns the RefundCancelationPolicy field value if set, zero value otherwise.
func (o *EntitlementInfo) GetRefundCancelationPolicy() string {
	if o == nil || IsNil(o.RefundCancelationPolicy) {
		var ret string
		return ret
	}
	return *o.RefundCancelationPolicy
}

// GetRefundCancelationPolicyOk returns a tuple with the RefundCancelationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetRefundCancelationPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.RefundCancelationPolicy) {
		return nil, false
	}
	return o.RefundCancelationPolicy, true
}

// HasRefundCancelationPolicy returns a boolean if a field has been set.
func (o *EntitlementInfo) HasRefundCancelationPolicy() bool {
	if o != nil && !IsNil(o.RefundCancelationPolicy) {
		return true
	}

	return false
}

// SetRefundCancelationPolicy gets a reference to the given string and assigns it to the RefundCancelationPolicy field.
func (o *EntitlementInfo) SetRefundCancelationPolicy(v string) {
	o.RefundCancelationPolicy = &v
}

// GetSellerNotes returns the SellerNotes field value if set, zero value otherwise.
func (o *EntitlementInfo) GetSellerNotes() string {
	if o == nil || IsNil(o.SellerNotes) {
		var ret string
		return ret
	}
	return *o.SellerNotes
}

// GetSellerNotesOk returns a tuple with the SellerNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementInfo) GetSellerNotesOk() (*string, bool) {
	if o == nil || IsNil(o.SellerNotes) {
		return nil, false
	}
	return o.SellerNotes, true
}

// HasSellerNotes returns a boolean if a field has been set.
func (o *EntitlementInfo) HasSellerNotes() bool {
	if o != nil && !IsNil(o.SellerNotes) {
		return true
	}

	return false
}

// SetSellerNotes gets a reference to the given string and assigns it to the SellerNotes field.
func (o *EntitlementInfo) SetSellerNotes(v string) {
	o.SellerNotes = &v
}

func (o EntitlementInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlibabaEntitlements) {
		toSerialize["alibabaEntitlements"] = o.AlibabaEntitlements
	}
	if !IsNil(o.AlibabaOrders) {
		toSerialize["alibabaOrders"] = o.AlibabaOrders
	}
	if !IsNil(o.AutoRenew) {
		toSerialize["autoRenew"] = o.AutoRenew
	}
	if !IsNil(o.AwsEntitlements) {
		toSerialize["awsEntitlements"] = o.AwsEntitlements
	}
	if !IsNil(o.AzureSubscriptions) {
		toSerialize["azureSubscriptions"] = o.AzureSubscriptions
	}
	if !IsNil(o.CollectableAmount) {
		toSerialize["collectableAmount"] = o.CollectableAmount
	}
	if !IsNil(o.CommitAmount) {
		toSerialize["commitAmount"] = o.CommitAmount
	}
	if !IsNil(o.Commits) {
		toSerialize["commits"] = o.Commits
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.DisbursedAmount) {
		toSerialize["disbursedAmount"] = o.DisbursedAmount
	}
	if !IsNil(o.EulaType) {
		toSerialize["eulaType"] = o.EulaType
	}
	if !IsNil(o.EulaUrl) {
		toSerialize["eulaUrl"] = o.EulaUrl
	}
	if !IsNil(o.GcpEntitlements) {
		toSerialize["gcpEntitlements"] = o.GcpEntitlements
	}
	if !IsNil(o.GcpPlans) {
		toSerialize["gcpPlans"] = o.GcpPlans
	}
	if !IsNil(o.InvoicedAmount) {
		toSerialize["invoicedAmount"] = o.InvoicedAmount
	}
	if !IsNil(o.PaymentInstallments) {
		toSerialize["paymentInstallments"] = o.PaymentInstallments
	}
	if !IsNil(o.RefundCancelationPolicy) {
		toSerialize["refundCancelationPolicy"] = o.RefundCancelationPolicy
	}
	if !IsNil(o.SellerNotes) {
		toSerialize["sellerNotes"] = o.SellerNotes
	}
	return toSerialize, nil
}

type NullableEntitlementInfo struct {
	value *EntitlementInfo
	isSet bool
}

func (v NullableEntitlementInfo) Get() *EntitlementInfo {
	return v.value
}

func (v *NullableEntitlementInfo) Set(val *EntitlementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementInfo(val *EntitlementInfo) *NullableEntitlementInfo {
	return &NullableEntitlementInfo{value: val, isSet: true}
}

func (v NullableEntitlementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


