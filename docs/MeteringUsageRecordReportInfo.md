# MeteringUsageRecordReportInfo

## Properties

 Name                                           | Type                                                                                                                                                                                                                       | Description                                                                                                                                                                                                              | Notes      
------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **AggregatedBillableRecords**                  | Pointer to [**[]AggregatedMeteringUsageRecord**](AggregatedMeteringUsageRecord.md)                                                                                                                                         | The aggregated billable records from the usage metering API v2.                                                                                                                                                          | [optional] 
 **AlibabaMeteringRequest**                     | Pointer to [**ClientPushMeteringDataRequest**](ClientPushMeteringDataRequest.md)                                                                                                                                           | The raw request to call Alibaba metering service.                                                                                                                                                                        | [optional] 
 **AlibabaMeteringResponse**                    | Pointer to [**ClientPushMeteringDataResponseBody**](ClientPushMeteringDataResponseBody.md)                                                                                                                                 | The raw response from Alibaba metering service.                                                                                                                                                                          | [optional] 
 **AwsMeteringRequest**                         | Pointer to [**AwsMarketplaceMeteringBatchMeterUsageInput**](AwsMarketplaceMeteringBatchMeterUsageInput.md)                                                                                                                 | The raw request to call AWS metering service.                                                                                                                                                                            | [optional] 
 **AwsMeteringResponse**                        | Pointer to [**MarketplacemeteringBatchMeterUsageOutput**](MarketplacemeteringBatchMeterUsageOutput.md)                                                                                                                     | The raw response from AWS metering service.                                                                                                                                                                              | [optional] 
 **AzureMeteringRequest**                       | Pointer to [**AzureMarketplaceMeteringBatchUsageEvent**](AzureMarketplaceMeteringBatchUsageEvent.md)                                                                                                                       | The raw request to call Azure metering service.                                                                                                                                                                          | [optional] 
 **AzureMeteringResponse**                      | Pointer to [**GithubComSugerioMarketplaceServiceThirdPartyAzureSdkMarketplacemeteringv1BatchUsageEventOkResponse**](GithubComSugerioMarketplaceServiceThirdPartyAzureSdkMarketplacemeteringv1BatchUsageEventOkResponse.md) | The raw response from Azure metering service.                                                                                                                                                                            | [optional] 
 **CommitAmount**                               | Pointer to **float32**                                                                                                                                                                                                     | The amount of the commit if applicable.                                                                                                                                                                                  | [optional] 
 **CreditAmount**                               | Pointer to **float32**                                                                                                                                                                                                     | The amount of the credit if applicable.                                                                                                                                                                                  | [optional] 
 **CreditRecords**                              | Pointer to **map[string]float32**                                                                                                                                                                                          | The credit usage records in the map of &lt;DimensionKey, Count&gt; for usage metering API v1.                                                                                                                            | [optional] 
 **DecimalParts**                               | Pointer to **map[string]float32**                                                                                                                                                                                          | The decimal parts of the usage dimension quantity in the map of &lt;DimensionKey, DecimalPart&gt;, before this usage record report.                                                                                      | [optional] 
 **DimensionCategories**                        | Pointer to **map[string]string**                                                                                                                                                                                           | The categories of the usage records in the map of &lt;DimensionKey, Category&gt;. The dimension category is required when reporting usage records to Alibaba Marketplace. It comes from the metering dimension category. | [optional] 
 **DimensionUnitListPrice**                     | Pointer to **map[string]float32**                                                                                                                                                                                          | The public list price of each dimension in the map of &lt;DimensionKey, UnitPrice&gt;.                                                                                                                                   | [optional] 
 **DimensionUnitPrice**                         | Pointer to **map[string]float32**                                                                                                                                                                                          | The unit price of each dimension in the map of &lt;DimensionKey, UnitPrice&gt;. It can be the negotiated price in the private offer or the public list price.                                                            | [optional] 
 **EndTime**                                    | Pointer to **time.Time**                                                                                                                                                                                                   | time in UTC when the UsageRecordReport ends                                                                                                                                                                              | [optional] 
 **GcpMeteringRequest**                         | Pointer to [**GcpMarketplaceMeteringOperation**](GcpMarketplaceMeteringOperation.md)                                                                                                                                       | The raw request to call GCP metering service.                                                                                                                                                                            | [optional] 
 **GcpMeteringResponse**                        | Pointer to [**ServicecontrolReportResponse**](ServicecontrolReportResponse.md)                                                                                                                                             | The raw response from GCP metering service.                                                                                                                                                                              | [optional] 
 **IncludedRecords**                            | Pointer to **map[string]float32**                                                                                                                                                                                          | The included usage records in the map of &lt;DimensionKey, Count&gt; for usage metering API v1.                                                                                                                          | [optional] 
 **Message**                                    | Pointer to **string**                                                                                                                                                                                                      |                                                                                                                                                                                                                          | [optional] 
 **NewDecimalParts**                            | Pointer to **map[string]float32**                                                                                                                                                                                          | The decimal parts of the usage dimension quantity in the map of &lt;DimensionKey, DecimalPart&gt;, after this usage record report.                                                                                       | [optional] 
 **Partner**                                    | Pointer to **string**                                                                                                                                                                                                      | The partner where this usage record report is sent to. Such as AWS, AZURE or GCP.                                                                                                                                        | [optional] 
 **RecordsToReportBeforeAdjustmentAtListPrice** | Pointer to **map[string]float32**                                                                                                                                                                                          | The usage records to report before the adjustment by the commit with additional usage at list price, in the map of &lt;DimensionKey, Count&gt;.                                                                          | [optional] 
 **ReportedRecords**                            | Pointer to **map[string]float32**                                                                                                                                                                                          | The reported usage records in the map of &lt;DimensionKey, Count&gt; for usage metering API v1.                                                                                                                          | [optional] 
 **StartTime**                                  | Pointer to **time.Time**                                                                                                                                                                                                   | time in UTC when the UsageRecordReport starts                                                                                                                                                                            | [optional] 
 **Status**                                     | Pointer to [**UsageRecordReportStatus**](UsageRecordReportStatus.md)                                                                                                                                                       |                                                                                                                                                                                                                          | [optional] 
 **UsageRecordGroupIds**                        | Pointer to **[]string**                                                                                                                                                                                                    | The IDs of UsageRecordGroups aggregated in this UsageRecordReport.                                                                                                                                                       | [optional] 
 **UsedCommitAmount**                           | Pointer to **float32**                                                                                                                                                                                                     | The amount of the used commit before this usage record report if applicable.                                                                                                                                             | [optional] 
 **UsedCommitAmountIncrement**                  | Pointer to **float32**                                                                                                                                                                                                     | The amount of the used commit increment in this usage record report if applicable.                                                                                                                                       | [optional] 
 **UsedCreditAmount**                           | Pointer to **float32**                                                                                                                                                                                                     | The amount of the used credit before this usage record report if applicable.                                                                                                                                             | [optional] 
 **UsedCreditAmountIncrement**                  | Pointer to **float32**                                                                                                                                                                                                     | The amount of the used credit increment in this usage record report if applicable.                                                                                                                                       | [optional] 

