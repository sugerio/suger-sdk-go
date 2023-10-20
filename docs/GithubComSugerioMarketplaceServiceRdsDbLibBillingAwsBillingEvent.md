# GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**AgreementID** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**BalanceImpacting** | Pointer to **int32** |  | [optional] 
**BankTraceID** | Pointer to **string** |  | [optional] 
**BillingAddressID** | Pointer to **string** |  | [optional] 
**BrokerID** | Pointer to **string** |  | [optional] 
**BuyerID** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DataFeedProductID** | Pointer to **string** |  | [optional] 
**DisbursementBillingEventID** | Pointer to **string** |  | [optional] 
**EndUserAccountID** | Pointer to **string** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**FromAccountID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InsertDate** | Pointer to [**SqlNullTime**](SqlNullTime.md) |  | [optional] 
**InvoiceDate** | Pointer to [**SqlNullTime**](SqlNullTime.md) |  | [optional] 
**InvoiceID** | Pointer to **string** |  | [optional] 
**OfferID** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**ParentBillingEventID** | Pointer to **string** |  | [optional] 
**PaymentDueDate** | Pointer to [**SqlNullTime**](SqlNullTime.md) |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 
**ToAccountID** | Pointer to **string** |  | [optional] 
**TransactionReferenceID** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**UsagePeriodEndDate** | Pointer to [**SqlNullTime**](SqlNullTime.md) |  | [optional] 
**UsagePeriodStartDate** | Pointer to [**SqlNullTime**](SqlNullTime.md) |  | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent

`func NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent() *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent`

NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent instantiates a new GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEventWithDefaults

`func NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEventWithDefaults() *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent`

NewGithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEventWithDefaults instantiates a new GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAgreementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetAgreementID() string`

GetAgreementID returns the AgreementID field if non-nil, zero value otherwise.

### GetAgreementIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetAgreementIDOk() (*string, bool)`

GetAgreementIDOk returns a tuple with the AgreementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetAgreementID(v string)`

SetAgreementID sets AgreementID field to given value.

### HasAgreementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasAgreementID() bool`

HasAgreementID returns a boolean if a field has been set.

### GetAmount

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceImpacting

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBalanceImpacting() int32`

GetBalanceImpacting returns the BalanceImpacting field if non-nil, zero value otherwise.

### GetBalanceImpactingOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBalanceImpactingOk() (*int32, bool)`

GetBalanceImpactingOk returns a tuple with the BalanceImpacting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceImpacting

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetBalanceImpacting(v int32)`

SetBalanceImpacting sets BalanceImpacting field to given value.

### HasBalanceImpacting

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasBalanceImpacting() bool`

HasBalanceImpacting returns a boolean if a field has been set.

### GetBankTraceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBankTraceID() string`

GetBankTraceID returns the BankTraceID field if non-nil, zero value otherwise.

### GetBankTraceIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBankTraceIDOk() (*string, bool)`

GetBankTraceIDOk returns a tuple with the BankTraceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTraceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetBankTraceID(v string)`

SetBankTraceID sets BankTraceID field to given value.

### HasBankTraceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasBankTraceID() bool`

HasBankTraceID returns a boolean if a field has been set.

### GetBillingAddressID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBillingAddressID() string`

GetBillingAddressID returns the BillingAddressID field if non-nil, zero value otherwise.

### GetBillingAddressIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBillingAddressIDOk() (*string, bool)`

GetBillingAddressIDOk returns a tuple with the BillingAddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetBillingAddressID(v string)`

SetBillingAddressID sets BillingAddressID field to given value.

### HasBillingAddressID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasBillingAddressID() bool`

HasBillingAddressID returns a boolean if a field has been set.

### GetBrokerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBrokerID() string`

GetBrokerID returns the BrokerID field if non-nil, zero value otherwise.

### GetBrokerIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBrokerIDOk() (*string, bool)`

GetBrokerIDOk returns a tuple with the BrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetBrokerID(v string)`

SetBrokerID sets BrokerID field to given value.

### HasBrokerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasBrokerID() bool`

HasBrokerID returns a boolean if a field has been set.

### GetBuyerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCurrency

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDataFeedProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetDataFeedProductID() string`

GetDataFeedProductID returns the DataFeedProductID field if non-nil, zero value otherwise.

### GetDataFeedProductIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetDataFeedProductIDOk() (*string, bool)`

GetDataFeedProductIDOk returns a tuple with the DataFeedProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFeedProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetDataFeedProductID(v string)`

SetDataFeedProductID sets DataFeedProductID field to given value.

### HasDataFeedProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasDataFeedProductID() bool`

HasDataFeedProductID returns a boolean if a field has been set.

### GetDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetDisbursementBillingEventID() string`

GetDisbursementBillingEventID returns the DisbursementBillingEventID field if non-nil, zero value otherwise.

### GetDisbursementBillingEventIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetDisbursementBillingEventIDOk() (*string, bool)`

