# CreateUsageRecordGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementID** | **string** |  | 
**Id** | Pointer to **string** | The uuid of the UsageRecordGroup (the size is up to 36 characters). Optional, if not provided, suger will generate one. | [optional] 
**MetaInfo** | Pointer to [**MeteringUsageRecordGroupMetaInfo**](MeteringUsageRecordGroupMetaInfo.md) |  | [optional] 
**OrganizationID** | **string** |  | 
**Records** | **map[string]float32** |  | 
**Timestamp** | Pointer to **time.Time** | The timestamp of when the usage records were generated. Optional, if not provided, the current report timestamp will be used. This is not the timestamp of when the usage records were reported to Suger. | [optional] 

## Methods

### NewCreateUsageRecordGroupParams

`func NewCreateUsageRecordGroupParams(entitlementID string, organizationID string, records map[string]float32, ) *CreateUsageRecordGroupParams`

NewCreateUsageRecordGroupParams instantiates a new CreateUsageRecordGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsageRecordGroupParamsWithDefaults

`func NewCreateUsageRecordGroupParamsWithDefaults() *CreateUsageRecordGroupParams`

NewCreateUsageRecordGroupParamsWithDefaults instantiates a new CreateUsageRecordGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementID

`func (o *CreateUsageRecordGroupParams) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *CreateUsageRecordGroupParams) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *CreateUsageRecordGroupParams) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.


### GetId

`func (o *CreateUsageRecordGroupParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateUsageRecordGroupParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateUsageRecordGroupParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateUsageRecordGroupParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaInfo

`func (o *CreateUsageRecordGroupParams) GetMetaInfo() MeteringUsageRecordGroupMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *CreateUsageRecordGroupParams) GetMetaInfoOk() (*MeteringUsageRecordGroupMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *CreateUsageRecordGroupParams) SetMetaInfo(v MeteringUsageRecordGroupMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *CreateUsageRecordGroupParams) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetOrganizationID

`func (o *CreateUsageRecordGroupParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *CreateUsageRecordGroupParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *CreateUsageRecordGroupParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.


### GetRecords

`func (o *CreateUsageRecordGroupParams) GetRecords() map[string]float32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *CreateUsageRecordGroupParams) GetRecordsOk() (*map[string]float32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *CreateUsageRecordGroupParams) SetRecords(v map[string]float32)`

SetRecords sets Records field to given value.


### GetTimestamp

`func (o *CreateUsageRecordGroupParams) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CreateUsageRecordGroupParams) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CreateUsageRecordGroupParams) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CreateUsageRecordGroupParams) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


