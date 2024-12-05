# UsageMeteringDimensionMappingValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertionMultiplier** | Pointer to **float32** | The convertion multiplier when mapping from the source dimension key to the destination dimensionKey by quantity mode. Not required if the mapping mode is AMOUNT. | [optional] 
**DimensionKey** | Pointer to **string** | The destination dimension key of the usage metering mapping. | [optional] 
**MappingMode** | Pointer to [**UsageMeteringDimensionMappingMode**](UsageMeteringDimensionMappingMode.md) | The conversion mode of UsageMeteringDimensionMapping. The default is QUANTITY if not available. | [optional] 

## Methods

### NewUsageMeteringDimensionMappingValue

`func NewUsageMeteringDimensionMappingValue() *UsageMeteringDimensionMappingValue`

NewUsageMeteringDimensionMappingValue instantiates a new UsageMeteringDimensionMappingValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeteringDimensionMappingValueWithDefaults

`func NewUsageMeteringDimensionMappingValueWithDefaults() *UsageMeteringDimensionMappingValue`

NewUsageMeteringDimensionMappingValueWithDefaults instantiates a new UsageMeteringDimensionMappingValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertionMultiplier

`func (o *UsageMeteringDimensionMappingValue) GetConvertionMultiplier() float32`

GetConvertionMultiplier returns the ConvertionMultiplier field if non-nil, zero value otherwise.

### GetConvertionMultiplierOk

`func (o *UsageMeteringDimensionMappingValue) GetConvertionMultiplierOk() (*float32, bool)`

GetConvertionMultiplierOk returns a tuple with the ConvertionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertionMultiplier

`func (o *UsageMeteringDimensionMappingValue) SetConvertionMultiplier(v float32)`

SetConvertionMultiplier sets ConvertionMultiplier field to given value.

### HasConvertionMultiplier

`func (o *UsageMeteringDimensionMappingValue) HasConvertionMultiplier() bool`

HasConvertionMultiplier returns a boolean if a field has been set.

### GetDimensionKey

`func (o *UsageMeteringDimensionMappingValue) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *UsageMeteringDimensionMappingValue) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *UsageMeteringDimensionMappingValue) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.

### HasDimensionKey

`func (o *UsageMeteringDimensionMappingValue) HasDimensionKey() bool`

HasDimensionKey returns a boolean if a field has been set.

### GetMappingMode

`func (o *UsageMeteringDimensionMappingValue) GetMappingMode() UsageMeteringDimensionMappingMode`

GetMappingMode returns the MappingMode field if non-nil, zero value otherwise.

### GetMappingModeOk

`func (o *UsageMeteringDimensionMappingValue) GetMappingModeOk() (*UsageMeteringDimensionMappingMode, bool)`

GetMappingModeOk returns a tuple with the MappingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingMode

`func (o *UsageMeteringDimensionMappingValue) SetMappingMode(v UsageMeteringDimensionMappingMode)`

SetMappingMode sets MappingMode field to given value.

### HasMappingMode

`func (o *UsageMeteringDimensionMappingValue) HasMappingMode() bool`

HasMappingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


