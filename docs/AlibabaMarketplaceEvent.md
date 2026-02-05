# AlibabaMarketplaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AlibabaMarketplaceAction**](AlibabaMarketplaceAction.md) |  | [optional] 
**AliUid** | Pointer to **string** | The Alibaba UID of the buyer&#39;s Alibaba Account. | [optional] 
**ExpiredOn** | Pointer to **time.Time** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**IsRefund** | Pointer to **bool** |  | [optional] 
**OrderBizId** | Pointer to **string** | Used as the Instance ID, and as the external ID for the Suger Entitlement. | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**SugerOrganizationId** | Pointer to **string** | Suger organization ID of this event. Populated by Suger Service. | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**TimeStamp** | Pointer to **time.Time** | When the event was received by the Suger. | [optional] 
**Token** | Pointer to **string** | SPI security token. | [optional] 
**Trial** | Pointer to **bool** | If true, the event is for a trial. | [optional] 

## Methods

### NewAlibabaMarketplaceEvent

`func NewAlibabaMarketplaceEvent() *AlibabaMarketplaceEvent`

NewAlibabaMarketplaceEvent instantiates a new AlibabaMarketplaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlibabaMarketplaceEventWithDefaults

`func NewAlibabaMarketplaceEventWithDefaults() *AlibabaMarketplaceEvent`

NewAlibabaMarketplaceEventWithDefaults instantiates a new AlibabaMarketplaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AlibabaMarketplaceEvent) GetAction() AlibabaMarketplaceAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AlibabaMarketplaceEvent) GetActionOk() (*AlibabaMarketplaceAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AlibabaMarketplaceEvent) SetAction(v AlibabaMarketplaceAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *AlibabaMarketplaceEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAliUid

`func (o *AlibabaMarketplaceEvent) GetAliUid() string`

GetAliUid returns the AliUid field if non-nil, zero value otherwise.

### GetAliUidOk

`func (o *AlibabaMarketplaceEvent) GetAliUidOk() (*string, bool)`

GetAliUidOk returns a tuple with the AliUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliUid

`func (o *AlibabaMarketplaceEvent) SetAliUid(v string)`

SetAliUid sets AliUid field to given value.

### HasAliUid

`func (o *AlibabaMarketplaceEvent) HasAliUid() bool`

HasAliUid returns a boolean if a field has been set.

### GetExpiredOn

`func (o *AlibabaMarketplaceEvent) GetExpiredOn() time.Time`

GetExpiredOn returns the ExpiredOn field if non-nil, zero value otherwise.

### GetExpiredOnOk

`func (o *AlibabaMarketplaceEvent) GetExpiredOnOk() (*time.Time, bool)`

GetExpiredOnOk returns a tuple with the ExpiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredOn

`func (o *AlibabaMarketplaceEvent) SetExpiredOn(v time.Time)`

SetExpiredOn sets ExpiredOn field to given value.

### HasExpiredOn

`func (o *AlibabaMarketplaceEvent) HasExpiredOn() bool`

HasExpiredOn returns a boolean if a field has been set.

### GetInstanceId

`func (o *AlibabaMarketplaceEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AlibabaMarketplaceEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AlibabaMarketplaceEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AlibabaMarketplaceEvent) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetIsRefund

`func (o *AlibabaMarketplaceEvent) GetIsRefund() bool`

GetIsRefund returns the IsRefund field if non-nil, zero value otherwise.

### GetIsRefundOk

`func (o *AlibabaMarketplaceEvent) GetIsRefundOk() (*bool, bool)`

GetIsRefundOk returns a tuple with the IsRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefund

`func (o *AlibabaMarketplaceEvent) SetIsRefund(v bool)`

SetIsRefund sets IsRefund field to given value.

### HasIsRefund

`func (o *AlibabaMarketplaceEvent) HasIsRefund() bool`

HasIsRefund returns a boolean if a field has been set.

### GetOrderBizId

`func (o *AlibabaMarketplaceEvent) GetOrderBizId() string`

GetOrderBizId returns the OrderBizId field if non-nil, zero value otherwise.

### GetOrderBizIdOk

`func (o *AlibabaMarketplaceEvent) GetOrderBizIdOk() (*string, bool)`

GetOrderBizIdOk returns a tuple with the OrderBizId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBizId

`func (o *AlibabaMarketplaceEvent) SetOrderBizId(v string)`

SetOrderBizId sets OrderBizId field to given value.

### HasOrderBizId

`func (o *AlibabaMarketplaceEvent) HasOrderBizId() bool`

HasOrderBizId returns a boolean if a field has been set.

### GetOrderId

`func (o *AlibabaMarketplaceEvent) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AlibabaMarketplaceEvent) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AlibabaMarketplaceEvent) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AlibabaMarketplaceEvent) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetProductCode

`func (o *AlibabaMarketplaceEvent) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *AlibabaMarketplaceEvent) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *AlibabaMarketplaceEvent) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *AlibabaMarketplaceEvent) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetSkuId

`func (o *AlibabaMarketplaceEvent) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *AlibabaMarketplaceEvent) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *AlibabaMarketplaceEvent) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *AlibabaMarketplaceEvent) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetSugerOrganizationId

`func (o *AlibabaMarketplaceEvent) GetSugerOrganizationId() string`

GetSugerOrganizationId returns the SugerOrganizationId field if non-nil, zero value otherwise.

### GetSugerOrganizationIdOk

`func (o *AlibabaMarketplaceEvent) GetSugerOrganizationIdOk() (*string, bool)`

GetSugerOrganizationIdOk returns a tuple with the SugerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerOrganizationId

`func (o *AlibabaMarketplaceEvent) SetSugerOrganizationId(v string)`

SetSugerOrganizationId sets SugerOrganizationId field to given value.

### HasSugerOrganizationId

`func (o *AlibabaMarketplaceEvent) HasSugerOrganizationId() bool`

HasSugerOrganizationId returns a boolean if a field has been set.

### GetTemplate

`func (o *AlibabaMarketplaceEvent) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlibabaMarketplaceEvent) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlibabaMarketplaceEvent) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AlibabaMarketplaceEvent) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimeStamp

`func (o *AlibabaMarketplaceEvent) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *AlibabaMarketplaceEvent) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *AlibabaMarketplaceEvent) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *AlibabaMarketplaceEvent) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetToken

`func (o *AlibabaMarketplaceEvent) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AlibabaMarketplaceEvent) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AlibabaMarketplaceEvent) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AlibabaMarketplaceEvent) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTrial

`func (o *AlibabaMarketplaceEvent) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *AlibabaMarketplaceEvent) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *AlibabaMarketplaceEvent) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *AlibabaMarketplaceEvent) HasTrial() bool`

HasTrial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


