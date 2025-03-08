# BillingInvoiceInfo

## Properties

 Name                               | Type                                                                                             | Description                                                                                                                                                                                    | Notes      
------------------------------------|--------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **AddFixedFees**                   | Pointer to [**[]InvoiceAddFixedFee**](InvoiceAddFixedFee.md)                                     | Adjust charge fields The fixed fees to be added to the invoice.                                                                                                                                | [optional] 
 **AddonDetail**                    | Pointer to [**BillingAddonRecord**](BillingAddonRecord.md)                                       |                                                                                                                                                                                                | [optional] 
 **AdjustDiscountByDimensions**     | Pointer to [**[]InvoiceAdjustDiscountByDimension**](InvoiceAdjustDiscountByDimension.md)         | add or adjust discount for a specific dimension                                                                                                                                                | [optional] 
 **AdjustMinimumSpendByDimensions** | Pointer to [**[]InvoiceAdjustMinimumSpendByDimension**](InvoiceAdjustMinimumSpendByDimension.md) | add or adjust minimum spend for a specific dimension                                                                                                                                           | [optional] 
 **AdjustOverallDiscount**          | Pointer to [**InvoiceAdjustOverallDiscount**](InvoiceAdjustOverallDiscount.md)                   | add or adjust overall discount calculate each dimension&#39;s discount first, then apply the overall discount                                                                                  | [optional] 
 **AdjustOverallMinimumSpend**      | Pointer to [**InvoiceAdjustOverallMinimumSpend**](InvoiceAdjustOverallMinimumSpend.md)           | add or adjust overall minimum spend calculate each dimension&#39;s minimum spend first, then apply the overall minimum spend                                                                   | [optional] 
 **BillableDimensionDetails**       | Pointer to [**[]BillableDimensionPriceModelDetail**](BillableDimensionPriceModelDetail.md)       |                                                                                                                                                                                                | [optional] 
 **CommitsRevenueDetails**          | Pointer to [**[]CommitRevenueDetail**](CommitRevenueDetail.md)                                   | Recurring flat fee for the invoice. There should be only one type fee for each invoice, commits, or usage.                                                                                     | [optional] 
 **CreationDate**                   | Pointer to **time.Time**                                                                         | The creation date of the invoice when the status of the invoice may be draft or issued. It may be different from the issue date.                                                               | [optional] 
 **Currency**                       | Pointer to **string**                                                                            |                                                                                                                                                                                                | [optional] 
 **DeductedCommitAmount**           | Pointer to **int32**                                                                             | The amount of the committed amount that has been deducted from the usage. It works only when IsMeteringOverageCommit is true.                                                                  | [optional] 
 **DeductedCommitInvoiceID**        | Pointer to **string**                                                                            | The ID of the commit invoice that has been deducted from the usage. It works only when IsMeteringOverageCommit is true.                                                                        | [optional] 
 **Description**                    | Pointer to **string**                                                                            |                                                                                                                                                                                                | [optional] 
 **DueAmount**                      | Pointer to **float32**                                                                           | Due amount &#x3D; SubtotalAmount + TaxAmount - AdjustOverallDiscount                                                                                                                           | [optional] 
 **DueDate**                        | Pointer to **time.Time**                                                                         | DueDate &#x3D; IssueDate + NetTerm                                                                                                                                                             | [optional] 
 **GracePeriodInDays**              | Pointer to **int32**                                                                             | Grace Period in number of days                                                                                                                                                                 | [optional] 
 **IsMeteringOverageCommit**        | Pointer to **bool**                                                                              | Whether the usage metering is charged for the amount that exceeds the committed amount from the entitlement.                                                                                   | [optional] 
 **IssueDate**                      | Pointer to **time.Time**                                                                         | IssueDate, issue invoice automatically when CreationDate + GracePeriod, or issue invoice manually IssueDate &gt;&#x3D; CreationDate &amp;&amp; IssueDate &lt;&#x3D; CreationDate + GracePeriod | [optional] 
 **Memo**                           | Pointer to **string**                                                                            |                                                                                                                                                                                                | [optional] 
 **NetTermsInDays**                 | Pointer to **int32**                                                                             | Net Terms period in number of days                                                                                                                                                             | [optional] 
 **PaymentInstallmentsDetail**      | Pointer to [**BillingPaymentInstallmentDetail**](BillingPaymentInstallmentDetail.md)             |                                                                                                                                                                                                | [optional] 
 **ReceiptUrl**                     | Pointer to **string**                                                                            | Invoice receipt url, it only exists when there are transactions.                                                                                                                               | [optional] 
 **SpaUrl**                         | Pointer to **string**                                                                            | SPA url with JWT.                                                                                                                                                                              | [optional] 
 **SubtotalAmount**                 | Pointer to **float32**                                                                           | Subtotal amount calculated from the user usage.                                                                                                                                                | [optional] 
 **TaxAmount**                      | Pointer to **float32**                                                                           |                                                                                                                                                                                                | [optional] 
 **TrialPeriodInDays**              | Pointer to **int32**                                                                             | Trial period in number of days                                                                                                                                                                 | [optional] 
 **UsageDailyRevenues**             | Pointer to [**[]BillableDimensionUsageDailyRevenue**](BillableDimensionUsageDailyRevenue.md)     | Billable dimension fees for the invoice.                                                                                                                                                       | [optional] 

