# AceEventBridgeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to [**AceEventBridgeEventDetail**](AceEventBridgeEventDetail.md) |  | [optional] 
**DetailType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** | \&quot;aws.partnercentral-selling\&quot; | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAceEventBridgeEvent

`func NewAceEventBridgeEvent() *AceEventBridgeEvent`

NewAceEventBridgeEvent instantiates a new AceEventBridgeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAceEventBridgeEventWithDefaults

`func NewAceEventBridgeEventWithDefaults() *AceEventBridgeEvent`

NewAceEventBridgeEventWithDefaults instantiates a new AceEventBridgeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AceEventBridgeEvent) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AceEventBridgeEvent) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AceEventBridgeEvent) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AceEventBridgeEvent) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetail

`func (o *AceEventBridgeEvent) GetDetail() AceEventBridgeEventDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AceEventBridgeEvent) GetDetailOk() (*AceEventBridgeEventDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AceEventBridgeEvent) SetDetail(v AceEventBridgeEventDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AceEventBridgeEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDetailType

`func (o *AceEventBridgeEvent) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AceEventBridgeEvent) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AceEventBridgeEvent) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.

### HasDetailType

`func (o *AceEventBridgeEvent) HasDetailType() bool`

HasDetailType returns a boolean if a field has been set.

### GetId

`func (o *AceEventBridgeEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AceEventBridgeEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AceEventBridgeEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AceEventBridgeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegion

`func (o *AceEventBridgeEvent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AceEventBridgeEvent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AceEventBridgeEvent) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AceEventBridgeEvent) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResources

`func (o *AceEventBridgeEvent) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AceEventBridgeEvent) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AceEventBridgeEvent) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AceEventBridgeEvent) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSource

`func (o *AceEventBridgeEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AceEventBridgeEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AceEventBridgeEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AceEventBridgeEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTime

`func (o *AceEventBridgeEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AceEventBridgeEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AceEventBridgeEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AceEventBridgeEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetVersion

`func (o *AceEventBridgeEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AceEventBridgeEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AceEventBridgeEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AceEventBridgeEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


