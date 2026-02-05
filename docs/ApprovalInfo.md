# ApprovalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalStatus** | Pointer to [**ApprovalStatus**](ApprovalStatus.md) | ApprovalStatus holds the current approval status of the offer | \&quot;Submitted\&quot; | \&quot;Approved\&quot; | \&quot;Declined\&quot; | \&quot;Action Required\&quot;; | [optional] 
**DecisionDate** | Pointer to **string** | DecisionDate is when the final approval/decline happened (nil when pending) Latest DecisionDate | [optional] 
**Message** | Pointer to **string** | Message is the reason or explanation provided when the approval status is set to Declined or Action Required. It always stores the latest message for the current status transition. Historical messages are stored in notification events. | [optional] 
**RequestDate** | Pointer to **string** | Latest  RequestDate | [optional] 

## Methods

### NewApprovalInfo

`func NewApprovalInfo() *ApprovalInfo`

NewApprovalInfo instantiates a new ApprovalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalInfoWithDefaults

`func NewApprovalInfoWithDefaults() *ApprovalInfo`

NewApprovalInfoWithDefaults instantiates a new ApprovalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalStatus

`func (o *ApprovalInfo) GetApprovalStatus() ApprovalStatus`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *ApprovalInfo) GetApprovalStatusOk() (*ApprovalStatus, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *ApprovalInfo) SetApprovalStatus(v ApprovalStatus)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *ApprovalInfo) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetDecisionDate

`func (o *ApprovalInfo) GetDecisionDate() string`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *ApprovalInfo) GetDecisionDateOk() (*string, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *ApprovalInfo) SetDecisionDate(v string)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *ApprovalInfo) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### GetMessage

`func (o *ApprovalInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApprovalInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApprovalInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApprovalInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRequestDate

`func (o *ApprovalInfo) GetRequestDate() string`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *ApprovalInfo) GetRequestDateOk() (*string, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *ApprovalInfo) SetRequestDate(v string)`

SetRequestDate sets RequestDate field to given value.

### HasRequestDate

`func (o *ApprovalInfo) HasRequestDate() bool`

HasRequestDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


