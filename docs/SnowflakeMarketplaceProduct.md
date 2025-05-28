# SnowflakeMarketplaceProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalRegions** | Pointer to **string** |  | [optional] 
**ApplicationPackageName** | Pointer to **string** |  | [optional] 
**AttachedShare** | Pointer to **string** |  | [optional] 
**Autofulfillment** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CustomizedContactInfo** | Pointer to **string** |  | [optional] 
**DefaultPricingPlan** | Pointer to [**SnowflakeMarketplaceProductDefaultPricingPlan**](SnowflakeMarketplaceProductDefaultPricingPlan.md) |  | [optional] 
**DetailedTargetAccounts** | Pointer to [**[]SnowflakeMarketplaceProductDetailedTargetAccount**](SnowflakeMarketplaceProductDetailedTargetAccount.md) |  | [optional] 
**Distribution** | Pointer to **string** |  | [optional] 
**EvaluationPlan** | Pointer to **string** |  | [optional] 
**FirstPublishedOn** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **string** |  | [optional] 
**FulfillmentType** | Pointer to **string** |  | [optional] 
**GlobalName** | Pointer to **string** |  | [optional] 
**IsMountlessQueryable** | Pointer to **bool** |  | [optional] 
**LastApprovedOn** | Pointer to **string** |  | [optional] 
**LastPublishedOn** | Pointer to **string** |  | [optional] 
**LastSubmittedOn** | Pointer to **string** |  | [optional] 
**ListingType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**SnowflakeMarketplaceProductMetadata**](SnowflakeMarketplaceProductMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PricingPlans** | Pointer to [**[]SnowflakeMarketplaceProductPricingPlan**](SnowflakeMarketplaceProductPricingPlan.md) |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**ProfileName** | Pointer to **string** |  | [optional] 
**PublishOnApproval** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to **string** |  | [optional] 
**RejectedOn** | Pointer to **string** |  | [optional] 
**RejectedReason** | Pointer to **string** |  | [optional] 
**ReplicationSchedule** | Pointer to **string** |  | [optional] 
**RetiredOn** | Pointer to **string** |  | [optional] 
**ScheduledDropTime** | Pointer to **string** |  | [optional] 
**ShareType** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TargetAccounts** | Pointer to **string** |  | [optional] 
**TrialDetails** | Pointer to [**SnowflakeMarketplaceTrialDetails**](SnowflakeMarketplaceTrialDetails.md) |  | [optional] 
**UnpublishedByAdminReason** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewSnowflakeMarketplaceProduct

`func NewSnowflakeMarketplaceProduct() *SnowflakeMarketplaceProduct`

NewSnowflakeMarketplaceProduct instantiates a new SnowflakeMarketplaceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceProductWithDefaults

`func NewSnowflakeMarketplaceProductWithDefaults() *SnowflakeMarketplaceProduct`

NewSnowflakeMarketplaceProductWithDefaults instantiates a new SnowflakeMarketplaceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalRegions

`func (o *SnowflakeMarketplaceProduct) GetAdditionalRegions() string`

GetAdditionalRegions returns the AdditionalRegions field if non-nil, zero value otherwise.

### GetAdditionalRegionsOk

`func (o *SnowflakeMarketplaceProduct) GetAdditionalRegionsOk() (*string, bool)`

GetAdditionalRegionsOk returns a tuple with the AdditionalRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRegions

`func (o *SnowflakeMarketplaceProduct) SetAdditionalRegions(v string)`

SetAdditionalRegions sets AdditionalRegions field to given value.

### HasAdditionalRegions

`func (o *SnowflakeMarketplaceProduct) HasAdditionalRegions() bool`

HasAdditionalRegions returns a boolean if a field has been set.

### GetApplicationPackageName

`func (o *SnowflakeMarketplaceProduct) GetApplicationPackageName() string`

GetApplicationPackageName returns the ApplicationPackageName field if non-nil, zero value otherwise.

### GetApplicationPackageNameOk

`func (o *SnowflakeMarketplaceProduct) GetApplicationPackageNameOk() (*string, bool)`

GetApplicationPackageNameOk returns a tuple with the ApplicationPackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPackageName

`func (o *SnowflakeMarketplaceProduct) SetApplicationPackageName(v string)`

SetApplicationPackageName sets ApplicationPackageName field to given value.

### HasApplicationPackageName

`func (o *SnowflakeMarketplaceProduct) HasApplicationPackageName() bool`

HasApplicationPackageName returns a boolean if a field has been set.

### GetAttachedShare

`func (o *SnowflakeMarketplaceProduct) GetAttachedShare() string`

GetAttachedShare returns the AttachedShare field if non-nil, zero value otherwise.

### GetAttachedShareOk

`func (o *SnowflakeMarketplaceProduct) GetAttachedShareOk() (*string, bool)`

GetAttachedShareOk returns a tuple with the AttachedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedShare

`func (o *SnowflakeMarketplaceProduct) SetAttachedShare(v string)`

SetAttachedShare sets AttachedShare field to given value.

### HasAttachedShare

`func (o *SnowflakeMarketplaceProduct) HasAttachedShare() bool`

HasAttachedShare returns a boolean if a field has been set.

### GetAutofulfillment

`func (o *SnowflakeMarketplaceProduct) GetAutofulfillment() bool`

GetAutofulfillment returns the Autofulfillment field if non-nil, zero value otherwise.

### GetAutofulfillmentOk

`func (o *SnowflakeMarketplaceProduct) GetAutofulfillmentOk() (*bool, bool)`

GetAutofulfillmentOk returns a tuple with the Autofulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofulfillment

`func (o *SnowflakeMarketplaceProduct) SetAutofulfillment(v bool)`

SetAutofulfillment sets Autofulfillment field to given value.

### HasAutofulfillment

`func (o *SnowflakeMarketplaceProduct) HasAutofulfillment() bool`

HasAutofulfillment returns a boolean if a field has been set.

### GetComment

`func (o *SnowflakeMarketplaceProduct) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SnowflakeMarketplaceProduct) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SnowflakeMarketplaceProduct) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SnowflakeMarketplaceProduct) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedOn

`func (o *SnowflakeMarketplaceProduct) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *SnowflakeMarketplaceProduct) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *SnowflakeMarketplaceProduct) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *SnowflakeMarketplaceProduct) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCustomizedContactInfo

