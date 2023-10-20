# ClientDescribeInstanceResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppJson** | Pointer to **string** |  | [optional] 
**AutoRenewal** | Pointer to **string** |  | [optional] 
**BeganOn** | Pointer to **int32** |  | [optional] 
**ComponentJson** | Pointer to **string** |  | [optional] 
**Constraints** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **int32** |  | [optional] 
**EndOn** | Pointer to **int32** |  | [optional] 
**ExtendJson** | Pointer to **string** |  | [optional] 
**HostJson** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **int32** |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] 
**Modules** | Pointer to [**ClientDescribeInstanceResponseBodyModules**](ClientDescribeInstanceResponseBodyModules.md) |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ProductSkuCode** | Pointer to **string** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**RelationalData** | Pointer to [**ClientDescribeInstanceResponseBodyRelationalData**](ClientDescribeInstanceResponseBodyRelationalData.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupplierName** | Pointer to **string** |  | [optional] 

## Methods

### NewClientDescribeInstanceResponseBody

`func NewClientDescribeInstanceResponseBody() *ClientDescribeInstanceResponseBody`

NewClientDescribeInstanceResponseBody instantiates a new ClientDescribeInstanceResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDescribeInstanceResponseBodyWithDefaults

`func NewClientDescribeInstanceResponseBodyWithDefaults() *ClientDescribeInstanceResponseBody`

NewClientDescribeInstanceResponseBodyWithDefaults instantiates a new ClientDescribeInstanceResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppJson

`func (o *ClientDescribeInstanceResponseBody) GetAppJson() string`

GetAppJson returns the AppJson field if non-nil, zero value otherwise.

### GetAppJsonOk

`func (o *ClientDescribeInstanceResponseBody) GetAppJsonOk() (*string, bool)`

GetAppJsonOk returns a tuple with the AppJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppJson

`func (o *ClientDescribeInstanceResponseBody) SetAppJson(v string)`

SetAppJson sets AppJson field to given value.

### HasAppJson

`func (o *ClientDescribeInstanceResponseBody) HasAppJson() bool`

HasAppJson returns a boolean if a field has been set.

### GetAutoRenewal

`func (o *ClientDescribeInstanceResponseBody) GetAutoRenewal() string`

GetAutoRenewal returns the AutoRenewal field if non-nil, zero value otherwise.

### GetAutoRenewalOk

`func (o *ClientDescribeInstanceResponseBody) GetAutoRenewalOk() (*string, bool)`

GetAutoRenewalOk returns a tuple with the AutoRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewal

`func (o *ClientDescribeInstanceResponseBody) SetAutoRenewal(v string)`

SetAutoRenewal sets AutoRenewal field to given value.

### HasAutoRenewal

`func (o *ClientDescribeInstanceResponseBody) HasAutoRenewal() bool`

HasAutoRenewal returns a boolean if a field has been set.

### GetBeganOn

`func (o *ClientDescribeInstanceResponseBody) GetBeganOn() int32`

GetBeganOn returns the BeganOn field if non-nil, zero value otherwise.

### GetBeganOnOk

`func (o *ClientDescribeInstanceResponseBody) GetBeganOnOk() (*int32, bool)`

GetBeganOnOk returns a tuple with the BeganOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeganOn

`func (o *ClientDescribeInstanceResponseBody) SetBeganOn(v int32)`

SetBeganOn sets BeganOn field to given value.

### HasBeganOn

`func (o *ClientDescribeInstanceResponseBody) HasBeganOn() bool`

HasBeganOn returns a boolean if a field has been set.

### GetComponentJson

`func (o *ClientDescribeInstanceResponseBody) GetComponentJson() string`

GetComponentJson returns the ComponentJson field if non-nil, zero value otherwise.

### GetComponentJsonOk

`func (o *ClientDescribeInstanceResponseBody) GetComponentJsonOk() (*string, bool)`

GetComponentJsonOk returns a tuple with the ComponentJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentJson

`func (o *ClientDescribeInstanceResponseBody) SetComponentJson(v string)`

SetComponentJson sets ComponentJson field to given value.

### HasComponentJson

`func (o *ClientDescribeInstanceResponseBody) HasComponentJson() bool`

HasComponentJson returns a boolean if a field has been set.

### GetConstraints

`func (o *ClientDescribeInstanceResponseBody) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ClientDescribeInstanceResponseBody) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ClientDescribeInstanceResponseBody) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *ClientDescribeInstanceResponseBody) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ClientDescribeInstanceResponseBody) GetCreatedOn() int32`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ClientDescribeInstanceResponseBody) GetCreatedOnOk() (*int32, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ClientDescribeInstanceResponseBody) SetCreatedOn(v int32)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ClientDescribeInstanceResponseBody) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetEndOn

`func (o *ClientDescribeInstanceResponseBody) GetEndOn() int32`

GetEndOn returns the EndOn field if non-nil, zero value otherwise.

### GetEndOnOk

`func (o *ClientDescribeInstanceResponseBody) GetEndOnOk() (*int32, bool)`

GetEndOnOk returns a tuple with the EndOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOn

`func (o *ClientDescribeInstanceResponseBody) SetEndOn(v int32)`

SetEndOn sets EndOn field to given value.

### HasEndOn

`func (o *ClientDescribeInstanceResponseBody) HasEndOn() bool`

HasEndOn returns a boolean if a field has been set.

