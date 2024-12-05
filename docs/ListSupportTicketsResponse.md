# ListSupportTicketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SupportTicket**](SupportTicket.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | Only available when the request is made with offset&#x3D;0. | [optional] 

## Methods

### NewListSupportTicketsResponse

`func NewListSupportTicketsResponse() *ListSupportTicketsResponse`

NewListSupportTicketsResponse instantiates a new ListSupportTicketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportTicketsResponseWithDefaults

`func NewListSupportTicketsResponseWithDefaults() *ListSupportTicketsResponse`

NewListSupportTicketsResponseWithDefaults instantiates a new ListSupportTicketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSupportTicketsResponse) GetItems() []SupportTicket`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSupportTicketsResponse) GetItemsOk() (*[]SupportTicket, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSupportTicketsResponse) SetItems(v []SupportTicket)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListSupportTicketsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListSupportTicketsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListSupportTicketsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListSupportTicketsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListSupportTicketsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


