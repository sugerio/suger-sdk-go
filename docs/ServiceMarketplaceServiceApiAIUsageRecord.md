# ServiceMarketplaceServiceApiAIUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostMicrocents** | Pointer to **int32** | CostMicrocents is the total cost in microcents. | [optional] 
**CostUsd** | Pointer to **float32** | CostUSD is the total cost in USD. | [optional] 
**HourBucket** | Pointer to **string** | HourBucket is the hour this record represents. | [optional] 
**InputTokens** | Pointer to **int32** | InputTokens is the total input tokens for this hour. | [optional] 
**IsByok** | Pointer to **bool** | IsBYOK indicates if the organization&#39;s own API key was used. | [optional] 
**Model** | Pointer to **string** | Model is the model name. | [optional] 
**OutputTokens** | Pointer to **int32** | OutputTokens is the total output tokens for this hour. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider (openai, anthropic, gemini). | [optional] 
**RequestCount** | Pointer to **int32** | RequestCount is the number of requests in this hour. | [optional] 
**UserId** | Pointer to **string** | UserID is the user who made the requests (empty if unattributed). | [optional] 

## Methods

### NewServiceMarketplaceServiceApiAIUsageRecord

`func NewServiceMarketplaceServiceApiAIUsageRecord() *ServiceMarketplaceServiceApiAIUsageRecord`

NewServiceMarketplaceServiceApiAIUsageRecord instantiates a new ServiceMarketplaceServiceApiAIUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMarketplaceServiceApiAIUsageRecordWithDefaults

`func NewServiceMarketplaceServiceApiAIUsageRecordWithDefaults() *ServiceMarketplaceServiceApiAIUsageRecord`

NewServiceMarketplaceServiceApiAIUsageRecordWithDefaults instantiates a new ServiceMarketplaceServiceApiAIUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostMicrocents

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetCostMicrocents() int32`

GetCostMicrocents returns the CostMicrocents field if non-nil, zero value otherwise.

### GetCostMicrocentsOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetCostMicrocentsOk() (*int32, bool)`

GetCostMicrocentsOk returns a tuple with the CostMicrocents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMicrocents

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetCostMicrocents(v int32)`

SetCostMicrocents sets CostMicrocents field to given value.

### HasCostMicrocents

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasCostMicrocents() bool`

HasCostMicrocents returns a boolean if a field has been set.

### GetCostUsd

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetHourBucket

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetHourBucket() string`

GetHourBucket returns the HourBucket field if non-nil, zero value otherwise.

### GetHourBucketOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetHourBucketOk() (*string, bool)`

GetHourBucketOk returns a tuple with the HourBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourBucket

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetHourBucket(v string)`

SetHourBucket sets HourBucket field to given value.

### HasHourBucket

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasHourBucket() bool`

HasHourBucket returns a boolean if a field has been set.

### GetInputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetIsByok

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetIsByok() bool`

GetIsByok returns the IsByok field if non-nil, zero value otherwise.

### GetIsByokOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetIsByokOk() (*bool, bool)`

GetIsByokOk returns a tuple with the IsByok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByok

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetIsByok(v bool)`

SetIsByok sets IsByok field to given value.

### HasIsByok

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasIsByok() bool`

HasIsByok returns a boolean if a field has been set.

### GetModel

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOutputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetProvider

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequestCount

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### GetUserId

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ServiceMarketplaceServiceApiAIUsageRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


