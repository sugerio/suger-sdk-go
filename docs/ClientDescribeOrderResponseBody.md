# ClientDescribeOrderResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountQuantity** | Pointer to **int32** |  | [optional] 
**AliUid** | Pointer to **int32** |  | [optional] 
**Components** | Pointer to **map[string]interface{}** |  | [optional] 
**CouponPrice** | Pointer to **float32** |  | [optional] 
**CreatedOn** | Pointer to **int32** |  | [optional] 
**InstanceIds** | Pointer to [**ClientDescribeOrderResponseBodyInstanceIds**](ClientDescribeOrderResponseBodyInstanceIds.md) |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**OriginalPrice** | Pointer to **float32** |  | [optional] 
**PaidOn** | Pointer to **int32** |  | [optional] 
**PayStatus** | Pointer to **string** |  | [optional] 
**PaymentPrice** | Pointer to **float32** |  | [optional] 
**PeriodType** | Pointer to **string** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ProductSkuCode** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**SupplierCompanyName** | Pointer to **string** |  | [optional] 
**SupplierTelephones** | Pointer to [**ClientDescribeOrderResponseBodySupplierTelephones**](ClientDescribeOrderResponseBodySupplierTelephones.md) |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 

## Methods

### NewClientDescribeOrderResponseBody

`func NewClientDescribeOrderResponseBody() *ClientDescribeOrderResponseBody`

NewClientDescribeOrderResponseBody instantiates a new ClientDescribeOrderResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDescribeOrderResponseBodyWithDefaults

`func NewClientDescribeOrderResponseBodyWithDefaults() *ClientDescribeOrderResponseBody`

NewClientDescribeOrderResponseBodyWithDefaults instantiates a new ClientDescribeOrderResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountQuantity

`func (o *ClientDescribeOrderResponseBody) GetAccountQuantity() int32`

GetAccountQuantity returns the AccountQuantity field if non-nil, zero value otherwise.

### GetAccountQuantityOk

`func (o *ClientDescribeOrderResponseBody) GetAccountQuantityOk() (*int32, bool)`

GetAccountQuantityOk returns a tuple with the AccountQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQuantity

`func (o *ClientDescribeOrderResponseBody) SetAccountQuantity(v int32)`

SetAccountQuantity sets AccountQuantity field to given value.

### HasAccountQuantity

`func (o *ClientDescribeOrderResponseBody) HasAccountQuantity() bool`

HasAccountQuantity returns a boolean if a field has been set.

### GetAliUid

`func (o *ClientDescribeOrderResponseBody) GetAliUid() int32`

GetAliUid returns the AliUid field if non-nil, zero value otherwise.

### GetAliUidOk

`func (o *ClientDescribeOrderResponseBody) GetAliUidOk() (*int32, bool)`

GetAliUidOk returns a tuple with the AliUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliUid

`func (o *ClientDescribeOrderResponseBody) SetAliUid(v int32)`

SetAliUid sets AliUid field to given value.

### HasAliUid

`func (o *ClientDescribeOrderResponseBody) HasAliUid() bool`

HasAliUid returns a boolean if a field has been set.

### GetComponents

