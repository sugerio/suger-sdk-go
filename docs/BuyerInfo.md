# BuyerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdyenBuyer** | Pointer to [**AdyenBuyer**](AdyenBuyer.md) | Buyer on Adyen | [optional] 
**AwsBuyer** | Pointer to [**AwsAccountIdentifier**](AwsAccountIdentifier.md) | Buyer from AWS Marketplace | [optional] 
**AzureBuyer** | Pointer to [**AzureADIdentifier**](AzureADIdentifier.md) | Buyer from Azure Marketplace | [optional] 
**CollectableAmount** | Pointer to **float32** | The amount that the seller can collect. It excludes the marketplace commision fee. | [optional] 
**CompanyInfo** | Pointer to [**CompanyInfo**](CompanyInfo.md) |  | [optional] 
**CustomerId** | Pointer to **string** | customerID of buyer on seller&#39;s side | [optional] 
**DisbursedAmount** | Pointer to **float32** | The amount that has been disbursed to the seller account. | [optional] 
**EmailAddress** | Pointer to **string** | The email address of the buyer. This was copied from the new client signup form. | [optional] 
**Fields** | Pointer to **map[string]interface{}** | Fields to store key-value pairs of buyer information. | [optional] 
**GcpBuyer** | Pointer to [**GcpMarketplaceUserAccount**](GcpMarketplaceUserAccount.md) | Buyer from GCP Marketplace | [optional] 
**GrossAmount** | Pointer to **float32** | The gross amount that the buyer has committed to pay, including usage metered amount. | [optional] 
**InvoicedAmount** | Pointer to **float32** | The amount that the buyer has got invoiced. | [optional] 
**LagoCustomerId** | Pointer to **string** | The lgo customer ID for the buyer if it is connected to a lago customer. | [optional] 
**LastModifiedBy** | Pointer to **string** | Last modifier user ID. | [optional] 
**MetronomeCustomerId** | Pointer to **string** | The metronome customer ID for the buyer if it is connected to a metronome customer. | [optional] 
**OrbCustomerId** | Pointer to **string** | The orb customer ID for the buyer if it is connected to a orb customer. | [optional] 
**PaymentConfig** | Pointer to [**PaymentConfig**](PaymentConfig.md) | Payment Config for billing. | [optional] 
**SpaUrl** | Pointer to **string** | Buyer SPA url, public page visited with jwt. | [optional] 
**StripeBuyer** | Pointer to [**StripeCustomer**](StripeCustomer.md) | Buyer as Customer on Stripe | [optional] 

## Methods

### NewBuyerInfo

`func NewBuyerInfo() *BuyerInfo`

NewBuyerInfo instantiates a new BuyerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerInfoWithDefaults

`func NewBuyerInfoWithDefaults() *BuyerInfo`

NewBuyerInfoWithDefaults instantiates a new BuyerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdyenBuyer

`func (o *BuyerInfo) GetAdyenBuyer() AdyenBuyer`

GetAdyenBuyer returns the AdyenBuyer field if non-nil, zero value otherwise.

### GetAdyenBuyerOk

`func (o *BuyerInfo) GetAdyenBuyerOk() (*AdyenBuyer, bool)`

GetAdyenBuyerOk returns a tuple with the AdyenBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdyenBuyer

`func (o *BuyerInfo) SetAdyenBuyer(v AdyenBuyer)`

SetAdyenBuyer sets AdyenBuyer field to given value.

### HasAdyenBuyer

`func (o *BuyerInfo) HasAdyenBuyer() bool`

HasAdyenBuyer returns a boolean if a field has been set.

### GetAwsBuyer

`func (o *BuyerInfo) GetAwsBuyer() AwsAccountIdentifier`

GetAwsBuyer returns the AwsBuyer field if non-nil, zero value otherwise.

### GetAwsBuyerOk

`func (o *BuyerInfo) GetAwsBuyerOk() (*AwsAccountIdentifier, bool)`

GetAwsBuyerOk returns a tuple with the AwsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBuyer

`func (o *BuyerInfo) SetAwsBuyer(v AwsAccountIdentifier)`

SetAwsBuyer sets AwsBuyer field to given value.

### HasAwsBuyer

