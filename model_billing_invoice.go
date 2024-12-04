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

// checks if the BillingInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInvoice{}

// BillingInvoice struct for BillingInvoice
type BillingInvoice struct {
	BuyerID       *string             `json:"buyerID,omitempty"`
	CreationTime  *time.Time          `json:"creationTime,omitempty"`
	EndDate       *time.Time          `json:"endDate,omitempty"`
	EntitlementID *string             `json:"entitlementID,omitempty"`
	Id            *string             `json:"id,omitempty"`
	Info          *BillingInvoiceInfo `json:"info,omitempty"`
	// The invoice file URL, provided as AWS S3 presigned URL with expiration time. Output only.
	InvoiceURL     *string               `json:"invoiceURL,omitempty"`
	LastUpdateTime *time.Time            `json:"lastUpdateTime,omitempty"`
	OrganizationID *string               `json:"organizationID,omitempty"`
	PaymentStatus  *BillingPaymentStatus `json:"paymentStatus,omitempty"`
	StartDate      *time.Time            `json:"startDate,omitempty"`
	Status         *BillingInvoiceStatus `json:"status,omitempty"`
	Type           *BillingInvoiceType   `json:"type,omitempty"`
}

// NewBillingInvoice instantiates a new BillingInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInvoice() *BillingInvoice {
	this := BillingInvoice{}
	return &this
}

// NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInvoiceWithDefaults() *BillingInvoice {
	this := BillingInvoice{}
	return &this
}

// GetBuyerID returns the BuyerID field value if set, zero value otherwise.
func (o *BillingInvoice) GetBuyerID() string {
	if o == nil || IsNil(o.BuyerID) {
		var ret string
		return ret
	}
	return *o.BuyerID
}

// GetBuyerIDOk returns a tuple with the BuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerID) {
		return nil, false
	}
	return o.BuyerID, true
}

// HasBuyerID returns a boolean if a field has been set.
func (o *BillingInvoice) HasBuyerID() bool {
	if o != nil && !IsNil(o.BuyerID) {
		return true
	}

	return false
}

// SetBuyerID gets a reference to the given string and assigns it to the BuyerID field.
func (o *BillingInvoice) SetBuyerID(v string) {
	o.BuyerID = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BillingInvoice) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BillingInvoice) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BillingInvoice) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingInvoice) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingInvoice) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingInvoice) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetEntitlementID returns the EntitlementID field value if set, zero value otherwise.
func (o *BillingInvoice) GetEntitlementID() string {
	if o == nil || IsNil(o.EntitlementID) {
		var ret string
		return ret
	}
	return *o.EntitlementID
}

// GetEntitlementIDOk returns a tuple with the EntitlementID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetEntitlementIDOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementID) {
		return nil, false
	}
	return o.EntitlementID, true
}

// HasEntitlementID returns a boolean if a field has been set.
func (o *BillingInvoice) HasEntitlementID() bool {
	if o != nil && !IsNil(o.EntitlementID) {
		return true
	}

	return false
}

// SetEntitlementID gets a reference to the given string and assigns it to the EntitlementID field.
func (o *BillingInvoice) SetEntitlementID(v string) {
	o.EntitlementID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingInvoice) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingInvoice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingInvoice) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BillingInvoice) GetInfo() BillingInvoiceInfo {
	if o == nil || IsNil(o.Info) {
		var ret BillingInvoiceInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetInfoOk() (*BillingInvoiceInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BillingInvoice) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given BillingInvoiceInfo and assigns it to the Info field.
func (o *BillingInvoice) SetInfo(v BillingInvoiceInfo) {
	o.Info = &v
}

// GetInvoiceURL returns the InvoiceURL field value if set, zero value otherwise.
func (o *BillingInvoice) GetInvoiceURL() string {
	if o == nil || IsNil(o.InvoiceURL) {
		var ret string
		return ret
	}
	return *o.InvoiceURL
}

// GetInvoiceURLOk returns a tuple with the InvoiceURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetInvoiceURLOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceURL) {
		return nil, false
	}
	return o.InvoiceURL, true
}

// HasInvoiceURL returns a boolean if a field has been set.
func (o *BillingInvoice) HasInvoiceURL() bool {
	if o != nil && !IsNil(o.InvoiceURL) {
		return true
	}

	return false
}

// SetInvoiceURL gets a reference to the given string and assigns it to the InvoiceURL field.
func (o *BillingInvoice) SetInvoiceURL(v string) {
	o.InvoiceURL = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *BillingInvoice) GetLastUpdateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetLastUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *BillingInvoice) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given time.Time and assigns it to the LastUpdateTime field.
func (o *BillingInvoice) SetLastUpdateTime(v time.Time) {
	o.LastUpdateTime = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *BillingInvoice) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *BillingInvoice) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *BillingInvoice) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *BillingInvoice) GetPaymentStatus() BillingPaymentStatus {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret BillingPaymentStatus
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetPaymentStatusOk() (*BillingPaymentStatus, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *BillingInvoice) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given BillingPaymentStatus and assigns it to the PaymentStatus field.
func (o *BillingInvoice) SetPaymentStatus(v BillingPaymentStatus) {
	o.PaymentStatus = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingInvoice) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingInvoice) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingInvoice) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingInvoice) GetStatus() BillingInvoiceStatus {
	if o == nil || IsNil(o.Status) {
		var ret BillingInvoiceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetStatusOk() (*BillingInvoiceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingInvoice) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BillingInvoiceStatus and assigns it to the Status field.
func (o *BillingInvoice) SetStatus(v BillingInvoiceStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingInvoice) GetType() BillingInvoiceType {
	if o == nil || IsNil(o.Type) {
		var ret BillingInvoiceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoice) GetTypeOk() (*BillingInvoiceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingInvoice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BillingInvoiceType and assigns it to the Type field.
func (o *BillingInvoice) SetType(v BillingInvoiceType) {
	o.Type = &v
}

func (o BillingInvoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerID) {
		toSerialize["buyerID"] = o.BuyerID
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
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
	if !IsNil(o.InvoiceURL) {
		toSerialize["invoiceURL"] = o.InvoiceURL
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["paymentStatus"] = o.PaymentStatus
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBillingInvoice struct {
	value *BillingInvoice
	isSet bool
}

func (v NullableBillingInvoice) Get() *BillingInvoice {
	return v.value
}

func (v *NullableBillingInvoice) Set(val *BillingInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInvoice(val *BillingInvoice) *NullableBillingInvoice {
	return &NullableBillingInvoice{value: val, isSet: true}
}

func (v NullableBillingInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}