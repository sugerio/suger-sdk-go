# GcpMarketplaceResellerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAccountId** | Pointer to **string** |  | [optional] 
**BillingAccountNickname** | Pointer to **string** |  | [optional] 
**BillingAccountOrgDisplayName** | Pointer to **string** |  | [optional] 
**BillingAccountType** | Pointer to **string** |  | [optional] 
**NotesToReseller** | Pointer to **string** |  | [optional] 
**PartnerAccountName** | Pointer to **string** | In the format of \&quot;\&quot;organizations/{GcpOrganizationID}/partnerAccounts/{partnerAccountID}\&quot; | [optional] 
**ResellOfferTemplateId** | Pointer to **string** |  | [optional] 
**ResellerContactEmail** | Pointer to **string** |  | [optional] 
**ResellerContactName** | Pointer to **string** |  | [optional] 
**ResellerPrivateOfferPlanId** | Pointer to **string** |  | [optional] 
**ResellerPrivateOfferPlanScope** | Pointer to **string** |  | [optional] 
**SubBillingAccount** | Pointer to **string** | In the format of \&quot;billingAccounts/...\&quot; | [optional] 

## Methods

### NewGcpMarketplaceResellerInfo

`func NewGcpMarketplaceResellerInfo() *GcpMarketplaceResellerInfo`

NewGcpMarketplaceResellerInfo instantiates a new GcpMarketplaceResellerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceResellerInfoWithDefaults

`func NewGcpMarketplaceResellerInfoWithDefaults() *GcpMarketplaceResellerInfo`

NewGcpMarketplaceResellerInfoWithDefaults instantiates a new GcpMarketplaceResellerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAccountId

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *GcpMarketplaceResellerInfo) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *GcpMarketplaceResellerInfo) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetBillingAccountNickname

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountNickname() string`

GetBillingAccountNickname returns the BillingAccountNickname field if non-nil, zero value otherwise.

### GetBillingAccountNicknameOk

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountNicknameOk() (*string, bool)`

GetBillingAccountNicknameOk returns a tuple with the BillingAccountNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountNickname

`func (o *GcpMarketplaceResellerInfo) SetBillingAccountNickname(v string)`

SetBillingAccountNickname sets BillingAccountNickname field to given value.

### HasBillingAccountNickname

`func (o *GcpMarketplaceResellerInfo) HasBillingAccountNickname() bool`

HasBillingAccountNickname returns a boolean if a field has been set.

### GetBillingAccountOrgDisplayName

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountOrgDisplayName() string`

GetBillingAccountOrgDisplayName returns the BillingAccountOrgDisplayName field if non-nil, zero value otherwise.

### GetBillingAccountOrgDisplayNameOk

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountOrgDisplayNameOk() (*string, bool)`

GetBillingAccountOrgDisplayNameOk returns a tuple with the BillingAccountOrgDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountOrgDisplayName

`func (o *GcpMarketplaceResellerInfo) SetBillingAccountOrgDisplayName(v string)`

SetBillingAccountOrgDisplayName sets BillingAccountOrgDisplayName field to given value.

### HasBillingAccountOrgDisplayName

`func (o *GcpMarketplaceResellerInfo) HasBillingAccountOrgDisplayName() bool`

HasBillingAccountOrgDisplayName returns a boolean if a field has been set.

