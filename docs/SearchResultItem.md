# SearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Tags** | Pointer to [**[]GithubComSugerioMarketplaceServicePkgSearchTag**](GithubComSugerioMarketplaceServicePkgSearchTag.md) | Tags provides structured tags for the search result | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchResultItem

`func NewSearchResultItem() *SearchResultItem`

NewSearchResultItem instantiates a new SearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultItemWithDefaults

`func NewSearchResultItemWithDefaults() *SearchResultItem`

NewSearchResultItemWithDefaults instantiates a new SearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SearchResultItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchResultItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchResultItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchResultItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchResultItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchResultItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *SearchResultItem) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SearchResultItem) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SearchResultItem) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SearchResultItem) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SearchResultItem) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SearchResultItem) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SearchResultItem) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SearchResultItem) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetScore

`func (o *SearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchResultItem) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTags

`func (o *SearchResultItem) GetTags() []GithubComSugerioMarketplaceServicePkgSearchTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchResultItem) GetTagsOk() (*[]GithubComSugerioMarketplaceServicePkgSearchTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchResultItem) SetTags(v []GithubComSugerioMarketplaceServicePkgSearchTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchResultItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SearchResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResultItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchResultItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


