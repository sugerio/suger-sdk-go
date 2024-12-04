# UpdateSupportTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | [**SupportTicketPriority**](SupportTicketPriority.md) |  | 
**Watchers** | **[]string** |  | 

## Methods

### NewUpdateSupportTicketRequest

`func NewUpdateSupportTicketRequest(priority SupportTicketPriority, watchers []string, ) *UpdateSupportTicketRequest`

NewUpdateSupportTicketRequest instantiates a new UpdateSupportTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSupportTicketRequestWithDefaults

`func NewUpdateSupportTicketRequestWithDefaults() *UpdateSupportTicketRequest`

NewUpdateSupportTicketRequestWithDefaults instantiates a new UpdateSupportTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *UpdateSupportTicketRequest) GetPriority() SupportTicketPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateSupportTicketRequest) GetPriorityOk() (*SupportTicketPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateSupportTicketRequest) SetPriority(v SupportTicketPriority)`

SetPriority sets Priority field to given value.


### GetWatchers

`func (o *UpdateSupportTicketRequest) GetWatchers() []string`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *UpdateSupportTicketRequest) GetWatchersOk() (*[]string, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *UpdateSupportTicketRequest) SetWatchers(v []string)`

SetWatchers sets Watchers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


