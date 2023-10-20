# AwsMarketplaceCppoPriceTermEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionUnitColumnNames** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | the dimension display name | [optional] 
**IsCustomDimension** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The dimension Key | [optional] 
**PricePerConsumptionUnit** | Pointer to **map[string]string** | Key: the unit in ConsumptionUnitColumnName, Value: the unit price | [optional] 
**PricingDimension** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsMarketplaceCppoPriceTermEntry

`func NewAwsMarketplaceCppoPriceTermEntry() *AwsMarketplaceCppoPriceTermEntry`

NewAwsMarketplaceCppoPriceTermEntry instantiates a new AwsMarketplaceCppoPriceTermEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCppoPriceTermEntryWithDefaults

`func NewAwsMarketplaceCppoPriceTermEntryWithDefaults() *AwsMarketplaceCppoPriceTermEntry`

NewAwsMarketplaceCppoPriceTermEntryWithDefaults instantiates a new AwsMarketplaceCppoPriceTermEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumptionUnitColumnNames

`func (o *AwsMarketplaceCppoPriceTermEntry) GetConsumptionUnitColumnNames() []string`

GetConsumptionUnitColumnNames returns the ConsumptionUnitColumnNames field if non-nil, zero value otherwise.

### GetConsumptionUnitColumnNamesOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetConsumptionUnitColumnNamesOk() (*[]string, bool)`

GetConsumptionUnitColumnNamesOk returns a tuple with the ConsumptionUnitColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionUnitColumnNames

`func (o *AwsMarketplaceCppoPriceTermEntry) SetConsumptionUnitColumnNames(v []string)`

SetConsumptionUnitColumnNames sets ConsumptionUnitColumnNames field to given value.

### HasConsumptionUnitColumnNames

`func (o *AwsMarketplaceCppoPriceTermEntry) HasConsumptionUnitColumnNames() bool`

HasConsumptionUnitColumnNames returns a boolean if a field has been set.

### GetDescription

`func (o *AwsMarketplaceCppoPriceTermEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsMarketplaceCppoPriceTermEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsMarketplaceCppoPriceTermEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AwsMarketplaceCppoPriceTermEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AwsMarketplaceCppoPriceTermEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AwsMarketplaceCppoPriceTermEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsCustomDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) GetIsCustomDimension() bool`

GetIsCustomDimension returns the IsCustomDimension field if non-nil, zero value otherwise.

### GetIsCustomDimensionOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetIsCustomDimensionOk() (*bool, bool)`

GetIsCustomDimensionOk returns a tuple with the IsCustomDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) SetIsCustomDimension(v bool)`

SetIsCustomDimension sets IsCustomDimension field to given value.

### HasIsCustomDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) HasIsCustomDimension() bool`

HasIsCustomDimension returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AwsMarketplaceCppoPriceTermEntry) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AwsMarketplaceCppoPriceTermEntry) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AwsMarketplaceCppoPriceTermEntry) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *AwsMarketplaceCppoPriceTermEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsMarketplaceCppoPriceTermEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsMarketplaceCppoPriceTermEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricePerConsumptionUnit

`func (o *AwsMarketplaceCppoPriceTermEntry) GetPricePerConsumptionUnit() map[string]string`

GetPricePerConsumptionUnit returns the PricePerConsumptionUnit field if non-nil, zero value otherwise.

### GetPricePerConsumptionUnitOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetPricePerConsumptionUnitOk() (*map[string]string, bool)`

GetPricePerConsumptionUnitOk returns a tuple with the PricePerConsumptionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerConsumptionUnit

`func (o *AwsMarketplaceCppoPriceTermEntry) SetPricePerConsumptionUnit(v map[string]string)`

SetPricePerConsumptionUnit sets PricePerConsumptionUnit field to given value.

### HasPricePerConsumptionUnit

`func (o *AwsMarketplaceCppoPriceTermEntry) HasPricePerConsumptionUnit() bool`

HasPricePerConsumptionUnit returns a boolean if a field has been set.

### GetPricingDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) GetPricingDimension() string`

GetPricingDimension returns the PricingDimension field if non-nil, zero value otherwise.

### GetPricingDimensionOk

`func (o *AwsMarketplaceCppoPriceTermEntry) GetPricingDimensionOk() (*string, bool)`

GetPricingDimensionOk returns a tuple with the PricingDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) SetPricingDimension(v string)`

SetPricingDimension sets PricingDimension field to given value.

### HasPricingDimension

`func (o *AwsMarketplaceCppoPriceTermEntry) HasPricingDimension() bool`

HasPricingDimension returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