## Methods

### NewMeteringUsageRecordReportInfo

`func NewMeteringUsageRecordReportInfo() *MeteringUsageRecordReportInfo`

NewMeteringUsageRecordReportInfo instantiates a new MeteringUsageRecordReportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordReportInfoWithDefaults

`func NewMeteringUsageRecordReportInfoWithDefaults() *MeteringUsageRecordReportInfo`

NewMeteringUsageRecordReportInfoWithDefaults instantiates a new MeteringUsageRecordReportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatedBillableRecords

`func (o *MeteringUsageRecordReportInfo) GetAggregatedBillableRecords() []AggregatedMeteringUsageRecord`

GetAggregatedBillableRecords returns the AggregatedBillableRecords field if non-nil, zero value otherwise.

### GetAggregatedBillableRecordsOk

`func (o *MeteringUsageRecordReportInfo) GetAggregatedBillableRecordsOk() (*[]AggregatedMeteringUsageRecord, bool)`

GetAggregatedBillableRecordsOk returns a tuple with the AggregatedBillableRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedBillableRecords

`func (o *MeteringUsageRecordReportInfo) SetAggregatedBillableRecords(v []AggregatedMeteringUsageRecord)`

SetAggregatedBillableRecords sets AggregatedBillableRecords field to given value.

### HasAggregatedBillableRecords

`func (o *MeteringUsageRecordReportInfo) HasAggregatedBillableRecords() bool`

HasAggregatedBillableRecords returns a boolean if a field has been set.

### GetAlibabaMeteringRequest

