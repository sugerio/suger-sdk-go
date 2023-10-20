# MeteringDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IncludedBaseQuantities** | Pointer to [**[]AzureIncludedBaseQuantity**](AzureIncludedBaseQuantity.md) | how many quantities of this dimension are included in the commit. | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Display name of the dimension. For GCP Marketplace, it is the metering metric ID without plan prefix. | [optional] 
**PlanId** | Pointer to **string** | The plan ID of the metering dimension. Applicable to GCP Marketplace only. No ISO duration suffix. | [optional] 
**PlanName** | Pointer to **string** | The name of the plan for the metering dimension. Applicable to GCP Marketplace only. It may contains the ISO duration suffix, such as P1Y. | [optional] 
**PriceTiers** | Pointer to [**[]GcpPriceTier**](GcpPriceTier.md) | The price tiers of the metering dimension. Applicable to GCP Marketplace only. | [optional] 
**Rate** | Pointer to **float32** | The unit price of this usage metering dimension. | [optional] 
**SkuId** | Pointer to **string** | The SKU ID of the metering dimension. Applicable to GCP Marketplace only. | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**UsageCount** | Pointer to [**UsageCount**](UsageCount.md) |  | [optional] 
**ValueType** | Pointer to [**ValueType**](ValueType.md) |  | [optional] 

## Methods

### NewMeteringDimension

`func NewMeteringDimension() *MeteringDimension`

NewMeteringDimension instantiates a new MeteringDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringDimensionWithDefaults

`func NewMeteringDimensionWithDefaults() *MeteringDimension`

NewMeteringDimensionWithDefaults instantiates a new MeteringDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *MeteringDimension) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MeteringDimension) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MeteringDimension) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MeteringDimension) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *MeteringDimension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeteringDimension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeteringDimension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeteringDimension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludedBaseQuantities

`func (o *MeteringDimension) GetIncludedBaseQuantities() []AzureIncludedBaseQuantity`

GetIncludedBaseQuantities returns the IncludedBaseQuantities field if non-nil, zero value otherwise.

### GetIncludedBaseQuantitiesOk

`func (o *MeteringDimension) GetIncludedBaseQuantitiesOk() (*[]AzureIncludedBaseQuantity, bool)`

GetIncludedBaseQuantitiesOk returns a tuple with the IncludedBaseQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedBaseQuantities

`func (o *MeteringDimension) SetIncludedBaseQuantities(v []AzureIncludedBaseQuantity)`

SetIncludedBaseQuantities sets IncludedBaseQuantities field to given value.

### HasIncludedBaseQuantities

`func (o *MeteringDimension) HasIncludedBaseQuantities() bool`

HasIncludedBaseQuantities returns a boolean if a field has been set.

### GetKey

`func (o *MeteringDimension) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MeteringDimension) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MeteringDimension) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MeteringDimension) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *MeteringDimension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeteringDimension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeteringDimension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeteringDimension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanId

`func (o *MeteringDimension) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *MeteringDimension) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *MeteringDimension) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *MeteringDimension) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *MeteringDimension) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *MeteringDimension) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *MeteringDimension) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *MeteringDimension) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPriceTiers

`func (o *MeteringDimension) GetPriceTiers() []GcpPriceTier`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *MeteringDimension) GetPriceTiersOk() (*[]GcpPriceTier, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *MeteringDimension) SetPriceTiers(v []GcpPriceTier)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *MeteringDimension) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetRate

`func (o *MeteringDimension) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MeteringDimension) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MeteringDimension) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MeteringDimension) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSkuId

`func (o *MeteringDimension) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MeteringDimension) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MeteringDimension) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MeteringDimension) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetTypes

`func (o *MeteringDimension) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *MeteringDimension) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *MeteringDimension) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *MeteringDimension) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUsageCount

`func (o *MeteringDimension) GetUsageCount() UsageCount`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *MeteringDimension) GetUsageCountOk() (*UsageCount, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *MeteringDimension) SetUsageCount(v UsageCount)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *MeteringDimension) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### GetValueType

`func (o *MeteringDimension) GetValueType() ValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MeteringDimension) GetValueTypeOk() (*ValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *MeteringDimension) SetValueType(v ValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *MeteringDimension) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