`func (o *SnowflakeMarketplaceProduct) GetCustomizedContactInfo() string`

GetCustomizedContactInfo returns the CustomizedContactInfo field if non-nil, zero value otherwise.

### GetCustomizedContactInfoOk

`func (o *SnowflakeMarketplaceProduct) GetCustomizedContactInfoOk() (*string, bool)`

GetCustomizedContactInfoOk returns a tuple with the CustomizedContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedContactInfo

`func (o *SnowflakeMarketplaceProduct) SetCustomizedContactInfo(v string)`

SetCustomizedContactInfo sets CustomizedContactInfo field to given value.

### HasCustomizedContactInfo

`func (o *SnowflakeMarketplaceProduct) HasCustomizedContactInfo() bool`

HasCustomizedContactInfo returns a boolean if a field has been set.

### GetDefaultPricingPlan

`func (o *SnowflakeMarketplaceProduct) GetDefaultPricingPlan() SnowflakeMarketplaceProductDefaultPricingPlan`

GetDefaultPricingPlan returns the DefaultPricingPlan field if non-nil, zero value otherwise.

### GetDefaultPricingPlanOk

`func (o *SnowflakeMarketplaceProduct) GetDefaultPricingPlanOk() (*SnowflakeMarketplaceProductDefaultPricingPlan, bool)`

GetDefaultPricingPlanOk returns a tuple with the DefaultPricingPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPricingPlan

`func (o *SnowflakeMarketplaceProduct) SetDefaultPricingPlan(v SnowflakeMarketplaceProductDefaultPricingPlan)`

SetDefaultPricingPlan sets DefaultPricingPlan field to given value.

### HasDefaultPricingPlan

`func (o *SnowflakeMarketplaceProduct) HasDefaultPricingPlan() bool`

HasDefaultPricingPlan returns a boolean if a field has been set.

### GetDetailedTargetAccounts