`func (o *BuyerInfo) HasAwsBuyer() bool`

HasAwsBuyer returns a boolean if a field has been set.

### GetAzureBuyer

`func (o *BuyerInfo) GetAzureBuyer() AzureADIdentifier`

GetAzureBuyer returns the AzureBuyer field if non-nil, zero value otherwise.

### GetAzureBuyerOk

`func (o *BuyerInfo) GetAzureBuyerOk() (*AzureADIdentifier, bool)`

GetAzureBuyerOk returns a tuple with the AzureBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBuyer

`func (o *BuyerInfo) SetAzureBuyer(v AzureADIdentifier)`

SetAzureBuyer sets AzureBuyer field to given value.

### HasAzureBuyer

`func (o *BuyerInfo) HasAzureBuyer() bool`

HasAzureBuyer returns a boolean if a field has been set.

### GetCollectableAmount

`func (o *BuyerInfo) GetCollectableAmount() float32`

GetCollectableAmount returns the CollectableAmount field if non-nil, zero value otherwise.

### GetCollectableAmountOk

`func (o *BuyerInfo) GetCollectableAmountOk() (*float32, bool)`

GetCollectableAmountOk returns a tuple with the CollectableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectableAmount

`func (o *BuyerInfo) SetCollectableAmount(v float32)`

SetCollectableAmount sets CollectableAmount field to given value.

### HasCollectableAmount

`func (o *BuyerInfo) HasCollectableAmount() bool`

HasCollectableAmount returns a boolean if a field has been set.

### GetCompanyInfo

`func (o *BuyerInfo) GetCompanyInfo() CompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *BuyerInfo) GetCompanyInfoOk() (*CompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *BuyerInfo) SetCompanyInfo(v CompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *BuyerInfo) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetCustomerId

`func (o *BuyerInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BuyerInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BuyerInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *BuyerInfo) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisbursedAmount

`func (o *BuyerInfo) GetDisbursedAmount() float32`

GetDisbursedAmount returns the DisbursedAmount field if non-nil, zero value otherwise.

### GetDisbursedAmountOk

`func (o *BuyerInfo) GetDisbursedAmountOk() (*float32, bool)`

GetDisbursedAmountOk returns a tuple with the DisbursedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursedAmount

`func (o *BuyerInfo) SetDisbursedAmount(v float32)`

SetDisbursedAmount sets DisbursedAmount field to given value.

### HasDisbursedAmount

`func (o *BuyerInfo) HasDisbursedAmount() bool`

HasDisbursedAmount returns a boolean if a field has been set.

### GetEmailAddress

`func (o *BuyerInfo) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *BuyerInfo) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *BuyerInfo) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *BuyerInfo) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFields

`func (o *BuyerInfo) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BuyerInfo) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BuyerInfo) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *BuyerInfo) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetGcpBuyer

`func (o *BuyerInfo) GetGcpBuyer() GcpMarketplaceUserAccount`

GetGcpBuyer returns the GcpBuyer field if non-nil, zero value otherwise.

### GetGcpBuyerOk

`func (o *BuyerInfo) GetGcpBuyerOk() (*GcpMarketplaceUserAccount, bool)`

GetGcpBuyerOk returns a tuple with the GcpBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBuyer

`func (o *BuyerInfo) SetGcpBuyer(v GcpMarketplaceUserAccount)`

SetGcpBuyer sets GcpBuyer field to given value.

### HasGcpBuyer

`func (o *BuyerInfo) HasGcpBuyer() bool`

HasGcpBuyer returns a boolean if a field has been set.

### GetGrossAmount

