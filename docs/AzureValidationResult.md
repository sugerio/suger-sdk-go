# AzureValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** |  | [optional] 
**MemberNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAzureValidationResult

`func NewAzureValidationResult() *AzureValidationResult`

NewAzureValidationResult instantiates a new AzureValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureValidationResultWithDefaults

`func NewAzureValidationResultWithDefaults() *AzureValidationResult`

NewAzureValidationResultWithDefaults instantiates a new AzureValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *AzureValidationResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AzureValidationResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AzureValidationResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AzureValidationResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMemberNames

`func (o *AzureValidationResult) GetMemberNames() []string`

GetMemberNames returns the MemberNames field if non-nil, zero value otherwise.

### GetMemberNamesOk

`func (o *AzureValidationResult) GetMemberNamesOk() (*[]string, bool)`

GetMemberNamesOk returns a tuple with the MemberNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberNames

`func (o *AzureValidationResult) SetMemberNames(v []string)`

SetMemberNames sets MemberNames field to given value.

### HasMemberNames

`func (o *AzureValidationResult) HasMemberNames() bool`

HasMemberNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


