# SnowflakeMarketplaceOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessEndTime** | Pointer to **string** |  | [optional] 
**AccessStartDatePreference** | Pointer to **string** |  | [optional] 
**AccessStartTime** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ContractDurationMonths** | Pointer to **int32** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**ContractValue** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to **float32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExpirationTime** | Pointer to **string** |  | [optional] 
**InvoiceStartDatePreference** | Pointer to **string** | invoice start date preference: FIRST_DAY_NEXT_MONTH, OFFER_ACCEPTED_DATE | [optional] 
**InvoiceStartTime** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PaymentTerms** | Pointer to [**SnowflakeMarketplaceOfferPaymentTerms**](SnowflakeMarketplaceOfferPaymentTerms.md) |  | [optional] 
**PricingPlanName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateUpdatedOn** | Pointer to **string** |  | [optional] 
**TargetConsumer** | Pointer to **string** |  | [optional] 
**TermsOfService** | Pointer to [**SnowflakeMarketplaceOfferTermsOfService**](SnowflakeMarketplaceOfferTermsOfService.md) | terms of service: {\&quot;type\&quot;:\&quot;DEFAULT\&quot;} | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewSnowflakeMarketplaceOffer

`func NewSnowflakeMarketplaceOffer() *SnowflakeMarketplaceOffer`

NewSnowflakeMarketplaceOffer instantiates a new SnowflakeMarketplaceOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceOfferWithDefaults

`func NewSnowflakeMarketplaceOfferWithDefaults() *SnowflakeMarketplaceOffer`

NewSnowflakeMarketplaceOfferWithDefaults instantiates a new SnowflakeMarketplaceOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessEndTime

`func (o *SnowflakeMarketplaceOffer) GetAccessEndTime() string`

GetAccessEndTime returns the AccessEndTime field if non-nil, zero value otherwise.

### GetAccessEndTimeOk

`func (o *SnowflakeMarketplaceOffer) GetAccessEndTimeOk() (*string, bool)`

GetAccessEndTimeOk returns a tuple with the AccessEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessEndTime

`func (o *SnowflakeMarketplaceOffer) SetAccessEndTime(v string)`

SetAccessEndTime sets AccessEndTime field to given value.

### HasAccessEndTime

`func (o *SnowflakeMarketplaceOffer) HasAccessEndTime() bool`

HasAccessEndTime returns a boolean if a field has been set.

### GetAccessStartDatePreference

`func (o *SnowflakeMarketplaceOffer) GetAccessStartDatePreference() string`

GetAccessStartDatePreference returns the AccessStartDatePreference field if non-nil, zero value otherwise.

### GetAccessStartDatePreferenceOk

`func (o *SnowflakeMarketplaceOffer) GetAccessStartDatePreferenceOk() (*string, bool)`

GetAccessStartDatePreferenceOk returns a tuple with the AccessStartDatePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStartDatePreference

`func (o *SnowflakeMarketplaceOffer) SetAccessStartDatePreference(v string)`

SetAccessStartDatePreference sets AccessStartDatePreference field to given value.

### HasAccessStartDatePreference

`func (o *SnowflakeMarketplaceOffer) HasAccessStartDatePreference() bool`

HasAccessStartDatePreference returns a boolean if a field has been set.

### GetAccessStartTime

`func (o *SnowflakeMarketplaceOffer) GetAccessStartTime() string`

GetAccessStartTime returns the AccessStartTime field if non-nil, zero value otherwise.

### GetAccessStartTimeOk

`func (o *SnowflakeMarketplaceOffer) GetAccessStartTimeOk() (*string, bool)`

GetAccessStartTimeOk returns a tuple with the AccessStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStartTime

`func (o *SnowflakeMarketplaceOffer) SetAccessStartTime(v string)`

SetAccessStartTime sets AccessStartTime field to given value.

### HasAccessStartTime

`func (o *SnowflakeMarketplaceOffer) HasAccessStartTime() bool`

HasAccessStartTime returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SnowflakeMarketplaceOffer) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SnowflakeMarketplaceOffer) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SnowflakeMarketplaceOffer) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SnowflakeMarketplaceOffer) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetComment

