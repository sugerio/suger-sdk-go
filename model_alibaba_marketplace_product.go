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

// checks if the AlibabaMarketplaceProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlibabaMarketplaceProduct{}

// AlibabaMarketplaceProduct struct for AlibabaMarketplaceProduct
type AlibabaMarketplaceProduct struct {
	AuditFailMsg     *string                            `json:"AuditFailMsg,omitempty"`
	AuditStatus      *string                            `json:"AuditStatus,omitempty"`
	AuditTime        *int32                             `json:"AuditTime,omitempty"`
	Code             *string                            `json:"Code,omitempty"`
	Description      *string                            `json:"Description,omitempty"`
	FrontCategoryId  *int32                             `json:"FrontCategoryId,omitempty"`
	GmtCreated       *int32                             `json:"GmtCreated,omitempty"`
	GmtModified      *int32                             `json:"GmtModified,omitempty"`
	Name             *string                            `json:"Name,omitempty"`
	PicUrl           *string                            `json:"PicUrl,omitempty"`
	ProductExtras    *AlibabaMarketplaceProductExtras   `json:"ProductExtras,omitempty"`
	ProductSkus      *AlibabaMarketplaceProductSkus     `json:"ProductSkus,omitempty"`
	RequestId        *string                            `json:"RequestId,omitempty"`
	Score            *float32                           `json:"Score,omitempty"`
	ShopInfo         *AlibabaMarketplaceProductShopInfo `json:"ShopInfo,omitempty"`
	ShortDescription *string                            `json:"ShortDescription,omitempty"`
	Status           *string                            `json:"Status,omitempty"`
	SupplierPk       *int32                             `json:"SupplierPk,omitempty"`
	Type             *string                            `json:"Type,omitempty"`
	UseCount         *int32                             `json:"UseCount,omitempty"`
}

// NewAlibabaMarketplaceProduct instantiates a new AlibabaMarketplaceProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaMarketplaceProduct() *AlibabaMarketplaceProduct {
	this := AlibabaMarketplaceProduct{}
	return &this
}

// NewAlibabaMarketplaceProductWithDefaults instantiates a new AlibabaMarketplaceProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaMarketplaceProductWithDefaults() *AlibabaMarketplaceProduct {
	this := AlibabaMarketplaceProduct{}
	return &this
}

// GetAuditFailMsg returns the AuditFailMsg field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetAuditFailMsg() string {
	if o == nil || IsNil(o.AuditFailMsg) {
		var ret string
		return ret
	}
	return *o.AuditFailMsg
}

// GetAuditFailMsgOk returns a tuple with the AuditFailMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetAuditFailMsgOk() (*string, bool) {
	if o == nil || IsNil(o.AuditFailMsg) {
		return nil, false
	}
	return o.AuditFailMsg, true
}

// HasAuditFailMsg returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasAuditFailMsg() bool {
	if o != nil && !IsNil(o.AuditFailMsg) {
		return true
	}

	return false
}

// SetAuditFailMsg gets a reference to the given string and assigns it to the AuditFailMsg field.
func (o *AlibabaMarketplaceProduct) SetAuditFailMsg(v string) {
	o.AuditFailMsg = &v
}

// GetAuditStatus returns the AuditStatus field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetAuditStatus() string {
	if o == nil || IsNil(o.AuditStatus) {
		var ret string
		return ret
	}
	return *o.AuditStatus
}

// GetAuditStatusOk returns a tuple with the AuditStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetAuditStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuditStatus) {
		return nil, false
	}
	return o.AuditStatus, true
}

// HasAuditStatus returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasAuditStatus() bool {
	if o != nil && !IsNil(o.AuditStatus) {
		return true
	}

	return false
}

// SetAuditStatus gets a reference to the given string and assigns it to the AuditStatus field.
func (o *AlibabaMarketplaceProduct) SetAuditStatus(v string) {
	o.AuditStatus = &v
}

// GetAuditTime returns the AuditTime field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetAuditTime() int32 {
	if o == nil || IsNil(o.AuditTime) {
		var ret int32
		return ret
	}
	return *o.AuditTime
}

// GetAuditTimeOk returns a tuple with the AuditTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetAuditTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditTime) {
		return nil, false
	}
	return o.AuditTime, true
}

// HasAuditTime returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasAuditTime() bool {
	if o != nil && !IsNil(o.AuditTime) {
		return true
	}

	return false
}

// SetAuditTime gets a reference to the given int32 and assigns it to the AuditTime field.
func (o *AlibabaMarketplaceProduct) SetAuditTime(v int32) {
	o.AuditTime = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlibabaMarketplaceProduct) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlibabaMarketplaceProduct) SetDescription(v string) {
	o.Description = &v
}

