# AzureIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmaFullSyncDone** | Pointer to **bool** | Is Azure Commercial Marketplace Analytics (CMA) full-sync done. | [optional] 
**Credential** | Pointer to [**AzureIntegrationCredential**](AzureIntegrationCredential.md) |  | [optional] 
**PartnerCenterReferralSyncPaused** | Pointer to **bool** | Is Microsoft Partner Center referral sync paused. | [optional] 
**RevenueRecordFullSyncDone** | Pointer to **bool** | Is AZURE Marketplace Revenue Record full-sync done. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to store the AzureIntegrationCredential in AWS Secret manager. for internal usage only. | [optional] 

## Methods

### NewAzureIntegration

`func NewAzureIntegration() *AzureIntegration`

NewAzureIntegration instantiates a new AzureIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureIntegrationWithDefaults

`func NewAzureIntegrationWithDefaults() *AzureIntegration`

NewAzureIntegrationWithDefaults instantiates a new AzureIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmaFullSyncDone

`func (o *AzureIntegration) GetCmaFullSyncDone() bool`

GetCmaFullSyncDone returns the CmaFullSyncDone field if non-nil, zero value otherwise.

### GetCmaFullSyncDoneOk

`func (o *AzureIntegration) GetCmaFullSyncDoneOk() (*bool, bool)`

GetCmaFullSyncDoneOk returns a tuple with the CmaFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmaFullSyncDone

`func (o *AzureIntegration) SetCmaFullSyncDone(v bool)`

SetCmaFullSyncDone sets CmaFullSyncDone field to given value.

### HasCmaFullSyncDone

`func (o *AzureIntegration) HasCmaFullSyncDone() bool`

HasCmaFullSyncDone returns a boolean if a field has been set.

### GetCredential

`func (o *AzureIntegration) GetCredential() AzureIntegrationCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AzureIntegration) GetCredentialOk() (*AzureIntegrationCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AzureIntegration) SetCredential(v AzureIntegrationCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AzureIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetPartnerCenterReferralSyncPaused

`func (o *AzureIntegration) GetPartnerCenterReferralSyncPaused() bool`

GetPartnerCenterReferralSyncPaused returns the PartnerCenterReferralSyncPaused field if non-nil, zero value otherwise.

### GetPartnerCenterReferralSyncPausedOk

`func (o *AzureIntegration) GetPartnerCenterReferralSyncPausedOk() (*bool, bool)`

GetPartnerCenterReferralSyncPausedOk returns a tuple with the PartnerCenterReferralSyncPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCenterReferralSyncPaused

`func (o *AzureIntegration) SetPartnerCenterReferralSyncPaused(v bool)`

SetPartnerCenterReferralSyncPaused sets PartnerCenterReferralSyncPaused field to given value.

### HasPartnerCenterReferralSyncPaused

`func (o *AzureIntegration) HasPartnerCenterReferralSyncPaused() bool`

HasPartnerCenterReferralSyncPaused returns a boolean if a field has been set.

### GetRevenueRecordFullSyncDone

`func (o *AzureIntegration) GetRevenueRecordFullSyncDone() bool`

GetRevenueRecordFullSyncDone returns the RevenueRecordFullSyncDone field if non-nil, zero value otherwise.

### GetRevenueRecordFullSyncDoneOk

`func (o *AzureIntegration) GetRevenueRecordFullSyncDoneOk() (*bool, bool)`

GetRevenueRecordFullSyncDoneOk returns a tuple with the RevenueRecordFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecordFullSyncDone

`func (o *AzureIntegration) SetRevenueRecordFullSyncDone(v bool)`

SetRevenueRecordFullSyncDone sets RevenueRecordFullSyncDone field to given value.

### HasRevenueRecordFullSyncDone

`func (o *AzureIntegration) HasRevenueRecordFullSyncDone() bool`

HasRevenueRecordFullSyncDone returns a boolean if a field has been set.

### GetSecretKey

`func (o *AzureIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AzureIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AzureIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AzureIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


