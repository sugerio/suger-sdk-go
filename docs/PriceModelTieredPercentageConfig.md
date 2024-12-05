# PriceModelTieredPercentageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstUnit** | Pointer to **float32** | Inclusive tier starting value | [optional] 
**FlatFee** | Pointer to **float32** |  | [optional] 
**LastUnit** | Pointer to **float32** | Exclusive tier ending value. If null, this is treated as the last tier | [optional] 
**PercentageRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewPriceModelTieredPercentageConfig

`func NewPriceModelTieredPercentageConfig() *PriceModelTieredPercentageConfig`

NewPriceModelTieredPercentageConfig instantiates a new PriceModelTieredPercentageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceModelTieredPercentageConfigWithDefaults

`func NewPriceModelTieredPercentageConfigWithDefaults() *PriceModelTieredPercentageConfig`

NewPriceModelTieredPercentageConfigWithDefaults instantiates a new PriceModelTieredPercentageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstUnit

`func (o *PriceModelTieredPercentageConfig) GetFirstUnit() float32`

GetFirstUnit returns the FirstUnit field if non-nil, zero value otherwise.

### GetFirstUnitOk

`func (o *PriceModelTieredPercentageConfig) GetFirstUnitOk() (*float32, bool)`

GetFirstUnitOk returns a tuple with the FirstUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnit

`func (o *PriceModelTieredPercentageConfig) SetFirstUnit(v float32)`

SetFirstUnit sets FirstUnit field to given value.

### HasFirstUnit

`func (o *PriceModelTieredPercentageConfig) HasFirstUnit() bool`

HasFirstUnit returns a boolean if a field has been set.

### GetFlatFee

`func (o *PriceModelTieredPercentageConfig) GetFlatFee() float32`

GetFlatFee returns the FlatFee field if non-nil, zero value otherwise.

### GetFlatFeeOk

`func (o *PriceModelTieredPercentageConfig) GetFlatFeeOk() (*float32, bool)`

GetFlatFeeOk returns a tuple with the FlatFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatFee

`func (o *PriceModelTieredPercentageConfig) SetFlatFee(v float32)`

SetFlatFee sets FlatFee field to given value.

### HasFlatFee

`func (o *PriceModelTieredPercentageConfig) HasFlatFee() bool`

HasFlatFee returns a boolean if a field has been set.

### GetLastUnit

`func (o *PriceModelTieredPercentageConfig) GetLastUnit() float32`

GetLastUnit returns the LastUnit field if non-nil, zero value otherwise.

### GetLastUnitOk

`func (o *PriceModelTieredPercentageConfig) GetLastUnitOk() (*float32, bool)`

GetLastUnitOk returns a tuple with the LastUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUnit

`func (o *PriceModelTieredPercentageConfig) SetLastUnit(v float32)`

SetLastUnit sets LastUnit field to given value.

### HasLastUnit

`func (o *PriceModelTieredPercentageConfig) HasLastUnit() bool`

HasLastUnit returns a boolean if a field has been set.

### GetPercentageRate

`func (o *PriceModelTieredPercentageConfig) GetPercentageRate() float32`

GetPercentageRate returns the PercentageRate field if non-nil, zero value otherwise.

### GetPercentageRateOk

`func (o *PriceModelTieredPercentageConfig) GetPercentageRateOk() (*float32, bool)`

GetPercentageRateOk returns a tuple with the PercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRate

`func (o *PriceModelTieredPercentageConfig) SetPercentageRate(v float32)`

SetPercentageRate sets PercentageRate field to given value.

### HasPercentageRate

`func (o *PriceModelTieredPercentageConfig) HasPercentageRate() bool`

HasPercentageRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


