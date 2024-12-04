# BillableDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableMetricId** | Pointer to **string** | The ID for the billable metric. | [optional] 
**Category** | Pointer to [**PriceModelCategory**](PriceModelCategory.md) | The category of the pricing model. This is used to determine which pricing model to use. | [optional] 
**Description** | Pointer to **string** | Description of the dimension. This is used in the UI to display the dimension. | [optional] 
**Discount** | Pointer to [**BillingDiscount**](BillingDiscount.md) | Discount for the dimension. | [optional] 
**Length** | Pointer to **int32** | The term length for the commit amount. Applicable to Direct only. | [optional] 
**MinimumCommit** | Pointer to **float32** | The minimum commit amount. Applicable to Direct only. Ignored if the value is 0 or less. | [optional] 
**MinimumCommitProrata** | Pointer to **bool** | If the minimum commit is appled with pro-rata. Applicable to Direct only. If true, the minimum commit amount will be prorated based on the usage period (starting time and ending time). | [optional] 
**MinimumCommitScope** | Pointer to [**BillingMinimumCommitScope**](BillingMinimumCommitScope.md) | The minimum commit scope. The default value is \&quot;DIMENSION\&quot; if not set. | [optional] 
**Name** | Pointer to **string** | Display name of the dimension. This is used in the UI to display the dimension. | [optional] 
**PriceModelBasic** | Pointer to [**PriceModelBasic**](PriceModelBasic.md) | The configuration for the Basic pricing model. Applicable to Direct only. | [optional] 
**PriceModelBulk** | Pointer to [**PriceModelBulk**](PriceModelBulk.md) | The configuration for the Package pricing model. Applicable to Direct only. | [optional] 
**PriceModelMatrix** | Pointer to [**PriceModelMatrix**](PriceModelMatrix.md) | The configuration for the Matrix pricing model. Applicable to Direct only. | [optional] 
**PriceModelPercentage** | Pointer to [**PriceModelPercentage**](PriceModelPercentage.md) | The configuration for the Percentage pricing model. Applicable to Direct only. | [optional] 
**PriceModelTiered** | Pointer to [**PriceModelTiered**](PriceModelTiered.md) | The configuration for the Tiered pricing model. Applicable to Direct only. | [optional] 
**PriceModelTieredPercentage** | Pointer to [**PriceModelTieredPercentage**](PriceModelTieredPercentage.md) | The configuration for the Tiered Percentage pricing model. Applicable to Direct only. | [optional] 
**PriceModelVolume** | Pointer to [**PriceModelVolume**](PriceModelVolume.md) | The configuration for the Bulk pricing model. Applicable to Direct only. | [optional] 
**TimeUnit** | Pointer to [**TimeUnit**](TimeUnit.md) | The term unit for the commit amount. Applicable to Direct only. | [optional] 

## Methods

### NewBillableDimension

`func NewBillableDimension() *BillableDimension`

NewBillableDimension instantiates a new BillableDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableDimensionWithDefaults

`func NewBillableDimensionWithDefaults() *BillableDimension`

NewBillableDimensionWithDefaults instantiates a new BillableDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableMetricId

`func (o *BillableDimension) GetBillableMetricId() string`

GetBillableMetricId returns the BillableMetricId field if non-nil, zero value otherwise.

### GetBillableMetricIdOk

`func (o *BillableDimension) GetBillableMetricIdOk() (*string, bool)`

GetBillableMetricIdOk returns a tuple with the BillableMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricId

`func (o *BillableDimension) SetBillableMetricId(v string)`

SetBillableMetricId sets BillableMetricId field to given value.

### HasBillableMetricId

`func (o *BillableDimension) HasBillableMetricId() bool`

HasBillableMetricId returns a boolean if a field has been set.

### GetCategory

`func (o *BillableDimension) GetCategory() PriceModelCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BillableDimension) GetCategoryOk() (*PriceModelCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BillableDimension) SetCategory(v PriceModelCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BillableDimension) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *BillableDimension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillableDimension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillableDimension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillableDimension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscount

`func (o *BillableDimension) GetDiscount() BillingDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillableDimension) GetDiscountOk() (*BillingDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillableDimension) SetDiscount(v BillingDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillableDimension) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetLength

