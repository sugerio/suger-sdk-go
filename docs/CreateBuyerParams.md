# CreateBuyerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdyenCustomerId** | Pointer to **string** | Adyen customerId of this buyer. If not provided but Partner is ADYEN, will create a new customer on Adyen. | [optional] 
**CompanyInfo** | Pointer to [**CompanyInfo**](CompanyInfo.md) | Optional. CompanyInfo of the buyer. | [optional] 
**CustomerId** | Pointer to **string** | The customer ID to recognize the cloud marketplace buyer in your internal system. | [optional] 
**Description** | Pointer to **string** | The description of the buyer. | [optional] 
**LagoCustomerId** | Pointer to **string** | Optional. The Lago Customer ID of the buyer. | [optional] 
**MetronomeCustomerId** | Pointer to **string** | Optional. The Metronome Customer ID of the buyer. | [optional] 
**Name** | Pointer to **string** | The name of the buyer. | [optional] 
**OrbCustomerId** | Pointer to **string** | Optional. The Orb Customer ID of the buyer. | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) | The channel partner where this buyer is billed. Only STRIPE &amp; ADYEN are supported at the moment. | [optional] 
**PaymentConfig** | Pointer to [**PaymentConfig**](PaymentConfig.md) | Payment config for billing. | [optional] 
**StripeCustomerId** | Pointer to **string** | Stripe customerId of this buyer. If not provided but Partner is STRIPE, will create a new customer on stripe. | [optional] 

## Methods

### NewCreateBuyerParams

`func NewCreateBuyerParams() *CreateBuyerParams`

NewCreateBuyerParams instantiates a new CreateBuyerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerParamsWithDefaults

`func NewCreateBuyerParamsWithDefaults() *CreateBuyerParams`

NewCreateBuyerParamsWithDefaults instantiates a new CreateBuyerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdyenCustomerId

`func (o *CreateBuyerParams) GetAdyenCustomerId() string`

GetAdyenCustomerId returns the AdyenCustomerId field if non-nil, zero value otherwise.

### GetAdyenCustomerIdOk

`func (o *CreateBuyerParams) GetAdyenCustomerIdOk() (*string, bool)`

GetAdyenCustomerIdOk returns a tuple with the AdyenCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdyenCustomerId

`func (o *CreateBuyerParams) SetAdyenCustomerId(v string)`

SetAdyenCustomerId sets AdyenCustomerId field to given value.

### HasAdyenCustomerId

`func (o *CreateBuyerParams) HasAdyenCustomerId() bool`

HasAdyenCustomerId returns a boolean if a field has been set.

### GetCompanyInfo

`func (o *CreateBuyerParams) GetCompanyInfo() CompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *CreateBuyerParams) GetCompanyInfoOk() (*CompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *CreateBuyerParams) SetCompanyInfo(v CompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *CreateBuyerParams) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetCustomerId

`func (o *CreateBuyerParams) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateBuyerParams) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateBuyerParams) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreateBuyerParams) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateBuyerParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBuyerParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBuyerParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBuyerParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLagoCustomerId

`func (o *CreateBuyerParams) GetLagoCustomerId() string`

GetLagoCustomerId returns the LagoCustomerId field if non-nil, zero value otherwise.

### GetLagoCustomerIdOk

`func (o *CreateBuyerParams) GetLagoCustomerIdOk() (*string, bool)`

GetLagoCustomerIdOk returns a tuple with the LagoCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoCustomerId

`func (o *CreateBuyerParams) SetLagoCustomerId(v string)`

SetLagoCustomerId sets LagoCustomerId field to given value.

### HasLagoCustomerId

`func (o *CreateBuyerParams) HasLagoCustomerId() bool`

HasLagoCustomerId returns a boolean if a field has been set.

### GetMetronomeCustomerId

`func (o *CreateBuyerParams) GetMetronomeCustomerId() string`

GetMetronomeCustomerId returns the MetronomeCustomerId field if non-nil, zero value otherwise.

### GetMetronomeCustomerIdOk

`func (o *CreateBuyerParams) GetMetronomeCustomerIdOk() (*string, bool)`

GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeCustomerId

`func (o *CreateBuyerParams) SetMetronomeCustomerId(v string)`

SetMetronomeCustomerId sets MetronomeCustomerId field to given value.

### HasMetronomeCustomerId

`func (o *CreateBuyerParams) HasMetronomeCustomerId() bool`

HasMetronomeCustomerId returns a boolean if a field has been set.

### GetName

`func (o *CreateBuyerParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBuyerParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBuyerParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBuyerParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrbCustomerId

`func (o *CreateBuyerParams) GetOrbCustomerId() string`

GetOrbCustomerId returns the OrbCustomerId field if non-nil, zero value otherwise.

### GetOrbCustomerIdOk

`func (o *CreateBuyerParams) GetOrbCustomerIdOk() (*string, bool)`

GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbCustomerId

`func (o *CreateBuyerParams) SetOrbCustomerId(v string)`

SetOrbCustomerId sets OrbCustomerId field to given value.

### HasOrbCustomerId

`func (o *CreateBuyerParams) HasOrbCustomerId() bool`

HasOrbCustomerId returns a boolean if a field has been set.

### GetPartner

`func (o *CreateBuyerParams) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *CreateBuyerParams) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *CreateBuyerParams) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *CreateBuyerParams) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPaymentConfig

`func (o *CreateBuyerParams) GetPaymentConfig() PaymentConfig`

GetPaymentConfig returns the PaymentConfig field if non-nil, zero value otherwise.

### GetPaymentConfigOk

`func (o *CreateBuyerParams) GetPaymentConfigOk() (*PaymentConfig, bool)`

GetPaymentConfigOk returns a tuple with the PaymentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfig

`func (o *CreateBuyerParams) SetPaymentConfig(v PaymentConfig)`

SetPaymentConfig sets PaymentConfig field to given value.

### HasPaymentConfig

`func (o *CreateBuyerParams) HasPaymentConfig() bool`

HasPaymentConfig returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *CreateBuyerParams) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateBuyerParams) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateBuyerParams) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *CreateBuyerParams) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


