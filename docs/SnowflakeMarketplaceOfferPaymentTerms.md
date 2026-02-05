# SnowflakeMarketplaceOfferPaymentTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** | Allowed payment methods. Accepted values are CREDIT_CARD, INSTALLMENT. | [optional] 
**InstallmentSchedule** | Pointer to [**SnowflakeMarketplacePlanInstallmentSchedule**](SnowflakeMarketplacePlanInstallmentSchedule.md) | The installment schedule for the offer. | [optional] 
**PaymentType** | Pointer to **string** | The pricing plan payment types. Accepted values are INVOICE, CREDIT_CARD, INSTALLMENT. | [optional] 

## Methods

### NewSnowflakeMarketplaceOfferPaymentTerms

`func NewSnowflakeMarketplaceOfferPaymentTerms() *SnowflakeMarketplaceOfferPaymentTerms`

NewSnowflakeMarketplaceOfferPaymentTerms instantiates a new SnowflakeMarketplaceOfferPaymentTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceOfferPaymentTermsWithDefaults

`func NewSnowflakeMarketplaceOfferPaymentTermsWithDefaults() *SnowflakeMarketplaceOfferPaymentTerms`

NewSnowflakeMarketplaceOfferPaymentTermsWithDefaults instantiates a new SnowflakeMarketplaceOfferPaymentTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *SnowflakeMarketplaceOfferPaymentTerms) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *SnowflakeMarketplaceOfferPaymentTerms) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetInstallmentSchedule

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetInstallmentSchedule() SnowflakeMarketplacePlanInstallmentSchedule`

GetInstallmentSchedule returns the InstallmentSchedule field if non-nil, zero value otherwise.

### GetInstallmentScheduleOk

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetInstallmentScheduleOk() (*SnowflakeMarketplacePlanInstallmentSchedule, bool)`

GetInstallmentScheduleOk returns a tuple with the InstallmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentSchedule

`func (o *SnowflakeMarketplaceOfferPaymentTerms) SetInstallmentSchedule(v SnowflakeMarketplacePlanInstallmentSchedule)`

SetInstallmentSchedule sets InstallmentSchedule field to given value.

### HasInstallmentSchedule

`func (o *SnowflakeMarketplaceOfferPaymentTerms) HasInstallmentSchedule() bool`

HasInstallmentSchedule returns a boolean if a field has been set.

### GetPaymentType

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *SnowflakeMarketplaceOfferPaymentTerms) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *SnowflakeMarketplaceOfferPaymentTerms) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *SnowflakeMarketplaceOfferPaymentTerms) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


