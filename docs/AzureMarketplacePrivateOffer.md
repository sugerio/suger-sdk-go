# AzureMarketplacePrivateOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**AcceptBy** | Pointer to **time.Time** | in format YYYY-MM-DD | [optional] 
**AcceptanceLinks** | Pointer to [**[]AzureMarketplacePrivateOfferAcceptanceLink**](AzureMarketplacePrivateOfferAcceptanceLink.md) |  | [optional] 
**Beneficiaries** | Pointer to [**[]AzureMarketplacePrivateOfferBeneficiary**](AzureMarketplacePrivateOfferBeneficiary.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] 
**End** | Pointer to **time.Time** | in format YYYY-MM-DD | [optional] 
**Id** | Pointer to **string** | in format of \&quot;private-offer/private-offer-durable-id\&quot; | [optional] 
**LastModified** | Pointer to **time.Time** | in format YYYY-MM-DD | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotificationContacts** | Pointer to **[]string** | array of email addresses of the users to be notified of any changes in the private offer status. | [optional] 
**OfferPricingType** | Pointer to [**AzureMarketplaceOfferPricingType**](AzureMarketplaceOfferPricingType.md) |  | [optional] 
**Partners** | Pointer to [**[]AzureMarketplacePrivateOfferPartner**](AzureMarketplacePrivateOfferPartner.md) |  | [optional] 
**PreparedBy** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to [**[]AzureMarketplacePrivateOfferPricing**](AzureMarketplacePrivateOfferPricing.md) | Up to 10 pricing entries are allowed. | [optional] 
**PrivateOfferType** | Pointer to [**AzureMarketplacePrivateOfferType**](AzureMarketplacePrivateOfferType.md) |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **time.Time** | in format YYYY-MM-DD, if VariableStartDate &#x3D; true, this field should be empty. | [optional] 
**State** | Pointer to [**AzureMarketplacePrivateOfferState**](AzureMarketplacePrivateOfferState.md) |  | [optional] 
**SubState** | Pointer to [**AzureMarketplacePrivateOfferSubState**](AzureMarketplacePrivateOfferSubState.md) |  | [optional] 
**TermsAndConditionsDocSasUrl** | Pointer to **string** | Only applicable to private offers with privateOfferType &#x3D; customerPromotion || cspPromotion | [optional] 
**TermsAndConditionsDocs** | Pointer to [**[]AzureMarketplacePrivateOfferTermsDoc**](AzureMarketplacePrivateOfferTermsDoc.md) | Only applicable to private offers with privateOfferType &#x3D; multipartyPromotionOriginator || multipartyPromotionChannelPartner | [optional] 
**UpgradedFrom** | Pointer to [**AzureMarketplacePrivateOfferPromotionReference**](AzureMarketplacePrivateOfferPromotionReference.md) |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 
**VariableStartDate** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureMarketplacePrivateOffer

`func NewAzureMarketplacePrivateOffer() *AzureMarketplacePrivateOffer`

NewAzureMarketplacePrivateOffer instantiates a new AzureMarketplacePrivateOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePrivateOfferWithDefaults

`func NewAzureMarketplacePrivateOfferWithDefaults() *AzureMarketplacePrivateOffer`

NewAzureMarketplacePrivateOfferWithDefaults instantiates a new AzureMarketplacePrivateOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplacePrivateOffer) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplacePrivateOffer) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplacePrivateOffer) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplacePrivateOffer) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAcceptBy

`func (o *AzureMarketplacePrivateOffer) GetAcceptBy() time.Time`

GetAcceptBy returns the AcceptBy field if non-nil, zero value otherwise.

### GetAcceptByOk

`func (o *AzureMarketplacePrivateOffer) GetAcceptByOk() (*time.Time, bool)`

GetAcceptByOk returns a tuple with the AcceptBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptBy

`func (o *AzureMarketplacePrivateOffer) SetAcceptBy(v time.Time)`

SetAcceptBy sets AcceptBy field to given value.

### HasAcceptBy

`func (o *AzureMarketplacePrivateOffer) HasAcceptBy() bool`

HasAcceptBy returns a boolean if a field has been set.

### GetAcceptanceLinks

`func (o *AzureMarketplacePrivateOffer) GetAcceptanceLinks() []AzureMarketplacePrivateOfferAcceptanceLink`

GetAcceptanceLinks returns the AcceptanceLinks field if non-nil, zero value otherwise.

### GetAcceptanceLinksOk

`func (o *AzureMarketplacePrivateOffer) GetAcceptanceLinksOk() (*[]AzureMarketplacePrivateOfferAcceptanceLink, bool)`

GetAcceptanceLinksOk returns a tuple with the AcceptanceLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceLinks

`func (o *AzureMarketplacePrivateOffer) SetAcceptanceLinks(v []AzureMarketplacePrivateOfferAcceptanceLink)`

SetAcceptanceLinks sets AcceptanceLinks field to given value.

### HasAcceptanceLinks

`func (o *AzureMarketplacePrivateOffer) HasAcceptanceLinks() bool`

HasAcceptanceLinks returns a boolean if a field has been set.

### GetBeneficiaries

