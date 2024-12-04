# AzureMarketplaceMeteringUsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | Dimension identifier | [optional] 
**EffectiveStartTime** | Pointer to **string** | Time in UTC when the usage event occurred | [optional] 
**PlanId** | Pointer to **string** | Plan associated with the purchased offer | [optional] 
**Quantity** | Pointer to **float32** | Number of units consumed | [optional] 
**ResourceId** | Pointer to **string** | subscriptionId property value for SaaS offer subscriptions; resourceUsageId property on the managed application resource for managed application offers. For managed applications, only use one of resourceId or resourceUri. | [optional] 
**ResourceUri** | Pointer to **string** | Resource URI for the managed app. Used with managed applications. Only use resourceUri or resourceId, but never both. | [optional] 

## Methods

### NewAzureMarketplaceMeteringUsageEvent

`func NewAzureMarketplaceMeteringUsageEvent() *AzureMarketplaceMeteringUsageEvent`

NewAzureMarketplaceMeteringUsageEvent instantiates a new AzureMarketplaceMeteringUsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceMeteringUsageEventWithDefaults

`func NewAzureMarketplaceMeteringUsageEventWithDefaults() *AzureMarketplaceMeteringUsageEvent`

NewAzureMarketplaceMeteringUsageEventWithDefaults instantiates a new AzureMarketplaceMeteringUsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *AzureMarketplaceMeteringUsageEvent) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *AzureMarketplaceMeteringUsageEvent) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *AzureMarketplaceMeteringUsageEvent) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetEffectiveStartTime

`func (o *AzureMarketplaceMeteringUsageEvent) GetEffectiveStartTime() string`

GetEffectiveStartTime returns the EffectiveStartTime field if non-nil, zero value otherwise.

### GetEffectiveStartTimeOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetEffectiveStartTimeOk() (*string, bool)`

GetEffectiveStartTimeOk returns a tuple with the EffectiveStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStartTime

`func (o *AzureMarketplaceMeteringUsageEvent) SetEffectiveStartTime(v string)`

SetEffectiveStartTime sets EffectiveStartTime field to given value.

### HasEffectiveStartTime

`func (o *AzureMarketplaceMeteringUsageEvent) HasEffectiveStartTime() bool`

HasEffectiveStartTime returns a boolean if a field has been set.

### GetPlanId

`func (o *AzureMarketplaceMeteringUsageEvent) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AzureMarketplaceMeteringUsageEvent) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AzureMarketplaceMeteringUsageEvent) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *AzureMarketplaceMeteringUsageEvent) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AzureMarketplaceMeteringUsageEvent) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AzureMarketplaceMeteringUsageEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetResourceId

`func (o *AzureMarketplaceMeteringUsageEvent) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AzureMarketplaceMeteringUsageEvent) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AzureMarketplaceMeteringUsageEvent) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceUri

`func (o *AzureMarketplaceMeteringUsageEvent) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *AzureMarketplaceMeteringUsageEvent) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *AzureMarketplaceMeteringUsageEvent) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *AzureMarketplaceMeteringUsageEvent) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


