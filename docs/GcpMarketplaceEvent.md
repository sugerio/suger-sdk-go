# GcpMarketplaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**GcpMarketplaceUserAccount**](GcpMarketplaceUserAccount.md) |  | [optional] 
**Entitlement** | Pointer to [**GcpMarketplaceEntitlement**](GcpMarketplaceEntitlement.md) |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**GcpMarketplacceEventType**](GcpMarketplacceEventType.md) |  | [optional] 
**ProviderId** | Pointer to **string** | GCP Partner ID of the SaaS Seller. | [optional] 
**PublishTime** | Pointer to **time.Time** | The Publish Time of the event. | [optional] 
**SugerOrganizationId** | Pointer to **string** | Populated by Suger Service. | [optional] 

## Methods

### NewGcpMarketplaceEvent

`func NewGcpMarketplaceEvent() *GcpMarketplaceEvent`

NewGcpMarketplaceEvent instantiates a new GcpMarketplaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceEventWithDefaults

`func NewGcpMarketplaceEventWithDefaults() *GcpMarketplaceEvent`

NewGcpMarketplaceEventWithDefaults instantiates a new GcpMarketplaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GcpMarketplaceEvent) GetAccount() GcpMarketplaceUserAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GcpMarketplaceEvent) GetAccountOk() (*GcpMarketplaceUserAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GcpMarketplaceEvent) SetAccount(v GcpMarketplaceUserAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GcpMarketplaceEvent) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetEntitlement

`func (o *GcpMarketplaceEvent) GetEntitlement() GcpMarketplaceEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *GcpMarketplaceEvent) GetEntitlementOk() (*GcpMarketplaceEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *GcpMarketplaceEvent) SetEntitlement(v GcpMarketplaceEntitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *GcpMarketplaceEvent) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetEventId

`func (o *GcpMarketplaceEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *GcpMarketplaceEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *GcpMarketplaceEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *GcpMarketplaceEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *GcpMarketplaceEvent) GetEventType() GcpMarketplacceEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GcpMarketplaceEvent) GetEventTypeOk() (*GcpMarketplacceEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GcpMarketplaceEvent) SetEventType(v GcpMarketplacceEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GcpMarketplaceEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetProviderId

`func (o *GcpMarketplaceEvent) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GcpMarketplaceEvent) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GcpMarketplaceEvent) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GcpMarketplaceEvent) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetPublishTime

`func (o *GcpMarketplaceEvent) GetPublishTime() time.Time`

GetPublishTime returns the PublishTime field if non-nil, zero value otherwise.

### GetPublishTimeOk

`func (o *GcpMarketplaceEvent) GetPublishTimeOk() (*time.Time, bool)`

GetPublishTimeOk returns a tuple with the PublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTime

`func (o *GcpMarketplaceEvent) SetPublishTime(v time.Time)`

SetPublishTime sets PublishTime field to given value.

### HasPublishTime

`func (o *GcpMarketplaceEvent) HasPublishTime() bool`

HasPublishTime returns a boolean if a field has been set.

### GetSugerOrganizationId

`func (o *GcpMarketplaceEvent) GetSugerOrganizationId() string`

GetSugerOrganizationId returns the SugerOrganizationId field if non-nil, zero value otherwise.

### GetSugerOrganizationIdOk

`func (o *GcpMarketplaceEvent) GetSugerOrganizationIdOk() (*string, bool)`

GetSugerOrganizationIdOk returns a tuple with the SugerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerOrganizationId

`func (o *GcpMarketplaceEvent) SetSugerOrganizationId(v string)`

SetSugerOrganizationId sets SugerOrganizationId field to given value.

### HasSugerOrganizationId

`func (o *GcpMarketplaceEvent) HasSugerOrganizationId() bool`

HasSugerOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