## Methods

### NewBillingInvoiceInfo

`func NewBillingInvoiceInfo() *BillingInvoiceInfo`

NewBillingInvoiceInfo instantiates a new BillingInvoiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceInfoWithDefaults

`func NewBillingInvoiceInfoWithDefaults() *BillingInvoiceInfo`

NewBillingInvoiceInfoWithDefaults instantiates a new BillingInvoiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddFixedFees

`func (o *BillingInvoiceInfo) GetAddFixedFees() []InvoiceAddFixedFee`

GetAddFixedFees returns the AddFixedFees field if non-nil, zero value otherwise.

### GetAddFixedFeesOk

`func (o *BillingInvoiceInfo) GetAddFixedFeesOk() (*[]InvoiceAddFixedFee, bool)`

GetAddFixedFeesOk returns a tuple with the AddFixedFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFixedFees

`func (o *BillingInvoiceInfo) SetAddFixedFees(v []InvoiceAddFixedFee)`

SetAddFixedFees sets AddFixedFees field to given value.

### HasAddFixedFees

`func (o *BillingInvoiceInfo) HasAddFixedFees() bool`

HasAddFixedFees returns a boolean if a field has been set.

### GetAddonDetail

`func (o *BillingInvoiceInfo) GetAddonDetail() BillingAddonRecord`

GetAddonDetail returns the AddonDetail field if non-nil, zero value otherwise.

### GetAddonDetailOk

`func (o *BillingInvoiceInfo) GetAddonDetailOk() (*BillingAddonRecord, bool)`

GetAddonDetailOk returns a tuple with the AddonDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonDetail

`func (o *BillingInvoiceInfo) SetAddonDetail(v BillingAddonRecord)`

SetAddonDetail sets AddonDetail field to given value.

### HasAddonDetail

`func (o *BillingInvoiceInfo) HasAddonDetail() bool`

HasAddonDetail returns a boolean if a field has been set.

### GetAdjustDiscountByDimensions

`func (o *BillingInvoiceInfo) GetAdjustDiscountByDimensions() []InvoiceAdjustDiscountByDimension`

GetAdjustDiscountByDimensions returns the AdjustDiscountByDimensions field if non-nil, zero value otherwise.

### GetAdjustDiscountByDimensionsOk

`func (o *BillingInvoiceInfo) GetAdjustDiscountByDimensionsOk() (*[]InvoiceAdjustDiscountByDimension, bool)`

