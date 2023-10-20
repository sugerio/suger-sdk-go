# ServicecontrolReportError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **string** | OperationId: The Operation.operation_id value from the request. | [optional] 
**Status** | Pointer to [**ServicecontrolStatus**](ServicecontrolStatus.md) |  | [optional] 

## Methods

### NewServicecontrolReportError

`func NewServicecontrolReportError() *ServicecontrolReportError`

NewServicecontrolReportError instantiates a new ServicecontrolReportError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicecontrolReportErrorWithDefaults

`func NewServicecontrolReportErrorWithDefaults() *ServicecontrolReportError`

NewServicecontrolReportErrorWithDefaults instantiates a new ServicecontrolReportError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *ServicecontrolReportError) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ServicecontrolReportError) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ServicecontrolReportError) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *ServicecontrolReportError) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetStatus

`func (o *ServicecontrolReportError) GetStatus() ServicecontrolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicecontrolReportError) GetStatusOk() (*ServicecontrolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicecontrolReportError) SetStatus(v ServicecontrolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicecontrolReportError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


