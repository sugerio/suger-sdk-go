# AlibabaMarketplaceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliyunUid** | Pointer to **string** | The uid of the seller&#39;s main Alibaba Account. | [optional] 
**AutoCancelSuspendedEntitlementsEnabled** | Pointer to **bool** | If true, the suspended entitlements will be automatically canceled after the specified days. | [optional] 
**Credential** | Pointer to [**AlibabaIntegrationCredential**](AlibabaIntegrationCredential.md) |  | [optional] 
**DaysUntilAutoCancelSuspendedEntitlements** | Pointer to **int32** | Specifies the number of days after which an entitlement in suspended status will be automatically canceled. Only applicable if AutoCancelSuspendedEntitlementsEnabled is true. If the value is 0 or negative, the auto cancel feature is disabled. | [optional] 
**ProductCodes** | Pointer to **[]string** | The codes of the products that the seller is selling on Alibaba Marketplace. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to store the AlibabaIntegrationCredential in AWS Secret manager. For internal usage only. | [optional] 
**UsageMeteringDisabled** | Pointer to **bool** | Disable Usage Metering to Alibaba Cloud Marketplace. If true, Suger stop to report usage records to Alibaba Cloud Marketplace. | [optional] 

## Methods

### NewAlibabaMarketplaceIntegration

`func NewAlibabaMarketplaceIntegration() *AlibabaMarketplaceIntegration`

NewAlibabaMarketplaceIntegration instantiates a new AlibabaMarketplaceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlibabaMarketplaceIntegrationWithDefaults

`func NewAlibabaMarketplaceIntegrationWithDefaults() *AlibabaMarketplaceIntegration`

NewAlibabaMarketplaceIntegrationWithDefaults instantiates a new AlibabaMarketplaceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliyunUid

`func (o *AlibabaMarketplaceIntegration) GetAliyunUid() string`

GetAliyunUid returns the AliyunUid field if non-nil, zero value otherwise.

### GetAliyunUidOk

`func (o *AlibabaMarketplaceIntegration) GetAliyunUidOk() (*string, bool)`

GetAliyunUidOk returns a tuple with the AliyunUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliyunUid

`func (o *AlibabaMarketplaceIntegration) SetAliyunUid(v string)`

SetAliyunUid sets AliyunUid field to given value.

### HasAliyunUid

`func (o *AlibabaMarketplaceIntegration) HasAliyunUid() bool`

HasAliyunUid returns a boolean if a field has been set.

### GetAutoCancelSuspendedEntitlementsEnabled

`func (o *AlibabaMarketplaceIntegration) GetAutoCancelSuspendedEntitlementsEnabled() bool`

GetAutoCancelSuspendedEntitlementsEnabled returns the AutoCancelSuspendedEntitlementsEnabled field if non-nil, zero value otherwise.

### GetAutoCancelSuspendedEntitlementsEnabledOk

`func (o *AlibabaMarketplaceIntegration) GetAutoCancelSuspendedEntitlementsEnabledOk() (*bool, bool)`

GetAutoCancelSuspendedEntitlementsEnabledOk returns a tuple with the AutoCancelSuspendedEntitlementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSuspendedEntitlementsEnabled

`func (o *AlibabaMarketplaceIntegration) SetAutoCancelSuspendedEntitlementsEnabled(v bool)`

SetAutoCancelSuspendedEntitlementsEnabled sets AutoCancelSuspendedEntitlementsEnabled field to given value.

### HasAutoCancelSuspendedEntitlementsEnabled

`func (o *AlibabaMarketplaceIntegration) HasAutoCancelSuspendedEntitlementsEnabled() bool`

HasAutoCancelSuspendedEntitlementsEnabled returns a boolean if a field has been set.

### GetCredential

`func (o *AlibabaMarketplaceIntegration) GetCredential() AlibabaIntegrationCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AlibabaMarketplaceIntegration) GetCredentialOk() (*AlibabaIntegrationCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AlibabaMarketplaceIntegration) SetCredential(v AlibabaIntegrationCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AlibabaMarketplaceIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDaysUntilAutoCancelSuspendedEntitlements

`func (o *AlibabaMarketplaceIntegration) GetDaysUntilAutoCancelSuspendedEntitlements() int32`

GetDaysUntilAutoCancelSuspendedEntitlements returns the DaysUntilAutoCancelSuspendedEntitlements field if non-nil, zero value otherwise.

### GetDaysUntilAutoCancelSuspendedEntitlementsOk

`func (o *AlibabaMarketplaceIntegration) GetDaysUntilAutoCancelSuspendedEntitlementsOk() (*int32, bool)`

GetDaysUntilAutoCancelSuspendedEntitlementsOk returns a tuple with the DaysUntilAutoCancelSuspendedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilAutoCancelSuspendedEntitlements

`func (o *AlibabaMarketplaceIntegration) SetDaysUntilAutoCancelSuspendedEntitlements(v int32)`

SetDaysUntilAutoCancelSuspendedEntitlements sets DaysUntilAutoCancelSuspendedEntitlements field to given value.

### HasDaysUntilAutoCancelSuspendedEntitlements

`func (o *AlibabaMarketplaceIntegration) HasDaysUntilAutoCancelSuspendedEntitlements() bool`

HasDaysUntilAutoCancelSuspendedEntitlements returns a boolean if a field has been set.

### GetProductCodes

`func (o *AlibabaMarketplaceIntegration) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *AlibabaMarketplaceIntegration) GetProductCodesOk() (*[]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCodes

`func (o *AlibabaMarketplaceIntegration) SetProductCodes(v []string)`

SetProductCodes sets ProductCodes field to given value.

### HasProductCodes

`func (o *AlibabaMarketplaceIntegration) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.

### GetSecretKey

`func (o *AlibabaMarketplaceIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AlibabaMarketplaceIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AlibabaMarketplaceIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AlibabaMarketplaceIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUsageMeteringDisabled

`func (o *AlibabaMarketplaceIntegration) GetUsageMeteringDisabled() bool`

GetUsageMeteringDisabled returns the UsageMeteringDisabled field if non-nil, zero value otherwise.

### GetUsageMeteringDisabledOk

`func (o *AlibabaMarketplaceIntegration) GetUsageMeteringDisabledOk() (*bool, bool)`

GetUsageMeteringDisabledOk returns a tuple with the UsageMeteringDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeteringDisabled

`func (o *AlibabaMarketplaceIntegration) SetUsageMeteringDisabled(v bool)`

SetUsageMeteringDisabled sets UsageMeteringDisabled field to given value.

### HasUsageMeteringDisabled

`func (o *AlibabaMarketplaceIntegration) HasUsageMeteringDisabled() bool`

HasUsageMeteringDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


