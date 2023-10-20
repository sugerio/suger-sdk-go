# RevenueReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to **string** |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 
**ReportDate** | Pointer to **time.Time** |  | [optional] 
**ReportType** | Pointer to [**RevenueReportType**](RevenueReportType.md) |  | [optional] 
**RevenueRecords** | Pointer to [**[]RevenueRecord**](RevenueRecord.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewRevenueReport

`func NewRevenueReport() *RevenueReport`

NewRevenueReport instantiates a new RevenueReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueReportWithDefaults

`func NewRevenueReportWithDefaults() *RevenueReport`

NewRevenueReportWithDefaults instantiates a new RevenueReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *RevenueReport) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *RevenueReport) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *RevenueReport) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *RevenueReport) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetEntitlementID

`func (o *RevenueReport) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *RevenueReport) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *RevenueReport) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *RevenueReport) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *RevenueReport) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *RevenueReport) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *RevenueReport) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *RevenueReport) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *RevenueReport) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *RevenueReport) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *RevenueReport) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *RevenueReport) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetProductID

`func (o *RevenueReport) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *RevenueReport) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *RevenueReport) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *RevenueReport) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetReportDate

`func (o *RevenueReport) GetReportDate() time.Time`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *RevenueReport) GetReportDateOk() (*time.Time, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *RevenueReport) SetReportDate(v time.Time)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *RevenueReport) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetReportType

`func (o *RevenueReport) GetReportType() RevenueReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *RevenueReport) GetReportTypeOk() (*RevenueReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *RevenueReport) SetReportType(v RevenueReportType)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *RevenueReport) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetRevenueRecords

`func (o *RevenueReport) GetRevenueRecords() []RevenueRecord`

GetRevenueRecords returns the RevenueRecords field if non-nil, zero value otherwise.

### GetRevenueRecordsOk

`func (o *RevenueReport) GetRevenueRecordsOk() (*[]RevenueRecord, bool)`

GetRevenueRecordsOk returns a tuple with the RevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueRecords

`func (o *RevenueReport) SetRevenueRecords(v []RevenueRecord)`

SetRevenueRecords sets RevenueRecords field to given value.

### HasRevenueRecords

`func (o *RevenueReport) HasRevenueRecords() bool`

HasRevenueRecords returns a boolean if a field has been set.

### GetService

`func (o *RevenueReport) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RevenueReport) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RevenueReport) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *RevenueReport) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


