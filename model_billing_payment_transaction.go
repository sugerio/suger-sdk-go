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

// checks if the BillingPaymentTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPaymentTransaction{}

// BillingPaymentTransaction struct for BillingPaymentTransaction
type BillingPaymentTransaction struct {
	Amount         *float32                       `json:"amount,omitempty"`
	BuyerID        *string                        `json:"buyerID,omitempty"`
	CreationTime   *time.Time                     `json:"creationTime,omitempty"`
	Currency       *string                        `json:"currency,omitempty"`
	EntitlementID  *string                        `json:"entitlementID,omitempty"`
	Id             *string                        `json:"id,omitempty"`
	Info           *BillingPaymentTransactionInfo `json:"info,omitempty"`
	InvoiceID      *string                        `json:"invoiceID,omitempty"`
	LastUpdateTime *time.Time                     `json:"lastUpdateTime,omitempty"`
	OrganizationID *string                        `json:"organizationID,omitempty"`
	ParentID       *string                        `json:"parentID,omitempty"`
	Partner        *Partner                       `json:"partner,omitempty"`
	Status         *BillingPaymentStatus          `json:"status,omitempty"`
	Type           *BillingPaymentTransactionType `json:"type,omitempty"`
	WalletID       *string                        `json:"walletID,omitempty"`
	WalletType     *BillingWalletType             `json:"walletType,omitempty"`
}

// NewBillingPaymentTransaction instantiates a new BillingPaymentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPaymentTransaction() *BillingPaymentTransaction {
	this := BillingPaymentTransaction{}
	return &this
}

// NewBillingPaymentTransactionWithDefaults instantiates a new BillingPaymentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPaymentTransactionWithDefaults() *BillingPaymentTransaction {
	this := BillingPaymentTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingPaymentTransaction) SetAmount(v float32) {
	o.Amount = &v
}

// GetBuyerID returns the BuyerID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetBuyerID() string {
	if o == nil || IsNil(o.BuyerID) {
		var ret string
		return ret
	}
	return *o.BuyerID
}

// GetBuyerIDOk returns a tuple with the BuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerID) {
		return nil, false
	}
	return o.BuyerID, true
}

// HasBuyerID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasBuyerID() bool {
	if o != nil && !IsNil(o.BuyerID) {
		return true
	}

	return false
}

// SetBuyerID gets a reference to the given string and assigns it to the BuyerID field.
func (o *BillingPaymentTransaction) SetBuyerID(v string) {
	o.BuyerID = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BillingPaymentTransaction) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingPaymentTransaction) SetCurrency(v string) {
	o.Currency = &v
}

// GetEntitlementID returns the EntitlementID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetEntitlementID() string {
	if o == nil || IsNil(o.EntitlementID) {
		var ret string
		return ret
	}
	return *o.EntitlementID
}

// GetEntitlementIDOk returns a tuple with the EntitlementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetEntitlementIDOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementID) {
		return nil, false
	}
	return o.EntitlementID, true
}

// HasEntitlementID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasEntitlementID() bool {
	if o != nil && !IsNil(o.EntitlementID) {
		return true
	}

	return false
}

