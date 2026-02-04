# AzureMarketplaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AzureMarketplaceEventAction**](AzureMarketplaceEventAction.md) |  | [optional] 
**ActivityId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The Operation Id. | [optional] 
**OfferId** | Pointer to **string** |  | [optional] 
**OperationRequestSource** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PublisherId** | Pointer to **string** |  | [optional] 
**PurchaseToken** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subscription** | Pointer to [**AzureMarketplaceSubscription**](AzureMarketplaceSubscription.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SugerOrganizationId** | Pointer to **string** | Populated by Suger Service. | [optional] 
**TimeStamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAzureMarketplaceEvent

`func NewAzureMarketplaceEvent() *AzureMarketplaceEvent`

NewAzureMarketplaceEvent instantiates a new AzureMarketplaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceEventWithDefaults

`func NewAzureMarketplaceEventWithDefaults() *AzureMarketplaceEvent`

NewAzureMarketplaceEventWithDefaults instantiates a new AzureMarketplaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AzureMarketplaceEvent) GetAction() AzureMarketplaceEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AzureMarketplaceEvent) GetActionOk() (*AzureMarketplaceEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AzureMarketplaceEvent) SetAction(v AzureMarketplaceEventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *AzureMarketplaceEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActivityId

`func (o *AzureMarketplaceEvent) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *AzureMarketplaceEvent) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *AzureMarketplaceEvent) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *AzureMarketplaceEvent) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOfferId

`func (o *AzureMarketplaceEvent) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *AzureMarketplaceEvent) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *AzureMarketplaceEvent) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *AzureMarketplaceEvent) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetOperationRequestSource

`func (o *AzureMarketplaceEvent) GetOperationRequestSource() string`

GetOperationRequestSource returns the OperationRequestSource field if non-nil, zero value otherwise.

### GetOperationRequestSourceOk

`func (o *AzureMarketplaceEvent) GetOperationRequestSourceOk() (*string, bool)`

GetOperationRequestSourceOk returns a tuple with the OperationRequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationRequestSource

`func (o *AzureMarketplaceEvent) SetOperationRequestSource(v string)`

SetOperationRequestSource sets OperationRequestSource field to given value.

### HasOperationRequestSource

`func (o *AzureMarketplaceEvent) HasOperationRequestSource() bool`

HasOperationRequestSource returns a boolean if a field has been set.

### GetPlanId

`func (o *AzureMarketplaceEvent) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AzureMarketplaceEvent) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AzureMarketplaceEvent) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AzureMarketplaceEvent) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPublisherId

`func (o *AzureMarketplaceEvent) GetPublisherId() string`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *AzureMarketplaceEvent) GetPublisherIdOk() (*string, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *AzureMarketplaceEvent) SetPublisherId(v string)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *AzureMarketplaceEvent) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetPurchaseToken

`func (o *AzureMarketplaceEvent) GetPurchaseToken() string`

GetPurchaseToken returns the PurchaseToken field if non-nil, zero value otherwise.

### GetPurchaseTokenOk

`func (o *AzureMarketplaceEvent) GetPurchaseTokenOk() (*string, bool)`

GetPurchaseTokenOk returns a tuple with the PurchaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseToken

`func (o *AzureMarketplaceEvent) SetPurchaseToken(v string)`

SetPurchaseToken sets PurchaseToken field to given value.

### HasPurchaseToken

`func (o *AzureMarketplaceEvent) HasPurchaseToken() bool`

HasPurchaseToken returns a boolean if a field has been set.

### GetQuantity

`func (o *AzureMarketplaceEvent) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AzureMarketplaceEvent) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AzureMarketplaceEvent) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AzureMarketplaceEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *AzureMarketplaceEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AzureMarketplaceEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AzureMarketplaceEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AzureMarketplaceEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *AzureMarketplaceEvent) GetSubscription() AzureMarketplaceSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AzureMarketplaceEvent) GetSubscriptionOk() (*AzureMarketplaceSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AzureMarketplaceEvent) SetSubscription(v AzureMarketplaceSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *AzureMarketplaceEvent) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureMarketplaceEvent) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureMarketplaceEvent) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureMarketplaceEvent) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureMarketplaceEvent) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSugerOrganizationId

`func (o *AzureMarketplaceEvent) GetSugerOrganizationId() string`

GetSugerOrganizationId returns the SugerOrganizationId field if non-nil, zero value otherwise.

### GetSugerOrganizationIdOk

`func (o *AzureMarketplaceEvent) GetSugerOrganizationIdOk() (*string, bool)`

GetSugerOrganizationIdOk returns a tuple with the SugerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerOrganizationId

`func (o *AzureMarketplaceEvent) SetSugerOrganizationId(v string)`

SetSugerOrganizationId sets SugerOrganizationId field to given value.

### HasSugerOrganizationId

`func (o *AzureMarketplaceEvent) HasSugerOrganizationId() bool`

HasSugerOrganizationId returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AzureMarketplaceEvent) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AzureMarketplaceEvent) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AzureMarketplaceEvent) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AzureMarketplaceEvent) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


