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

// checks if the ClientDescribeOrderResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientDescribeOrderResponseBody{}

// ClientDescribeOrderResponseBody struct for ClientDescribeOrderResponseBody
type ClientDescribeOrderResponseBody struct {
	AccountQuantity     *int32                                             `json:"AccountQuantity,omitempty"`
	AliUid              *int32                                             `json:"AliUid,omitempty"`
	Components          map[string]interface{}                             `json:"Components,omitempty"`
	CouponPrice         *float32                                           `json:"CouponPrice,omitempty"`
	CreatedOn           *int32                                             `json:"CreatedOn,omitempty"`
	InstanceIds         *ClientDescribeOrderResponseBodyInstanceIds        `json:"InstanceIds,omitempty"`
	OrderId             *int32                                             `json:"OrderId,omitempty"`
	OrderStatus         *string                                            `json:"OrderStatus,omitempty"`
	OrderType           *string                                            `json:"OrderType,omitempty"`
	OriginalPrice       *float32                                           `json:"OriginalPrice,omitempty"`
	PaidOn              *int32                                             `json:"PaidOn,omitempty"`
	PayStatus           *string                                            `json:"PayStatus,omitempty"`
	PaymentPrice        *float32                                           `json:"PaymentPrice,omitempty"`
	PeriodType          *string                                            `json:"PeriodType,omitempty"`
	ProductCode         *string                                            `json:"ProductCode,omitempty"`
	ProductName         *string                                            `json:"ProductName,omitempty"`
	ProductSkuCode      *string                                            `json:"ProductSkuCode,omitempty"`
	Quantity            *int32                                             `json:"Quantity,omitempty"`
	RequestId           *string                                            `json:"RequestId,omitempty"`
	SupplierCompanyName *string                                            `json:"SupplierCompanyName,omitempty"`
	SupplierTelephones  *ClientDescribeOrderResponseBodySupplierTelephones `json:"SupplierTelephones,omitempty"`
	TotalPrice          *float32                                           `json:"TotalPrice,omitempty"`
}

// NewClientDescribeOrderResponseBody instantiates a new ClientDescribeOrderResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientDescribeOrderResponseBody() *ClientDescribeOrderResponseBody {
	this := ClientDescribeOrderResponseBody{}
	return &this
}

// NewClientDescribeOrderResponseBodyWithDefaults instantiates a new ClientDescribeOrderResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientDescribeOrderResponseBodyWithDefaults() *ClientDescribeOrderResponseBody {
	this := ClientDescribeOrderResponseBody{}
	return &this
}

// GetAccountQuantity returns the AccountQuantity field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetAccountQuantity() int32 {
	if o == nil || IsNil(o.AccountQuantity) {
		var ret int32
		return ret
	}
	return *o.AccountQuantity
}

// GetAccountQuantityOk returns a tuple with the AccountQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetAccountQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountQuantity) {
		return nil, false
	}
	return o.AccountQuantity, true
}

// HasAccountQuantity returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasAccountQuantity() bool {
	if o != nil && !IsNil(o.AccountQuantity) {
		return true
	}

	return false
}

// SetAccountQuantity gets a reference to the given int32 and assigns it to the AccountQuantity field.
func (o *ClientDescribeOrderResponseBody) SetAccountQuantity(v int32) {
	o.AccountQuantity = &v
}

// GetAliUid returns the AliUid field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetAliUid() int32 {
	if o == nil || IsNil(o.AliUid) {
		var ret int32
		return ret
	}
	return *o.AliUid
}

// GetAliUidOk returns a tuple with the AliUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetAliUidOk() (*int32, bool) {
	if o == nil || IsNil(o.AliUid) {
		return nil, false
	}
	return o.AliUid, true
}

// HasAliUid returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasAliUid() bool {
	if o != nil && !IsNil(o.AliUid) {
		return true
	}

	return false
}