`func (o *BillableDimension) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *BillableDimension) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *BillableDimension) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *BillableDimension) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMinimumCommit

`func (o *BillableDimension) GetMinimumCommit() float32`

GetMinimumCommit returns the MinimumCommit field if non-nil, zero value otherwise.

### GetMinimumCommitOk

`func (o *BillableDimension) GetMinimumCommitOk() (*float32, bool)`

GetMinimumCommitOk returns a tuple with the MinimumCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCommit

`func (o *BillableDimension) SetMinimumCommit(v float32)`

SetMinimumCommit sets MinimumCommit field to given value.

### HasMinimumCommit

`func (o *BillableDimension) HasMinimumCommit() bool`

HasMinimumCommit returns a boolean if a field has been set.

### GetMinimumCommitProrata

`func (o *BillableDimension) GetMinimumCommitProrata() bool`

GetMinimumCommitProrata returns the MinimumCommitProrata field if non-nil, zero value otherwise.

### GetMinimumCommitProrataOk

`func (o *BillableDimension) GetMinimumCommitProrataOk() (*bool, bool)`

GetMinimumCommitProrataOk returns a tuple with the MinimumCommitProrata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCommitProrata

`func (o *BillableDimension) SetMinimumCommitProrata(v bool)`

SetMinimumCommitProrata sets MinimumCommitProrata field to given value.

### HasMinimumCommitProrata

`func (o *BillableDimension) HasMinimumCommitProrata() bool`

HasMinimumCommitProrata returns a boolean if a field has been set.

### GetMinimumCommitScope

`func (o *BillableDimension) GetMinimumCommitScope() BillingMinimumCommitScope`

GetMinimumCommitScope returns the MinimumCommitScope field if non-nil, zero value otherwise.

### GetMinimumCommitScopeOk

`func (o *BillableDimension) GetMinimumCommitScopeOk() (*BillingMinimumCommitScope, bool)`

GetMinimumCommitScopeOk returns a tuple with the MinimumCommitScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCommitScope

`func (o *BillableDimension) SetMinimumCommitScope(v BillingMinimumCommitScope)`

SetMinimumCommitScope sets MinimumCommitScope field to given value.

### HasMinimumCommitScope

`func (o *BillableDimension) HasMinimumCommitScope() bool`

HasMinimumCommitScope returns a boolean if a field has been set.

### GetName

`func (o *BillableDimension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableDimension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableDimension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableDimension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceModelBasic

`func (o *BillableDimension) GetPriceModelBasic() PriceModelBasic`

GetPriceModelBasic returns the PriceModelBasic field if non-nil, zero value otherwise.

### GetPriceModelBasicOk

`func (o *BillableDimension) GetPriceModelBasicOk() (*PriceModelBasic, bool)`

GetPriceModelBasicOk returns a tuple with the PriceModelBasic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelBasic

`func (o *BillableDimension) SetPriceModelBasic(v PriceModelBasic)`

SetPriceModelBasic sets PriceModelBasic field to given value.

### HasPriceModelBasic

`func (o *BillableDimension) HasPriceModelBasic() bool`

HasPriceModelBasic returns a boolean if a field has been set.

### GetPriceModelBulk

`func (o *BillableDimension) GetPriceModelBulk() PriceModelBulk`

GetPriceModelBulk returns the PriceModelBulk field if non-nil, zero value otherwise.

### GetPriceModelBulkOk

`func (o *BillableDimension) GetPriceModelBulkOk() (*PriceModelBulk, bool)`

GetPriceModelBulkOk returns a tuple with the PriceModelBulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelBulk

`func (o *BillableDimension) SetPriceModelBulk(v PriceModelBulk)`

SetPriceModelBulk sets PriceModelBulk field to given value.

### HasPriceModelBulk

`func (o *BillableDimension) HasPriceModelBulk() bool`

HasPriceModelBulk returns a boolean if a field has been set.

### GetPriceModelMatrix

`func (o *BillableDimension) GetPriceModelMatrix() PriceModelMatrix`

GetPriceModelMatrix returns the PriceModelMatrix field if non-nil, zero value otherwise.

### GetPriceModelMatrixOk

`func (o *BillableDimension) GetPriceModelMatrixOk() (*PriceModelMatrix, bool)`

GetPriceModelMatrixOk returns a tuple with the PriceModelMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelMatrix

`func (o *BillableDimension) SetPriceModelMatrix(v PriceModelMatrix)`

SetPriceModelMatrix sets PriceModelMatrix field to given value.

### HasPriceModelMatrix

`func (o *BillableDimension) HasPriceModelMatrix() bool`

HasPriceModelMatrix returns a boolean if a field has been set.

### GetPriceModelPercentage

`func (o *BillableDimension) GetPriceModelPercentage() PriceModelPercentage`

GetPriceModelPercentage returns the PriceModelPercentage field if non-nil, zero value otherwise.

### GetPriceModelPercentageOk

`func (o *BillableDimension) GetPriceModelPercentageOk() (*PriceModelPercentage, bool)`

GetPriceModelPercentageOk returns a tuple with the PriceModelPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelPercentage

`func (o *BillableDimension) SetPriceModelPercentage(v PriceModelPercentage)`

SetPriceModelPercentage sets PriceModelPercentage field to given value.

### HasPriceModelPercentage

`func (o *BillableDimension) HasPriceModelPercentage() bool`

HasPriceModelPercentage returns a boolean if a field has been set.

### GetPriceModelTiered

`func (o *BillableDimension) GetPriceModelTiered() PriceModelTiered`

GetPriceModelTiered returns the PriceModelTiered field if non-nil, zero value otherwise.

### GetPriceModelTieredOk

`func (o *BillableDimension) GetPriceModelTieredOk() (*PriceModelTiered, bool)`

GetPriceModelTieredOk returns a tuple with the PriceModelTiered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelTiered

`func (o *BillableDimension) SetPriceModelTiered(v PriceModelTiered)`

SetPriceModelTiered sets PriceModelTiered field to given value.

### HasPriceModelTiered

`func (o *BillableDimension) HasPriceModelTiered() bool`

HasPriceModelTiered returns a boolean if a field has been set.

### GetPriceModelTieredPercentage

`func (o *BillableDimension) GetPriceModelTieredPercentage() PriceModelTieredPercentage`

GetPriceModelTieredPercentage returns the PriceModelTieredPercentage field if non-nil, zero value otherwise.

### GetPriceModelTieredPercentageOk

`func (o *BillableDimension) GetPriceModelTieredPercentageOk() (*PriceModelTieredPercentage, bool)`

GetPriceModelTieredPercentageOk returns a tuple with the PriceModelTieredPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelTieredPercentage

`func (o *BillableDimension) SetPriceModelTieredPercentage(v PriceModelTieredPercentage)`

SetPriceModelTieredPercentage sets PriceModelTieredPercentage field to given value.

### HasPriceModelTieredPercentage

`func (o *BillableDimension) HasPriceModelTieredPercentage() bool`

HasPriceModelTieredPercentage returns a boolean if a field has been set.

### GetPriceModelVolume

`func (o *BillableDimension) GetPriceModelVolume() PriceModelVolume`

GetPriceModelVolume returns the PriceModelVolume field if non-nil, zero value otherwise.

### GetPriceModelVolumeOk

`func (o *BillableDimension) GetPriceModelVolumeOk() (*PriceModelVolume, bool)`

GetPriceModelVolumeOk returns a tuple with the PriceModelVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelVolume

`func (o *BillableDimension) SetPriceModelVolume(v PriceModelVolume)`

SetPriceModelVolume sets PriceModelVolume field to given value.

### HasPriceModelVolume

`func (o *BillableDimension) HasPriceModelVolume() bool`

HasPriceModelVolume returns a boolean if a field has been set.

### GetTimeUnit

`func (o *BillableDimension) GetTimeUnit() TimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *BillableDimension) GetTimeUnitOk() (*TimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *BillableDimension) SetTimeUnit(v TimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *BillableDimension) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