// SetEntitlementID gets a reference to the given string and assigns it to the EntitlementID field.
func (o *BillingPaymentTransaction) SetEntitlementID(v string) {
	o.EntitlementID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingPaymentTransaction) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetInfo() BillingPaymentTransactionInfo {
	if o == nil || IsNil(o.Info) {
		var ret BillingPaymentTransactionInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetInfoOk() (*BillingPaymentTransactionInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given BillingPaymentTransactionInfo and assigns it to the Info field.
func (o *BillingPaymentTransaction) SetInfo(v BillingPaymentTransactionInfo) {
	o.Info = &v
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetInvoiceID() string {
	if o == nil || IsNil(o.InvoiceID) {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetInvoiceIDOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceID) {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasInvoiceID() bool {
	if o != nil && !IsNil(o.InvoiceID) {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *BillingPaymentTransaction) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetLastUpdateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetLastUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given time.Time and assigns it to the LastUpdateTime field.
func (o *BillingPaymentTransaction) SetLastUpdateTime(v time.Time) {
	o.LastUpdateTime = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *BillingPaymentTransaction) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetParentID returns the ParentID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetParentID() string {
	if o == nil || IsNil(o.ParentID) {
		var ret string
		return ret
	}
	return *o.ParentID
}

// GetParentIDOk returns a tuple with the ParentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetParentIDOk() (*string, bool) {
	if o == nil || IsNil(o.ParentID) {
		return nil, false
	}
	return o.ParentID, true
}

// HasParentID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasParentID() bool {
	if o != nil && !IsNil(o.ParentID) {
		return true
	}

	return false
}

// SetParentID gets a reference to the given string and assigns it to the ParentID field.
func (o *BillingPaymentTransaction) SetParentID(v string) {
	o.ParentID = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *BillingPaymentTransaction) SetPartner(v Partner) {
	o.Partner = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetStatus() BillingPaymentStatus {
	if o == nil || IsNil(o.Status) {
		var ret BillingPaymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetStatusOk() (*BillingPaymentStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BillingPaymentStatus and assigns it to the Status field.
func (o *BillingPaymentTransaction) SetStatus(v BillingPaymentStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetType() BillingPaymentTransactionType {
	if o == nil || IsNil(o.Type) {
		var ret BillingPaymentTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetTypeOk() (*BillingPaymentTransactionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BillingPaymentTransactionType and assigns it to the Type field.
func (o *BillingPaymentTransaction) SetType(v BillingPaymentTransactionType) {
	o.Type = &v
}

// GetWalletID returns the WalletID field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetWalletID() string {
	if o == nil || IsNil(o.WalletID) {
		var ret string
		return ret
	}
	return *o.WalletID
}

// GetWalletIDOk returns a tuple with the WalletID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetWalletIDOk() (*string, bool) {
	if o == nil || IsNil(o.WalletID) {
		return nil, false
	}
	return o.WalletID, true
}

// HasWalletID returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasWalletID() bool {
	if o != nil && !IsNil(o.WalletID) {
		return true
	}

	return false
}

// SetWalletID gets a reference to the given string and assigns it to the WalletID field.
func (o *BillingPaymentTransaction) SetWalletID(v string) {
	o.WalletID = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *BillingPaymentTransaction) GetWalletType() BillingWalletType {
	if o == nil || IsNil(o.WalletType) {
		var ret BillingWalletType
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentTransaction) GetWalletTypeOk() (*BillingWalletType, bool) {
	if o == nil || IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *BillingPaymentTransaction) HasWalletType() bool {
	if o != nil && !IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given BillingWalletType and assigns it to the WalletType field.
func (o *BillingPaymentTransaction) SetWalletType(v BillingWalletType) {
	o.WalletType = &v
}

func (o BillingPaymentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPaymentTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BuyerID) {
		toSerialize["buyerID"] = o.BuyerID
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
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
	if !IsNil(o.InvoiceID) {
		toSerialize["invoiceID"] = o.InvoiceID
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.ParentID) {
		toSerialize["parentID"] = o.ParentID
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WalletID) {
		toSerialize["walletID"] = o.WalletID
	}
	if !IsNil(o.WalletType) {
		toSerialize["walletType"] = o.WalletType
	}
	return toSerialize, nil
}

type NullableBillingPaymentTransaction struct {
	value *BillingPaymentTransaction
	isSet bool
}

func (v NullableBillingPaymentTransaction) Get() *BillingPaymentTransaction {
	return v.value
}

func (v *NullableBillingPaymentTransaction) Set(val *BillingPaymentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentTransaction(val *BillingPaymentTransaction) *NullableBillingPaymentTransaction {
	return &NullableBillingPaymentTransaction{value: val, isSet: true}
}

func (v NullableBillingPaymentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