`func (o *SnowflakeMarketplaceProduct) GetDetailedTargetAccounts() []SnowflakeMarketplaceProductDetailedTargetAccount`

GetDetailedTargetAccounts returns the DetailedTargetAccounts field if non-nil, zero value otherwise.

### GetDetailedTargetAccountsOk

`func (o *SnowflakeMarketplaceProduct) GetDetailedTargetAccountsOk() (*[]SnowflakeMarketplaceProductDetailedTargetAccount, bool)`

GetDetailedTargetAccountsOk returns a tuple with the DetailedTargetAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTargetAccounts

`func (o *SnowflakeMarketplaceProduct) SetDetailedTargetAccounts(v []SnowflakeMarketplaceProductDetailedTargetAccount)`

SetDetailedTargetAccounts sets DetailedTargetAccounts field to given value.

### HasDetailedTargetAccounts

`func (o *SnowflakeMarketplaceProduct) HasDetailedTargetAccounts() bool`

HasDetailedTargetAccounts returns a boolean if a field has been set.

### GetDistribution

`func (o *SnowflakeMarketplaceProduct) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *SnowflakeMarketplaceProduct) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *SnowflakeMarketplaceProduct) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *SnowflakeMarketplaceProduct) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetEvaluationPlan

`func (o *SnowflakeMarketplaceProduct) GetEvaluationPlan() string`

GetEvaluationPlan returns the EvaluationPlan field if non-nil, zero value otherwise.

### GetEvaluationPlanOk

`func (o *SnowflakeMarketplaceProduct) GetEvaluationPlanOk() (*string, bool)`

GetEvaluationPlanOk returns a tuple with the EvaluationPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPlan

`func (o *SnowflakeMarketplaceProduct) SetEvaluationPlan(v string)`

SetEvaluationPlan sets EvaluationPlan field to given value.

### HasEvaluationPlan

`func (o *SnowflakeMarketplaceProduct) HasEvaluationPlan() bool`

HasEvaluationPlan returns a boolean if a field has been set.

### GetFirstPublishedOn

`func (o *SnowflakeMarketplaceProduct) GetFirstPublishedOn() string`

GetFirstPublishedOn returns the FirstPublishedOn field if non-nil, zero value otherwise.

### GetFirstPublishedOnOk

`func (o *SnowflakeMarketplaceProduct) GetFirstPublishedOnOk() (*string, bool)`

GetFirstPublishedOnOk returns a tuple with the FirstPublishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPublishedOn

`func (o *SnowflakeMarketplaceProduct) SetFirstPublishedOn(v string)`

SetFirstPublishedOn sets FirstPublishedOn field to given value.

### HasFirstPublishedOn

`func (o *SnowflakeMarketplaceProduct) HasFirstPublishedOn() bool`

HasFirstPublishedOn returns a boolean if a field has been set.

### GetFlags

`func (o *SnowflakeMarketplaceProduct) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SnowflakeMarketplaceProduct) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SnowflakeMarketplaceProduct) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SnowflakeMarketplaceProduct) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFulfillmentType

`func (o *SnowflakeMarketplaceProduct) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *SnowflakeMarketplaceProduct) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *SnowflakeMarketplaceProduct) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *SnowflakeMarketplaceProduct) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetGlobalName

`func (o *SnowflakeMarketplaceProduct) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *SnowflakeMarketplaceProduct) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *SnowflakeMarketplaceProduct) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *SnowflakeMarketplaceProduct) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetIsMountlessQueryable

`func (o *SnowflakeMarketplaceProduct) GetIsMountlessQueryable() bool`

GetIsMountlessQueryable returns the IsMountlessQueryable field if non-nil, zero value otherwise.

### GetIsMountlessQueryableOk

`func (o *SnowflakeMarketplaceProduct) GetIsMountlessQueryableOk() (*bool, bool)`

GetIsMountlessQueryableOk returns a tuple with the IsMountlessQueryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountlessQueryable

`func (o *SnowflakeMarketplaceProduct) SetIsMountlessQueryable(v bool)`

SetIsMountlessQueryable sets IsMountlessQueryable field to given value.

### HasIsMountlessQueryable

`func (o *SnowflakeMarketplaceProduct) HasIsMountlessQueryable() bool`

