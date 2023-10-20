# MeteringUsageRecordGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**MetaInfo** | Pointer to [**MeteringUsageRecordGroupMetaInfo**](MeteringUsageRecordGroupMetaInfo.md) |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to **string** |  | [optional] 
**Records** | Pointer to **map[string]float32** |  | [optional] 
**SerialID** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UsageRecordReportID** | Pointer to **string** |  | [optional] 

## Methods

### NewMeteringUsageRecordGroup

`func NewMeteringUsageRecordGroup() *MeteringUsageRecordGroup`

NewMeteringUsageRecordGroup instantiates a new MeteringUsageRecordGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordGroupWithDefaults

`func NewMeteringUsageRecordGroupWithDefaults() *MeteringUsageRecordGroup`

NewMeteringUsageRecordGroupWithDefaults instantiates a new MeteringUsageRecordGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *MeteringUsageRecordGroup) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *MeteringUsageRecordGroup) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *MeteringUsageRecordGroup) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *MeteringUsageRecordGroup) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *MeteringUsageRecordGroup) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MeteringUsageRecordGroup) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MeteringUsageRecordGroup) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MeteringUsageRecordGroup) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEntitlementID

`func (o *MeteringUsageRecordGroup) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *MeteringUsageRecordGroup) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *MeteringUsageRecordGroup) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *MeteringUsageRecordGroup) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetId

`func (o *MeteringUsageRecordGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeteringUsageRecordGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeteringUsageRecordGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeteringUsageRecordGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *MeteringUsageRecordGroup) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *MeteringUsageRecordGroup) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *MeteringUsageRecordGroup) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *MeteringUsageRecordGroup) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMetaInfo

`func (o *MeteringUsageRecordGroup) GetMetaInfo() MeteringUsageRecordGroupMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *MeteringUsageRecordGroup) GetMetaInfoOk() (*MeteringUsageRecordGroupMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *MeteringUsageRecordGroup) SetMetaInfo(v MeteringUsageRecordGroupMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *MeteringUsageRecordGroup) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetOrganizationID

`func (o *MeteringUsageRecordGroup) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *MeteringUsageRecordGroup) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *MeteringUsageRecordGroup) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *MeteringUsageRecordGroup) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *MeteringUsageRecordGroup) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *MeteringUsageRecordGroup) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *MeteringUsageRecordGroup) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *MeteringUsageRecordGroup) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetRecords

`func (o *MeteringUsageRecordGroup) GetRecords() map[string]float32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *MeteringUsageRecordGroup) GetRecordsOk() (*map[string]float32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *MeteringUsageRecordGroup) SetRecords(v map[string]float32)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *MeteringUsageRecordGroup) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetSerialID

`func (o *MeteringUsageRecordGroup) GetSerialID() int32`

GetSerialID returns the SerialID field if non-nil, zero value otherwise.

### GetSerialIDOk

`func (o *MeteringUsageRecordGroup) GetSerialIDOk() (*int32, bool)`

GetSerialIDOk returns a tuple with the SerialID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialID

`func (o *MeteringUsageRecordGroup) SetSerialID(v int32)`

SetSerialID sets SerialID field to given value.

### HasSerialID

`func (o *MeteringUsageRecordGroup) HasSerialID() bool`

HasSerialID returns a boolean if a field has been set.

### GetStatus

`func (o *MeteringUsageRecordGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeteringUsageRecordGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeteringUsageRecordGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeteringUsageRecordGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsageRecordReportID

`func (o *MeteringUsageRecordGroup) GetUsageRecordReportID() string`

GetUsageRecordReportID returns the UsageRecordReportID field if non-nil, zero value otherwise.

### GetUsageRecordReportIDOk

`func (o *MeteringUsageRecordGroup) GetUsageRecordReportIDOk() (*string, bool)`

GetUsageRecordReportIDOk returns a tuple with the UsageRecordReportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordReportID

`func (o *MeteringUsageRecordGroup) SetUsageRecordReportID(v string)`

SetUsageRecordReportID sets UsageRecordReportID field to given value.

### HasUsageRecordReportID

`func (o *MeteringUsageRecordGroup) HasUsageRecordReportID() bool`

HasUsageRecordReportID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


