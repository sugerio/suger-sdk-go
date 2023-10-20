# AwsMarketplaceEventBridgeEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **string** |  | [optional] 
**EventCategory** | Pointer to **string** |  | [optional] 
**EventID** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**EventSource** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventVersion** | Pointer to **string** |  | [optional] 
**ManagementEvent** | Pointer to **bool** |  | [optional] 
**Manufacturer** | Pointer to [**AwsMarketplaceEventBridgeEventAccount**](AwsMarketplaceEventBridgeEventAccount.md) |  | [optional] 
**Offer** | Pointer to [**AwsMarketplaceEventBridgeEventOffer**](AwsMarketplaceEventBridgeEventOffer.md) |  | [optional] 
**Product** | Pointer to [**AwsMarketplaceEventBridgeEventProduct**](AwsMarketplaceEventBridgeEventProduct.md) |  | [optional] 
**RequestID** | Pointer to **string** |  | [optional] 
**RequestParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseElements** | Pointer to **map[string]interface{}** |  | [optional] 
**SellerOfRecord** | Pointer to [**AwsMarketplaceEventBridgeEventAccount**](AwsMarketplaceEventBridgeEventAccount.md) |  | [optional] 
**TargetedBuyerAccountIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAwsMarketplaceEventBridgeEventDetail

`func NewAwsMarketplaceEventBridgeEventDetail() *AwsMarketplaceEventBridgeEventDetail`

NewAwsMarketplaceEventBridgeEventDetail instantiates a new AwsMarketplaceEventBridgeEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceEventBridgeEventDetailWithDefaults

`func NewAwsMarketplaceEventBridgeEventDetailWithDefaults() *AwsMarketplaceEventBridgeEventDetail`

NewAwsMarketplaceEventBridgeEventDetailWithDefaults instantiates a new AwsMarketplaceEventBridgeEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *AwsMarketplaceEventBridgeEventDetail) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AwsMarketplaceEventBridgeEventDetail) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AwsMarketplaceEventBridgeEventDetail) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetEventCategory

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.

### HasEventCategory

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventCategory() bool`

HasEventCategory returns a boolean if a field has been set.

### GetEventID

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventName

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventSource

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetEventType

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventVersion

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventVersion() string`

GetEventVersion returns the EventVersion field if non-nil, zero value otherwise.

### GetEventVersionOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetEventVersionOk() (*string, bool)`

GetEventVersionOk returns a tuple with the EventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventVersion

`func (o *AwsMarketplaceEventBridgeEventDetail) SetEventVersion(v string)`

SetEventVersion sets EventVersion field to given value.

### HasEventVersion

`func (o *AwsMarketplaceEventBridgeEventDetail) HasEventVersion() bool`

HasEventVersion returns a boolean if a field has been set.

### GetManagementEvent

`func (o *AwsMarketplaceEventBridgeEventDetail) GetManagementEvent() bool`

GetManagementEvent returns the ManagementEvent field if non-nil, zero value otherwise.

### GetManagementEventOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetManagementEventOk() (*bool, bool)`

GetManagementEventOk returns a tuple with the ManagementEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementEvent

`func (o *AwsMarketplaceEventBridgeEventDetail) SetManagementEvent(v bool)`

SetManagementEvent sets ManagementEvent field to given value.

### HasManagementEvent

`func (o *AwsMarketplaceEventBridgeEventDetail) HasManagementEvent() bool`

HasManagementEvent returns a boolean if a field has been set.

### GetManufacturer

`func (o *AwsMarketplaceEventBridgeEventDetail) GetManufacturer() AwsMarketplaceEventBridgeEventAccount`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetManufacturerOk() (*AwsMarketplaceEventBridgeEventAccount, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *AwsMarketplaceEventBridgeEventDetail) SetManufacturer(v AwsMarketplaceEventBridgeEventAccount)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *AwsMarketplaceEventBridgeEventDetail) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetOffer

`func (o *AwsMarketplaceEventBridgeEventDetail) GetOffer() AwsMarketplaceEventBridgeEventOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetOfferOk() (*AwsMarketplaceEventBridgeEventOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *AwsMarketplaceEventBridgeEventDetail) SetOffer(v AwsMarketplaceEventBridgeEventOffer)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *AwsMarketplaceEventBridgeEventDetail) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetProduct

`func (o *AwsMarketplaceEventBridgeEventDetail) GetProduct() AwsMarketplaceEventBridgeEventProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetProductOk() (*AwsMarketplaceEventBridgeEventProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AwsMarketplaceEventBridgeEventDetail) SetProduct(v AwsMarketplaceEventBridgeEventProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AwsMarketplaceEventBridgeEventDetail) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRequestID

`func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *AwsMarketplaceEventBridgeEventDetail) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *AwsMarketplaceEventBridgeEventDetail) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetRequestParameters

`func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestParameters() map[string]interface{}`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestParametersOk() (*map[string]interface{}, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *AwsMarketplaceEventBridgeEventDetail) SetRequestParameters(v map[string]interface{})`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *AwsMarketplaceEventBridgeEventDetail) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### GetResponseElements

`func (o *AwsMarketplaceEventBridgeEventDetail) GetResponseElements() map[string]interface{}`

GetResponseElements returns the ResponseElements field if non-nil, zero value otherwise.

### GetResponseElementsOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetResponseElementsOk() (*map[string]interface{}, bool)`

GetResponseElementsOk returns a tuple with the ResponseElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseElements

`func (o *AwsMarketplaceEventBridgeEventDetail) SetResponseElements(v map[string]interface{})`

SetResponseElements sets ResponseElements field to given value.

### HasResponseElements

`func (o *AwsMarketplaceEventBridgeEventDetail) HasResponseElements() bool`

HasResponseElements returns a boolean if a field has been set.

### GetSellerOfRecord

`func (o *AwsMarketplaceEventBridgeEventDetail) GetSellerOfRecord() AwsMarketplaceEventBridgeEventAccount`

GetSellerOfRecord returns the SellerOfRecord field if non-nil, zero value otherwise.

### GetSellerOfRecordOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetSellerOfRecordOk() (*AwsMarketplaceEventBridgeEventAccount, bool)`

GetSellerOfRecordOk returns a tuple with the SellerOfRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerOfRecord

`func (o *AwsMarketplaceEventBridgeEventDetail) SetSellerOfRecord(v AwsMarketplaceEventBridgeEventAccount)`

SetSellerOfRecord sets SellerOfRecord field to given value.

### HasSellerOfRecord

`func (o *AwsMarketplaceEventBridgeEventDetail) HasSellerOfRecord() bool`

HasSellerOfRecord returns a boolean if a field has been set.

### GetTargetedBuyerAccountIds

`func (o *AwsMarketplaceEventBridgeEventDetail) GetTargetedBuyerAccountIds() []string`

GetTargetedBuyerAccountIds returns the TargetedBuyerAccountIds field if non-nil, zero value otherwise.

### GetTargetedBuyerAccountIdsOk

`func (o *AwsMarketplaceEventBridgeEventDetail) GetTargetedBuyerAccountIdsOk() (*[]string, bool)`

GetTargetedBuyerAccountIdsOk returns a tuple with the TargetedBuyerAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedBuyerAccountIds

`func (o *AwsMarketplaceEventBridgeEventDetail) SetTargetedBuyerAccountIds(v []string)`

SetTargetedBuyerAccountIds sets TargetedBuyerAccountIds field to given value.

### HasTargetedBuyerAccountIds

`func (o *AwsMarketplaceEventBridgeEventDetail) HasTargetedBuyerAccountIds() bool`

HasTargetedBuyerAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