HasIsMountlessQueryable returns a boolean if a field has been set.

### GetLastApprovedOn

`func (o *SnowflakeMarketplaceProduct) GetLastApprovedOn() string`

GetLastApprovedOn returns the LastApprovedOn field if non-nil, zero value otherwise.

### GetLastApprovedOnOk

`func (o *SnowflakeMarketplaceProduct) GetLastApprovedOnOk() (*string, bool)`

GetLastApprovedOnOk returns a tuple with the LastApprovedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApprovedOn

`func (o *SnowflakeMarketplaceProduct) SetLastApprovedOn(v string)`

SetLastApprovedOn sets LastApprovedOn field to given value.

### HasLastApprovedOn

`func (o *SnowflakeMarketplaceProduct) HasLastApprovedOn() bool`

HasLastApprovedOn returns a boolean if a field has been set.

### GetLastPublishedOn

`func (o *SnowflakeMarketplaceProduct) GetLastPublishedOn() string`

GetLastPublishedOn returns the LastPublishedOn field if non-nil, zero value otherwise.

### GetLastPublishedOnOk

`func (o *SnowflakeMarketplaceProduct) GetLastPublishedOnOk() (*string, bool)`

GetLastPublishedOnOk returns a tuple with the LastPublishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublishedOn

`func (o *SnowflakeMarketplaceProduct) SetLastPublishedOn(v string)`

SetLastPublishedOn sets LastPublishedOn field to given value.

### HasLastPublishedOn

`func (o *SnowflakeMarketplaceProduct) HasLastPublishedOn() bool`

HasLastPublishedOn returns a boolean if a field has been set.

### GetLastSubmittedOn

`func (o *SnowflakeMarketplaceProduct) GetLastSubmittedOn() string`

GetLastSubmittedOn returns the LastSubmittedOn field if non-nil, zero value otherwise.

### GetLastSubmittedOnOk

`func (o *SnowflakeMarketplaceProduct) GetLastSubmittedOnOk() (*string, bool)`

GetLastSubmittedOnOk returns a tuple with the LastSubmittedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSubmittedOn

`func (o *SnowflakeMarketplaceProduct) SetLastSubmittedOn(v string)`

SetLastSubmittedOn sets LastSubmittedOn field to given value.

### HasLastSubmittedOn

`func (o *SnowflakeMarketplaceProduct) HasLastSubmittedOn() bool`

HasLastSubmittedOn returns a boolean if a field has been set.

### GetListingType

`func (o *SnowflakeMarketplaceProduct) GetListingType() string`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *SnowflakeMarketplaceProduct) GetListingTypeOk() (*string, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *SnowflakeMarketplaceProduct) SetListingType(v string)`

SetListingType sets ListingType field to given value.

### HasListingType

`func (o *SnowflakeMarketplaceProduct) HasListingType() bool`

HasListingType returns a boolean if a field has been set.

### GetMetadata

`func (o *SnowflakeMarketplaceProduct) GetMetadata() SnowflakeMarketplaceProductMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnowflakeMarketplaceProduct) GetMetadataOk() (*SnowflakeMarketplaceProductMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnowflakeMarketplaceProduct) SetMetadata(v SnowflakeMarketplaceProductMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SnowflakeMarketplaceProduct) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SnowflakeMarketplaceProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnowflakeMarketplaceProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnowflakeMarketplaceProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnowflakeMarketplaceProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricingPlans

`func (o *SnowflakeMarketplaceProduct) GetPricingPlans() []SnowflakeMarketplaceProductPricingPlan`

GetPricingPlans returns the PricingPlans field if non-nil, zero value otherwise.

### GetPricingPlansOk

`func (o *SnowflakeMarketplaceProduct) GetPricingPlansOk() (*[]SnowflakeMarketplaceProductPricingPlan, bool)`

GetPricingPlansOk returns a tuple with the PricingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlans

`func (o *SnowflakeMarketplaceProduct) SetPricingPlans(v []SnowflakeMarketplaceProductPricingPlan)`

SetPricingPlans sets PricingPlans field to given value.

### HasPricingPlans

