# GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent

## Properties

 Name                           | Type                                                         | Description | Notes      
--------------------------------|--------------------------------------------------------------|-------------|------------
 **Action**                     | Pointer to **string**                                        |             | [optional] 
 **AgreementID**                | Pointer to **string**                                        |             | [optional] 
 **Amount**                     | Pointer to **float32**                                       |             | [optional] 
 **BalanceImpacting**           | Pointer to **int32**                                         |             | [optional] 
 **BankTraceID**                | Pointer to **string**                                        |             | [optional] 
 **BillingAddressID**           | Pointer to **string**                                        |             | [optional] 
 **BrokerID**                   | Pointer to **string**                                        |             | [optional] 
 **BuyerID**                    | Pointer to **string**                                        |             | [optional] 
 **Currency**                   | Pointer to **string**                                        |             | [optional] 
 **DataFeedProductID**          | Pointer to **string**                                        |             | [optional] 
 **DisbursementBillingEventID** | Pointer to **string**                                        |             | [optional] 
 **EndUserAccountID**           | Pointer to **string**                                        |             | [optional] 
 **EntitlementID**              | Pointer to **string**                                        |             | [optional] 
 **FromAccountID**              | Pointer to **string**                                        |             | [optional] 
 **Id**                         | Pointer to **string**                                        |             | [optional] 
 **InsertDate**                 | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md) |             | [optional] 
 **InvoiceDate**                | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md) |             | [optional] 
 **InvoiceID**                  | Pointer to **string**                                        |             | [optional] 
 **OfferID**                    | Pointer to **string**                                        |             | [optional] 
 **OrganizationID**             | Pointer to **string**                                        |             | [optional] 
 **ParentBillingEventID**       | Pointer to **string**                                        |             | [optional] 
 **PaymentDueDate**             | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md) |             | [optional] 
 **ProductID**                  | Pointer to **string**                                        |             | [optional] 
 **ToAccountID**                | Pointer to **string**                                        |             | [optional] 
 **TransactionReferenceID**     | Pointer to **string**                                        |             | [optional] 
 **TransactionType**            | Pointer to **string**                                        |             | [optional] 
 **UsagePeriodEndDate**         | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md) |             | [optional] 
 **UsagePeriodStartDate**       | Pointer to [**DatabaseSqlNullTime**](DatabaseSqlNullTime.md) |             | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent

`func NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent() *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent`

NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent instantiates a new
GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEventWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEventWithDefaults() *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent`

NewGithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEventWithDefaults instantiates a new
GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAgreementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetAgreementID() string`

GetAgreementID returns the AgreementID field if non-nil, zero value otherwise.

### GetAgreementIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetAgreementIDOk() (*string, bool)`

GetAgreementIDOk returns a tuple with the AgreementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetAgreementID(v string)`

SetAgreementID sets AgreementID field to given value.

### HasAgreementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasAgreementID() bool`

HasAgreementID returns a boolean if a field has been set.

### GetAmount

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceImpacting

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBalanceImpacting() int32`

GetBalanceImpacting returns the BalanceImpacting field if non-nil, zero value otherwise.

### GetBalanceImpactingOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBalanceImpactingOk() (*int32, bool)`

GetBalanceImpactingOk returns a tuple with the BalanceImpacting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceImpacting

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetBalanceImpacting(v int32)`

SetBalanceImpacting sets BalanceImpacting field to given value.

### HasBalanceImpacting

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasBalanceImpacting() bool`

HasBalanceImpacting returns a boolean if a field has been set.

### GetBankTraceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBankTraceID() string`

GetBankTraceID returns the BankTraceID field if non-nil, zero value otherwise.

### GetBankTraceIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBankTraceIDOk() (*string, bool)`

GetBankTraceIDOk returns a tuple with the BankTraceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTraceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetBankTraceID(v string)`

SetBankTraceID sets BankTraceID field to given value.

### HasBankTraceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasBankTraceID() bool`

HasBankTraceID returns a boolean if a field has been set.

### GetBillingAddressID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBillingAddressID() string`

GetBillingAddressID returns the BillingAddressID field if non-nil, zero value otherwise.

### GetBillingAddressIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBillingAddressIDOk() (*string, bool)`

GetBillingAddressIDOk returns a tuple with the BillingAddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetBillingAddressID(v string)`

SetBillingAddressID sets BillingAddressID field to given value.

### HasBillingAddressID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasBillingAddressID() bool`

HasBillingAddressID returns a boolean if a field has been set.

### GetBrokerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBrokerID() string`

GetBrokerID returns the BrokerID field if non-nil, zero value otherwise.

### GetBrokerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBrokerIDOk() (*string, bool)`

GetBrokerIDOk returns a tuple with the BrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetBrokerID(v string)`

SetBrokerID sets BrokerID field to given value.

### HasBrokerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasBrokerID() bool`

HasBrokerID returns a boolean if a field has been set.

### GetBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCurrency

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDataFeedProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetDataFeedProductID() string`

GetDataFeedProductID returns the DataFeedProductID field if non-nil, zero value otherwise.

### GetDataFeedProductIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetDataFeedProductIDOk() (*string, bool)`

GetDataFeedProductIDOk returns a tuple with the DataFeedProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFeedProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetDataFeedProductID(v string)`

SetDataFeedProductID sets DataFeedProductID field to given value.

### HasDataFeedProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasDataFeedProductID() bool`

