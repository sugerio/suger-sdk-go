# PriceModelMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUnitAmount** | Pointer to **float32** |  | [optional] 
**Matrix** | Pointer to [**[]PriceModelMatrixConfigGroup**](PriceModelMatrixConfigGroup.md) | The matrix of the pricing model | [optional] 

## Methods

### NewPriceModelMatrix

`func NewPriceModelMatrix() *PriceModelMatrix`

NewPriceModelMatrix instantiates a new PriceModelMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceModelMatrixWithDefaults

`func NewPriceModelMatrixWithDefaults() *PriceModelMatrix`

NewPriceModelMatrixWithDefaults instantiates a new PriceModelMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUnitAmount

`func (o *PriceModelMatrix) GetDefaultUnitAmount() float32`

GetDefaultUnitAmount returns the DefaultUnitAmount field if non-nil, zero value otherwise.

### GetDefaultUnitAmountOk

`func (o *PriceModelMatrix) GetDefaultUnitAmountOk() (*float32, bool)`

GetDefaultUnitAmountOk returns a tuple with the DefaultUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnitAmount

`func (o *PriceModelMatrix) SetDefaultUnitAmount(v float32)`

SetDefaultUnitAmount sets DefaultUnitAmount field to given value.

### HasDefaultUnitAmount

`func (o *PriceModelMatrix) HasDefaultUnitAmount() bool`

HasDefaultUnitAmount returns a boolean if a field has been set.

### GetMatrix

`func (o *PriceModelMatrix) GetMatrix() []PriceModelMatrixConfigGroup`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *PriceModelMatrix) GetMatrixOk() (*[]PriceModelMatrixConfigGroup, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *PriceModelMatrix) SetMatrix(v []PriceModelMatrixConfigGroup)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *PriceModelMatrix) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