`func (o *SnowflakeMarketplaceProduct) HasPricingPlans() bool`

HasPricingPlans returns a boolean if a field has been set.

### GetPrivate

`func (o *SnowflakeMarketplaceProduct) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SnowflakeMarketplaceProduct) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SnowflakeMarketplaceProduct) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SnowflakeMarketplaceProduct) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetProfileName

`func (o *SnowflakeMarketplaceProduct) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *SnowflakeMarketplaceProduct) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *SnowflakeMarketplaceProduct) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *SnowflakeMarketplaceProduct) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetPublishOnApproval

`func (o *SnowflakeMarketplaceProduct) GetPublishOnApproval() bool`

GetPublishOnApproval returns the PublishOnApproval field if non-nil, zero value otherwise.

### GetPublishOnApprovalOk

`func (o *SnowflakeMarketplaceProduct) GetPublishOnApprovalOk() (*bool, bool)`

GetPublishOnApprovalOk returns a tuple with the PublishOnApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishOnApproval

`func (o *SnowflakeMarketplaceProduct) SetPublishOnApproval(v bool)`

SetPublishOnApproval sets PublishOnApproval field to given value.

### HasPublishOnApproval

`func (o *SnowflakeMarketplaceProduct) HasPublishOnApproval() bool`

HasPublishOnApproval returns a boolean if a field has been set.

### GetRegions

`func (o *SnowflakeMarketplaceProduct) GetRegions() string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *SnowflakeMarketplaceProduct) GetRegionsOk() (*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *SnowflakeMarketplaceProduct) SetRegions(v string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *SnowflakeMarketplaceProduct) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRejectedOn

`func (o *SnowflakeMarketplaceProduct) GetRejectedOn() string`

GetRejectedOn returns the RejectedOn field if non-nil, zero value otherwise.

### GetRejectedOnOk

`func (o *SnowflakeMarketplaceProduct) GetRejectedOnOk() (*string, bool)`

GetRejectedOnOk returns a tuple with the RejectedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedOn

`func (o *SnowflakeMarketplaceProduct) SetRejectedOn(v string)`

SetRejectedOn sets RejectedOn field to given value.

### HasRejectedOn

`func (o *SnowflakeMarketplaceProduct) HasRejectedOn() bool`

HasRejectedOn returns a boolean if a field has been set.

### GetRejectedReason

`func (o *SnowflakeMarketplaceProduct) GetRejectedReason() string`

GetRejectedReason returns the RejectedReason field if non-nil, zero value otherwise.

### GetRejectedReasonOk

`func (o *SnowflakeMarketplaceProduct) GetRejectedReasonOk() (*string, bool)`

GetRejectedReasonOk returns a tuple with the RejectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReason

`func (o *SnowflakeMarketplaceProduct) SetRejectedReason(v string)`

SetRejectedReason sets RejectedReason field to given value.

### HasRejectedReason

`func (o *SnowflakeMarketplaceProduct) HasRejectedReason() bool`

HasRejectedReason returns a boolean if a field has been set.

### GetReplicationSchedule

`func (o *SnowflakeMarketplaceProduct) GetReplicationSchedule() string`

GetReplicationSchedule returns the ReplicationSchedule field if non-nil, zero value otherwise.

### GetReplicationScheduleOk

`func (o *SnowflakeMarketplaceProduct) GetReplicationScheduleOk() (*string, bool)`

GetReplicationScheduleOk returns a tuple with the ReplicationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSchedule

`func (o *SnowflakeMarketplaceProduct) SetReplicationSchedule(v string)`

SetReplicationSchedule sets ReplicationSchedule field to given value.

### HasReplicationSchedule

`func (o *SnowflakeMarketplaceProduct) HasReplicationSchedule() bool`

HasReplicationSchedule returns a boolean if a field has been set.

### GetRetiredOn

`func (o *SnowflakeMarketplaceProduct) GetRetiredOn() string`

GetRetiredOn returns the RetiredOn field if non-nil, zero value otherwise.

### GetRetiredOnOk

`func (o *SnowflakeMarketplaceProduct) GetRetiredOnOk() (*string, bool)`

GetRetiredOnOk returns a tuple with the RetiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetiredOn

`func (o *SnowflakeMarketplaceProduct) SetRetiredOn(v string)`

SetRetiredOn sets RetiredOn field to given value.

### HasRetiredOn

`func (o *SnowflakeMarketplaceProduct) HasRetiredOn() bool`

HasRetiredOn returns a boolean if a field has been set.

### GetScheduledDropTime

`func (o *SnowflakeMarketplaceProduct) GetScheduledDropTime() string`

GetScheduledDropTime returns the ScheduledDropTime field if non-nil, zero value otherwise.

### GetScheduledDropTimeOk

`func (o *SnowflakeMarketplaceProduct) GetScheduledDropTimeOk() (*string, bool)`

GetScheduledDropTimeOk returns a tuple with the ScheduledDropTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDropTime

`func (o *SnowflakeMarketplaceProduct) SetScheduledDropTime(v string)`

SetScheduledDropTime sets ScheduledDropTime field to given value.

### HasScheduledDropTime

`func (o *SnowflakeMarketplaceProduct) HasScheduledDropTime() bool`

HasScheduledDropTime returns a boolean if a field has been set.

### GetShareType

`func (o *SnowflakeMarketplaceProduct) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *SnowflakeMarketplaceProduct) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *SnowflakeMarketplaceProduct) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *SnowflakeMarketplaceProduct) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### GetState

