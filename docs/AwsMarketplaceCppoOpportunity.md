# AwsMarketplaceCppoOpportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerIds** | Pointer to **[]string** | The AWS Account ID of buyer that are specified by the ISV/Seller in this CPPO Opportunity. | [optional] 
**BuyerNames** | Pointer to **[]string** |  | [optional] 
**ContractDurationInDays** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** | The AWS Account ID of the ISV/Seller who create this CPPO Opportunity. | [optional] 
**CreatedDate** | Pointer to **string** | in format of \&quot;Wed Sep 13 17:50:07 UTC 2023\&quot; | [optional] 
**CustomPriceTerms** | Pointer to [**AwsMarketplaceCppoPriceTerms**](AwsMarketplaceCppoPriceTerms.md) |  | [optional] 
**Discount** | Pointer to **string** | in format of \&quot;10.0\&quot; (10%) | [optional] 
**DiscountPercent** | Pointer to **float32** |  | [optional] 
**DiscountType** | Pointer to [**AwsMarketplaceCppoDiscountType**](AwsMarketplaceCppoDiscountType.md) |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ExpirationDate** | Pointer to **string** | in format of \&quot;Thu Nov 30 23:59:59 UTC 2023\&quot; | [optional] 
**ListingFeeRenewal** | Pointer to **bool** |  | [optional] 
**ManufacturerId** | Pointer to **string** | The AWS Account ID of the ISV/Seller | [optional] 
**ManufacturerName** | Pointer to **string** | The name of the ISV/Seller | [optional] 
**OffersCount** | Pointer to **int32** |  | [optional] 
**OpportunityDiscription** | Pointer to **string** |  | [optional] 
**OpportunityDurationType** | Pointer to [**AwsMarketplaceCppoDurationType**](AwsMarketplaceCppoDurationType.md) |  | [optional] 
**OpportunityEula** | Pointer to [**AwsMarketplaceCppoOpportunityEula**](AwsMarketplaceCppoOpportunityEula.md) |  | [optional] 
**OpportunityId** | Pointer to **string** |  | [optional] 
**OpportunityName** | Pointer to **string** |  | [optional] 
**OpportunityRcmp** | Pointer to [**AwsMarketplaceCppoOpportunityEula**](AwsMarketplaceCppoOpportunityEula.md) |  | [optional] 
**PartnerId** | Pointer to **string** | The AWS Account ID of the Channel Partner | [optional] 
**PartnerName** | Pointer to **string** | The name of the Channel Partner | [optional] 
**PaymentTerms** | Pointer to [**AwsMarketplaceCppoPaymentTerms**](AwsMarketplaceCppoPaymentTerms.md) |  | [optional] 
**ProductId** | Pointer to **string** | The AWS Product ID from the ISV/Seller in this CPPO Opportunity. | [optional] 
**ProductName** | Pointer to **string** | The AWS Product Name from the ISV/Seller in this CPPO Opportunity. | [optional] 
**ProductType** | Pointer to **string** | The type of the AWS Product from the ISV/Seller in this CPPO Opportunity. | [optional] 
**Sppo** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | The status of the CPPO Opportunity. | [optional] 
**UsageAllowed** | Pointer to **int32** | How many times the CPPO Opportunity can be allowed to create CPPO Private Offer by the Channel Partner. | [optional] 

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

### GetBuyerIds

`func (o *AwsMarketplaceCppoOpportunity) GetBuyerIds() []string`

GetBuyerIds returns the BuyerIds field if non-nil, zero value otherwise.

### GetBuyerIdsOk

`func (o *AwsMarketplaceCppoOpportunity) GetBuyerIdsOk() (*[]string, bool)`

GetBuyerIdsOk returns a tuple with the BuyerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerIds

`func (o *AwsMarketplaceCppoOpportunity) SetBuyerIds(v []string)`

SetBuyerIds sets BuyerIds field to given value.

### HasBuyerIds

`func (o *AwsMarketplaceCppoOpportunity) HasBuyerIds() bool`

HasBuyerIds returns a boolean if a field has been set.

### GetBuyerNames

`func (o *AwsMarketplaceCppoOpportunity) GetBuyerNames() []string`

GetBuyerNames returns the BuyerNames field if non-nil, zero value otherwise.

### GetBuyerNamesOk

`func (o *AwsMarketplaceCppoOpportunity) GetBuyerNamesOk() (*[]string, bool)`

GetBuyerNamesOk returns a tuple with the BuyerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerNames

`func (o *AwsMarketplaceCppoOpportunity) SetBuyerNames(v []string)`

SetBuyerNames sets BuyerNames field to given value.

### HasBuyerNames

`func (o *AwsMarketplaceCppoOpportunity) HasBuyerNames() bool`

HasBuyerNames returns a boolean if a field has been set.

### GetContractDurationInDays

`func (o *AwsMarketplaceCppoOpportunity) GetContractDurationInDays() int32`

