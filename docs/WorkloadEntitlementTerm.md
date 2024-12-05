# WorkloadEntitlementTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerID** | Pointer to **string** |  | [optional] 
**CommitAmount** | Pointer to **float32** |  | [optional] 
**CreditAmount** | Pointer to **float32** |  | [optional] 
**EndTime** | Pointer to **time.Time** | nullable | [optional] 
**EntitlementID** | Pointer to **string** |  | [optional] 
**EntitlementInfo** | Pointer to [**EntitlementInfo**](EntitlementInfo.md) |  | [optional] 
**ExternalEntitlementID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**EntitlementTermInfo**](EntitlementTermInfo.md) |  | [optional] 
**OfferID** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**ProductID** | Pointer to **string** |  | [optional] 
**ReportedAmount** | Pointer to **float32** |  | [optional] 
**Service** | Pointer to [**PartnerService**](PartnerService.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**UsedCommitAmount** | Pointer to **float32** |  | [optional] 
**UsedCreditAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewWorkloadEntitlementTerm

`func NewWorkloadEntitlementTerm() *WorkloadEntitlementTerm`

NewWorkloadEntitlementTerm instantiates a new WorkloadEntitlementTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEntitlementTermWithDefaults

`func NewWorkloadEntitlementTermWithDefaults() *WorkloadEntitlementTerm`

NewWorkloadEntitlementTermWithDefaults instantiates a new WorkloadEntitlementTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerID

`func (o *WorkloadEntitlementTerm) GetBuyerID() string`

GetBuyerID returns the BuyerID field if non-nil, zero value otherwise.

### GetBuyerIDOk

`func (o *WorkloadEntitlementTerm) GetBuyerIDOk() (*string, bool)`

GetBuyerIDOk returns a tuple with the BuyerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerID

`func (o *WorkloadEntitlementTerm) SetBuyerID(v string)`

SetBuyerID sets BuyerID field to given value.

### HasBuyerID

`func (o *WorkloadEntitlementTerm) HasBuyerID() bool`

HasBuyerID returns a boolean if a field has been set.

### GetCommitAmount

`func (o *WorkloadEntitlementTerm) GetCommitAmount() float32`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *WorkloadEntitlementTerm) GetCommitAmountOk() (*float32, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *WorkloadEntitlementTerm) SetCommitAmount(v float32)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *WorkloadEntitlementTerm) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *WorkloadEntitlementTerm) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *WorkloadEntitlementTerm) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *WorkloadEntitlementTerm) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *WorkloadEntitlementTerm) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetEndTime

`func (o *WorkloadEntitlementTerm) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkloadEntitlementTerm) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkloadEntitlementTerm) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WorkloadEntitlementTerm) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntitlementID

`func (o *WorkloadEntitlementTerm) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *WorkloadEntitlementTerm) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *WorkloadEntitlementTerm) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *WorkloadEntitlementTerm) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetEntitlementInfo

`func (o *WorkloadEntitlementTerm) GetEntitlementInfo() EntitlementInfo`

GetEntitlementInfo returns the EntitlementInfo field if non-nil, zero value otherwise.

### GetEntitlementInfoOk

`func (o *WorkloadEntitlementTerm) GetEntitlementInfoOk() (*EntitlementInfo, bool)`

GetEntitlementInfoOk returns a tuple with the EntitlementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementInfo

`func (o *WorkloadEntitlementTerm) SetEntitlementInfo(v EntitlementInfo)`

SetEntitlementInfo sets EntitlementInfo field to given value.

### HasEntitlementInfo

`func (o *WorkloadEntitlementTerm) HasEntitlementInfo() bool`

HasEntitlementInfo returns a boolean if a field has been set.

### GetExternalEntitlementID

`func (o *WorkloadEntitlementTerm) GetExternalEntitlementID() string`

GetExternalEntitlementID returns the ExternalEntitlementID field if non-nil, zero value otherwise.

### GetExternalEntitlementIDOk

`func (o *WorkloadEntitlementTerm) GetExternalEntitlementIDOk() (*string, bool)`

GetExternalEntitlementIDOk returns a tuple with the ExternalEntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntitlementID

`func (o *WorkloadEntitlementTerm) SetExternalEntitlementID(v string)`

SetExternalEntitlementID sets ExternalEntitlementID field to given value.

### HasExternalEntitlementID

