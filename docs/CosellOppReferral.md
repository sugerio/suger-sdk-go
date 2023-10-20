# CosellOppReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to [**CosellOpp**](CosellOpp.md) |  | [optional] 
**DestinationID** | Pointer to **string** |  | [optional] 
**DestinationPartner** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **string** |  | [optional] 
**MetaInfo** | Pointer to [**CosellOppMeta**](CosellOppMeta.md) |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**CosellOpp**](CosellOpp.md) |  | [optional] 
**SourceID** | Pointer to **string** |  | [optional] 
**SourcePartner** | Pointer to **string** |  | [optional] 

## Methods

### NewCosellOppReferral

`func NewCosellOppReferral() *CosellOppReferral`

NewCosellOppReferral instantiates a new CosellOppReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosellOppReferralWithDefaults

`func NewCosellOppReferralWithDefaults() *CosellOppReferral`

NewCosellOppReferralWithDefaults instantiates a new CosellOppReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *CosellOppReferral) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CosellOppReferral) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CosellOppReferral) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CosellOppReferral) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDestination

`func (o *CosellOppReferral) GetDestination() CosellOpp`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CosellOppReferral) GetDestinationOk() (*CosellOpp, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CosellOppReferral) SetDestination(v CosellOpp)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CosellOppReferral) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationID

`func (o *CosellOppReferral) GetDestinationID() string`

GetDestinationID returns the DestinationID field if non-nil, zero value otherwise.

### GetDestinationIDOk

`func (o *CosellOppReferral) GetDestinationIDOk() (*string, bool)`

GetDestinationIDOk returns a tuple with the DestinationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationID

`func (o *CosellOppReferral) SetDestinationID(v string)`

SetDestinationID sets DestinationID field to given value.

### HasDestinationID

`func (o *CosellOppReferral) HasDestinationID() bool`

HasDestinationID returns a boolean if a field has been set.

### GetDestinationPartner

`func (o *CosellOppReferral) GetDestinationPartner() string`

GetDestinationPartner returns the DestinationPartner field if non-nil, zero value otherwise.

### GetDestinationPartnerOk

`func (o *CosellOppReferral) GetDestinationPartnerOk() (*string, bool)`

GetDestinationPartnerOk returns a tuple with the DestinationPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPartner

`func (o *CosellOppReferral) SetDestinationPartner(v string)`

SetDestinationPartner sets DestinationPartner field to given value.

### HasDestinationPartner

`func (o *CosellOppReferral) HasDestinationPartner() bool`

HasDestinationPartner returns a boolean if a field has been set.

### GetId

`func (o *CosellOppReferral) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CosellOppReferral) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CosellOppReferral) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CosellOppReferral) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *CosellOppReferral) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *CosellOppReferral) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *CosellOppReferral) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *CosellOppReferral) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMetaInfo

`func (o *CosellOppReferral) GetMetaInfo() CosellOppMeta`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *CosellOppReferral) GetMetaInfoOk() (*CosellOppMeta, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *CosellOppReferral) SetMetaInfo(v CosellOppMeta)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *CosellOppReferral) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CosellOppReferral) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CosellOppReferral) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CosellOppReferral) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CosellOppReferral) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSource

`func (o *CosellOppReferral) GetSource() CosellOpp`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CosellOppReferral) GetSourceOk() (*CosellOpp, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CosellOppReferral) SetSource(v CosellOpp)`

SetSource sets Source field to given value.

### HasSource

`func (o *CosellOppReferral) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceID

`func (o *CosellOppReferral) GetSourceID() string`

GetSourceID returns the SourceID field if non-nil, zero value otherwise.

### GetSourceIDOk

`func (o *CosellOppReferral) GetSourceIDOk() (*string, bool)`

GetSourceIDOk returns a tuple with the SourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceID

`func (o *CosellOppReferral) SetSourceID(v string)`

SetSourceID sets SourceID field to given value.

### HasSourceID

`func (o *CosellOppReferral) HasSourceID() bool`

HasSourceID returns a boolean if a field has been set.

### GetSourcePartner

`func (o *CosellOppReferral) GetSourcePartner() string`

GetSourcePartner returns the SourcePartner field if non-nil, zero value otherwise.

### GetSourcePartnerOk

`func (o *CosellOppReferral) GetSourcePartnerOk() (*string, bool)`

GetSourcePartnerOk returns a tuple with the SourcePartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePartner

`func (o *CosellOppReferral) SetSourcePartner(v string)`

SetSourcePartner sets SourcePartner field to given value.

### HasSourcePartner

`func (o *CosellOppReferral) HasSourcePartner() bool`

HasSourcePartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