GetContractDurationInDays returns the ContractDurationInDays field if non-nil, zero value otherwise.

### GetContractDurationInDaysOk

`func (o *AwsMarketplaceCppoOpportunity) GetContractDurationInDaysOk() (*int32, bool)`

GetContractDurationInDaysOk returns a tuple with the ContractDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDurationInDays

`func (o *AwsMarketplaceCppoOpportunity) SetContractDurationInDays(v int32)`

SetContractDurationInDays sets ContractDurationInDays field to given value.

### HasContractDurationInDays

`func (o *AwsMarketplaceCppoOpportunity) HasContractDurationInDays() bool`

HasContractDurationInDays returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AwsMarketplaceCppoOpportunity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AwsMarketplaceCppoOpportunity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AwsMarketplaceCppoOpportunity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AwsMarketplaceCppoOpportunity) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

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

### GetCustomPriceTerms

`func (o *AwsMarketplaceCppoOpportunity) GetCustomPriceTerms() AwsMarketplaceCppoPriceTerms`

GetCustomPriceTerms returns the CustomPriceTerms field if non-nil, zero value otherwise.

### GetCustomPriceTermsOk

`func (o *AwsMarketplaceCppoOpportunity) GetCustomPriceTermsOk() (*AwsMarketplaceCppoPriceTerms, bool)`

GetCustomPriceTermsOk returns a tuple with the CustomPriceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPriceTerms

`func (o *AwsMarketplaceCppoOpportunity) SetCustomPriceTerms(v AwsMarketplaceCppoPriceTerms)`

SetCustomPriceTerms sets CustomPriceTerms field to given value.

### HasCustomPriceTerms

`func (o *AwsMarketplaceCppoOpportunity) HasCustomPriceTerms() bool`

HasCustomPriceTerms returns a boolean if a field has been set.

### GetDiscount

`func (o *AwsMarketplaceCppoOpportunity) GetDiscount() string`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *AwsMarketplaceCppoOpportunity) GetDiscountOk() (*string, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *AwsMarketplaceCppoOpportunity) SetDiscount(v string)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *AwsMarketplaceCppoOpportunity) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountPercent

`func (o *AwsMarketplaceCppoOpportunity) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *AwsMarketplaceCppoOpportunity) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *AwsMarketplaceCppoOpportunity) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *AwsMarketplaceCppoOpportunity) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

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

### GetErrors

`func (o *AwsMarketplaceCppoOpportunity) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AwsMarketplaceCppoOpportunity) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AwsMarketplaceCppoOpportunity) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AwsMarketplaceCppoOpportunity) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AwsMarketplaceCppoOpportunity) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AwsMarketplaceCppoOpportunity) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AwsMarketplaceCppoOpportunity) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AwsMarketplaceCppoOpportunity) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetListingFeeRenewal

`func (o *AwsMarketplaceCppoOpportunity) GetListingFeeRenewal() bool`

GetListingFeeRenewal returns the ListingFeeRenewal field if non-nil, zero value otherwise.

### GetListingFeeRenewalOk

`func (o *AwsMarketplaceCppoOpportunity) GetListingFeeRenewalOk() (*bool, bool)`

GetListingFeeRenewalOk returns a tuple with the ListingFeeRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingFeeRenewal

`func (o *AwsMarketplaceCppoOpportunity) SetListingFeeRenewal(v bool)`

SetListingFeeRenewal sets ListingFeeRenewal field to given value.

### HasListingFeeRenewal

`func (o *AwsMarketplaceCppoOpportunity) HasListingFeeRenewal() bool`

HasListingFeeRenewal returns a boolean if a field has been set.

### GetManufacturerId

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *AwsMarketplaceCppoOpportunity) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *AwsMarketplaceCppoOpportunity) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetManufacturerName

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerName() string`

GetManufacturerName returns the ManufacturerName field if non-nil, zero value otherwise.

### GetManufacturerNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetManufacturerNameOk() (*string, bool)`

GetManufacturerNameOk returns a tuple with the ManufacturerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerName

`func (o *AwsMarketplaceCppoOpportunity) SetManufacturerName(v string)`

SetManufacturerName sets ManufacturerName field to given value.

### HasManufacturerName

`func (o *AwsMarketplaceCppoOpportunity) HasManufacturerName() bool`

HasManufacturerName returns a boolean if a field has been set.

### GetOffersCount

`func (o *AwsMarketplaceCppoOpportunity) GetOffersCount() int32`

GetOffersCount returns the OffersCount field if non-nil, zero value otherwise.

### GetOffersCountOk

`func (o *AwsMarketplaceCppoOpportunity) GetOffersCountOk() (*int32, bool)`

GetOffersCountOk returns a tuple with the OffersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffersCount

`func (o *AwsMarketplaceCppoOpportunity) SetOffersCount(v int32)`

SetOffersCount sets OffersCount field to given value.

### HasOffersCount

