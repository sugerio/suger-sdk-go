# GithubComSugerioMarketplaceServicePkgOrmGlobalCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City holds the value of the \&quot;city\&quot; field. | [optional] 
**ContactEmail** | Pointer to **string** | ContactEmail holds the value of the \&quot;contact_email\&quot; field. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-3 country codes such as USA, JPN, CHN | [optional] 
**CreationTime** | Pointer to **string** | CreationTime holds the value of the \&quot;creation_time\&quot; field. | [optional] 
**Domain** | Pointer to **string** | Domain holds the value of the \&quot;domain\&quot; field. | [optional] 
**EmployeeCount** | Pointer to **int32** | EmployeeCount holds the value of the \&quot;employee_count\&quot; field. | [optional] 
**FoundedYear** | Pointer to **string** | Year founded (e.g. 2021) | [optional] 
**Id** | Pointer to **string** | ID of the ent. | [optional] 
**Industry** | Pointer to **string** | Industry holds the value of the \&quot;industry\&quot; field. | [optional] 
**IsVendor** | Pointer to **bool** | Indicates if this company is a vendor in the marketplace catalog | [optional] 
**LastUpdateTime** | Pointer to **string** | LastUpdateTime holds the value of the \&quot;last_update_time\&quot; field. | [optional] 
**LinkedIn** | Pointer to **string** | The link to the linkedin profile | [optional] 
**MetaInfo** | Pointer to [**CompanyMetaInfo**](CompanyMetaInfo.md) | Additional information about the company populated by Suger, such as partner engagement scores | [optional] 
**Name** | Pointer to **string** | Name holds the value of the \&quot;name\&quot; field. | [optional] 
**PartnerGlobalCompanyContactIds** | Pointer to **[]string** | Array of unique global_company_contact.id values. Used to track partner contacts associated with this buyer company. | [optional] 
**PostalCode** | Pointer to **string** | PostalCode holds the value of the \&quot;postal_code\&quot; field. | [optional] 
**Provider** | Pointer to **string** | The data provider | [optional] 
**ProviderCompanyID** | Pointer to **string** | The external id from the provider | [optional] 
**S3KeyLogo** | Pointer to **string** | S3KeyLogo holds the value of the \&quot;s3_key_logo\&quot; field. | [optional] 
**S3KeyProviderData** | Pointer to **string** | The normalized data from provider | [optional] 
**ShortDescription** | Pointer to **string** | ShortDescription holds the value of the \&quot;short_description\&quot; field. | [optional] 
**StateProvince** | Pointer to **string** | State or province | [optional] 
**Status** | Pointer to **string** | active/stale/archived | [optional] 
**Street** | Pointer to **string** | Street holds the value of the \&quot;street\&quot; field. | [optional] 
**Type** | Pointer to **string** | Company type (e.g. Private, Public, Nonprofit, Franchise) | [optional] 
**Website** | Pointer to **string** | The main website | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompany

`func NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompany() *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany`

NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompany instantiates a new GithubComSugerioMarketplaceServicePkgOrmGlobalCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompanyWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompanyWithDefaults() *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany`

NewGithubComSugerioMarketplaceServicePkgOrmGlobalCompanyWithDefaults instantiates a new GithubComSugerioMarketplaceServicePkgOrmGlobalCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContactEmail

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCountry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmployeeCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetEmployeeCount() int32`

GetEmployeeCount returns the EmployeeCount field if non-nil, zero value otherwise.

### GetEmployeeCountOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetEmployeeCountOk() (*int32, bool)`

GetEmployeeCountOk returns a tuple with the EmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetEmployeeCount(v int32)`

SetEmployeeCount sets EmployeeCount field to given value.

### HasEmployeeCount

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasEmployeeCount() bool`

HasEmployeeCount returns a boolean if a field has been set.

### GetFoundedYear

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetFoundedYear() string`

GetFoundedYear returns the FoundedYear field if non-nil, zero value otherwise.

### GetFoundedYearOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetFoundedYearOk() (*string, bool)`

GetFoundedYearOk returns a tuple with the FoundedYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundedYear

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetFoundedYear(v string)`

SetFoundedYear sets FoundedYear field to given value.

### HasFoundedYear

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasFoundedYear() bool`

HasFoundedYear returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndustry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetIsVendor

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetIsVendor() bool`

GetIsVendor returns the IsVendor field if non-nil, zero value otherwise.

### GetIsVendorOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetIsVendorOk() (*bool, bool)`

GetIsVendorOk returns a tuple with the IsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendor

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetIsVendor(v bool)`

SetIsVendor sets IsVendor field to given value.

### HasIsVendor

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasIsVendor() bool`

HasIsVendor returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLinkedIn

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetMetaInfo() CompanyMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetMetaInfoOk() (*CompanyMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetMetaInfo(v CompanyMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerGlobalCompanyContactIds

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetPartnerGlobalCompanyContactIds() []string`

GetPartnerGlobalCompanyContactIds returns the PartnerGlobalCompanyContactIds field if non-nil, zero value otherwise.

### GetPartnerGlobalCompanyContactIdsOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetPartnerGlobalCompanyContactIdsOk() (*[]string, bool)`

GetPartnerGlobalCompanyContactIdsOk returns a tuple with the PartnerGlobalCompanyContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerGlobalCompanyContactIds

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetPartnerGlobalCompanyContactIds(v []string)`

SetPartnerGlobalCompanyContactIds sets PartnerGlobalCompanyContactIds field to given value.

### HasPartnerGlobalCompanyContactIds

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasPartnerGlobalCompanyContactIds() bool`

HasPartnerGlobalCompanyContactIds returns a boolean if a field has been set.

### GetPostalCode

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetProvider

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderCompanyID

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetProviderCompanyID() string`

GetProviderCompanyID returns the ProviderCompanyID field if non-nil, zero value otherwise.

### GetProviderCompanyIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetProviderCompanyIDOk() (*string, bool)`

GetProviderCompanyIDOk returns a tuple with the ProviderCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCompanyID

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetProviderCompanyID(v string)`

SetProviderCompanyID sets ProviderCompanyID field to given value.

### HasProviderCompanyID

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasProviderCompanyID() bool`

HasProviderCompanyID returns a boolean if a field has been set.

### GetS3KeyLogo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetS3KeyLogo() string`

GetS3KeyLogo returns the S3KeyLogo field if non-nil, zero value otherwise.

### GetS3KeyLogoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetS3KeyLogoOk() (*string, bool)`

GetS3KeyLogoOk returns a tuple with the S3KeyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyLogo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetS3KeyLogo(v string)`

SetS3KeyLogo sets S3KeyLogo field to given value.

### HasS3KeyLogo

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasS3KeyLogo() bool`

HasS3KeyLogo returns a boolean if a field has been set.

### GetS3KeyProviderData

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetS3KeyProviderData() string`

GetS3KeyProviderData returns the S3KeyProviderData field if non-nil, zero value otherwise.

### GetS3KeyProviderDataOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetS3KeyProviderDataOk() (*string, bool)`

GetS3KeyProviderDataOk returns a tuple with the S3KeyProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyProviderData

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetS3KeyProviderData(v string)`

SetS3KeyProviderData sets S3KeyProviderData field to given value.

### HasS3KeyProviderData

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasS3KeyProviderData() bool`

HasS3KeyProviderData returns a boolean if a field has been set.

### GetShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetStateProvince

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreet

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebsite

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *GithubComSugerioMarketplaceServicePkgOrmGlobalCompany) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