// SetAliUid gets a reference to the given int32 and assigns it to the AliUid field.
func (o *ClientDescribeOrderResponseBody) SetAliUid(v int32) {
	o.AliUid = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetComponents() map[string]interface{} {
	if o == nil || IsNil(o.Components) {
		var ret map[string]interface{}
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetComponentsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Components) {
		return map[string]interface{}{}, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given map[string]interface{} and assigns it to the Components field.
func (o *ClientDescribeOrderResponseBody) SetComponents(v map[string]interface{}) {
	o.Components = v
}

// GetCouponPrice returns the CouponPrice field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetCouponPrice() float32 {
	if o == nil || IsNil(o.CouponPrice) {
		var ret float32
		return ret
	}
	return *o.CouponPrice
}

// GetCouponPriceOk returns a tuple with the CouponPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetCouponPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.CouponPrice) {
		return nil, false
	}
	return o.CouponPrice, true
}

// HasCouponPrice returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasCouponPrice() bool {
	if o != nil && !IsNil(o.CouponPrice) {
		return true
	}

	return false
}

// SetCouponPrice gets a reference to the given float32 and assigns it to the CouponPrice field.
func (o *ClientDescribeOrderResponseBody) SetCouponPrice(v float32) {
	o.CouponPrice = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetCreatedOn() int32 {
	if o == nil || IsNil(o.CreatedOn) {
		var ret int32
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetCreatedOnOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given int32 and assigns it to the CreatedOn field.
func (o *ClientDescribeOrderResponseBody) SetCreatedOn(v int32) {
	o.CreatedOn = &v
}

// GetInstanceIds returns the InstanceIds field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetInstanceIds() ClientDescribeOrderResponseBodyInstanceIds {
	if o == nil || IsNil(o.InstanceIds) {
		var ret ClientDescribeOrderResponseBodyInstanceIds
		return ret
	}
	return *o.InstanceIds
}

// GetInstanceIdsOk returns a tuple with the InstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetInstanceIdsOk() (*ClientDescribeOrderResponseBodyInstanceIds, bool) {
	if o == nil || IsNil(o.InstanceIds) {
		return nil, false
	}
	return o.InstanceIds, true
}

// HasInstanceIds returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasInstanceIds() bool {
	if o != nil && !IsNil(o.InstanceIds) {
		return true
	}

	return false
}

// SetInstanceIds gets a reference to the given ClientDescribeOrderResponseBodyInstanceIds and assigns it to the InstanceIds field.
func (o *ClientDescribeOrderResponseBody) SetInstanceIds(v ClientDescribeOrderResponseBodyInstanceIds) {
	o.InstanceIds = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *ClientDescribeOrderResponseBody) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *ClientDescribeOrderResponseBody) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *ClientDescribeOrderResponseBody) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetOriginalPrice() float32 {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret float32
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetOriginalPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given float32 and assigns it to the OriginalPrice field.
func (o *ClientDescribeOrderResponseBody) SetOriginalPrice(v float32) {
	o.OriginalPrice = &v
}

// GetPaidOn returns the PaidOn field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetPaidOn() int32 {
	if o == nil || IsNil(o.PaidOn) {
		var ret int32
		return ret
	}
	return *o.PaidOn
}

// GetPaidOnOk returns a tuple with the PaidOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetPaidOnOk() (*int32, bool) {
	if o == nil || IsNil(o.PaidOn) {
		return nil, false
	}
	return o.PaidOn, true
}

// HasPaidOn returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasPaidOn() bool {
	if o != nil && !IsNil(o.PaidOn) {
		return true
	}

	return false
}

// SetPaidOn gets a reference to the given int32 and assigns it to the PaidOn field.
func (o *ClientDescribeOrderResponseBody) SetPaidOn(v int32) {
	o.PaidOn = &v
}

// GetPayStatus returns the PayStatus field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetPayStatus() string {
	if o == nil || IsNil(o.PayStatus) {
		var ret string
		return ret
	}
	return *o.PayStatus
}

// GetPayStatusOk returns a tuple with the PayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetPayStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PayStatus) {
		return nil, false
	}
	return o.PayStatus, true
}