`func (o *SnowflakeMarketplaceProduct) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnowflakeMarketplaceProduct) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SnowflakeMarketplaceProduct) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SnowflakeMarketplaceProduct) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTargetAccounts

`func (o *SnowflakeMarketplaceProduct) GetTargetAccounts() string`

GetTargetAccounts returns the TargetAccounts field if non-nil, zero value otherwise.

### GetTargetAccountsOk

`func (o *SnowflakeMarketplaceProduct) GetTargetAccountsOk() (*string, bool)`

GetTargetAccountsOk returns a tuple with the TargetAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccounts

`func (o *SnowflakeMarketplaceProduct) SetTargetAccounts(v string)`

SetTargetAccounts sets TargetAccounts field to given value.

### HasTargetAccounts

`func (o *SnowflakeMarketplaceProduct) HasTargetAccounts() bool`

HasTargetAccounts returns a boolean if a field has been set.

### GetTrialDetails

`func (o *SnowflakeMarketplaceProduct) GetTrialDetails() SnowflakeMarketplaceTrialDetails`

GetTrialDetails returns the TrialDetails field if non-nil, zero value otherwise.

### GetTrialDetailsOk

`func (o *SnowflakeMarketplaceProduct) GetTrialDetailsOk() (*SnowflakeMarketplaceTrialDetails, bool)`

GetTrialDetailsOk returns a tuple with the TrialDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDetails

`func (o *SnowflakeMarketplaceProduct) SetTrialDetails(v SnowflakeMarketplaceTrialDetails)`

SetTrialDetails sets TrialDetails field to given value.

### HasTrialDetails

`func (o *SnowflakeMarketplaceProduct) HasTrialDetails() bool`

HasTrialDetails returns a boolean if a field has been set.

### GetUnpublishedByAdminReason

`func (o *SnowflakeMarketplaceProduct) GetUnpublishedByAdminReason() string`

GetUnpublishedByAdminReason returns the UnpublishedByAdminReason field if non-nil, zero value otherwise.

### GetUnpublishedByAdminReasonOk

`func (o *SnowflakeMarketplaceProduct) GetUnpublishedByAdminReasonOk() (*string, bool)`

GetUnpublishedByAdminReasonOk returns a tuple with the UnpublishedByAdminReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedByAdminReason

`func (o *SnowflakeMarketplaceProduct) SetUnpublishedByAdminReason(v string)`

SetUnpublishedByAdminReason sets UnpublishedByAdminReason field to given value.

### HasUnpublishedByAdminReason

`func (o *SnowflakeMarketplaceProduct) HasUnpublishedByAdminReason() bool`

HasUnpublishedByAdminReason returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *SnowflakeMarketplaceProduct) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *SnowflakeMarketplaceProduct) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *SnowflakeMarketplaceProduct) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *SnowflakeMarketplaceProduct) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