HasDataFeedProductID returns a boolean if a field has been set.

### GetDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetDisbursementBillingEventID() string`

GetDisbursementBillingEventID returns the DisbursementBillingEventID field if non-nil, zero value otherwise.

### GetDisbursementBillingEventIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetDisbursementBillingEventIDOk() (*string, bool)`

GetDisbursementBillingEventIDOk returns a tuple with the DisbursementBillingEventID field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetDisbursementBillingEventID(v string)`

SetDisbursementBillingEventID sets DisbursementBillingEventID field to given value.

### HasDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasDisbursementBillingEventID() bool`

HasDisbursementBillingEventID returns a boolean if a field has been set.

### GetEndUserAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetEndUserAccountID() string`

GetEndUserAccountID returns the EndUserAccountID field if non-nil, zero value otherwise.

### GetEndUserAccountIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetEndUserAccountIDOk() (*string, bool)`

GetEndUserAccountIDOk returns a tuple with the EndUserAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetEndUserAccountID(v string)`

SetEndUserAccountID sets EndUserAccountID field to given value.

### HasEndUserAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasEndUserAccountID() bool`

HasEndUserAccountID returns a boolean if a field has been set.

### GetEntitlementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetFromAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetFromAccountID() string`

GetFromAccountID returns the FromAccountID field if non-nil, zero value otherwise.

### GetFromAccountIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetFromAccountIDOk() (*string, bool)`

GetFromAccountIDOk returns a tuple with the FromAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetFromAccountID(v string)`

SetFromAccountID sets FromAccountID field to given value.

### HasFromAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasFromAccountID() bool`

HasFromAccountID returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInsertDate() DatabaseSqlNullTime`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInsertDateOk() (*DatabaseSqlNullTime, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetInsertDate(v DatabaseSqlNullTime)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInvoiceDate() DatabaseSqlNullTime`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInvoiceDateOk() (*DatabaseSqlNullTime, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetInvoiceDate(v DatabaseSqlNullTime)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetOfferID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetOfferID() string`

GetOfferID returns the OfferID field if non-nil, zero value otherwise.

### GetOfferIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetOfferIDOk() (*string, bool)`

GetOfferIDOk returns a tuple with the OfferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetOfferID(v string)`

SetOfferID sets OfferID field to given value.

### HasOfferID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasOfferID() bool`

HasOfferID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetParentBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetParentBillingEventID() string`

GetParentBillingEventID returns the ParentBillingEventID field if non-nil, zero value otherwise.

### GetParentBillingEventIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetParentBillingEventIDOk() (*string, bool)`

GetParentBillingEventIDOk returns a tuple with the ParentBillingEventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetParentBillingEventID(v string)`

SetParentBillingEventID sets ParentBillingEventID field to given value.

### HasParentBillingEventID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasParentBillingEventID() bool`

HasParentBillingEventID returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetPaymentDueDate() DatabaseSqlNullTime`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetPaymentDueDateOk() (*DatabaseSqlNullTime, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetPaymentDueDate(v DatabaseSqlNullTime)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetToAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetToAccountID() string`

GetToAccountID returns the ToAccountID field if non-nil, zero value otherwise.

### GetToAccountIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetToAccountIDOk() (*string, bool)`

GetToAccountIDOk returns a tuple with the ToAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetToAccountID(v string)`

SetToAccountID sets ToAccountID field to given value.

### HasToAccountID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasToAccountID() bool`

HasToAccountID returns a boolean if a field has been set.

### GetTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetTransactionReferenceID() string`

GetTransactionReferenceID returns the TransactionReferenceID field if non-nil, zero value otherwise.

### GetTransactionReferenceIDOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetTransactionReferenceIDOk() (*string, bool)`

GetTransactionReferenceIDOk returns a tuple with the TransactionReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetTransactionReferenceID(v string)`

SetTransactionReferenceID sets TransactionReferenceID field to given value.

### HasTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasTransactionReferenceID() bool`

HasTransactionReferenceID returns a boolean if a field has been set.

### GetTransactionType

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetUsagePeriodEndDate() DatabaseSqlNullTime`

GetUsagePeriodEndDate returns the UsagePeriodEndDate field if non-nil, zero value otherwise.

### GetUsagePeriodEndDateOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetUsagePeriodEndDateOk() (*DatabaseSqlNullTime, bool)`

GetUsagePeriodEndDateOk returns a tuple with the UsagePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetUsagePeriodEndDate(v DatabaseSqlNullTime)`

SetUsagePeriodEndDate sets UsagePeriodEndDate field to given value.

### HasUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasUsagePeriodEndDate() bool`

HasUsagePeriodEndDate returns a boolean if a field has been set.

### GetUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetUsagePeriodStartDate() DatabaseSqlNullTime`

GetUsagePeriodStartDate returns the UsagePeriodStartDate field if non-nil, zero value otherwise.

### GetUsagePeriodStartDateOk

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) GetUsagePeriodStartDateOk() (*DatabaseSqlNullTime, bool)`

GetUsagePeriodStartDateOk returns a tuple with the UsagePeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) SetUsagePeriodStartDate(v DatabaseSqlNullTime)`

SetUsagePeriodStartDate sets UsagePeriodStartDate field to given value.

### HasUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent) HasUsagePeriodStartDate() bool`

HasUsagePeriodStartDate returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


