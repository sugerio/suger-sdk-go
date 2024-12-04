# GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | Dimension identifier | [optional] 
**EffectiveStartTime** | Pointer to **string** | Time in UTC when the usage event occurred | [optional] 
**MessageTime** | Pointer to **string** | Time this message was created in UTC | [optional] 
**PlanId** | Pointer to **string** | Plan associated with the purchased offer | [optional] 
**Quantity** | Pointer to **float32** | Number of units consumed | [optional] 
**ResourceId** | Pointer to **string** | Identifier of the resource against which usage is emitted | [optional] 
**ResourceUri** | Pointer to **string** | Identifier of the managed app resource against which usage is emitted | [optional] 
**Status** | Pointer to [**GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum**](GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum.md) | Status of the operation. | [optional] 
**UsageEventId** | Pointer to **string** | Unique identifier associated with the usage event | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse

`func NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse() *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse`

NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse instantiates a new GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponseWithDefaults

`func NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponseWithDefaults() *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse`

NewGithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponseWithDefaults instantiates a new GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetEffectiveStartTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetEffectiveStartTime() string`

GetEffectiveStartTime returns the EffectiveStartTime field if non-nil, zero value otherwise.

### GetEffectiveStartTimeOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetEffectiveStartTimeOk() (*string, bool)`

GetEffectiveStartTimeOk returns a tuple with the EffectiveStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStartTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetEffectiveStartTime(v string)`

SetEffectiveStartTime sets EffectiveStartTime field to given value.

### HasEffectiveStartTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasEffectiveStartTime() bool`

HasEffectiveStartTime returns a boolean if a field has been set.

### GetMessageTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetMessageTime() string`

GetMessageTime returns the MessageTime field if non-nil, zero value otherwise.

### GetMessageTimeOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetMessageTimeOk() (*string, bool)`

GetMessageTimeOk returns a tuple with the MessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetMessageTime(v string)`

SetMessageTime sets MessageTime field to given value.

### HasMessageTime

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasMessageTime() bool`

HasMessageTime returns a boolean if a field has been set.

### GetPlanId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetResourceId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceUri

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetStatus() GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetStatusOk() (*GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetStatus(v GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsageEventId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.

### HasUsageEventId

`func (o *GithubComSugerioMarketplaceServiceAzureSdkMarketplacemeteringv1UsageEventOkResponse) HasUsageEventId() bool`

HasUsageEventId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


