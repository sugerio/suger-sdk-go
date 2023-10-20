# CreateIntegrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**Info** | [**IntegrationInfo**](IntegrationInfo.md) |  | 
**OrganizationID** | **string** |  | 
**Partner** | [**Partner**](Partner.md) |  | 
**Service** | [**PartnerService**](PartnerService.md) |  | 

## Methods

### NewCreateIntegrationParams

`func NewCreateIntegrationParams(info IntegrationInfo, organizationID string, partner Partner, service PartnerService, ) *CreateIntegrationParams`

NewCreateIntegrationParams instantiates a new CreateIntegrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationParamsWithDefaults

`func NewCreateIntegrationParamsWithDefaults() *CreateIntegrationParams`

NewCreateIntegrationParamsWithDefaults instantiates a new CreateIntegrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *CreateIntegrationParams) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateIntegrationParams) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateIntegrationParams) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateIntegrationParams) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetInfo

`func (o *CreateIntegrationParams) GetInfo() IntegrationInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateIntegrationParams) GetInfoOk() (*IntegrationInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateIntegrationParams) SetInfo(v IntegrationInfo)`

SetInfo sets Info field to given value.


### GetOrganizationID

`func (o *CreateIntegrationParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *CreateIntegrationParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *CreateIntegrationParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.


### GetPartner

`func (o *CreateIntegrationParams) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *CreateIntegrationParams) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *CreateIntegrationParams) SetPartner(v Partner)`

SetPartner sets Partner field to given value.


### GetService

`func (o *CreateIntegrationParams) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateIntegrationParams) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateIntegrationParams) SetService(v PartnerService)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