// GetFrontCategoryId returns the FrontCategoryId field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetFrontCategoryId() int32 {
	if o == nil || IsNil(o.FrontCategoryId) {
		var ret int32
		return ret
	}
	return *o.FrontCategoryId
}

// GetFrontCategoryIdOk returns a tuple with the FrontCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetFrontCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FrontCategoryId) {
		return nil, false
	}
	return o.FrontCategoryId, true
}

// HasFrontCategoryId returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasFrontCategoryId() bool {
	if o != nil && !IsNil(o.FrontCategoryId) {
		return true
	}

	return false
}

// SetFrontCategoryId gets a reference to the given int32 and assigns it to the FrontCategoryId field.
func (o *AlibabaMarketplaceProduct) SetFrontCategoryId(v int32) {
	o.FrontCategoryId = &v
}

// GetGmtCreated returns the GmtCreated field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetGmtCreated() int32 {
	if o == nil || IsNil(o.GmtCreated) {
		var ret int32
		return ret
	}
	return *o.GmtCreated
}

// GetGmtCreatedOk returns a tuple with the GmtCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetGmtCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.GmtCreated) {
		return nil, false
	}
	return o.GmtCreated, true
}

// HasGmtCreated returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasGmtCreated() bool {
	if o != nil && !IsNil(o.GmtCreated) {
		return true
	}

	return false
}

// SetGmtCreated gets a reference to the given int32 and assigns it to the GmtCreated field.
func (o *AlibabaMarketplaceProduct) SetGmtCreated(v int32) {
	o.GmtCreated = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetGmtModified() int32 {
	if o == nil || IsNil(o.GmtModified) {
		var ret int32
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetGmtModifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given int32 and assigns it to the GmtModified field.
func (o *AlibabaMarketplaceProduct) SetGmtModified(v int32) {
	o.GmtModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlibabaMarketplaceProduct) SetName(v string) {
	o.Name = &v
}

// GetPicUrl returns the PicUrl field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetPicUrl() string {
	if o == nil || IsNil(o.PicUrl) {
		var ret string
		return ret
	}
	return *o.PicUrl
}

// GetPicUrlOk returns a tuple with the PicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetPicUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PicUrl) {
		return nil, false
	}
	return o.PicUrl, true
}

// HasPicUrl returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasPicUrl() bool {
	if o != nil && !IsNil(o.PicUrl) {
		return true
	}

	return false
}

// SetPicUrl gets a reference to the given string and assigns it to the PicUrl field.
func (o *AlibabaMarketplaceProduct) SetPicUrl(v string) {
	o.PicUrl = &v
}

// GetProductExtras returns the ProductExtras field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetProductExtras() AlibabaMarketplaceProductExtras {
	if o == nil || IsNil(o.ProductExtras) {
		var ret AlibabaMarketplaceProductExtras
		return ret
	}
	return *o.ProductExtras
}

// GetProductExtrasOk returns a tuple with the ProductExtras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetProductExtrasOk() (*AlibabaMarketplaceProductExtras, bool) {
	if o == nil || IsNil(o.ProductExtras) {
		return nil, false
	}
	return o.ProductExtras, true
}

// HasProductExtras returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasProductExtras() bool {
	if o != nil && !IsNil(o.ProductExtras) {
		return true
	}

	return false
}

// SetProductExtras gets a reference to the given AlibabaMarketplaceProductExtras and assigns it to the ProductExtras field.
func (o *AlibabaMarketplaceProduct) SetProductExtras(v AlibabaMarketplaceProductExtras) {
	o.ProductExtras = &v
}

// GetProductSkus returns the ProductSkus field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetProductSkus() AlibabaMarketplaceProductSkus {
	if o == nil || IsNil(o.ProductSkus) {
		var ret AlibabaMarketplaceProductSkus
		return ret
	}
	return *o.ProductSkus
}

// GetProductSkusOk returns a tuple with the ProductSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetProductSkusOk() (*AlibabaMarketplaceProductSkus, bool) {
	if o == nil || IsNil(o.ProductSkus) {
		return nil, false
	}
	return o.ProductSkus, true
}

// HasProductSkus returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasProductSkus() bool {
	if o != nil && !IsNil(o.ProductSkus) {
		return true
	}

	return false
}

// SetProductSkus gets a reference to the given AlibabaMarketplaceProductSkus and assigns it to the ProductSkus field.
func (o *AlibabaMarketplaceProduct) SetProductSkus(v AlibabaMarketplaceProductSkus) {
	o.ProductSkus = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *AlibabaMarketplaceProduct) SetRequestId(v string) {
	o.RequestId = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *AlibabaMarketplaceProduct) SetScore(v float32) {
	o.Score = &v
}

