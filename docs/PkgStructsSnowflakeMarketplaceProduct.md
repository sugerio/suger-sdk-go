# PkgStructsSnowflakeMarketplaceProduct

## Properties

 Name                        | Type                                                             | Description | Notes      
-----------------------------|------------------------------------------------------------------|-------------|------------
 **Comment**                 | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **CreatedOn**               | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md)     |             | [optional] 
 **DetailedTargetAccounts**  | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **Distribution**            | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **GlobalName**              | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **IsApplication**           | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **IsByRequest**             | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **IsLimitedTrial**          | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **IsMonetized**             | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **IsMountlessQueryable**    | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **IsTargeted**              | Pointer to [**DatabaseSqlNullBool**](DatabaseSqlNullBool.md)     |             | [optional] 
 **Owner**                   | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **OwnerRoleType**           | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **Profile**                 | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **PublishedOn**             | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md)     |             | [optional] 
 **Regions**                 | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **RejectedOn**              | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md)     |             | [optional] 
 **ReviewState**             | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **State**                   | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **Subtitle**                | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **TargetAccounts**          | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **Title**                   | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **UniformListingLocator**   | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **UpdatedOn**               | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md)     |             | [optional] 
 **Name**                    | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 
 **OrganizationProfileName** | Pointer to [**DatabaseSqlNullString**](DatabaseSqlNullString.md) |             | [optional] 

## Methods

### NewPkgStructsSnowflakeMarketplaceProduct

`func NewPkgStructsSnowflakeMarketplaceProduct() *PkgStructsSnowflakeMarketplaceProduct`

NewPkgStructsSnowflakeMarketplaceProduct instantiates a new PkgStructsSnowflakeMarketplaceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgStructsSnowflakeMarketplaceProductWithDefaults

`func NewPkgStructsSnowflakeMarketplaceProductWithDefaults() *PkgStructsSnowflakeMarketplaceProduct`

NewPkgStructsSnowflakeMarketplaceProductWithDefaults instantiates a new PkgStructsSnowflakeMarketplaceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetComment() DatabaseSqlNullString`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetCommentOk() (*DatabaseSqlNullString, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetComment(v DatabaseSqlNullString)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetCreatedOn() DatabaseSqlNullTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetCreatedOnOk() (*DatabaseSqlNullTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetCreatedOn(v DatabaseSqlNullTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDetailedTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetDetailedTargetAccounts() DatabaseSqlNullString`

GetDetailedTargetAccounts returns the DetailedTargetAccounts field if non-nil, zero value otherwise.

### GetDetailedTargetAccountsOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetDetailedTargetAccountsOk() (*DatabaseSqlNullString, bool)`

GetDetailedTargetAccountsOk returns a tuple with the DetailedTargetAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetDetailedTargetAccounts(v DatabaseSqlNullString)`

SetDetailedTargetAccounts sets DetailedTargetAccounts field to given value.

### HasDetailedTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasDetailedTargetAccounts() bool`

HasDetailedTargetAccounts returns a boolean if a field has been set.

