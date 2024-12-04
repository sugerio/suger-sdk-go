# GcpMarketplaceDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DocumentBody** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** | such as \&quot;PARTNER_EULA\&quot; | [optional] 
**ExternalGoogleLink** | Pointer to [**GcpMarketplaceExternalGoogleLink**](GcpMarketplaceExternalGoogleLink.md) | The external link to Standard EULA such as https://cloud.google.com/terms/marketplace/eula-standard-v2-01272021. | [optional] 
**Name** | Pointer to **string** | in format of \&quot;projects/{projectNumber}/agreements/{agreementId}/documents/{documentId}\&quot; | [optional] 
**UnstructuredDocument** | Pointer to [**GcpMarketplaceUnstructuredDocument**](GcpMarketplaceUnstructuredDocument.md) |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGcpMarketplaceDocument

`func NewGcpMarketplaceDocument() *GcpMarketplaceDocument`

NewGcpMarketplaceDocument instantiates a new GcpMarketplaceDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceDocumentWithDefaults

`func NewGcpMarketplaceDocumentWithDefaults() *GcpMarketplaceDocument`

NewGcpMarketplaceDocumentWithDefaults instantiates a new GcpMarketplaceDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GcpMarketplaceDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GcpMarketplaceDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GcpMarketplaceDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GcpMarketplaceDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentBody

`func (o *GcpMarketplaceDocument) GetDocumentBody() string`

GetDocumentBody returns the DocumentBody field if non-nil, zero value otherwise.

### GetDocumentBodyOk

`func (o *GcpMarketplaceDocument) GetDocumentBodyOk() (*string, bool)`

GetDocumentBodyOk returns a tuple with the DocumentBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBody

`func (o *GcpMarketplaceDocument) SetDocumentBody(v string)`

SetDocumentBody sets DocumentBody field to given value.

### HasDocumentBody

`func (o *GcpMarketplaceDocument) HasDocumentBody() bool`

HasDocumentBody returns a boolean if a field has been set.

### GetDocumentType

`func (o *GcpMarketplaceDocument) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GcpMarketplaceDocument) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GcpMarketplaceDocument) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GcpMarketplaceDocument) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetExternalGoogleLink

`func (o *GcpMarketplaceDocument) GetExternalGoogleLink() GcpMarketplaceExternalGoogleLink`

GetExternalGoogleLink returns the ExternalGoogleLink field if non-nil, zero value otherwise.

### GetExternalGoogleLinkOk

`func (o *GcpMarketplaceDocument) GetExternalGoogleLinkOk() (*GcpMarketplaceExternalGoogleLink, bool)`

GetExternalGoogleLinkOk returns a tuple with the ExternalGoogleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGoogleLink

`func (o *GcpMarketplaceDocument) SetExternalGoogleLink(v GcpMarketplaceExternalGoogleLink)`

SetExternalGoogleLink sets ExternalGoogleLink field to given value.

### HasExternalGoogleLink

`func (o *GcpMarketplaceDocument) HasExternalGoogleLink() bool`

HasExternalGoogleLink returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceDocument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceDocument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnstructuredDocument

`func (o *GcpMarketplaceDocument) GetUnstructuredDocument() GcpMarketplaceUnstructuredDocument`

GetUnstructuredDocument returns the UnstructuredDocument field if non-nil, zero value otherwise.

### GetUnstructuredDocumentOk

`func (o *GcpMarketplaceDocument) GetUnstructuredDocumentOk() (*GcpMarketplaceUnstructuredDocument, bool)`

GetUnstructuredDocumentOk returns a tuple with the UnstructuredDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredDocument

`func (o *GcpMarketplaceDocument) SetUnstructuredDocument(v GcpMarketplaceUnstructuredDocument)`

SetUnstructuredDocument sets UnstructuredDocument field to given value.

### HasUnstructuredDocument

`func (o *GcpMarketplaceDocument) HasUnstructuredDocument() bool`

HasUnstructuredDocument returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GcpMarketplaceDocument) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GcpMarketplaceDocument) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GcpMarketplaceDocument) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GcpMarketplaceDocument) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


