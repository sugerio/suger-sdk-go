# SupportTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]SupportTicketAttachment**](SupportTicketAttachment.md) |  | [optional] 
**CloseTime** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]SupportTicketComment**](SupportTicketComment.md) |  | [optional] 
**CreationTime** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DueTime** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**SupportTicketPriority**](SupportTicketPriority.md) |  | [optional] 
**Status** | Pointer to [**SupportTicketStatus**](SupportTicketStatus.md) |  | [optional] 
**Watchers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSupportTicket

`func NewSupportTicket() *SupportTicket`

NewSupportTicket instantiates a new SupportTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketWithDefaults

`func NewSupportTicketWithDefaults() *SupportTicket`

NewSupportTicketWithDefaults instantiates a new SupportTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *SupportTicket) GetAttachments() []SupportTicketAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SupportTicket) GetAttachmentsOk() (*[]SupportTicketAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SupportTicket) SetAttachments(v []SupportTicketAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SupportTicket) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCloseTime

`func (o *SupportTicket) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SupportTicket) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SupportTicket) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SupportTicket) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetComments

`func (o *SupportTicket) GetComments() []SupportTicketComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SupportTicket) GetCommentsOk() (*[]SupportTicketComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SupportTicket) SetComments(v []SupportTicketComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *SupportTicket) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreationTime

`func (o *SupportTicket) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SupportTicket) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SupportTicket) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SupportTicket) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *SupportTicket) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SupportTicket) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SupportTicket) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SupportTicket) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *SupportTicket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportTicket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportTicket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportTicket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueTime

`func (o *SupportTicket) GetDueTime() string`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *SupportTicket) GetDueTimeOk() (*string, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *SupportTicket) SetDueTime(v string)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *SupportTicket) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetId

`func (o *SupportTicket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SupportTicket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportTicket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportTicket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportTicket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SupportTicket) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SupportTicket) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SupportTicket) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SupportTicket) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPriority

`func (o *SupportTicket) GetPriority() SupportTicketPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SupportTicket) GetPriorityOk() (*SupportTicketPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SupportTicket) SetPriority(v SupportTicketPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SupportTicket) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *SupportTicket) GetStatus() SupportTicketStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportTicket) GetStatusOk() (*SupportTicketStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportTicket) SetStatus(v SupportTicketStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportTicket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWatchers

`func (o *SupportTicket) GetWatchers() []string`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *SupportTicket) GetWatchersOk() (*[]string, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *SupportTicket) SetWatchers(v []string)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *SupportTicket) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


