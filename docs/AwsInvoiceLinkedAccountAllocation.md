# AwsInvoiceLinkedAccountAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**AllocatedTotal** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]AwsInvoiceLineItemDetail**](AwsInvoiceLineItemDetail.md) |  | [optional] 

## Methods

### NewAwsInvoiceLinkedAccountAllocation

`func NewAwsInvoiceLinkedAccountAllocation() *AwsInvoiceLinkedAccountAllocation`

NewAwsInvoiceLinkedAccountAllocation instantiates a new AwsInvoiceLinkedAccountAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsInvoiceLinkedAccountAllocationWithDefaults

`func NewAwsInvoiceLinkedAccountAllocationWithDefaults() *AwsInvoiceLinkedAccountAllocation`

NewAwsInvoiceLinkedAccountAllocationWithDefaults instantiates a new AwsInvoiceLinkedAccountAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsInvoiceLinkedAccountAllocation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsInvoiceLinkedAccountAllocation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsInvoiceLinkedAccountAllocation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AwsInvoiceLinkedAccountAllocation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *AwsInvoiceLinkedAccountAllocation) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AwsInvoiceLinkedAccountAllocation) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AwsInvoiceLinkedAccountAllocation) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AwsInvoiceLinkedAccountAllocation) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAllocatedTotal

`func (o *AwsInvoiceLinkedAccountAllocation) GetAllocatedTotal() float32`

GetAllocatedTotal returns the AllocatedTotal field if non-nil, zero value otherwise.

### GetAllocatedTotalOk

`func (o *AwsInvoiceLinkedAccountAllocation) GetAllocatedTotalOk() (*float32, bool)`

GetAllocatedTotalOk returns a tuple with the AllocatedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedTotal

`func (o *AwsInvoiceLinkedAccountAllocation) SetAllocatedTotal(v float32)`

SetAllocatedTotal sets AllocatedTotal field to given value.

### HasAllocatedTotal

`func (o *AwsInvoiceLinkedAccountAllocation) HasAllocatedTotal() bool`

HasAllocatedTotal returns a boolean if a field has been set.

### GetCurrency

`func (o *AwsInvoiceLinkedAccountAllocation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AwsInvoiceLinkedAccountAllocation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AwsInvoiceLinkedAccountAllocation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AwsInvoiceLinkedAccountAllocation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDetails

`func (o *AwsInvoiceLinkedAccountAllocation) GetDetails() []AwsInvoiceLineItemDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AwsInvoiceLinkedAccountAllocation) GetDetailsOk() (*[]AwsInvoiceLineItemDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AwsInvoiceLinkedAccountAllocation) SetDetails(v []AwsInvoiceLineItemDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AwsInvoiceLinkedAccountAllocation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


