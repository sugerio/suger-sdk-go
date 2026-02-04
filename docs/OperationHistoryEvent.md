# OperationHistoryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**EventID** | Pointer to **int32** |  | [optional] 
**EventTime** | Pointer to **time.Time** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**TaskID** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WorkerMayIgnore** | Pointer to **bool** |  | [optional] 

## Methods

### NewOperationHistoryEvent

`func NewOperationHistoryEvent() *OperationHistoryEvent`

NewOperationHistoryEvent instantiates a new OperationHistoryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationHistoryEventWithDefaults

`func NewOperationHistoryEventWithDefaults() *OperationHistoryEvent`

NewOperationHistoryEventWithDefaults instantiates a new OperationHistoryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *OperationHistoryEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OperationHistoryEvent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OperationHistoryEvent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *OperationHistoryEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEventID

`func (o *OperationHistoryEvent) GetEventID() int32`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *OperationHistoryEvent) GetEventIDOk() (*int32, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *OperationHistoryEvent) SetEventID(v int32)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *OperationHistoryEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventTime

`func (o *OperationHistoryEvent) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *OperationHistoryEvent) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *OperationHistoryEvent) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *OperationHistoryEvent) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *OperationHistoryEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OperationHistoryEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OperationHistoryEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OperationHistoryEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTaskID

`func (o *OperationHistoryEvent) GetTaskID() int32`

GetTaskID returns the TaskID field if non-nil, zero value otherwise.

### GetTaskIDOk

`func (o *OperationHistoryEvent) GetTaskIDOk() (*int32, bool)`

GetTaskIDOk returns a tuple with the TaskID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskID

`func (o *OperationHistoryEvent) SetTaskID(v int32)`

SetTaskID sets TaskID field to given value.

### HasTaskID

`func (o *OperationHistoryEvent) HasTaskID() bool`

HasTaskID returns a boolean if a field has been set.

### GetVersion

`func (o *OperationHistoryEvent) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OperationHistoryEvent) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OperationHistoryEvent) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OperationHistoryEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkerMayIgnore

`func (o *OperationHistoryEvent) GetWorkerMayIgnore() bool`

GetWorkerMayIgnore returns the WorkerMayIgnore field if non-nil, zero value otherwise.

### GetWorkerMayIgnoreOk

`func (o *OperationHistoryEvent) GetWorkerMayIgnoreOk() (*bool, bool)`

GetWorkerMayIgnoreOk returns a tuple with the WorkerMayIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerMayIgnore

`func (o *OperationHistoryEvent) SetWorkerMayIgnore(v bool)`

SetWorkerMayIgnore sets WorkerMayIgnore field to given value.

### HasWorkerMayIgnore

`func (o *OperationHistoryEvent) HasWorkerMayIgnore() bool`

HasWorkerMayIgnore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