// GetShopInfo returns the ShopInfo field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetShopInfo() AlibabaMarketplaceProductShopInfo {
	if o == nil || IsNil(o.ShopInfo) {
		var ret AlibabaMarketplaceProductShopInfo
		return ret
	}
	return *o.ShopInfo
}

// GetShopInfoOk returns a tuple with the ShopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetShopInfoOk() (*AlibabaMarketplaceProductShopInfo, bool) {
	if o == nil || IsNil(o.ShopInfo) {
		return nil, false
	}
	return o.ShopInfo, true
}

// HasShopInfo returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasShopInfo() bool {
	if o != nil && !IsNil(o.ShopInfo) {
		return true
	}

	return false
}

// SetShopInfo gets a reference to the given AlibabaMarketplaceProductShopInfo and assigns it to the ShopInfo field.
func (o *AlibabaMarketplaceProduct) SetShopInfo(v AlibabaMarketplaceProductShopInfo) {
	o.ShopInfo = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *AlibabaMarketplaceProduct) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlibabaMarketplaceProduct) SetStatus(v string) {
	o.Status = &v
}

// GetSupplierPk returns the SupplierPk field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetSupplierPk() int32 {
	if o == nil || IsNil(o.SupplierPk) {
		var ret int32
		return ret
	}
	return *o.SupplierPk
}

// GetSupplierPkOk returns a tuple with the SupplierPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetSupplierPkOk() (*int32, bool) {
	if o == nil || IsNil(o.SupplierPk) {
		return nil, false
	}
	return o.SupplierPk, true
}

// HasSupplierPk returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasSupplierPk() bool {
	if o != nil && !IsNil(o.SupplierPk) {
		return true
	}

	return false
}

// SetSupplierPk gets a reference to the given int32 and assigns it to the SupplierPk field.
func (o *AlibabaMarketplaceProduct) SetSupplierPk(v int32) {
	o.SupplierPk = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlibabaMarketplaceProduct) SetType(v string) {
	o.Type = &v
}

// GetUseCount returns the UseCount field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProduct) GetUseCount() int32 {
	if o == nil || IsNil(o.UseCount) {
		var ret int32
		return ret
	}
	return *o.UseCount
}

// GetUseCountOk returns a tuple with the UseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProduct) GetUseCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UseCount) {
		return nil, false
	}
	return o.UseCount, true
}

// HasUseCount returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProduct) HasUseCount() bool {
	if o != nil && !IsNil(o.UseCount) {
		return true
	}

	return false
}

// SetUseCount gets a reference to the given int32 and assigns it to the UseCount field.
func (o *AlibabaMarketplaceProduct) SetUseCount(v int32) {
	o.UseCount = &v
}

func (o AlibabaMarketplaceProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaMarketplaceProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditFailMsg) {
		toSerialize["AuditFailMsg"] = o.AuditFailMsg
	}
	if !IsNil(o.AuditStatus) {
		toSerialize["AuditStatus"] = o.AuditStatus
	}
	if !IsNil(o.AuditTime) {
		toSerialize["AuditTime"] = o.AuditTime
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.FrontCategoryId) {
		toSerialize["FrontCategoryId"] = o.FrontCategoryId
	}
	if !IsNil(o.GmtCreated) {
		toSerialize["GmtCreated"] = o.GmtCreated
	}
	if !IsNil(o.GmtModified) {
		toSerialize["GmtModified"] = o.GmtModified
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PicUrl) {
		toSerialize["PicUrl"] = o.PicUrl
	}
	if !IsNil(o.ProductExtras) {
		toSerialize["ProductExtras"] = o.ProductExtras
	}
	if !IsNil(o.ProductSkus) {
		toSerialize["ProductSkus"] = o.ProductSkus
	}
	if !IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	if !IsNil(o.Score) {
		toSerialize["Score"] = o.Score
	}
	if !IsNil(o.ShopInfo) {
		toSerialize["ShopInfo"] = o.ShopInfo
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["ShortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.SupplierPk) {
		toSerialize["SupplierPk"] = o.SupplierPk
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.UseCount) {
		toSerialize["UseCount"] = o.UseCount
	}
	return toSerialize, nil
}

type NullableAlibabaMarketplaceProduct struct {
	value *AlibabaMarketplaceProduct
	isSet bool
}

func (v NullableAlibabaMarketplaceProduct) Get() *AlibabaMarketplaceProduct {
	return v.value
}

func (v *NullableAlibabaMarketplaceProduct) Set(val *AlibabaMarketplaceProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaMarketplaceProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaMarketplaceProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaMarketplaceProduct(val *AlibabaMarketplaceProduct) *NullableAlibabaMarketplaceProduct {
	return &NullableAlibabaMarketplaceProduct{value: val, isSet: true}
}

func (v NullableAlibabaMarketplaceProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaMarketplaceProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
