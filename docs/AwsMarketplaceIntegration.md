# AwsMarketplaceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicInfoFullSyncDone** | Pointer to **bool** | Is AWS Marketplace Basic Info (Private Offers, Entitlements/Agreements) full-sync done. | [optional] 
**CreateEntitlementBeforeNotificationEnabled** | Pointer to **bool** | If enabled, Suger will create an active entitlement based on the default offer when the new client is redirected from AWS Marketplace, but the notification evnet from AWS Marketplace is not received yet by Suger. If disabled, Suger will not create the entitlement when the new client is redirected from AWS Marketplace. So the redirect URL won&#39;t contain the sugerEntitlementId parameter. | [optional] 
**EnableMarketplaceBetaApi** | Pointer to **bool** | Enable the use of AWS Marketplace beta APIs in the backend. If true, Suger will use the beta APIs. | [optional] 
**EventBridgeRuleName** | Pointer to **string** | AWS EventBridge rule name for AWS Marketplace events. | [optional] 
**ExternalID** | Pointer to **string** | The external ID for assuming IAM role. If empty, means no external ID set or needed. Otherwise, it should be auth_id in table identity.organization. | [optional] 
**IamRoleArn** | Pointer to **string** | The AWS IAM role for Suger service to assume to access the client&#39;s AWS services. | [optional] 
**MarketplaceStartDate** | Pointer to **time.Time** | AWS Marketplace start date which comes from MDFS Full-Sync. | [optional] 
**McasFullSyncDone** | Pointer to **bool** | Is AWS Marketplace Commerce Analytics Service (MCAS) full-sync done. | [optional] 
**McasIamRoleArn** | Pointer to **string** | IAM role ARN to allow AWS Marketplace to write to the S3 bucket and publish notifications to the SNS topic. | [optional] 
**McasS3Bucket** | Pointer to **string** | S3 bucket for AWS Marketplace Commerce Analytics Service (MCAS) | [optional] 
**McasSnsTopic** | Pointer to **string** | SNS topic ARN for AWS Marketplace Commerce Analytics Service (MCAS) | [optional] 
**McasSyncDisabled** | Pointer to **bool** | Disable Sync with AWS MCAS. If true, Suger stop to sync with MCAS. | [optional] 
**MdfsFullSyncDone** | Pointer to **bool** | Is AWS Marketplace Data Feeds Service (MDFS) full-sync done. | [optional] 
**MdfsKmsKeyArn** | Pointer to **string** | KMS Key ARN for the S3 bucket of AWS Marketplace Data Feeds Service (MDFS) | [optional] 
**MdfsS3BucketArn** | Pointer to **string** | S3 bucket ARN for AWS Marketplace Data Feeds Service (MDFS) | [optional] 
**MdfsSyncDisabled** | Pointer to **bool** | Disable Sync with AWS MDFS. If true, Suger stop to sync with MDFS. | [optional] 
**PolicyArns** | Pointer to **[]string** | The policy ARNs in the IAM role. | [optional] 
**RevenueRecordFullSyncDone** | Pointer to **bool** | Is AWS Marketplace Revenue Record full-sync done. | [optional] 
**SignupRedirectWithoutEntitlementEnabled** | Pointer to **bool** | If enabled, Suger will redirect the new client to the signup page even the entitlement is not found. If disabled, Suger will redirect the new client to the error page if the entitlement is not found. | [optional] 
**UsageMeteringDisabled** | Pointer to **bool** | Disable Usage Metering to AWS Marketplace. If true, Suger stop to report usage records to AWS Marketplace. | [optional] 

## Methods

### NewAwsMarketplaceIntegration

`func NewAwsMarketplaceIntegration() *AwsMarketplaceIntegration`

NewAwsMarketplaceIntegration instantiates a new AwsMarketplaceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceIntegrationWithDefaults

`func NewAwsMarketplaceIntegrationWithDefaults() *AwsMarketplaceIntegration`

NewAwsMarketplaceIntegrationWithDefaults instantiates a new AwsMarketplaceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicInfoFullSyncDone

`func (o *AwsMarketplaceIntegration) GetBasicInfoFullSyncDone() bool`

GetBasicInfoFullSyncDone returns the BasicInfoFullSyncDone field if non-nil, zero value otherwise.

