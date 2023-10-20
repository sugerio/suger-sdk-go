# GcpIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableChromeSync** | Pointer to **bool** | Enable GCP marketplace sync from GCP Chrome. If true, Suger will sync GCP Marketplace Product &amp; Private Offer from GCP Chrome. | [optional] 
**EnableManualApproveEntitlement** | Pointer to **bool** | Enable manually approve the GCP Marketplace Entitlement. If true, Suger will not automatically approve the GCP Marketplace Entitlement. Util the GCP Marketplace Entitlement is manually approved, it will not be activated. | [optional] 
**GcpProjectId** | Pointer to **string** |  | [optional] 
**GcpProjectNumber** | Pointer to **string** |  | [optional] 
**IdentityProviderId** | Pointer to **string** |  | [optional] 
**LoginUrl** | Pointer to **string** | The Login URL for GCP Marketplace buyers to login to your service with or without SSO JWT token. If not set, the login will be redirected to the Product&#39;s fulfillment URL. | [optional] 
**PartnerId** | Pointer to **string** | The GCP Marketplace Partner ID, it is also called as Provider ID somewhere. | [optional] 
**PubsubSubscription** | Pointer to **string** | The resource name of the Pub/Sub subscription to receive notifications from Google cloud marketplace. | [optional] 
**PubsubTopic** | Pointer to **string** | The resource name of the Pub/Sub topic to receive notifications from Google when a user signs up for your service, purchases a plan, or changes an existing plan. | [optional] 
**ReportBucket** | Pointer to **string** | The GCP storage bucket name to store the GCP Marketplace reports. | [optional] 
**ReportFullSyncDone** | Pointer to **bool** | Is GCP Marketplace Report full-sync done. | [optional] 
**ReportStartDate** | Pointer to **time.Time** | The UTC date when GCP Marketplace reports start to generate. | [optional] 
**ReportSyncDisabled** | Pointer to **bool** | Disable GCP Marketplace Report sync. If true, Suger stop to sync GCP Marketplace reports. | [optional] 
**RevenueRecordFullSyncDone** | Pointer to **bool** | Is GCP Marketplace Revenue Record full-sync done. | [optional] 
**RevenueRecordSyncDisabled** | Pointer to **bool** | Disable GCP Marketplace Revenue Record sync. If true, Suger stop to sync GCP Marketplace Revenue Records. | [optional] 
**ServiceAccountEmail** | Pointer to **string** |  | [optional] 
**ServiceNames** | Pointer to **[]string** | The array of service resource names of the listings in GCP Marketplace. | [optional] 
**SignupRedirectWithoutEntitlementEnabled** | Pointer to **bool** | If enabled, Suger will redirect the new client to the signup page even the entitlement is not found. If disabled, Suger will redirect the new client to the error page if the entitlement is not found. | [optional] 
**UsageMeteringDisabled** | Pointer to **bool** | Disable Usage Metering to GCP Marketplace. If true, Suger stop to report usage records to GCP Marketplace. | [optional] 
**WorkloadIdentityPoolId** | Pointer to **string** |  | [optional] 

## Methods

### NewGcpIntegration

`func NewGcpIntegration() *GcpIntegration`

NewGcpIntegration instantiates a new GcpIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpIntegrationWithDefaults

`func NewGcpIntegrationWithDefaults() *GcpIntegration`

NewGcpIntegrationWithDefaults instantiates a new GcpIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableChromeSync

`func (o *GcpIntegration) GetEnableChromeSync() bool`

GetEnableChromeSync returns the EnableChromeSync field if non-nil, zero value otherwise.

### GetEnableChromeSyncOk

`func (o *GcpIntegration) GetEnableChromeSyncOk() (*bool, bool)`

GetEnableChromeSyncOk returns a tuple with the EnableChromeSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChromeSync

`func (o *GcpIntegration) SetEnableChromeSync(v bool)`

SetEnableChromeSync sets EnableChromeSync field to given value.

### HasEnableChromeSync

`func (o *GcpIntegration) HasEnableChromeSync() bool`

HasEnableChromeSync returns a boolean if a field has been set.

### GetEnableManualApproveEntitlement

`func (o *GcpIntegration) GetEnableManualApproveEntitlement() bool`

GetEnableManualApproveEntitlement returns the EnableManualApproveEntitlement field if non-nil, zero value otherwise.

### GetEnableManualApproveEntitlementOk

`func (o *GcpIntegration) GetEnableManualApproveEntitlementOk() (*bool, bool)`

GetEnableManualApproveEntitlementOk returns a tuple with the EnableManualApproveEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManualApproveEntitlement

`func (o *GcpIntegration) SetEnableManualApproveEntitlement(v bool)`

SetEnableManualApproveEntitlement sets EnableManualApproveEntitlement field to given value.

### HasEnableManualApproveEntitlement

`func (o *GcpIntegration) HasEnableManualApproveEntitlement() bool`

