# AzureProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availabilities** | Pointer to [**[]AzureProductAvailability**](AzureProductAvailability.md) |  | [optional] 
**Branches** | Pointer to [**[]AzureProductBranch**](AzureProductBranch.md) |  | [optional] 
**ExternalIDs** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsModularPublishing** | Pointer to **bool** |  | [optional] 
**Listings** | Pointer to [**[]AzureProductListing**](AzureProductListing.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackageConfigurations** | Pointer to [**[]AzureProductPackageConfiguration**](AzureProductPackageConfiguration.md) |  | [optional] 
**Plans** | Pointer to [**[]AzurePriceAndAvailabilityPrivateOfferPlan**](AzurePriceAndAvailabilityPrivateOfferPlan.md) | All plans under this product | [optional] 
**Properties** | Pointer to [**[]AzureProductProperty**](AzureProductProperty.md) |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Setup** | Pointer to [**AzureProductSetup**](AzureProductSetup.md) |  | [optional] 
**Submissions** | Pointer to [**[]AzureProductSubmission**](AzureProductSubmission.md) |  | [optional] 
**Variants** | Pointer to [**[]AzureProductVariant**](AzureProductVariant.md) |  | [optional] 

## Methods

### NewAzureProduct

`func NewAzureProduct() *AzureProduct`

NewAzureProduct instantiates a new AzureProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductWithDefaults

`func NewAzureProductWithDefaults() *AzureProduct`

NewAzureProductWithDefaults instantiates a new AzureProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilities

`func (o *AzureProduct) GetAvailabilities() []AzureProductAvailability`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *AzureProduct) GetAvailabilitiesOk() (*[]AzureProductAvailability, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *AzureProduct) SetAvailabilities(v []AzureProductAvailability)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *AzureProduct) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.

### GetBranches

`func (o *AzureProduct) GetBranches() []AzureProductBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *AzureProduct) GetBranchesOk() (*[]AzureProductBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *AzureProduct) SetBranches(v []AzureProductBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *AzureProduct) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetExternalIDs

`func (o *AzureProduct) GetExternalIDs() []AzureTypeValue`

GetExternalIDs returns the ExternalIDs field if non-nil, zero value otherwise.

### GetExternalIDsOk

`func (o *AzureProduct) GetExternalIDsOk() (*[]AzureTypeValue, bool)`

GetExternalIDsOk returns a tuple with the ExternalIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIDs

`func (o *AzureProduct) SetExternalIDs(v []AzureTypeValue)`

SetExternalIDs sets ExternalIDs field to given value.

### HasExternalIDs

`func (o *AzureProduct) HasExternalIDs() bool`

HasExternalIDs returns a boolean if a field has been set.

### GetId

`func (o *AzureProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsModularPublishing

`func (o *AzureProduct) GetIsModularPublishing() bool`

GetIsModularPublishing returns the IsModularPublishing field if non-nil, zero value otherwise.

### GetIsModularPublishingOk

`func (o *AzureProduct) GetIsModularPublishingOk() (*bool, bool)`

GetIsModularPublishingOk returns a tuple with the IsModularPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModularPublishing

`func (o *AzureProduct) SetIsModularPublishing(v bool)`

SetIsModularPublishing sets IsModularPublishing field to given value.

### HasIsModularPublishing

`func (o *AzureProduct) HasIsModularPublishing() bool`

HasIsModularPublishing returns a boolean if a field has been set.

### GetListings

`func (o *AzureProduct) GetListings() []AzureProductListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *AzureProduct) GetListingsOk() (*[]AzureProductListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *AzureProduct) SetListings(v []AzureProductListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *AzureProduct) HasListings() bool`

HasListings returns a boolean if a field has been set.

### GetName

`func (o *AzureProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageConfigurations

`func (o *AzureProduct) GetPackageConfigurations() []AzureProductPackageConfiguration`

GetPackageConfigurations returns the PackageConfigurations field if non-nil, zero value otherwise.

### GetPackageConfigurationsOk

`func (o *AzureProduct) GetPackageConfigurationsOk() (*[]AzureProductPackageConfiguration, bool)`

GetPackageConfigurationsOk returns a tuple with the PackageConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageConfigurations

`func (o *AzureProduct) SetPackageConfigurations(v []AzureProductPackageConfiguration)`

SetPackageConfigurations sets PackageConfigurations field to given value.

### HasPackageConfigurations

`func (o *AzureProduct) HasPackageConfigurations() bool`

HasPackageConfigurations returns a boolean if a field has been set.

### GetPlans

`func (o *AzureProduct) GetPlans() []AzurePriceAndAvailabilityPrivateOfferPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AzureProduct) GetPlansOk() (*[]AzurePriceAndAvailabilityPrivateOfferPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AzureProduct) SetPlans(v []AzurePriceAndAvailabilityPrivateOfferPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AzureProduct) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetProperties

`func (o *AzureProduct) GetProperties() []AzureProductProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AzureProduct) GetPropertiesOk() (*[]AzureProductProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AzureProduct) SetProperties(v []AzureProductProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AzureProduct) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProduct) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProduct) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProduct) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProduct) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSetup

`func (o *AzureProduct) GetSetup() AzureProductSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AzureProduct) GetSetupOk() (*AzureProductSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AzureProduct) SetSetup(v AzureProductSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AzureProduct) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSubmissions

`func (o *AzureProduct) GetSubmissions() []AzureProductSubmission`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *AzureProduct) GetSubmissionsOk() (*[]AzureProductSubmission, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *AzureProduct) SetSubmissions(v []AzureProductSubmission)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *AzureProduct) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.

### GetVariants

`func (o *AzureProduct) GetVariants() []AzureProductVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *AzureProduct) GetVariantsOk() (*[]AzureProductVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *AzureProduct) SetVariants(v []AzureProductVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *AzureProduct) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


