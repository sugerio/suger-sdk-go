# AzureMarketplaceSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** | Date-time string | [optional] 
**DeprecationSchedule** | Pointer to [**AzureMarketplaceDeprecationSchedule**](AzureMarketplaceDeprecationSchedule.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**AzureMarketplaceResourceTarget**](AzureMarketplaceResourceTarget.md) |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceSubmission

`func NewAzureMarketplaceSubmission() *AzureMarketplaceSubmission`

NewAzureMarketplaceSubmission instantiates a new AzureMarketplaceSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceSubmissionWithDefaults

`func NewAzureMarketplaceSubmissionWithDefaults() *AzureMarketplaceSubmission`

NewAzureMarketplaceSubmissionWithDefaults instantiates a new AzureMarketplaceSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceSubmission) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceSubmission) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceSubmission) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceSubmission) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCreated

`func (o *AzureMarketplaceSubmission) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AzureMarketplaceSubmission) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AzureMarketplaceSubmission) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AzureMarketplaceSubmission) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeprecationSchedule

`func (o *AzureMarketplaceSubmission) GetDeprecationSchedule() AzureMarketplaceDeprecationSchedule`

GetDeprecationSchedule returns the DeprecationSchedule field if non-nil, zero value otherwise.

### GetDeprecationScheduleOk

`func (o *AzureMarketplaceSubmission) GetDeprecationScheduleOk() (*AzureMarketplaceDeprecationSchedule, bool)`

GetDeprecationScheduleOk returns a tuple with the DeprecationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationSchedule

`func (o *AzureMarketplaceSubmission) SetDeprecationSchedule(v AzureMarketplaceDeprecationSchedule)`

SetDeprecationSchedule sets DeprecationSchedule field to given value.

### HasDeprecationSchedule

`func (o *AzureMarketplaceSubmission) HasDeprecationSchedule() bool`

HasDeprecationSchedule returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceSubmission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceSubmission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceSubmission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceSubmission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplaceSubmission) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplaceSubmission) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplaceSubmission) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplaceSubmission) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceSubmission) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceSubmission) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceSubmission) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceSubmission) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceSubmission) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceSubmission) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceSubmission) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceSubmission) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResult

`func (o *AzureMarketplaceSubmission) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AzureMarketplaceSubmission) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AzureMarketplaceSubmission) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AzureMarketplaceSubmission) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *AzureMarketplaceSubmission) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AzureMarketplaceSubmission) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AzureMarketplaceSubmission) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AzureMarketplaceSubmission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *AzureMarketplaceSubmission) GetTarget() AzureMarketplaceResourceTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AzureMarketplaceSubmission) GetTargetOk() (*AzureMarketplaceResourceTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AzureMarketplaceSubmission) SetTarget(v AzureMarketplaceResourceTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AzureMarketplaceSubmission) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceSubmission) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceSubmission) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceSubmission) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceSubmission) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


