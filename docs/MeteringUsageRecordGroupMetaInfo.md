# MeteringUsageRecordGroupMetaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginRecords** | Pointer to **map[string]float32** | The original records reported by the customer before convertion. If no dimension mapping is applied, this field is the same as the records field. | [optional] 
**Source** | Pointer to [**UsageRecordGroupSource**](UsageRecordGroupSource.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** | The timestamp (UTC)) of when the usage records were generated. Optional, if not provided, the current report timestamp will be used. | [optional] 

## Methods

### NewMeteringUsageRecordGroupMetaInfo

`func NewMeteringUsageRecordGroupMetaInfo() *MeteringUsageRecordGroupMetaInfo`

NewMeteringUsageRecordGroupMetaInfo instantiates a new MeteringUsageRecordGroupMetaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordGroupMetaInfoWithDefaults

`func NewMeteringUsageRecordGroupMetaInfoWithDefaults() *MeteringUsageRecordGroupMetaInfo`

NewMeteringUsageRecordGroupMetaInfoWithDefaults instantiates a new MeteringUsageRecordGroupMetaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecords() map[string]float32`

GetOriginRecords returns the OriginRecords field if non-nil, zero value otherwise.

### GetOriginRecordsOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecordsOk() (*map[string]float32, bool)`

GetOriginRecordsOk returns a tuple with the OriginRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) SetOriginRecords(v map[string]float32)`

SetOriginRecords sets OriginRecords field to given value.

### HasOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) HasOriginRecords() bool`

HasOriginRecords returns a boolean if a field has been set.

### GetSource

`func (o *MeteringUsageRecordGroupMetaInfo) GetSource() UsageRecordGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetSourceOk() (*UsageRecordGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MeteringUsageRecordGroupMetaInfo) SetSource(v UsageRecordGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MeteringUsageRecordGroupMetaInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


