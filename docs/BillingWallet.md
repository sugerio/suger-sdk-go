# BillingWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **time.Time** | nullable | [optional] 
**ExternalID** | Pointer to **string** | The payment method id in payment provider, such as stripe payment method id. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BillingWalletInfo**](BillingWalletInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**BillingWalletStatus**](BillingWalletStatus.md) |  | [optional] 
**TotalAmount** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to [**BillingWalletType**](BillingWalletType.md) |  | [optional] 
**UsedAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingWallet

`func NewBillingWallet() *BillingWallet`

NewBillingWallet instantiates a new BillingWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWalletWithDefaults

`func NewBillingWalletWithDefaults() *BillingWallet`

NewBillingWalletWithDefaults instantiates a new BillingWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *BillingWallet) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *BillingWallet) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *BillingWallet) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *BillingWallet) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *BillingWallet) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BillingWallet) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BillingWallet) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BillingWallet) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingWallet) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingWallet) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingWallet) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingWallet) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpireTime

`func (o *BillingWallet) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *BillingWallet) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *BillingWallet) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *BillingWallet) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetExternalID

`func (o *BillingWallet) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BillingWallet) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BillingWallet) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BillingWallet) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetId

`func (o *BillingWallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingWallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingWallet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BillingWallet) GetInfo() BillingWalletInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingWallet) GetInfoOk() (*BillingWalletInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingWallet) SetInfo(v BillingWalletInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingWallet) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *BillingWallet) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *BillingWallet) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *BillingWallet) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *BillingWallet) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *BillingWallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingWallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingWallet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingWallet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationID

`func (o *BillingWallet) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *BillingWallet) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *BillingWallet) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *BillingWallet) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *BillingWallet) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *BillingWallet) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *BillingWallet) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *BillingWallet) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetStartTime

`func (o *BillingWallet) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BillingWallet) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BillingWallet) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BillingWallet) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *BillingWallet) GetStatus() BillingWalletStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingWallet) GetStatusOk() (*BillingWalletStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingWallet) SetStatus(v BillingWalletStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingWallet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalAmount

`func (o *BillingWallet) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *BillingWallet) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *BillingWallet) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *BillingWallet) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetType

`func (o *BillingWallet) GetType() BillingWalletType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingWallet) GetTypeOk() (*BillingWalletType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingWallet) SetType(v BillingWalletType)`

SetType sets Type field to given value.

### HasType

`func (o *BillingWallet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsedAmount

`func (o *BillingWallet) GetUsedAmount() float32`

GetUsedAmount returns the UsedAmount field if non-nil, zero value otherwise.

### GetUsedAmountOk

`func (o *BillingWallet) GetUsedAmountOk() (*float32, bool)`

GetUsedAmountOk returns a tuple with the UsedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAmount

`func (o *BillingWallet) SetUsedAmount(v float32)`

SetUsedAmount sets UsedAmount field to given value.

### HasUsedAmount

`func (o *BillingWallet) HasUsedAmount() bool`

HasUsedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