`func (o *AzureMarketplacePrivateOffer) GetBeneficiaries() []AzureMarketplacePrivateOfferBeneficiary`

GetBeneficiaries returns the Beneficiaries field if non-nil, zero value otherwise.

### GetBeneficiariesOk

`func (o *AzureMarketplacePrivateOffer) GetBeneficiariesOk() (*[]AzureMarketplacePrivateOfferBeneficiary, bool)`

GetBeneficiariesOk returns a tuple with the Beneficiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaries

`func (o *AzureMarketplacePrivateOffer) SetBeneficiaries(v []AzureMarketplacePrivateOfferBeneficiary)`

SetBeneficiaries sets Beneficiaries field to given value.

### HasBeneficiaries

`func (o *AzureMarketplacePrivateOffer) HasBeneficiaries() bool`

HasBeneficiaries returns a boolean if a field has been set.

### GetETag

`func (o *AzureMarketplacePrivateOffer) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *AzureMarketplacePrivateOffer) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *AzureMarketplacePrivateOffer) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *AzureMarketplacePrivateOffer) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetEnd

`func (o *AzureMarketplacePrivateOffer) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AzureMarketplacePrivateOffer) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AzureMarketplacePrivateOffer) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AzureMarketplacePrivateOffer) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplacePrivateOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplacePrivateOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplacePrivateOffer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplacePrivateOffer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *AzureMarketplacePrivateOffer) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AzureMarketplacePrivateOffer) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AzureMarketplacePrivateOffer) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AzureMarketplacePrivateOffer) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *AzureMarketplacePrivateOffer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureMarketplacePrivateOffer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureMarketplacePrivateOffer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureMarketplacePrivateOffer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationContacts

`func (o *AzureMarketplacePrivateOffer) GetNotificationContacts() []string`

GetNotificationContacts returns the NotificationContacts field if non-nil, zero value otherwise.

### GetNotificationContactsOk

`func (o *AzureMarketplacePrivateOffer) GetNotificationContactsOk() (*[]string, bool)`

GetNotificationContactsOk returns a tuple with the NotificationContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationContacts

`func (o *AzureMarketplacePrivateOffer) SetNotificationContacts(v []string)`

SetNotificationContacts sets NotificationContacts field to given value.

### HasNotificationContacts

`func (o *AzureMarketplacePrivateOffer) HasNotificationContacts() bool`

HasNotificationContacts returns a boolean if a field has been set.

### GetOfferPricingType

`func (o *AzureMarketplacePrivateOffer) GetOfferPricingType() AzureMarketplaceOfferPricingType`

GetOfferPricingType returns the OfferPricingType field if non-nil, zero value otherwise.

### GetOfferPricingTypeOk

`func (o *AzureMarketplacePrivateOffer) GetOfferPricingTypeOk() (*AzureMarketplaceOfferPricingType, bool)`

GetOfferPricingTypeOk returns a tuple with the OfferPricingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPricingType

`func (o *AzureMarketplacePrivateOffer) SetOfferPricingType(v AzureMarketplaceOfferPricingType)`

SetOfferPricingType sets OfferPricingType field to given value.

### HasOfferPricingType

`func (o *AzureMarketplacePrivateOffer) HasOfferPricingType() bool`

HasOfferPricingType returns a boolean if a field has been set.

### GetPartners

`func (o *AzureMarketplacePrivateOffer) GetPartners() []AzureMarketplacePrivateOfferPartner`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *AzureMarketplacePrivateOffer) GetPartnersOk() (*[]AzureMarketplacePrivateOfferPartner, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *AzureMarketplacePrivateOffer) SetPartners(v []AzureMarketplacePrivateOfferPartner)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *AzureMarketplacePrivateOffer) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetPreparedBy

`func (o *AzureMarketplacePrivateOffer) GetPreparedBy() string`

GetPreparedBy returns the PreparedBy field if non-nil, zero value otherwise.

### GetPreparedByOk

`func (o *AzureMarketplacePrivateOffer) GetPreparedByOk() (*string, bool)`

GetPreparedByOk returns a tuple with the PreparedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedBy

`func (o *AzureMarketplacePrivateOffer) SetPreparedBy(v string)`

SetPreparedBy sets PreparedBy field to given value.

### HasPreparedBy

`func (o *AzureMarketplacePrivateOffer) HasPreparedBy() bool`

HasPreparedBy returns a boolean if a field has been set.

### GetPricing