`func (o *MeteringUsageRecordReportInfo) GetAlibabaMeteringRequest() ClientPushMeteringDataRequest`

GetAlibabaMeteringRequest returns the AlibabaMeteringRequest field if non-nil, zero value otherwise.

### GetAlibabaMeteringRequestOk

`func (o *MeteringUsageRecordReportInfo) GetAlibabaMeteringRequestOk() (*ClientPushMeteringDataRequest, bool)`

GetAlibabaMeteringRequestOk returns a tuple with the AlibabaMeteringRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaMeteringRequest

`func (o *MeteringUsageRecordReportInfo) SetAlibabaMeteringRequest(v ClientPushMeteringDataRequest)`

SetAlibabaMeteringRequest sets AlibabaMeteringRequest field to given value.

### HasAlibabaMeteringRequest

`func (o *MeteringUsageRecordReportInfo) HasAlibabaMeteringRequest() bool`

HasAlibabaMeteringRequest returns a boolean if a field has been set.

### GetAlibabaMeteringResponse

`func (o *MeteringUsageRecordReportInfo) GetAlibabaMeteringResponse() ClientPushMeteringDataResponseBody`

GetAlibabaMeteringResponse returns the AlibabaMeteringResponse field if non-nil, zero value otherwise.

### GetAlibabaMeteringResponseOk

`func (o *MeteringUsageRecordReportInfo) GetAlibabaMeteringResponseOk() (*ClientPushMeteringDataResponseBody, bool)`

GetAlibabaMeteringResponseOk returns a tuple with the AlibabaMeteringResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaMeteringResponse

`func (o *MeteringUsageRecordReportInfo) SetAlibabaMeteringResponse(v ClientPushMeteringDataResponseBody)`

SetAlibabaMeteringResponse sets AlibabaMeteringResponse field to given value.

### HasAlibabaMeteringResponse

`func (o *MeteringUsageRecordReportInfo) HasAlibabaMeteringResponse() bool`

HasAlibabaMeteringResponse returns a boolean if a field has been set.

### GetAwsMeteringRequest

`func (o *MeteringUsageRecordReportInfo) GetAwsMeteringRequest() AwsMarketplaceMeteringBatchMeterUsageInput`

GetAwsMeteringRequest returns the AwsMeteringRequest field if non-nil, zero value otherwise.

### GetAwsMeteringRequestOk

`func (o *MeteringUsageRecordReportInfo) GetAwsMeteringRequestOk() (*AwsMarketplaceMeteringBatchMeterUsageInput, bool)`

GetAwsMeteringRequestOk returns a tuple with the AwsMeteringRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMeteringRequest

`func (o *MeteringUsageRecordReportInfo) SetAwsMeteringRequest(v AwsMarketplaceMeteringBatchMeterUsageInput)`

SetAwsMeteringRequest sets AwsMeteringRequest field to given value.

### HasAwsMeteringRequest

`func (o *MeteringUsageRecordReportInfo) HasAwsMeteringRequest() bool`

HasAwsMeteringRequest returns a boolean if a field has been set.

### GetAwsMeteringResponse

`func (o *MeteringUsageRecordReportInfo) GetAwsMeteringResponse() MarketplacemeteringBatchMeterUsageOutput`

GetAwsMeteringResponse returns the AwsMeteringResponse field if non-nil, zero value otherwise.

### GetAwsMeteringResponseOk

`func (o *MeteringUsageRecordReportInfo) GetAwsMeteringResponseOk() (*MarketplacemeteringBatchMeterUsageOutput, bool)`

GetAwsMeteringResponseOk returns a tuple with the AwsMeteringResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMeteringResponse

`func (o *MeteringUsageRecordReportInfo) SetAwsMeteringResponse(v MarketplacemeteringBatchMeterUsageOutput)`

SetAwsMeteringResponse sets AwsMeteringResponse field to given value.

### HasAwsMeteringResponse

`func (o *MeteringUsageRecordReportInfo) HasAwsMeteringResponse() bool`

HasAwsMeteringResponse returns a boolean if a field has been set.

### GetAzureMeteringRequest

`func (o *MeteringUsageRecordReportInfo) GetAzureMeteringRequest() AzureMarketplaceMeteringBatchUsageEvent`

GetAzureMeteringRequest returns the AzureMeteringRequest field if non-nil, zero value otherwise.

### GetAzureMeteringRequestOk

