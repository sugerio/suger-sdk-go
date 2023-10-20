# OrbPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableMetric** | Pointer to [**OrbBillableMetric**](OrbBillableMetric.md) |  | [optional] 
**BpsConfig** | Pointer to [**OrbPriceModelConfigBPS**](OrbPriceModelConfigBPS.md) |  | [optional] 
**BulkBpsConfig** | Pointer to [**OrbPriceModelConfigBULKBPS**](OrbPriceModelConfigBULKBPS.md) |  | [optional] 
**BulkConfig** | Pointer to [**OrbPriceModelConfigBULK**](OrbPriceModelConfigBULK.md) |  | [optional] 
**Cadence** | Pointer to [**OrbCadence**](OrbCadence.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**OrbPriceDiscount**](OrbPriceDiscount.md) |  | [optional] 
**FixedPriceQuantity** | Pointer to **int32** | If the Price represents a fixed cost, this represents the quantity of units applied. Mutually exclusive with billable_metric. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**OrbItem**](OrbItem.md) |  | [optional] 
**MatrixConfig** | Pointer to [**OrbPriceModelConfigMATRIX**](OrbPriceModelConfigMATRIX.md) |  | [optional] 
**Maximum** | Pointer to [**OrbPriceMaximum**](OrbPriceMaximum.md) |  | [optional] 
**Minimum** | Pointer to [**OrbPriceMinimum**](OrbPriceMinimum.md) |  | [optional] 
**ModelType** | Pointer to [**OrbPriceModelType**](OrbPriceModelType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackageConfig** | Pointer to [**OrbPriceModelConfigPACKAGE**](OrbPriceModelConfigPACKAGE.md) |  | [optional] 
**PlanPhaseOrder** | Pointer to **int32** | The phase order which includes this price, only applicable to a plan with phases. | [optional] 
**TieredBpsConfig** | Pointer to [**OrbPriceModelConfigTIEREDBPS**](OrbPriceModelConfigTIEREDBPS.md) |  | [optional] 
**TieredConfig** | Pointer to [**OrbPriceModelConfigTIERED**](OrbPriceModelConfigTIERED.md) |  | [optional] 
**UnitConfig** | Pointer to [**OrbPriceModelConfigUNIT**](OrbPriceModelConfigUNIT.md) |  | [optional] 

## Methods

### NewOrbPrice

`func NewOrbPrice() *OrbPrice`

NewOrbPrice instantiates a new OrbPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceWithDefaults

`func NewOrbPriceWithDefaults() *OrbPrice`

NewOrbPriceWithDefaults instantiates a new OrbPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableMetric

`func (o *OrbPrice) GetBillableMetric() OrbBillableMetric`

GetBillableMetric returns the BillableMetric field if non-nil, zero value otherwise.

### GetBillableMetricOk

`func (o *OrbPrice) GetBillableMetricOk() (*OrbBillableMetric, bool)`

GetBillableMetricOk returns a tuple with the BillableMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetric

`func (o *OrbPrice) SetBillableMetric(v OrbBillableMetric)`

SetBillableMetric sets BillableMetric field to given value.

### HasBillableMetric

`func (o *OrbPrice) HasBillableMetric() bool`

HasBillableMetric returns a boolean if a field has been set.

### GetBpsConfig

`func (o *OrbPrice) GetBpsConfig() OrbPriceModelConfigBPS`

GetBpsConfig returns the BpsConfig field if non-nil, zero value otherwise.

### GetBpsConfigOk

`func (o *OrbPrice) GetBpsConfigOk() (*OrbPriceModelConfigBPS, bool)`

GetBpsConfigOk returns a tuple with the BpsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpsConfig

`func (o *OrbPrice) SetBpsConfig(v OrbPriceModelConfigBPS)`

SetBpsConfig sets BpsConfig field to given value.

### HasBpsConfig

`func (o *OrbPrice) HasBpsConfig() bool`

HasBpsConfig returns a boolean if a field has been set.

### GetBulkBpsConfig

`func (o *OrbPrice) GetBulkBpsConfig() OrbPriceModelConfigBULKBPS`

GetBulkBpsConfig returns the BulkBpsConfig field if non-nil, zero value otherwise.

### GetBulkBpsConfigOk

`func (o *OrbPrice) GetBulkBpsConfigOk() (*OrbPriceModelConfigBULKBPS, bool)`

GetBulkBpsConfigOk returns a tuple with the BulkBpsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkBpsConfig

`func (o *OrbPrice) SetBulkBpsConfig(v OrbPriceModelConfigBULKBPS)`

SetBulkBpsConfig sets BulkBpsConfig field to given value.

### HasBulkBpsConfig

`func (o *OrbPrice) HasBulkBpsConfig() bool`

HasBulkBpsConfig returns a boolean if a field has been set.

### GetBulkConfig

`func (o *OrbPrice) GetBulkConfig() OrbPriceModelConfigBULK`

GetBulkConfig returns the BulkConfig field if non-nil, zero value otherwise.

### GetBulkConfigOk

`func (o *OrbPrice) GetBulkConfigOk() (*OrbPriceModelConfigBULK, bool)`

GetBulkConfigOk returns a tuple with the BulkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkConfig

`func (o *OrbPrice) SetBulkConfig(v OrbPriceModelConfigBULK)`

SetBulkConfig sets BulkConfig field to given value.

### HasBulkConfig

`func (o *OrbPrice) HasBulkConfig() bool`

HasBulkConfig returns a boolean if a field has been set.

### GetCadence

`func (o *OrbPrice) GetCadence() OrbCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *OrbPrice) GetCadenceOk() (*OrbCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *OrbPrice) SetCadence(v OrbCadence)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *OrbPrice) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrbPrice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrbPrice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrbPrice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrbPrice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *OrbPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrbPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrbPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrbPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *OrbPrice) GetDiscount() OrbPriceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrbPrice) GetDiscountOk() (*OrbPriceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrbPrice) SetDiscount(v OrbPriceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrbPrice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetFixedPriceQuantity

`func (o *OrbPrice) GetFixedPriceQuantity() int32`

GetFixedPriceQuantity returns the FixedPriceQuantity field if non-nil, zero value otherwise.

### GetFixedPriceQuantityOk

`func (o *OrbPrice) GetFixedPriceQuantityOk() (*int32, bool)`

GetFixedPriceQuantityOk returns a tuple with the FixedPriceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPriceQuantity

`func (o *OrbPrice) SetFixedPriceQuantity(v int32)`

SetFixedPriceQuantity sets FixedPriceQuantity field to given value.

### HasFixedPriceQuantity

`func (o *OrbPrice) HasFixedPriceQuantity() bool`

HasFixedPriceQuantity returns a boolean if a field has been set.

### GetId

`func (o *OrbPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrbPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrbPrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrbPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItem

`func (o *OrbPrice) GetItem() OrbItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *OrbPrice) GetItemOk() (*OrbItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *OrbPrice) SetItem(v OrbItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *OrbPrice) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetMatrixConfig

`func (o *OrbPrice) GetMatrixConfig() OrbPriceModelConfigMATRIX`

GetMatrixConfig returns the MatrixConfig field if non-nil, zero value otherwise.

### GetMatrixConfigOk

`func (o *OrbPrice) GetMatrixConfigOk() (*OrbPriceModelConfigMATRIX, bool)`

GetMatrixConfigOk returns a tuple with the MatrixConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrixConfig

`func (o *OrbPrice) SetMatrixConfig(v OrbPriceModelConfigMATRIX)`

SetMatrixConfig sets MatrixConfig field to given value.

### HasMatrixConfig

`func (o *OrbPrice) HasMatrixConfig() bool`

HasMatrixConfig returns a boolean if a field has been set.

### GetMaximum

`func (o *OrbPrice) GetMaximum() OrbPriceMaximum`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *OrbPrice) GetMaximumOk() (*OrbPriceMaximum, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *OrbPrice) SetMaximum(v OrbPriceMaximum)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *OrbPrice) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *OrbPrice) GetMinimum() OrbPriceMinimum`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OrbPrice) GetMinimumOk() (*OrbPriceMinimum, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OrbPrice) SetMinimum(v OrbPriceMinimum)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *OrbPrice) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetModelType

`func (o *OrbPrice) GetModelType() OrbPriceModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *OrbPrice) GetModelTypeOk() (*OrbPriceModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *OrbPrice) SetModelType(v OrbPriceModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *OrbPrice) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetName

`func (o *OrbPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrbPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrbPrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrbPrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageConfig

`func (o *OrbPrice) GetPackageConfig() OrbPriceModelConfigPACKAGE`

GetPackageConfig returns the PackageConfig field if non-nil, zero value otherwise.

### GetPackageConfigOk

`func (o *OrbPrice) GetPackageConfigOk() (*OrbPriceModelConfigPACKAGE, bool)`

GetPackageConfigOk returns a tuple with the PackageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageConfig

`func (o *OrbPrice) SetPackageConfig(v OrbPriceModelConfigPACKAGE)`

SetPackageConfig sets PackageConfig field to given value.

### HasPackageConfig

`func (o *OrbPrice) HasPackageConfig() bool`

HasPackageConfig returns a boolean if a field has been set.

### GetPlanPhaseOrder

`func (o *OrbPrice) GetPlanPhaseOrder() int32`

GetPlanPhaseOrder returns the PlanPhaseOrder field if non-nil, zero value otherwise.

### GetPlanPhaseOrderOk

`func (o *OrbPrice) GetPlanPhaseOrderOk() (*int32, bool)`

GetPlanPhaseOrderOk returns a tuple with the PlanPhaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPhaseOrder

`func (o *OrbPrice) SetPlanPhaseOrder(v int32)`

SetPlanPhaseOrder sets PlanPhaseOrder field to given value.

### HasPlanPhaseOrder

`func (o *OrbPrice) HasPlanPhaseOrder() bool`

HasPlanPhaseOrder returns a boolean if a field has been set.

### GetTieredBpsConfig

`func (o *OrbPrice) GetTieredBpsConfig() OrbPriceModelConfigTIEREDBPS`

GetTieredBpsConfig returns the TieredBpsConfig field if non-nil, zero value otherwise.

### GetTieredBpsConfigOk

`func (o *OrbPrice) GetTieredBpsConfigOk() (*OrbPriceModelConfigTIEREDBPS, bool)`

GetTieredBpsConfigOk returns a tuple with the TieredBpsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredBpsConfig

`func (o *OrbPrice) SetTieredBpsConfig(v OrbPriceModelConfigTIEREDBPS)`

SetTieredBpsConfig sets TieredBpsConfig field to given value.

### HasTieredBpsConfig

`func (o *OrbPrice) HasTieredBpsConfig() bool`

HasTieredBpsConfig returns a boolean if a field has been set.

### GetTieredConfig

`func (o *OrbPrice) GetTieredConfig() OrbPriceModelConfigTIERED`

GetTieredConfig returns the TieredConfig field if non-nil, zero value otherwise.

### GetTieredConfigOk

`func (o *OrbPrice) GetTieredConfigOk() (*OrbPriceModelConfigTIERED, bool)`

GetTieredConfigOk returns a tuple with the TieredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredConfig

`func (o *OrbPrice) SetTieredConfig(v OrbPriceModelConfigTIERED)`

SetTieredConfig sets TieredConfig field to given value.

### HasTieredConfig

`func (o *OrbPrice) HasTieredConfig() bool`

HasTieredConfig returns a boolean if a field has been set.

### GetUnitConfig

`func (o *OrbPrice) GetUnitConfig() OrbPriceModelConfigUNIT`

GetUnitConfig returns the UnitConfig field if non-nil, zero value otherwise.

### GetUnitConfigOk

`func (o *OrbPrice) GetUnitConfigOk() (*OrbPriceModelConfigUNIT, bool)`

GetUnitConfigOk returns a tuple with the UnitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitConfig

`func (o *OrbPrice) SetUnitConfig(v OrbPriceModelConfigUNIT)`

SetUnitConfig sets UnitConfig field to given value.

### HasUnitConfig

`func (o *OrbPrice) HasUnitConfig() bool`

HasUnitConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


