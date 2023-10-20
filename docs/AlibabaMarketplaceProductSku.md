# AlibabaMarketplaceProductSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **string** | POSTPAY or PREPAY | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Constraints** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Modules** | Pointer to [**AlibabaMarketplaceProductSkuModules**](AlibabaMarketplaceProductSkuModules.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrderPeriods** | Pointer to [**AlibabaMarketplaceProductSkuOrderPeriods**](AlibabaMarketplaceProductSkuOrderPeriods.md) |  | [optional] 

## Methods

### NewAlibabaMarketplaceProductSku

`func NewAlibabaMarketplaceProductSku() *AlibabaMarketplaceProductSku`

NewAlibabaMarketplaceProductSku instantiates a new AlibabaMarketplaceProductSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlibabaMarketplaceProductSkuWithDefaults

`func NewAlibabaMarketplaceProductSkuWithDefaults() *AlibabaMarketplaceProductSku`

NewAlibabaMarketplaceProductSkuWithDefaults instantiates a new AlibabaMarketplaceProductSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *AlibabaMarketplaceProductSku) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *AlibabaMarketplaceProductSku) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *AlibabaMarketplaceProductSku) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *AlibabaMarketplaceProductSku) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetCode

`func (o *AlibabaMarketplaceProductSku) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AlibabaMarketplaceProductSku) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AlibabaMarketplaceProductSku) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AlibabaMarketplaceProductSku) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConstraints

`func (o *AlibabaMarketplaceProductSku) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AlibabaMarketplaceProductSku) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AlibabaMarketplaceProductSku) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AlibabaMarketplaceProductSku) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetHidden

`func (o *AlibabaMarketplaceProductSku) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *AlibabaMarketplaceProductSku) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *AlibabaMarketplaceProductSku) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *AlibabaMarketplaceProductSku) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetModules

`func (o *AlibabaMarketplaceProductSku) GetModules() AlibabaMarketplaceProductSkuModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AlibabaMarketplaceProductSku) GetModulesOk() (*AlibabaMarketplaceProductSkuModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AlibabaMarketplaceProductSku) SetModules(v AlibabaMarketplaceProductSkuModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *AlibabaMarketplaceProductSku) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetName

`func (o *AlibabaMarketplaceProductSku) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlibabaMarketplaceProductSku) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlibabaMarketplaceProductSku) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlibabaMarketplaceProductSku) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderPeriods

`func (o *AlibabaMarketplaceProductSku) GetOrderPeriods() AlibabaMarketplaceProductSkuOrderPeriods`

GetOrderPeriods returns the OrderPeriods field if non-nil, zero value otherwise.

### GetOrderPeriodsOk

`func (o *AlibabaMarketplaceProductSku) GetOrderPeriodsOk() (*AlibabaMarketplaceProductSkuOrderPeriods, bool)`

GetOrderPeriodsOk returns a tuple with the OrderPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPeriods

`func (o *AlibabaMarketplaceProductSku) SetOrderPeriods(v AlibabaMarketplaceProductSkuOrderPeriods)`

SetOrderPeriods sets OrderPeriods field to given value.

### HasOrderPeriods

`func (o *AlibabaMarketplaceProductSku) HasOrderPeriods() bool`

HasOrderPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


