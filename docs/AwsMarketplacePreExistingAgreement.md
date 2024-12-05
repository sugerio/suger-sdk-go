# AwsMarketplacePreExistingAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquisitionChannel** | Pointer to [**AwsRenewalOfferType**](AwsRenewalOfferType.md) | Indicates if the existing agreement was signed outside AWS Marketplace or within AWS Marketplace. one of values [\&quot;External\&quot;, \&quot;AwsMarketplace\&quot;] | [optional] 
**PricingModel** | Pointer to [**AwsMarketplaceCatalogPricingModel**](AwsMarketplaceCatalogPricingModel.md) | Indicates which pricing model the existing agreement uses. | [optional] 

## Methods

### NewAwsMarketplacePreExistingAgreement

`func NewAwsMarketplacePreExistingAgreement() *AwsMarketplacePreExistingAgreement`

NewAwsMarketplacePreExistingAgreement instantiates a new AwsMarketplacePreExistingAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplacePreExistingAgreementWithDefaults

`func NewAwsMarketplacePreExistingAgreementWithDefaults() *AwsMarketplacePreExistingAgreement`

NewAwsMarketplacePreExistingAgreementWithDefaults instantiates a new AwsMarketplacePreExistingAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquisitionChannel

`func (o *AwsMarketplacePreExistingAgreement) GetAcquisitionChannel() AwsRenewalOfferType`

GetAcquisitionChannel returns the AcquisitionChannel field if non-nil, zero value otherwise.

### GetAcquisitionChannelOk

`func (o *AwsMarketplacePreExistingAgreement) GetAcquisitionChannelOk() (*AwsRenewalOfferType, bool)`

GetAcquisitionChannelOk returns a tuple with the AcquisitionChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquisitionChannel

`func (o *AwsMarketplacePreExistingAgreement) SetAcquisitionChannel(v AwsRenewalOfferType)`

SetAcquisitionChannel sets AcquisitionChannel field to given value.

### HasAcquisitionChannel

`func (o *AwsMarketplacePreExistingAgreement) HasAcquisitionChannel() bool`

HasAcquisitionChannel returns a boolean if a field has been set.

### GetPricingModel

`func (o *AwsMarketplacePreExistingAgreement) GetPricingModel() AwsMarketplaceCatalogPricingModel`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *AwsMarketplacePreExistingAgreement) GetPricingModelOk() (*AwsMarketplaceCatalogPricingModel, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *AwsMarketplacePreExistingAgreement) SetPricingModel(v AwsMarketplaceCatalogPricingModel)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *AwsMarketplacePreExistingAgreement) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