GetAdjustDiscountByDimensionsOk returns a tuple with the AdjustDiscountByDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustDiscountByDimensions

`func (o *BillingInvoiceInfo) SetAdjustDiscountByDimensions(v []InvoiceAdjustDiscountByDimension)`

SetAdjustDiscountByDimensions sets AdjustDiscountByDimensions field to given value.

### HasAdjustDiscountByDimensions

`func (o *BillingInvoiceInfo) HasAdjustDiscountByDimensions() bool`

HasAdjustDiscountByDimensions returns a boolean if a field has been set.

### GetAdjustMinimumSpendByDimensions

`func (o *BillingInvoiceInfo) GetAdjustMinimumSpendByDimensions() []InvoiceAdjustMinimumSpendByDimension`

GetAdjustMinimumSpendByDimensions returns the AdjustMinimumSpendByDimensions field if non-nil, zero value otherwise.

### GetAdjustMinimumSpendByDimensionsOk

`func (o *BillingInvoiceInfo) GetAdjustMinimumSpendByDimensionsOk() (*[]InvoiceAdjustMinimumSpendByDimension, bool)`

GetAdjustMinimumSpendByDimensionsOk returns a tuple with the AdjustMinimumSpendByDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustMinimumSpendByDimensions

`func (o *BillingInvoiceInfo) SetAdjustMinimumSpendByDimensions(v []InvoiceAdjustMinimumSpendByDimension)`

SetAdjustMinimumSpendByDimensions sets AdjustMinimumSpendByDimensions field to given value.

### HasAdjustMinimumSpendByDimensions

`func (o *BillingInvoiceInfo) HasAdjustMinimumSpendByDimensions() bool`

HasAdjustMinimumSpendByDimensions returns a boolean if a field has been set.

### GetAdjustOverallDiscount

`func (o *BillingInvoiceInfo) GetAdjustOverallDiscount() InvoiceAdjustOverallDiscount`

GetAdjustOverallDiscount returns the AdjustOverallDiscount field if non-nil, zero value otherwise.

### GetAdjustOverallDiscountOk

`func (o *BillingInvoiceInfo) GetAdjustOverallDiscountOk() (*InvoiceAdjustOverallDiscount, bool)`

GetAdjustOverallDiscountOk returns a tuple with the AdjustOverallDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustOverallDiscount

`func (o *BillingInvoiceInfo) SetAdjustOverallDiscount(v InvoiceAdjustOverallDiscount)`

SetAdjustOverallDiscount sets AdjustOverallDiscount field to given value.

### HasAdjustOverallDiscount

`func (o *BillingInvoiceInfo) HasAdjustOverallDiscount() bool`

HasAdjustOverallDiscount returns a boolean if a field has been set.

### GetAdjustOverallMinimumSpend

`func (o *BillingInvoiceInfo) GetAdjustOverallMinimumSpend() InvoiceAdjustOverallMinimumSpend`

GetAdjustOverallMinimumSpend returns the AdjustOverallMinimumSpend field if non-nil, zero value otherwise.

### GetAdjustOverallMinimumSpendOk

`func (o *BillingInvoiceInfo) GetAdjustOverallMinimumSpendOk() (*InvoiceAdjustOverallMinimumSpend, bool)`

GetAdjustOverallMinimumSpendOk returns a tuple with the AdjustOverallMinimumSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustOverallMinimumSpend

`func (o *BillingInvoiceInfo) SetAdjustOverallMinimumSpend(v InvoiceAdjustOverallMinimumSpend)`

SetAdjustOverallMinimumSpend sets AdjustOverallMinimumSpend field to given value.

### HasAdjustOverallMinimumSpend

`func (o *BillingInvoiceInfo) HasAdjustOverallMinimumSpend() bool`

HasAdjustOverallMinimumSpend returns a boolean if a field has been set.

### GetBillableDimensionDetails

