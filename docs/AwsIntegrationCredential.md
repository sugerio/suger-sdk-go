# AwsIntegrationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | The access key ID of the IAM user for Suger service to access the client&#39;s AWS services. | [optional] 
**IamUserArn** | Pointer to **string** | The ARN of the IAM user for Suger service to access the client&#39;s AWS services. | [optional] 
**SecretAccessKey** | Pointer to **string** | The secret access key of the IAM user for Suger service to access the client&#39;s AWS services. | [optional] 

## Methods

### NewAwsIntegrationCredential

`func NewAwsIntegrationCredential() *AwsIntegrationCredential`

NewAwsIntegrationCredential instantiates a new AwsIntegrationCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIntegrationCredentialWithDefaults

`func NewAwsIntegrationCredentialWithDefaults() *AwsIntegrationCredential`

NewAwsIntegrationCredentialWithDefaults instantiates a new AwsIntegrationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *AwsIntegrationCredential) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AwsIntegrationCredential) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AwsIntegrationCredential) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AwsIntegrationCredential) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetIamUserArn

`func (o *AwsIntegrationCredential) GetIamUserArn() string`

GetIamUserArn returns the IamUserArn field if non-nil, zero value otherwise.

### GetIamUserArnOk

`func (o *AwsIntegrationCredential) GetIamUserArnOk() (*string, bool)`

GetIamUserArnOk returns a tuple with the IamUserArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamUserArn

`func (o *AwsIntegrationCredential) SetIamUserArn(v string)`

SetIamUserArn sets IamUserArn field to given value.

### HasIamUserArn

`func (o *AwsIntegrationCredential) HasIamUserArn() bool`

HasIamUserArn returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *AwsIntegrationCredential) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AwsIntegrationCredential) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AwsIntegrationCredential) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AwsIntegrationCredential) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


