# OrbPlanPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**OrbPriceDiscount**](OrbPriceDiscount.md) |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**DurationUnit** | Pointer to [**OrbCadence**](OrbCadence.md) |  | [optional] 
**Maximum** | Pointer to [**OrbPriceMaximum**](OrbPriceMaximum.md) |  | [optional] 
**Minimum** | Pointer to [**OrbPriceMinimum**](OrbPriceMinimum.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrbPlanPhase

`func NewOrbPlanPhase() *OrbPlanPhase`

NewOrbPlanPhase instantiates a new OrbPlanPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPlanPhaseWithDefaults

`func NewOrbPlanPhaseWithDefaults() *OrbPlanPhase`

NewOrbPlanPhaseWithDefaults instantiates a new OrbPlanPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrbPlanPhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrbPlanPhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrbPlanPhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrbPlanPhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscount

`func (o *OrbPlanPhase) GetDiscount() OrbPriceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrbPlanPhase) GetDiscountOk() (*OrbPriceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrbPlanPhase) SetDiscount(v OrbPriceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrbPlanPhase) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDuration

`func (o *OrbPlanPhase) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrbPlanPhase) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrbPlanPhase) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrbPlanPhase) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDurationUnit

`func (o *OrbPlanPhase) GetDurationUnit() OrbCadence`

GetDurationUnit returns the DurationUnit field if non-nil, zero value otherwise.

### GetDurationUnitOk

`func (o *OrbPlanPhase) GetDurationUnitOk() (*OrbCadence, bool)`

GetDurationUnitOk returns a tuple with the DurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnit

`func (o *OrbPlanPhase) SetDurationUnit(v OrbCadence)`

SetDurationUnit sets DurationUnit field to given value.

### HasDurationUnit

`func (o *OrbPlanPhase) HasDurationUnit() bool`

HasDurationUnit returns a boolean if a field has been set.

### GetMaximum

`func (o *OrbPlanPhase) GetMaximum() OrbPriceMaximum`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *OrbPlanPhase) GetMaximumOk() (*OrbPriceMaximum, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *OrbPlanPhase) SetMaximum(v OrbPriceMaximum)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *OrbPlanPhase) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *OrbPlanPhase) GetMinimum() OrbPriceMinimum`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OrbPlanPhase) GetMinimumOk() (*OrbPriceMinimum, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OrbPlanPhase) SetMinimum(v OrbPriceMinimum)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *OrbPlanPhase) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetName

`func (o *OrbPlanPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrbPlanPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrbPlanPhase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrbPlanPhase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *OrbPlanPhase) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrbPlanPhase) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrbPlanPhase) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrbPlanPhase) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


