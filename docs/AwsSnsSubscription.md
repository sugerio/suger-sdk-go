# AwsSnsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AwsSnsSubscriptionStatus**](AwsSnsSubscriptionStatus.md) |  | [optional] 
**SubscriptionArn** | Pointer to **string** |  | [optional] 
**TopicArn** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsSnsSubscription

`func NewAwsSnsSubscription() *AwsSnsSubscription`

NewAwsSnsSubscription instantiates a new AwsSnsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSnsSubscriptionWithDefaults

`func NewAwsSnsSubscriptionWithDefaults() *AwsSnsSubscription`

NewAwsSnsSubscriptionWithDefaults instantiates a new AwsSnsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *AwsSnsSubscription) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AwsSnsSubscription) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AwsSnsSubscription) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AwsSnsSubscription) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetProtocol

`func (o *AwsSnsSubscription) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AwsSnsSubscription) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AwsSnsSubscription) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AwsSnsSubscription) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *AwsSnsSubscription) GetStatus() AwsSnsSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsSnsSubscription) GetStatusOk() (*AwsSnsSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsSnsSubscription) SetStatus(v AwsSnsSubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsSnsSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionArn

`func (o *AwsSnsSubscription) GetSubscriptionArn() string`

GetSubscriptionArn returns the SubscriptionArn field if non-nil, zero value otherwise.

### GetSubscriptionArnOk

`func (o *AwsSnsSubscription) GetSubscriptionArnOk() (*string, bool)`

GetSubscriptionArnOk returns a tuple with the SubscriptionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionArn

`func (o *AwsSnsSubscription) SetSubscriptionArn(v string)`

SetSubscriptionArn sets SubscriptionArn field to given value.

### HasSubscriptionArn

`func (o *AwsSnsSubscription) HasSubscriptionArn() bool`

HasSubscriptionArn returns a boolean if a field has been set.

### GetTopicArn

`func (o *AwsSnsSubscription) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *AwsSnsSubscription) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *AwsSnsSubscription) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.

### HasTopicArn

`func (o *AwsSnsSubscription) HasTopicArn() bool`

HasTopicArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


