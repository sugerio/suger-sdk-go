# NotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**NotificationEventAction**](NotificationEventAction.md) |  | [optional] 
**CcContactIds** | Pointer to **[]string** | Cc contactIds that will receive this notification | [optional] 
**ContactIds** | Pointer to **[]string** | ContactIds that will receive this notification | [optional] 
**EntityID** | Pointer to **string** |  | [optional] 
**EntityStatus** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**EventID** | Pointer to **string** | notification event id. | [optional] 
**EventStatus** | Pointer to [**NotificationEventStatus**](NotificationEventStatus.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | timestamp of the event when it is updated. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** | suger organization id. | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** | timestamp of the event when it is scheduled or created. | [optional] 
**Title** | Pointer to **string** | The title of the notification event such as email subject. | [optional] 
**TrackEvents** | Pointer to [**[]TrackEvent**](TrackEvent.md) | The track events of the notification event. | [optional] 

## Methods

### NewNotificationEvent

`func NewNotificationEvent() *NotificationEvent`

NewNotificationEvent instantiates a new NotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventWithDefaults

`func NewNotificationEventWithDefaults() *NotificationEvent`

NewNotificationEventWithDefaults instantiates a new NotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NotificationEvent) GetAction() NotificationEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NotificationEvent) GetActionOk() (*NotificationEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NotificationEvent) SetAction(v NotificationEventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *NotificationEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCcContactIds

`func (o *NotificationEvent) GetCcContactIds() []string`

GetCcContactIds returns the CcContactIds field if non-nil, zero value otherwise.

### GetCcContactIdsOk

`func (o *NotificationEvent) GetCcContactIdsOk() (*[]string, bool)`

GetCcContactIdsOk returns a tuple with the CcContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcContactIds

`func (o *NotificationEvent) SetCcContactIds(v []string)`

SetCcContactIds sets CcContactIds field to given value.

### HasCcContactIds

`func (o *NotificationEvent) HasCcContactIds() bool`

HasCcContactIds returns a boolean if a field has been set.

### GetContactIds

`func (o *NotificationEvent) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *NotificationEvent) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *NotificationEvent) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *NotificationEvent) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetEntityID

`func (o *NotificationEvent) GetEntityID() string`

GetEntityID returns the EntityID field if non-nil, zero value otherwise.

### GetEntityIDOk

`func (o *NotificationEvent) GetEntityIDOk() (*string, bool)`

GetEntityIDOk returns a tuple with the EntityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityID

`func (o *NotificationEvent) SetEntityID(v string)`

SetEntityID sets EntityID field to given value.

### HasEntityID

`func (o *NotificationEvent) HasEntityID() bool`

HasEntityID returns a boolean if a field has been set.

### GetEntityStatus

`func (o *NotificationEvent) GetEntityStatus() string`

GetEntityStatus returns the EntityStatus field if non-nil, zero value otherwise.

### GetEntityStatusOk

`func (o *NotificationEvent) GetEntityStatusOk() (*string, bool)`

GetEntityStatusOk returns a tuple with the EntityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityStatus

`func (o *NotificationEvent) SetEntityStatus(v string)`

SetEntityStatus sets EntityStatus field to given value.

### HasEntityStatus

`func (o *NotificationEvent) HasEntityStatus() bool`

HasEntityStatus returns a boolean if a field has been set.

### GetEntityType

`func (o *NotificationEvent) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *NotificationEvent) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *NotificationEvent) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *NotificationEvent) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEventID

`func (o *NotificationEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *NotificationEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *NotificationEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *NotificationEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventStatus

`func (o *NotificationEvent) GetEventStatus() NotificationEventStatus`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *NotificationEvent) GetEventStatusOk() (*NotificationEventStatus, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *NotificationEvent) SetEventStatus(v NotificationEventStatus)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *NotificationEvent) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *NotificationEvent) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *NotificationEvent) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *NotificationEvent) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *NotificationEvent) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMessage

`func (o *NotificationEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrganizationID

`func (o *NotificationEvent) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *NotificationEvent) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *NotificationEvent) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *NotificationEvent) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *NotificationEvent) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *NotificationEvent) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *NotificationEvent) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *NotificationEvent) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetTimestamp

`func (o *NotificationEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NotificationEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NotificationEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NotificationEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *NotificationEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTrackEvents

`func (o *NotificationEvent) GetTrackEvents() []TrackEvent`

GetTrackEvents returns the TrackEvents field if non-nil, zero value otherwise.

### GetTrackEventsOk

`func (o *NotificationEvent) GetTrackEventsOk() (*[]TrackEvent, bool)`

GetTrackEventsOk returns a tuple with the TrackEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackEvents

`func (o *NotificationEvent) SetTrackEvents(v []TrackEvent)`

SetTrackEvents sets TrackEvents field to given value.

### HasTrackEvents

`func (o *NotificationEvent) HasTrackEvents() bool`

HasTrackEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


