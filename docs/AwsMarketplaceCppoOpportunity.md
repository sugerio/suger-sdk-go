# AwsMarketplaceCppoOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to [**[]AwsProductDimension**](AwsProductDimension.md) |  | [optional] 
**ManufacturerAccountId** | Pointer to **string** |  | [optional] 
**ManufacturerLegalName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OfferDetails** | Pointer to [**AwsMarketplaceCppoOpportunityOfferDetails**](AwsMarketplaceCppoOpportunityOfferDetails.md) |  | [optional] 
**PreExistingBuyerAgreement** | Pointer to [**AwsMarketplacePreExistingAgreement**](AwsMarketplacePreExistingAgreement.md) |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]AwsMarketplaceCppoOpportunityRule**](AwsMarketplaceCppoOpportunityRule.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to [**[]AwsMarketplaceCppoOpportunityTerm**](AwsMarketplaceCppoOpportunityTerm.md) |  | [optional] 
**DiscountType** | Pointer to [**AwsMarketplaceCppoDiscountType**](AwsMarketplaceCppoDiscountType.md) | The following fields are not from aws catalog API, only used for cppo_out offer create. They shouldn&#39;t be read in other places because they will absent when fetch opportunity from aws catalog API. | [optional] 
**OpportunityDurationType** | Pointer to [**AwsMarketplaceCppoDurationType**](AwsMarketplaceCppoDurationType.md) |  | [optional] 
**OpportunityId** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsMarketplaceCppoOpportunity

`func NewAwsMarketplaceCppoOpportunity() *AwsMarketplaceCppoOpportunity`

NewAwsMarketplaceCppoOpportunity instantiates a new AwsMarketplaceCppoOpportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCppoOpportunityWithDefaults

`func NewAwsMarketplaceCppoOpportunityWithDefaults() *AwsMarketplaceCppoOpportunity`

NewAwsMarketplaceCppoOpportunityWithDefaults instantiates a new AwsMarketplaceCppoOpportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *AwsMarketplaceCppoOpportunity) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AwsMarketplaceCppoOpportunity) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AwsMarketplaceCppoOpportunity) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AwsMarketplaceCppoOpportunity) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDescription

`func (o *AwsMarketplaceCppoOpportunity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsMarketplaceCppoOpportunity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsMarketplaceCppoOpportunity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsMarketplaceCppoOpportunity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensions

`func (o *AwsMarketplaceCppoOpportunity) GetDimensions() []AwsProductDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AwsMarketplaceCppoOpportunity) GetDimensionsOk() (*[]AwsProductDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AwsMarketplaceCppoOpportunity) SetDimensions(v []AwsProductDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AwsMarketplaceCppoOpportunity) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetManufacturerAccountId

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerAccountId() string`

GetManufacturerAccountId returns the ManufacturerAccountId field if non-nil, zero value otherwise.

### GetManufacturerAccountIdOk

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerAccountIdOk() (*string, bool)`

GetManufacturerAccountIdOk returns a tuple with the ManufacturerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerAccountId

`func (o *AwsMarketplaceCppoOpportunity) SetManufacturerAccountId(v string)`

SetManufacturerAccountId sets ManufacturerAccountId field to given value.

### HasManufacturerAccountId

`func (o *AwsMarketplaceCppoOpportunity) HasManufacturerAccountId() bool`

HasManufacturerAccountId returns a boolean if a field has been set.

### GetManufacturerLegalName

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerLegalName() string`

GetManufacturerLegalName returns the ManufacturerLegalName field if non-nil, zero value otherwise.

### GetManufacturerLegalNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerLegalNameOk() (*string, bool)`

GetManufacturerLegalNameOk returns a tuple with the ManufacturerLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerLegalName

`func (o *AwsMarketplaceCppoOpportunity) SetManufacturerLegalName(v string)`

SetManufacturerLegalName sets ManufacturerLegalName field to given value.

### HasManufacturerLegalName

`func (o *AwsMarketplaceCppoOpportunity) HasManufacturerLegalName() bool`

HasManufacturerLegalName returns a boolean if a field has been set.

### GetName

`func (o *AwsMarketplaceCppoOpportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsMarketplaceCppoOpportunity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsMarketplaceCppoOpportunity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferDetails

`func (o *AwsMarketplaceCppoOpportunity) GetOfferDetails() AwsMarketplaceCppoOpportunityOfferDetails`

GetOfferDetails returns the OfferDetails field if non-nil, zero value otherwise.

### GetOfferDetailsOk

`func (o *AwsMarketplaceCppoOpportunity) GetOfferDetailsOk() (*AwsMarketplaceCppoOpportunityOfferDetails, bool)`

GetOfferDetailsOk returns a tuple with the OfferDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferDetails

`func (o *AwsMarketplaceCppoOpportunity) SetOfferDetails(v AwsMarketplaceCppoOpportunityOfferDetails)`

SetOfferDetails sets OfferDetails field to given value.

### HasOfferDetails

`func (o *AwsMarketplaceCppoOpportunity) HasOfferDetails() bool`

HasOfferDetails returns a boolean if a field has been set.

### GetPreExistingBuyerAgreement

`func (o *AwsMarketplaceCppoOpportunity) GetPreExistingBuyerAgreement() AwsMarketplacePreExistingAgreement`

GetPreExistingBuyerAgreement returns the PreExistingBuyerAgreement field if non-nil, zero value otherwise.

### GetPreExistingBuyerAgreementOk

`func (o *AwsMarketplaceCppoOpportunity) GetPreExistingBuyerAgreementOk() (*AwsMarketplacePreExistingAgreement, bool)`

GetPreExistingBuyerAgreementOk returns a tuple with the PreExistingBuyerAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreExistingBuyerAgreement

`func (o *AwsMarketplaceCppoOpportunity) SetPreExistingBuyerAgreement(v AwsMarketplacePreExistingAgreement)`

SetPreExistingBuyerAgreement sets PreExistingBuyerAgreement field to given value.

### HasPreExistingBuyerAgreement

`func (o *AwsMarketplaceCppoOpportunity) HasPreExistingBuyerAgreement() bool`

HasPreExistingBuyerAgreement returns a boolean if a field has been set.

### GetProductId

`func (o *AwsMarketplaceCppoOpportunity) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AwsMarketplaceCppoOpportunity) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AwsMarketplaceCppoOpportunity) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AwsMarketplaceCppoOpportunity) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *AwsMarketplaceCppoOpportunity) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *AwsMarketplaceCppoOpportunity) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *AwsMarketplaceCppoOpportunity) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetRules

`func (o *AwsMarketplaceCppoOpportunity) GetRules() []AwsMarketplaceCppoOpportunityRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AwsMarketplaceCppoOpportunity) GetRulesOk() (*[]AwsMarketplaceCppoOpportunityRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AwsMarketplaceCppoOpportunity) SetRules(v []AwsMarketplaceCppoOpportunityRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AwsMarketplaceCppoOpportunity) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *AwsMarketplaceCppoOpportunity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsMarketplaceCppoOpportunity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsMarketplaceCppoOpportunity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsMarketplaceCppoOpportunity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTerms

`func (o *AwsMarketplaceCppoOpportunity) GetTerms() []AwsMarketplaceCppoOpportunityTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *AwsMarketplaceCppoOpportunity) GetTermsOk() (*[]AwsMarketplaceCppoOpportunityTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *AwsMarketplaceCppoOpportunity) SetTerms(v []AwsMarketplaceCppoOpportunityTerm)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *AwsMarketplaceCppoOpportunity) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetDiscountType

`func (o *AwsMarketplaceCppoOpportunity) GetDiscountType() AwsMarketplaceCppoDiscountType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *AwsMarketplaceCppoOpportunity) GetDiscountTypeOk() (*AwsMarketplaceCppoDiscountType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *AwsMarketplaceCppoOpportunity) SetDiscountType(v AwsMarketplaceCppoDiscountType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *AwsMarketplaceCppoOpportunity) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetOpportunityDurationType

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityDurationType() AwsMarketplaceCppoDurationType`

GetOpportunityDurationType returns the OpportunityDurationType field if non-nil, zero value otherwise.

### GetOpportunityDurationTypeOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityDurationTypeOk() (*AwsMarketplaceCppoDurationType, bool)`

GetOpportunityDurationTypeOk returns a tuple with the OpportunityDurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityDurationType

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityDurationType(v AwsMarketplaceCppoDurationType)`

SetOpportunityDurationType sets OpportunityDurationType field to given value.

### HasOpportunityDurationType

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityDurationType() bool`

HasOpportunityDurationType returns a boolean if a field has been set.

### GetOpportunityId

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityId() string`

GetOpportunityId returns the OpportunityId field if non-nil, zero value otherwise.

### GetOpportunityIdOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityIdOk() (*string, bool)`

GetOpportunityIdOk returns a tuple with the OpportunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityId

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityId(v string)`

SetOpportunityId sets OpportunityId field to given value.

### HasOpportunityId

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityId() bool`

HasOpportunityId returns a boolean if a field has been set.

### GetPartnerId

`func (o *AwsMarketplaceCppoOpportunity) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *AwsMarketplaceCppoOpportunity) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *AwsMarketplaceCppoOpportunity) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *AwsMarketplaceCppoOpportunity) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