`func (o *MeteringUsageRecordReportInfo) GetAzureMeteringRequestOk() (*AzureMarketplaceMeteringBatchUsageEvent, bool)`

GetAzureMeteringRequestOk returns a tuple with the AzureMeteringRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureMeteringRequest

`func (o *MeteringUsageRecordReportInfo) SetAzureMeteringRequest(v AzureMarketplaceMeteringBatchUsageEvent)`

SetAzureMeteringRequest sets AzureMeteringRequest field to given value.

### HasAzureMeteringRequest

`func (o *MeteringUsageRecordReportInfo) HasAzureMeteringRequest() bool`

HasAzureMeteringRequest returns a boolean if a field has been set.

### GetAzureMeteringResponse

`func (o *MeteringUsageRecordReportInfo) GetAzureMeteringResponse() GithubComSugerioMarketplaceServiceThirdPartyAzureSdkMarketplacemeteringv1BatchUsageEventOkResponse`

GetAzureMeteringResponse returns the AzureMeteringResponse field if non-nil, zero value otherwise.

### GetAzureMeteringResponseOk

`func (o *MeteringUsageRecordReportInfo) GetAzureMeteringResponseOk() (*GithubComSugerioMarketplaceServiceThirdPartyAzureSdkMarketplacemeteringv1BatchUsageEventOkResponse, bool)`

GetAzureMeteringResponseOk returns a tuple with the AzureMeteringResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureMeteringResponse

`func (o *MeteringUsageRecordReportInfo) SetAzureMeteringResponse(v GithubComSugerioMarketplaceServiceThirdPartyAzureSdkMarketplacemeteringv1BatchUsageEventOkResponse)`

SetAzureMeteringResponse sets AzureMeteringResponse field to given value.

### HasAzureMeteringResponse

`func (o *MeteringUsageRecordReportInfo) HasAzureMeteringResponse() bool`

HasAzureMeteringResponse returns a boolean if a field has been set.

### GetCommitAmount

`func (o *MeteringUsageRecordReportInfo) GetCommitAmount() float32`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *MeteringUsageRecordReportInfo) GetCommitAmountOk() (*float32, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *MeteringUsageRecordReportInfo) SetCommitAmount(v float32)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *MeteringUsageRecordReportInfo) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *MeteringUsageRecordReportInfo) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *MeteringUsageRecordReportInfo) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *MeteringUsageRecordReportInfo) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *MeteringUsageRecordReportInfo) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCreditRecords

`func (o *MeteringUsageRecordReportInfo) GetCreditRecords() map[string]float32`

GetCreditRecords returns the CreditRecords field if non-nil, zero value otherwise.

### GetCreditRecordsOk

`func (o *MeteringUsageRecordReportInfo) GetCreditRecordsOk() (*map[string]float32, bool)`

GetCreditRecordsOk returns a tuple with the CreditRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRecords

`func (o *MeteringUsageRecordReportInfo) SetCreditRecords(v map[string]float32)`

SetCreditRecords sets CreditRecords field to given value.

### HasCreditRecords

`func (o *MeteringUsageRecordReportInfo) HasCreditRecords() bool`

HasCreditRecords returns a boolean if a field has been set.

### GetDecimalParts

`func (o *MeteringUsageRecordReportInfo) GetDecimalParts() map[string]float32`

GetDecimalParts returns the DecimalParts field if non-nil, zero value otherwise.

### GetDecimalPartsOk

`func (o *MeteringUsageRecordReportInfo) GetDecimalPartsOk() (*map[string]float32, bool)`

GetDecimalPartsOk returns a tuple with the DecimalParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalParts

`func (o *MeteringUsageRecordReportInfo) SetDecimalParts(v map[string]float32)`

SetDecimalParts sets DecimalParts field to given value.

### HasDecimalParts

`func (o *MeteringUsageRecordReportInfo) HasDecimalParts() bool`

HasDecimalParts returns a boolean if a field has been set.

### GetDimensionCategories

`func (o *MeteringUsageRecordReportInfo) GetDimensionCategories() map[string]string`

GetDimensionCategories returns the DimensionCategories field if non-nil, zero value otherwise.

### GetDimensionCategoriesOk

`func (o *MeteringUsageRecordReportInfo) GetDimensionCategoriesOk() (*map[string]string, bool)`

GetDimensionCategoriesOk returns a tuple with the DimensionCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionCategories

`func (o *MeteringUsageRecordReportInfo) SetDimensionCategories(v map[string]string)`