`func (o *BuyerInfo) GetGrossAmount() float32`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *BuyerInfo) GetGrossAmountOk() (*float32, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *BuyerInfo) SetGrossAmount(v float32)`

SetGrossAmount sets GrossAmount field to given value.

### HasGrossAmount

`func (o *BuyerInfo) HasGrossAmount() bool`

HasGrossAmount returns a boolean if a field has been set.

### GetInvoicedAmount

`func (o *BuyerInfo) GetInvoicedAmount() float32`

GetInvoicedAmount returns the InvoicedAmount field if non-nil, zero value otherwise.

### GetInvoicedAmountOk

`func (o *BuyerInfo) GetInvoicedAmountOk() (*float32, bool)`

GetInvoicedAmountOk returns a tuple with the InvoicedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedAmount

`func (o *BuyerInfo) SetInvoicedAmount(v float32)`

SetInvoicedAmount sets InvoicedAmount field to given value.

### HasInvoicedAmount

`func (o *BuyerInfo) HasInvoicedAmount() bool`

HasInvoicedAmount returns a boolean if a field has been set.

### GetLagoCustomerId

`func (o *BuyerInfo) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *BuyerInfo) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *BuyerInfo) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.

### HasLagoCustomerId

`func (o *BuyerInfo) HasLagoCustomerId() bool`

HasLagoCustomerId returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *BuyerInfo) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BuyerInfo) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BuyerInfo) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BuyerInfo) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetMetronomeCustomerId

`func (o *BuyerInfo) GetMetronomeCustomerId() string`

GetMetronomeCustomerId returns the MetronomeCustomerId field if non-nil, zero value otherwise.

### GetMetronomeCustomerIdOk

`func (o *BuyerInfo) GetMetronomeCustomerIdOk() (*string, bool)`

GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeCustomerId

`func (o *BuyerInfo) SetMetronomeCustomerId(v string)`

SetMetronomeCustomerId sets MetronomeCustomerId field to given value.

### HasMetronomeCustomerId

`func (o *BuyerInfo) HasMetronomeCustomerId() bool`

HasMetronomeCustomerId returns a boolean if a field has been set.

### GetOrbCustomerId

`func (o *BuyerInfo) GetOrbCustomerId() string`

GetOrbCustomerId returns the OrbCustomerId field if non-nil, zero value otherwise.

### GetOrbCustomerIdOk

`func (o *BuyerInfo) GetOrbCustomerIdOk() (*string, bool)`

GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbCustomerId

`func (o *BuyerInfo) SetOrbCustomerId(v string)`

SetOrbCustomerId sets OrbCustomerId field to given value.

### HasOrbCustomerId

`func (o *BuyerInfo) HasOrbCustomerId() bool`

HasOrbCustomerId returns a boolean if a field has been set.

### GetPaymentConfig

`func (o *BuyerInfo) GetPaymentConfig() PaymentConfig`

GetPaymentConfig returns the PaymentConfig field if non-nil, zero value otherwise.

### GetPaymentConfigOk

`func (o *BuyerInfo) GetPaymentConfigOk() (*PaymentConfig, bool)`

GetPaymentConfigOk returns a tuple with the PaymentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfig

`func (o *BuyerInfo) SetPaymentConfig(v PaymentConfig)`

SetPaymentConfig sets PaymentConfig field to given value.

### HasPaymentConfig

`func (o *BuyerInfo) HasPaymentConfig() bool`

HasPaymentConfig returns a boolean if a field has been set.

### GetSpaUrl

`func (o *BuyerInfo) GetSpaUrl() string`

GetSpaUrl returns the SpaUrl field if non-nil, zero value otherwise.

### GetSpaUrlOk

`func (o *BuyerInfo) GetSpaUrlOk() (*string, bool)`

GetSpaUrlOk returns a tuple with the SpaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUrl

`func (o *BuyerInfo) SetSpaUrl(v string)`

SetSpaUrl sets SpaUrl field to given value.

### HasSpaUrl

`func (o *BuyerInfo) HasSpaUrl() bool`

HasSpaUrl returns a boolean if a field has been set.

### GetStripeBuyer

`func (o *BuyerInfo) GetStripeBuyer() StripeCustomer`

GetStripeBuyer returns the StripeBuyer field if non-nil, zero value otherwise.

### GetStripeBuyerOk

`func (o *BuyerInfo) GetStripeBuyerOk() (*StripeCustomer, bool)`

GetStripeBuyerOk returns a tuple with the StripeBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeBuyer

`func (o *BuyerInfo) SetStripeBuyer(v StripeCustomer)`

SetStripeBuyer sets StripeBuyer field to given value.

### HasStripeBuyer

`func (o *BuyerInfo) HasStripeBuyer() bool`

HasStripeBuyer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


