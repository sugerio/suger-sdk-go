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

// checks if the BillingWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingWallet{}

// BillingWallet struct for BillingWallet
type BillingWallet struct {
	BuyerID      *string    `json:"buyerID,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	Currency     *string    `json:"currency,omitempty"`
	// nullable
	ExpireTime *time.Time `json:"expireTime,omitempty"`
	// The payment method id in payment provider, such as stripe payment method id.
	ExternalID     *string              `json:"externalID,omitempty"`
	Id             *string              `json:"id,omitempty"`
	Info           *BillingWalletInfo   `json:"info,omitempty"`
	LastUpdateTime *time.Time           `json:"lastUpdateTime,omitempty"`
	Name           *string              `json:"name,omitempty"`
	OrganizationID *string              `json:"organizationID,omitempty"`
	Partner        *Partner             `json:"partner,omitempty"`
	StartTime      *time.Time           `json:"startTime,omitempty"`
	Status         *BillingWalletStatus `json:"status,omitempty"`
	TotalAmount    *float32             `json:"totalAmount,omitempty"`
	Type           *BillingWalletType   `json:"type,omitempty"`
	UsedAmount     *float32             `json:"usedAmount,omitempty"`
}

// NewBillingWallet instantiates a new BillingWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingWallet() *BillingWallet {
	this := BillingWallet{}
	return &this
}

// NewBillingWalletWithDefaults instantiates a new BillingWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingWalletWithDefaults() *BillingWallet {
	this := BillingWallet{}
	return &this
}

// GetBuyerID returns the BuyerID field value if set, zero value otherwise.
func (o *BillingWallet) GetBuyerID() string {
	if o == nil || IsNil(o.BuyerID) {
		var ret string
		return ret
	}
	return *o.BuyerID
}

// GetBuyerIDOk returns a tuple with the BuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerID) {
		return nil, false
	}
	return o.BuyerID, true
}

// HasBuyerID returns a boolean if a field has been set.
func (o *BillingWallet) HasBuyerID() bool {
	if o != nil && !IsNil(o.BuyerID) {
		return true
	}

	return false
}

// SetBuyerID gets a reference to the given string and assigns it to the BuyerID field.
func (o *BillingWallet) SetBuyerID(v string) {
	o.BuyerID = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BillingWallet) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BillingWallet) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BillingWallet) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingWallet) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingWallet) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingWallet) SetCurrency(v string) {
	o.Currency = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *BillingWallet) GetExpireTime() time.Time {
	if o == nil || IsNil(o.ExpireTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetExpireTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *BillingWallet) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given time.Time and assigns it to the ExpireTime field.
func (o *BillingWallet) SetExpireTime(v time.Time) {
	o.ExpireTime = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *BillingWallet) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *BillingWallet) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *BillingWallet) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingWallet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingWallet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingWallet) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BillingWallet) GetInfo() BillingWalletInfo {
	if o == nil || IsNil(o.Info) {
		var ret BillingWalletInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetInfoOk() (*BillingWalletInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BillingWallet) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given BillingWalletInfo and assigns it to the Info field.
func (o *BillingWallet) SetInfo(v BillingWalletInfo) {
	o.Info = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *BillingWallet) GetLastUpdateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetLastUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *BillingWallet) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given time.Time and assigns it to the LastUpdateTime field.
func (o *BillingWallet) SetLastUpdateTime(v time.Time) {
	o.LastUpdateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingWallet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingWallet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingWallet) SetName(v string) {
	o.Name = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *BillingWallet) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *BillingWallet) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *BillingWallet) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *BillingWallet) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *BillingWallet) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *BillingWallet) SetPartner(v Partner) {
	o.Partner = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *BillingWallet) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *BillingWallet) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *BillingWallet) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingWallet) GetStatus() BillingWalletStatus {
	if o == nil || IsNil(o.Status) {
		var ret BillingWalletStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetStatusOk() (*BillingWalletStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingWallet) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BillingWalletStatus and assigns it to the Status field.
func (o *BillingWallet) SetStatus(v BillingWalletStatus) {
	o.Status = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *BillingWallet) GetTotalAmount() float32 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret float32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetTotalAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *BillingWallet) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given float32 and assigns it to the TotalAmount field.
func (o *BillingWallet) SetTotalAmount(v float32) {
	o.TotalAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingWallet) GetType() BillingWalletType {
	if o == nil || IsNil(o.Type) {
		var ret BillingWalletType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetTypeOk() (*BillingWalletType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingWallet) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BillingWalletType and assigns it to the Type field.
func (o *BillingWallet) SetType(v BillingWalletType) {
	o.Type = &v
}

// GetUsedAmount returns the UsedAmount field value if set, zero value otherwise.
func (o *BillingWallet) GetUsedAmount() float32 {
	if o == nil || IsNil(o.UsedAmount) {
		var ret float32
		return ret
	}
	return *o.UsedAmount
}

// GetUsedAmountOk returns a tuple with the UsedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingWallet) GetUsedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.UsedAmount) {
		return nil, false
	}
	return o.UsedAmount, true
}

// HasUsedAmount returns a boolean if a field has been set.
func (o *BillingWallet) HasUsedAmount() bool {
	if o != nil && !IsNil(o.UsedAmount) {
		return true
	}

	return false
}

// SetUsedAmount gets a reference to the given float32 and assigns it to the UsedAmount field.
func (o *BillingWallet) SetUsedAmount(v float32) {
	o.UsedAmount = &v
}

func (o BillingWallet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerID) {
		toSerialize["buyerID"] = o.BuyerID
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UsedAmount) {
		toSerialize["usedAmount"] = o.UsedAmount
	}
	return toSerialize, nil
}

type NullableBillingWallet struct {
	value *BillingWallet
	isSet bool
}

func (v NullableBillingWallet) Get() *BillingWallet {
	return v.value
}

func (v *NullableBillingWallet) Set(val *BillingWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingWallet(val *BillingWallet) *NullableBillingWallet {
	return &NullableBillingWallet{value: val, isSet: true}
}

func (v NullableBillingWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}