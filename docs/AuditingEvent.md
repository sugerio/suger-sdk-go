# AuditingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AceEventBridgeEvent** | Pointer to [**AceEventBridgeEvent**](AceEventBridgeEvent.md) | Nullable, applicable when eventType &#x3D; AWS_ACE_EVENT_BRIDGE. | [optional] 
**AlibabaMarketplaceEvent** | Pointer to [**AlibabaMarketplaceEvent**](AlibabaMarketplaceEvent.md) | Nullable, applicable when eventType &#x3D; ALIBABA_MARKETPLACE. | [optional] 
**AwsMarketplaceEvent** | Pointer to [**AwsMarketplaceEvent**](AwsMarketplaceEvent.md) | Nullable, applicable when eventType &#x3D; AWS_MARKETPLACE. | [optional] 
**AwsMarketplaceEventBridgeEvent** | Pointer to [**AwsMarketplaceEventBridgeEvent**](AwsMarketplaceEventBridgeEvent.md) | Nullable, applicable when eventType &#x3D; AWS_EVENT_BRIDGE. | [optional] 
**AzureMarketplaceEvent** | Pointer to [**AzureMarketplaceEvent**](AzureMarketplaceEvent.md) | Nullable, applicable when eventType &#x3D; AZURE_MARKETPLACE. | [optional] 
**CreationTime** | Pointer to **time.Time** | When the event is received and audited. | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**GcpMarketplaceEvent** | Pointer to [**GcpMarketplaceEvent**](GcpMarketplaceEvent.md) | Nullable, applicable when eventType &#x3D; GCP_MARKETPLACE. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | when the event is updated. | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**OtherTypeEvent** | Pointer to **map[string]interface{}** | Nullable, applicable when eventType is other types. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAuditingEvent

`func NewAuditingEvent() *AuditingEvent`

NewAuditingEvent instantiates a new AuditingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditingEventWithDefaults

`func NewAuditingEventWithDefaults() *AuditingEvent`

NewAuditingEventWithDefaults instantiates a new AuditingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAceEventBridgeEvent

`func (o *AuditingEvent) GetAceEventBridgeEvent() AceEventBridgeEvent`

GetAceEventBridgeEvent returns the AceEventBridgeEvent field if non-nil, zero value otherwise.

### GetAceEventBridgeEventOk

`func (o *AuditingEvent) GetAceEventBridgeEventOk() (*AceEventBridgeEvent, bool)`

GetAceEventBridgeEventOk returns a tuple with the AceEventBridgeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceEventBridgeEvent

`func (o *AuditingEvent) SetAceEventBridgeEvent(v AceEventBridgeEvent)`

SetAceEventBridgeEvent sets AceEventBridgeEvent field to given value.

### HasAceEventBridgeEvent

`func (o *AuditingEvent) HasAceEventBridgeEvent() bool`

HasAceEventBridgeEvent returns a boolean if a field has been set.

### GetAlibabaMarketplaceEvent

`func (o *AuditingEvent) GetAlibabaMarketplaceEvent() AlibabaMarketplaceEvent`

GetAlibabaMarketplaceEvent returns the AlibabaMarketplaceEvent field if non-nil, zero value otherwise.

### GetAlibabaMarketplaceEventOk

`func (o *AuditingEvent) GetAlibabaMarketplaceEventOk() (*AlibabaMarketplaceEvent, bool)`

GetAlibabaMarketplaceEventOk returns a tuple with the AlibabaMarketplaceEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaMarketplaceEvent

`func (o *AuditingEvent) SetAlibabaMarketplaceEvent(v AlibabaMarketplaceEvent)`

SetAlibabaMarketplaceEvent sets AlibabaMarketplaceEvent field to given value.

### HasAlibabaMarketplaceEvent

`func (o *AuditingEvent) HasAlibabaMarketplaceEvent() bool`

HasAlibabaMarketplaceEvent returns a boolean if a field has been set.

### GetAwsMarketplaceEvent

`func (o *AuditingEvent) GetAwsMarketplaceEvent() AwsMarketplaceEvent`

GetAwsMarketplaceEvent returns the AwsMarketplaceEvent field if non-nil, zero value otherwise.

### GetAwsMarketplaceEventOk

`func (o *AuditingEvent) GetAwsMarketplaceEventOk() (*AwsMarketplaceEvent, bool)`

GetAwsMarketplaceEventOk returns a tuple with the AwsMarketplaceEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarketplaceEvent

`func (o *AuditingEvent) SetAwsMarketplaceEvent(v AwsMarketplaceEvent)`

SetAwsMarketplaceEvent sets AwsMarketplaceEvent field to given value.

