# AwsAccountIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS Account ID of the buyer in AWS Marketplace | [optional] 
**AwsCustomerID** | Pointer to **string** | The AWS Customer ID of the buyer in AWS Marketplace | [optional] 
**CompanyInfo** | Pointer to [**CompanyInfo**](CompanyInfo.md) |  | [optional] 
**DataFeedAccountID** | Pointer to **string** | The Account ID in AWS Marketplace Data Feed service | [optional] 

## Methods

### NewAwsAccountIdentifier

`func NewAwsAccountIdentifier() *AwsAccountIdentifier`

NewAwsAccountIdentifier instantiates a new AwsAccountIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAccountIdentifierWithDefaults

`func NewAwsAccountIdentifierWithDefaults() *AwsAccountIdentifier`

NewAwsAccountIdentifierWithDefaults instantiates a new AwsAccountIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *AwsAccountIdentifier) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *AwsAccountIdentifier) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *AwsAccountIdentifier) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *AwsAccountIdentifier) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsCustomerID

`func (o *AwsAccountIdentifier) GetAwsCustomerID() string`

GetAwsCustomerID returns the AwsCustomerID field if non-nil, zero value otherwise.

### GetAwsCustomerIDOk

`func (o *AwsAccountIdentifier) GetAwsCustomerIDOk() (*string, bool)`

GetAwsCustomerIDOk returns a tuple with the AwsCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCustomerID

`func (o *AwsAccountIdentifier) SetAwsCustomerID(v string)`

SetAwsCustomerID sets AwsCustomerID field to given value.

### HasAwsCustomerID

`func (o *AwsAccountIdentifier) HasAwsCustomerID() bool`

HasAwsCustomerID returns a boolean if a field has been set.

### GetCompanyInfo

`func (o *AwsAccountIdentifier) GetCompanyInfo() CompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *AwsAccountIdentifier) GetCompanyInfoOk() (*CompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *AwsAccountIdentifier) SetCompanyInfo(v CompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.

### HasCompanyInfo

`func (o *AwsAccountIdentifier) HasCompanyInfo() bool`

HasCompanyInfo returns a boolean if a field has been set.

### GetDataFeedAccountID

`func (o *AwsAccountIdentifier) GetDataFeedAccountID() string`

GetDataFeedAccountID returns the DataFeedAccountID field if non-nil, zero value otherwise.

### GetDataFeedAccountIDOk

`func (o *AwsAccountIdentifier) GetDataFeedAccountIDOk() (*string, bool)`

GetDataFeedAccountIDOk returns a tuple with the DataFeedAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFeedAccountID

`func (o *AwsAccountIdentifier) SetDataFeedAccountID(v string)`

SetDataFeedAccountID sets DataFeedAccountID field to given value.

### HasDataFeedAccountID

`func (o *AwsAccountIdentifier) HasDataFeedAccountID() bool`

HasDataFeedAccountID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


