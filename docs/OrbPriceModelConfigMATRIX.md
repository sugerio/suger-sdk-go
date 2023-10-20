# OrbPriceModelConfigMATRIX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUnitAmount** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to **[]string** | First and (optional) second dimension grouping keys | [optional] 
**MatrixValues** | Pointer to [**[]OrbMatrixPriceValue**](OrbMatrixPriceValue.md) |  | [optional] 

## Methods

### NewOrbPriceModelConfigMATRIX

`func NewOrbPriceModelConfigMATRIX() *OrbPriceModelConfigMATRIX`

NewOrbPriceModelConfigMATRIX instantiates a new OrbPriceModelConfigMATRIX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceModelConfigMATRIXWithDefaults

`func NewOrbPriceModelConfigMATRIXWithDefaults() *OrbPriceModelConfigMATRIX`

NewOrbPriceModelConfigMATRIXWithDefaults instantiates a new OrbPriceModelConfigMATRIX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUnitAmount

`func (o *OrbPriceModelConfigMATRIX) GetDefaultUnitAmount() string`

GetDefaultUnitAmount returns the DefaultUnitAmount field if non-nil, zero value otherwise.

### GetDefaultUnitAmountOk

`func (o *OrbPriceModelConfigMATRIX) GetDefaultUnitAmountOk() (*string, bool)`

GetDefaultUnitAmountOk returns a tuple with the DefaultUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnitAmount

`func (o *OrbPriceModelConfigMATRIX) SetDefaultUnitAmount(v string)`

SetDefaultUnitAmount sets DefaultUnitAmount field to given value.

### HasDefaultUnitAmount

`func (o *OrbPriceModelConfigMATRIX) HasDefaultUnitAmount() bool`

HasDefaultUnitAmount returns a boolean if a field has been set.

### GetDimensions

`func (o *OrbPriceModelConfigMATRIX) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OrbPriceModelConfigMATRIX) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OrbPriceModelConfigMATRIX) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *OrbPriceModelConfigMATRIX) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetMatrixValues

`func (o *OrbPriceModelConfigMATRIX) GetMatrixValues() []OrbMatrixPriceValue`

GetMatrixValues returns the MatrixValues field if non-nil, zero value otherwise.

### GetMatrixValuesOk

`func (o *OrbPriceModelConfigMATRIX) GetMatrixValuesOk() (*[]OrbMatrixPriceValue, bool)`

GetMatrixValuesOk returns a tuple with the MatrixValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrixValues

`func (o *OrbPriceModelConfigMATRIX) SetMatrixValues(v []OrbMatrixPriceValue)`

SetMatrixValues sets MatrixValues field to given value.

### HasMatrixValues

`func (o *OrbPriceModelConfigMATRIX) HasMatrixValues() bool`

HasMatrixValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


