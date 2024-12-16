/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suger

import (
	"encoding/json"
)

// checks if the AwsMarketplaceEventBridgeEventDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceEventBridgeEventDetail{}

// AwsMarketplaceEventBridgeEventDetail struct for AwsMarketplaceEventBridgeEventDetail
type AwsMarketplaceEventBridgeEventDetail struct {
	Catalog         *string `json:"catalog,omitempty"`
	EventCategory   *string `json:"eventCategory,omitempty"`
	EventID         *string `json:"eventID,omitempty"`
	EventName       *string `json:"eventName,omitempty"`
	EventSource     *string `json:"eventSource,omitempty"`
	EventType       *string `json:"eventType,omitempty"`
	EventVersion    *string `json:"eventVersion,omitempty"`
	ManagementEvent *bool   `json:"managementEvent,omitempty"`
	// The seller/ISV's AWS Account Id.
	Manufacturer      *AwsMarketplaceEventBridgeEventAccount `json:"manufacturer,omitempty"`
	Offer             *AwsMarketplaceEventBridgeEventOffer   `json:"offer,omitempty"`
	Product           *AwsMarketplaceEventBridgeEventProduct `json:"product,omitempty"`
	RequestID         *string                                `json:"requestID,omitempty"`
	RequestParameters map[string]interface{}                 `json:"requestParameters,omitempty"`
	ResponseElements  map[string]interface{}                 `json:"responseElements,omitempty"`
	// For private offer created by a channel partner, this is the channel partner's AWS Account Id. For private offer created by a seller/ISV, this is the seller/ISV's AWS Account Id.
	SellerOfRecord          *AwsMarketplaceEventBridgeEventAccount `json:"sellerOfRecord,omitempty"`
	TargetedBuyerAccountIds []string                               `json:"targetedBuyerAccountIds,omitempty"`
}

// NewAwsMarketplaceEventBridgeEventDetail instantiates a new AwsMarketplaceEventBridgeEventDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceEventBridgeEventDetail() *AwsMarketplaceEventBridgeEventDetail {
	this := AwsMarketplaceEventBridgeEventDetail{}
	return &this
}

// NewAwsMarketplaceEventBridgeEventDetailWithDefaults instantiates a new AwsMarketplaceEventBridgeEventDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceEventBridgeEventDetailWithDefaults() *AwsMarketplaceEventBridgeEventDetail {
	this := AwsMarketplaceEventBridgeEventDetail{}
	return &this
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetCatalog() string {
	if o == nil || IsNil(o.Catalog) {
		var ret string
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetCatalogOk() (*string, bool) {
	if o == nil || IsNil(o.Catalog) {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasCatalog() bool {
	if o != nil && !IsNil(o.Catalog) {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given string and assigns it to the Catalog field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetCatalog(v string) {
	o.Catalog = &v
}

// GetEventCategory returns the EventCategory field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventCategory() string {
	if o == nil || IsNil(o.EventCategory) {
		var ret string
		return ret
	}
	return *o.EventCategory
}

// GetEventCategoryOk returns a tuple with the EventCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.EventCategory) {
		return nil, false
	}
	return o.EventCategory, true
}

// HasEventCategory returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventCategory() bool {
	if o != nil && !IsNil(o.EventCategory) {
		return true
	}

	return false
}

// SetEventCategory gets a reference to the given string and assigns it to the EventCategory field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventCategory(v string) {
	o.EventCategory = &v
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventID() string {
	if o == nil || IsNil(o.EventID) {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventIDOk() (*string, bool) {
	if o == nil || IsNil(o.EventID) {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventID() bool {
	if o != nil && !IsNil(o.EventID) {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventID(v string) {
	o.EventID = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventName(v string) {
	o.EventName = &v
}

// GetEventSource returns the EventSource field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventSource() string {
	if o == nil || IsNil(o.EventSource) {
		var ret string
		return ret
	}
	return *o.EventSource
}

// GetEventSourceOk returns a tuple with the EventSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventSourceOk() (*string, bool) {
	if o == nil || IsNil(o.EventSource) {
		return nil, false
	}
	return o.EventSource, true
}

// HasEventSource returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventSource() bool {
	if o != nil && !IsNil(o.EventSource) {
		return true
	}

	return false
}

// SetEventSource gets a reference to the given string and assigns it to the EventSource field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventSource(v string) {
	o.EventSource = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventType(v string) {
	o.EventType = &v
}

// GetEventVersion returns the EventVersion field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventVersion() string {
	if o == nil || IsNil(o.EventVersion) {
		var ret string
		return ret
	}
	return *o.EventVersion
}

// GetEventVersionOk returns a tuple with the EventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetEventVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EventVersion) {
		return nil, false
	}
	return o.EventVersion, true
}

// HasEventVersion returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasEventVersion() bool {
	if o != nil && !IsNil(o.EventVersion) {
		return true
	}

	return false
}

// SetEventVersion gets a reference to the given string and assigns it to the EventVersion field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetEventVersion(v string) {
	o.EventVersion = &v
}

// GetManagementEvent returns the ManagementEvent field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetManagementEvent() bool {
	if o == nil || IsNil(o.ManagementEvent) {
		var ret bool
		return ret
	}
	return *o.ManagementEvent
}

