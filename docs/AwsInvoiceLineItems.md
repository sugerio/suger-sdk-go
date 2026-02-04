# AwsInvoiceLineItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]AwsInvoiceLineItemDetail**](AwsInvoiceLineItemDetail.md) |  | [optional] 
**LinkedAccountAllocations** | Pointer to [**[]AwsInvoiceLinkedAccountAllocation**](AwsInvoiceLinkedAccountAllocation.md) |  | [optional] 

## Methods

### NewAwsInvoiceLineItems

`func NewAwsInvoiceLineItems() *AwsInvoiceLineItems`

NewAwsInvoiceLineItems instantiates a new AwsInvoiceLineItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsInvoiceLineItemsWithDefaults

`func NewAwsInvoiceLineItemsWithDefaults() *AwsInvoiceLineItems`

NewAwsInvoiceLineItemsWithDefaults instantiates a new AwsInvoiceLineItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *AwsInvoiceLineItems) GetDetails() []AwsInvoiceLineItemDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AwsInvoiceLineItems) GetDetailsOk() (*[]AwsInvoiceLineItemDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AwsInvoiceLineItems) SetDetails(v []AwsInvoiceLineItemDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AwsInvoiceLineItems) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLinkedAccountAllocations

`func (o *AwsInvoiceLineItems) GetLinkedAccountAllocations() []AwsInvoiceLinkedAccountAllocation`

GetLinkedAccountAllocations returns the LinkedAccountAllocations field if non-nil, zero value otherwise.

### GetLinkedAccountAllocationsOk

`func (o *AwsInvoiceLineItems) GetLinkedAccountAllocationsOk() (*[]AwsInvoiceLinkedAccountAllocation, bool)`

GetLinkedAccountAllocationsOk returns a tuple with the LinkedAccountAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountAllocations

`func (o *AwsInvoiceLineItems) SetLinkedAccountAllocations(v []AwsInvoiceLinkedAccountAllocation)`

SetLinkedAccountAllocations sets LinkedAccountAllocations field to given value.

### HasLinkedAccountAllocations

`func (o *AwsInvoiceLineItems) HasLinkedAccountAllocations() bool`

HasLinkedAccountAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