### GetDistribution

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetDistribution() DatabaseSqlNullString`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetDistributionOk() (*DatabaseSqlNullString, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetDistribution(v DatabaseSqlNullString)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetGlobalName

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetGlobalName() DatabaseSqlNullString`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetGlobalNameOk() (*DatabaseSqlNullString, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetGlobalName(v DatabaseSqlNullString)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetIsApplication

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsApplication() DatabaseSqlNullBool`

GetIsApplication returns the IsApplication field if non-nil, zero value otherwise.

### GetIsApplicationOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsApplicationOk() (*DatabaseSqlNullBool, bool)`

GetIsApplicationOk returns a tuple with the IsApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApplication

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsApplication(v DatabaseSqlNullBool)`

SetIsApplication sets IsApplication field to given value.

### HasIsApplication

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsApplication() bool`

HasIsApplication returns a boolean if a field has been set.

### GetIsByRequest

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsByRequest() DatabaseSqlNullBool`

GetIsByRequest returns the IsByRequest field if non-nil, zero value otherwise.

### GetIsByRequestOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsByRequestOk() (*DatabaseSqlNullBool, bool)`

GetIsByRequestOk returns a tuple with the IsByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByRequest

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsByRequest(v DatabaseSqlNullBool)`

SetIsByRequest sets IsByRequest field to given value.

### HasIsByRequest

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsByRequest() bool`

HasIsByRequest returns a boolean if a field has been set.

### GetIsLimitedTrial

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsLimitedTrial() DatabaseSqlNullBool`

GetIsLimitedTrial returns the IsLimitedTrial field if non-nil, zero value otherwise.

### GetIsLimitedTrialOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsLimitedTrialOk() (*DatabaseSqlNullBool, bool)`

GetIsLimitedTrialOk returns a tuple with the IsLimitedTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimitedTrial

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsLimitedTrial(v DatabaseSqlNullBool)`

SetIsLimitedTrial sets IsLimitedTrial field to given value.

### HasIsLimitedTrial

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsLimitedTrial() bool`

HasIsLimitedTrial returns a boolean if a field has been set.

### GetIsMonetized

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsMonetized() DatabaseSqlNullBool`

GetIsMonetized returns the IsMonetized field if non-nil, zero value otherwise.

### GetIsMonetizedOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsMonetizedOk() (*DatabaseSqlNullBool, bool)`

GetIsMonetizedOk returns a tuple with the IsMonetized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonetized

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsMonetized(v DatabaseSqlNullBool)`

SetIsMonetized sets IsMonetized field to given value.

### HasIsMonetized

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsMonetized() bool`

HasIsMonetized returns a boolean if a field has been set.

### GetIsMountlessQueryable

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsMountlessQueryable() DatabaseSqlNullBool`

GetIsMountlessQueryable returns the IsMountlessQueryable field if non-nil, zero value otherwise.

### GetIsMountlessQueryableOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsMountlessQueryableOk() (*DatabaseSqlNullBool, bool)`

GetIsMountlessQueryableOk returns a tuple with the IsMountlessQueryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountlessQueryable

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsMountlessQueryable(v DatabaseSqlNullBool)`

SetIsMountlessQueryable sets IsMountlessQueryable field to given value.

### HasIsMountlessQueryable

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsMountlessQueryable() bool`

HasIsMountlessQueryable returns a boolean if a field has been set.

### GetIsTargeted

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsTargeted() DatabaseSqlNullBool`

GetIsTargeted returns the IsTargeted field if non-nil, zero value otherwise.

### GetIsTargetedOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetIsTargetedOk() (*DatabaseSqlNullBool, bool)`

GetIsTargetedOk returns a tuple with the IsTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTargeted

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetIsTargeted(v DatabaseSqlNullBool)`

SetIsTargeted sets IsTargeted field to given value.

### HasIsTargeted

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasIsTargeted() bool`

HasIsTargeted returns a boolean if a field has been set.

### GetOwner

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOwner() DatabaseSqlNullString`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOwnerOk() (*DatabaseSqlNullString, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetOwner(v DatabaseSqlNullString)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerRoleType

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOwnerRoleType() DatabaseSqlNullString`

GetOwnerRoleType returns the OwnerRoleType field if non-nil, zero value otherwise.

### GetOwnerRoleTypeOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOwnerRoleTypeOk() (*DatabaseSqlNullString, bool)`

GetOwnerRoleTypeOk returns a tuple with the OwnerRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleType

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetOwnerRoleType(v DatabaseSqlNullString)`

SetOwnerRoleType sets OwnerRoleType field to given value.

### HasOwnerRoleType

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasOwnerRoleType() bool`

HasOwnerRoleType returns a boolean if a field has been set.

### GetProfile

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetProfile() DatabaseSqlNullString`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetProfileOk() (*DatabaseSqlNullString, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetProfile(v DatabaseSqlNullString)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPublishedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetPublishedOn() DatabaseSqlNullTime`

GetPublishedOn returns the PublishedOn field if non-nil, zero value otherwise.

### GetPublishedOnOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetPublishedOnOk() (*DatabaseSqlNullTime, bool)`

GetPublishedOnOk returns a tuple with the PublishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetPublishedOn(v DatabaseSqlNullTime)`

SetPublishedOn sets PublishedOn field to given value.

### HasPublishedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasPublishedOn() bool`

HasPublishedOn returns a boolean if a field has been set.

### GetRegions

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetRegions() DatabaseSqlNullString`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetRegionsOk() (*DatabaseSqlNullString, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetRegions(v DatabaseSqlNullString)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRejectedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetRejectedOn() DatabaseSqlNullTime`

GetRejectedOn returns the RejectedOn field if non-nil, zero value otherwise.

### GetRejectedOnOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetRejectedOnOk() (*DatabaseSqlNullTime, bool)`

GetRejectedOnOk returns a tuple with the RejectedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetRejectedOn(v DatabaseSqlNullTime)`

SetRejectedOn sets RejectedOn field to given value.

### HasRejectedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasRejectedOn() bool`

HasRejectedOn returns a boolean if a field has been set.

### GetReviewState

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetReviewState() DatabaseSqlNullString`

GetReviewState returns the ReviewState field if non-nil, zero value otherwise.

### GetReviewStateOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetReviewStateOk() (*DatabaseSqlNullString, bool)`

GetReviewStateOk returns a tuple with the ReviewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewState

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetReviewState(v DatabaseSqlNullString)`

SetReviewState sets ReviewState field to given value.

### HasReviewState

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasReviewState() bool`

HasReviewState returns a boolean if a field has been set.

### GetState

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetState() DatabaseSqlNullString`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetStateOk() (*DatabaseSqlNullString, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetState(v DatabaseSqlNullString)`

SetState sets State field to given value.

### HasState

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubtitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetSubtitle() DatabaseSqlNullString`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetSubtitleOk() (*DatabaseSqlNullString, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetSubtitle(v DatabaseSqlNullString)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetTargetAccounts() DatabaseSqlNullString`

GetTargetAccounts returns the TargetAccounts field if non-nil, zero value otherwise.

### GetTargetAccountsOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetTargetAccountsOk() (*DatabaseSqlNullString, bool)`

GetTargetAccountsOk returns a tuple with the TargetAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetTargetAccounts(v DatabaseSqlNullString)`

SetTargetAccounts sets TargetAccounts field to given value.

### HasTargetAccounts

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasTargetAccounts() bool`

HasTargetAccounts returns a boolean if a field has been set.

### GetTitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetTitle() DatabaseSqlNullString`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetTitleOk() (*DatabaseSqlNullString, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetTitle(v DatabaseSqlNullString)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUniformListingLocator

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetUniformListingLocator() DatabaseSqlNullString`

GetUniformListingLocator returns the UniformListingLocator field if non-nil, zero value otherwise.

### GetUniformListingLocatorOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetUniformListingLocatorOk() (*DatabaseSqlNullString, bool)`

GetUniformListingLocatorOk returns a tuple with the UniformListingLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniformListingLocator

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetUniformListingLocator(v DatabaseSqlNullString)`

SetUniformListingLocator sets UniformListingLocator field to given value.

### HasUniformListingLocator

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasUniformListingLocator() bool`

HasUniformListingLocator returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetUpdatedOn() DatabaseSqlNullTime`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetUpdatedOnOk() (*DatabaseSqlNullTime, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetUpdatedOn(v DatabaseSqlNullTime)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetName

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetName() DatabaseSqlNullString`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetNameOk() (*DatabaseSqlNullString, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetName(v DatabaseSqlNullString)`

SetName sets Name field to given value.

### HasName

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationProfileName

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOrganizationProfileName() DatabaseSqlNullString`

GetOrganizationProfileName returns the OrganizationProfileName field if non-nil, zero value otherwise.

### GetOrganizationProfileNameOk

`func (o *PkgStructsSnowflakeMarketplaceProduct) GetOrganizationProfileNameOk() (*DatabaseSqlNullString, bool)`

GetOrganizationProfileNameOk returns a tuple with the OrganizationProfileName field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileName

`func (o *PkgStructsSnowflakeMarketplaceProduct) SetOrganizationProfileName(v DatabaseSqlNullString)`

SetOrganizationProfileName sets OrganizationProfileName field to given value.

### HasOrganizationProfileName

`func (o *PkgStructsSnowflakeMarketplaceProduct) HasOrganizationProfileName() bool`

HasOrganizationProfileName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