// GetManagementEventOk returns a tuple with the ManagementEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetManagementEventOk() (*bool, bool) {
	if o == nil || IsNil(o.ManagementEvent) {
		return nil, false
	}
	return o.ManagementEvent, true
}

// HasManagementEvent returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasManagementEvent() bool {
	if o != nil && !IsNil(o.ManagementEvent) {
		return true
	}

	return false
}

// SetManagementEvent gets a reference to the given bool and assigns it to the ManagementEvent field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetManagementEvent(v bool) {
	o.ManagementEvent = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetManufacturer() AwsMarketplaceEventBridgeEventAccount {
	if o == nil || IsNil(o.Manufacturer) {
		var ret AwsMarketplaceEventBridgeEventAccount
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetManufacturerOk() (*AwsMarketplaceEventBridgeEventAccount, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given AwsMarketplaceEventBridgeEventAccount and assigns it to the Manufacturer field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetManufacturer(v AwsMarketplaceEventBridgeEventAccount) {
	o.Manufacturer = &v
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetOffer() AwsMarketplaceEventBridgeEventOffer {
	if o == nil || IsNil(o.Offer) {
		var ret AwsMarketplaceEventBridgeEventOffer
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetOfferOk() (*AwsMarketplaceEventBridgeEventOffer, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given AwsMarketplaceEventBridgeEventOffer and assigns it to the Offer field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetOffer(v AwsMarketplaceEventBridgeEventOffer) {
	o.Offer = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetProduct() AwsMarketplaceEventBridgeEventProduct {
	if o == nil || IsNil(o.Product) {
		var ret AwsMarketplaceEventBridgeEventProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetProductOk() (*AwsMarketplaceEventBridgeEventProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AwsMarketplaceEventBridgeEventProduct and assigns it to the Product field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetProduct(v AwsMarketplaceEventBridgeEventProduct) {
	o.Product = &v
}

// GetRequestID returns the RequestID field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestID() string {
	if o == nil || IsNil(o.RequestID) {
		var ret string
		return ret
	}
	return *o.RequestID
}

// GetRequestIDOk returns a tuple with the RequestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestIDOk() (*string, bool) {
	if o == nil || IsNil(o.RequestID) {
		return nil, false
	}
	return o.RequestID, true
}

// HasRequestID returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasRequestID() bool {
	if o != nil && !IsNil(o.RequestID) {
		return true
	}

	return false
}

// SetRequestID gets a reference to the given string and assigns it to the RequestID field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetRequestID(v string) {
	o.RequestID = &v
}

// GetRequestParameters returns the RequestParameters field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestParameters() map[string]interface{} {
	if o == nil || IsNil(o.RequestParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.RequestParameters
}

// GetRequestParametersOk returns a tuple with the RequestParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetRequestParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RequestParameters) {
		return map[string]interface{}{}, false
	}
	return o.RequestParameters, true
}

// HasRequestParameters returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasRequestParameters() bool {
	if o != nil && !IsNil(o.RequestParameters) {
		return true
	}

	return false
}

// SetRequestParameters gets a reference to the given map[string]interface{} and assigns it to the RequestParameters field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetRequestParameters(v map[string]interface{}) {
	o.RequestParameters = v
}

// GetResponseElements returns the ResponseElements field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetResponseElements() map[string]interface{} {
	if o == nil || IsNil(o.ResponseElements) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResponseElements
}

// GetResponseElementsOk returns a tuple with the ResponseElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetResponseElementsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResponseElements) {
		return map[string]interface{}{}, false
	}
	return o.ResponseElements, true
}

// HasResponseElements returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasResponseElements() bool {
	if o != nil && !IsNil(o.ResponseElements) {
		return true
	}

	return false
}

