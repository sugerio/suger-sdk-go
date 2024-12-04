# NewUsageRecordGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillableRecords** | Pointer to [**[]MeteringUsageRecord**](MeteringUsageRecord.md) | for usage metering API v2, don&#39;t use it together with the records v1. | [optional] 
**EntitlementID** | **string** |  | 
**MetaInfo** | Pointer to [**MeteringUsageRecordGroupMetaInfo**](MeteringUsageRecordGroupMetaInfo.md) | read-only, don&#39;t set it when validating or reporting the usage record group. | [optional] 
**Records** | **map[string]float32** | for usage metering API v1, don&#39;t use it together with the billableRecords v2. | 
**Timestamp** | Pointer to **time.Time** | The timestamp of when the usage records were generated. Optional, if not provided, the current report timestamp will be used. This is not the timestamp of when the usage records were reported to Suger. | [optional] 

## Methods

### NewNewUsageRecordGroup

`func NewNewUsageRecordGroup(entitlementID string, records map[string]float32, ) *NewUsageRecordGroup`

NewNewUsageRecordGroup instantiates a new NewUsageRecordGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUsageRecordGroupWithDefaults

`func NewNewUsageRecordGroupWithDefaults() *NewUsageRecordGroup`

NewNewUsageRecordGroupWithDefaults instantiates a new NewUsageRecordGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillableRecords

`func (o *NewUsageRecordGroup) GetBillableRecords() []MeteringUsageRecord`

GetBillableRecords returns the BillableRecords field if non-nil, zero value otherwise.

### GetBillableRecordsOk

`func (o *NewUsageRecordGroup) GetBillableRecordsOk() (*[]MeteringUsageRecord, bool)`

GetBillableRecordsOk returns a tuple with the BillableRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableRecords

`func (o *NewUsageRecordGroup) SetBillableRecords(v []MeteringUsageRecord)`

SetBillableRecords sets BillableRecords field to given value.

### HasBillableRecords

`func (o *NewUsageRecordGroup) HasBillableRecords() bool`

HasBillableRecords returns a boolean if a field has been set.

### GetEntitlementID

`func (o *NewUsageRecordGroup) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *NewUsageRecordGroup) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *NewUsageRecordGroup) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.


### GetMetaInfo

`func (o *NewUsageRecordGroup) GetMetaInfo() MeteringUsageRecordGroupMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *NewUsageRecordGroup) GetMetaInfoOk() (*MeteringUsageRecordGroupMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *NewUsageRecordGroup) SetMetaInfo(v MeteringUsageRecordGroupMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *NewUsageRecordGroup) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetRecords

`func (o *NewUsageRecordGroup) GetRecords() map[string]float32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *NewUsageRecordGroup) GetRecordsOk() (*map[string]float32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *NewUsageRecordGroup) SetRecords(v map[string]float32)`

SetRecords sets Records field to given value.


### GetTimestamp

`func (o *NewUsageRecordGroup) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NewUsageRecordGroup) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NewUsageRecordGroup) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NewUsageRecordGroup) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


