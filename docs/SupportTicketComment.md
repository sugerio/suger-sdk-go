# SupportTicketComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**[]SupportTicketCommentDetail**](SupportTicketCommentDetail.md) |  | [optional] 
**CommentText** | Pointer to **string** | When creating a new comment, only CommentText is required. | [optional] 
**Creator** | Pointer to [**SupportTicketUser**](SupportTicketUser.md) | who created the comment | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewSupportTicketComment

`func NewSupportTicketComment() *SupportTicketComment`

NewSupportTicketComment instantiates a new SupportTicketComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketCommentWithDefaults

`func NewSupportTicketCommentWithDefaults() *SupportTicketComment`

NewSupportTicketCommentWithDefaults instantiates a new SupportTicketComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SupportTicketComment) GetComment() []SupportTicketCommentDetail`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SupportTicketComment) GetCommentOk() (*[]SupportTicketCommentDetail, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SupportTicketComment) SetComment(v []SupportTicketCommentDetail)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SupportTicketComment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCommentText

`func (o *SupportTicketComment) GetCommentText() string`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *SupportTicketComment) GetCommentTextOk() (*string, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *SupportTicketComment) SetCommentText(v string)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *SupportTicketComment) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.

### GetCreator

`func (o *SupportTicketComment) GetCreator() SupportTicketUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SupportTicketComment) GetCreatorOk() (*SupportTicketUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SupportTicketComment) SetCreator(v SupportTicketUser)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SupportTicketComment) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDate

`func (o *SupportTicketComment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SupportTicketComment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SupportTicketComment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SupportTicketComment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *SupportTicketComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketComment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketComment) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


