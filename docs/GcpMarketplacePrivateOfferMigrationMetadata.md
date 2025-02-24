# GcpMarketplacePrivateOfferMigrationMetadata

## Properties

 Name                            | Type                  | Description                                                                                                                                      | Notes      
---------------------------------|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **InventoryFlavorExternalName** | Pointer to **string** | Plan name maybe with term suffix, such as \&quot;plan-name-P1Y\&quot;                                                                            | [optional] 
 **ProductExternalName**         | Pointer to **string** | in format of \&quot;product-service-id.endpoints.gcp-project-id.cloud.goog\&quot;                                                                | [optional] 
 **ProjectId**                   | Pointer to **string** | The GCP project ID of the GCP marketplace integration.                                                                                           | [optional] 
 **ProjectNumber**               | Pointer to **string** | The GCP project number of the provider.                                                                                                          | [optional] 
 **ProviderId**                  | Pointer to **string** | The GCP provider ID / partner ID of the GCP marketplace integration. In most cases, it is the same as the project ID. But it could be different. | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferMigrationMetadata

`func NewGcpMarketplacePrivateOfferMigrationMetadata() *GcpMarketplacePrivateOfferMigrationMetadata`

NewGcpMarketplacePrivateOfferMigrationMetadata instantiates a new GcpMarketplacePrivateOfferMigrationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferMigrationMetadataWithDefaults

`func NewGcpMarketplacePrivateOfferMigrationMetadataWithDefaults() *GcpMarketplacePrivateOfferMigrationMetadata`

NewGcpMarketplacePrivateOfferMigrationMetadataWithDefaults instantiates a new GcpMarketplacePrivateOfferMigrationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryFlavorExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetInventoryFlavorExternalName() string`

GetInventoryFlavorExternalName returns the InventoryFlavorExternalName field if non-nil, zero value otherwise.

### GetInventoryFlavorExternalNameOk

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetInventoryFlavorExternalNameOk() (*string, bool)`

GetInventoryFlavorExternalNameOk returns a tuple with the InventoryFlavorExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryFlavorExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) SetInventoryFlavorExternalName(v string)`

SetInventoryFlavorExternalName sets InventoryFlavorExternalName field to given value.

### HasInventoryFlavorExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) HasInventoryFlavorExternalName() bool`

HasInventoryFlavorExternalName returns a boolean if a field has been set.

### GetProductExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProductExternalName() string`

GetProductExternalName returns the ProductExternalName field if non-nil, zero value otherwise.

### GetProductExternalNameOk

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProductExternalNameOk() (*string, bool)`

GetProductExternalNameOk returns a tuple with the ProductExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) SetProductExternalName(v string)`

SetProductExternalName sets ProductExternalName field to given value.

### HasProductExternalName

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) HasProductExternalName() bool`

HasProductExternalName returns a boolean if a field has been set.

### GetProjectId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectNumber

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProjectNumber() string`

GetProjectNumber returns the ProjectNumber field if non-nil, zero value otherwise.

### GetProjectNumberOk

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProjectNumberOk() (*string, bool)`

GetProjectNumberOk returns a tuple with the ProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectNumber

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) SetProjectNumber(v string)`

SetProjectNumber sets ProjectNumber field to given value.

### HasProjectNumber

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) HasProjectNumber() bool`

HasProjectNumber returns a boolean if a field has been set.

### GetProviderId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GcpMarketplacePrivateOfferMigrationMetadata) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


