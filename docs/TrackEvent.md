# TrackEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**TrackEventActionType**](TrackEventActionType.md) |  | [optional] 
**ContactId** | Pointer to **string** | The ID of the contact who triggered the track event if applicable. | [optional] 
**Timestamp** | Pointer to **time.Time** | timestamp of the track event happened. | [optional] 

## Methods

### NewTrackEvent

`func NewTrackEvent() *TrackEvent`

NewTrackEvent instantiates a new TrackEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackEventWithDefaults

`func NewTrackEventWithDefaults() *TrackEvent`

NewTrackEventWithDefaults instantiates a new TrackEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TrackEvent) GetAction() TrackEventActionType`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TrackEvent) GetActionOk() (*TrackEventActionType, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TrackEvent) SetAction(v TrackEventActionType)`

SetAction sets Action field to given value.

### HasAction

`func (o *TrackEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetContactId

`func (o *TrackEvent) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *TrackEvent) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *TrackEvent) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *TrackEvent) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TrackEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrackEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrackEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TrackEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


