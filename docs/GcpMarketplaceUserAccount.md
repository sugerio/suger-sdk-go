# GcpMarketplaceUserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]GcpMarketplaceUserAccountApproval**](GcpMarketplaceUserAccountApproval.md) | The approvals for this account, that are permitted or have been completed. | [optional] 
**BillingAccountId** | Pointer to **string** | The buyer&#39;s GCP billing account ID. | [optional] 
**CreateTime** | Pointer to **time.Time** | RFC3339 UTC timestamp | [optional] 
**Id** | Pointer to **string** | GCP Marketplace User Account ID. | [optional] 
**InputProperties** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **string** | The resource name of the account. Account names have the form providers/{provider_id}/accounts/{account_id}. | [optional] 
**Provider** | Pointer to **string** | The identifier of the service provider (SaaS Seller) that this account was created against. | [optional] 
**State** | Pointer to [**GcpMarketplaceUserAccountState**](GcpMarketplaceUserAccountState.md) | The state of the account. An account might not be able to make a purchase if the billing account is suspended. | [optional] 
**UpdateTime** | Pointer to **time.Time** | RFC3339 UTC timestamp | [optional] 
**UserInfo** | Pointer to [**GcpUserInfo**](GcpUserInfo.md) |  | [optional] 

## Methods

### NewGcpMarketplaceUserAccount

`func NewGcpMarketplaceUserAccount() *GcpMarketplaceUserAccount`

NewGcpMarketplaceUserAccount instantiates a new GcpMarketplaceUserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceUserAccountWithDefaults

`func NewGcpMarketplaceUserAccountWithDefaults() *GcpMarketplaceUserAccount`

NewGcpMarketplaceUserAccountWithDefaults instantiates a new GcpMarketplaceUserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *GcpMarketplaceUserAccount) GetApprovals() []GcpMarketplaceUserAccountApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *GcpMarketplaceUserAccount) GetApprovalsOk() (*[]GcpMarketplaceUserAccountApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *GcpMarketplaceUserAccount) SetApprovals(v []GcpMarketplaceUserAccountApproval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *GcpMarketplaceUserAccount) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetBillingAccountId

`func (o *GcpMarketplaceUserAccount) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *GcpMarketplaceUserAccount) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *GcpMarketplaceUserAccount) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *GcpMarketplaceUserAccount) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetCreateTime

`func (o *GcpMarketplaceUserAccount) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GcpMarketplaceUserAccount) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GcpMarketplaceUserAccount) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GcpMarketplaceUserAccount) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *GcpMarketplaceUserAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpMarketplaceUserAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpMarketplaceUserAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GcpMarketplaceUserAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputProperties

`func (o *GcpMarketplaceUserAccount) GetInputProperties() []int32`

GetInputProperties returns the InputProperties field if non-nil, zero value otherwise.

### GetInputPropertiesOk

`func (o *GcpMarketplaceUserAccount) GetInputPropertiesOk() (*[]int32, bool)`

GetInputPropertiesOk returns a tuple with the InputProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputProperties

`func (o *GcpMarketplaceUserAccount) SetInputProperties(v []int32)`

SetInputProperties sets InputProperties field to given value.

### HasInputProperties

`func (o *GcpMarketplaceUserAccount) HasInputProperties() bool`

HasInputProperties returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceUserAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceUserAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceUserAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceUserAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *GcpMarketplaceUserAccount) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GcpMarketplaceUserAccount) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GcpMarketplaceUserAccount) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GcpMarketplaceUserAccount) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetState

`func (o *GcpMarketplaceUserAccount) GetState() GcpMarketplaceUserAccountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GcpMarketplaceUserAccount) GetStateOk() (*GcpMarketplaceUserAccountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GcpMarketplaceUserAccount) SetState(v GcpMarketplaceUserAccountState)`

SetState sets State field to given value.

### HasState

`func (o *GcpMarketplaceUserAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GcpMarketplaceUserAccount) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GcpMarketplaceUserAccount) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GcpMarketplaceUserAccount) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GcpMarketplaceUserAccount) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUserInfo

`func (o *GcpMarketplaceUserAccount) GetUserInfo() GcpUserInfo`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *GcpMarketplaceUserAccount) GetUserInfoOk() (*GcpUserInfo, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *GcpMarketplaceUserAccount) SetUserInfo(v GcpUserInfo)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *GcpMarketplaceUserAccount) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


