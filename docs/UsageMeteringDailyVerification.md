# UsageMeteringDailyVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledAmount** | Pointer to **float32** | The amount of the usage metering billed by the cloud marketplace metering service. | [optional] 
**BilledDate** | Pointer to **time.Time** | The date when the usage metering is billed by the cloud marketplace metering service. | [optional] 
**BilledQuantity** | Pointer to **float32** | The quantity of the usage metering billed by the cloud marketplace metering service. | [optional] 
**IsAmountMatched** | Pointer to **bool** | Whether the amount is matched between reported &amp; billed usage metering. | [optional] 
**IsQuantityMatched** | Pointer to **bool** | Whether the quantity is matched between reported &amp; billed usage metering. | [optional] 
**Key** | Pointer to **string** | The dimension key of the usage metering. | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**ReportedAmount** | Pointer to **float32** | The amount of the usage metering reported to the cloud marketplace. | [optional] 
**ReportedDate** | Pointer to **time.Time** | The date when the usage metering is reported to the cloud marketplace. | [optional] 
**ReportedQuantity** | Pointer to **float32** | The quantity of the usage metering reported to the cloud marketplace. | [optional] 

## Methods

### NewUsageMeteringDailyVerification

`func NewUsageMeteringDailyVerification() *UsageMeteringDailyVerification`

NewUsageMeteringDailyVerification instantiates a new UsageMeteringDailyVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeteringDailyVerificationWithDefaults

`func NewUsageMeteringDailyVerificationWithDefaults() *UsageMeteringDailyVerification`

NewUsageMeteringDailyVerificationWithDefaults instantiates a new UsageMeteringDailyVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledAmount

`func (o *UsageMeteringDailyVerification) GetBilledAmount() float32`

GetBilledAmount returns the BilledAmount field if non-nil, zero value otherwise.

### GetBilledAmountOk

`func (o *UsageMeteringDailyVerification) GetBilledAmountOk() (*float32, bool)`

GetBilledAmountOk returns a tuple with the BilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledAmount

`func (o *UsageMeteringDailyVerification) SetBilledAmount(v float32)`

SetBilledAmount sets BilledAmount field to given value.

### HasBilledAmount

`func (o *UsageMeteringDailyVerification) HasBilledAmount() bool`

HasBilledAmount returns a boolean if a field has been set.

### GetBilledDate

`func (o *UsageMeteringDailyVerification) GetBilledDate() time.Time`

GetBilledDate returns the BilledDate field if non-nil, zero value otherwise.

### GetBilledDateOk

`func (o *UsageMeteringDailyVerification) GetBilledDateOk() (*time.Time, bool)`

GetBilledDateOk returns a tuple with the BilledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledDate

`func (o *UsageMeteringDailyVerification) SetBilledDate(v time.Time)`

SetBilledDate sets BilledDate field to given value.

### HasBilledDate

`func (o *UsageMeteringDailyVerification) HasBilledDate() bool`

HasBilledDate returns a boolean if a field has been set.

### GetBilledQuantity

`func (o *UsageMeteringDailyVerification) GetBilledQuantity() float32`

GetBilledQuantity returns the BilledQuantity field if non-nil, zero value otherwise.

### GetBilledQuantityOk

`func (o *UsageMeteringDailyVerification) GetBilledQuantityOk() (*float32, bool)`

GetBilledQuantityOk returns a tuple with the BilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledQuantity

`func (o *UsageMeteringDailyVerification) SetBilledQuantity(v float32)`

SetBilledQuantity sets BilledQuantity field to given value.

### HasBilledQuantity

`func (o *UsageMeteringDailyVerification) HasBilledQuantity() bool`

HasBilledQuantity returns a boolean if a field has been set.

### GetIsAmountMatched

`func (o *UsageMeteringDailyVerification) GetIsAmountMatched() bool`

GetIsAmountMatched returns the IsAmountMatched field if non-nil, zero value otherwise.

### GetIsAmountMatchedOk

`func (o *UsageMeteringDailyVerification) GetIsAmountMatchedOk() (*bool, bool)`

GetIsAmountMatchedOk returns a tuple with the IsAmountMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountMatched

`func (o *UsageMeteringDailyVerification) SetIsAmountMatched(v bool)`

SetIsAmountMatched sets IsAmountMatched field to given value.

### HasIsAmountMatched

`func (o *UsageMeteringDailyVerification) HasIsAmountMatched() bool`

HasIsAmountMatched returns a boolean if a field has been set.

### GetIsQuantityMatched

`func (o *UsageMeteringDailyVerification) GetIsQuantityMatched() bool`

GetIsQuantityMatched returns the IsQuantityMatched field if non-nil, zero value otherwise.

### GetIsQuantityMatchedOk

`func (o *UsageMeteringDailyVerification) GetIsQuantityMatchedOk() (*bool, bool)`

GetIsQuantityMatchedOk returns a tuple with the IsQuantityMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuantityMatched

`func (o *UsageMeteringDailyVerification) SetIsQuantityMatched(v bool)`

SetIsQuantityMatched sets IsQuantityMatched field to given value.

### HasIsQuantityMatched

`func (o *UsageMeteringDailyVerification) HasIsQuantityMatched() bool`

HasIsQuantityMatched returns a boolean if a field has been set.

### GetKey

`func (o *UsageMeteringDailyVerification) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UsageMeteringDailyVerification) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UsageMeteringDailyVerification) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UsageMeteringDailyVerification) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPartner

`func (o *UsageMeteringDailyVerification) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *UsageMeteringDailyVerification) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *UsageMeteringDailyVerification) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *UsageMeteringDailyVerification) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetReportedAmount

`func (o *UsageMeteringDailyVerification) GetReportedAmount() float32`

GetReportedAmount returns the ReportedAmount field if non-nil, zero value otherwise.

### GetReportedAmountOk

`func (o *UsageMeteringDailyVerification) GetReportedAmountOk() (*float32, bool)`

GetReportedAmountOk returns a tuple with the ReportedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedAmount

`func (o *UsageMeteringDailyVerification) SetReportedAmount(v float32)`

SetReportedAmount sets ReportedAmount field to given value.

### HasReportedAmount

`func (o *UsageMeteringDailyVerification) HasReportedAmount() bool`

HasReportedAmount returns a boolean if a field has been set.

### GetReportedDate

`func (o *UsageMeteringDailyVerification) GetReportedDate() time.Time`

GetReportedDate returns the ReportedDate field if non-nil, zero value otherwise.

### GetReportedDateOk

`func (o *UsageMeteringDailyVerification) GetReportedDateOk() (*time.Time, bool)`

GetReportedDateOk returns a tuple with the ReportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDate

`func (o *UsageMeteringDailyVerification) SetReportedDate(v time.Time)`

SetReportedDate sets ReportedDate field to given value.

### HasReportedDate

`func (o *UsageMeteringDailyVerification) HasReportedDate() bool`

HasReportedDate returns a boolean if a field has been set.

### GetReportedQuantity

`func (o *UsageMeteringDailyVerification) GetReportedQuantity() float32`

GetReportedQuantity returns the ReportedQuantity field if non-nil, zero value otherwise.

### GetReportedQuantityOk

`func (o *UsageMeteringDailyVerification) GetReportedQuantityOk() (*float32, bool)`

GetReportedQuantityOk returns a tuple with the ReportedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedQuantity

`func (o *UsageMeteringDailyVerification) SetReportedQuantity(v float32)`

SetReportedQuantity sets ReportedQuantity field to given value.

### HasReportedQuantity

`func (o *UsageMeteringDailyVerification) HasReportedQuantity() bool`

HasReportedQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


