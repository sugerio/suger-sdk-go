# GcpMarketplacePrivateOfferReplacementMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CotermAlignment** | Pointer to **string** |  | [optional] 
**ReplacedOffer** | Pointer to **string** | The resource name of the private offer being replaced. in format of \&quot;projects/{projectNumber}/services/{productServiceName}/privateOffers/{privateOfferId}\&quot; | [optional] 
**ReplacedOfferEndTime** | Pointer to **time.Time** |  | [optional] 
**ReplacementPolicy** | Pointer to **string** |  | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferReplacementMetadata

`func NewGcpMarketplacePrivateOfferReplacementMetadata() *GcpMarketplacePrivateOfferReplacementMetadata`

NewGcpMarketplacePrivateOfferReplacementMetadata instantiates a new GcpMarketplacePrivateOfferReplacementMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferReplacementMetadataWithDefaults

`func NewGcpMarketplacePrivateOfferReplacementMetadataWithDefaults() *GcpMarketplacePrivateOfferReplacementMetadata`

NewGcpMarketplacePrivateOfferReplacementMetadataWithDefaults instantiates a new GcpMarketplacePrivateOfferReplacementMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCotermAlignment

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetCotermAlignment() string`

GetCotermAlignment returns the CotermAlignment field if non-nil, zero value otherwise.

### GetCotermAlignmentOk

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetCotermAlignmentOk() (*string, bool)`

GetCotermAlignmentOk returns a tuple with the CotermAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCotermAlignment

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) SetCotermAlignment(v string)`

SetCotermAlignment sets CotermAlignment field to given value.

### HasCotermAlignment

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) HasCotermAlignment() bool`

HasCotermAlignment returns a boolean if a field has been set.

### GetReplacedOffer

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacedOffer() string`

GetReplacedOffer returns the ReplacedOffer field if non-nil, zero value otherwise.

### GetReplacedOfferOk

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacedOfferOk() (*string, bool)`

GetReplacedOfferOk returns a tuple with the ReplacedOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedOffer

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) SetReplacedOffer(v string)`

SetReplacedOffer sets ReplacedOffer field to given value.

### HasReplacedOffer

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) HasReplacedOffer() bool`

HasReplacedOffer returns a boolean if a field has been set.

### GetReplacedOfferEndTime

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacedOfferEndTime() time.Time`

GetReplacedOfferEndTime returns the ReplacedOfferEndTime field if non-nil, zero value otherwise.

### GetReplacedOfferEndTimeOk

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacedOfferEndTimeOk() (*time.Time, bool)`

GetReplacedOfferEndTimeOk returns a tuple with the ReplacedOfferEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedOfferEndTime

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) SetReplacedOfferEndTime(v time.Time)`

SetReplacedOfferEndTime sets ReplacedOfferEndTime field to given value.

### HasReplacedOfferEndTime

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) HasReplacedOfferEndTime() bool`

HasReplacedOfferEndTime returns a boolean if a field has been set.

### GetReplacementPolicy

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacementPolicy() string`

GetReplacementPolicy returns the ReplacementPolicy field if non-nil, zero value otherwise.

### GetReplacementPolicyOk

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) GetReplacementPolicyOk() (*string, bool)`

GetReplacementPolicyOk returns a tuple with the ReplacementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementPolicy

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) SetReplacementPolicy(v string)`

SetReplacementPolicy sets ReplacementPolicy field to given value.

### HasReplacementPolicy

`func (o *GcpMarketplacePrivateOfferReplacementMetadata) HasReplacementPolicy() bool`

HasReplacementPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


