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

// checks if the SupportTicketAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicketAttachment{}

// SupportTicketAttachment struct for SupportTicketAttachment
type SupportTicketAttachment struct {
	Date            *string `json:"date,omitempty"`
	Id              *string `json:"id,omitempty"`
	Mimetype        *string `json:"mimetype,omitempty"`
	Size            *int32  `json:"size,omitempty"`
	ThumbnailLarge  *string `json:"thumbnail_large,omitempty"`
	ThumbnailMedium *string `json:"thumbnail_medium,omitempty"`
	ThumbnailSmall  *string `json:"thumbnail_small,omitempty"`
	Title           *string `json:"title,omitempty"`
	Url             *string `json:"url,omitempty"`
}

// NewSupportTicketAttachment instantiates a new SupportTicketAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicketAttachment() *SupportTicketAttachment {
	this := SupportTicketAttachment{}
	return &this
}

// NewSupportTicketAttachmentWithDefaults instantiates a new SupportTicketAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketAttachmentWithDefaults() *SupportTicketAttachment {
	this := SupportTicketAttachment{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SupportTicketAttachment) SetDate(v string) {
	o.Date = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportTicketAttachment) SetId(v string) {
	o.Id = &v
}

// GetMimetype returns the Mimetype field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetMimetype() string {
	if o == nil || IsNil(o.Mimetype) {
		var ret string
		return ret
	}
	return *o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetMimetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Mimetype) {
		return nil, false
	}
	return o.Mimetype, true
}

// HasMimetype returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasMimetype() bool {
	if o != nil && !IsNil(o.Mimetype) {
		return true
	}

	return false
}

// SetMimetype gets a reference to the given string and assigns it to the Mimetype field.
func (o *SupportTicketAttachment) SetMimetype(v string) {
	o.Mimetype = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *SupportTicketAttachment) SetSize(v int32) {
	o.Size = &v
}

// GetThumbnailLarge returns the ThumbnailLarge field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetThumbnailLarge() string {
	if o == nil || IsNil(o.ThumbnailLarge) {
		var ret string
		return ret
	}
	return *o.ThumbnailLarge
}

// GetThumbnailLargeOk returns a tuple with the ThumbnailLarge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetThumbnailLargeOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailLarge) {
		return nil, false
	}
	return o.ThumbnailLarge, true
}

// HasThumbnailLarge returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasThumbnailLarge() bool {
	if o != nil && !IsNil(o.ThumbnailLarge) {
		return true
	}

	return false
}

// SetThumbnailLarge gets a reference to the given string and assigns it to the ThumbnailLarge field.
func (o *SupportTicketAttachment) SetThumbnailLarge(v string) {
	o.ThumbnailLarge = &v
}

// GetThumbnailMedium returns the ThumbnailMedium field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetThumbnailMedium() string {
	if o == nil || IsNil(o.ThumbnailMedium) {
		var ret string
		return ret
	}
	return *o.ThumbnailMedium
}

// GetThumbnailMediumOk returns a tuple with the ThumbnailMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetThumbnailMediumOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailMedium) {
		return nil, false
	}
	return o.ThumbnailMedium, true
}

// HasThumbnailMedium returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasThumbnailMedium() bool {
	if o != nil && !IsNil(o.ThumbnailMedium) {
		return true
	}

	return false
}

// SetThumbnailMedium gets a reference to the given string and assigns it to the ThumbnailMedium field.
func (o *SupportTicketAttachment) SetThumbnailMedium(v string) {
	o.ThumbnailMedium = &v
}

// GetThumbnailSmall returns the ThumbnailSmall field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetThumbnailSmall() string {
	if o == nil || IsNil(o.ThumbnailSmall) {
		var ret string
		return ret
	}
	return *o.ThumbnailSmall
}

// GetThumbnailSmallOk returns a tuple with the ThumbnailSmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetThumbnailSmallOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailSmall) {
		return nil, false
	}
	return o.ThumbnailSmall, true
}

// HasThumbnailSmall returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasThumbnailSmall() bool {
	if o != nil && !IsNil(o.ThumbnailSmall) {
		return true
	}

	return false
}

// SetThumbnailSmall gets a reference to the given string and assigns it to the ThumbnailSmall field.
func (o *SupportTicketAttachment) SetThumbnailSmall(v string) {
	o.ThumbnailSmall = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SupportTicketAttachment) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SupportTicketAttachment) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketAttachment) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SupportTicketAttachment) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SupportTicketAttachment) SetUrl(v string) {
	o.Url = &v
}

func (o SupportTicketAttachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicketAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mimetype) {
		toSerialize["mimetype"] = o.Mimetype
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.ThumbnailLarge) {
		toSerialize["thumbnail_large"] = o.ThumbnailLarge
	}
	if !IsNil(o.ThumbnailMedium) {
		toSerialize["thumbnail_medium"] = o.ThumbnailMedium
	}
	if !IsNil(o.ThumbnailSmall) {
		toSerialize["thumbnail_small"] = o.ThumbnailSmall
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSupportTicketAttachment struct {
	value *SupportTicketAttachment
	isSet bool
}

func (v NullableSupportTicketAttachment) Get() *SupportTicketAttachment {
	return v.value
}

func (v *NullableSupportTicketAttachment) Set(val *SupportTicketAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketAttachment(val *SupportTicketAttachment) *NullableSupportTicketAttachment {
	return &NullableSupportTicketAttachment{value: val, isSet: true}
}

func (v NullableSupportTicketAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
