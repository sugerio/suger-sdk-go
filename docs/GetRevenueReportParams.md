# GetRevenueReportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** | Optional, if available, return the report for the Buyer. | [optional] 
**EntitlementID** | Pointer to **string** | Optional, if available, return the report for the Entitlement. | [optional] 
**OrganizationID** | **string** | Required. If the productID &amp; entitlementID are emtpy, return the report for the entire Organization. | 
**Partner** | **string** |  | 
**ProductID** | Pointer to **string** | Optional, if available, return the report for the Product. | [optional] 
**ReportType** | [**RevenueReportType**](RevenueReportType.md) |  | 
**Service** | **string** |  | 

## Methods

### NewGetRevenueReportParams

`func NewGetRevenueReportParams(organizationID string, partner string, reportType RevenueReportType, service string, ) *GetRevenueReportParams`

NewGetRevenueReportParams instantiates a new GetRevenueReportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevenueReportParamsWithDefaults

`func NewGetRevenueReportParamsWithDefaults() *GetRevenueReportParams`

NewGetRevenueReportParamsWithDefaults instantiates a new GetRevenueReportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *GetRevenueReportParams) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *GetRevenueReportParams) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *GetRevenueReportParams) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *GetRevenueReportParams) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetEntitlementID

`func (o *GetRevenueReportParams) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *GetRevenueReportParams) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *GetRevenueReportParams) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *GetRevenueReportParams) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GetRevenueReportParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GetRevenueReportParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GetRevenueReportParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.


### GetPartner

`func (o *GetRevenueReportParams) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *GetRevenueReportParams) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *GetRevenueReportParams) SetPartner(v string)`

SetPartner sets Partner field to given value.


### GetProductID

`func (o *GetRevenueReportParams) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *GetRevenueReportParams) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *GetRevenueReportParams) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *GetRevenueReportParams) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetReportType

`func (o *GetRevenueReportParams) GetReportType() RevenueReportType`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *GetRevenueReportParams) GetReportTypeOk() (*RevenueReportType, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *GetRevenueReportParams) SetReportType(v RevenueReportType)`

SetReportType sets ReportType field to given value.


### GetService

`func (o *GetRevenueReportParams) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetRevenueReportParams) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetRevenueReportParams) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


