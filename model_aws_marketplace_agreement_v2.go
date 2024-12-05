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

// checks if the AwsMarketplaceAgreementV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceAgreementV2{}

// AwsMarketplaceAgreementV2 struct for AwsMarketplaceAgreementV2
type AwsMarketplaceAgreementV2 struct {
	AcceptanceTime *time.Time `json:"acceptanceTime,omitempty"`
	// AWS Marketplace Agreement Id
	AgreementId   *string `json:"agreementId,omitempty"`
	AgreementType *string `json:"agreementType,omitempty"`
	// The AWS Account Id of the buyer in AWS Marketplace
	BuyerAccountId *string    `json:"buyerAccountId,omitempty"`
	EndTime        *time.Time `json:"endTime,omitempty"`
	// AWS Marketplace Offer Id
	OfferId *string `json:"offerId,omitempty"`
	// AWS Marketplace Product Id
	ProductId   *string `json:"productId,omitempty"`
	ProductType *string `json:"productType,omitempty"`
	// The AWS Account Id of the seller in AWS Marketplace
	SellerAccountId *string                        `json:"sellerAccountId,omitempty"`
	StartTime       *time.Time                     `json:"startTime,omitempty"`
	Status          *AwsMarketplaceAgreementStatus `json:"status,omitempty"`
}

// NewAwsMarketplaceAgreementV2 instantiates a new AwsMarketplaceAgreementV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceAgreementV2() *AwsMarketplaceAgreementV2 {
	this := AwsMarketplaceAgreementV2{}
	return &this
}

// NewAwsMarketplaceAgreementV2WithDefaults instantiates a new AwsMarketplaceAgreementV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceAgreementV2WithDefaults() *AwsMarketplaceAgreementV2 {
	this := AwsMarketplaceAgreementV2{}
	return &this
}

// GetAcceptanceTime returns the AcceptanceTime field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetAcceptanceTime() time.Time {
	if o == nil || IsNil(o.AcceptanceTime) {
		var ret time.Time
		return ret
	}
	return *o.AcceptanceTime
}

// GetAcceptanceTimeOk returns a tuple with the AcceptanceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetAcceptanceTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AcceptanceTime) {
		return nil, false
	}
	return o.AcceptanceTime, true
}

// HasAcceptanceTime returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasAcceptanceTime() bool {
	if o != nil && !IsNil(o.AcceptanceTime) {
		return true
	}

	return false
}

// SetAcceptanceTime gets a reference to the given time.Time and assigns it to the AcceptanceTime field.
func (o *AwsMarketplaceAgreementV2) SetAcceptanceTime(v time.Time) {
	o.AcceptanceTime = &v
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AwsMarketplaceAgreementV2) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAgreementType returns the AgreementType field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetAgreementType() string {
	if o == nil || IsNil(o.AgreementType) {
		var ret string
		return ret
	}
	return *o.AgreementType
}

// GetAgreementTypeOk returns a tuple with the AgreementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetAgreementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementType) {
		return nil, false
	}
	return o.AgreementType, true
}

// HasAgreementType returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasAgreementType() bool {
	if o != nil && !IsNil(o.AgreementType) {
		return true
	}

	return false
}

// SetAgreementType gets a reference to the given string and assigns it to the AgreementType field.
func (o *AwsMarketplaceAgreementV2) SetAgreementType(v string) {
	o.AgreementType = &v
}

// GetBuyerAccountId returns the BuyerAccountId field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetBuyerAccountId() string {
	if o == nil || IsNil(o.BuyerAccountId) {
		var ret string
		return ret
	}
	return *o.BuyerAccountId
}

// GetBuyerAccountIdOk returns a tuple with the BuyerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetBuyerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerAccountId) {
		return nil, false
	}
	return o.BuyerAccountId, true
}

// HasBuyerAccountId returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasBuyerAccountId() bool {
	if o != nil && !IsNil(o.BuyerAccountId) {
		return true
	}

	return false
}

// SetBuyerAccountId gets a reference to the given string and assigns it to the BuyerAccountId field.
func (o *AwsMarketplaceAgreementV2) SetBuyerAccountId(v string) {
	o.BuyerAccountId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *AwsMarketplaceAgreementV2) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetOfferId() string {
	if o == nil || IsNil(o.OfferId) {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.OfferId) {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasOfferId() bool {
	if o != nil && !IsNil(o.OfferId) {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *AwsMarketplaceAgreementV2) SetOfferId(v string) {
	o.OfferId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AwsMarketplaceAgreementV2) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AwsMarketplaceAgreementV2) SetProductType(v string) {
	o.ProductType = &v
}

// GetSellerAccountId returns the SellerAccountId field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetSellerAccountId() string {
	if o == nil || IsNil(o.SellerAccountId) {
		var ret string
		return ret
	}
	return *o.SellerAccountId
}

// GetSellerAccountIdOk returns a tuple with the SellerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetSellerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerAccountId) {
		return nil, false
	}
	return o.SellerAccountId, true
}

// HasSellerAccountId returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasSellerAccountId() bool {
	if o != nil && !IsNil(o.SellerAccountId) {
		return true
	}

	return false
}

// SetSellerAccountId gets a reference to the given string and assigns it to the SellerAccountId field.
func (o *AwsMarketplaceAgreementV2) SetSellerAccountId(v string) {
	o.SellerAccountId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AwsMarketplaceAgreementV2) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AwsMarketplaceAgreementV2) GetStatus() AwsMarketplaceAgreementStatus {
	if o == nil || IsNil(o.Status) {
		var ret AwsMarketplaceAgreementStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceAgreementV2) GetStatusOk() (*AwsMarketplaceAgreementStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AwsMarketplaceAgreementV2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AwsMarketplaceAgreementStatus and assigns it to the Status field.
func (o *AwsMarketplaceAgreementV2) SetStatus(v AwsMarketplaceAgreementStatus) {
	o.Status = &v
}

func (o AwsMarketplaceAgreementV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceAgreementV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptanceTime) {
		toSerialize["acceptanceTime"] = o.AcceptanceTime
	}
	if !IsNil(o.AgreementId) {
		toSerialize["agreementId"] = o.AgreementId
	}
	if !IsNil(o.AgreementType) {
		toSerialize["agreementType"] = o.AgreementType
	}
	if !IsNil(o.BuyerAccountId) {
		toSerialize["buyerAccountId"] = o.BuyerAccountId
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.OfferId) {
		toSerialize["offerId"] = o.OfferId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.SellerAccountId) {
		toSerialize["sellerAccountId"] = o.SellerAccountId
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceAgreementV2 struct {
	value *AwsMarketplaceAgreementV2
	isSet bool
}

func (v NullableAwsMarketplaceAgreementV2) Get() *AwsMarketplaceAgreementV2 {
	return v.value
}

func (v *NullableAwsMarketplaceAgreementV2) Set(val *AwsMarketplaceAgreementV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceAgreementV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceAgreementV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceAgreementV2(val *AwsMarketplaceAgreementV2) *NullableAwsMarketplaceAgreementV2 {
	return &NullableAwsMarketplaceAgreementV2{value: val, isSet: true}
}

func (v NullableAwsMarketplaceAgreementV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceAgreementV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