GetDisbursementBillingEventIDOk returns a tuple with the DisbursementBillingEventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetDisbursementBillingEventID(v string)`

SetDisbursementBillingEventID sets DisbursementBillingEventID field to given value.

### HasDisbursementBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasDisbursementBillingEventID() bool`

HasDisbursementBillingEventID returns a boolean if a field has been set.

### GetEndUserAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetEndUserAccountID() string`

GetEndUserAccountID returns the EndUserAccountID field if non-nil, zero value otherwise.

### GetEndUserAccountIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetEndUserAccountIDOk() (*string, bool)`

GetEndUserAccountIDOk returns a tuple with the EndUserAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetEndUserAccountID(v string)`

SetEndUserAccountID sets EndUserAccountID field to given value.

### HasEndUserAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasEndUserAccountID() bool`

HasEndUserAccountID returns a boolean if a field has been set.

### GetEntitlementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetFromAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetFromAccountID() string`

GetFromAccountID returns the FromAccountID field if non-nil, zero value otherwise.

### GetFromAccountIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetFromAccountIDOk() (*string, bool)`

GetFromAccountIDOk returns a tuple with the FromAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetFromAccountID(v string)`

SetFromAccountID sets FromAccountID field to given value.

### HasFromAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasFromAccountID() bool`

HasFromAccountID returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInsertDate() SqlNullTime`

GetInsertDate returns the InsertDate field if non-nil, zero value otherwise.

### GetInsertDateOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInsertDateOk() (*SqlNullTime, bool)`

GetInsertDateOk returns a tuple with the InsertDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetInsertDate(v SqlNullTime)`

SetInsertDate sets InsertDate field to given value.

### HasInsertDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasInsertDate() bool`

HasInsertDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInvoiceDate() SqlNullTime`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInvoiceDateOk() (*SqlNullTime, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetInvoiceDate(v SqlNullTime)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetOfferID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetOfferID() string`

GetOfferID returns the OfferID field if non-nil, zero value otherwise.

### GetOfferIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetOfferIDOk() (*string, bool)`

GetOfferIDOk returns a tuple with the OfferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetOfferID(v string)`

SetOfferID sets OfferID field to given value.

### HasOfferID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasOfferID() bool`

HasOfferID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetParentBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetParentBillingEventID() string`

GetParentBillingEventID returns the ParentBillingEventID field if non-nil, zero value otherwise.

### GetParentBillingEventIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetParentBillingEventIDOk() (*string, bool)`

GetParentBillingEventIDOk returns a tuple with the ParentBillingEventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetParentBillingEventID(v string)`

SetParentBillingEventID sets ParentBillingEventID field to given value.

### HasParentBillingEventID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasParentBillingEventID() bool`

HasParentBillingEventID returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetPaymentDueDate() SqlNullTime`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetPaymentDueDateOk() (*SqlNullTime, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetPaymentDueDate(v SqlNullTime)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetToAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetToAccountID() string`

GetToAccountID returns the ToAccountID field if non-nil, zero value otherwise.

### GetToAccountIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetToAccountIDOk() (*string, bool)`

GetToAccountIDOk returns a tuple with the ToAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetToAccountID(v string)`

SetToAccountID sets ToAccountID field to given value.

### HasToAccountID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasToAccountID() bool`

HasToAccountID returns a boolean if a field has been set.

### GetTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetTransactionReferenceID() string`

GetTransactionReferenceID returns the TransactionReferenceID field if non-nil, zero value otherwise.

### GetTransactionReferenceIDOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetTransactionReferenceIDOk() (*string, bool)`

GetTransactionReferenceIDOk returns a tuple with the TransactionReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetTransactionReferenceID(v string)`

SetTransactionReferenceID sets TransactionReferenceID field to given value.

### HasTransactionReferenceID

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasTransactionReferenceID() bool`

HasTransactionReferenceID returns a boolean if a field has been set.

### GetTransactionType

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetUsagePeriodEndDate() SqlNullTime`

GetUsagePeriodEndDate returns the UsagePeriodEndDate field if non-nil, zero value otherwise.

### GetUsagePeriodEndDateOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetUsagePeriodEndDateOk() (*SqlNullTime, bool)`

GetUsagePeriodEndDateOk returns a tuple with the UsagePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetUsagePeriodEndDate(v SqlNullTime)`

SetUsagePeriodEndDate sets UsagePeriodEndDate field to given value.

### HasUsagePeriodEndDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasUsagePeriodEndDate() bool`

HasUsagePeriodEndDate returns a boolean if a field has been set.

### GetUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetUsagePeriodStartDate() SqlNullTime`

GetUsagePeriodStartDate returns the UsagePeriodStartDate field if non-nil, zero value otherwise.

### GetUsagePeriodStartDateOk

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) GetUsagePeriodStartDateOk() (*SqlNullTime, bool)`

GetUsagePeriodStartDateOk returns a tuple with the UsagePeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) SetUsagePeriodStartDate(v SqlNullTime)`

SetUsagePeriodStartDate sets UsagePeriodStartDate field to given value.

### HasUsagePeriodStartDate

`func (o *GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent) HasUsagePeriodStartDate() bool`

HasUsagePeriodStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


