# GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** | BuyerID holds the value of the \&quot;buyer_id\&quot; field. | [optional] 
**CreationTime** | Pointer to **time.Time** | CreationTime holds the value of the \&quot;creation_time\&quot; field. | [optional] 
**EndTime** | Pointer to **time.Time** | EndTime holds the value of the \&quot;end_time\&quot; field. | [optional] 
**EntitlementTermID** | Pointer to **string** | EntitlementTermID holds the value of the \&quot;entitlement_term_id\&quot; field. | [optional] 
**ExternalBuyerID** | Pointer to **string** | ExternalBuyerID holds the value of the \&quot;external_buyer_id\&quot; field. | [optional] 
**ExternalID** | Pointer to **string** | ExternalID holds the value of the \&quot;external_id\&quot; field. | [optional] 
**ExternalProductID** | Pointer to **string** | ExternalProductID holds the value of the \&quot;external_product_id\&quot; field. | [optional] 
**Id** | Pointer to **string** | ID of the ent. | [optional] 
**Info** | Pointer to [**EntitlementInfo**](EntitlementInfo.md) | Info holds the value of the \&quot;info\&quot; field. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | LastUpdateTime holds the value of the \&quot;last_update_time\&quot; field. | [optional] 
**MetaInfo** | Pointer to [**WorkloadMetaInfo**](WorkloadMetaInfo.md) | MetaInfo holds the value of the \&quot;meta_info\&quot; field. | [optional] 
**Name** | Pointer to **string** | Name holds the value of the \&quot;name\&quot; field. | [optional] 
**OfferID** | Pointer to **string** | OfferID holds the value of the \&quot;offer_id\&quot; field. | [optional] 
**OrganizationID** | Pointer to **string** | OrganizationID holds the value of the \&quot;organization_id\&quot; field. | [optional] 
**Partner** | Pointer to **string** | Partner holds the value of the \&quot;partner\&quot; field. | [optional] 
**PartnerID** | Pointer to **string** | PartnerID holds the value of the \&quot;partner_id\&quot; field. | [optional] 
**ProductID** | Pointer to **string** | ProductID holds the value of the \&quot;product_id\&quot; field. | [optional] 
**Service** | Pointer to **string** | Service holds the value of the \&quot;service\&quot; field. | [optional] 
**StartTime** | Pointer to **time.Time** | StartTime holds the value of the \&quot;start_time\&quot; field. | [optional] 
**Status** | Pointer to **string** | Status holds the value of the \&quot;status\&quot; field. | [optional] 
**Type** | Pointer to **string** | Type holds the value of the \&quot;type\&quot; field. | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement

`func NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement() *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement`

NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement instantiates a new GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlementWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlementWithDefaults() *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement`

NewGithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlementWithDefaults instantiates a new GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntitlementTermID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetEntitlementTermID() string`

GetEntitlementTermID returns the EntitlementTermID field if non-nil, zero value otherwise.

### GetEntitlementTermIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetEntitlementTermIDOk() (*string, bool)`

GetEntitlementTermIDOk returns a tuple with the EntitlementTermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTermID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetEntitlementTermID(v string)`

SetEntitlementTermID sets EntitlementTermID field to given value.

### HasEntitlementTermID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasEntitlementTermID() bool`

HasEntitlementTermID returns a boolean if a field has been set.

### GetExternalBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalBuyerID() string`

GetExternalBuyerID returns the ExternalBuyerID field if non-nil, zero value otherwise.

### GetExternalBuyerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalBuyerIDOk() (*string, bool)`

GetExternalBuyerIDOk returns a tuple with the ExternalBuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetExternalBuyerID(v string)`

SetExternalBuyerID sets ExternalBuyerID field to given value.

### HasExternalBuyerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasExternalBuyerID() bool`

HasExternalBuyerID returns a boolean if a field has been set.

### GetExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetExternalProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalProductID() string`

GetExternalProductID returns the ExternalProductID field if non-nil, zero value otherwise.

### GetExternalProductIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetExternalProductIDOk() (*string, bool)`

GetExternalProductIDOk returns a tuple with the ExternalProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetExternalProductID(v string)`

SetExternalProductID sets ExternalProductID field to given value.

### HasExternalProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasExternalProductID() bool`

HasExternalProductID returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetInfo() EntitlementInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetInfoOk() (*EntitlementInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetInfo(v EntitlementInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetMetaInfo() WorkloadMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetMetaInfoOk() (*WorkloadMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetMetaInfo(v WorkloadMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetOfferID() string`

GetOfferID returns the OfferID field if non-nil, zero value otherwise.

### GetOfferIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetOfferIDOk() (*string, bool)`

GetOfferIDOk returns a tuple with the OfferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetOfferID(v string)`

SetOfferID sets OfferID field to given value.

### HasOfferID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasOfferID() bool`

HasOfferID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetPartnerID() string`

GetPartnerID returns the PartnerID field if non-nil, zero value otherwise.

### GetPartnerIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetPartnerIDOk() (*string, bool)`

GetPartnerIDOk returns a tuple with the PartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetPartnerID(v string)`

SetPartnerID sets PartnerID field to given value.

### HasPartnerID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasPartnerID() bool`

HasPartnerID returns a boolean if a field has been set.

### GetProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetService

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStartTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GithubComSugerioMarketplaceServicePkgOrmWorkloadEntitlement) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