### GetBasicInfoFullSyncDoneOk

`func (o *AwsMarketplaceIntegration) GetBasicInfoFullSyncDoneOk() (*bool, bool)`

GetBasicInfoFullSyncDoneOk returns a tuple with the BasicInfoFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicInfoFullSyncDone

`func (o *AwsMarketplaceIntegration) SetBasicInfoFullSyncDone(v bool)`

SetBasicInfoFullSyncDone sets BasicInfoFullSyncDone field to given value.

### HasBasicInfoFullSyncDone

`func (o *AwsMarketplaceIntegration) HasBasicInfoFullSyncDone() bool`

HasBasicInfoFullSyncDone returns a boolean if a field has been set.

### GetCreateEntitlementBeforeNotificationEnabled

`func (o *AwsMarketplaceIntegration) GetCreateEntitlementBeforeNotificationEnabled() bool`

GetCreateEntitlementBeforeNotificationEnabled returns the CreateEntitlementBeforeNotificationEnabled field if non-nil, zero value otherwise.

### GetCreateEntitlementBeforeNotificationEnabledOk

`func (o *AwsMarketplaceIntegration) GetCreateEntitlementBeforeNotificationEnabledOk() (*bool, bool)`

GetCreateEntitlementBeforeNotificationEnabledOk returns a tuple with the CreateEntitlementBeforeNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEntitlementBeforeNotificationEnabled

`func (o *AwsMarketplaceIntegration) SetCreateEntitlementBeforeNotificationEnabled(v bool)`

SetCreateEntitlementBeforeNotificationEnabled sets CreateEntitlementBeforeNotificationEnabled field to given value.

### HasCreateEntitlementBeforeNotificationEnabled

`func (o *AwsMarketplaceIntegration) HasCreateEntitlementBeforeNotificationEnabled() bool`

HasCreateEntitlementBeforeNotificationEnabled returns a boolean if a field has been set.

### GetEnableMarketplaceBetaApi

`func (o *AwsMarketplaceIntegration) GetEnableMarketplaceBetaApi() bool`

GetEnableMarketplaceBetaApi returns the EnableMarketplaceBetaApi field if non-nil, zero value otherwise.

### GetEnableMarketplaceBetaApiOk

`func (o *AwsMarketplaceIntegration) GetEnableMarketplaceBetaApiOk() (*bool, bool)`

GetEnableMarketplaceBetaApiOk returns a tuple with the EnableMarketplaceBetaApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMarketplaceBetaApi

`func (o *AwsMarketplaceIntegration) SetEnableMarketplaceBetaApi(v bool)`

SetEnableMarketplaceBetaApi sets EnableMarketplaceBetaApi field to given value.

### HasEnableMarketplaceBetaApi

`func (o *AwsMarketplaceIntegration) HasEnableMarketplaceBetaApi() bool`

HasEnableMarketplaceBetaApi returns a boolean if a field has been set.

### GetEventBridgeRuleName

`func (o *AwsMarketplaceIntegration) GetEventBridgeRuleName() string`

GetEventBridgeRuleName returns the EventBridgeRuleName field if non-nil, zero value otherwise.

### GetEventBridgeRuleNameOk

`func (o *AwsMarketplaceIntegration) GetEventBridgeRuleNameOk() (*string, bool)`

GetEventBridgeRuleNameOk returns a tuple with the EventBridgeRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBridgeRuleName

`func (o *AwsMarketplaceIntegration) SetEventBridgeRuleName(v string)`

SetEventBridgeRuleName sets EventBridgeRuleName field to given value.

### HasEventBridgeRuleName

`func (o *AwsMarketplaceIntegration) HasEventBridgeRuleName() bool`

HasEventBridgeRuleName returns a boolean if a field has been set.

### GetExternalID

