# TypesEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIdentifier** | Pointer to **string** | The customer identifier is a handle to each unique customer in an application. Customer identifiers are obtained through the ResolveCustomer operation in AWS Marketplace Metering Service. | [optional] 
**Dimension** | Pointer to **string** | The dimension for which the given entitlement applies. Dimensions represent categories of capacity in a product and are specified when the product is listed in AWS Marketplace. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date represents the minimum date through which this entitlement is expected to remain valid. For contractual products listed on AWS Marketplace, the expiration date is the date at which the customer will renew or cancel their contract. Customers who are opting to renew their contract will still have entitlements with an expiration date. | [optional] 
**ProductCode** | Pointer to **string** | The product code for which the given entitlement applies. Product codes are provided by AWS Marketplace when the product listing is created. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The EntitlementValue represents the amount of capacity that the customer is entitled to for the product. | [optional] 

## Methods

### NewTypesEntitlement

`func NewTypesEntitlement() *TypesEntitlement`

NewTypesEntitlement instantiates a new TypesEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesEntitlementWithDefaults

`func NewTypesEntitlementWithDefaults() *TypesEntitlement`

NewTypesEntitlementWithDefaults instantiates a new TypesEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIdentifier

`func (o *TypesEntitlement) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *TypesEntitlement) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *TypesEntitlement) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *TypesEntitlement) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetDimension

`func (o *TypesEntitlement) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *TypesEntitlement) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *TypesEntitlement) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *TypesEntitlement) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetExpirationDate

`func (o *TypesEntitlement) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TypesEntitlement) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TypesEntitlement) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TypesEntitlement) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetProductCode

`func (o *TypesEntitlement) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *TypesEntitlement) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *TypesEntitlement) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *TypesEntitlement) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetValue

`func (o *TypesEntitlement) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TypesEntitlement) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TypesEntitlement) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TypesEntitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