`func (o *BillingInvoiceInfo) GetBillableDimensionDetails() []BillableDimensionPriceModelDetail`

GetBillableDimensionDetails returns the BillableDimensionDetails field if non-nil, zero value otherwise.

### GetBillableDimensionDetailsOk

`func (o *BillingInvoiceInfo) GetBillableDimensionDetailsOk() (*[]BillableDimensionPriceModelDetail, bool)`

GetBillableDimensionDetailsOk returns a tuple with the BillableDimensionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDimensionDetails

`func (o *BillingInvoiceInfo) SetBillableDimensionDetails(v []BillableDimensionPriceModelDetail)`

SetBillableDimensionDetails sets BillableDimensionDetails field to given value.

### HasBillableDimensionDetails

`func (o *BillingInvoiceInfo) HasBillableDimensionDetails() bool`

HasBillableDimensionDetails returns a boolean if a field has been set.

### GetCommitsRevenueDetails

`func (o *BillingInvoiceInfo) GetCommitsRevenueDetails() []CommitRevenueDetail`

GetCommitsRevenueDetails returns the CommitsRevenueDetails field if non-nil, zero value otherwise.

### GetCommitsRevenueDetailsOk

`func (o *BillingInvoiceInfo) GetCommitsRevenueDetailsOk() (*[]CommitRevenueDetail, bool)`

GetCommitsRevenueDetailsOk returns a tuple with the CommitsRevenueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsRevenueDetails

`func (o *BillingInvoiceInfo) SetCommitsRevenueDetails(v []CommitRevenueDetail)`

SetCommitsRevenueDetails sets CommitsRevenueDetails field to given value.

### HasCommitsRevenueDetails

`func (o *BillingInvoiceInfo) HasCommitsRevenueDetails() bool`

HasCommitsRevenueDetails returns a boolean if a field has been set.

### GetCreationDate

`func (o *BillingInvoiceInfo) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *BillingInvoiceInfo) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *BillingInvoiceInfo) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *BillingInvoiceInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInvoiceInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoiceInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoiceInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInvoiceInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeductedCommitAmount

`func (o *BillingInvoiceInfo) GetDeductedCommitAmount() int32`

GetDeductedCommitAmount returns the DeductedCommitAmount field if non-nil, zero value otherwise.

### GetDeductedCommitAmountOk

`func (o *BillingInvoiceInfo) GetDeductedCommitAmountOk() (*int32, bool)`

GetDeductedCommitAmountOk returns a tuple with the DeductedCommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductedCommitAmount

`func (o *BillingInvoiceInfo) SetDeductedCommitAmount(v int32)`

SetDeductedCommitAmount sets DeductedCommitAmount field to given value.

### HasDeductedCommitAmount

`func (o *BillingInvoiceInfo) HasDeductedCommitAmount() bool`

HasDeductedCommitAmount returns a boolean if a field has been set.

### GetDeductedCommitInvoiceID

`func (o *BillingInvoiceInfo) GetDeductedCommitInvoiceID() string`

GetDeductedCommitInvoiceID returns the DeductedCommitInvoiceID field if non-nil, zero value otherwise.

### GetDeductedCommitInvoiceIDOk

`func (o *BillingInvoiceInfo) GetDeductedCommitInvoiceIDOk() (*string, bool)`

GetDeductedCommitInvoiceIDOk returns a tuple with the DeductedCommitInvoiceID field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetDeductedCommitInvoiceID

`func (o *BillingInvoiceInfo) SetDeductedCommitInvoiceID(v string)`

SetDeductedCommitInvoiceID sets DeductedCommitInvoiceID field to given value.

### HasDeductedCommitInvoiceID

`func (o *BillingInvoiceInfo) HasDeductedCommitInvoiceID() bool`

HasDeductedCommitInvoiceID returns a boolean if a field has been set.

### GetDescription

`func (o *BillingInvoiceInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingInvoiceInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingInvoiceInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingInvoiceInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueAmount

`func (o *BillingInvoiceInfo) GetDueAmount() float32`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *BillingInvoiceInfo) GetDueAmountOk() (*float32, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *BillingInvoiceInfo) SetDueAmount(v float32)`

SetDueAmount sets DueAmount field to given value.

### HasDueAmount

`func (o *BillingInvoiceInfo) HasDueAmount() bool`

HasDueAmount returns a boolean if a field has been set.

### GetDueDate

`func (o *BillingInvoiceInfo) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingInvoiceInfo) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingInvoiceInfo) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingInvoiceInfo) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetGracePeriodInDays

