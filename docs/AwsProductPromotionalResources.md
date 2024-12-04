# AwsProductPromotionalResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalResources** | Pointer to [**[]AwsProductAdditionalResource**](AwsProductAdditionalResource.md) |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**VideoUrls** | Pointer to **[]string** | Currently, AWS only support 1 url in the array. | [optional] 

## Methods

### NewAwsProductPromotionalResources

`func NewAwsProductPromotionalResources() *AwsProductPromotionalResources`

NewAwsProductPromotionalResources instantiates a new AwsProductPromotionalResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProductPromotionalResourcesWithDefaults

`func NewAwsProductPromotionalResourcesWithDefaults() *AwsProductPromotionalResources`

NewAwsProductPromotionalResourcesWithDefaults instantiates a new AwsProductPromotionalResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalResources

`func (o *AwsProductPromotionalResources) GetAdditionalResources() []AwsProductAdditionalResource`

GetAdditionalResources returns the AdditionalResources field if non-nil, zero value otherwise.

### GetAdditionalResourcesOk

`func (o *AwsProductPromotionalResources) GetAdditionalResourcesOk() (*[]AwsProductAdditionalResource, bool)`

GetAdditionalResourcesOk returns a tuple with the AdditionalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResources

`func (o *AwsProductPromotionalResources) SetAdditionalResources(v []AwsProductAdditionalResource)`

SetAdditionalResources sets AdditionalResources field to given value.

### HasAdditionalResources

`func (o *AwsProductPromotionalResources) HasAdditionalResources() bool`

HasAdditionalResources returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AwsProductPromotionalResources) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AwsProductPromotionalResources) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AwsProductPromotionalResources) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AwsProductPromotionalResources) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetVideoUrls

`func (o *AwsProductPromotionalResources) GetVideoUrls() []string`

GetVideoUrls returns the VideoUrls field if non-nil, zero value otherwise.

### GetVideoUrlsOk

`func (o *AwsProductPromotionalResources) GetVideoUrlsOk() (*[]string, bool)`

GetVideoUrlsOk returns a tuple with the VideoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrls

`func (o *AwsProductPromotionalResources) SetVideoUrls(v []string)`

SetVideoUrls sets VideoUrls field to given value.

### HasVideoUrls

`func (o *AwsProductPromotionalResources) HasVideoUrls() bool`

HasVideoUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


