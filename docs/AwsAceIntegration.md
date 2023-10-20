# AwsAceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**AwsIntegrationCredential**](AwsIntegrationCredential.md) |  | [optional] 
**EnableAssumeIamRole** | Pointer to **bool** | Enable assume IAM role to access the client&#39;s AWS ACE S3 bucket. If false, Suger will use the AwsIntegrationCredential to access the client&#39;s AWS ACE S3 bucket. | [optional] 
**IamRoleArn** | Pointer to **string** | The AWS IAM role for Suger service to assume to access the client&#39;s AWS ACE S3 bucket. Applicable only if EnableAssumeIamRole is true. | [optional] 
**PartnerId** | Pointer to **string** | The partner ID of the ISV/Seller in AWS Partner Network. | [optional] 
**Paused** | Pointer to **bool** | Is the integration paused. | [optional] 
**S3BucketName** | Pointer to **string** | The Name of the S3 bucket for AWS APN Customer Engagement Program (ACE) to sync the leads &amp; opportunities. | [optional] 
**S3BucketRegion** | Pointer to **string** | The region of the S3 bucket for AWS APN Customer Engagement Program (ACE) to sync the leads &amp; opportunities. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to store the AwsIntegrationCredential in AWS Secret manager. for internal usage only. | [optional] 

## Methods

### NewAwsAceIntegration

`func NewAwsAceIntegration() *AwsAceIntegration`

NewAwsAceIntegration instantiates a new AwsAceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAceIntegrationWithDefaults

`func NewAwsAceIntegrationWithDefaults() *AwsAceIntegration`

NewAwsAceIntegrationWithDefaults instantiates a new AwsAceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AwsAceIntegration) GetCredential() AwsIntegrationCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AwsAceIntegration) GetCredentialOk() (*AwsIntegrationCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AwsAceIntegration) SetCredential(v AwsIntegrationCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AwsAceIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetEnableAssumeIamRole

`func (o *AwsAceIntegration) GetEnableAssumeIamRole() bool`

GetEnableAssumeIamRole returns the EnableAssumeIamRole field if non-nil, zero value otherwise.

### GetEnableAssumeIamRoleOk

`func (o *AwsAceIntegration) GetEnableAssumeIamRoleOk() (*bool, bool)`

GetEnableAssumeIamRoleOk returns a tuple with the EnableAssumeIamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAssumeIamRole

`func (o *AwsAceIntegration) SetEnableAssumeIamRole(v bool)`

SetEnableAssumeIamRole sets EnableAssumeIamRole field to given value.

### HasEnableAssumeIamRole

`func (o *AwsAceIntegration) HasEnableAssumeIamRole() bool`

HasEnableAssumeIamRole returns a boolean if a field has been set.

### GetIamRoleArn

`func (o *AwsAceIntegration) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsAceIntegration) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsAceIntegration) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsAceIntegration) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### GetPartnerId

`func (o *AwsAceIntegration) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *AwsAceIntegration) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *AwsAceIntegration) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *AwsAceIntegration) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPaused

`func (o *AwsAceIntegration) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *AwsAceIntegration) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *AwsAceIntegration) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *AwsAceIntegration) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetS3BucketName

`func (o *AwsAceIntegration) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *AwsAceIntegration) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *AwsAceIntegration) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.

### HasS3BucketName

`func (o *AwsAceIntegration) HasS3BucketName() bool`

HasS3BucketName returns a boolean if a field has been set.

### GetS3BucketRegion

`func (o *AwsAceIntegration) GetS3BucketRegion() string`

GetS3BucketRegion returns the S3BucketRegion field if non-nil, zero value otherwise.

### GetS3BucketRegionOk

`func (o *AwsAceIntegration) GetS3BucketRegionOk() (*string, bool)`

GetS3BucketRegionOk returns a tuple with the S3BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketRegion

`func (o *AwsAceIntegration) SetS3BucketRegion(v string)`

SetS3BucketRegion sets S3BucketRegion field to given value.

### HasS3BucketRegion

`func (o *AwsAceIntegration) HasS3BucketRegion() bool`

HasS3BucketRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *AwsAceIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AwsAceIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AwsAceIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AwsAceIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


