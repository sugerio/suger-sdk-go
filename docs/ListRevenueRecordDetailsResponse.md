# ListRevenueRecordDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** |  | [optional] 
**RevenueRecordDetails** | Pointer to [**[]RevenueRecordDetail**](RevenueRecordDetail.md) |  | [optional] 

## Methods

### NewListRevenueRecordDetailsResponse

`func NewListRevenueRecordDetailsResponse() *ListRevenueRecordDetailsResponse`

NewListRevenueRecordDetailsResponse instantiates a new ListRevenueRecordDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRevenueRecordDetailsResponseWithDefaults

`func NewListRevenueRecordDetailsResponseWithDefaults() *ListRevenueRecordDetailsResponse`

NewListRevenueRecordDetailsResponseWithDefaults instantiates a new ListRevenueRecordDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListRevenueRecordDetailsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListRevenueRecordDetailsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListRevenueRecordDetailsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListRevenueRecordDetailsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetRevenueRecordDetails

`func (o *ListRevenueRecordDetailsResponse) GetRevenueRecordDetails() []RevenueRecordDetail`

GetRevenueRecordDetails returns the RevenueRecordDetails field if non-nil, zero value otherwise.

### GetRevenueRecordDetailsOk

`func (o *ListRevenueRecordDetailsResponse) GetRevenueRecordDetailsOk() (*[]RevenueRecordDetail, bool)`

GetRevenueRecordDetailsOk returns a tuple with the RevenueRecordDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecordDetails

`func (o *ListRevenueRecordDetailsResponse) SetRevenueRecordDetails(v []RevenueRecordDetail)`

SetRevenueRecordDetails sets RevenueRecordDetails field to given value.

### HasRevenueRecordDetails

`func (o *ListRevenueRecordDetailsResponse) HasRevenueRecordDetails() bool`

HasRevenueRecordDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


