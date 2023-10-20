# ServicecontrolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Code: The status code, which should be an enum value of google.rpc.Code. | [optional] 
**Details** | Pointer to **[][]int32** | Details: A list of messages that carry the error details. There is a common set of message types for APIs to use. | [optional] 
**Message** | Pointer to **string** | Message: A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client. | [optional] 

## Methods

### NewServicecontrolStatus

`func NewServicecontrolStatus() *ServicecontrolStatus`

NewServicecontrolStatus instantiates a new ServicecontrolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicecontrolStatusWithDefaults

`func NewServicecontrolStatusWithDefaults() *ServicecontrolStatus`

NewServicecontrolStatusWithDefaults instantiates a new ServicecontrolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ServicecontrolStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServicecontrolStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServicecontrolStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServicecontrolStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *ServicecontrolStatus) GetDetails() [][]int32`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServicecontrolStatus) GetDetailsOk() (*[][]int32, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServicecontrolStatus) SetDetails(v [][]int32)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServicecontrolStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *ServicecontrolStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServicecontrolStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServicecontrolStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServicecontrolStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


