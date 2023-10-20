# PartnerUsageMeteringConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionMapping** | Pointer to [**map[string]UsageMeteringDimensionMappingValue**](UsageMeteringDimensionMappingValue.md) | The mapping of the source dimension key to the destination dimension key of the usage metering. | [optional] 
**EnableCommitWithAdditionalUsageAtListPrice** | Pointer to **bool** | Enable the commit (discount) with additional usage metering at list price. Only applicable if EnableCommitWithAdditionalUsageMetering is true. The default is false, which means the commit with additional usage metering at the discounted price in the private offer. If set to true, the additional usage is metered at the list price (the price in public product listing) instead of the discounted price. | [optional] 
**EnableCommitWithAdditionalUsageMetering** | Pointer to **bool** | Enable the commit with additional usage metering. The default is false, which means all usage records are reported to partner no matter how much is the commit. If set to true, the usage records will be reported to partner only if the current commit has been exhausted. | [optional] 
**EnableDimensionMapping** | Pointer to **bool** | Enable the dimension mapping for the usage metering. The default is false, which means no dimension conversion and just use the origin dimension. | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 

## Methods

### NewPartnerUsageMeteringConfig

`func NewPartnerUsageMeteringConfig() *PartnerUsageMeteringConfig`

NewPartnerUsageMeteringConfig instantiates a new PartnerUsageMeteringConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerUsageMeteringConfigWithDefaults

`func NewPartnerUsageMeteringConfigWithDefaults() *PartnerUsageMeteringConfig`

NewPartnerUsageMeteringConfigWithDefaults instantiates a new PartnerUsageMeteringConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionMapping

`func (o *PartnerUsageMeteringConfig) GetDimensionMapping() map[string]UsageMeteringDimensionMappingValue`

GetDimensionMapping returns the DimensionMapping field if non-nil, zero value otherwise.

### GetDimensionMappingOk

`func (o *PartnerUsageMeteringConfig) GetDimensionMappingOk() (*map[string]UsageMeteringDimensionMappingValue, bool)`

GetDimensionMappingOk returns a tuple with the DimensionMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionMapping

`func (o *PartnerUsageMeteringConfig) SetDimensionMapping(v map[string]UsageMeteringDimensionMappingValue)`

SetDimensionMapping sets DimensionMapping field to given value.

### HasDimensionMapping

`func (o *PartnerUsageMeteringConfig) HasDimensionMapping() bool`

HasDimensionMapping returns a boolean if a field has been set.

### GetEnableCommitWithAdditionalUsageAtListPrice

`func (o *PartnerUsageMeteringConfig) GetEnableCommitWithAdditionalUsageAtListPrice() bool`

GetEnableCommitWithAdditionalUsageAtListPrice returns the EnableCommitWithAdditionalUsageAtListPrice field if non-nil, zero value otherwise.

### GetEnableCommitWithAdditionalUsageAtListPriceOk

`func (o *PartnerUsageMeteringConfig) GetEnableCommitWithAdditionalUsageAtListPriceOk() (*bool, bool)`

GetEnableCommitWithAdditionalUsageAtListPriceOk returns a tuple with the EnableCommitWithAdditionalUsageAtListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCommitWithAdditionalUsageAtListPrice

`func (o *PartnerUsageMeteringConfig) SetEnableCommitWithAdditionalUsageAtListPrice(v bool)`

SetEnableCommitWithAdditionalUsageAtListPrice sets EnableCommitWithAdditionalUsageAtListPrice field to given value.

### HasEnableCommitWithAdditionalUsageAtListPrice

`func (o *PartnerUsageMeteringConfig) HasEnableCommitWithAdditionalUsageAtListPrice() bool`

HasEnableCommitWithAdditionalUsageAtListPrice returns a boolean if a field has been set.

### GetEnableCommitWithAdditionalUsageMetering

`func (o *PartnerUsageMeteringConfig) GetEnableCommitWithAdditionalUsageMetering() bool`

GetEnableCommitWithAdditionalUsageMetering returns the EnableCommitWithAdditionalUsageMetering field if non-nil, zero value otherwise.

### GetEnableCommitWithAdditionalUsageMeteringOk

`func (o *PartnerUsageMeteringConfig) GetEnableCommitWithAdditionalUsageMeteringOk() (*bool, bool)`

GetEnableCommitWithAdditionalUsageMeteringOk returns a tuple with the EnableCommitWithAdditionalUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCommitWithAdditionalUsageMetering

`func (o *PartnerUsageMeteringConfig) SetEnableCommitWithAdditionalUsageMetering(v bool)`

SetEnableCommitWithAdditionalUsageMetering sets EnableCommitWithAdditionalUsageMetering field to given value.

### HasEnableCommitWithAdditionalUsageMetering

`func (o *PartnerUsageMeteringConfig) HasEnableCommitWithAdditionalUsageMetering() bool`

HasEnableCommitWithAdditionalUsageMetering returns a boolean if a field has been set.

### GetEnableDimensionMapping

`func (o *PartnerUsageMeteringConfig) GetEnableDimensionMapping() bool`

GetEnableDimensionMapping returns the EnableDimensionMapping field if non-nil, zero value otherwise.

### GetEnableDimensionMappingOk

`func (o *PartnerUsageMeteringConfig) GetEnableDimensionMappingOk() (*bool, bool)`

GetEnableDimensionMappingOk returns a tuple with the EnableDimensionMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDimensionMapping

`func (o *PartnerUsageMeteringConfig) SetEnableDimensionMapping(v bool)`

SetEnableDimensionMapping sets EnableDimensionMapping field to given value.

### HasEnableDimensionMapping

`func (o *PartnerUsageMeteringConfig) HasEnableDimensionMapping() bool`

HasEnableDimensionMapping returns a boolean if a field has been set.

### GetPartner

`func (o *PartnerUsageMeteringConfig) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *PartnerUsageMeteringConfig) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *PartnerUsageMeteringConfig) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *PartnerUsageMeteringConfig) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


