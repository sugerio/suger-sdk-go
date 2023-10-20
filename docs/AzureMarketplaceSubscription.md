# AzureMarketplaceSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedCustomerOperations** | Pointer to **[]string** |  | [optional] 
**AutoRenew** | Pointer to **bool** |  | [optional] 
**Beneficiary** | Pointer to [**AzureADIdentifier**](AzureADIdentifier.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**FulfillmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsFreeTrial** | Pointer to **bool** |  | [optional] 
**IsTest** | Pointer to **bool** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OfferId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PublisherId** | Pointer to **string** |  | [optional] 
**Purchaser** | Pointer to [**AzureADIdentifier**](AzureADIdentifier.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**SaasSubscriptionStatus** | Pointer to [**AzureMarketplaceSubscriptionStatus**](AzureMarketplaceSubscriptionStatus.md) |  | [optional] 
**SandboxType** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**SessionMode** | Pointer to **string** |  | [optional] 
**StoreFront** | Pointer to **string** |  | [optional] 
**Term** | Pointer to [**AzureTerm**](AzureTerm.md) |  | [optional] 

## Methods

### NewAzureMarketplaceSubscription

`func NewAzureMarketplaceSubscription() *AzureMarketplaceSubscription`

NewAzureMarketplaceSubscription instantiates a new AzureMarketplaceSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceSubscriptionWithDefaults

`func NewAzureMarketplaceSubscriptionWithDefaults() *AzureMarketplaceSubscription`

NewAzureMarketplaceSubscriptionWithDefaults instantiates a new AzureMarketplaceSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedCustomerOperations

`func (o *AzureMarketplaceSubscription) GetAllowedCustomerOperations() []string`

GetAllowedCustomerOperations returns the AllowedCustomerOperations field if non-nil, zero value otherwise.

### GetAllowedCustomerOperationsOk

`func (o *AzureMarketplaceSubscription) GetAllowedCustomerOperationsOk() (*[]string, bool)`

GetAllowedCustomerOperationsOk returns a tuple with the AllowedCustomerOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCustomerOperations

`func (o *AzureMarketplaceSubscription) SetAllowedCustomerOperations(v []string)`

SetAllowedCustomerOperations sets AllowedCustomerOperations field to given value.

### HasAllowedCustomerOperations

`func (o *AzureMarketplaceSubscription) HasAllowedCustomerOperations() bool`

HasAllowedCustomerOperations returns a boolean if a field has been set.

### GetAutoRenew

`func (o *AzureMarketplaceSubscription) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *AzureMarketplaceSubscription) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *AzureMarketplaceSubscription) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *AzureMarketplaceSubscription) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetBeneficiary

`func (o *AzureMarketplaceSubscription) GetBeneficiary() AzureADIdentifier`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *AzureMarketplaceSubscription) GetBeneficiaryOk() (*AzureADIdentifier, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *AzureMarketplaceSubscription) SetBeneficiary(v AzureADIdentifier)`

SetBeneficiary sets Beneficiary field to given value.

### HasBeneficiary

`func (o *AzureMarketplaceSubscription) HasBeneficiary() bool`

HasBeneficiary returns a boolean if a field has been set.

### GetCreated

`func (o *AzureMarketplaceSubscription) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AzureMarketplaceSubscription) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AzureMarketplaceSubscription) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AzureMarketplaceSubscription) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFulfillmentId

`func (o *AzureMarketplaceSubscription) GetFulfillmentId() string`

GetFulfillmentId returns the FulfillmentId field if non-nil, zero value otherwise.

### GetFulfillmentIdOk

`func (o *AzureMarketplaceSubscription) GetFulfillmentIdOk() (*string, bool)`

GetFulfillmentIdOk returns a tuple with the FulfillmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentId

`func (o *AzureMarketplaceSubscription) SetFulfillmentId(v string)`

SetFulfillmentId sets FulfillmentId field to given value.

### HasFulfillmentId

`func (o *AzureMarketplaceSubscription) HasFulfillmentId() bool`

HasFulfillmentId returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFreeTrial

`func (o *AzureMarketplaceSubscription) GetIsFreeTrial() bool`

GetIsFreeTrial returns the IsFreeTrial field if non-nil, zero value otherwise.

### GetIsFreeTrialOk

`func (o *AzureMarketplaceSubscription) GetIsFreeTrialOk() (*bool, bool)`

GetIsFreeTrialOk returns a tuple with the IsFreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeTrial

`func (o *AzureMarketplaceSubscription) SetIsFreeTrial(v bool)`

SetIsFreeTrial sets IsFreeTrial field to given value.

### HasIsFreeTrial

`func (o *AzureMarketplaceSubscription) HasIsFreeTrial() bool`

HasIsFreeTrial returns a boolean if a field has been set.

### GetIsTest

`func (o *AzureMarketplaceSubscription) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *AzureMarketplaceSubscription) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *AzureMarketplaceSubscription) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *AzureMarketplaceSubscription) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetLastModified

