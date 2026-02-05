# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildWorkflowOperations** | Pointer to [**[]Operation**](Operation.md) | Only populated if it has child workflows. Otherwise it is nil. | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** | Operation ID. | [optional] 
**Memo** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** | Run ID. | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OperationType**](OperationType.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildWorkflowOperations

`func (o *Operation) GetChildWorkflowOperations() []Operation`

GetChildWorkflowOperations returns the ChildWorkflowOperations field if non-nil, zero value otherwise.

### GetChildWorkflowOperationsOk

`func (o *Operation) GetChildWorkflowOperationsOk() (*[]Operation, bool)`

GetChildWorkflowOperationsOk returns a tuple with the ChildWorkflowOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildWorkflowOperations

`func (o *Operation) SetChildWorkflowOperations(v []Operation)`

SetChildWorkflowOperations sets ChildWorkflowOperations field to given value.

### HasChildWorkflowOperations

`func (o *Operation) HasChildWorkflowOperations() bool`

HasChildWorkflowOperations returns a boolean if a field has been set.

### GetEndTime

`func (o *Operation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Operation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Operation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Operation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *Operation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Operation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Operation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Operation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemo

`func (o *Operation) GetMemo() map[string]map[string]interface{}`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Operation) GetMemoOk() (*map[string]map[string]interface{}, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Operation) SetMemo(v map[string]map[string]interface{})`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Operation) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMessage

`func (o *Operation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Operation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Operation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Operation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *Operation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Operation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Operation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Operation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunId

`func (o *Operation) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Operation) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Operation) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *Operation) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStartTime

`func (o *Operation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Operation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Operation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Operation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Operation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Operation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Operation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Operation) GetType() OperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Operation) GetTypeOk() (*OperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Operation) SetType(v OperationType)`

SetType sets Type field to given value.

### HasType

`func (o *Operation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


