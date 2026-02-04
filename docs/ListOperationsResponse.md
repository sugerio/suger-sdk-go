# ListOperationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetToken** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]Operation**](Operation.md) |  | [optional] 

## Methods

### NewListOperationsResponse

`func NewListOperationsResponse() *ListOperationsResponse`

NewListOperationsResponse instantiates a new ListOperationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperationsResponseWithDefaults

`func NewListOperationsResponseWithDefaults() *ListOperationsResponse`

NewListOperationsResponseWithDefaults instantiates a new ListOperationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsetToken

`func (o *ListOperationsResponse) GetOffsetToken() string`

GetOffsetToken returns the OffsetToken field if non-nil, zero value otherwise.

### GetOffsetTokenOk

`func (o *ListOperationsResponse) GetOffsetTokenOk() (*string, bool)`

GetOffsetTokenOk returns a tuple with the OffsetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetToken

`func (o *ListOperationsResponse) SetOffsetToken(v string)`

SetOffsetToken sets OffsetToken field to given value.

### HasOffsetToken

`func (o *ListOperationsResponse) HasOffsetToken() bool`

HasOffsetToken returns a boolean if a field has been set.

### GetOperations

`func (o *ListOperationsResponse) GetOperations() []Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ListOperationsResponse) GetOperationsOk() (*[]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ListOperationsResponse) SetOperations(v []Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ListOperationsResponse) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