`func (o *BillingInvoiceInfo) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *BillingInvoiceInfo) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *BillingInvoiceInfo) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *BillingInvoiceInfo) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetIsMeteringOverageCommit

`func (o *BillingInvoiceInfo) GetIsMeteringOverageCommit() bool`

GetIsMeteringOverageCommit returns the IsMeteringOverageCommit field if non-nil, zero value otherwise.

### GetIsMeteringOverageCommitOk

`func (o *BillingInvoiceInfo) GetIsMeteringOverageCommitOk() (*bool, bool)`

GetIsMeteringOverageCommitOk returns a tuple with the IsMeteringOverageCommit field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetIsMeteringOverageCommit

`func (o *BillingInvoiceInfo) SetIsMeteringOverageCommit(v bool)`

SetIsMeteringOverageCommit sets IsMeteringOverageCommit field to given value.

### HasIsMeteringOverageCommit

`func (o *BillingInvoiceInfo) HasIsMeteringOverageCommit() bool`

HasIsMeteringOverageCommit returns a boolean if a field has been set.

### GetIssueDate

`func (o *BillingInvoiceInfo) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *BillingInvoiceInfo) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *BillingInvoiceInfo) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *BillingInvoiceInfo) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetMemo

`func (o *BillingInvoiceInfo) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *BillingInvoiceInfo) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *BillingInvoiceInfo) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *BillingInvoiceInfo) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetNetTermsInDays

`func (o *BillingInvoiceInfo) GetNetTermsInDays() int32`

GetNetTermsInDays returns the NetTermsInDays field if non-nil, zero value otherwise.

### GetNetTermsInDaysOk

`func (o *BillingInvoiceInfo) GetNetTermsInDaysOk() (*int32, bool)`

GetNetTermsInDaysOk returns a tuple with the NetTermsInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTermsInDays

`func (o *BillingInvoiceInfo) SetNetTermsInDays(v int32)`

SetNetTermsInDays sets NetTermsInDays field to given value.

### HasNetTermsInDays

`func (o *BillingInvoiceInfo) HasNetTermsInDays() bool`

HasNetTermsInDays returns a boolean if a field has been set.

### GetPaymentInstallmentsDetail

`func (o *BillingInvoiceInfo) GetPaymentInstallmentsDetail() BillingPaymentInstallmentDetail`

GetPaymentInstallmentsDetail returns the PaymentInstallmentsDetail field if non-nil, zero value otherwise.

### GetPaymentInstallmentsDetailOk

`func (o *BillingInvoiceInfo) GetPaymentInstallmentsDetailOk() (*BillingPaymentInstallmentDetail, bool)`

GetPaymentInstallmentsDetailOk returns a tuple with the PaymentInstallmentsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstallmentsDetail

`func (o *BillingInvoiceInfo) SetPaymentInstallmentsDetail(v BillingPaymentInstallmentDetail)`

SetPaymentInstallmentsDetail sets PaymentInstallmentsDetail field to given value.

### HasPaymentInstallmentsDetail

`func (o *BillingInvoiceInfo) HasPaymentInstallmentsDetail() bool`

HasPaymentInstallmentsDetail returns a boolean if a field has been set.

### GetReceiptUrl

`func (o *BillingInvoiceInfo) GetReceiptUrl() string`

GetReceiptUrl returns the ReceiptUrl field if non-nil, zero value otherwise.

### GetReceiptUrlOk

`func (o *BillingInvoiceInfo) GetReceiptUrlOk() (*string, bool)`

GetReceiptUrlOk returns a tuple with the ReceiptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptUrl

`func (o *BillingInvoiceInfo) SetReceiptUrl(v string)`

SetReceiptUrl sets ReceiptUrl field to given value.

### HasReceiptUrl

`func (o *BillingInvoiceInfo) HasReceiptUrl() bool`

HasReceiptUrl returns a boolean if a field has been set.

### GetSpaUrl

`func (o *BillingInvoiceInfo) GetSpaUrl() string`

GetSpaUrl returns the SpaUrl field if non-nil, zero value otherwise.

### GetSpaUrlOk

`func (o *BillingInvoiceInfo) GetSpaUrlOk() (*string, bool)`

GetSpaUrlOk returns a tuple with the SpaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUrl

`func (o *BillingInvoiceInfo) SetSpaUrl(v string)`

SetSpaUrl sets SpaUrl field to given value.

### HasSpaUrl

`func (o *BillingInvoiceInfo) HasSpaUrl() bool`

HasSpaUrl returns a boolean if a field has been set.

### GetSubtotalAmount

`func (o *BillingInvoiceInfo) GetSubtotalAmount() float32`

GetSubtotalAmount returns the SubtotalAmount field if non-nil, zero value otherwise.

### GetSubtotalAmountOk

`func (o *BillingInvoiceInfo) GetSubtotalAmountOk() (*float32, bool)`

GetSubtotalAmountOk returns a tuple with the SubtotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalAmount

`func (o *BillingInvoiceInfo) SetSubtotalAmount(v float32)`

SetSubtotalAmount sets SubtotalAmount field to given value.

### HasSubtotalAmount

`func (o *BillingInvoiceInfo) HasSubtotalAmount() bool`

HasSubtotalAmount returns a boolean if a field has been set.

### GetTaxAmount

`func (o *BillingInvoiceInfo) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *BillingInvoiceInfo) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *BillingInvoiceInfo) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *BillingInvoiceInfo) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTrialPeriodInDays

