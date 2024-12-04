# TrialConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialPeriod** | Pointer to **int32** |  | [optional] 
**TrialPeriodUnit** | Pointer to [**TimeUnit**](TimeUnit.md) |  | [optional] 

## Methods

### NewTrialConfig

`func NewTrialConfig() *TrialConfig`

NewTrialConfig instantiates a new TrialConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialConfigWithDefaults

`func NewTrialConfigWithDefaults() *TrialConfig`

NewTrialConfigWithDefaults instantiates a new TrialConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialPeriod

`func (o *TrialConfig) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *TrialConfig) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *TrialConfig) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *TrialConfig) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetTrialPeriodUnit

`func (o *TrialConfig) GetTrialPeriodUnit() TimeUnit`

GetTrialPeriodUnit returns the TrialPeriodUnit field if non-nil, zero value otherwise.

### GetTrialPeriodUnitOk

`func (o *TrialConfig) GetTrialPeriodUnitOk() (*TimeUnit, bool)`

GetTrialPeriodUnitOk returns a tuple with the TrialPeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodUnit

`func (o *TrialConfig) SetTrialPeriodUnit(v TimeUnit)`

SetTrialPeriodUnit sets TrialPeriodUnit field to given value.

### HasTrialPeriodUnit

`func (o *TrialConfig) HasTrialPeriodUnit() bool`

HasTrialPeriodUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