HasEnableManualApproveEntitlement returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *GcpIntegration) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *GcpIntegration) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *GcpIntegration) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *GcpIntegration) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *GcpIntegration) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *GcpIntegration) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *GcpIntegration) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *GcpIntegration) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetIdentityProviderId

`func (o *GcpIntegration) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *GcpIntegration) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *GcpIntegration) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *GcpIntegration) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### GetLoginUrl

`func (o *GcpIntegration) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *GcpIntegration) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *GcpIntegration) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *GcpIntegration) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetPartnerId

`func (o *GcpIntegration) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *GcpIntegration) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *GcpIntegration) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *GcpIntegration) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPubsubSubscription

`func (o *GcpIntegration) GetPubsubSubscription() string`

GetPubsubSubscription returns the PubsubSubscription field if non-nil, zero value otherwise.

### GetPubsubSubscriptionOk

`func (o *GcpIntegration) GetPubsubSubscriptionOk() (*string, bool)`

GetPubsubSubscriptionOk returns a tuple with the PubsubSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubSubscription

`func (o *GcpIntegration) SetPubsubSubscription(v string)`

SetPubsubSubscription sets PubsubSubscription field to given value.

### HasPubsubSubscription

`func (o *GcpIntegration) HasPubsubSubscription() bool`

HasPubsubSubscription returns a boolean if a field has been set.

### GetPubsubTopic

`func (o *GcpIntegration) GetPubsubTopic() string`

GetPubsubTopic returns the PubsubTopic field if non-nil, zero value otherwise.

### GetPubsubTopicOk

`func (o *GcpIntegration) GetPubsubTopicOk() (*string, bool)`

GetPubsubTopicOk returns a tuple with the PubsubTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubTopic

`func (o *GcpIntegration) SetPubsubTopic(v string)`

SetPubsubTopic sets PubsubTopic field to given value.

### HasPubsubTopic

`func (o *GcpIntegration) HasPubsubTopic() bool`

HasPubsubTopic returns a boolean if a field has been set.

### GetReportBucket

`func (o *GcpIntegration) GetReportBucket() string`

GetReportBucket returns the ReportBucket field if non-nil, zero value otherwise.

### GetReportBucketOk

`func (o *GcpIntegration) GetReportBucketOk() (*string, bool)`

GetReportBucketOk returns a tuple with the ReportBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportBucket

`func (o *GcpIntegration) SetReportBucket(v string)`

SetReportBucket sets ReportBucket field to given value.

### HasReportBucket

`func (o *GcpIntegration) HasReportBucket() bool`

HasReportBucket returns a boolean if a field has been set.

### GetReportFullSyncDone

`func (o *GcpIntegration) GetReportFullSyncDone() bool`

GetReportFullSyncDone returns the ReportFullSyncDone field if non-nil, zero value otherwise.

### GetReportFullSyncDoneOk

`func (o *GcpIntegration) GetReportFullSyncDoneOk() (*bool, bool)`

GetReportFullSyncDoneOk returns a tuple with the ReportFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFullSyncDone

`func (o *GcpIntegration) SetReportFullSyncDone(v bool)`

SetReportFullSyncDone sets ReportFullSyncDone field to given value.

### HasReportFullSyncDone

`func (o *GcpIntegration) HasReportFullSyncDone() bool`

HasReportFullSyncDone returns a boolean if a field has been set.

### GetReportStartDate

`func (o *GcpIntegration) GetReportStartDate() time.Time`

GetReportStartDate returns the ReportStartDate field if non-nil, zero value otherwise.

### GetReportStartDateOk

`func (o *GcpIntegration) GetReportStartDateOk() (*time.Time, bool)`

GetReportStartDateOk returns a tuple with the ReportStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStartDate

`func (o *GcpIntegration) SetReportStartDate(v time.Time)`

SetReportStartDate sets ReportStartDate field to given value.

### HasReportStartDate

`func (o *GcpIntegration) HasReportStartDate() bool`

HasReportStartDate returns a boolean if a field has been set.

### GetReportSyncDisabled

`func (o *GcpIntegration) GetReportSyncDisabled() bool`

GetReportSyncDisabled returns the ReportSyncDisabled field if non-nil, zero value otherwise.

### GetReportSyncDisabledOk

`func (o *GcpIntegration) GetReportSyncDisabledOk() (*bool, bool)`

GetReportSyncDisabledOk returns a tuple with the ReportSyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportSyncDisabled

`func (o *GcpIntegration) SetReportSyncDisabled(v bool)`

SetReportSyncDisabled sets ReportSyncDisabled field to given value.

### HasReportSyncDisabled

`func (o *GcpIntegration) HasReportSyncDisabled() bool`

HasReportSyncDisabled returns a boolean if a field has been set.

### GetRevenueRecordFullSyncDone

`func (o *GcpIntegration) GetRevenueRecordFullSyncDone() bool`

GetRevenueRecordFullSyncDone returns the RevenueRecordFullSyncDone field if non-nil, zero value otherwise.

### GetRevenueRecordFullSyncDoneOk

`func (o *GcpIntegration) GetRevenueRecordFullSyncDoneOk() (*bool, bool)`

GetRevenueRecordFullSyncDoneOk returns a tuple with the RevenueRecordFullSyncDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecordFullSyncDone

`func (o *GcpIntegration) SetRevenueRecordFullSyncDone(v bool)`

SetRevenueRecordFullSyncDone sets RevenueRecordFullSyncDone field to given value.

### HasRevenueRecordFullSyncDone

`func (o *GcpIntegration) HasRevenueRecordFullSyncDone() bool`

HasRevenueRecordFullSyncDone returns a boolean if a field has been set.

### GetRevenueRecordSyncDisabled

`func (o *GcpIntegration) GetRevenueRecordSyncDisabled() bool`

GetRevenueRecordSyncDisabled returns the RevenueRecordSyncDisabled field if non-nil, zero value otherwise.

### GetRevenueRecordSyncDisabledOk

`func (o *GcpIntegration) GetRevenueRecordSyncDisabledOk() (*bool, bool)`

GetRevenueRecordSyncDisabledOk returns a tuple with the RevenueRecordSyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecordSyncDisabled

`func (o *GcpIntegration) SetRevenueRecordSyncDisabled(v bool)`

SetRevenueRecordSyncDisabled sets RevenueRecordSyncDisabled field to given value.

### HasRevenueRecordSyncDisabled

`func (o *GcpIntegration) HasRevenueRecordSyncDisabled() bool`

HasRevenueRecordSyncDisabled returns a boolean if a field has been set.

### GetServiceAccountEmail

`func (o *GcpIntegration) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GcpIntegration) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GcpIntegration) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.

