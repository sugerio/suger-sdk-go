# StripeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewStripeError

`func NewStripeError() *StripeError`

NewStripeError instantiates a new StripeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeErrorWithDefaults

`func NewStripeErrorWithDefaults() *StripeError`

NewStripeErrorWithDefaults instantiates a new StripeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *StripeError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StripeError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StripeError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StripeError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *StripeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StripeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StripeError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StripeError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParam

`func (o *StripeError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *StripeError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *StripeError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *StripeError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetStatus

`func (o *StripeError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StripeError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StripeError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StripeError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *StripeError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


