# TypesUsageRecordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeteringRecordId** | Pointer to **string** | The MeteringRecordId is a unique identifier for this metering event. | [optional] 
**Status** | Pointer to [**TypesUsageRecordResultStatus**](TypesUsageRecordResultStatus.md) |  | [optional] 
**UsageRecord** | Pointer to [**TypesUsageRecord**](TypesUsageRecord.md) |  | [optional] 

## Methods

### NewTypesUsageRecordResult

`func NewTypesUsageRecordResult() *TypesUsageRecordResult`

NewTypesUsageRecordResult instantiates a new TypesUsageRecordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesUsageRecordResultWithDefaults

`func NewTypesUsageRecordResultWithDefaults() *TypesUsageRecordResult`

NewTypesUsageRecordResultWithDefaults instantiates a new TypesUsageRecordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeteringRecordId

`func (o *TypesUsageRecordResult) GetMeteringRecordId() string`

GetMeteringRecordId returns the MeteringRecordId field if non-nil, zero value otherwise.

### GetMeteringRecordIdOk

`func (o *TypesUsageRecordResult) GetMeteringRecordIdOk() (*string, bool)`

GetMeteringRecordIdOk returns a tuple with the MeteringRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringRecordId

`func (o *TypesUsageRecordResult) SetMeteringRecordId(v string)`

SetMeteringRecordId sets MeteringRecordId field to given value.

### HasMeteringRecordId

`func (o *TypesUsageRecordResult) HasMeteringRecordId() bool`

HasMeteringRecordId returns a boolean if a field has been set.

### GetStatus

`func (o *TypesUsageRecordResult) GetStatus() TypesUsageRecordResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesUsageRecordResult) GetStatusOk() (*TypesUsageRecordResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesUsageRecordResult) SetStatus(v TypesUsageRecordResultStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesUsageRecordResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsageRecord

`func (o *TypesUsageRecordResult) GetUsageRecord() TypesUsageRecord`

GetUsageRecord returns the UsageRecord field if non-nil, zero value otherwise.

### GetUsageRecordOk

`func (o *TypesUsageRecordResult) GetUsageRecordOk() (*TypesUsageRecord, bool)`

GetUsageRecordOk returns a tuple with the UsageRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecord

`func (o *TypesUsageRecordResult) SetUsageRecord(v TypesUsageRecord)`

SetUsageRecord sets UsageRecord field to given value.

### HasUsageRecord

`func (o *TypesUsageRecordResult) HasUsageRecord() bool`

HasUsageRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


