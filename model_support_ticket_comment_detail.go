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

// checks if the SupportTicketCommentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicketCommentDetail{}

// SupportTicketCommentDetail struct for SupportTicketCommentDetail
type SupportTicketCommentDetail struct {
	Attachment *SupportTicketAttachment `json:"attachment,omitempty"`
	Frame      *SupportTicketFrame      `json:"frame,omitempty"`
	Image      *SupportTicketImage      `json:"image,omitempty"`
	Text       *string                  `json:"text,omitempty"`
	Type       *string                  `json:"type,omitempty"`
}

// NewSupportTicketCommentDetail instantiates a new SupportTicketCommentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicketCommentDetail() *SupportTicketCommentDetail {
	this := SupportTicketCommentDetail{}
	return &this
}

// NewSupportTicketCommentDetailWithDefaults instantiates a new SupportTicketCommentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketCommentDetailWithDefaults() *SupportTicketCommentDetail {
	this := SupportTicketCommentDetail{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *SupportTicketCommentDetail) GetAttachment() SupportTicketAttachment {
	if o == nil || IsNil(o.Attachment) {
		var ret SupportTicketAttachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketCommentDetail) GetAttachmentOk() (*SupportTicketAttachment, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *SupportTicketCommentDetail) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given SupportTicketAttachment and assigns it to the Attachment field.
func (o *SupportTicketCommentDetail) SetAttachment(v SupportTicketAttachment) {
	o.Attachment = &v
}

// GetFrame returns the Frame field value if set, zero value otherwise.
func (o *SupportTicketCommentDetail) GetFrame() SupportTicketFrame {
	if o == nil || IsNil(o.Frame) {
		var ret SupportTicketFrame
		return ret
	}
	return *o.Frame
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketCommentDetail) GetFrameOk() (*SupportTicketFrame, bool) {
	if o == nil || IsNil(o.Frame) {
		return nil, false
	}
	return o.Frame, true
}

// HasFrame returns a boolean if a field has been set.
func (o *SupportTicketCommentDetail) HasFrame() bool {
	if o != nil && !IsNil(o.Frame) {
		return true
	}

	return false
}

// SetFrame gets a reference to the given SupportTicketFrame and assigns it to the Frame field.
func (o *SupportTicketCommentDetail) SetFrame(v SupportTicketFrame) {
	o.Frame = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SupportTicketCommentDetail) GetImage() SupportTicketImage {
	if o == nil || IsNil(o.Image) {
		var ret SupportTicketImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketCommentDetail) GetImageOk() (*SupportTicketImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SupportTicketCommentDetail) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given SupportTicketImage and assigns it to the Image field.
func (o *SupportTicketCommentDetail) SetImage(v SupportTicketImage) {
	o.Image = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SupportTicketCommentDetail) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketCommentDetail) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SupportTicketCommentDetail) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SupportTicketCommentDetail) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SupportTicketCommentDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicketCommentDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SupportTicketCommentDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SupportTicketCommentDetail) SetType(v string) {
	o.Type = &v
}

func (o SupportTicketCommentDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicketCommentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	if !IsNil(o.Frame) {
		toSerialize["frame"] = o.Frame
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSupportTicketCommentDetail struct {
	value *SupportTicketCommentDetail
	isSet bool
}

func (v NullableSupportTicketCommentDetail) Get() *SupportTicketCommentDetail {
	return v.value
}

func (v *NullableSupportTicketCommentDetail) Set(val *SupportTicketCommentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketCommentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketCommentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketCommentDetail(val *SupportTicketCommentDetail) *NullableSupportTicketCommentDetail {
	return &NullableSupportTicketCommentDetail{value: val, isSet: true}
}

func (v NullableSupportTicketCommentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketCommentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