`func (o *BillingInvoiceInfo) GetTrialPeriodInDays() int32`

GetTrialPeriodInDays returns the TrialPeriodInDays field if non-nil, zero value otherwise.

### GetTrialPeriodInDaysOk

`func (o *BillingInvoiceInfo) GetTrialPeriodInDaysOk() (*int32, bool)`

GetTrialPeriodInDaysOk returns a tuple with the TrialPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodInDays

`func (o *BillingInvoiceInfo) SetTrialPeriodInDays(v int32)`

SetTrialPeriodInDays sets TrialPeriodInDays field to given value.

### HasTrialPeriodInDays

`func (o *BillingInvoiceInfo) HasTrialPeriodInDays() bool`

HasTrialPeriodInDays returns a boolean if a field has been set.

### GetUsageDailyRevenues

`func (o *BillingInvoiceInfo) GetUsageDailyRevenues() []BillableDimensionUsageDailyRevenue`

GetUsageDailyRevenues returns the UsageDailyRevenues field if non-nil, zero value otherwise.

### GetUsageDailyRevenuesOk

`func (o *BillingInvoiceInfo) GetUsageDailyRevenuesOk() (*[]BillableDimensionUsageDailyRevenue, bool)`

GetUsageDailyRevenuesOk returns a tuple with the UsageDailyRevenues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDailyRevenues

`func (o *BillingInvoiceInfo) SetUsageDailyRevenues(v []BillableDimensionUsageDailyRevenue)`

SetUsageDailyRevenues sets UsageDailyRevenues field to given value.

### HasUsageDailyRevenues

`func (o *BillingInvoiceInfo) HasUsageDailyRevenues() bool`

HasUsageDailyRevenues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


