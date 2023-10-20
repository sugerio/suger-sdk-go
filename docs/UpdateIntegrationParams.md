# UpdateIntegrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | [**IntegrationInfo**](IntegrationInfo.md) |  | 
**OrganizationID** | **string** |  | 
**Partner** | [**Partner**](Partner.md) |  | 
**Service** | [**PartnerService**](PartnerService.md) |  | 

## Methods

### NewUpdateIntegrationParams

`func NewUpdateIntegrationParams(info IntegrationInfo, organizationID string, partner Partner, service PartnerService, ) *UpdateIntegrationParams`

NewUpdateIntegrationParams instantiates a new UpdateIntegrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationParamsWithDefaults

`func NewUpdateIntegrationParamsWithDefaults() *UpdateIntegrationParams`

NewUpdateIntegrationParamsWithDefaults instantiates a new UpdateIntegrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *UpdateIntegrationParams) GetInfo() IntegrationInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UpdateIntegrationParams) GetInfoOk() (*IntegrationInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UpdateIntegrationParams) SetInfo(v IntegrationInfo)`

SetInfo sets Info field to given value.


### GetOrganizationID

`func (o *UpdateIntegrationParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *UpdateIntegrationParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *UpdateIntegrationParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.


### GetPartner

`func (o *UpdateIntegrationParams) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *UpdateIntegrationParams) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *UpdateIntegrationParams) SetPartner(v Partner)`

SetPartner sets Partner field to given value.


### GetService

`func (o *UpdateIntegrationParams) GetService() PartnerService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UpdateIntegrationParams) GetServiceOk() (*PartnerService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UpdateIntegrationParams) SetService(v PartnerService)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


