# PriceModelTieredConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstUnit** | Pointer to **float32** | Inclusive tier starting value | [optional] 
**FlatFee** | Pointer to **float32** |  | [optional] 
**LastUnit** | Pointer to **float32** | Exclusive tier ending value. If null, this is treated as the last tier | [optional] 
**UnitAmount** | Pointer to **float32** | Amount per unit | [optional] 

## Methods

### NewPriceModelTieredConfig

`func NewPriceModelTieredConfig() *PriceModelTieredConfig`

NewPriceModelTieredConfig instantiates a new PriceModelTieredConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceModelTieredConfigWithDefaults

`func NewPriceModelTieredConfigWithDefaults() *PriceModelTieredConfig`

NewPriceModelTieredConfigWithDefaults instantiates a new PriceModelTieredConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstUnit

`func (o *PriceModelTieredConfig) GetFirstUnit() float32`

GetFirstUnit returns the FirstUnit field if non-nil, zero value otherwise.

### GetFirstUnitOk

`func (o *PriceModelTieredConfig) GetFirstUnitOk() (*float32, bool)`

GetFirstUnitOk returns a tuple with the FirstUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnit

`func (o *PriceModelTieredConfig) SetFirstUnit(v float32)`

SetFirstUnit sets FirstUnit field to given value.

### HasFirstUnit

`func (o *PriceModelTieredConfig) HasFirstUnit() bool`

HasFirstUnit returns a boolean if a field has been set.

### GetFlatFee

`func (o *PriceModelTieredConfig) GetFlatFee() float32`

GetFlatFee returns the FlatFee field if non-nil, zero value otherwise.

### GetFlatFeeOk

`func (o *PriceModelTieredConfig) GetFlatFeeOk() (*float32, bool)`

GetFlatFeeOk returns a tuple with the FlatFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatFee

`func (o *PriceModelTieredConfig) SetFlatFee(v float32)`

SetFlatFee sets FlatFee field to given value.

### HasFlatFee

`func (o *PriceModelTieredConfig) HasFlatFee() bool`

HasFlatFee returns a boolean if a field has been set.

### GetLastUnit

`func (o *PriceModelTieredConfig) GetLastUnit() float32`

GetLastUnit returns the LastUnit field if non-nil, zero value otherwise.

### GetLastUnitOk

`func (o *PriceModelTieredConfig) GetLastUnitOk() (*float32, bool)`

GetLastUnitOk returns a tuple with the LastUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUnit

`func (o *PriceModelTieredConfig) SetLastUnit(v float32)`

SetLastUnit sets LastUnit field to given value.

### HasLastUnit

`func (o *PriceModelTieredConfig) HasLastUnit() bool`

HasLastUnit returns a boolean if a field has been set.

### GetUnitAmount

`func (o *PriceModelTieredConfig) GetUnitAmount() float32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PriceModelTieredConfig) GetUnitAmountOk() (*float32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PriceModelTieredConfig) SetUnitAmount(v float32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *PriceModelTieredConfig) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


