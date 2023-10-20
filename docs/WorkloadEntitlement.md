# WorkloadEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** | nullable | [optional] 
**EntitlementTermID** | Pointer to **string** |  | [optional] 
**ExternalBuyerID** | Pointer to **string** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**ExternalProductID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**EntitlementInfo**](EntitlementInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**MetaInfo** | Pointer to [**WorkloadMetaInfo**](WorkloadMetaInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OfferID** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**PartnerService**](PartnerService.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**EntitlementStatus**](EntitlementStatus.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkloadEntitlement

`func NewWorkloadEntitlement() *WorkloadEntitlement`

NewWorkloadEntitlement instantiates a new WorkloadEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEntitlementWithDefaults

`func NewWorkloadEntitlementWithDefaults() *WorkloadEntitlement`

NewWorkloadEntitlementWithDefaults instantiates a new WorkloadEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *WorkloadEntitlement) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *WorkloadEntitlement) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *WorkloadEntitlement) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *WorkloadEntitlement) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCreationTime

`func (o *WorkloadEntitlement) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WorkloadEntitlement) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WorkloadEntitlement) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WorkloadEntitlement) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkloadEntitlement) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkloadEntitlement) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkloadEntitlement) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkloadEntitlement) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntitlementTermID

`func (o *WorkloadEntitlement) GetEntitlementTermID() string`

GetEntitlementTermID returns the EntitlementTermID field if non-nil, zero value otherwise.

### GetEntitlementTermIDOk

`func (o *WorkloadEntitlement) GetEntitlementTermIDOk() (*string, bool)`

GetEntitlementTermIDOk returns a tuple with the EntitlementTermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTermID

`func (o *WorkloadEntitlement) SetEntitlementTermID(v string)`

SetEntitlementTermID sets EntitlementTermID field to given value.

### HasEntitlementTermID

`func (o *WorkloadEntitlement) HasEntitlementTermID() bool`

HasEntitlementTermID returns a boolean if a field has been set.

### GetExternalBuyerID

`func (o *WorkloadEntitlement) GetExternalBuyerID() string`

GetExternalBuyerID returns the ExternalBuyerID field if non-nil, zero value otherwise.

### GetExternalBuyerIDOk

`func (o *WorkloadEntitlement) GetExternalBuyerIDOk() (*string, bool)`

GetExternalBuyerIDOk returns a tuple with the ExternalBuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBuyerID

`func (o *WorkloadEntitlement) SetExternalBuyerID(v string)`

SetExternalBuyerID sets ExternalBuyerID field to given value.

### HasExternalBuyerID

`func (o *WorkloadEntitlement) HasExternalBuyerID() bool`

HasExternalBuyerID returns a boolean if a field has been set.

### GetExternalID

`func (o *WorkloadEntitlement) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *WorkloadEntitlement) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *WorkloadEntitlement) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *WorkloadEntitlement) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetExternalProductID

`func (o *WorkloadEntitlement) GetExternalProductID() string`

GetExternalProductID returns the ExternalProductID field if non-nil, zero value otherwise.

### GetExternalProductIDOk

`func (o *WorkloadEntitlement) GetExternalProductIDOk() (*string, bool)`

GetExternalProductIDOk returns a tuple with the ExternalProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductID

`func (o *WorkloadEntitlement) SetExternalProductID(v string)`

SetExternalProductID sets ExternalProductID field to given value.

### HasExternalProductID

`func (o *WorkloadEntitlement) HasExternalProductID() bool`

HasExternalProductID returns a boolean if a field has been set.

### GetId

`func (o *WorkloadEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkloadEntitlement) GetInfo() EntitlementInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkloadEntitlement) GetInfoOk() (*EntitlementInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkloadEntitlement) SetInfo(v EntitlementInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkloadEntitlement) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *WorkloadEntitlement) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *WorkloadEntitlement) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *WorkloadEntitlement) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *WorkloadEntitlement) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMetaInfo

`func (o *WorkloadEntitlement) GetMetaInfo() WorkloadMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *WorkloadEntitlement) GetMetaInfoOk() (*WorkloadMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *WorkloadEntitlement) SetMetaInfo(v WorkloadMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *WorkloadEntitlement) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *WorkloadEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferID

`func (o *WorkloadEntitlement) GetOfferID() string`

GetOfferID returns the OfferID field if non-nil, zero value otherwise.

### GetOfferIDOk

`func (o *WorkloadEntitlement) GetOfferIDOk() (*string, bool)`

GetOfferIDOk returns a tuple with the OfferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferID

`func (o *WorkloadEntitlement) SetOfferID(v string)`

SetOfferID sets OfferID field to given value.

### HasOfferID

`func (o *WorkloadEntitlement) HasOfferID() bool`

HasOfferID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *WorkloadEntitlement) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *WorkloadEntitlement) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *WorkloadEntitlement) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *WorkloadEntitlement) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *WorkloadEntitlement) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *WorkloadEntitlement) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *WorkloadEntitlement) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *WorkloadEntitlement) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetProductID

`func (o *WorkloadEntitlement) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *WorkloadEntitlement) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *WorkloadEntitlement) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *WorkloadEntitlement) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetService

`func (o *WorkloadEntitlement) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkloadEntitlement) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkloadEntitlement) SetService(v PartnerService)`

SetService sets Service field to given value.

### HasService

`func (o *WorkloadEntitlement) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkloadEntitlement) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkloadEntitlement) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkloadEntitlement) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkloadEntitlement) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadEntitlement) GetStatus() EntitlementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadEntitlement) GetStatusOk() (*EntitlementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadEntitlement) SetStatus(v EntitlementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadEntitlement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *WorkloadEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkloadEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkloadEntitlement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkloadEntitlement) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