### GetBillingAccountType

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountType() string`

GetBillingAccountType returns the BillingAccountType field if non-nil, zero value otherwise.

### GetBillingAccountTypeOk

`func (o *GcpMarketplaceResellerInfo) GetBillingAccountTypeOk() (*string, bool)`

GetBillingAccountTypeOk returns a tuple with the BillingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountType

`func (o *GcpMarketplaceResellerInfo) SetBillingAccountType(v string)`

SetBillingAccountType sets BillingAccountType field to given value.

### HasBillingAccountType

`func (o *GcpMarketplaceResellerInfo) HasBillingAccountType() bool`

HasBillingAccountType returns a boolean if a field has been set.

### GetNotesToReseller

`func (o *GcpMarketplaceResellerInfo) GetNotesToReseller() string`

GetNotesToReseller returns the NotesToReseller field if non-nil, zero value otherwise.

### GetNotesToResellerOk

`func (o *GcpMarketplaceResellerInfo) GetNotesToResellerOk() (*string, bool)`

GetNotesToResellerOk returns a tuple with the NotesToReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesToReseller

`func (o *GcpMarketplaceResellerInfo) SetNotesToReseller(v string)`

SetNotesToReseller sets NotesToReseller field to given value.

### HasNotesToReseller

`func (o *GcpMarketplaceResellerInfo) HasNotesToReseller() bool`

HasNotesToReseller returns a boolean if a field has been set.

### GetPartnerAccountName

`func (o *GcpMarketplaceResellerInfo) GetPartnerAccountName() string`

GetPartnerAccountName returns the PartnerAccountName field if non-nil, zero value otherwise.

### GetPartnerAccountNameOk

`func (o *GcpMarketplaceResellerInfo) GetPartnerAccountNameOk() (*string, bool)`

GetPartnerAccountNameOk returns a tuple with the PartnerAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAccountName

`func (o *GcpMarketplaceResellerInfo) SetPartnerAccountName(v string)`

SetPartnerAccountName sets PartnerAccountName field to given value.

### HasPartnerAccountName

`func (o *GcpMarketplaceResellerInfo) HasPartnerAccountName() bool`

HasPartnerAccountName returns a boolean if a field has been set.

### GetResellOfferTemplateId

`func (o *GcpMarketplaceResellerInfo) GetResellOfferTemplateId() string`

GetResellOfferTemplateId returns the ResellOfferTemplateId field if non-nil, zero value otherwise.

### GetResellOfferTemplateIdOk

`func (o *GcpMarketplaceResellerInfo) GetResellOfferTemplateIdOk() (*string, bool)`

GetResellOfferTemplateIdOk returns a tuple with the ResellOfferTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellOfferTemplateId

`func (o *GcpMarketplaceResellerInfo) SetResellOfferTemplateId(v string)`

SetResellOfferTemplateId sets ResellOfferTemplateId field to given value.

### HasResellOfferTemplateId

`func (o *GcpMarketplaceResellerInfo) HasResellOfferTemplateId() bool`

HasResellOfferTemplateId returns a boolean if a field has been set.

### GetResellerContactEmail

`func (o *GcpMarketplaceResellerInfo) GetResellerContactEmail() string`

GetResellerContactEmail returns the ResellerContactEmail field if non-nil, zero value otherwise.

### GetResellerContactEmailOk

`func (o *GcpMarketplaceResellerInfo) GetResellerContactEmailOk() (*string, bool)`

GetResellerContactEmailOk returns a tuple with the ResellerContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerContactEmail

`func (o *GcpMarketplaceResellerInfo) SetResellerContactEmail(v string)`

SetResellerContactEmail sets ResellerContactEmail field to given value.

### HasResellerContactEmail

`func (o *GcpMarketplaceResellerInfo) HasResellerContactEmail() bool`

HasResellerContactEmail returns a boolean if a field has been set.

### GetResellerContactName

`func (o *GcpMarketplaceResellerInfo) GetResellerContactName() string`

GetResellerContactName returns the ResellerContactName field if non-nil, zero value otherwise.

### GetResellerContactNameOk

`func (o *GcpMarketplaceResellerInfo) GetResellerContactNameOk() (*string, bool)`

GetResellerContactNameOk returns a tuple with the ResellerContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerContactName

`func (o *GcpMarketplaceResellerInfo) SetResellerContactName(v string)`

SetResellerContactName sets ResellerContactName field to given value.

### HasResellerContactName

`func (o *GcpMarketplaceResellerInfo) HasResellerContactName() bool`

HasResellerContactName returns a boolean if a field has been set.

### GetResellerPrivateOfferPlanId

`func (o *GcpMarketplaceResellerInfo) GetResellerPrivateOfferPlanId() string`

GetResellerPrivateOfferPlanId returns the ResellerPrivateOfferPlanId field if non-nil, zero value otherwise.

### GetResellerPrivateOfferPlanIdOk

`func (o *GcpMarketplaceResellerInfo) GetResellerPrivateOfferPlanIdOk() (*string, bool)`

GetResellerPrivateOfferPlanIdOk returns a tuple with the ResellerPrivateOfferPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerPrivateOfferPlanId

`func (o *GcpMarketplaceResellerInfo) SetResellerPrivateOfferPlanId(v string)`

SetResellerPrivateOfferPlanId sets ResellerPrivateOfferPlanId field to given value.

### HasResellerPrivateOfferPlanId

`func (o *GcpMarketplaceResellerInfo) HasResellerPrivateOfferPlanId() bool`

HasResellerPrivateOfferPlanId returns a boolean if a field has been set.

### GetResellerPrivateOfferPlanScope

`func (o *GcpMarketplaceResellerInfo) GetResellerPrivateOfferPlanScope() string`

GetResellerPrivateOfferPlanScope returns the ResellerPrivateOfferPlanScope field if non-nil, zero value otherwise.

### GetResellerPrivateOfferPlanScopeOk

`func (o *GcpMarketplaceResellerInfo) GetResellerPrivateOfferPlanScopeOk() (*string, bool)`

GetResellerPrivateOfferPlanScopeOk returns a tuple with the ResellerPrivateOfferPlanScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerPrivateOfferPlanScope

`func (o *GcpMarketplaceResellerInfo) SetResellerPrivateOfferPlanScope(v string)`

SetResellerPrivateOfferPlanScope sets ResellerPrivateOfferPlanScope field to given value.

### HasResellerPrivateOfferPlanScope

`func (o *GcpMarketplaceResellerInfo) HasResellerPrivateOfferPlanScope() bool`

HasResellerPrivateOfferPlanScope returns a boolean if a field has been set.

### GetSubBillingAccount

`func (o *GcpMarketplaceResellerInfo) GetSubBillingAccount() string`

GetSubBillingAccount returns the SubBillingAccount field if non-nil, zero value otherwise.

### GetSubBillingAccountOk

`func (o *GcpMarketplaceResellerInfo) GetSubBillingAccountOk() (*string, bool)`

GetSubBillingAccountOk returns a tuple with the SubBillingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBillingAccount

`func (o *GcpMarketplaceResellerInfo) SetSubBillingAccount(v string)`

SetSubBillingAccount sets SubBillingAccount field to given value.

### HasSubBillingAccount

`func (o *GcpMarketplaceResellerInfo) HasSubBillingAccount() bool`

HasSubBillingAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