`func (o *AzureMarketplacePrivateOffer) GetPricing() []AzureMarketplacePrivateOfferPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AzureMarketplacePrivateOffer) GetPricingOk() (*[]AzureMarketplacePrivateOfferPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AzureMarketplacePrivateOffer) SetPricing(v []AzureMarketplacePrivateOfferPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AzureMarketplacePrivateOffer) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetPrivateOfferType

`func (o *AzureMarketplacePrivateOffer) GetPrivateOfferType() AzureMarketplacePrivateOfferType`

GetPrivateOfferType returns the PrivateOfferType field if non-nil, zero value otherwise.

### GetPrivateOfferTypeOk

`func (o *AzureMarketplacePrivateOffer) GetPrivateOfferTypeOk() (*AzureMarketplacePrivateOfferType, bool)`

GetPrivateOfferTypeOk returns a tuple with the PrivateOfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOfferType

`func (o *AzureMarketplacePrivateOffer) SetPrivateOfferType(v AzureMarketplacePrivateOfferType)`

SetPrivateOfferType sets PrivateOfferType field to given value.

### HasPrivateOfferType

`func (o *AzureMarketplacePrivateOffer) HasPrivateOfferType() bool`

HasPrivateOfferType returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplacePrivateOffer) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplacePrivateOffer) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplacePrivateOffer) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplacePrivateOffer) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetStart

`func (o *AzureMarketplacePrivateOffer) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AzureMarketplacePrivateOffer) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AzureMarketplacePrivateOffer) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AzureMarketplacePrivateOffer) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetState

`func (o *AzureMarketplacePrivateOffer) GetState() AzureMarketplacePrivateOfferState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzureMarketplacePrivateOffer) GetStateOk() (*AzureMarketplacePrivateOfferState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzureMarketplacePrivateOffer) SetState(v AzureMarketplacePrivateOfferState)`

SetState sets State field to given value.

### HasState

`func (o *AzureMarketplacePrivateOffer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubState

`func (o *AzureMarketplacePrivateOffer) GetSubState() AzureMarketplacePrivateOfferSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *AzureMarketplacePrivateOffer) GetSubStateOk() (*AzureMarketplacePrivateOfferSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *AzureMarketplacePrivateOffer) SetSubState(v AzureMarketplacePrivateOfferSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *AzureMarketplacePrivateOffer) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetTermsAndConditionsDocSasUrl

`func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocSasUrl() string`

GetTermsAndConditionsDocSasUrl returns the TermsAndConditionsDocSasUrl field if non-nil, zero value otherwise.

### GetTermsAndConditionsDocSasUrlOk

`func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocSasUrlOk() (*string, bool)`

GetTermsAndConditionsDocSasUrlOk returns a tuple with the TermsAndConditionsDocSasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsDocSasUrl

`func (o *AzureMarketplacePrivateOffer) SetTermsAndConditionsDocSasUrl(v string)`

SetTermsAndConditionsDocSasUrl sets TermsAndConditionsDocSasUrl field to given value.

### HasTermsAndConditionsDocSasUrl

`func (o *AzureMarketplacePrivateOffer) HasTermsAndConditionsDocSasUrl() bool`

HasTermsAndConditionsDocSasUrl returns a boolean if a field has been set.

### GetTermsAndConditionsDocs

`func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocs() []AzureMarketplacePrivateOfferTermsDoc`

GetTermsAndConditionsDocs returns the TermsAndConditionsDocs field if non-nil, zero value otherwise.

### GetTermsAndConditionsDocsOk

`func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocsOk() (*[]AzureMarketplacePrivateOfferTermsDoc, bool)`

GetTermsAndConditionsDocsOk returns a tuple with the TermsAndConditionsDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionsDocs

`func (o *AzureMarketplacePrivateOffer) SetTermsAndConditionsDocs(v []AzureMarketplacePrivateOfferTermsDoc)`

SetTermsAndConditionsDocs sets TermsAndConditionsDocs field to given value.

### HasTermsAndConditionsDocs

`func (o *AzureMarketplacePrivateOffer) HasTermsAndConditionsDocs() bool`

HasTermsAndConditionsDocs returns a boolean if a field has been set.

### GetUpgradedFrom

`func (o *AzureMarketplacePrivateOffer) GetUpgradedFrom() AzureMarketplacePrivateOfferPromotionReference`

GetUpgradedFrom returns the UpgradedFrom field if non-nil, zero value otherwise.

### GetUpgradedFromOk

`func (o *AzureMarketplacePrivateOffer) GetUpgradedFromOk() (*AzureMarketplacePrivateOfferPromotionReference, bool)`

GetUpgradedFromOk returns a tuple with the UpgradedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedFrom

`func (o *AzureMarketplacePrivateOffer) SetUpgradedFrom(v AzureMarketplacePrivateOfferPromotionReference)`

SetUpgradedFrom sets UpgradedFrom field to given value.

### HasUpgradedFrom

`func (o *AzureMarketplacePrivateOffer) HasUpgradedFrom() bool`

HasUpgradedFrom returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplacePrivateOffer) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplacePrivateOffer) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplacePrivateOffer) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplacePrivateOffer) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetVariableStartDate

`func (o *AzureMarketplacePrivateOffer) GetVariableStartDate() bool`

GetVariableStartDate returns the VariableStartDate field if non-nil, zero value otherwise.

### GetVariableStartDateOk

`func (o *AzureMarketplacePrivateOffer) GetVariableStartDateOk() (*bool, bool)`

GetVariableStartDateOk returns a tuple with the VariableStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableStartDate

`func (o *AzureMarketplacePrivateOffer) SetVariableStartDate(v bool)`

SetVariableStartDate sets VariableStartDate field to given value.

### HasVariableStartDate

`func (o *AzureMarketplacePrivateOffer) HasVariableStartDate() bool`

HasVariableStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


