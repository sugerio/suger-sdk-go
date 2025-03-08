# RevenueRecord

## Properties

 Name                     | Type                                                     | Description                                                                | Notes      
--------------------------|----------------------------------------------------------|----------------------------------------------------------------------------|------------
 **Amount**               | Pointer to **float32**                                   | The revenue amount for the revenue report                                  | [optional] 
 **BuyerID**              | Pointer to **string**                                    |                                                                            | [optional] 
 **CollectableAmount**    | Pointer to **float32**                                   | The revenue amount that the seller/ISV can collect.                        | [optional] 
 **Currency**             | Pointer to **string**                                    | The currency of the revenue in ISO 4217 format, such as \&quot;USD\&quot;. | [optional] 
 **Date**                 | Pointer to **time.Time**                                 | The date for the revenue report                                            | [optional] 
 **DisburseAmount**       | Pointer to **float32**                                   |                                                                            | [optional] 
 **DisburseDate**         | Pointer to **time.Time**                                 |                                                                            | [optional] 
 **EntitlementID**        | Pointer to **string**                                    |                                                                            | [optional] 
 **Id**                   | Pointer to **string**                                    |                                                                            | [optional] 
 **Info**                 | Pointer to [**RevenueRecordInfo**](RevenueRecordInfo.md) |                                                                            | [optional] 
 **InvoiceAmount**        | Pointer to **float32**                                   |                                                                            | [optional] 
 **InvoiceDate**          | Pointer to **time.Time**                                 |                                                                            | [optional] 
 **OrganizationID**       | Pointer to **string**                                    |                                                                            | [optional] 
 **Partner**              | Pointer to **string**                                    | The value is from type Partner string                                      | [optional] 
 **PaymentDueDate**       | Pointer to **time.Time**                                 |                                                                            | [optional] 
 **ProductID**            | Pointer to **string**                                    |                                                                            | [optional] 
 **RefundDisburseAmount** | Pointer to **float32**                                   |                                                                            | [optional] 
 **RefundDisburseDate**   | Pointer to **time.Time**                                 |                                                                            | [optional] 
 **RefundInvoiceAmount**  | Pointer to **float32**                                   |                                                                            | [optional] 
 **RefundInvoiceDate**    | Pointer to **time.Time**                                 |                                                                            | [optional] 
 **TaxAmount**            | Pointer to **float32**                                   |                                                                            | [optional] 

## Methods

### NewRevenueRecord

`func NewRevenueRecord() *RevenueRecord`

NewRevenueRecord instantiates a new RevenueRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueRecordWithDefaults

`func NewRevenueRecordWithDefaults() *RevenueRecord`

NewRevenueRecordWithDefaults instantiates a new RevenueRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RevenueRecord) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RevenueRecord) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RevenueRecord) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RevenueRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBuyerID

`func (o *RevenueRecord) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *RevenueRecord) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *RevenueRecord) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *RevenueRecord) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCollectableAmount

`func (o *RevenueRecord) GetCollectableAmount() float32`

GetCollectableAmount returns the CollectableAmount field if non-nil, zero value otherwise.

### GetCollectableAmountOk

`func (o *RevenueRecord) GetCollectableAmountOk() (*float32, bool)`

GetCollectableAmountOk returns a tuple with the CollectableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectableAmount

`func (o *RevenueRecord) SetCollectableAmount(v float32)`

SetCollectableAmount sets CollectableAmount field to given value.

### HasCollectableAmount

`func (o *RevenueRecord) HasCollectableAmount() bool`

HasCollectableAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *RevenueRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RevenueRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RevenueRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RevenueRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *RevenueRecord) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RevenueRecord) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RevenueRecord) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *RevenueRecord) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDisburseAmount

`func (o *RevenueRecord) GetDisburseAmount() float32`

GetDisburseAmount returns the DisburseAmount field if non-nil, zero value otherwise.

### GetDisburseAmountOk

`func (o *RevenueRecord) GetDisburseAmountOk() (*float32, bool)`

GetDisburseAmountOk returns a tuple with the DisburseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisburseAmount

`func (o *RevenueRecord) SetDisburseAmount(v float32)`

SetDisburseAmount sets DisburseAmount field to given value.

### HasDisburseAmount

`func (o *RevenueRecord) HasDisburseAmount() bool`

HasDisburseAmount returns a boolean if a field has been set.

### GetDisburseDate

`func (o *RevenueRecord) GetDisburseDate() time.Time`

GetDisburseDate returns the DisburseDate field if non-nil, zero value otherwise.

### GetDisburseDateOk

`func (o *RevenueRecord) GetDisburseDateOk() (*time.Time, bool)`

GetDisburseDateOk returns a tuple with the DisburseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisburseDate

`func (o *RevenueRecord) SetDisburseDate(v time.Time)`

SetDisburseDate sets DisburseDate field to given value.

### HasDisburseDate

`func (o *RevenueRecord) HasDisburseDate() bool`

HasDisburseDate returns a boolean if a field has been set.

### GetEntitlementID

`func (o *RevenueRecord) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *RevenueRecord) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *RevenueRecord) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *RevenueRecord) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetId

`func (o *RevenueRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevenueRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevenueRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RevenueRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *RevenueRecord) GetInfo() RevenueRecordInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RevenueRecord) GetInfoOk() (*RevenueRecordInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RevenueRecord) SetInfo(v RevenueRecordInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RevenueRecord) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetInvoiceAmount

`func (o *RevenueRecord) GetInvoiceAmount() float32`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *RevenueRecord) GetInvoiceAmountOk() (*float32, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *RevenueRecord) SetInvoiceAmount(v float32)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *RevenueRecord) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *RevenueRecord) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *RevenueRecord) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *RevenueRecord) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *RevenueRecord) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetOrganizationID

`func (o *RevenueRecord) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *RevenueRecord) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *RevenueRecord) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *RevenueRecord) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *RevenueRecord) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *RevenueRecord) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *RevenueRecord) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *RevenueRecord) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *RevenueRecord) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *RevenueRecord) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *RevenueRecord) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *RevenueRecord) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetProductID

`func (o *RevenueRecord) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *RevenueRecord) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *RevenueRecord) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *RevenueRecord) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetRefundDisburseAmount

`func (o *RevenueRecord) GetRefundDisburseAmount() float32`

GetRefundDisburseAmount returns the RefundDisburseAmount field if non-nil, zero value otherwise.

### GetRefundDisburseAmountOk

`func (o *RevenueRecord) GetRefundDisburseAmountOk() (*float32, bool)`

GetRefundDisburseAmountOk returns a tuple with the RefundDisburseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDisburseAmount

`func (o *RevenueRecord) SetRefundDisburseAmount(v float32)`

SetRefundDisburseAmount sets RefundDisburseAmount field to given value.

### HasRefundDisburseAmount

`func (o *RevenueRecord) HasRefundDisburseAmount() bool`

HasRefundDisburseAmount returns a boolean if a field has been set.

### GetRefundDisburseDate

`func (o *RevenueRecord) GetRefundDisburseDate() time.Time`

GetRefundDisburseDate returns the RefundDisburseDate field if non-nil, zero value otherwise.

### GetRefundDisburseDateOk

`func (o *RevenueRecord) GetRefundDisburseDateOk() (*time.Time, bool)`

GetRefundDisburseDateOk returns a tuple with the RefundDisburseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDisburseDate

`func (o *RevenueRecord) SetRefundDisburseDate(v time.Time)`

SetRefundDisburseDate sets RefundDisburseDate field to given value.

### HasRefundDisburseDate

`func (o *RevenueRecord) HasRefundDisburseDate() bool`

HasRefundDisburseDate returns a boolean if a field has been set.

### GetRefundInvoiceAmount

`func (o *RevenueRecord) GetRefundInvoiceAmount() float32`

GetRefundInvoiceAmount returns the RefundInvoiceAmount field if non-nil, zero value otherwise.

### GetRefundInvoiceAmountOk

`func (o *RevenueRecord) GetRefundInvoiceAmountOk() (*float32, bool)`

GetRefundInvoiceAmountOk returns a tuple with the RefundInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInvoiceAmount

`func (o *RevenueRecord) SetRefundInvoiceAmount(v float32)`

SetRefundInvoiceAmount sets RefundInvoiceAmount field to given value.

### HasRefundInvoiceAmount

`func (o *RevenueRecord) HasRefundInvoiceAmount() bool`

HasRefundInvoiceAmount returns a boolean if a field has been set.

### GetRefundInvoiceDate

`func (o *RevenueRecord) GetRefundInvoiceDate() time.Time`

GetRefundInvoiceDate returns the RefundInvoiceDate field if non-nil, zero value otherwise.

### GetRefundInvoiceDateOk

`func (o *RevenueRecord) GetRefundInvoiceDateOk() (*time.Time, bool)`

GetRefundInvoiceDateOk returns a tuple with the RefundInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInvoiceDate

`func (o *RevenueRecord) SetRefundInvoiceDate(v time.Time)`

SetRefundInvoiceDate sets RefundInvoiceDate field to given value.

### HasRefundInvoiceDate

`func (o *RevenueRecord) HasRefundInvoiceDate() bool`

HasRefundInvoiceDate returns a boolean if a field has been set.

### GetTaxAmount

`func (o *RevenueRecord) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *RevenueRecord) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *RevenueRecord) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *RevenueRecord) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