### HasServiceAccountEmail

`func (o *GcpIntegration) HasServiceAccountEmail() bool`

HasServiceAccountEmail returns a boolean if a field has been set.

### GetServiceNames

`func (o *GcpIntegration) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *GcpIntegration) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *GcpIntegration) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *GcpIntegration) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetSignupRedirectWithoutEntitlementEnabled

`func (o *GcpIntegration) GetSignupRedirectWithoutEntitlementEnabled() bool`

GetSignupRedirectWithoutEntitlementEnabled returns the SignupRedirectWithoutEntitlementEnabled field if non-nil, zero value otherwise.

### GetSignupRedirectWithoutEntitlementEnabledOk

`func (o *GcpIntegration) GetSignupRedirectWithoutEntitlementEnabledOk() (*bool, bool)`

GetSignupRedirectWithoutEntitlementEnabledOk returns a tuple with the SignupRedirectWithoutEntitlementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupRedirectWithoutEntitlementEnabled

`func (o *GcpIntegration) SetSignupRedirectWithoutEntitlementEnabled(v bool)`

SetSignupRedirectWithoutEntitlementEnabled sets SignupRedirectWithoutEntitlementEnabled field to given value.

### HasSignupRedirectWithoutEntitlementEnabled

`func (o *GcpIntegration) HasSignupRedirectWithoutEntitlementEnabled() bool`

HasSignupRedirectWithoutEntitlementEnabled returns a boolean if a field has been set.

### GetUsageMeteringDisabled

`func (o *GcpIntegration) GetUsageMeteringDisabled() bool`

GetUsageMeteringDisabled returns the UsageMeteringDisabled field if non-nil, zero value otherwise.

### GetUsageMeteringDisabledOk

`func (o *GcpIntegration) GetUsageMeteringDisabledOk() (*bool, bool)`

GetUsageMeteringDisabledOk returns a tuple with the UsageMeteringDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeteringDisabled

`func (o *GcpIntegration) SetUsageMeteringDisabled(v bool)`

SetUsageMeteringDisabled sets UsageMeteringDisabled field to given value.

### HasUsageMeteringDisabled

`func (o *GcpIntegration) HasUsageMeteringDisabled() bool`

HasUsageMeteringDisabled returns a boolean if a field has been set.

### GetWorkloadIdentityPoolId

`func (o *GcpIntegration) GetWorkloadIdentityPoolId() string`

GetWorkloadIdentityPoolId returns the WorkloadIdentityPoolId field if non-nil, zero value otherwise.

### GetWorkloadIdentityPoolIdOk

`func (o *GcpIntegration) GetWorkloadIdentityPoolIdOk() (*string, bool)`

GetWorkloadIdentityPoolIdOk returns a tuple with the WorkloadIdentityPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityPoolId

`func (o *GcpIntegration) SetWorkloadIdentityPoolId(v string)`

SetWorkloadIdentityPoolId sets WorkloadIdentityPoolId field to given value.

### HasWorkloadIdentityPoolId

`func (o *GcpIntegration) HasWorkloadIdentityPoolId() bool`

HasWorkloadIdentityPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


