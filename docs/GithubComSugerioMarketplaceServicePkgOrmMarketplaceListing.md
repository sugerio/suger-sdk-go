# GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | Categories holds the value of the \&quot;categories\&quot; field. | [optional] 
**CompanyDescription** | Pointer to **string** | CompanyDescription holds the value of the \&quot;company_description\&quot; field. | [optional] 
**CompanyDomain** | Pointer to **string** | CompanyDomain holds the value of the \&quot;company_domain\&quot; field. | [optional] 
**CompanyId** | Pointer to **string** | Reference to Company.id where is_vendor&#x3D;true | [optional] 
**CompanyName** | Pointer to **string** | CompanyName holds the value of the \&quot;company_name\&quot; field. | [optional] 
**CreationTime** | Pointer to **string** | CreationTime holds the value of the \&quot;creation_time\&quot; field. | [optional] 
**DeliveryMethod** | Pointer to **string** | DeliveryMethod holds the value of the \&quot;delivery_method\&quot; field. | [optional] 
**HasFreeTrial** | Pointer to **bool** | HasFreeTrial holds the value of the \&quot;has_free_trial\&quot; field. | [optional] 
**HasUsageMetering** | Pointer to **bool** | HasUsageMetering holds the value of the \&quot;has_usage_metering\&quot; field. | [optional] 
**Id** | Pointer to **string** | ID of the ent. | [optional] 
**Info** | Pointer to [**GithubComSugerioMarketplaceServicePkgStructsMarketplaceListingInfo**](GithubComSugerioMarketplaceServicePkgStructsMarketplaceListingInfo.md) | Info holds the value of the \&quot;info\&quot; field. | [optional] 
**LastUpdateTime** | Pointer to **string** | LastUpdateTime holds the value of the \&quot;last_update_time\&quot; field. | [optional] 
**Link** | Pointer to **string** | Link holds the value of the \&quot;link\&quot; field. | [optional] 
**ListingID** | Pointer to **string** | ListingID holds the value of the \&quot;listing_id\&quot; field. | [optional] 
**Name** | Pointer to **string** | Name holds the value of the \&quot;name\&quot; field. | [optional] 
**Partner** | Pointer to **string** | Partner holds the value of the \&quot;partner\&quot; field. | [optional] 
**ReviewCount** | Pointer to **int32** | ReviewCount holds the value of the \&quot;review_count\&quot; field. | [optional] 
**ReviewScore** | Pointer to **float32** | ReviewScore holds the value of the \&quot;review_score\&quot; field. | [optional] 
**ShortDescription** | Pointer to **string** | ShortDescription holds the value of the \&quot;short_description\&quot; field. | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing

`func NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing() *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing`

NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListing instantiates a new GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListingWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListingWithDefaults() *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing`

NewGithubComSugerioMarketplaceServicePkgOrmMarketplaceListingWithDefaults instantiates a new GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCompanyId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetHasFreeTrial

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetHasFreeTrial() bool`

GetHasFreeTrial returns the HasFreeTrial field if non-nil, zero value otherwise.

### GetHasFreeTrialOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetHasFreeTrialOk() (*bool, bool)`

GetHasFreeTrialOk returns a tuple with the HasFreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFreeTrial

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetHasFreeTrial(v bool)`

SetHasFreeTrial sets HasFreeTrial field to given value.

### HasHasFreeTrial

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasHasFreeTrial() bool`

HasHasFreeTrial returns a boolean if a field has been set.

### GetHasUsageMetering

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetHasUsageMetering() bool`

GetHasUsageMetering returns the HasUsageMetering field if non-nil, zero value otherwise.

### GetHasUsageMeteringOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetHasUsageMeteringOk() (*bool, bool)`

GetHasUsageMeteringOk returns a tuple with the HasUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsageMetering

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetHasUsageMetering(v bool)`

SetHasUsageMetering sets HasUsageMetering field to given value.

### HasHasUsageMetering

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasHasUsageMetering() bool`

HasHasUsageMetering returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetInfo() GithubComSugerioMarketplaceServicePkgStructsMarketplaceListingInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetInfoOk() (*GithubComSugerioMarketplaceServicePkgStructsMarketplaceListingInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetInfo(v GithubComSugerioMarketplaceServicePkgStructsMarketplaceListingInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLink

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetListingID

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetListingID() string`

GetListingID returns the ListingID field if non-nil, zero value otherwise.

### GetListingIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetListingIDOk() (*string, bool)`

GetListingIDOk returns a tuple with the ListingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingID

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetListingID(v string)`

SetListingID sets ListingID field to given value.

### HasListingID

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasListingID() bool`

HasListingID returns a boolean if a field has been set.

### GetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetPartner() string`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetPartnerOk() (*string, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetPartner(v string)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetReviewCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetReviewCount() int32`

GetReviewCount returns the ReviewCount field if non-nil, zero value otherwise.

### GetReviewCountOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetReviewCountOk() (*int32, bool)`

GetReviewCountOk returns a tuple with the ReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetReviewCount(v int32)`

SetReviewCount sets ReviewCount field to given value.

### HasReviewCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasReviewCount() bool`

HasReviewCount returns a boolean if a field has been set.

### GetReviewScore

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetReviewScore() float32`

GetReviewScore returns the ReviewScore field if non-nil, zero value otherwise.

### GetReviewScoreOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetReviewScoreOk() (*float32, bool)`

GetReviewScoreOk returns a tuple with the ReviewScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewScore

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetReviewScore(v float32)`

SetReviewScore sets ReviewScore field to given value.

### HasReviewScore

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasReviewScore() bool`

HasReviewScore returns a boolean if a field has been set.

### GetShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmMarketplaceListing) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


