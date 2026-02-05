# GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAWSAccountId** | Pointer to **string** | The CustomerAWSAccountID parameter specifies the AWS account ID of the buyer. | [optional] 
**CustomerIdentifier** | Pointer to **string** | The customer identifier is a handle to each unique customer in an application. Customer identifiers are obtained through the ResolveCustomer operation in AWS Marketplace Metering Service. | [optional] 
**Dimension** | Pointer to **string** | The dimension for which the given entitlement applies. Dimensions represent categories of capacity in a product and are specified when the product is listed in AWS Marketplace. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date represents the minimum date through which this entitlement is expected to remain valid. For contractual products listed on AWS Marketplace, the expiration date is the date at which the customer will renew or cancel their contract. Customers who are opting to renew their contract will still have entitlements with an expiration date. | [optional] 
**LicenseArn** | Pointer to **string** |  | [optional] 
**ProductCode** | Pointer to **string** | The product code for which the given entitlement applies. Product codes are provided by AWS Marketplace when the product listing is created. | [optional] 
**Value** | Pointer to [**GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementValue**](GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementValue.md) | The EntitlementValue represents the amount of capacity that the customer is entitled to for the product. | [optional] 

## Methods

### NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement() *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement`

NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementWithDefaults

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementWithDefaults() *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement`

NewGithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementWithDefaults instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetCustomerAWSAccountId() string`

GetCustomerAWSAccountId returns the CustomerAWSAccountId field if non-nil, zero value otherwise.

### GetCustomerAWSAccountIdOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetCustomerAWSAccountIdOk() (*string, bool)`

GetCustomerAWSAccountIdOk returns a tuple with the CustomerAWSAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetCustomerAWSAccountId(v string)`

SetCustomerAWSAccountId sets CustomerAWSAccountId field to given value.

### HasCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasCustomerAWSAccountId() bool`

HasCustomerAWSAccountId returns a boolean if a field has been set.

### GetCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetLicenseArn() string`

GetLicenseArn returns the LicenseArn field if non-nil, zero value otherwise.

### GetLicenseArnOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetLicenseArnOk() (*string, bool)`

GetLicenseArnOk returns a tuple with the LicenseArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetLicenseArn(v string)`

SetLicenseArn sets LicenseArn field to given value.

### HasLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasLicenseArn() bool`

HasLicenseArn returns a boolean if a field has been set.

### GetProductCode

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetValue

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetValue() GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) GetValueOk() (*GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) SetValue(v GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlementValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplaceentitlementserviceTypesEntitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


