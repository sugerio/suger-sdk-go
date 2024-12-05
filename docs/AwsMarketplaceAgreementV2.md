# AwsMarketplaceAgreementV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceTime** | Pointer to **time.Time** |  | [optional] 
**AgreementId** | Pointer to **string** | AWS Marketplace Agreement Id | [optional] 
**AgreementType** | Pointer to **string** |  | [optional] 
**BuyerAccountId** | Pointer to **string** | The AWS Account Id of the buyer in AWS Marketplace | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**OfferId** | Pointer to **string** | AWS Marketplace Offer Id | [optional] 
**ProductId** | Pointer to **string** | AWS Marketplace Product Id | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**SellerAccountId** | Pointer to **string** | The AWS Account Id of the seller in AWS Marketplace | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**AwsMarketplaceAgreementStatus**](AwsMarketplaceAgreementStatus.md) |  | [optional] 

## Methods

### NewAwsMarketplaceAgreementV2

`func NewAwsMarketplaceAgreementV2() *AwsMarketplaceAgreementV2`

NewAwsMarketplaceAgreementV2 instantiates a new AwsMarketplaceAgreementV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceAgreementV2WithDefaults

`func NewAwsMarketplaceAgreementV2WithDefaults() *AwsMarketplaceAgreementV2`

NewAwsMarketplaceAgreementV2WithDefaults instantiates a new AwsMarketplaceAgreementV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceTime

`func (o *AwsMarketplaceAgreementV2) GetAcceptanceTime() time.Time`

GetAcceptanceTime returns the AcceptanceTime field if non-nil, zero value otherwise.

### GetAcceptanceTimeOk

`func (o *AwsMarketplaceAgreementV2) GetAcceptanceTimeOk() (*time.Time, bool)`

GetAcceptanceTimeOk returns a tuple with the AcceptanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceTime

`func (o *AwsMarketplaceAgreementV2) SetAcceptanceTime(v time.Time)`

SetAcceptanceTime sets AcceptanceTime field to given value.

### HasAcceptanceTime

`func (o *AwsMarketplaceAgreementV2) HasAcceptanceTime() bool`

HasAcceptanceTime returns a boolean if a field has been set.

### GetAgreementId

`func (o *AwsMarketplaceAgreementV2) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AwsMarketplaceAgreementV2) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AwsMarketplaceAgreementV2) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AwsMarketplaceAgreementV2) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetAgreementType

`func (o *AwsMarketplaceAgreementV2) GetAgreementType() string`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *AwsMarketplaceAgreementV2) GetAgreementTypeOk() (*string, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *AwsMarketplaceAgreementV2) SetAgreementType(v string)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *AwsMarketplaceAgreementV2) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetBuyerAccountId

`func (o *AwsMarketplaceAgreementV2) GetBuyerAccountId() string`

GetBuyerAccountId returns the BuyerAccountId field if non-nil, zero value otherwise.

### GetBuyerAccountIdOk

`func (o *AwsMarketplaceAgreementV2) GetBuyerAccountIdOk() (*string, bool)`

GetBuyerAccountIdOk returns a tuple with the BuyerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAccountId

`func (o *AwsMarketplaceAgreementV2) SetBuyerAccountId(v string)`

SetBuyerAccountId sets BuyerAccountId field to given value.

### HasBuyerAccountId

`func (o *AwsMarketplaceAgreementV2) HasBuyerAccountId() bool`

HasBuyerAccountId returns a boolean if a field has been set.

### GetEndTime

`func (o *AwsMarketplaceAgreementV2) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AwsMarketplaceAgreementV2) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AwsMarketplaceAgreementV2) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AwsMarketplaceAgreementV2) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetOfferId

`func (o *AwsMarketplaceAgreementV2) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *AwsMarketplaceAgreementV2) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *AwsMarketplaceAgreementV2) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *AwsMarketplaceAgreementV2) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetProductId

`func (o *AwsMarketplaceAgreementV2) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AwsMarketplaceAgreementV2) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AwsMarketplaceAgreementV2) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AwsMarketplaceAgreementV2) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductType

`func (o *AwsMarketplaceAgreementV2) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AwsMarketplaceAgreementV2) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AwsMarketplaceAgreementV2) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AwsMarketplaceAgreementV2) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSellerAccountId

`func (o *AwsMarketplaceAgreementV2) GetSellerAccountId() string`

GetSellerAccountId returns the SellerAccountId field if non-nil, zero value otherwise.

### GetSellerAccountIdOk

`func (o *AwsMarketplaceAgreementV2) GetSellerAccountIdOk() (*string, bool)`

GetSellerAccountIdOk returns a tuple with the SellerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAccountId

`func (o *AwsMarketplaceAgreementV2) SetSellerAccountId(v string)`

SetSellerAccountId sets SellerAccountId field to given value.

### HasSellerAccountId

`func (o *AwsMarketplaceAgreementV2) HasSellerAccountId() bool`

HasSellerAccountId returns a boolean if a field has been set.

### GetStartTime

`func (o *AwsMarketplaceAgreementV2) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AwsMarketplaceAgreementV2) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AwsMarketplaceAgreementV2) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AwsMarketplaceAgreementV2) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *AwsMarketplaceAgreementV2) GetStatus() AwsMarketplaceAgreementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsMarketplaceAgreementV2) GetStatusOk() (*AwsMarketplaceAgreementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsMarketplaceAgreementV2) SetStatus(v AwsMarketplaceAgreementStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsMarketplaceAgreementV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


