# AzureTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeDuration** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**TermUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureTerm

`func NewAzureTerm() *AzureTerm`

NewAzureTerm instantiates a new AzureTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTermWithDefaults

`func NewAzureTermWithDefaults() *AzureTerm`

NewAzureTermWithDefaults instantiates a new AzureTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeDuration

`func (o *AzureTerm) GetChargeDuration() string`

GetChargeDuration returns the ChargeDuration field if non-nil, zero value otherwise.

### GetChargeDurationOk

`func (o *AzureTerm) GetChargeDurationOk() (*string, bool)`

GetChargeDurationOk returns a tuple with the ChargeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDuration

`func (o *AzureTerm) SetChargeDuration(v string)`

SetChargeDuration sets ChargeDuration field to given value.

### HasChargeDuration

`func (o *AzureTerm) HasChargeDuration() bool`

HasChargeDuration returns a boolean if a field has been set.

### GetEndDate

`func (o *AzureTerm) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AzureTerm) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AzureTerm) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AzureTerm) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *AzureTerm) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AzureTerm) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AzureTerm) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AzureTerm) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTermUnit

`func (o *AzureTerm) GetTermUnit() string`

GetTermUnit returns the TermUnit field if non-nil, zero value otherwise.

### GetTermUnitOk

`func (o *AzureTerm) GetTermUnitOk() (*string, bool)`

GetTermUnitOk returns a tuple with the TermUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermUnit

`func (o *AzureTerm) SetTermUnit(v string)`

SetTermUnit sets TermUnit field to given value.

### HasTermUnit

`func (o *AzureTerm) HasTermUnit() bool`

HasTermUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


