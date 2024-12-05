# BillingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BillingInvoiceInfo**](BillingInvoiceInfo.md) |  | [optional] 
**InvoiceURL** | Pointer to **string** | The invoice file URL, provided as AWS S3 presigned URL with expiration time. Output only. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to [**BillingPaymentStatus**](BillingPaymentStatus.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**BillingInvoiceStatus**](BillingInvoiceStatus.md) |  | [optional] 
**Type** | Pointer to [**BillingInvoiceType**](BillingInvoiceType.md) |  | [optional] 

## Methods

### NewBillingInvoice

`func NewBillingInvoice() *BillingInvoice`

NewBillingInvoice instantiates a new BillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceWithDefaults

`func NewBillingInvoiceWithDefaults() *BillingInvoice`

NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *BillingInvoice) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *BillingInvoice) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *BillingInvoice) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *BillingInvoice) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *BillingInvoice) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BillingInvoice) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BillingInvoice) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BillingInvoice) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingInvoice) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingInvoice) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingInvoice) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingInvoice) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntitlementID

`func (o *BillingInvoice) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *BillingInvoice) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *BillingInvoice) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *BillingInvoice) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetId

`func (o *BillingInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BillingInvoice) GetInfo() BillingInvoiceInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingInvoice) GetInfoOk() (*BillingInvoiceInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingInvoice) SetInfo(v BillingInvoiceInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingInvoice) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetInvoiceURL

`func (o *BillingInvoice) GetInvoiceURL() string`

GetInvoiceURL returns the InvoiceURL field if non-nil, zero value otherwise.

### GetInvoiceURLOk

`func (o *BillingInvoice) GetInvoiceURLOk() (*string, bool)`

GetInvoiceURLOk returns a tuple with the InvoiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceURL

`func (o *BillingInvoice) SetInvoiceURL(v string)`

SetInvoiceURL sets InvoiceURL field to given value.

### HasInvoiceURL

`func (o *BillingInvoice) HasInvoiceURL() bool`

HasInvoiceURL returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *BillingInvoice) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *BillingInvoice) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *BillingInvoice) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *BillingInvoice) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetOrganizationID

`func (o *BillingInvoice) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *BillingInvoice) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *BillingInvoice) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *BillingInvoice) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *BillingInvoice) GetPaymentStatus() BillingPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *BillingInvoice) GetPaymentStatusOk() (*BillingPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *BillingInvoice) SetPaymentStatus(v BillingPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *BillingInvoice) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingInvoice) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingInvoice) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingInvoice) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingInvoice) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInvoice) GetStatus() BillingInvoiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInvoice) GetStatusOk() (*BillingInvoiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInvoice) SetStatus(v BillingInvoiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *BillingInvoice) GetType() BillingInvoiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingInvoice) GetTypeOk() (*BillingInvoiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingInvoice) SetType(v BillingInvoiceType)`

SetType sets Type field to given value.

### HasType

`func (o *BillingInvoice) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


