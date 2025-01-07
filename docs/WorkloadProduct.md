# WorkloadProduct

## Properties

 Name               | Type                                                   | Description | Notes      
--------------------|--------------------------------------------------------|-------------|------------
 **CreatedBy**      | Pointer to **string**                                  |             | [optional] 
 **CreationTime**   | Pointer to **time.Time**                               |             | [optional] 
 **ExternalID**     | Pointer to **string**                                  |             | [optional] 
 **FulfillmentUrl** | Pointer to **string**                                  |             | [optional] 
 **Id**             | Pointer to **string**                                  |             | [optional] 
 **Info**           | Pointer to [**ProductInfo**](ProductInfo.md)           |             | [optional] 
 **LastUpdateTime** | Pointer to **time.Time**                               |             | [optional] 
 **LastUpdatedBy**  | Pointer to **string**                                  |             | [optional] 
 **MetaInfo**       | Pointer to [**WorkloadMetaInfo**](WorkloadMetaInfo.md) |             | [optional] 
 **Name**           | Pointer to **string**                                  |             | [optional] 
 **OrganizationID** | Pointer to **string**                                  |             | [optional] 
 **Partner**        | Pointer to [**Partner**](Partner.md)                   |             | [optional] 
 **PartnerID**      | Pointer to **string**                                  |             | [optional] 
 **ProductType**    | Pointer to **string**                                  |             | [optional] 
 **Service**        | Pointer to [**PartnerService**](PartnerService.md)     |             | [optional] 
 **Status**         | Pointer to **string**                                  |             | [optional] 

## Methods

### NewWorkloadProduct

`func NewWorkloadProduct() *WorkloadProduct`

NewWorkloadProduct instantiates a new WorkloadProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadProductWithDefaults

`func NewWorkloadProductWithDefaults() *WorkloadProduct`

NewWorkloadProductWithDefaults instantiates a new WorkloadProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *WorkloadProduct) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkloadProduct) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkloadProduct) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkloadProduct) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreationTime

`func (o *WorkloadProduct) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WorkloadProduct) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WorkloadProduct) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WorkloadProduct) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetExternalID

`func (o *WorkloadProduct) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *WorkloadProduct) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *WorkloadProduct) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *WorkloadProduct) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetFulfillmentUrl

`func (o *WorkloadProduct) GetFulfillmentUrl() string`

GetFulfillmentUrl returns the FulfillmentUrl field if non-nil, zero value otherwise.

### GetFulfillmentUrlOk

`func (o *WorkloadProduct) GetFulfillmentUrlOk() (*string, bool)`

GetFulfillmentUrlOk returns a tuple with the FulfillmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUrl

`func (o *WorkloadProduct) SetFulfillmentUrl(v string)`

SetFulfillmentUrl sets FulfillmentUrl field to given value.

### HasFulfillmentUrl

`func (o *WorkloadProduct) HasFulfillmentUrl() bool`

HasFulfillmentUrl returns a boolean if a field has been set.

### GetId

`func (o *WorkloadProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkloadProduct) GetInfo() ProductInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkloadProduct) GetInfoOk() (*ProductInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkloadProduct) SetInfo(v ProductInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkloadProduct) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *WorkloadProduct) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *WorkloadProduct) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *WorkloadProduct) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *WorkloadProduct) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *WorkloadProduct) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *WorkloadProduct) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *WorkloadProduct) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *WorkloadProduct) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetMetaInfo

`func (o *WorkloadProduct) GetMetaInfo() WorkloadMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *WorkloadProduct) GetMetaInfoOk() (*WorkloadMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *WorkloadProduct) SetMetaInfo(v WorkloadMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *WorkloadProduct) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *WorkloadProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationID

`func (o *WorkloadProduct) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *WorkloadProduct) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *WorkloadProduct) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *WorkloadProduct) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *WorkloadProduct) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *WorkloadProduct) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *WorkloadProduct) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *WorkloadProduct) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerID

`func (o *WorkloadProduct) GetPartnerID() string`

GetPartnerID returns the PartnerID field if non-nil, zero value otherwise.

### GetPartnerIDOk

`func (o *WorkloadProduct) GetPartnerIDOk() (*string, bool)`

GetPartnerIDOk returns a tuple with the PartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerID

`func (o *WorkloadProduct) SetPartnerID(v string)`

SetPartnerID sets PartnerID field to given value.

### HasPartnerID

`func (o *WorkloadProduct) HasPartnerID() bool`

HasPartnerID returns a boolean if a field has been set.

### GetProductType

`func (o *WorkloadProduct) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *WorkloadProduct) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *WorkloadProduct) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *WorkloadProduct) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetService

`func (o *WorkloadProduct) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkloadProduct) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkloadProduct) SetService(v PartnerService)`

SetService sets Service field to given value.

### HasService

`func (o *WorkloadProduct) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadProduct) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadProduct) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadProduct) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadProduct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


