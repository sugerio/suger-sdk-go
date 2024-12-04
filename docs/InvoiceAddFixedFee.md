# InvoiceAddFixedFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoiceAddFixedFee

`func NewInvoiceAddFixedFee() *InvoiceAddFixedFee`

NewInvoiceAddFixedFee instantiates a new InvoiceAddFixedFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAddFixedFeeWithDefaults

`func NewInvoiceAddFixedFeeWithDefaults() *InvoiceAddFixedFee`

NewInvoiceAddFixedFeeWithDefaults instantiates a new InvoiceAddFixedFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *InvoiceAddFixedFee) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceAddFixedFee) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceAddFixedFee) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceAddFixedFee) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceAddFixedFee) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceAddFixedFee) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceAddFixedFee) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceAddFixedFee) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRate

`func (o *InvoiceAddFixedFee) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InvoiceAddFixedFee) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InvoiceAddFixedFee) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *InvoiceAddFixedFee) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetReason

`func (o *InvoiceAddFixedFee) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceAddFixedFee) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceAddFixedFee) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceAddFixedFee) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStartDate

`func (o *InvoiceAddFixedFee) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceAddFixedFee) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceAddFixedFee) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceAddFixedFee) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