`func (o *WorkloadEntitlementTerm) HasExternalEntitlementID() bool`

HasExternalEntitlementID returns a boolean if a field has been set.

### GetId

`func (o *WorkloadEntitlementTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadEntitlementTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadEntitlementTerm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadEntitlementTerm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkloadEntitlementTerm) GetInfo() EntitlementTermInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkloadEntitlementTerm) GetInfoOk() (*EntitlementTermInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkloadEntitlementTerm) SetInfo(v EntitlementTermInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkloadEntitlementTerm) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOfferID

`func (o *WorkloadEntitlementTerm) GetOfferID() string`

GetOfferID returns the OfferID field if non-nil, zero value otherwise.

### GetOfferIDOk

`func (o *WorkloadEntitlementTerm) GetOfferIDOk() (*string, bool)`

GetOfferIDOk returns a tuple with the OfferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferID

`func (o *WorkloadEntitlementTerm) SetOfferID(v string)`

SetOfferID sets OfferID field to given value.

### HasOfferID

`func (o *WorkloadEntitlementTerm) HasOfferID() bool`

HasOfferID returns a boolean if a field has been set.

### GetOrganizationID

`func (o *WorkloadEntitlementTerm) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *WorkloadEntitlementTerm) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *WorkloadEntitlementTerm) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *WorkloadEntitlementTerm) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *WorkloadEntitlementTerm) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *WorkloadEntitlementTerm) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *WorkloadEntitlementTerm) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *WorkloadEntitlementTerm) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetProductID

`func (o *WorkloadEntitlementTerm) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *WorkloadEntitlementTerm) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *WorkloadEntitlementTerm) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *WorkloadEntitlementTerm) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetReportedAmount

`func (o *WorkloadEntitlementTerm) GetReportedAmount() float32`

GetReportedAmount returns the ReportedAmount field if non-nil, zero value otherwise.

### GetReportedAmountOk

`func (o *WorkloadEntitlementTerm) GetReportedAmountOk() (*float32, bool)`

GetReportedAmountOk returns a tuple with the ReportedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedAmount

`func (o *WorkloadEntitlementTerm) SetReportedAmount(v float32)`

SetReportedAmount sets ReportedAmount field to given value.

### HasReportedAmount

`func (o *WorkloadEntitlementTerm) HasReportedAmount() bool`

HasReportedAmount returns a boolean if a field has been set.

### GetService

`func (o *WorkloadEntitlementTerm) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WorkloadEntitlementTerm) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WorkloadEntitlementTerm) SetService(v PartnerService)`

SetService sets Service field to given value.

### HasService

`func (o *WorkloadEntitlementTerm) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStartTime

`func (o *WorkloadEntitlementTerm) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkloadEntitlementTerm) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkloadEntitlementTerm) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WorkloadEntitlementTerm) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUsedCommitAmount

`func (o *WorkloadEntitlementTerm) GetUsedCommitAmount() float32`

GetUsedCommitAmount returns the UsedCommitAmount field if non-nil, zero value otherwise.

### GetUsedCommitAmountOk

`func (o *WorkloadEntitlementTerm) GetUsedCommitAmountOk() (*float32, bool)`

GetUsedCommitAmountOk returns a tuple with the UsedCommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCommitAmount

`func (o *WorkloadEntitlementTerm) SetUsedCommitAmount(v float32)`

SetUsedCommitAmount sets UsedCommitAmount field to given value.

### HasUsedCommitAmount

`func (o *WorkloadEntitlementTerm) HasUsedCommitAmount() bool`

HasUsedCommitAmount returns a boolean if a field has been set.

### GetUsedCreditAmount

`func (o *WorkloadEntitlementTerm) GetUsedCreditAmount() float32`

GetUsedCreditAmount returns the UsedCreditAmount field if non-nil, zero value otherwise.

### GetUsedCreditAmountOk

`func (o *WorkloadEntitlementTerm) GetUsedCreditAmountOk() (*float32, bool)`

GetUsedCreditAmountOk returns a tuple with the UsedCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCreditAmount

`func (o *WorkloadEntitlementTerm) SetUsedCreditAmount(v float32)`

SetUsedCreditAmount sets UsedCreditAmount field to given value.

### HasUsedCreditAmount

`func (o *WorkloadEntitlementTerm) HasUsedCreditAmount() bool`

HasUsedCreditAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


