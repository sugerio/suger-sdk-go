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
	"time"
)

// checks if the WorkloadEntitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadEntitlement{}

// WorkloadEntitlement struct for WorkloadEntitlement
type WorkloadEntitlement struct {
	BuyerID      *string    `json:"buyerID,omitempty"`
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// nullable
	EndTime           *time.Time         `json:"endTime,omitempty"`
	EntitlementTermID *string            `json:"entitlementTermID,omitempty"`
	ExternalBuyerID   *string            `json:"externalBuyerID,omitempty"`
	ExternalID        *string            `json:"externalID,omitempty"`
	ExternalProductID *string            `json:"externalProductID,omitempty"`
	Id                *string            `json:"id,omitempty"`
	Info              *EntitlementInfo   `json:"info,omitempty"`
	LastUpdateTime    *time.Time         `json:"lastUpdateTime,omitempty"`
	MetaInfo          *WorkloadMetaInfo  `json:"metaInfo,omitempty"`
	Name              *string            `json:"name,omitempty"`
	OfferID           *string            `json:"offerID,omitempty"`
	OrganizationID    *string            `json:"organizationID,omitempty"`
	Partner           *Partner           `json:"partner,omitempty"`
	ProductID         *string            `json:"productID,omitempty"`
	Service           *PartnerService    `json:"service,omitempty"`
	StartTime         *time.Time         `json:"startTime,omitempty"`
	Status            *EntitlementStatus `json:"status,omitempty"`
	Type              *string            `json:"type,omitempty"`
}

// NewWorkloadEntitlement instantiates a new WorkloadEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadEntitlement() *WorkloadEntitlement {
	this := WorkloadEntitlement{}
	return &this
}

// NewWorkloadEntitlementWithDefaults instantiates a new WorkloadEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadEntitlementWithDefaults() *WorkloadEntitlement {
	this := WorkloadEntitlement{}
	return &this
}

// GetBuyerID returns the BuyerID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetBuyerID() string {
	if o == nil || IsNil(o.BuyerID) {
		var ret string
		return ret
	}
	return *o.BuyerID
}

// GetBuyerIDOk returns a tuple with the BuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerID) {
		return nil, false
	}
	return o.BuyerID, true
}

// HasBuyerID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasBuyerID() bool {
	if o != nil && !IsNil(o.BuyerID) {
		return true
	}

	return false
}

// SetBuyerID gets a reference to the given string and assigns it to the BuyerID field.
func (o *WorkloadEntitlement) SetBuyerID(v string) {
	o.BuyerID = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *WorkloadEntitlement) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WorkloadEntitlement) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetEntitlementTermID returns the EntitlementTermID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetEntitlementTermID() string {
	if o == nil || IsNil(o.EntitlementTermID) {
		var ret string
		return ret
	}
	return *o.EntitlementTermID
}

// GetEntitlementTermIDOk returns a tuple with the EntitlementTermID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetEntitlementTermIDOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementTermID) {
		return nil, false
	}
	return o.EntitlementTermID, true
}

// HasEntitlementTermID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasEntitlementTermID() bool {
	if o != nil && !IsNil(o.EntitlementTermID) {
		return true
	}

	return false
}

// SetEntitlementTermID gets a reference to the given string and assigns it to the EntitlementTermID field.
func (o *WorkloadEntitlement) SetEntitlementTermID(v string) {
	o.EntitlementTermID = &v
}

// GetExternalBuyerID returns the ExternalBuyerID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetExternalBuyerID() string {
	if o == nil || IsNil(o.ExternalBuyerID) {
		var ret string
		return ret
	}
	return *o.ExternalBuyerID
}

// GetExternalBuyerIDOk returns a tuple with the ExternalBuyerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetExternalBuyerIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalBuyerID) {
		return nil, false
	}
	return o.ExternalBuyerID, true
}

// HasExternalBuyerID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasExternalBuyerID() bool {
	if o != nil && !IsNil(o.ExternalBuyerID) {
		return true
	}

	return false
}

// SetExternalBuyerID gets a reference to the given string and assigns it to the ExternalBuyerID field.
func (o *WorkloadEntitlement) SetExternalBuyerID(v string) {
	o.ExternalBuyerID = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *WorkloadEntitlement) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetExternalProductID returns the ExternalProductID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetExternalProductID() string {
	if o == nil || IsNil(o.ExternalProductID) {
		var ret string
		return ret
	}
	return *o.ExternalProductID
}

// GetExternalProductIDOk returns a tuple with the ExternalProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetExternalProductIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalProductID) {
		return nil, false
	}
	return o.ExternalProductID, true
}

// HasExternalProductID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasExternalProductID() bool {
	if o != nil && !IsNil(o.ExternalProductID) {
		return true
	}

	return false
}

