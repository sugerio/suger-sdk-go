# ListRevenueRecordsResponse

## Properties

 Name               | Type                                               | Description              | Notes      
--------------------|----------------------------------------------------|--------------------------|------------
 **NextOffset**     | Pointer to **int32**                               |                          | [optional] 
 **RevenueRecords** | Pointer to [**[]RevenueRecord**](RevenueRecord.md) | list of revenue records. | [optional] 

## Methods

### NewListRevenueRecordsResponse

`func NewListRevenueRecordsResponse() *ListRevenueRecordsResponse`

NewListRevenueRecordsResponse instantiates a new ListRevenueRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRevenueRecordsResponseWithDefaults

`func NewListRevenueRecordsResponseWithDefaults() *ListRevenueRecordsResponse`

NewListRevenueRecordsResponseWithDefaults instantiates a new ListRevenueRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListRevenueRecordsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListRevenueRecordsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListRevenueRecordsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListRevenueRecordsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetRevenueRecords

`func (o *ListRevenueRecordsResponse) GetRevenueRecords() []RevenueRecord`

GetRevenueRecords returns the RevenueRecords field if non-nil, zero value otherwise.

### GetRevenueRecordsOk

`func (o *ListRevenueRecordsResponse) GetRevenueRecordsOk() (*[]RevenueRecord, bool)`

GetRevenueRecordsOk returns a tuple with the RevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecords

`func (o *ListRevenueRecordsResponse) SetRevenueRecords(v []RevenueRecord)`

SetRevenueRecords sets RevenueRecords field to given value.

### HasRevenueRecords

`func (o *ListRevenueRecordsResponse) HasRevenueRecords() bool`

HasRevenueRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


