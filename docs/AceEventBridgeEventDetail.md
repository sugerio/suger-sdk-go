# AceEventBridgeEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** | \&quot;AWS\&quot; or \&quot;Sandbox\&quot; | [optional] 
**EngagementInvitation** | Pointer to [**AceEventEngagementInvitation**](AceEventEngagementInvitation.md) | Engagement invitation-related events | [optional] 
**Opportunity** | Pointer to [**AceEventOpportunity**](AceEventOpportunity.md) | Opportunity-related events | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewAceEventBridgeEventDetail

`func NewAceEventBridgeEventDetail() *AceEventBridgeEventDetail`

NewAceEventBridgeEventDetail instantiates a new AceEventBridgeEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAceEventBridgeEventDetailWithDefaults

`func NewAceEventBridgeEventDetailWithDefaults() *AceEventBridgeEventDetail`

NewAceEventBridgeEventDetailWithDefaults instantiates a new AceEventBridgeEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *AceEventBridgeEventDetail) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AceEventBridgeEventDetail) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AceEventBridgeEventDetail) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AceEventBridgeEventDetail) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetEngagementInvitation

`func (o *AceEventBridgeEventDetail) GetEngagementInvitation() AceEventEngagementInvitation`

GetEngagementInvitation returns the EngagementInvitation field if non-nil, zero value otherwise.

### GetEngagementInvitationOk

`func (o *AceEventBridgeEventDetail) GetEngagementInvitationOk() (*AceEventEngagementInvitation, bool)`

GetEngagementInvitationOk returns a tuple with the EngagementInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementInvitation

`func (o *AceEventBridgeEventDetail) SetEngagementInvitation(v AceEventEngagementInvitation)`

SetEngagementInvitation sets EngagementInvitation field to given value.

### HasEngagementInvitation

`func (o *AceEventBridgeEventDetail) HasEngagementInvitation() bool`

HasEngagementInvitation returns a boolean if a field has been set.

### GetOpportunity

`func (o *AceEventBridgeEventDetail) GetOpportunity() AceEventOpportunity`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *AceEventBridgeEventDetail) GetOpportunityOk() (*AceEventOpportunity, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *AceEventBridgeEventDetail) SetOpportunity(v AceEventOpportunity)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *AceEventBridgeEventDetail) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *AceEventBridgeEventDetail) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AceEventBridgeEventDetail) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AceEventBridgeEventDetail) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AceEventBridgeEventDetail) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