// SetExternalProductID gets a reference to the given string and assigns it to the ExternalProductID field.
func (o *WorkloadEntitlement) SetExternalProductID(v string) {
	o.ExternalProductID = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkloadEntitlement) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetInfo() EntitlementInfo {
	if o == nil || IsNil(o.Info) {
		var ret EntitlementInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetInfoOk() (*EntitlementInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given EntitlementInfo and assigns it to the Info field.
func (o *WorkloadEntitlement) SetInfo(v EntitlementInfo) {
	o.Info = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetLastUpdateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetLastUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given time.Time and assigns it to the LastUpdateTime field.
func (o *WorkloadEntitlement) SetLastUpdateTime(v time.Time) {
	o.LastUpdateTime = &v
}

// GetMetaInfo returns the MetaInfo field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetMetaInfo() WorkloadMetaInfo {
	if o == nil || IsNil(o.MetaInfo) {
		var ret WorkloadMetaInfo
		return ret
	}
	return *o.MetaInfo
}

// GetMetaInfoOk returns a tuple with the MetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetMetaInfoOk() (*WorkloadMetaInfo, bool) {
	if o == nil || IsNil(o.MetaInfo) {
		return nil, false
	}
	return o.MetaInfo, true
}

// HasMetaInfo returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasMetaInfo() bool {
	if o != nil && !IsNil(o.MetaInfo) {
		return true
	}

	return false
}

// SetMetaInfo gets a reference to the given WorkloadMetaInfo and assigns it to the MetaInfo field.
func (o *WorkloadEntitlement) SetMetaInfo(v WorkloadMetaInfo) {
	o.MetaInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkloadEntitlement) SetName(v string) {
	o.Name = &v
}

// GetOfferID returns the OfferID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetOfferID() string {
	if o == nil || IsNil(o.OfferID) {
		var ret string
		return ret
	}
	return *o.OfferID
}

// GetOfferIDOk returns a tuple with the OfferID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetOfferIDOk() (*string, bool) {
	if o == nil || IsNil(o.OfferID) {
		return nil, false
	}
	return o.OfferID, true
}

// HasOfferID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasOfferID() bool {
	if o != nil && !IsNil(o.OfferID) {
		return true
	}

	return false
}

// SetOfferID gets a reference to the given string and assigns it to the OfferID field.
func (o *WorkloadEntitlement) SetOfferID(v string) {
	o.OfferID = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *WorkloadEntitlement) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetPartner() Partner {
	if o == nil || IsNil(o.Partner) {
		var ret Partner
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetPartnerOk() (*Partner, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given Partner and assigns it to the Partner field.
func (o *WorkloadEntitlement) SetPartner(v Partner) {
	o.Partner = &v
}

// GetProductID returns the ProductID field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetProductID() string {
	if o == nil || IsNil(o.ProductID) {
		var ret string
		return ret
	}
	return *o.ProductID
}

// GetProductIDOk returns a tuple with the ProductID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetProductIDOk() (*string, bool) {
	if o == nil || IsNil(o.ProductID) {
		return nil, false
	}
	return o.ProductID, true
}

// HasProductID returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasProductID() bool {
	if o != nil && !IsNil(o.ProductID) {
		return true
	}

	return false
}

// SetProductID gets a reference to the given string and assigns it to the ProductID field.
func (o *WorkloadEntitlement) SetProductID(v string) {
	o.ProductID = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetService() PartnerService {
	if o == nil || IsNil(o.Service) {
		var ret PartnerService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetServiceOk() (*PartnerService, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given PartnerService and assigns it to the Service field.
func (o *WorkloadEntitlement) SetService(v PartnerService) {
	o.Service = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WorkloadEntitlement) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetStatus() EntitlementStatus {
	if o == nil || IsNil(o.Status) {
		var ret EntitlementStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetStatusOk() (*EntitlementStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EntitlementStatus and assigns it to the Status field.
func (o *WorkloadEntitlement) SetStatus(v EntitlementStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkloadEntitlement) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEntitlement) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkloadEntitlement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkloadEntitlement) SetType(v string) {
	o.Type = &v
}

func (o WorkloadEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadEntitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerID) {
		toSerialize["buyerID"] = o.BuyerID
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.EntitlementTermID) {
		toSerialize["entitlementTermID"] = o.EntitlementTermID
	}
	if !IsNil(o.ExternalBuyerID) {
		toSerialize["externalBuyerID"] = o.ExternalBuyerID
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.ExternalProductID) {
		toSerialize["externalProductID"] = o.ExternalProductID
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.MetaInfo) {
		toSerialize["metaInfo"] = o.MetaInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OfferID) {
		toSerialize["offerID"] = o.OfferID
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.ProductID) {
		toSerialize["productID"] = o.ProductID
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWorkloadEntitlement struct {
	value *WorkloadEntitlement
	isSet bool
}

func (v NullableWorkloadEntitlement) Get() *WorkloadEntitlement {
	return v.value
}

func (v *NullableWorkloadEntitlement) Set(val *WorkloadEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadEntitlement(val *WorkloadEntitlement) *NullableWorkloadEntitlement {
	return &NullableWorkloadEntitlement{value: val, isSet: true}
}

func (v NullableWorkloadEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