`func (o *ClientDescribeOrderResponseBody) GetComponents() map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ClientDescribeOrderResponseBody) GetComponentsOk() (*map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ClientDescribeOrderResponseBody) SetComponents(v map[string]interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ClientDescribeOrderResponseBody) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCouponPrice

`func (o *ClientDescribeOrderResponseBody) GetCouponPrice() float32`

GetCouponPrice returns the CouponPrice field if non-nil, zero value otherwise.

### GetCouponPriceOk

`func (o *ClientDescribeOrderResponseBody) GetCouponPriceOk() (*float32, bool)`

GetCouponPriceOk returns a tuple with the CouponPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponPrice

`func (o *ClientDescribeOrderResponseBody) SetCouponPrice(v float32)`

SetCouponPrice sets CouponPrice field to given value.

### HasCouponPrice

`func (o *ClientDescribeOrderResponseBody) HasCouponPrice() bool`

HasCouponPrice returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ClientDescribeOrderResponseBody) GetCreatedOn() int32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ClientDescribeOrderResponseBody) GetCreatedOnOk() (*int32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ClientDescribeOrderResponseBody) SetCreatedOn(v int32)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ClientDescribeOrderResponseBody) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetInstanceIds

`func (o *ClientDescribeOrderResponseBody) GetInstanceIds() ClientDescribeOrderResponseBodyInstanceIds`

GetInstanceIds returns the InstanceIds field if non-nil, zero value otherwise.

### GetInstanceIdsOk

`func (o *ClientDescribeOrderResponseBody) GetInstanceIdsOk() (*ClientDescribeOrderResponseBodyInstanceIds, bool)`

GetInstanceIdsOk returns a tuple with the InstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceIds

`func (o *ClientDescribeOrderResponseBody) SetInstanceIds(v ClientDescribeOrderResponseBodyInstanceIds)`

SetInstanceIds sets InstanceIds field to given value.

### HasInstanceIds

`func (o *ClientDescribeOrderResponseBody) HasInstanceIds() bool`

HasInstanceIds returns a boolean if a field has been set.

### GetOrderId

`func (o *ClientDescribeOrderResponseBody) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ClientDescribeOrderResponseBody) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ClientDescribeOrderResponseBody) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ClientDescribeOrderResponseBody) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *ClientDescribeOrderResponseBody) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *ClientDescribeOrderResponseBody) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *ClientDescribeOrderResponseBody) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *ClientDescribeOrderResponseBody) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderType

`func (o *ClientDescribeOrderResponseBody) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ClientDescribeOrderResponseBody) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ClientDescribeOrderResponseBody) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ClientDescribeOrderResponseBody) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *ClientDescribeOrderResponseBody) GetOriginalPrice() float32`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *ClientDescribeOrderResponseBody) GetOriginalPriceOk() (*float32, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *ClientDescribeOrderResponseBody) SetOriginalPrice(v float32)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *ClientDescribeOrderResponseBody) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetPaidOn

`func (o *ClientDescribeOrderResponseBody) GetPaidOn() int32`

GetPaidOn returns the PaidOn field if non-nil, zero value otherwise.

### GetPaidOnOk

`func (o *ClientDescribeOrderResponseBody) GetPaidOnOk() (*int32, bool)`

GetPaidOnOk returns a tuple with the PaidOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidOn

`func (o *ClientDescribeOrderResponseBody) SetPaidOn(v int32)`

SetPaidOn sets PaidOn field to given value.

### HasPaidOn

`func (o *ClientDescribeOrderResponseBody) HasPaidOn() bool`

HasPaidOn returns a boolean if a field has been set.

### GetPayStatus

`func (o *ClientDescribeOrderResponseBody) GetPayStatus() string`

GetPayStatus returns the PayStatus field if non-nil, zero value otherwise.

### GetPayStatusOk

`func (o *ClientDescribeOrderResponseBody) GetPayStatusOk() (*string, bool)`

GetPayStatusOk returns a tuple with the PayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayStatus

`func (o *ClientDescribeOrderResponseBody) SetPayStatus(v string)`

SetPayStatus sets PayStatus field to given value.

### HasPayStatus

`func (o *ClientDescribeOrderResponseBody) HasPayStatus() bool`

HasPayStatus returns a boolean if a field has been set.

### GetPaymentPrice

`func (o *ClientDescribeOrderResponseBody) GetPaymentPrice() float32`

GetPaymentPrice returns the PaymentPrice field if non-nil, zero value otherwise.

### GetPaymentPriceOk

`func (o *ClientDescribeOrderResponseBody) GetPaymentPriceOk() (*float32, bool)`

GetPaymentPriceOk returns a tuple with the PaymentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPrice

