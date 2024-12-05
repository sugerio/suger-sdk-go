# CancellationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelDate** | Pointer to **string** | The cancellation date of the entitlement. It is required when the type is SpecificDate. | [optional] 
**CreationDate** | Pointer to **string** | When this cancellation schedule is created. | [optional] 
**Note** | Pointer to **string** | The cancellation note. Max 500 characters. | [optional] 
**Type** | Pointer to [**CancellationScheduleType**](CancellationScheduleType.md) | Cancellation type | [optional] 

## Methods

### NewCancellationSchedule

`func NewCancellationSchedule() *CancellationSchedule`

NewCancellationSchedule instantiates a new CancellationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationScheduleWithDefaults

`func NewCancellationScheduleWithDefaults() *CancellationSchedule`

NewCancellationScheduleWithDefaults instantiates a new CancellationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelDate

`func (o *CancellationSchedule) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *CancellationSchedule) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *CancellationSchedule) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *CancellationSchedule) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *CancellationSchedule) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CancellationSchedule) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CancellationSchedule) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CancellationSchedule) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetNote

`func (o *CancellationSchedule) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CancellationSchedule) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CancellationSchedule) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CancellationSchedule) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetType

`func (o *CancellationSchedule) GetType() CancellationScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancellationSchedule) GetTypeOk() (*CancellationScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancellationSchedule) SetType(v CancellationScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *CancellationSchedule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


