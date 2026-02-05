# AwsProductDeliveryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiAlias** | Pointer to **string** | Exclusive Fields For AWS AMI Product | [optional] 
**ApiType** | Pointer to **string** |  | [optional] 
**CompatibleServices** | Pointer to **[]string** |  | [optional] 
**Endpoints** | Pointer to [**[]AwsProductDeliveryOptionApiEndpoint**](AwsProductDeliveryOptionApiEndpoint.md) |  | [optional] 
**FulfillmentUrl** | Pointer to **string** | Exclusive Fields For AWS SaaS Product | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**QuickLaunchEnabled** | Pointer to **bool** |  | [optional] 
**Recommendations** | Pointer to **map[string]interface{}** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | Exclusive Fields For AWS Container Product | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsageInstructions** | Pointer to **string** | Exclusive Fields For SaaS AI Product | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsProductDeliveryOption

`func NewAwsProductDeliveryOption() *AwsProductDeliveryOption`

NewAwsProductDeliveryOption instantiates a new AwsProductDeliveryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProductDeliveryOptionWithDefaults

`func NewAwsProductDeliveryOptionWithDefaults() *AwsProductDeliveryOption`

NewAwsProductDeliveryOptionWithDefaults instantiates a new AwsProductDeliveryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiAlias

`func (o *AwsProductDeliveryOption) GetAmiAlias() string`

GetAmiAlias returns the AmiAlias field if non-nil, zero value otherwise.

### GetAmiAliasOk

`func (o *AwsProductDeliveryOption) GetAmiAliasOk() (*string, bool)`

GetAmiAliasOk returns a tuple with the AmiAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiAlias

`func (o *AwsProductDeliveryOption) SetAmiAlias(v string)`

SetAmiAlias sets AmiAlias field to given value.

### HasAmiAlias

`func (o *AwsProductDeliveryOption) HasAmiAlias() bool`

HasAmiAlias returns a boolean if a field has been set.

### GetApiType

`func (o *AwsProductDeliveryOption) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *AwsProductDeliveryOption) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *AwsProductDeliveryOption) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *AwsProductDeliveryOption) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetCompatibleServices

`func (o *AwsProductDeliveryOption) GetCompatibleServices() []string`

GetCompatibleServices returns the CompatibleServices field if non-nil, zero value otherwise.

### GetCompatibleServicesOk

`func (o *AwsProductDeliveryOption) GetCompatibleServicesOk() (*[]string, bool)`

GetCompatibleServicesOk returns a tuple with the CompatibleServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleServices

`func (o *AwsProductDeliveryOption) SetCompatibleServices(v []string)`

SetCompatibleServices sets CompatibleServices field to given value.

### HasCompatibleServices

`func (o *AwsProductDeliveryOption) HasCompatibleServices() bool`

HasCompatibleServices returns a boolean if a field has been set.

### GetEndpoints

`func (o *AwsProductDeliveryOption) GetEndpoints() []AwsProductDeliveryOptionApiEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *AwsProductDeliveryOption) GetEndpointsOk() (*[]AwsProductDeliveryOptionApiEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *AwsProductDeliveryOption) SetEndpoints(v []AwsProductDeliveryOptionApiEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *AwsProductDeliveryOption) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetFulfillmentUrl

`func (o *AwsProductDeliveryOption) GetFulfillmentUrl() string`

GetFulfillmentUrl returns the FulfillmentUrl field if non-nil, zero value otherwise.

### GetFulfillmentUrlOk

`func (o *AwsProductDeliveryOption) GetFulfillmentUrlOk() (*string, bool)`

GetFulfillmentUrlOk returns a tuple with the FulfillmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUrl

`func (o *AwsProductDeliveryOption) SetFulfillmentUrl(v string)`

SetFulfillmentUrl sets FulfillmentUrl field to given value.

### HasFulfillmentUrl

`func (o *AwsProductDeliveryOption) HasFulfillmentUrl() bool`

HasFulfillmentUrl returns a boolean if a field has been set.

### GetId

`func (o *AwsProductDeliveryOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsProductDeliveryOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsProductDeliveryOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsProductDeliveryOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuickLaunchEnabled

`func (o *AwsProductDeliveryOption) GetQuickLaunchEnabled() bool`

GetQuickLaunchEnabled returns the QuickLaunchEnabled field if non-nil, zero value otherwise.

### GetQuickLaunchEnabledOk

`func (o *AwsProductDeliveryOption) GetQuickLaunchEnabledOk() (*bool, bool)`

GetQuickLaunchEnabledOk returns a tuple with the QuickLaunchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickLaunchEnabled

`func (o *AwsProductDeliveryOption) SetQuickLaunchEnabled(v bool)`

SetQuickLaunchEnabled sets QuickLaunchEnabled field to given value.

### HasQuickLaunchEnabled

`func (o *AwsProductDeliveryOption) HasQuickLaunchEnabled() bool`

HasQuickLaunchEnabled returns a boolean if a field has been set.

### GetRecommendations

`func (o *AwsProductDeliveryOption) GetRecommendations() map[string]interface{}`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *AwsProductDeliveryOption) GetRecommendationsOk() (*map[string]interface{}, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *AwsProductDeliveryOption) SetRecommendations(v map[string]interface{})`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *AwsProductDeliveryOption) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetShortDescription

`func (o *AwsProductDeliveryOption) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AwsProductDeliveryOption) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AwsProductDeliveryOption) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *AwsProductDeliveryOption) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSourceId

`func (o *AwsProductDeliveryOption) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AwsProductDeliveryOption) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AwsProductDeliveryOption) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AwsProductDeliveryOption) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetTitle

`func (o *AwsProductDeliveryOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AwsProductDeliveryOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AwsProductDeliveryOption) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AwsProductDeliveryOption) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AwsProductDeliveryOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsProductDeliveryOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsProductDeliveryOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AwsProductDeliveryOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsageInstructions

`func (o *AwsProductDeliveryOption) GetUsageInstructions() string`

GetUsageInstructions returns the UsageInstructions field if non-nil, zero value otherwise.

### GetUsageInstructionsOk

`func (o *AwsProductDeliveryOption) GetUsageInstructionsOk() (*string, bool)`

GetUsageInstructionsOk returns a tuple with the UsageInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageInstructions

`func (o *AwsProductDeliveryOption) SetUsageInstructions(v string)`

SetUsageInstructions sets UsageInstructions field to given value.

### HasUsageInstructions

`func (o *AwsProductDeliveryOption) HasUsageInstructions() bool`

HasUsageInstructions returns a boolean if a field has been set.

### GetVisibility

`func (o *AwsProductDeliveryOption) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AwsProductDeliveryOption) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AwsProductDeliveryOption) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AwsProductDeliveryOption) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