// HasPayStatus returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasPayStatus() bool {
	if o != nil && !IsNil(o.PayStatus) {
		return true
	}

	return false
}

// SetPayStatus gets a reference to the given string and assigns it to the PayStatus field.
func (o *ClientDescribeOrderResponseBody) SetPayStatus(v string) {
	o.PayStatus = &v
}

// GetPaymentPrice returns the PaymentPrice field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetPaymentPrice() float32 {
	if o == nil || IsNil(o.PaymentPrice) {
		var ret float32
		return ret
	}
	return *o.PaymentPrice
}

// GetPaymentPriceOk returns a tuple with the PaymentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetPaymentPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.PaymentPrice) {
		return nil, false
	}
	return o.PaymentPrice, true
}

// HasPaymentPrice returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasPaymentPrice() bool {
	if o != nil && !IsNil(o.PaymentPrice) {
		return true
	}

	return false
}

// SetPaymentPrice gets a reference to the given float32 and assigns it to the PaymentPrice field.
func (o *ClientDescribeOrderResponseBody) SetPaymentPrice(v float32) {
	o.PaymentPrice = &v
}

// GetPeriodType returns the PeriodType field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetPeriodType() string {
	if o == nil || IsNil(o.PeriodType) {
		var ret string
		return ret
	}
	return *o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetPeriodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodType) {
		return nil, false
	}
	return o.PeriodType, true
}

// HasPeriodType returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasPeriodType() bool {
	if o != nil && !IsNil(o.PeriodType) {
		return true
	}

	return false
}

// SetPeriodType gets a reference to the given string and assigns it to the PeriodType field.
func (o *ClientDescribeOrderResponseBody) SetPeriodType(v string) {
	o.PeriodType = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *ClientDescribeOrderResponseBody) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *ClientDescribeOrderResponseBody) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductSkuCode returns the ProductSkuCode field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetProductSkuCode() string {
	if o == nil || IsNil(o.ProductSkuCode) {
		var ret string
		return ret
	}
	return *o.ProductSkuCode
}

// GetProductSkuCodeOk returns a tuple with the ProductSkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetProductSkuCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductSkuCode) {
		return nil, false
	}
	return o.ProductSkuCode, true
}

// HasProductSkuCode returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasProductSkuCode() bool {
	if o != nil && !IsNil(o.ProductSkuCode) {
		return true
	}

	return false
}

// SetProductSkuCode gets a reference to the given string and assigns it to the ProductSkuCode field.
func (o *ClientDescribeOrderResponseBody) SetProductSkuCode(v string) {
	o.ProductSkuCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ClientDescribeOrderResponseBody) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ClientDescribeOrderResponseBody) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSupplierCompanyName returns the SupplierCompanyName field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetSupplierCompanyName() string {
	if o == nil || IsNil(o.SupplierCompanyName) {
		var ret string
		return ret
	}
	return *o.SupplierCompanyName
}

// GetSupplierCompanyNameOk returns a tuple with the SupplierCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetSupplierCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.SupplierCompanyName) {
		return nil, false
	}
	return o.SupplierCompanyName, true
}

// HasSupplierCompanyName returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasSupplierCompanyName() bool {
	if o != nil && !IsNil(o.SupplierCompanyName) {
		return true
	}

	return false
}

// SetSupplierCompanyName gets a reference to the given string and assigns it to the SupplierCompanyName field.
func (o *ClientDescribeOrderResponseBody) SetSupplierCompanyName(v string) {
	o.SupplierCompanyName = &v
}

// GetSupplierTelephones returns the SupplierTelephones field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetSupplierTelephones() ClientDescribeOrderResponseBodySupplierTelephones {
	if o == nil || IsNil(o.SupplierTelephones) {
		var ret ClientDescribeOrderResponseBodySupplierTelephones
		return ret
	}
	return *o.SupplierTelephones
}

