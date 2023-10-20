# BuyerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsBuyer** | Pointer to [**AwsAccountIdentifier**](AwsAccountIdentifier.md) |  | [optional] 
**AzureBuyer** | Pointer to [**AzureADIdentifier**](AzureADIdentifier.md) |  | [optional] 
**CollectableAmount** | Pointer to **float32** | The amount that the seller can collect. It excludes the marketplace commision fee. | [optional] 
**CustomerId** | Pointer to **string** | customerID of buyer on seller&#39;s side | [optional] 
**DisbursedAmount** | Pointer to **float32** | The amount that has been disbursed to the seller account. | [optional] 
**GcpBuyer** | Pointer to [**GcpMarketplaceUserAccount**](GcpMarketplaceUserAccount.md) |  | [optional] 
**InvoicedAmount** | Pointer to **float32** | The amount that the buyer has got invoiced. | [optional] 
**MetronomeCustomerId** | Pointer to **string** | The metronome customer ID for the buyer if it is connected to a metronome customer. | [optional] 
**OrbCustomerId** | Pointer to **string** | The orb customer ID for the buyer if it is connected to a orb customer. | [optional] 

## Methods

### NewBuyerInfo

`func NewBuyerInfo() *BuyerInfo`

NewBuyerInfo instantiates a new BuyerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerInfoWithDefaults

`func NewBuyerInfoWithDefaults() *BuyerInfo`

NewBuyerInfoWithDefaults instantiates a new BuyerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsBuyer

`func (o *BuyerInfo) GetAwsBuyer() AwsAccountIdentifier`

GetAwsBuyer returns the AwsBuyer field if non-nil, zero value otherwise.

### GetAwsBuyerOk

`func (o *BuyerInfo) GetAwsBuyerOk() (*AwsAccountIdentifier, bool)`

GetAwsBuyerOk returns a tuple with the AwsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBuyer

`func (o *BuyerInfo) SetAwsBuyer(v AwsAccountIdentifier)`

SetAwsBuyer sets AwsBuyer field to given value.

### HasAwsBuyer

`func (o *BuyerInfo) HasAwsBuyer() bool`

HasAwsBuyer returns a boolean if a field has been set.

### GetAzureBuyer

`func (o *BuyerInfo) GetAzureBuyer() AzureADIdentifier`

GetAzureBuyer returns the AzureBuyer field if non-nil, zero value otherwise.

### GetAzureBuyerOk

`func (o *BuyerInfo) GetAzureBuyerOk() (*AzureADIdentifier, bool)`

GetAzureBuyerOk returns a tuple with the AzureBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBuyer

`func (o *BuyerInfo) SetAzureBuyer(v AzureADIdentifier)`

SetAzureBuyer sets AzureBuyer field to given value.

### HasAzureBuyer

`func (o *BuyerInfo) HasAzureBuyer() bool`

HasAzureBuyer returns a boolean if a field has been set.

### GetCollectableAmount

`func (o *BuyerInfo) GetCollectableAmount() float32`

GetCollectableAmount returns the CollectableAmount field if non-nil, zero value otherwise.

### GetCollectableAmountOk

`func (o *BuyerInfo) GetCollectableAmountOk() (*float32, bool)`

GetCollectableAmountOk returns a tuple with the CollectableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectableAmount

`func (o *BuyerInfo) SetCollectableAmount(v float32)`

SetCollectableAmount sets CollectableAmount field to given value.

### HasCollectableAmount

`func (o *BuyerInfo) HasCollectableAmount() bool`

HasCollectableAmount returns a boolean if a field has been set.

### GetCustomerId

`func (o *BuyerInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BuyerInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BuyerInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *BuyerInfo) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDisbursedAmount

`func (o *BuyerInfo) GetDisbursedAmount() float32`

GetDisbursedAmount returns the DisbursedAmount field if non-nil, zero value otherwise.

### GetDisbursedAmountOk

`func (o *BuyerInfo) GetDisbursedAmountOk() (*float32, bool)`

GetDisbursedAmountOk returns a tuple with the DisbursedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursedAmount

`func (o *BuyerInfo) SetDisbursedAmount(v float32)`

SetDisbursedAmount sets DisbursedAmount field to given value.

### HasDisbursedAmount

`func (o *BuyerInfo) HasDisbursedAmount() bool`

HasDisbursedAmount returns a boolean if a field has been set.

### GetGcpBuyer

`func (o *BuyerInfo) GetGcpBuyer() GcpMarketplaceUserAccount`

GetGcpBuyer returns the GcpBuyer field if non-nil, zero value otherwise.

### GetGcpBuyerOk

`func (o *BuyerInfo) GetGcpBuyerOk() (*GcpMarketplaceUserAccount, bool)`

GetGcpBuyerOk returns a tuple with the GcpBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBuyer

`func (o *BuyerInfo) SetGcpBuyer(v GcpMarketplaceUserAccount)`

SetGcpBuyer sets GcpBuyer field to given value.

### HasGcpBuyer

`func (o *BuyerInfo) HasGcpBuyer() bool`

HasGcpBuyer returns a boolean if a field has been set.

### GetInvoicedAmount

`func (o *BuyerInfo) GetInvoicedAmount() float32`

GetInvoicedAmount returns the InvoicedAmount field if non-nil, zero value otherwise.

### GetInvoicedAmountOk

`func (o *BuyerInfo) GetInvoicedAmountOk() (*float32, bool)`

GetInvoicedAmountOk returns a tuple with the InvoicedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedAmount

`func (o *BuyerInfo) SetInvoicedAmount(v float32)`

SetInvoicedAmount sets InvoicedAmount field to given value.

### HasInvoicedAmount

`func (o *BuyerInfo) HasInvoicedAmount() bool`

HasInvoicedAmount returns a boolean if a field has been set.

### GetMetronomeCustomerId

`func (o *BuyerInfo) GetMetronomeCustomerId() string`

GetMetronomeCustomerId returns the MetronomeCustomerId field if non-nil, zero value otherwise.

### GetMetronomeCustomerIdOk

`func (o *BuyerInfo) GetMetronomeCustomerIdOk() (*string, bool)`

GetMetronomeCustomerIdOk returns a tuple with the MetronomeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeCustomerId

`func (o *BuyerInfo) SetMetronomeCustomerId(v string)`

SetMetronomeCustomerId sets MetronomeCustomerId field to given value.

### HasMetronomeCustomerId

`func (o *BuyerInfo) HasMetronomeCustomerId() bool`

HasMetronomeCustomerId returns a boolean if a field has been set.

### GetOrbCustomerId

`func (o *BuyerInfo) GetOrbCustomerId() string`

GetOrbCustomerId returns the OrbCustomerId field if non-nil, zero value otherwise.

### GetOrbCustomerIdOk

`func (o *BuyerInfo) GetOrbCustomerIdOk() (*string, bool)`

GetOrbCustomerIdOk returns a tuple with the OrbCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbCustomerId

`func (o *BuyerInfo) SetOrbCustomerId(v string)`

SetOrbCustomerId sets OrbCustomerId field to given value.

### HasOrbCustomerId

`func (o *BuyerInfo) HasOrbCustomerId() bool`

HasOrbCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


