# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** | Contact | [optional] 
**Country** | Pointer to **string** | Location information | [optional] 
**CreationTime** | Pointer to **string** | Timestamps | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**EmployeeCount** | Pointer to **int32** |  | [optional] 
**FoundedYear** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Core identification fields | [optional] 
**Industry** | Pointer to **string** |  | [optional] 
**IsVendor** | Pointer to **bool** | Vendor flag | [optional] 
**LastUpdateTime** | Pointer to **string** |  | [optional] 
**LinkedIn** | Pointer to **string** | The link to the linkedin profile | [optional] 
**MetaInfo** | Pointer to [**CompanyMetaInfo**](CompanyMetaInfo.md) | Additional information about the company populated by Suger, such as partner engagement scores | [optional] 
**Name** | Pointer to **string** | Basic company information | [optional] 
**PartnerGlobalCompanyContactIds** | Pointer to **[]string** | Partner contacts associated with this buyer company (array of global_company_contact.id values) | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**S3KeyLogo** | Pointer to **string** | Media and social | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**StateProvince** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**EnrichmentDataStatus**](EnrichmentDataStatus.md) | active/stale/archived | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Private, Public, Nonprofit, Franchise | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *Company) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Company) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Company) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Company) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContactEmail

`func (o *Company) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Company) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Company) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Company) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCountry

`func (o *Company) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Company) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Company) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Company) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreationTime

`func (o *Company) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Company) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Company) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Company) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDomain

`func (o *Company) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Company) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Company) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Company) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmployeeCount

`func (o *Company) GetEmployeeCount() int32`

GetEmployeeCount returns the EmployeeCount field if non-nil, zero value otherwise.

### GetEmployeeCountOk

`func (o *Company) GetEmployeeCountOk() (*int32, bool)`

GetEmployeeCountOk returns a tuple with the EmployeeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeCount

`func (o *Company) SetEmployeeCount(v int32)`

SetEmployeeCount sets EmployeeCount field to given value.

### HasEmployeeCount

`func (o *Company) HasEmployeeCount() bool`

HasEmployeeCount returns a boolean if a field has been set.

### GetFoundedYear

`func (o *Company) GetFoundedYear() string`

GetFoundedYear returns the FoundedYear field if non-nil, zero value otherwise.

### GetFoundedYearOk

`func (o *Company) GetFoundedYearOk() (*string, bool)`

GetFoundedYearOk returns a tuple with the FoundedYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundedYear

`func (o *Company) SetFoundedYear(v string)`

SetFoundedYear sets FoundedYear field to given value.

### HasFoundedYear

`func (o *Company) HasFoundedYear() bool`

HasFoundedYear returns a boolean if a field has been set.

### GetId

`func (o *Company) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndustry

`func (o *Company) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Company) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Company) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *Company) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetIsVendor

`func (o *Company) GetIsVendor() bool`

GetIsVendor returns the IsVendor field if non-nil, zero value otherwise.

### GetIsVendorOk

`func (o *Company) GetIsVendorOk() (*bool, bool)`

GetIsVendorOk returns a tuple with the IsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendor

`func (o *Company) SetIsVendor(v bool)`

SetIsVendor sets IsVendor field to given value.

### HasIsVendor

`func (o *Company) HasIsVendor() bool`

HasIsVendor returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *Company) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *Company) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *Company) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *Company) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLinkedIn

`func (o *Company) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *Company) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *Company) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *Company) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetMetaInfo

`func (o *Company) GetMetaInfo() CompanyMetaInfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *Company) GetMetaInfoOk() (*CompanyMetaInfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *Company) SetMetaInfo(v CompanyMetaInfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *Company) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerGlobalCompanyContactIds

`func (o *Company) GetPartnerGlobalCompanyContactIds() []string`

GetPartnerGlobalCompanyContactIds returns the PartnerGlobalCompanyContactIds field if non-nil, zero value otherwise.

### GetPartnerGlobalCompanyContactIdsOk

`func (o *Company) GetPartnerGlobalCompanyContactIdsOk() (*[]string, bool)`

GetPartnerGlobalCompanyContactIdsOk returns a tuple with the PartnerGlobalCompanyContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerGlobalCompanyContactIds

`func (o *Company) SetPartnerGlobalCompanyContactIds(v []string)`

SetPartnerGlobalCompanyContactIds sets PartnerGlobalCompanyContactIds field to given value.

### HasPartnerGlobalCompanyContactIds

`func (o *Company) HasPartnerGlobalCompanyContactIds() bool`

HasPartnerGlobalCompanyContactIds returns a boolean if a field has been set.

### GetPostalCode

`func (o *Company) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Company) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Company) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Company) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetS3KeyLogo

`func (o *Company) GetS3KeyLogo() string`

GetS3KeyLogo returns the S3KeyLogo field if non-nil, zero value otherwise.

### GetS3KeyLogoOk

`func (o *Company) GetS3KeyLogoOk() (*string, bool)`

GetS3KeyLogoOk returns a tuple with the S3KeyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyLogo

`func (o *Company) SetS3KeyLogo(v string)`

SetS3KeyLogo sets S3KeyLogo field to given value.

### HasS3KeyLogo

`func (o *Company) HasS3KeyLogo() bool`

HasS3KeyLogo returns a boolean if a field has been set.

### GetShortDescription

`func (o *Company) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Company) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Company) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Company) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetStateProvince

`func (o *Company) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *Company) GetStateProvinceOk() (*string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProvince

`func (o *Company) SetStateProvince(v string)`

SetStateProvince sets StateProvince field to given value.

### HasStateProvince

`func (o *Company) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### GetStatus

`func (o *Company) GetStatus() EnrichmentDataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Company) GetStatusOk() (*EnrichmentDataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Company) SetStatus(v EnrichmentDataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Company) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreet

`func (o *Company) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Company) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Company) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *Company) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetType

`func (o *Company) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Company) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Company) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Company) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebsite

`func (o *Company) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Company) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Company) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Company) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


