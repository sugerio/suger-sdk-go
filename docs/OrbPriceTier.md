# OrbPriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bps** | Pointer to **float32** |  | [optional] 
**FirstUnit** | Pointer to **string** | The following fields applicable only to UNIT price model | [optional] 
**LastUnit** | Pointer to **string** |  | [optional] 
**MaximumAmount** | Pointer to **string** |  | [optional] 
**MaximumUnits** | Pointer to **float32** | The following fields applicable only to BULK price model | [optional] 
**MinimumAmount** | Pointer to **string** | The following fields applicable only to BPS price model | [optional] 
**PerUnitMaximum** | Pointer to **string** |  | [optional] 
**UnitAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewOrbPriceTier

`func NewOrbPriceTier() *OrbPriceTier`

NewOrbPriceTier instantiates a new OrbPriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceTierWithDefaults

`func NewOrbPriceTierWithDefaults() *OrbPriceTier`

NewOrbPriceTierWithDefaults instantiates a new OrbPriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBps

`func (o *OrbPriceTier) GetBps() float32`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *OrbPriceTier) GetBpsOk() (*float32, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *OrbPriceTier) SetBps(v float32)`

SetBps sets Bps field to given value.

### HasBps

`func (o *OrbPriceTier) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetFirstUnit

`func (o *OrbPriceTier) GetFirstUnit() string`

GetFirstUnit returns the FirstUnit field if non-nil, zero value otherwise.

### GetFirstUnitOk

`func (o *OrbPriceTier) GetFirstUnitOk() (*string, bool)`

GetFirstUnitOk returns a tuple with the FirstUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnit

`func (o *OrbPriceTier) SetFirstUnit(v string)`

SetFirstUnit sets FirstUnit field to given value.

### HasFirstUnit

`func (o *OrbPriceTier) HasFirstUnit() bool`

HasFirstUnit returns a boolean if a field has been set.

### GetLastUnit

`func (o *OrbPriceTier) GetLastUnit() string`

GetLastUnit returns the LastUnit field if non-nil, zero value otherwise.

### GetLastUnitOk

`func (o *OrbPriceTier) GetLastUnitOk() (*string, bool)`

GetLastUnitOk returns a tuple with the LastUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUnit

`func (o *OrbPriceTier) SetLastUnit(v string)`

SetLastUnit sets LastUnit field to given value.

### HasLastUnit

`func (o *OrbPriceTier) HasLastUnit() bool`

HasLastUnit returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *OrbPriceTier) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *OrbPriceTier) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *OrbPriceTier) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *OrbPriceTier) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### GetMaximumUnits

`func (o *OrbPriceTier) GetMaximumUnits() float32`

GetMaximumUnits returns the MaximumUnits field if non-nil, zero value otherwise.

### GetMaximumUnitsOk

`func (o *OrbPriceTier) GetMaximumUnitsOk() (*float32, bool)`

GetMaximumUnitsOk returns a tuple with the MaximumUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumUnits

`func (o *OrbPriceTier) SetMaximumUnits(v float32)`

SetMaximumUnits sets MaximumUnits field to given value.

### HasMaximumUnits

`func (o *OrbPriceTier) HasMaximumUnits() bool`

HasMaximumUnits returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *OrbPriceTier) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *OrbPriceTier) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *OrbPriceTier) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *OrbPriceTier) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetPerUnitMaximum

`func (o *OrbPriceTier) GetPerUnitMaximum() string`

GetPerUnitMaximum returns the PerUnitMaximum field if non-nil, zero value otherwise.

### GetPerUnitMaximumOk

`func (o *OrbPriceTier) GetPerUnitMaximumOk() (*string, bool)`

GetPerUnitMaximumOk returns a tuple with the PerUnitMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnitMaximum

`func (o *OrbPriceTier) SetPerUnitMaximum(v string)`

SetPerUnitMaximum sets PerUnitMaximum field to given value.

### HasPerUnitMaximum

`func (o *OrbPriceTier) HasPerUnitMaximum() bool`

HasPerUnitMaximum returns a boolean if a field has been set.

### GetUnitAmount

`func (o *OrbPriceTier) GetUnitAmount() string`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *OrbPriceTier) GetUnitAmountOk() (*string, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *OrbPriceTier) SetUnitAmount(v string)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *OrbPriceTier) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


