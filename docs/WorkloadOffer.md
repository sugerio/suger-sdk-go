# WorkloadOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**ContactIds** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** | nullable | [optional] 
**ExpireTime** | Pointer to **time.Time** | nullable | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**OfferInfo**](OfferInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**MetaInfo** | Pointer to [**WorkloadMetaInfo**](WorkloadMetaInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OfferType** | Pointer to [**OfferType**](OfferType.md) |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**PartnerService**](PartnerService.md) |  | [optional] 
**Status** | Pointer to [**OfferStatus**](OfferStatus.md) |  | [optional] 

## Methods

### NewWorkloadOffer

`func NewWorkloadOffer() *WorkloadOffer`

NewWorkloadOffer instantiates a new WorkloadOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadOfferWithDefaults

`func NewWorkloadOfferWithDefaults() *WorkloadOffer`

NewWorkloadOfferWithDefaults instantiates a new WorkloadOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *WorkloadOffer) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *WorkloadOffer) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *WorkloadOffer) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *WorkloadOffer) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetContactIds

`func (o *WorkloadOffer) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *WorkloadOffer) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *WorkloadOffer) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *WorkloadOffer) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkloadOffer) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkloadOffer) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkloadOffer) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkloadOffer) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreationTime

`func (o *WorkloadOffer) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WorkloadOffer) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WorkloadOffer) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WorkloadOffer) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkloadOffer) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkloadOffer) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkloadOffer) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkloadOffer) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *WorkloadOffer) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *WorkloadOffer) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *WorkloadOffer) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *WorkloadOffer) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetExternalID

`func (o *WorkloadOffer) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *WorkloadOffer) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *WorkloadOffer) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *WorkloadOffer) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetId

`func (o *WorkloadOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadOffer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadOffer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkloadOffer) GetInfo() OfferInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkloadOffer) GetInfoOk() (*OfferInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkloadOffer) SetInfo(v OfferInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkloadOffer) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *WorkloadOffer) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *WorkloadOffer) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *WorkloadOffer) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *WorkloadOffer) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *WorkloadOffer) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *WorkloadOffer) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *WorkloadOffer) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *WorkloadOffer) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetMetaInfo

`func (o *WorkloadOffer) GetMetaInfo() WorkloadMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *WorkloadOffer) GetMetaInfoOk() (*WorkloadMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *WorkloadOffer) SetMetaInfo(v WorkloadMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *WorkloadOffer) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *WorkloadOffer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadOffer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadOffer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadOffer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferType

`func (o *WorkloadOffer) GetOfferType() OfferType`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *WorkloadOffer) GetOfferTypeOk() (*OfferType, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *WorkloadOffer) SetOfferType(v OfferType)`

SetOfferType sets OfferType field to given value.

### HasOfferType

`func (o *WorkloadOffer) HasOfferType() bool`

HasOfferType returns a boolean if a field has been set.

### GetOrganizationID

`func (o *WorkloadOffer) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *WorkloadOffer) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *WorkloadOffer) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *WorkloadOffer) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *WorkloadOffer) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *WorkloadOffer) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *WorkloadOffer) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *WorkloadOffer) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetProductID

`func (o *WorkloadOffer) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *WorkloadOffer) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *WorkloadOffer) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *WorkloadOffer) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetService

`func (o *WorkloadOffer) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkloadOffer) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkloadOffer) SetService(v PartnerService)`

SetService sets Service field to given value.

### HasService

`func (o *WorkloadOffer) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadOffer) GetStatus() OfferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadOffer) GetStatusOk() (*OfferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadOffer) SetStatus(v OfferStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadOffer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