// SetResponseElements gets a reference to the given map[string]interface{} and assigns it to the ResponseElements field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetResponseElements(v map[string]interface{}) {
	o.ResponseElements = v
}

// GetSellerOfRecord returns the SellerOfRecord field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetSellerOfRecord() AwsMarketplaceEventBridgeEventAccount {
	if o == nil || IsNil(o.SellerOfRecord) {
		var ret AwsMarketplaceEventBridgeEventAccount
		return ret
	}
	return *o.SellerOfRecord
}

// GetSellerOfRecordOk returns a tuple with the SellerOfRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetSellerOfRecordOk() (*AwsMarketplaceEventBridgeEventAccount, bool) {
	if o == nil || IsNil(o.SellerOfRecord) {
		return nil, false
	}
	return o.SellerOfRecord, true
}

// HasSellerOfRecord returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasSellerOfRecord() bool {
	if o != nil && !IsNil(o.SellerOfRecord) {
		return true
	}

	return false
}

// SetSellerOfRecord gets a reference to the given AwsMarketplaceEventBridgeEventAccount and assigns it to the SellerOfRecord field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetSellerOfRecord(v AwsMarketplaceEventBridgeEventAccount) {
	o.SellerOfRecord = &v
}

// GetTargetedBuyerAccountIds returns the TargetedBuyerAccountIds field value if set, zero value otherwise.
func (o *AwsMarketplaceEventBridgeEventDetail) GetTargetedBuyerAccountIds() []string {
	if o == nil || IsNil(o.TargetedBuyerAccountIds) {
		var ret []string
		return ret
	}
	return o.TargetedBuyerAccountIds
}

// GetTargetedBuyerAccountIdsOk returns a tuple with the TargetedBuyerAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) GetTargetedBuyerAccountIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetedBuyerAccountIds) {
		return nil, false
	}
	return o.TargetedBuyerAccountIds, true
}

// HasTargetedBuyerAccountIds returns a boolean if a field has been set.
func (o *AwsMarketplaceEventBridgeEventDetail) HasTargetedBuyerAccountIds() bool {
	if o != nil && !IsNil(o.TargetedBuyerAccountIds) {
		return true
	}

	return false
}

// SetTargetedBuyerAccountIds gets a reference to the given []string and assigns it to the TargetedBuyerAccountIds field.
func (o *AwsMarketplaceEventBridgeEventDetail) SetTargetedBuyerAccountIds(v []string) {
	o.TargetedBuyerAccountIds = v
}

func (o AwsMarketplaceEventBridgeEventDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceEventBridgeEventDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Catalog) {
		toSerialize["catalog"] = o.Catalog
	}
	if !IsNil(o.EventCategory) {
		toSerialize["eventCategory"] = o.EventCategory
	}
	if !IsNil(o.EventID) {
		toSerialize["eventID"] = o.EventID
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.EventSource) {
		toSerialize["eventSource"] = o.EventSource
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EventVersion) {
		toSerialize["eventVersion"] = o.EventVersion
	}
	if !IsNil(o.ManagementEvent) {
		toSerialize["managementEvent"] = o.ManagementEvent
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.RequestID) {
		toSerialize["requestID"] = o.RequestID
	}
	if !IsNil(o.RequestParameters) {
		toSerialize["requestParameters"] = o.RequestParameters
	}
	if !IsNil(o.ResponseElements) {
		toSerialize["responseElements"] = o.ResponseElements
	}
	if !IsNil(o.SellerOfRecord) {
		toSerialize["sellerOfRecord"] = o.SellerOfRecord
	}
	if !IsNil(o.TargetedBuyerAccountIds) {
		toSerialize["targetedBuyerAccountIds"] = o.TargetedBuyerAccountIds
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceEventBridgeEventDetail struct {
	value *AwsMarketplaceEventBridgeEventDetail
	isSet bool
}

func (v NullableAwsMarketplaceEventBridgeEventDetail) Get() *AwsMarketplaceEventBridgeEventDetail {
	return v.value
}

func (v *NullableAwsMarketplaceEventBridgeEventDetail) Set(val *AwsMarketplaceEventBridgeEventDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceEventBridgeEventDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceEventBridgeEventDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceEventBridgeEventDetail(val *AwsMarketplaceEventBridgeEventDetail) *NullableAwsMarketplaceEventBridgeEventDetail {
	return &NullableAwsMarketplaceEventBridgeEventDetail{value: val, isSet: true}
}

func (v NullableAwsMarketplaceEventBridgeEventDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceEventBridgeEventDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