`func (o *SnowflakeMarketplaceOffer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SnowflakeMarketplaceOffer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SnowflakeMarketplaceOffer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SnowflakeMarketplaceOffer) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContractDurationMonths

`func (o *SnowflakeMarketplaceOffer) GetContractDurationMonths() int32`

GetContractDurationMonths returns the ContractDurationMonths field if non-nil, zero value otherwise.

### GetContractDurationMonthsOk

`func (o *SnowflakeMarketplaceOffer) GetContractDurationMonthsOk() (*int32, bool)`

GetContractDurationMonthsOk returns a tuple with the ContractDurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDurationMonths

`func (o *SnowflakeMarketplaceOffer) SetContractDurationMonths(v int32)`

SetContractDurationMonths sets ContractDurationMonths field to given value.

### HasContractDurationMonths

`func (o *SnowflakeMarketplaceOffer) HasContractDurationMonths() bool`

HasContractDurationMonths returns a boolean if a field has been set.

### GetContractType

`func (o *SnowflakeMarketplaceOffer) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *SnowflakeMarketplaceOffer) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *SnowflakeMarketplaceOffer) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *SnowflakeMarketplaceOffer) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetContractValue

`func (o *SnowflakeMarketplaceOffer) GetContractValue() string`

GetContractValue returns the ContractValue field if non-nil, zero value otherwise.

### GetContractValueOk

`func (o *SnowflakeMarketplaceOffer) GetContractValueOk() (*string, bool)`

GetContractValueOk returns a tuple with the ContractValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractValue

`func (o *SnowflakeMarketplaceOffer) SetContractValue(v string)`

SetContractValue sets ContractValue field to given value.

### HasContractValue

`func (o *SnowflakeMarketplaceOffer) HasContractValue() bool`

HasContractValue returns a boolean if a field has been set.

### GetDiscount

