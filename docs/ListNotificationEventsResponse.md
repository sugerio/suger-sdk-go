# ListNotificationEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** | If it is nil, it means there is no more records. | [optional] 
**NotificationEvents** | Pointer to [**[]NotificationEvent**](NotificationEvent.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | Only available when the request is made with offset&#x3D;0. | [optional] 

## Methods

### NewListNotificationEventsResponse

`func NewListNotificationEventsResponse() *ListNotificationEventsResponse`

NewListNotificationEventsResponse instantiates a new ListNotificationEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationEventsResponseWithDefaults

`func NewListNotificationEventsResponseWithDefaults() *ListNotificationEventsResponse`

NewListNotificationEventsResponseWithDefaults instantiates a new ListNotificationEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListNotificationEventsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListNotificationEventsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListNotificationEventsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListNotificationEventsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetNotificationEvents

`func (o *ListNotificationEventsResponse) GetNotificationEvents() []NotificationEvent`

GetNotificationEvents returns the NotificationEvents field if non-nil, zero value otherwise.

### GetNotificationEventsOk

`func (o *ListNotificationEventsResponse) GetNotificationEventsOk() (*[]NotificationEvent, bool)`

GetNotificationEventsOk returns a tuple with the NotificationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEvents

`func (o *ListNotificationEventsResponse) SetNotificationEvents(v []NotificationEvent)`

SetNotificationEvents sets NotificationEvents field to given value.

### HasNotificationEvents

`func (o *ListNotificationEventsResponse) HasNotificationEvents() bool`

HasNotificationEvents returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListNotificationEventsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListNotificationEventsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListNotificationEventsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListNotificationEventsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


