# MeteringUsageRecordReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**EntitlementTermID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**MeteringUsageRecordReportInfo**](MeteringUsageRecordReportInfo.md) |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to **string** |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 

## Methods

### NewMeteringUsageRecordReport

`func NewMeteringUsageRecordReport() *MeteringUsageRecordReport`

NewMeteringUsageRecordReport instantiates a new MeteringUsageRecordReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordReportWithDefaults

`func NewMeteringUsageRecordReportWithDefaults() *MeteringUsageRecordReport`

NewMeteringUsageRecordReportWithDefaults instantiates a new MeteringUsageRecordReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *MeteringUsageRecordReport) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *MeteringUsageRecordReport) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *MeteringUsageRecordReport) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *MeteringUsageRecordReport) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *MeteringUsageRecordReport) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MeteringUsageRecordReport) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MeteringUsageRecordReport) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MeteringUsageRecordReport) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEntitlementID

`func (o *MeteringUsageRecordReport) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *MeteringUsageRecordReport) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *MeteringUsageRecordReport) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *MeteringUsageRecordReport) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetEntitlementTermID

`func (o *MeteringUsageRecordReport) GetEntitlementTermID() string`

GetEntitlementTermID returns the EntitlementTermID field if non-nil, zero value otherwise.

### GetEntitlementTermIDOk

`func (o *MeteringUsageRecordReport) GetEntitlementTermIDOk() (*string, bool)`

GetEntitlementTermIDOk returns a tuple with the EntitlementTermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTermID

`func (o *MeteringUsageRecordReport) SetEntitlementTermID(v string)`

SetEntitlementTermID sets EntitlementTermID field to given value.

### HasEntitlementTermID

`func (o *MeteringUsageRecordReport) HasEntitlementTermID() bool`

HasEntitlementTermID returns a boolean if a field has been set.

### GetId

`func (o *MeteringUsageRecordReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeteringUsageRecordReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeteringUsageRecordReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeteringUsageRecordReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *MeteringUsageRecordReport) GetInfo() MeteringUsageRecordReportInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MeteringUsageRecordReport) GetInfoOk() (*MeteringUsageRecordReportInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MeteringUsageRecordReport) SetInfo(v MeteringUsageRecordReportInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MeteringUsageRecordReport) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOrganizationID

`func (o *MeteringUsageRecordReport) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *MeteringUsageRecordReport) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *MeteringUsageRecordReport) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *MeteringUsageRecordReport) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *MeteringUsageRecordReport) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *MeteringUsageRecordReport) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *MeteringUsageRecordReport) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *MeteringUsageRecordReport) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetProductID

`func (o *MeteringUsageRecordReport) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *MeteringUsageRecordReport) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *MeteringUsageRecordReport) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *MeteringUsageRecordReport) HasProductID() bool`

HasProductID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


