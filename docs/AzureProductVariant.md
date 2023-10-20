# AzureProductVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureGovernmentCertifications** | Pointer to [**[]AzureGovernmentCertification**](AzureGovernmentCertification.md) |  | [optional] 
**CloudAvailabilities** | Pointer to **[]string** |  | [optional] 
**ConversionPaths** | Pointer to **string** |  | [optional] 
**ExtendedProperties** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**FeatureAvailabilities** | Pointer to [**[]AzureProductFeatureAvailability**](AzureProductFeatureAvailability.md) | Not original fields. They are populated by other API calls | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LeadGenID** | Pointer to **string** |  | [optional] 
**ReferenceVariantID** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureProductVariant

`func NewAzureProductVariant() *AzureProductVariant`

NewAzureProductVariant instantiates a new AzureProductVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductVariantWithDefaults

`func NewAzureProductVariantWithDefaults() *AzureProductVariant`

NewAzureProductVariantWithDefaults instantiates a new AzureProductVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureGovernmentCertifications

`func (o *AzureProductVariant) GetAzureGovernmentCertifications() []AzureGovernmentCertification`

GetAzureGovernmentCertifications returns the AzureGovernmentCertifications field if non-nil, zero value otherwise.

### GetAzureGovernmentCertificationsOk

`func (o *AzureProductVariant) GetAzureGovernmentCertificationsOk() (*[]AzureGovernmentCertification, bool)`

GetAzureGovernmentCertificationsOk returns a tuple with the AzureGovernmentCertifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureGovernmentCertifications

`func (o *AzureProductVariant) SetAzureGovernmentCertifications(v []AzureGovernmentCertification)`

SetAzureGovernmentCertifications sets AzureGovernmentCertifications field to given value.

### HasAzureGovernmentCertifications

`func (o *AzureProductVariant) HasAzureGovernmentCertifications() bool`

HasAzureGovernmentCertifications returns a boolean if a field has been set.

### GetCloudAvailabilities

`func (o *AzureProductVariant) GetCloudAvailabilities() []string`

GetCloudAvailabilities returns the CloudAvailabilities field if non-nil, zero value otherwise.

### GetCloudAvailabilitiesOk

`func (o *AzureProductVariant) GetCloudAvailabilitiesOk() (*[]string, bool)`

GetCloudAvailabilitiesOk returns a tuple with the CloudAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAvailabilities

`func (o *AzureProductVariant) SetCloudAvailabilities(v []string)`

SetCloudAvailabilities sets CloudAvailabilities field to given value.

### HasCloudAvailabilities

`func (o *AzureProductVariant) HasCloudAvailabilities() bool`

HasCloudAvailabilities returns a boolean if a field has been set.

### GetConversionPaths

`func (o *AzureProductVariant) GetConversionPaths() string`

GetConversionPaths returns the ConversionPaths field if non-nil, zero value otherwise.

### GetConversionPathsOk

`func (o *AzureProductVariant) GetConversionPathsOk() (*string, bool)`

GetConversionPathsOk returns a tuple with the ConversionPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionPaths

`func (o *AzureProductVariant) SetConversionPaths(v string)`

SetConversionPaths sets ConversionPaths field to given value.

### HasConversionPaths

`func (o *AzureProductVariant) HasConversionPaths() bool`

HasConversionPaths returns a boolean if a field has been set.

### GetExtendedProperties

`func (o *AzureProductVariant) GetExtendedProperties() []AzureTypeValue`

GetExtendedProperties returns the ExtendedProperties field if non-nil, zero value otherwise.

### GetExtendedPropertiesOk

`func (o *AzureProductVariant) GetExtendedPropertiesOk() (*[]AzureTypeValue, bool)`

GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProperties

`func (o *AzureProductVariant) SetExtendedProperties(v []AzureTypeValue)`

SetExtendedProperties sets ExtendedProperties field to given value.

### HasExtendedProperties

`func (o *AzureProductVariant) HasExtendedProperties() bool`

HasExtendedProperties returns a boolean if a field has been set.

### GetExternalID

`func (o *AzureProductVariant) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *AzureProductVariant) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *AzureProductVariant) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *AzureProductVariant) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetFeatureAvailabilities

`func (o *AzureProductVariant) GetFeatureAvailabilities() []AzureProductFeatureAvailability`

GetFeatureAvailabilities returns the FeatureAvailabilities field if non-nil, zero value otherwise.

### GetFeatureAvailabilitiesOk

`func (o *AzureProductVariant) GetFeatureAvailabilitiesOk() (*[]AzureProductFeatureAvailability, bool)`

GetFeatureAvailabilitiesOk returns a tuple with the FeatureAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureAvailabilities

`func (o *AzureProductVariant) SetFeatureAvailabilities(v []AzureProductFeatureAvailability)`

SetFeatureAvailabilities sets FeatureAvailabilities field to given value.

### HasFeatureAvailabilities

`func (o *AzureProductVariant) HasFeatureAvailabilities() bool`

HasFeatureAvailabilities returns a boolean if a field has been set.

### GetFriendlyName

`func (o *AzureProductVariant) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AzureProductVariant) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AzureProductVariant) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AzureProductVariant) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *AzureProductVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProductVariant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProductVariant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProductVariant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeadGenID

`func (o *AzureProductVariant) GetLeadGenID() string`

GetLeadGenID returns the LeadGenID field if non-nil, zero value otherwise.

### GetLeadGenIDOk

`func (o *AzureProductVariant) GetLeadGenIDOk() (*string, bool)`

GetLeadGenIDOk returns a tuple with the LeadGenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadGenID

`func (o *AzureProductVariant) SetLeadGenID(v string)`

SetLeadGenID sets LeadGenID field to given value.

### HasLeadGenID

`func (o *AzureProductVariant) HasLeadGenID() bool`

HasLeadGenID returns a boolean if a field has been set.

### GetReferenceVariantID

`func (o *AzureProductVariant) GetReferenceVariantID() string`

GetReferenceVariantID returns the ReferenceVariantID field if non-nil, zero value otherwise.

### GetReferenceVariantIDOk

`func (o *AzureProductVariant) GetReferenceVariantIDOk() (*string, bool)`

GetReferenceVariantIDOk returns a tuple with the ReferenceVariantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceVariantID

`func (o *AzureProductVariant) SetReferenceVariantID(v string)`

SetReferenceVariantID sets ReferenceVariantID field to given value.

### HasReferenceVariantID

`func (o *AzureProductVariant) HasReferenceVariantID() bool`

HasReferenceVariantID returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProductVariant) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProductVariant) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProductVariant) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProductVariant) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetState

`func (o *AzureProductVariant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzureProductVariant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzureProductVariant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AzureProductVariant) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