### HasAwsMarketplaceEvent

`func (o *AuditingEvent) HasAwsMarketplaceEvent() bool`

HasAwsMarketplaceEvent returns a boolean if a field has been set.

### GetAwsMarketplaceEventBridgeEvent

`func (o *AuditingEvent) GetAwsMarketplaceEventBridgeEvent() AwsMarketplaceEventBridgeEvent`

GetAwsMarketplaceEventBridgeEvent returns the AwsMarketplaceEventBridgeEvent field if non-nil, zero value otherwise.

### GetAwsMarketplaceEventBridgeEventOk

`func (o *AuditingEvent) GetAwsMarketplaceEventBridgeEventOk() (*AwsMarketplaceEventBridgeEvent, bool)`

GetAwsMarketplaceEventBridgeEventOk returns a tuple with the AwsMarketplaceEventBridgeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarketplaceEventBridgeEvent

`func (o *AuditingEvent) SetAwsMarketplaceEventBridgeEvent(v AwsMarketplaceEventBridgeEvent)`

SetAwsMarketplaceEventBridgeEvent sets AwsMarketplaceEventBridgeEvent field to given value.

### HasAwsMarketplaceEventBridgeEvent

`func (o *AuditingEvent) HasAwsMarketplaceEventBridgeEvent() bool`

HasAwsMarketplaceEventBridgeEvent returns a boolean if a field has been set.

### GetAzureMarketplaceEvent

`func (o *AuditingEvent) GetAzureMarketplaceEvent() AzureMarketplaceEvent`

GetAzureMarketplaceEvent returns the AzureMarketplaceEvent field if non-nil, zero value otherwise.

### GetAzureMarketplaceEventOk

`func (o *AuditingEvent) GetAzureMarketplaceEventOk() (*AzureMarketplaceEvent, bool)`

GetAzureMarketplaceEventOk returns a tuple with the AzureMarketplaceEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureMarketplaceEvent

`func (o *AuditingEvent) SetAzureMarketplaceEvent(v AzureMarketplaceEvent)`

SetAzureMarketplaceEvent sets AzureMarketplaceEvent field to given value.

### HasAzureMarketplaceEvent

`func (o *AuditingEvent) HasAzureMarketplaceEvent() bool`

HasAzureMarketplaceEvent returns a boolean if a field has been set.

### GetCreationTime

`func (o *AuditingEvent) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AuditingEvent) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AuditingEvent) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AuditingEvent) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEventType

`func (o *AuditingEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AuditingEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AuditingEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AuditingEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetGcpMarketplaceEvent

`func (o *AuditingEvent) GetGcpMarketplaceEvent() GcpMarketplaceEvent`

GetGcpMarketplaceEvent returns the GcpMarketplaceEvent field if non-nil, zero value otherwise.

### GetGcpMarketplaceEventOk

`func (o *AuditingEvent) GetGcpMarketplaceEventOk() (*GcpMarketplaceEvent, bool)`

GetGcpMarketplaceEventOk returns a tuple with the GcpMarketplaceEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpMarketplaceEvent

`func (o *AuditingEvent) SetGcpMarketplaceEvent(v GcpMarketplaceEvent)`

SetGcpMarketplaceEvent sets GcpMarketplaceEvent field to given value.

### HasGcpMarketplaceEvent

`func (o *AuditingEvent) HasGcpMarketplaceEvent() bool`

HasGcpMarketplaceEvent returns a boolean if a field has been set.

### GetId

`func (o *AuditingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *AuditingEvent) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *AuditingEvent) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *AuditingEvent) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *AuditingEvent) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetOrganizationID

`func (o *AuditingEvent) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *AuditingEvent) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *AuditingEvent) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *AuditingEvent) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetOtherTypeEvent

`func (o *AuditingEvent) GetOtherTypeEvent() map[string]interface{}`

GetOtherTypeEvent returns the OtherTypeEvent field if non-nil, zero value otherwise.

### GetOtherTypeEventOk

`func (o *AuditingEvent) GetOtherTypeEventOk() (*map[string]interface{}, bool)`

GetOtherTypeEventOk returns a tuple with the OtherTypeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherTypeEvent

`func (o *AuditingEvent) SetOtherTypeEvent(v map[string]interface{})`

SetOtherTypeEvent sets OtherTypeEvent field to given value.

### HasOtherTypeEvent

`func (o *AuditingEvent) HasOtherTypeEvent() bool`

HasOtherTypeEvent returns a boolean if a field has been set.

### GetStatus

`func (o *AuditingEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditingEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditingEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditingEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


