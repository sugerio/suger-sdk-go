# GithubComSugerioMarketplaceServicePkgOrmProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | CreatedBy holds the value of the \&quot;created_by\&quot; field. | [optional] 
**CreationTime** | Pointer to **time.Time** | CreationTime holds the value of the \&quot;creation_time\&quot; field. | [optional] 
**ExternalID** | Pointer to **string** | ExternalID holds the value of the \&quot;external_id\&quot; field. | [optional] 
**FulfillmentUrl** | Pointer to **string** | FulfillmentURL holds the value of the \&quot;fulfillment_url\&quot; field. | [optional] 
**Id** | Pointer to **string** | ID of the ent. | [optional] 
**Info** | Pointer to [**ProductInfo**](ProductInfo.md) | Info holds the value of the \&quot;info\&quot; field. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | LastUpdateTime holds the value of the \&quot;last_update_time\&quot; field. | [optional] 
**LastUpdatedBy** | Pointer to **string** | LastUpdatedBy holds the value of the \&quot;last_updated_by\&quot; field. | [optional] 
**MetaInfo** | Pointer to [**WorkloadMetaInfo**](WorkloadMetaInfo.md) | MetaInfo holds the value of the \&quot;meta_info\&quot; field. | [optional] 
**Name** | Pointer to **string** | Name holds the value of the \&quot;name\&quot; field. | [optional] 
**OrganizationID** | Pointer to **string** | OrganizationID holds the value of the \&quot;organization_id\&quot; field. | [optional] 
**Partner** | Pointer to **string** | Partner holds the value of the \&quot;partner\&quot; field. | [optional] 
**PartnerID** | Pointer to **string** | PartnerID holds the value of the \&quot;partner_id\&quot; field. | [optional] 
**ProductType** | Pointer to [**GithubComSugerioMarketplaceServicePkgOrmProductProductType**](GithubComSugerioMarketplaceServicePkgOrmProductProductType.md) | ProductType holds the value of the \&quot;product_type\&quot; field. | [optional] 
**SellerID** | Pointer to **string** | SellerID holds the value of the \&quot;seller_id\&quot; field. | [optional] 
**Service** | Pointer to **string** | Service holds the value of the \&quot;service\&quot; field. | [optional] 
**Status** | Pointer to **string** | Status holds the value of the \&quot;status\&quot; field. | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgOrmProduct

`func NewGithubComSugerioMarketplaceServicePkgOrmProduct() *GithubComSugerioMarketplaceServicePkgOrmProduct`

NewGithubComSugerioMarketplaceServicePkgOrmProduct instantiates a new GithubComSugerioMarketplaceServicePkgOrmProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgOrmProductWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgOrmProductWithDefaults() *GithubComSugerioMarketplaceServicePkgOrmProduct`

NewGithubComSugerioMarketplaceServicePkgOrmProductWithDefaults instantiates a new GithubComSugerioMarketplaceServicePkgOrmProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetFulfillmentUrl

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetFulfillmentUrl() string`

GetFulfillmentUrl returns the FulfillmentUrl field if non-nil, zero value otherwise.

### GetFulfillmentUrlOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetFulfillmentUrlOk() (*string, bool)`

GetFulfillmentUrlOk returns a tuple with the FulfillmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentUrl

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetFulfillmentUrl(v string)`

SetFulfillmentUrl sets FulfillmentUrl field to given value.

### HasFulfillmentUrl

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasFulfillmentUrl() bool`

HasFulfillmentUrl returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetInfo() ProductInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetInfoOk() (*ProductInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetInfo(v ProductInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetMetaInfo() WorkloadMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetMetaInfoOk() (*WorkloadMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetMetaInfo(v WorkloadMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetPartnerID() string`

GetPartnerID returns the PartnerID field if non-nil, zero value otherwise.

### GetPartnerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetPartnerIDOk() (*string, bool)`

GetPartnerIDOk returns a tuple with the PartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetPartnerID(v string)`

SetPartnerID sets PartnerID field to given value.

### HasPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasPartnerID() bool`

HasPartnerID returns a boolean if a field has been set.

### GetProductType

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetProductType() GithubComSugerioMarketplaceServicePkgOrmProductProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetProductTypeOk() (*GithubComSugerioMarketplaceServicePkgOrmProductProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetProductType(v GithubComSugerioMarketplaceServicePkgOrmProductProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSellerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetSellerID() string`

GetSellerID returns the SellerID field if non-nil, zero value otherwise.

### GetSellerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetSellerIDOk() (*string, bool)`

GetSellerIDOk returns a tuple with the SellerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetSellerID(v string)`

SetSellerID sets SellerID field to given value.

### HasSellerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasSellerID() bool`

HasSellerID returns a boolean if a field has been set.

### GetService

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmProduct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


