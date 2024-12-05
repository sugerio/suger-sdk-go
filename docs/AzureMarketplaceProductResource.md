# AzureMarketplaceProductResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLeads** | Pointer to [**AzureMarketplaceCustomerLeads**](AzureMarketplaceCustomerLeads.md) |  | [optional] 
**Listing** | Pointer to [**AzureMarketplaceListing**](AzureMarketplaceListing.md) |  | [optional] 
**ListingAssets** | Pointer to [**[]AzureMarketplaceListingAsset**](AzureMarketplaceListingAsset.md) |  | [optional] 
**Plans** | Pointer to [**[]AzureMarketplacePlanResource**](AzureMarketplacePlanResource.md) |  | [optional] 
**PriceAndAvailabilityCustomMeter** | Pointer to [**AzureMarketplacePriceAndAvailabilityCustomMeter**](AzureMarketplacePriceAndAvailabilityCustomMeter.md) |  | [optional] 
**PriceAndAvailabilityOffer** | Pointer to [**AzureMarketplacePriceAndAvailabilityOffer**](AzureMarketplacePriceAndAvailabilityOffer.md) |  | [optional] 
**Product** | Pointer to [**AzureMarketplaceProduct**](AzureMarketplaceProduct.md) |  | [optional] 
**Property** | Pointer to [**AzureMarketplaceProperty**](AzureMarketplaceProperty.md) |  | [optional] 
**Reseller** | Pointer to [**AzureMarketplaceReseller**](AzureMarketplaceReseller.md) |  | [optional] 
**Setup** | Pointer to [**AzureCommercialMarketplaceSetup**](AzureCommercialMarketplaceSetup.md) |  | [optional] 
**Submission** | Pointer to [**AzureMarketplaceSubmission**](AzureMarketplaceSubmission.md) |  | [optional] 
**TechnicalConfiguration** | Pointer to [**AzureMarketplaceSaasTechnicalConfiguration**](AzureMarketplaceSaasTechnicalConfiguration.md) |  | [optional] 

## Methods

### NewAzureMarketplaceProductResource

`func NewAzureMarketplaceProductResource() *AzureMarketplaceProductResource`

NewAzureMarketplaceProductResource instantiates a new AzureMarketplaceProductResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceProductResourceWithDefaults

`func NewAzureMarketplaceProductResourceWithDefaults() *AzureMarketplaceProductResource`

NewAzureMarketplaceProductResourceWithDefaults instantiates a new AzureMarketplaceProductResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLeads

`func (o *AzureMarketplaceProductResource) GetCustomerLeads() AzureMarketplaceCustomerLeads`

GetCustomerLeads returns the CustomerLeads field if non-nil, zero value otherwise.

### GetCustomerLeadsOk

`func (o *AzureMarketplaceProductResource) GetCustomerLeadsOk() (*AzureMarketplaceCustomerLeads, bool)`

GetCustomerLeadsOk returns a tuple with the CustomerLeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLeads

`func (o *AzureMarketplaceProductResource) SetCustomerLeads(v AzureMarketplaceCustomerLeads)`

SetCustomerLeads sets CustomerLeads field to given value.

### HasCustomerLeads

`func (o *AzureMarketplaceProductResource) HasCustomerLeads() bool`

HasCustomerLeads returns a boolean if a field has been set.

### GetListing

`func (o *AzureMarketplaceProductResource) GetListing() AzureMarketplaceListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *AzureMarketplaceProductResource) GetListingOk() (*AzureMarketplaceListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *AzureMarketplaceProductResource) SetListing(v AzureMarketplaceListing)`

SetListing sets Listing field to given value.

### HasListing