`func (o *ClientDescribeOrderResponseBody) SetPaymentPrice(v float32)`

SetPaymentPrice sets PaymentPrice field to given value.

### HasPaymentPrice

`func (o *ClientDescribeOrderResponseBody) HasPaymentPrice() bool`

HasPaymentPrice returns a boolean if a field has been set.

### GetPeriodType

`func (o *ClientDescribeOrderResponseBody) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *ClientDescribeOrderResponseBody) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *ClientDescribeOrderResponseBody) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *ClientDescribeOrderResponseBody) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetProductCode

`func (o *ClientDescribeOrderResponseBody) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ClientDescribeOrderResponseBody) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ClientDescribeOrderResponseBody) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ClientDescribeOrderResponseBody) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *ClientDescribeOrderResponseBody) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ClientDescribeOrderResponseBody) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ClientDescribeOrderResponseBody) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ClientDescribeOrderResponseBody) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductSkuCode

`func (o *ClientDescribeOrderResponseBody) GetProductSkuCode() string`

GetProductSkuCode returns the ProductSkuCode field if non-nil, zero value otherwise.

### GetProductSkuCodeOk

`func (o *ClientDescribeOrderResponseBody) GetProductSkuCodeOk() (*string, bool)`

GetProductSkuCodeOk returns a tuple with the ProductSkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSkuCode

`func (o *ClientDescribeOrderResponseBody) SetProductSkuCode(v string)`

SetProductSkuCode sets ProductSkuCode field to given value.

### HasProductSkuCode

`func (o *ClientDescribeOrderResponseBody) HasProductSkuCode() bool`

HasProductSkuCode returns a boolean if a field has been set.

### GetQuantity

`func (o *ClientDescribeOrderResponseBody) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ClientDescribeOrderResponseBody) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ClientDescribeOrderResponseBody) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ClientDescribeOrderResponseBody) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRequestId

`func (o *ClientDescribeOrderResponseBody) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClientDescribeOrderResponseBody) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClientDescribeOrderResponseBody) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ClientDescribeOrderResponseBody) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSupplierCompanyName

`func (o *ClientDescribeOrderResponseBody) GetSupplierCompanyName() string`

GetSupplierCompanyName returns the SupplierCompanyName field if non-nil, zero value otherwise.

### GetSupplierCompanyNameOk

`func (o *ClientDescribeOrderResponseBody) GetSupplierCompanyNameOk() (*string, bool)`

GetSupplierCompanyNameOk returns a tuple with the SupplierCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierCompanyName

`func (o *ClientDescribeOrderResponseBody) SetSupplierCompanyName(v string)`

SetSupplierCompanyName sets SupplierCompanyName field to given value.

### HasSupplierCompanyName

`func (o *ClientDescribeOrderResponseBody) HasSupplierCompanyName() bool`

HasSupplierCompanyName returns a boolean if a field has been set.

### GetSupplierTelephones

`func (o *ClientDescribeOrderResponseBody) GetSupplierTelephones() ClientDescribeOrderResponseBodySupplierTelephones`

GetSupplierTelephones returns the SupplierTelephones field if non-nil, zero value otherwise.

### GetSupplierTelephonesOk

`func (o *ClientDescribeOrderResponseBody) GetSupplierTelephonesOk() (*ClientDescribeOrderResponseBodySupplierTelephones, bool)`

GetSupplierTelephonesOk returns a tuple with the SupplierTelephones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierTelephones

`func (o *ClientDescribeOrderResponseBody) SetSupplierTelephones(v ClientDescribeOrderResponseBodySupplierTelephones)`

SetSupplierTelephones sets SupplierTelephones field to given value.

### HasSupplierTelephones

`func (o *ClientDescribeOrderResponseBody) HasSupplierTelephones() bool`

HasSupplierTelephones returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ClientDescribeOrderResponseBody) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ClientDescribeOrderResponseBody) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ClientDescribeOrderResponseBody) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *ClientDescribeOrderResponseBody) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