`func (o *AwsMarketplaceCppoOpportunity) HasOffersCount() bool`

HasOffersCount returns a boolean if a field has been set.

### GetOpportunityDiscription

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityDiscription() string`

GetOpportunityDiscription returns the OpportunityDiscription field if non-nil, zero value otherwise.

### GetOpportunityDiscriptionOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityDiscriptionOk() (*string, bool)`

GetOpportunityDiscriptionOk returns a tuple with the OpportunityDiscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityDiscription

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityDiscription(v string)`

SetOpportunityDiscription sets OpportunityDiscription field to given value.

### HasOpportunityDiscription

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityDiscription() bool`

HasOpportunityDiscription returns a boolean if a field has been set.

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

### GetOpportunityEula

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityEula() AwsMarketplaceCppoOpportunityEula`

GetOpportunityEula returns the OpportunityEula field if non-nil, zero value otherwise.

### GetOpportunityEulaOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityEulaOk() (*AwsMarketplaceCppoOpportunityEula, bool)`

GetOpportunityEulaOk returns a tuple with the OpportunityEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityEula

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityEula(v AwsMarketplaceCppoOpportunityEula)`

SetOpportunityEula sets OpportunityEula field to given value.

### HasOpportunityEula

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityEula() bool`

HasOpportunityEula returns a boolean if a field has been set.

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

### GetOpportunityName

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityName() string`

GetOpportunityName returns the OpportunityName field if non-nil, zero value otherwise.

### GetOpportunityNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityNameOk() (*string, bool)`

GetOpportunityNameOk returns a tuple with the OpportunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityName

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityName(v string)`

SetOpportunityName sets OpportunityName field to given value.

### HasOpportunityName

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityName() bool`

HasOpportunityName returns a boolean if a field has been set.

### GetOpportunityRcmp

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityRcmp() AwsMarketplaceCppoOpportunityEula`

GetOpportunityRcmp returns the OpportunityRcmp field if non-nil, zero value otherwise.

### GetOpportunityRcmpOk

`func (o *AwsMarketplaceCppoOpportunity) GetOpportunityRcmpOk() (*AwsMarketplaceCppoOpportunityEula, bool)`

GetOpportunityRcmpOk returns a tuple with the OpportunityRcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityRcmp

`func (o *AwsMarketplaceCppoOpportunity) SetOpportunityRcmp(v AwsMarketplaceCppoOpportunityEula)`

SetOpportunityRcmp sets OpportunityRcmp field to given value.

### HasOpportunityRcmp

`func (o *AwsMarketplaceCppoOpportunity) HasOpportunityRcmp() bool`

HasOpportunityRcmp returns a boolean if a field has been set.

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

### GetPartnerName

`func (o *AwsMarketplaceCppoOpportunity) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *AwsMarketplaceCppoOpportunity) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *AwsMarketplaceCppoOpportunity) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *AwsMarketplaceCppoOpportunity) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *AwsMarketplaceCppoOpportunity) GetPaymentTerms() AwsMarketplaceCppoPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AwsMarketplaceCppoOpportunity) GetPaymentTermsOk() (*AwsMarketplaceCppoPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AwsMarketplaceCppoOpportunity) SetPaymentTerms(v AwsMarketplaceCppoPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AwsMarketplaceCppoOpportunity) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

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

### GetProductType

`func (o *AwsMarketplaceCppoOpportunity) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AwsMarketplaceCppoOpportunity) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AwsMarketplaceCppoOpportunity) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AwsMarketplaceCppoOpportunity) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSppo

`func (o *AwsMarketplaceCppoOpportunity) GetSppo() bool`

GetSppo returns the Sppo field if non-nil, zero value otherwise.

### GetSppoOk

`func (o *AwsMarketplaceCppoOpportunity) GetSppoOk() (*bool, bool)`

GetSppoOk returns a tuple with the Sppo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSppo

`func (o *AwsMarketplaceCppoOpportunity) SetSppo(v bool)`

SetSppo sets Sppo field to given value.

### HasSppo

`func (o *AwsMarketplaceCppoOpportunity) HasSppo() bool`

HasSppo returns a boolean if a field has been set.

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

### GetUsageAllowed

`func (o *AwsMarketplaceCppoOpportunity) GetUsageAllowed() int32`

GetUsageAllowed returns the UsageAllowed field if non-nil, zero value otherwise.

### GetUsageAllowedOk

`func (o *AwsMarketplaceCppoOpportunity) GetUsageAllowedOk() (*int32, bool)`

GetUsageAllowedOk returns a tuple with the UsageAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAllowed

`func (o *AwsMarketplaceCppoOpportunity) SetUsageAllowed(v int32)`

SetUsageAllowed sets UsageAllowed field to given value.

### HasUsageAllowed

`func (o *AwsMarketplaceCppoOpportunity) HasUsageAllowed() bool`

HasUsageAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


