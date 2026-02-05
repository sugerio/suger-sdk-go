# AwsMarketplaceEventBridgeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The seller/ISV AWS Account Id. | [optional] 
**Detail** | Pointer to [**AwsMarketplaceEventBridgeEventDetail**](AwsMarketplaceEventBridgeEventDetail.md) |  | [optional] 
**DetailType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** | \&quot;aws.marketplacecatalog\&quot; | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsMarketplaceEventBridgeEvent

`func NewAwsMarketplaceEventBridgeEvent() *AwsMarketplaceEventBridgeEvent`

NewAwsMarketplaceEventBridgeEvent instantiates a new AwsMarketplaceEventBridgeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceEventBridgeEventWithDefaults

`func NewAwsMarketplaceEventBridgeEventWithDefaults() *AwsMarketplaceEventBridgeEvent`

NewAwsMarketplaceEventBridgeEventWithDefaults instantiates a new AwsMarketplaceEventBridgeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AwsMarketplaceEventBridgeEvent) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AwsMarketplaceEventBridgeEvent) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AwsMarketplaceEventBridgeEvent) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AwsMarketplaceEventBridgeEvent) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetail

`func (o *AwsMarketplaceEventBridgeEvent) GetDetail() AwsMarketplaceEventBridgeEventDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AwsMarketplaceEventBridgeEvent) GetDetailOk() (*AwsMarketplaceEventBridgeEventDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AwsMarketplaceEventBridgeEvent) SetDetail(v AwsMarketplaceEventBridgeEventDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AwsMarketplaceEventBridgeEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDetailType

`func (o *AwsMarketplaceEventBridgeEvent) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AwsMarketplaceEventBridgeEvent) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AwsMarketplaceEventBridgeEvent) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.

### HasDetailType

`func (o *AwsMarketplaceEventBridgeEvent) HasDetailType() bool`

HasDetailType returns a boolean if a field has been set.

### GetId

`func (o *AwsMarketplaceEventBridgeEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsMarketplaceEventBridgeEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsMarketplaceEventBridgeEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsMarketplaceEventBridgeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegion

`func (o *AwsMarketplaceEventBridgeEvent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsMarketplaceEventBridgeEvent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsMarketplaceEventBridgeEvent) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsMarketplaceEventBridgeEvent) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResources

`func (o *AwsMarketplaceEventBridgeEvent) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsMarketplaceEventBridgeEvent) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsMarketplaceEventBridgeEvent) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsMarketplaceEventBridgeEvent) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSource

`func (o *AwsMarketplaceEventBridgeEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AwsMarketplaceEventBridgeEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AwsMarketplaceEventBridgeEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AwsMarketplaceEventBridgeEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTime

`func (o *AwsMarketplaceEventBridgeEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AwsMarketplaceEventBridgeEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AwsMarketplaceEventBridgeEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AwsMarketplaceEventBridgeEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetVersion

`func (o *AwsMarketplaceEventBridgeEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AwsMarketplaceEventBridgeEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AwsMarketplaceEventBridgeEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AwsMarketplaceEventBridgeEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


