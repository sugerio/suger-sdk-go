# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SearchResultItem**](SearchResultItem.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchResponse) GetItems() []SearchResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchResponse) GetItemsOk() (*[]SearchResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchResponse) SetItems(v []SearchResultItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextCursor

`func (o *SearchResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SearchResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SearchResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *SearchResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


