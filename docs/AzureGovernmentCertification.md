# AzureGovernmentCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**ValidationResults** | Pointer to [**[]AzureValidationResult**](AzureValidationResult.md) |  | [optional] 

## Methods

### NewAzureGovernmentCertification

`func NewAzureGovernmentCertification() *AzureGovernmentCertification`

NewAzureGovernmentCertification instantiates a new AzureGovernmentCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureGovernmentCertificationWithDefaults

`func NewAzureGovernmentCertificationWithDefaults() *AzureGovernmentCertification`

NewAzureGovernmentCertificationWithDefaults instantiates a new AzureGovernmentCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AzureGovernmentCertification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AzureGovernmentCertification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AzureGovernmentCertification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AzureGovernmentCertification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUri

`func (o *AzureGovernmentCertification) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AzureGovernmentCertification) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AzureGovernmentCertification) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AzureGovernmentCertification) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetValidationResults

`func (o *AzureGovernmentCertification) GetValidationResults() []AzureValidationResult`

GetValidationResults returns the ValidationResults field if non-nil, zero value otherwise.

### GetValidationResultsOk

`func (o *AzureGovernmentCertification) GetValidationResultsOk() (*[]AzureValidationResult, bool)`

GetValidationResultsOk returns a tuple with the ValidationResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResults

`func (o *AzureGovernmentCertification) SetValidationResults(v []AzureValidationResult)`

SetValidationResults sets ValidationResults field to given value.

### HasValidationResults

`func (o *AzureGovernmentCertification) HasValidationResults() bool`

HasValidationResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