SetDimensionCategories sets DimensionCategories field to given value.

### HasDimensionCategories

`func (o *MeteringUsageRecordReportInfo) HasDimensionCategories() bool`

HasDimensionCategories returns a boolean if a field has been set.

### GetDimensionUnitListPrice

`func (o *MeteringUsageRecordReportInfo) GetDimensionUnitListPrice() map[string]float32`

GetDimensionUnitListPrice returns the DimensionUnitListPrice field if non-nil, zero value otherwise.

### GetDimensionUnitListPriceOk

`func (o *MeteringUsageRecordReportInfo) GetDimensionUnitListPriceOk() (*map[string]float32, bool)`

GetDimensionUnitListPriceOk returns a tuple with the DimensionUnitListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnitListPrice

`func (o *MeteringUsageRecordReportInfo) SetDimensionUnitListPrice(v map[string]float32)`

SetDimensionUnitListPrice sets DimensionUnitListPrice field to given value.

### HasDimensionUnitListPrice

`func (o *MeteringUsageRecordReportInfo) HasDimensionUnitListPrice() bool`

HasDimensionUnitListPrice returns a boolean if a field has been set.

### GetDimensionUnitPrice

`func (o *MeteringUsageRecordReportInfo) GetDimensionUnitPrice() map[string]float32`

GetDimensionUnitPrice returns the DimensionUnitPrice field if non-nil, zero value otherwise.

### GetDimensionUnitPriceOk

`func (o *MeteringUsageRecordReportInfo) GetDimensionUnitPriceOk() (*map[string]float32, bool)`

GetDimensionUnitPriceOk returns a tuple with the DimensionUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnitPrice

`func (o *MeteringUsageRecordReportInfo) SetDimensionUnitPrice(v map[string]float32)`

SetDimensionUnitPrice sets DimensionUnitPrice field to given value.

### HasDimensionUnitPrice

`func (o *MeteringUsageRecordReportInfo) HasDimensionUnitPrice() bool`

HasDimensionUnitPrice returns a boolean if a field has been set.

### GetEndTime

`func (o *MeteringUsageRecordReportInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MeteringUsageRecordReportInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MeteringUsageRecordReportInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MeteringUsageRecordReportInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetGcpMeteringRequest

`func (o *MeteringUsageRecordReportInfo) GetGcpMeteringRequest() GcpMarketplaceMeteringOperation`

GetGcpMeteringRequest returns the GcpMeteringRequest field if non-nil, zero value otherwise.

### GetGcpMeteringRequestOk

`func (o *MeteringUsageRecordReportInfo) GetGcpMeteringRequestOk() (*GcpMarketplaceMeteringOperation, bool)`

GetGcpMeteringRequestOk returns a tuple with the GcpMeteringRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpMeteringRequest

`func (o *MeteringUsageRecordReportInfo) SetGcpMeteringRequest(v GcpMarketplaceMeteringOperation)`

SetGcpMeteringRequest sets GcpMeteringRequest field to given value.

### HasGcpMeteringRequest

`func (o *MeteringUsageRecordReportInfo) HasGcpMeteringRequest() bool`

HasGcpMeteringRequest returns a boolean if a field has been set.

### GetGcpMeteringResponse

`func (o *MeteringUsageRecordReportInfo) GetGcpMeteringResponse() ServicecontrolReportResponse`

GetGcpMeteringResponse returns the GcpMeteringResponse field if non-nil, zero value otherwise.

### GetGcpMeteringResponseOk

`func (o *MeteringUsageRecordReportInfo) GetGcpMeteringResponseOk() (*ServicecontrolReportResponse, bool)`

GetGcpMeteringResponseOk returns a tuple with the GcpMeteringResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpMeteringResponse

`func (o *MeteringUsageRecordReportInfo) SetGcpMeteringResponse(v ServicecontrolReportResponse)`

SetGcpMeteringResponse sets GcpMeteringResponse field to given value.

### HasGcpMeteringResponse

`func (o *MeteringUsageRecordReportInfo) HasGcpMeteringResponse() bool`

HasGcpMeteringResponse returns a boolean if a field has been set.

### GetIncludedRecords

`func (o *MeteringUsageRecordReportInfo) GetIncludedRecords() map[string]float32`

GetIncludedRecords returns the IncludedRecords field if non-nil, zero value otherwise.

### GetIncludedRecordsOk

`func (o *MeteringUsageRecordReportInfo) GetIncludedRecordsOk() (*map[string]float32, bool)`

GetIncludedRecordsOk returns a tuple with the IncludedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedRecords

`func (o *MeteringUsageRecordReportInfo) SetIncludedRecords(v map[string]float32)`

SetIncludedRecords sets IncludedRecords field to given value.

### HasIncludedRecords

`func (o *MeteringUsageRecordReportInfo) HasIncludedRecords() bool`

HasIncludedRecords returns a boolean if a field has been set.

### GetMessage

`func (o *MeteringUsageRecordReportInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MeteringUsageRecordReportInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MeteringUsageRecordReportInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MeteringUsageRecordReportInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNewDecimalParts

`func (o *MeteringUsageRecordReportInfo) GetNewDecimalParts() map[string]float32`

GetNewDecimalParts returns the NewDecimalParts field if non-nil, zero value otherwise.

### GetNewDecimalPartsOk

`func (o *MeteringUsageRecordReportInfo) GetNewDecimalPartsOk() (*map[string]float32, bool)`

GetNewDecimalPartsOk returns a tuple with the NewDecimalParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDecimalParts

`func (o *MeteringUsageRecordReportInfo) SetNewDecimalParts(v map[string]float32)`

SetNewDecimalParts sets NewDecimalParts field to given value.

### HasNewDecimalParts

`func (o *MeteringUsageRecordReportInfo) HasNewDecimalParts() bool`

HasNewDecimalParts returns a boolean if a field has been set.

### GetPartner

`func (o *MeteringUsageRecordReportInfo) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *MeteringUsageRecordReportInfo) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *MeteringUsageRecordReportInfo) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *MeteringUsageRecordReportInfo) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetRecordsToReportBeforeAdjustmentAtListPrice

`func (o *MeteringUsageRecordReportInfo) GetRecordsToReportBeforeAdjustmentAtListPrice() map[string]float32`

GetRecordsToReportBeforeAdjustmentAtListPrice returns the RecordsToReportBeforeAdjustmentAtListPrice field if non-nil, zero value otherwise.

### GetRecordsToReportBeforeAdjustmentAtListPriceOk

`func (o *MeteringUsageRecordReportInfo) GetRecordsToReportBeforeAdjustmentAtListPriceOk() (*map[string]float32, bool)`

GetRecordsToReportBeforeAdjustmentAtListPriceOk returns a tuple with the RecordsToReportBeforeAdjustmentAtListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsToReportBeforeAdjustmentAtListPrice

`func (o *MeteringUsageRecordReportInfo) SetRecordsToReportBeforeAdjustmentAtListPrice(v map[string]float32)`

SetRecordsToReportBeforeAdjustmentAtListPrice sets RecordsToReportBeforeAdjustmentAtListPrice field to given value.

### HasRecordsToReportBeforeAdjustmentAtListPrice

`func (o *MeteringUsageRecordReportInfo) HasRecordsToReportBeforeAdjustmentAtListPrice() bool`

HasRecordsToReportBeforeAdjustmentAtListPrice returns a boolean if a field has been set.

### GetReportedRecords

`func (o *MeteringUsageRecordReportInfo) GetReportedRecords() map[string]float32`

GetReportedRecords returns the ReportedRecords field if non-nil, zero value otherwise.

### GetReportedRecordsOk

`func (o *MeteringUsageRecordReportInfo) GetReportedRecordsOk() (*map[string]float32, bool)`

GetReportedRecordsOk returns a tuple with the ReportedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedRecords

`func (o *MeteringUsageRecordReportInfo) SetReportedRecords(v map[string]float32)`

SetReportedRecords sets ReportedRecords field to given value.

### HasReportedRecords

`func (o *MeteringUsageRecordReportInfo) HasReportedRecords() bool`

HasReportedRecords returns a boolean if a field has been set.

### GetStartTime

`func (o *MeteringUsageRecordReportInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MeteringUsageRecordReportInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MeteringUsageRecordReportInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MeteringUsageRecordReportInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *MeteringUsageRecordReportInfo) GetStatus() UsageRecordReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeteringUsageRecordReportInfo) GetStatusOk() (*UsageRecordReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeteringUsageRecordReportInfo) SetStatus(v UsageRecordReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeteringUsageRecordReportInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsageRecordGroupIds

`func (o *MeteringUsageRecordReportInfo) GetUsageRecordGroupIds() []string`

GetUsageRecordGroupIds returns the UsageRecordGroupIds field if non-nil, zero value otherwise.

### GetUsageRecordGroupIdsOk

`func (o *MeteringUsageRecordReportInfo) GetUsageRecordGroupIdsOk() (*[]string, bool)`

GetUsageRecordGroupIdsOk returns a tuple with the UsageRecordGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordGroupIds

`func (o *MeteringUsageRecordReportInfo) SetUsageRecordGroupIds(v []string)`

SetUsageRecordGroupIds sets UsageRecordGroupIds field to given value.

### HasUsageRecordGroupIds

`func (o *MeteringUsageRecordReportInfo) HasUsageRecordGroupIds() bool`

HasUsageRecordGroupIds returns a boolean if a field has been set.

### GetUsedCommitAmount

`func (o *MeteringUsageRecordReportInfo) GetUsedCommitAmount() float32`

GetUsedCommitAmount returns the UsedCommitAmount field if non-nil, zero value otherwise.

### GetUsedCommitAmountOk

`func (o *MeteringUsageRecordReportInfo) GetUsedCommitAmountOk() (*float32, bool)`

GetUsedCommitAmountOk returns a tuple with the UsedCommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCommitAmount

`func (o *MeteringUsageRecordReportInfo) SetUsedCommitAmount(v float32)`

SetUsedCommitAmount sets UsedCommitAmount field to given value.

### HasUsedCommitAmount

`func (o *MeteringUsageRecordReportInfo) HasUsedCommitAmount() bool`

HasUsedCommitAmount returns a boolean if a field has been set.

### GetUsedCommitAmountIncrement

`func (o *MeteringUsageRecordReportInfo) GetUsedCommitAmountIncrement() float32`

GetUsedCommitAmountIncrement returns the UsedCommitAmountIncrement field if non-nil, zero value otherwise.

### GetUsedCommitAmountIncrementOk

`func (o *MeteringUsageRecordReportInfo) GetUsedCommitAmountIncrementOk() (*float32, bool)`

GetUsedCommitAmountIncrementOk returns a tuple with the UsedCommitAmountIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCommitAmountIncrement

`func (o *MeteringUsageRecordReportInfo) SetUsedCommitAmountIncrement(v float32)`

SetUsedCommitAmountIncrement sets UsedCommitAmountIncrement field to given value.

### HasUsedCommitAmountIncrement

`func (o *MeteringUsageRecordReportInfo) HasUsedCommitAmountIncrement() bool`

HasUsedCommitAmountIncrement returns a boolean if a field has been set.

### GetUsedCreditAmount

`func (o *MeteringUsageRecordReportInfo) GetUsedCreditAmount() float32`

GetUsedCreditAmount returns the UsedCreditAmount field if non-nil, zero value otherwise.

### GetUsedCreditAmountOk

`func (o *MeteringUsageRecordReportInfo) GetUsedCreditAmountOk() (*float32, bool)`

GetUsedCreditAmountOk returns a tuple with the UsedCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCreditAmount

`func (o *MeteringUsageRecordReportInfo) SetUsedCreditAmount(v float32)`

SetUsedCreditAmount sets UsedCreditAmount field to given value.

### HasUsedCreditAmount

`func (o *MeteringUsageRecordReportInfo) HasUsedCreditAmount() bool`

HasUsedCreditAmount returns a boolean if a field has been set.

### GetUsedCreditAmountIncrement

`func (o *MeteringUsageRecordReportInfo) GetUsedCreditAmountIncrement() float32`

GetUsedCreditAmountIncrement returns the UsedCreditAmountIncrement field if non-nil, zero value otherwise.

### GetUsedCreditAmountIncrementOk

`func (o *MeteringUsageRecordReportInfo) GetUsedCreditAmountIncrementOk() (*float32, bool)`

GetUsedCreditAmountIncrementOk returns a tuple with the UsedCreditAmountIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCreditAmountIncrement

`func (o *MeteringUsageRecordReportInfo) SetUsedCreditAmountIncrement(v float32)`

SetUsedCreditAmountIncrement sets UsedCreditAmountIncrement field to given value.

### HasUsedCreditAmountIncrement

`func (o *MeteringUsageRecordReportInfo) HasUsedCreditAmountIncrement() bool`

HasUsedCreditAmountIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


