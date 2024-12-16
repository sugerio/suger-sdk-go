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

// checks if the OrbBillableMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbBillableMetric{}

// OrbBillableMetric struct for OrbBillableMetric
type OrbBillableMetric struct {
	Description *string                  `json:"description,omitempty"`
	Id          *string                  `json:"id,omitempty"`
	Metadata    *map[string]string       `json:"metadata,omitempty"`
	Name        *string                  `json:"name,omitempty"`
	Status      *OrbBillableMetricStatus `json:"status,omitempty"`
}

// NewOrbBillableMetric instantiates a new OrbBillableMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbBillableMetric() *OrbBillableMetric {
	this := OrbBillableMetric{}
	return &this
}

// NewOrbBillableMetricWithDefaults instantiates a new OrbBillableMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbBillableMetricWithDefaults() *OrbBillableMetric {
	this := OrbBillableMetric{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrbBillableMetric) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbBillableMetric) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrbBillableMetric) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrbBillableMetric) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrbBillableMetric) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbBillableMetric) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrbBillableMetric) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrbBillableMetric) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrbBillableMetric) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbBillableMetric) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrbBillableMetric) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OrbBillableMetric) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrbBillableMetric) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbBillableMetric) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrbBillableMetric) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrbBillableMetric) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrbBillableMetric) GetStatus() OrbBillableMetricStatus {
	if o == nil || IsNil(o.Status) {
		var ret OrbBillableMetricStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbBillableMetric) GetStatusOk() (*OrbBillableMetricStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrbBillableMetric) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrbBillableMetricStatus and assigns it to the Status field.
func (o *OrbBillableMetric) SetStatus(v OrbBillableMetricStatus) {
	o.Status = &v
}

func (o OrbBillableMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbBillableMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOrbBillableMetric struct {
	value *OrbBillableMetric
	isSet bool
}

func (v NullableOrbBillableMetric) Get() *OrbBillableMetric {
	return v.value
}

func (v *NullableOrbBillableMetric) Set(val *OrbBillableMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbBillableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbBillableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbBillableMetric(val *OrbBillableMetric) *NullableOrbBillableMetric {
	return &NullableOrbBillableMetric{value: val, isSet: true}
}

func (v NullableOrbBillableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbBillableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
