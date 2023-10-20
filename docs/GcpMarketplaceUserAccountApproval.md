# GcpMarketplaceUserAccountApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** | An explanation for the state of the approval. | [optional] 
**State** | Pointer to [**GcpMarketplaceUserAccountApprovalState**](GcpMarketplaceUserAccountApprovalState.md) |  | [optional] 
**UpdateTime** | Pointer to **string** | RFC3339 UTC timestamp | [optional] 

## Methods

### NewGcpMarketplaceUserAccountApproval

`func NewGcpMarketplaceUserAccountApproval() *GcpMarketplaceUserAccountApproval`

NewGcpMarketplaceUserAccountApproval instantiates a new GcpMarketplaceUserAccountApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceUserAccountApprovalWithDefaults

`func NewGcpMarketplaceUserAccountApprovalWithDefaults() *GcpMarketplaceUserAccountApproval`

NewGcpMarketplaceUserAccountApprovalWithDefaults instantiates a new GcpMarketplaceUserAccountApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpMarketplaceUserAccountApproval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceUserAccountApproval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceUserAccountApproval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceUserAccountApproval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *GcpMarketplaceUserAccountApproval) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GcpMarketplaceUserAccountApproval) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GcpMarketplaceUserAccountApproval) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GcpMarketplaceUserAccountApproval) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *GcpMarketplaceUserAccountApproval) GetState() GcpMarketplaceUserAccountApprovalState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GcpMarketplaceUserAccountApproval) GetStateOk() (*GcpMarketplaceUserAccountApprovalState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GcpMarketplaceUserAccountApproval) SetState(v GcpMarketplaceUserAccountApprovalState)`

SetState sets State field to given value.

### HasState

`func (o *GcpMarketplaceUserAccountApproval) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GcpMarketplaceUserAccountApproval) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GcpMarketplaceUserAccountApproval) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GcpMarketplaceUserAccountApproval) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GcpMarketplaceUserAccountApproval) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