`func (o *AzureMarketplaceProductResource) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetListingAssets

`func (o *AzureMarketplaceProductResource) GetListingAssets() []AzureMarketplaceListingAsset`

GetListingAssets returns the ListingAssets field if non-nil, zero value otherwise.

### GetListingAssetsOk

`func (o *AzureMarketplaceProductResource) GetListingAssetsOk() (*[]AzureMarketplaceListingAsset, bool)`

GetListingAssetsOk returns a tuple with the ListingAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingAssets

`func (o *AzureMarketplaceProductResource) SetListingAssets(v []AzureMarketplaceListingAsset)`

SetListingAssets sets ListingAssets field to given value.

### HasListingAssets

`func (o *AzureMarketplaceProductResource) HasListingAssets() bool`

HasListingAssets returns a boolean if a field has been set.

### GetPlans

`func (o *AzureMarketplaceProductResource) GetPlans() []AzureMarketplacePlanResource`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AzureMarketplaceProductResource) GetPlansOk() (*[]AzureMarketplacePlanResource, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AzureMarketplaceProductResource) SetPlans(v []AzureMarketplacePlanResource)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AzureMarketplaceProductResource) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetPriceAndAvailabilityCustomMeter

`func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityCustomMeter() AzureMarketplacePriceAndAvailabilityCustomMeter`

GetPriceAndAvailabilityCustomMeter returns the PriceAndAvailabilityCustomMeter field if non-nil, zero value otherwise.

### GetPriceAndAvailabilityCustomMeterOk

`func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityCustomMeterOk() (*AzureMarketplacePriceAndAvailabilityCustomMeter, bool)`

GetPriceAndAvailabilityCustomMeterOk returns a tuple with the PriceAndAvailabilityCustomMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAndAvailabilityCustomMeter

`func (o *AzureMarketplaceProductResource) SetPriceAndAvailabilityCustomMeter(v AzureMarketplacePriceAndAvailabilityCustomMeter)`

SetPriceAndAvailabilityCustomMeter sets PriceAndAvailabilityCustomMeter field to given value.

### HasPriceAndAvailabilityCustomMeter

`func (o *AzureMarketplaceProductResource) HasPriceAndAvailabilityCustomMeter() bool`

HasPriceAndAvailabilityCustomMeter returns a boolean if a field has been set.

### GetPriceAndAvailabilityOffer

`func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityOffer() AzureMarketplacePriceAndAvailabilityOffer`

GetPriceAndAvailabilityOffer returns the PriceAndAvailabilityOffer field if non-nil, zero value otherwise.

### GetPriceAndAvailabilityOfferOk

`func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityOfferOk() (*AzureMarketplacePriceAndAvailabilityOffer, bool)`

GetPriceAndAvailabilityOfferOk returns a tuple with the PriceAndAvailabilityOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAndAvailabilityOffer

`func (o *AzureMarketplaceProductResource) SetPriceAndAvailabilityOffer(v AzureMarketplacePriceAndAvailabilityOffer)`

SetPriceAndAvailabilityOffer sets PriceAndAvailabilityOffer field to given value.

### HasPriceAndAvailabilityOffer

`func (o *AzureMarketplaceProductResource) HasPriceAndAvailabilityOffer() bool`

HasPriceAndAvailabilityOffer returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceProductResource) GetProduct() AzureMarketplaceProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceProductResource) GetProductOk() (*AzureMarketplaceProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceProductResource) SetProduct(v AzureMarketplaceProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceProductResource) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperty

`func (o *AzureMarketplaceProductResource) GetProperty() AzureMarketplaceProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AzureMarketplaceProductResource) GetPropertyOk() (*AzureMarketplaceProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AzureMarketplaceProductResource) SetProperty(v AzureMarketplaceProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AzureMarketplaceProductResource) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetReseller

`func (o *AzureMarketplaceProductResource) GetReseller() AzureMarketplaceReseller`

GetReseller returns the Reseller field if non-nil, zero value otherwise.

### GetResellerOk

`func (o *AzureMarketplaceProductResource) GetResellerOk() (*AzureMarketplaceReseller, bool)`

GetResellerOk returns a tuple with the Reseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReseller

`func (o *AzureMarketplaceProductResource) SetReseller(v AzureMarketplaceReseller)`

SetReseller sets Reseller field to given value.

### HasReseller

`func (o *AzureMarketplaceProductResource) HasReseller() bool`

HasReseller returns a boolean if a field has been set.

### GetSetup

`func (o *AzureMarketplaceProductResource) GetSetup() AzureCommercialMarketplaceSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AzureMarketplaceProductResource) GetSetupOk() (*AzureCommercialMarketplaceSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AzureMarketplaceProductResource) SetSetup(v AzureCommercialMarketplaceSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AzureMarketplaceProductResource) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSubmission

`func (o *AzureMarketplaceProductResource) GetSubmission() AzureMarketplaceSubmission`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *AzureMarketplaceProductResource) GetSubmissionOk() (*AzureMarketplaceSubmission, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *AzureMarketplaceProductResource) SetSubmission(v AzureMarketplaceSubmission)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *AzureMarketplaceProductResource) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetTechnicalConfiguration

`func (o *AzureMarketplaceProductResource) GetTechnicalConfiguration() AzureMarketplaceSaasTechnicalConfiguration`

GetTechnicalConfiguration returns the TechnicalConfiguration field if non-nil, zero value otherwise.

### GetTechnicalConfigurationOk

`func (o *AzureMarketplaceProductResource) GetTechnicalConfigurationOk() (*AzureMarketplaceSaasTechnicalConfiguration, bool)`

GetTechnicalConfigurationOk returns a tuple with the TechnicalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalConfiguration

`func (o *AzureMarketplaceProductResource) SetTechnicalConfiguration(v AzureMarketplaceSaasTechnicalConfiguration)`

SetTechnicalConfiguration sets TechnicalConfiguration field to given value.

### HasTechnicalConfiguration

`func (o *AzureMarketplaceProductResource) HasTechnicalConfiguration() bool`

HasTechnicalConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


