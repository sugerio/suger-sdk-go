# BillingPaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BillingPaymentTransactionInfo**](BillingPaymentTransactionInfo.md) |  | [optional] 
**InvoiceID** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**ParentID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**Status** | Pointer to [**BillingPaymentStatus**](BillingPaymentStatus.md) |  | [optional] 
**Type** | Pointer to [**BillingPaymentTransactionType**](BillingPaymentTransactionType.md) |  | [optional] 
**WalletID** | Pointer to **string** |  | [optional] 
**WalletType** | Pointer to [**BillingWalletType**](BillingWalletType.md) |  | [optional] 

## Methods

### NewBillingPaymentTransaction

`func NewBillingPaymentTransaction() *BillingPaymentTransaction`

NewBillingPaymentTransaction instantiates a new BillingPaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPaymentTransactionWithDefaults

`func NewBillingPaymentTransactionWithDefaults() *BillingPaymentTransaction`

NewBillingPaymentTransactionWithDefaults instantiates a new BillingPaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingPaymentTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingPaymentTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingPaymentTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingPaymentTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBuyerID

`func (o *BillingPaymentTransaction) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *BillingPaymentTransaction) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *BillingPaymentTransaction) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *BillingPaymentTransaction) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *BillingPaymentTransaction) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BillingPaymentTransaction) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BillingPaymentTransaction) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BillingPaymentTransaction) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingPaymentTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingPaymentTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingPaymentTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingPaymentTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEntitlementID

`func (o *BillingPaymentTransaction) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *BillingPaymentTransaction) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *BillingPaymentTransaction) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *BillingPaymentTransaction) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetId

`func (o *BillingPaymentTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPaymentTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPaymentTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingPaymentTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BillingPaymentTransaction) GetInfo() BillingPaymentTransactionInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingPaymentTransaction) GetInfoOk() (*BillingPaymentTransactionInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingPaymentTransaction) SetInfo(v BillingPaymentTransactionInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingPaymentTransaction) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetInvoiceID

`func (o *BillingPaymentTransaction) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *BillingPaymentTransaction) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *BillingPaymentTransaction) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *BillingPaymentTransaction) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *BillingPaymentTransaction) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *BillingPaymentTransaction) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *BillingPaymentTransaction) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *BillingPaymentTransaction) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetOrganizationID

`func (o *BillingPaymentTransaction) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *BillingPaymentTransaction) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *BillingPaymentTransaction) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *BillingPaymentTransaction) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetParentID

`func (o *BillingPaymentTransaction) GetParentID() string`

GetParentID returns the ParentID field if non-nil, zero value otherwise.

### GetParentIDOk

`func (o *BillingPaymentTransaction) GetParentIDOk() (*string, bool)`

GetParentIDOk returns a tuple with the ParentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentID

`func (o *BillingPaymentTransaction) SetParentID(v string)`

SetParentID sets ParentID field to given value.

### HasParentID

`func (o *BillingPaymentTransaction) HasParentID() bool`

HasParentID returns a boolean if a field has been set.

### GetPartner

`func (o *BillingPaymentTransaction) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *BillingPaymentTransaction) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *BillingPaymentTransaction) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *BillingPaymentTransaction) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetStatus

`func (o *BillingPaymentTransaction) GetStatus() BillingPaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingPaymentTransaction) GetStatusOk() (*BillingPaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingPaymentTransaction) SetStatus(v BillingPaymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingPaymentTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *BillingPaymentTransaction) GetType() BillingPaymentTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingPaymentTransaction) GetTypeOk() (*BillingPaymentTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingPaymentTransaction) SetType(v BillingPaymentTransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *BillingPaymentTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWalletID

`func (o *BillingPaymentTransaction) GetWalletID() string`

GetWalletID returns the WalletID field if non-nil, zero value otherwise.

### GetWalletIDOk

`func (o *BillingPaymentTransaction) GetWalletIDOk() (*string, bool)`

GetWalletIDOk returns a tuple with the WalletID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletID

`func (o *BillingPaymentTransaction) SetWalletID(v string)`

SetWalletID sets WalletID field to given value.

### HasWalletID

`func (o *BillingPaymentTransaction) HasWalletID() bool`

HasWalletID returns a boolean if a field has been set.

### GetWalletType

`func (o *BillingPaymentTransaction) GetWalletType() BillingWalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *BillingPaymentTransaction) GetWalletTypeOk() (*BillingWalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *BillingPaymentTransaction) SetWalletType(v BillingWalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *BillingPaymentTransaction) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


