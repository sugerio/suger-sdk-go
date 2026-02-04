# AwsMarketplacePurchaseConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultipleDimensionSelection** | Pointer to [**AwsMarketplaceCatalogConstraintsEnum**](AwsMarketplaceCatalogConstraintsEnum.md) | Determines if buyers are allowed to select multiple dimensions in the rate card. Possible values are \&quot;Allowed\&quot; and \&quot;Disallowed\&quot;. Default value is \&quot;Allowed\&quot;. | [optional] 
**QuantityConfiguration** | Pointer to [**AwsMarketplaceCatalogConstraintsEnum**](AwsMarketplaceCatalogConstraintsEnum.md) | Determines if acceptors are allowed to configure quantity for each dimension in rate card. Possible values are \&quot;Allowed\&quot; and \&quot;Disallowed\&quot;. Default value is \&quot;Allowed\&quot;. | [optional] 

## Methods

### NewAwsMarketplacePurchaseConstraints

`func NewAwsMarketplacePurchaseConstraints() *AwsMarketplacePurchaseConstraints`

NewAwsMarketplacePurchaseConstraints instantiates a new AwsMarketplacePurchaseConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplacePurchaseConstraintsWithDefaults

`func NewAwsMarketplacePurchaseConstraintsWithDefaults() *AwsMarketplacePurchaseConstraints`

NewAwsMarketplacePurchaseConstraintsWithDefaults instantiates a new AwsMarketplacePurchaseConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultipleDimensionSelection

`func (o *AwsMarketplacePurchaseConstraints) GetMultipleDimensionSelection() AwsMarketplaceCatalogConstraintsEnum`

GetMultipleDimensionSelection returns the MultipleDimensionSelection field if non-nil, zero value otherwise.

### GetMultipleDimensionSelectionOk

`func (o *AwsMarketplacePurchaseConstraints) GetMultipleDimensionSelectionOk() (*AwsMarketplaceCatalogConstraintsEnum, bool)`

GetMultipleDimensionSelectionOk returns a tuple with the MultipleDimensionSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleDimensionSelection

`func (o *AwsMarketplacePurchaseConstraints) SetMultipleDimensionSelection(v AwsMarketplaceCatalogConstraintsEnum)`

SetMultipleDimensionSelection sets MultipleDimensionSelection field to given value.

### HasMultipleDimensionSelection

`func (o *AwsMarketplacePurchaseConstraints) HasMultipleDimensionSelection() bool`

HasMultipleDimensionSelection returns a boolean if a field has been set.

### GetQuantityConfiguration

`func (o *AwsMarketplacePurchaseConstraints) GetQuantityConfiguration() AwsMarketplaceCatalogConstraintsEnum`

GetQuantityConfiguration returns the QuantityConfiguration field if non-nil, zero value otherwise.

### GetQuantityConfigurationOk

`func (o *AwsMarketplacePurchaseConstraints) GetQuantityConfigurationOk() (*AwsMarketplaceCatalogConstraintsEnum, bool)`

GetQuantityConfigurationOk returns a tuple with the QuantityConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfiguration

`func (o *AwsMarketplacePurchaseConstraints) SetQuantityConfiguration(v AwsMarketplaceCatalogConstraintsEnum)`

SetQuantityConfiguration sets QuantityConfiguration field to given value.

### HasQuantityConfiguration

`func (o *AwsMarketplacePurchaseConstraints) HasQuantityConfiguration() bool`

HasQuantityConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


