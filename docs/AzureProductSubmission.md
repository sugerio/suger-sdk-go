# AzureProductSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreResourcesReady** | Pointer to **bool** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PendingUpdateInfo** | Pointer to [**AzurePendingUpdateInfo**](AzurePendingUpdateInfo.md) |  | [optional] 
**PublishedTimeInUtc** | Pointer to **time.Time** |  | [optional] 
**ReleaseNumber** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubState** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**VariantResources** | Pointer to [**[]AzureVariantResource**](AzureVariantResource.md) |  | [optional] 

## Methods

### NewAzureProductSubmission

`func NewAzureProductSubmission() *AzureProductSubmission`

NewAzureProductSubmission instantiates a new AzureProductSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductSubmissionWithDefaults

`func NewAzureProductSubmissionWithDefaults() *AzureProductSubmission`

NewAzureProductSubmissionWithDefaults instantiates a new AzureProductSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreResourcesReady

`func (o *AzureProductSubmission) GetAreResourcesReady() bool`

GetAreResourcesReady returns the AreResourcesReady field if non-nil, zero value otherwise.

### GetAreResourcesReadyOk

`func (o *AzureProductSubmission) GetAreResourcesReadyOk() (*bool, bool)`

GetAreResourcesReadyOk returns a tuple with the AreResourcesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreResourcesReady

`func (o *AzureProductSubmission) SetAreResourcesReady(v bool)`

SetAreResourcesReady sets AreResourcesReady field to given value.

### HasAreResourcesReady

`func (o *AzureProductSubmission) HasAreResourcesReady() bool`

HasAreResourcesReady returns a boolean if a field has been set.

### GetFriendlyName

`func (o *AzureProductSubmission) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AzureProductSubmission) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AzureProductSubmission) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AzureProductSubmission) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *AzureProductSubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProductSubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProductSubmission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProductSubmission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPendingUpdateInfo

`func (o *AzureProductSubmission) GetPendingUpdateInfo() AzurePendingUpdateInfo`

GetPendingUpdateInfo returns the PendingUpdateInfo field if non-nil, zero value otherwise.

### GetPendingUpdateInfoOk

`func (o *AzureProductSubmission) GetPendingUpdateInfoOk() (*AzurePendingUpdateInfo, bool)`

GetPendingUpdateInfoOk returns a tuple with the PendingUpdateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateInfo

`func (o *AzureProductSubmission) SetPendingUpdateInfo(v AzurePendingUpdateInfo)`

SetPendingUpdateInfo sets PendingUpdateInfo field to given value.

### HasPendingUpdateInfo

`func (o *AzureProductSubmission) HasPendingUpdateInfo() bool`

HasPendingUpdateInfo returns a boolean if a field has been set.

### GetPublishedTimeInUtc

`func (o *AzureProductSubmission) GetPublishedTimeInUtc() time.Time`

GetPublishedTimeInUtc returns the PublishedTimeInUtc field if non-nil, zero value otherwise.

### GetPublishedTimeInUtcOk

`func (o *AzureProductSubmission) GetPublishedTimeInUtcOk() (*time.Time, bool)`

GetPublishedTimeInUtcOk returns a tuple with the PublishedTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedTimeInUtc

`func (o *AzureProductSubmission) SetPublishedTimeInUtc(v time.Time)`

SetPublishedTimeInUtc sets PublishedTimeInUtc field to given value.

### HasPublishedTimeInUtc

`func (o *AzureProductSubmission) HasPublishedTimeInUtc() bool`

HasPublishedTimeInUtc returns a boolean if a field has been set.

### GetReleaseNumber

`func (o *AzureProductSubmission) GetReleaseNumber() int32`

GetReleaseNumber returns the ReleaseNumber field if non-nil, zero value otherwise.

### GetReleaseNumberOk

`func (o *AzureProductSubmission) GetReleaseNumberOk() (*int32, bool)`

GetReleaseNumberOk returns a tuple with the ReleaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNumber

`func (o *AzureProductSubmission) SetReleaseNumber(v int32)`

SetReleaseNumber sets ReleaseNumber field to given value.

### HasReleaseNumber

`func (o *AzureProductSubmission) HasReleaseNumber() bool`

HasReleaseNumber returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProductSubmission) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProductSubmission) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProductSubmission) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProductSubmission) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResources

`func (o *AzureProductSubmission) GetResources() []AzureTypeValue`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AzureProductSubmission) GetResourcesOk() (*[]AzureTypeValue, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AzureProductSubmission) SetResources(v []AzureTypeValue)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AzureProductSubmission) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetState

`func (o *AzureProductSubmission) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzureProductSubmission) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzureProductSubmission) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AzureProductSubmission) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubState

`func (o *AzureProductSubmission) GetSubState() string`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *AzureProductSubmission) GetSubStateOk() (*string, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *AzureProductSubmission) SetSubState(v string)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *AzureProductSubmission) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetTargets

`func (o *AzureProductSubmission) GetTargets() []AzureTypeValue`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AzureProductSubmission) GetTargetsOk() (*[]AzureTypeValue, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AzureProductSubmission) SetTargets(v []AzureTypeValue)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AzureProductSubmission) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetVariantResources

`func (o *AzureProductSubmission) GetVariantResources() []AzureVariantResource`

GetVariantResources returns the VariantResources field if non-nil, zero value otherwise.

### GetVariantResourcesOk

`func (o *AzureProductSubmission) GetVariantResourcesOk() (*[]AzureVariantResource, bool)`

GetVariantResourcesOk returns a tuple with the VariantResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantResources

`func (o *AzureProductSubmission) SetVariantResources(v []AzureVariantResource)`

SetVariantResources sets VariantResources field to given value.

### HasVariantResources

`func (o *AzureProductSubmission) HasVariantResources() bool`

HasVariantResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


