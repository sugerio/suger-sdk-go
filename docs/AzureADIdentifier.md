# AzureADIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAccountId** | Pointer to **string** | Azure Billing Account ID | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**EmailId** | Pointer to **string** | Email address | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** | Azure License Type | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Puid** | Pointer to **string** | ID of the user, used as External ID of suger IdentityBuyer. | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureADIdentifier

`func NewAzureADIdentifier() *AzureADIdentifier`

NewAzureADIdentifier instantiates a new AzureADIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureADIdentifierWithDefaults

`func NewAzureADIdentifierWithDefaults() *AzureADIdentifier`

NewAzureADIdentifierWithDefaults instantiates a new AzureADIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAccountId

`func (o *AzureADIdentifier) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AzureADIdentifier) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AzureADIdentifier) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AzureADIdentifier) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AzureADIdentifier) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AzureADIdentifier) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AzureADIdentifier) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AzureADIdentifier) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmailId

`func (o *AzureADIdentifier) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *AzureADIdentifier) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *AzureADIdentifier) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.

### HasEmailId

`func (o *AzureADIdentifier) HasEmailId() bool`

HasEmailId returns a boolean if a field has been set.

### GetFirstName

`func (o *AzureADIdentifier) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AzureADIdentifier) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AzureADIdentifier) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AzureADIdentifier) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AzureADIdentifier) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AzureADIdentifier) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AzureADIdentifier) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AzureADIdentifier) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLicenseType

`func (o *AzureADIdentifier) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *AzureADIdentifier) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *AzureADIdentifier) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *AzureADIdentifier) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetObjectId

`func (o *AzureADIdentifier) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AzureADIdentifier) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AzureADIdentifier) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *AzureADIdentifier) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPuid

`func (o *AzureADIdentifier) GetPuid() string`

GetPuid returns the Puid field if non-nil, zero value otherwise.

### GetPuidOk

`func (o *AzureADIdentifier) GetPuidOk() (*string, bool)`

GetPuidOk returns a tuple with the Puid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuid

`func (o *AzureADIdentifier) SetPuid(v string)`

SetPuid sets Puid field to given value.

### HasPuid

`func (o *AzureADIdentifier) HasPuid() bool`

HasPuid returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureADIdentifier) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureADIdentifier) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureADIdentifier) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureADIdentifier) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