// GetSupplierTelephonesOk returns a tuple with the SupplierTelephones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetSupplierTelephonesOk() (*ClientDescribeOrderResponseBodySupplierTelephones, bool) {
	if o == nil || IsNil(o.SupplierTelephones) {
		return nil, false
	}
	return o.SupplierTelephones, true
}

// HasSupplierTelephones returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasSupplierTelephones() bool {
	if o != nil && !IsNil(o.SupplierTelephones) {
		return true
	}

	return false
}

// SetSupplierTelephones gets a reference to the given ClientDescribeOrderResponseBodySupplierTelephones and assigns it to the SupplierTelephones field.
func (o *ClientDescribeOrderResponseBody) SetSupplierTelephones(v ClientDescribeOrderResponseBodySupplierTelephones) {
	o.SupplierTelephones = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBody) GetTotalPrice() float32 {
	if o == nil || IsNil(o.TotalPrice) {
		var ret float32
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBody) GetTotalPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPrice) {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBody) HasTotalPrice() bool {
	if o != nil && !IsNil(o.TotalPrice) {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given float32 and assigns it to the TotalPrice field.
func (o *ClientDescribeOrderResponseBody) SetTotalPrice(v float32) {
	o.TotalPrice = &v
}

func (o ClientDescribeOrderResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientDescribeOrderResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountQuantity) {
		toSerialize["AccountQuantity"] = o.AccountQuantity
	}
	if !IsNil(o.AliUid) {
		toSerialize["AliUid"] = o.AliUid
	}
	if !IsNil(o.Components) {
		toSerialize["Components"] = o.Components
	}
	if !IsNil(o.CouponPrice) {
		toSerialize["CouponPrice"] = o.CouponPrice
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["CreatedOn"] = o.CreatedOn
	}
	if !IsNil(o.InstanceIds) {
		toSerialize["InstanceIds"] = o.InstanceIds
	}
	if !IsNil(o.OrderId) {
		toSerialize["OrderId"] = o.OrderId
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["OrderStatus"] = o.OrderStatus
	}
	if !IsNil(o.OrderType) {
		toSerialize["OrderType"] = o.OrderType
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["OriginalPrice"] = o.OriginalPrice
	}
	if !IsNil(o.PaidOn) {
		toSerialize["PaidOn"] = o.PaidOn
	}
	if !IsNil(o.PayStatus) {
		toSerialize["PayStatus"] = o.PayStatus
	}
	if !IsNil(o.PaymentPrice) {
		toSerialize["PaymentPrice"] = o.PaymentPrice
	}
	if !IsNil(o.PeriodType) {
		toSerialize["PeriodType"] = o.PeriodType
	}
	if !IsNil(o.ProductCode) {
		toSerialize["ProductCode"] = o.ProductCode
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}
	if !IsNil(o.ProductSkuCode) {
		toSerialize["ProductSkuCode"] = o.ProductSkuCode
	}
	if !IsNil(o.Quantity) {
		toSerialize["Quantity"] = o.Quantity
	}
	if !IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	if !IsNil(o.SupplierCompanyName) {
		toSerialize["SupplierCompanyName"] = o.SupplierCompanyName
	}
	if !IsNil(o.SupplierTelephones) {
		toSerialize["SupplierTelephones"] = o.SupplierTelephones
	}
	if !IsNil(o.TotalPrice) {
		toSerialize["TotalPrice"] = o.TotalPrice
	}
	return toSerialize, nil
}

type NullableClientDescribeOrderResponseBody struct {
	value *ClientDescribeOrderResponseBody
	isSet bool
}

func (v NullableClientDescribeOrderResponseBody) Get() *ClientDescribeOrderResponseBody {
	return v.value
}

func (v *NullableClientDescribeOrderResponseBody) Set(val *ClientDescribeOrderResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableClientDescribeOrderResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableClientDescribeOrderResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientDescribeOrderResponseBody(val *ClientDescribeOrderResponseBody) *NullableClientDescribeOrderResponseBody {
	return &NullableClientDescribeOrderResponseBody{value: val, isSet: true}
}

func (v NullableClientDescribeOrderResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientDescribeOrderResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