### GetExtendJson

`func (o *ClientDescribeInstanceResponseBody) GetExtendJson() string`

GetExtendJson returns the ExtendJson field if non-nil, zero value otherwise.

### GetExtendJsonOk

`func (o *ClientDescribeInstanceResponseBody) GetExtendJsonOk() (*string, bool)`

GetExtendJsonOk returns a tuple with the ExtendJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendJson

`func (o *ClientDescribeInstanceResponseBody) SetExtendJson(v string)`

SetExtendJson sets ExtendJson field to given value.

### HasExtendJson

`func (o *ClientDescribeInstanceResponseBody) HasExtendJson() bool`

HasExtendJson returns a boolean if a field has been set.

### GetHostJson

`func (o *ClientDescribeInstanceResponseBody) GetHostJson() string`

GetHostJson returns the HostJson field if non-nil, zero value otherwise.

### GetHostJsonOk

`func (o *ClientDescribeInstanceResponseBody) GetHostJsonOk() (*string, bool)`

GetHostJsonOk returns a tuple with the HostJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostJson

`func (o *ClientDescribeInstanceResponseBody) SetHostJson(v string)`

SetHostJson sets HostJson field to given value.

### HasHostJson

`func (o *ClientDescribeInstanceResponseBody) HasHostJson() bool`

HasHostJson returns a boolean if a field has been set.

### GetInstanceId

`func (o *ClientDescribeInstanceResponseBody) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClientDescribeInstanceResponseBody) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClientDescribeInstanceResponseBody) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClientDescribeInstanceResponseBody) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetIsTrial

`func (o *ClientDescribeInstanceResponseBody) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *ClientDescribeInstanceResponseBody) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *ClientDescribeInstanceResponseBody) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *ClientDescribeInstanceResponseBody) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetModules

`func (o *ClientDescribeInstanceResponseBody) GetModules() ClientDescribeInstanceResponseBodyModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ClientDescribeInstanceResponseBody) GetModulesOk() (*ClientDescribeInstanceResponseBodyModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ClientDescribeInstanceResponseBody) SetModules(v ClientDescribeInstanceResponseBodyModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ClientDescribeInstanceResponseBody) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetOrderId

`func (o *ClientDescribeInstanceResponseBody) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ClientDescribeInstanceResponseBody) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ClientDescribeInstanceResponseBody) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ClientDescribeInstanceResponseBody) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetProductCode

`func (o *ClientDescribeInstanceResponseBody) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ClientDescribeInstanceResponseBody) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ClientDescribeInstanceResponseBody) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *ClientDescribeInstanceResponseBody) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *ClientDescribeInstanceResponseBody) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ClientDescribeInstanceResponseBody) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ClientDescribeInstanceResponseBody) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ClientDescribeInstanceResponseBody) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductSkuCode

`func (o *ClientDescribeInstanceResponseBody) GetProductSkuCode() string`

GetProductSkuCode returns the ProductSkuCode field if non-nil, zero value otherwise.

### GetProductSkuCodeOk

`func (o *ClientDescribeInstanceResponseBody) GetProductSkuCodeOk() (*string, bool)`

GetProductSkuCodeOk returns a tuple with the ProductSkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSkuCode

`func (o *ClientDescribeInstanceResponseBody) SetProductSkuCode(v string)`

SetProductSkuCode sets ProductSkuCode field to given value.

### HasProductSkuCode

`func (o *ClientDescribeInstanceResponseBody) HasProductSkuCode() bool`

HasProductSkuCode returns a boolean if a field has been set.

### GetProductType

`func (o *ClientDescribeInstanceResponseBody) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ClientDescribeInstanceResponseBody) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ClientDescribeInstanceResponseBody) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ClientDescribeInstanceResponseBody) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRelationalData

`func (o *ClientDescribeInstanceResponseBody) GetRelationalData() ClientDescribeInstanceResponseBodyRelationalData`

GetRelationalData returns the RelationalData field if non-nil, zero value otherwise.

### GetRelationalDataOk

`func (o *ClientDescribeInstanceResponseBody) GetRelationalDataOk() (*ClientDescribeInstanceResponseBodyRelationalData, bool)`

GetRelationalDataOk returns a tuple with the RelationalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationalData

`func (o *ClientDescribeInstanceResponseBody) SetRelationalData(v ClientDescribeInstanceResponseBodyRelationalData)`

SetRelationalData sets RelationalData field to given value.

### HasRelationalData

`func (o *ClientDescribeInstanceResponseBody) HasRelationalData() bool`

HasRelationalData returns a boolean if a field has been set.

### GetStatus

`func (o *ClientDescribeInstanceResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientDescribeInstanceResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientDescribeInstanceResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientDescribeInstanceResponseBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupplierName

`func (o *ClientDescribeInstanceResponseBody) GetSupplierName() string`

GetSupplierName returns the SupplierName field if non-nil, zero value otherwise.

### GetSupplierNameOk

`func (o *ClientDescribeInstanceResponseBody) GetSupplierNameOk() (*string, bool)`

GetSupplierNameOk returns a tuple with the SupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierName

`func (o *ClientDescribeInstanceResponseBody) SetSupplierName(v string)`

SetSupplierName sets SupplierName field to given value.

### HasSupplierName

`func (o *ClientDescribeInstanceResponseBody) HasSupplierName() bool`

HasSupplierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


