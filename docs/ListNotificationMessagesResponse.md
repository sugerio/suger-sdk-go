# ListNotificationMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** | The next offset to use in the next request to get the next page of notification messages. If this field is null, there are no more notification messages to get. | [optional] 
**NotificationMessages** | Pointer to [**[]NotificationMessage**](NotificationMessage.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | The total number of notification messages. Only available when the request is made with the first offset &#x3D; 0. | [optional] 

## Methods

### NewListNotificationMessagesResponse

`func NewListNotificationMessagesResponse() *ListNotificationMessagesResponse`

NewListNotificationMessagesResponse instantiates a new ListNotificationMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationMessagesResponseWithDefaults

`func NewListNotificationMessagesResponseWithDefaults() *ListNotificationMessagesResponse`

NewListNotificationMessagesResponseWithDefaults instantiates a new ListNotificationMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListNotificationMessagesResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListNotificationMessagesResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListNotificationMessagesResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListNotificationMessagesResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetNotificationMessages

`func (o *ListNotificationMessagesResponse) GetNotificationMessages() []NotificationMessage`

GetNotificationMessages returns the NotificationMessages field if non-nil, zero value otherwise.

### GetNotificationMessagesOk

`func (o *ListNotificationMessagesResponse) GetNotificationMessagesOk() (*[]NotificationMessage, bool)`

GetNotificationMessagesOk returns a tuple with the NotificationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessages

`func (o *ListNotificationMessagesResponse) SetNotificationMessages(v []NotificationMessage)`

SetNotificationMessages sets NotificationMessages field to given value.

### HasNotificationMessages

`func (o *ListNotificationMessagesResponse) HasNotificationMessages() bool`

HasNotificationMessages returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListNotificationMessagesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListNotificationMessagesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListNotificationMessagesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListNotificationMessagesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