`func (o *AzureMarketplaceSubscription) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AzureMarketplaceSubscription) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AzureMarketplaceSubscription) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AzureMarketplaceSubscription) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *AzureMarketplaceSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureMarketplaceSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureMarketplaceSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureMarketplaceSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferId

`func (o *AzureMarketplaceSubscription) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *AzureMarketplaceSubscription) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *AzureMarketplaceSubscription) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *AzureMarketplaceSubscription) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetPlanId

`func (o *AzureMarketplaceSubscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AzureMarketplaceSubscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AzureMarketplaceSubscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AzureMarketplaceSubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPublisherId

`func (o *AzureMarketplaceSubscription) GetPublisherId() string`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *AzureMarketplaceSubscription) GetPublisherIdOk() (*string, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *AzureMarketplaceSubscription) SetPublisherId(v string)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *AzureMarketplaceSubscription) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetPurchaser

`func (o *AzureMarketplaceSubscription) GetPurchaser() AzureADIdentifier`

GetPurchaser returns the Purchaser field if non-nil, zero value otherwise.

### GetPurchaserOk

`func (o *AzureMarketplaceSubscription) GetPurchaserOk() (*AzureADIdentifier, bool)`

GetPurchaserOk returns a tuple with the Purchaser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaser

`func (o *AzureMarketplaceSubscription) SetPurchaser(v AzureADIdentifier)`

SetPurchaser sets Purchaser field to given value.

### HasPurchaser

`func (o *AzureMarketplaceSubscription) HasPurchaser() bool`

HasPurchaser returns a boolean if a field has been set.

### GetQuantity

`func (o *AzureMarketplaceSubscription) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AzureMarketplaceSubscription) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AzureMarketplaceSubscription) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AzureMarketplaceSubscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSaasSubscriptionStatus

`func (o *AzureMarketplaceSubscription) GetSaasSubscriptionStatus() AzureMarketplaceSubscriptionStatus`

GetSaasSubscriptionStatus returns the SaasSubscriptionStatus field if non-nil, zero value otherwise.

### GetSaasSubscriptionStatusOk

`func (o *AzureMarketplaceSubscription) GetSaasSubscriptionStatusOk() (*AzureMarketplaceSubscriptionStatus, bool)`

GetSaasSubscriptionStatusOk returns a tuple with the SaasSubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasSubscriptionStatus

`func (o *AzureMarketplaceSubscription) SetSaasSubscriptionStatus(v AzureMarketplaceSubscriptionStatus)`

SetSaasSubscriptionStatus sets SaasSubscriptionStatus field to given value.

### HasSaasSubscriptionStatus

`func (o *AzureMarketplaceSubscription) HasSaasSubscriptionStatus() bool`

HasSaasSubscriptionStatus returns a boolean if a field has been set.

### GetSandboxType

`func (o *AzureMarketplaceSubscription) GetSandboxType() string`

GetSandboxType returns the SandboxType field if non-nil, zero value otherwise.

### GetSandboxTypeOk

`func (o *AzureMarketplaceSubscription) GetSandboxTypeOk() (*string, bool)`

GetSandboxTypeOk returns a tuple with the SandboxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxType

`func (o *AzureMarketplaceSubscription) SetSandboxType(v string)`

SetSandboxType sets SandboxType field to given value.

### HasSandboxType

`func (o *AzureMarketplaceSubscription) HasSandboxType() bool`

HasSandboxType returns a boolean if a field has been set.

### GetSessionId

`func (o *AzureMarketplaceSubscription) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AzureMarketplaceSubscription) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AzureMarketplaceSubscription) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AzureMarketplaceSubscription) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSessionMode

`func (o *AzureMarketplaceSubscription) GetSessionMode() string`

GetSessionMode returns the SessionMode field if non-nil, zero value otherwise.

### GetSessionModeOk

`func (o *AzureMarketplaceSubscription) GetSessionModeOk() (*string, bool)`

GetSessionModeOk returns a tuple with the SessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMode

`func (o *AzureMarketplaceSubscription) SetSessionMode(v string)`

SetSessionMode sets SessionMode field to given value.

### HasSessionMode

`func (o *AzureMarketplaceSubscription) HasSessionMode() bool`

HasSessionMode returns a boolean if a field has been set.

### GetStoreFront

`func (o *AzureMarketplaceSubscription) GetStoreFront() string`

GetStoreFront returns the StoreFront field if non-nil, zero value otherwise.

### GetStoreFrontOk

`func (o *AzureMarketplaceSubscription) GetStoreFrontOk() (*string, bool)`

GetStoreFrontOk returns a tuple with the StoreFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFront

`func (o *AzureMarketplaceSubscription) SetStoreFront(v string)`

SetStoreFront sets StoreFront field to given value.

### HasStoreFront

`func (o *AzureMarketplaceSubscription) HasStoreFront() bool`

HasStoreFront returns a boolean if a field has been set.

### GetTerm

`func (o *AzureMarketplaceSubscription) GetTerm() AzureTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *AzureMarketplaceSubscription) GetTermOk() (*AzureTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *AzureMarketplaceSubscription) SetTerm(v AzureTerm)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *AzureMarketplaceSubscription) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