`func (o *SnowflakeMarketplaceOffer) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SnowflakeMarketplaceOffer) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SnowflakeMarketplaceOffer) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SnowflakeMarketplaceOffer) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDisplayName

`func (o *SnowflakeMarketplaceOffer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SnowflakeMarketplaceOffer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SnowflakeMarketplaceOffer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SnowflakeMarketplaceOffer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpirationTime

`func (o *SnowflakeMarketplaceOffer) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *SnowflakeMarketplaceOffer) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *SnowflakeMarketplaceOffer) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *SnowflakeMarketplaceOffer) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetInvoiceStartDatePreference

`func (o *SnowflakeMarketplaceOffer) GetInvoiceStartDatePreference() string`

GetInvoiceStartDatePreference returns the InvoiceStartDatePreference field if non-nil, zero value otherwise.

### GetInvoiceStartDatePreferenceOk

`func (o *SnowflakeMarketplaceOffer) GetInvoiceStartDatePreferenceOk() (*string, bool)`

GetInvoiceStartDatePreferenceOk returns a tuple with the InvoiceStartDatePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStartDatePreference

`func (o *SnowflakeMarketplaceOffer) SetInvoiceStartDatePreference(v string)`

SetInvoiceStartDatePreference sets InvoiceStartDatePreference field to given value.

### HasInvoiceStartDatePreference

`func (o *SnowflakeMarketplaceOffer) HasInvoiceStartDatePreference() bool`

HasInvoiceStartDatePreference returns a boolean if a field has been set.

### GetInvoiceStartTime

`func (o *SnowflakeMarketplaceOffer) GetInvoiceStartTime() string`

GetInvoiceStartTime returns the InvoiceStartTime field if non-nil, zero value otherwise.

### GetInvoiceStartTimeOk

`func (o *SnowflakeMarketplaceOffer) GetInvoiceStartTimeOk() (*string, bool)`

GetInvoiceStartTimeOk returns a tuple with the InvoiceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStartTime

`func (o *SnowflakeMarketplaceOffer) SetInvoiceStartTime(v string)`

SetInvoiceStartTime sets InvoiceStartTime field to given value.

### HasInvoiceStartTime

`func (o *SnowflakeMarketplaceOffer) HasInvoiceStartTime() bool`

HasInvoiceStartTime returns a boolean if a field has been set.

### GetIsDefault

`func (o *SnowflakeMarketplaceOffer) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SnowflakeMarketplaceOffer) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SnowflakeMarketplaceOffer) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SnowflakeMarketplaceOffer) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *SnowflakeMarketplaceOffer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnowflakeMarketplaceOffer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnowflakeMarketplaceOffer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnowflakeMarketplaceOffer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *SnowflakeMarketplaceOffer) GetPaymentTerms() SnowflakeMarketplaceOfferPaymentTerms`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *SnowflakeMarketplaceOffer) GetPaymentTermsOk() (*SnowflakeMarketplaceOfferPaymentTerms, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *SnowflakeMarketplaceOffer) SetPaymentTerms(v SnowflakeMarketplaceOfferPaymentTerms)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *SnowflakeMarketplaceOffer) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetPricingPlanName

`func (o *SnowflakeMarketplaceOffer) GetPricingPlanName() string`

GetPricingPlanName returns the PricingPlanName field if non-nil, zero value otherwise.

### GetPricingPlanNameOk

`func (o *SnowflakeMarketplaceOffer) GetPricingPlanNameOk() (*string, bool)`

GetPricingPlanNameOk returns a tuple with the PricingPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPlanName

`func (o *SnowflakeMarketplaceOffer) SetPricingPlanName(v string)`

SetPricingPlanName sets PricingPlanName field to given value.

### HasPricingPlanName

`func (o *SnowflakeMarketplaceOffer) HasPricingPlanName() bool`

HasPricingPlanName returns a boolean if a field has been set.

### GetState

`func (o *SnowflakeMarketplaceOffer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnowflakeMarketplaceOffer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SnowflakeMarketplaceOffer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SnowflakeMarketplaceOffer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateUpdatedOn

`func (o *SnowflakeMarketplaceOffer) GetStateUpdatedOn() string`

GetStateUpdatedOn returns the StateUpdatedOn field if non-nil, zero value otherwise.

### GetStateUpdatedOnOk

`func (o *SnowflakeMarketplaceOffer) GetStateUpdatedOnOk() (*string, bool)`

GetStateUpdatedOnOk returns a tuple with the StateUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdatedOn

`func (o *SnowflakeMarketplaceOffer) SetStateUpdatedOn(v string)`

SetStateUpdatedOn sets StateUpdatedOn field to given value.

### HasStateUpdatedOn

`func (o *SnowflakeMarketplaceOffer) HasStateUpdatedOn() bool`

HasStateUpdatedOn returns a boolean if a field has been set.

### GetTargetConsumer

`func (o *SnowflakeMarketplaceOffer) GetTargetConsumer() string`

GetTargetConsumer returns the TargetConsumer field if non-nil, zero value otherwise.

### GetTargetConsumerOk

`func (o *SnowflakeMarketplaceOffer) GetTargetConsumerOk() (*string, bool)`

GetTargetConsumerOk returns a tuple with the TargetConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConsumer

`func (o *SnowflakeMarketplaceOffer) SetTargetConsumer(v string)`

SetTargetConsumer sets TargetConsumer field to given value.

### HasTargetConsumer

`func (o *SnowflakeMarketplaceOffer) HasTargetConsumer() bool`

HasTargetConsumer returns a boolean if a field has been set.

### GetTermsOfService

`func (o *SnowflakeMarketplaceOffer) GetTermsOfService() SnowflakeMarketplaceOfferTermsOfService`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *SnowflakeMarketplaceOffer) GetTermsOfServiceOk() (*SnowflakeMarketplaceOfferTermsOfService, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *SnowflakeMarketplaceOffer) SetTermsOfService(v SnowflakeMarketplaceOfferTermsOfService)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *SnowflakeMarketplaceOffer) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *SnowflakeMarketplaceOffer) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *SnowflakeMarketplaceOffer) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *SnowflakeMarketplaceOffer) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *SnowflakeMarketplaceOffer) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