`func (o *AwsMarketplaceIntegration) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *AwsMarketplaceIntegration) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *AwsMarketplaceIntegration) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *AwsMarketplaceIntegration) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetIamRoleArn

`func (o *AwsMarketplaceIntegration) GetIamRoleArn() string`

GetIamRoleArn returns the IamRoleArn field if non-nil, zero value otherwise.

### GetIamRoleArnOk

`func (o *AwsMarketplaceIntegration) GetIamRoleArnOk() (*string, bool)`

GetIamRoleArnOk returns a tuple with the IamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleArn

`func (o *AwsMarketplaceIntegration) SetIamRoleArn(v string)`

SetIamRoleArn sets IamRoleArn field to given value.

### HasIamRoleArn

`func (o *AwsMarketplaceIntegration) HasIamRoleArn() bool`

HasIamRoleArn returns a boolean if a field has been set.

### GetMarketplaceStartDate

`func (o *AwsMarketplaceIntegration) GetMarketplaceStartDate() time.Time`

GetMarketplaceStartDate returns the MarketplaceStartDate field if non-nil, zero value otherwise.

### GetMarketplaceStartDateOk

`func (o *AwsMarketplaceIntegration) GetMarketplaceStartDateOk() (*time.Time, bool)`

GetMarketplaceStartDateOk returns a tuple with the MarketplaceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceStartDate

`func (o *AwsMarketplaceIntegration) SetMarketplaceStartDate(v time.Time)`

SetMarketplaceStartDate sets MarketplaceStartDate field to given value.

### HasMarketplaceStartDate

`func (o *AwsMarketplaceIntegration) HasMarketplaceStartDate() bool`

HasMarketplaceStartDate returns a boolean if a field has been set.

### GetMcasFullSyncDone

`func (o *AwsMarketplaceIntegration) GetMcasFullSyncDone() bool`

GetMcasFullSyncDone returns the McasFullSyncDone field if non-nil, zero value otherwise.

### GetMcasFullSyncDoneOk

`func (o *AwsMarketplaceIntegration) GetMcasFullSyncDoneOk() (*bool, bool)`

GetMcasFullSyncDoneOk returns a tuple with the McasFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcasFullSyncDone

`func (o *AwsMarketplaceIntegration) SetMcasFullSyncDone(v bool)`

SetMcasFullSyncDone sets McasFullSyncDone field to given value.

### HasMcasFullSyncDone

`func (o *AwsMarketplaceIntegration) HasMcasFullSyncDone() bool`

HasMcasFullSyncDone returns a boolean if a field has been set.

### GetMcasIamRoleArn

`func (o *AwsMarketplaceIntegration) GetMcasIamRoleArn() string`

GetMcasIamRoleArn returns the McasIamRoleArn field if non-nil, zero value otherwise.

### GetMcasIamRoleArnOk

`func (o *AwsMarketplaceIntegration) GetMcasIamRoleArnOk() (*string, bool)`

GetMcasIamRoleArnOk returns a tuple with the McasIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcasIamRoleArn

`func (o *AwsMarketplaceIntegration) SetMcasIamRoleArn(v string)`

SetMcasIamRoleArn sets McasIamRoleArn field to given value.

### HasMcasIamRoleArn

`func (o *AwsMarketplaceIntegration) HasMcasIamRoleArn() bool`

HasMcasIamRoleArn returns a boolean if a field has been set.

### GetMcasS3Bucket

`func (o *AwsMarketplaceIntegration) GetMcasS3Bucket() string`

GetMcasS3Bucket returns the McasS3Bucket field if non-nil, zero value otherwise.

### GetMcasS3BucketOk

`func (o *AwsMarketplaceIntegration) GetMcasS3BucketOk() (*string, bool)`

GetMcasS3BucketOk returns a tuple with the McasS3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcasS3Bucket

`func (o *AwsMarketplaceIntegration) SetMcasS3Bucket(v string)`

SetMcasS3Bucket sets McasS3Bucket field to given value.

### HasMcasS3Bucket

`func (o *AwsMarketplaceIntegration) HasMcasS3Bucket() bool`

HasMcasS3Bucket returns a boolean if a field has been set.

### GetMcasSnsTopic

`func (o *AwsMarketplaceIntegration) GetMcasSnsTopic() string`

GetMcasSnsTopic returns the McasSnsTopic field if non-nil, zero value otherwise.

### GetMcasSnsTopicOk

`func (o *AwsMarketplaceIntegration) GetMcasSnsTopicOk() (*string, bool)`

GetMcasSnsTopicOk returns a tuple with the McasSnsTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcasSnsTopic

`func (o *AwsMarketplaceIntegration) SetMcasSnsTopic(v string)`

SetMcasSnsTopic sets McasSnsTopic field to given value.

### HasMcasSnsTopic

`func (o *AwsMarketplaceIntegration) HasMcasSnsTopic() bool`

HasMcasSnsTopic returns a boolean if a field has been set.

### GetMcasSyncDisabled

`func (o *AwsMarketplaceIntegration) GetMcasSyncDisabled() bool`

GetMcasSyncDisabled returns the McasSyncDisabled field if non-nil, zero value otherwise.

### GetMcasSyncDisabledOk

`func (o *AwsMarketplaceIntegration) GetMcasSyncDisabledOk() (*bool, bool)`

GetMcasSyncDisabledOk returns a tuple with the McasSyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcasSyncDisabled

`func (o *AwsMarketplaceIntegration) SetMcasSyncDisabled(v bool)`

SetMcasSyncDisabled sets McasSyncDisabled field to given value.

### HasMcasSyncDisabled

`func (o *AwsMarketplaceIntegration) HasMcasSyncDisabled() bool`

HasMcasSyncDisabled returns a boolean if a field has been set.

### GetMdfsFullSyncDone

`func (o *AwsMarketplaceIntegration) GetMdfsFullSyncDone() bool`

GetMdfsFullSyncDone returns the MdfsFullSyncDone field if non-nil, zero value otherwise.

### GetMdfsFullSyncDoneOk

`func (o *AwsMarketplaceIntegration) GetMdfsFullSyncDoneOk() (*bool, bool)`

GetMdfsFullSyncDoneOk returns a tuple with the MdfsFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfsFullSyncDone

`func (o *AwsMarketplaceIntegration) SetMdfsFullSyncDone(v bool)`

SetMdfsFullSyncDone sets MdfsFullSyncDone field to given value.

### HasMdfsFullSyncDone

`func (o *AwsMarketplaceIntegration) HasMdfsFullSyncDone() bool`

HasMdfsFullSyncDone returns a boolean if a field has been set.

### GetMdfsKmsKeyArn

`func (o *AwsMarketplaceIntegration) GetMdfsKmsKeyArn() string`

GetMdfsKmsKeyArn returns the MdfsKmsKeyArn field if non-nil, zero value otherwise.

### GetMdfsKmsKeyArnOk

`func (o *AwsMarketplaceIntegration) GetMdfsKmsKeyArnOk() (*string, bool)`

GetMdfsKmsKeyArnOk returns a tuple with the MdfsKmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfsKmsKeyArn

`func (o *AwsMarketplaceIntegration) SetMdfsKmsKeyArn(v string)`

SetMdfsKmsKeyArn sets MdfsKmsKeyArn field to given value.

### HasMdfsKmsKeyArn

`func (o *AwsMarketplaceIntegration) HasMdfsKmsKeyArn() bool`

HasMdfsKmsKeyArn returns a boolean if a field has been set.

### GetMdfsS3BucketArn

`func (o *AwsMarketplaceIntegration) GetMdfsS3BucketArn() string`

GetMdfsS3BucketArn returns the MdfsS3BucketArn field if non-nil, zero value otherwise.

### GetMdfsS3BucketArnOk

`func (o *AwsMarketplaceIntegration) GetMdfsS3BucketArnOk() (*string, bool)`

GetMdfsS3BucketArnOk returns a tuple with the MdfsS3BucketArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfsS3BucketArn

`func (o *AwsMarketplaceIntegration) SetMdfsS3BucketArn(v string)`

SetMdfsS3BucketArn sets MdfsS3BucketArn field to given value.

### HasMdfsS3BucketArn

`func (o *AwsMarketplaceIntegration) HasMdfsS3BucketArn() bool`

HasMdfsS3BucketArn returns a boolean if a field has been set.

### GetMdfsSyncDisabled

`func (o *AwsMarketplaceIntegration) GetMdfsSyncDisabled() bool`

GetMdfsSyncDisabled returns the MdfsSyncDisabled field if non-nil, zero value otherwise.

### GetMdfsSyncDisabledOk

`func (o *AwsMarketplaceIntegration) GetMdfsSyncDisabledOk() (*bool, bool)`

GetMdfsSyncDisabledOk returns a tuple with the MdfsSyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdfsSyncDisabled

`func (o *AwsMarketplaceIntegration) SetMdfsSyncDisabled(v bool)`

SetMdfsSyncDisabled sets MdfsSyncDisabled field to given value.

### HasMdfsSyncDisabled

`func (o *AwsMarketplaceIntegration) HasMdfsSyncDisabled() bool`

HasMdfsSyncDisabled returns a boolean if a field has been set.

### GetPolicyArns

`func (o *AwsMarketplaceIntegration) GetPolicyArns() []string`

GetPolicyArns returns the PolicyArns field if non-nil, zero value otherwise.

### GetPolicyArnsOk

`func (o *AwsMarketplaceIntegration) GetPolicyArnsOk() (*[]string, bool)`

GetPolicyArnsOk returns a tuple with the PolicyArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyArns

`func (o *AwsMarketplaceIntegration) SetPolicyArns(v []string)`

SetPolicyArns sets PolicyArns field to given value.

### HasPolicyArns

`func (o *AwsMarketplaceIntegration) HasPolicyArns() bool`

HasPolicyArns returns a boolean if a field has been set.

### GetRevenueRecordFullSyncDone

`func (o *AwsMarketplaceIntegration) GetRevenueRecordFullSyncDone() bool`

GetRevenueRecordFullSyncDone returns the RevenueRecordFullSyncDone field if non-nil, zero value otherwise.

### GetRevenueRecordFullSyncDoneOk

`func (o *AwsMarketplaceIntegration) GetRevenueRecordFullSyncDoneOk() (*bool, bool)`

GetRevenueRecordFullSyncDoneOk returns a tuple with the RevenueRecordFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecordFullSyncDone

`func (o *AwsMarketplaceIntegration) SetRevenueRecordFullSyncDone(v bool)`

SetRevenueRecordFullSyncDone sets RevenueRecordFullSyncDone field to given value.

### HasRevenueRecordFullSyncDone

`func (o *AwsMarketplaceIntegration) HasRevenueRecordFullSyncDone() bool`

HasRevenueRecordFullSyncDone returns a boolean if a field has been set.

### GetSignupRedirectWithoutEntitlementEnabled

`func (o *AwsMarketplaceIntegration) GetSignupRedirectWithoutEntitlementEnabled() bool`

GetSignupRedirectWithoutEntitlementEnabled returns the SignupRedirectWithoutEntitlementEnabled field if non-nil, zero value otherwise.

### GetSignupRedirectWithoutEntitlementEnabledOk

`func (o *AwsMarketplaceIntegration) GetSignupRedirectWithoutEntitlementEnabledOk() (*bool, bool)`

GetSignupRedirectWithoutEntitlementEnabledOk returns a tuple with the SignupRedirectWithoutEntitlementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupRedirectWithoutEntitlementEnabled

`func (o *AwsMarketplaceIntegration) SetSignupRedirectWithoutEntitlementEnabled(v bool)`

SetSignupRedirectWithoutEntitlementEnabled sets SignupRedirectWithoutEntitlementEnabled field to given value.

### HasSignupRedirectWithoutEntitlementEnabled

`func (o *AwsMarketplaceIntegration) HasSignupRedirectWithoutEntitlementEnabled() bool`

HasSignupRedirectWithoutEntitlementEnabled returns a boolean if a field has been set.

### GetUsageMeteringDisabled

`func (o *AwsMarketplaceIntegration) GetUsageMeteringDisabled() bool`

GetUsageMeteringDisabled returns the UsageMeteringDisabled field if non-nil, zero value otherwise.

### GetUsageMeteringDisabledOk

`func (o *AwsMarketplaceIntegration) GetUsageMeteringDisabledOk() (*bool, bool)`

GetUsageMeteringDisabledOk returns a tuple with the UsageMeteringDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeteringDisabled

`func (o *AwsMarketplaceIntegration) SetUsageMeteringDisabled(v bool)`

SetUsageMeteringDisabled sets UsageMeteringDisabled field to given value.

### HasUsageMeteringDisabled

`func (o *AwsMarketplaceIntegration) HasUsageMeteringDisabled() bool`

HasUsageMeteringDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


