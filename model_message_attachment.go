/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
)

// checks if the MessageAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttachment{}

// MessageAttachment struct for MessageAttachment
type MessageAttachment struct {
	// File's content as byte array (or a Base64 string)
	BinaryContent string `json:"BinaryContent"`
	// Display name of the file
	Name string `json:"Name"`
	// MIME content type
	ContentType *string `json:"ContentType,omitempty"`
	// Size of your attachment (in bytes).
	Size *int32 `json:"Size,omitempty"`
}

// NewMessageAttachment instantiates a new MessageAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttachment(binaryContent string, name string) *MessageAttachment {
	this := MessageAttachment{}
	this.BinaryContent = binaryContent
	this.Name = name
	return &this
}

// NewMessageAttachmentWithDefaults instantiates a new MessageAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttachmentWithDefaults() *MessageAttachment {
	this := MessageAttachment{}
	return &this
}

// GetBinaryContent returns the BinaryContent field value
func (o *MessageAttachment) GetBinaryContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryContent
}

// GetBinaryContentOk returns a tuple with the BinaryContent field value
// and a boolean to check if the value has been set.
func (o *MessageAttachment) GetBinaryContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinaryContent, true
}

// SetBinaryContent sets field value
func (o *MessageAttachment) SetBinaryContent(v string) {
	o.BinaryContent = v
}

// GetName returns the Name field value
func (o *MessageAttachment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageAttachment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageAttachment) SetName(v string) {
	o.Name = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *MessageAttachment) GetContentType() string {
	if o == nil || isNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil || isNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MessageAttachment) HasContentType() bool {
	if o != nil && !isNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *MessageAttachment) SetContentType(v string) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MessageAttachment) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttachment) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MessageAttachment) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MessageAttachment) SetSize(v int32) {
	o.Size = &v
}

func (o MessageAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BinaryContent"] = o.BinaryContent
	toSerialize["Name"] = o.Name
	if !isNil(o.ContentType) {
		toSerialize["ContentType"] = o.ContentType
	}
	if !isNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	return toSerialize, nil
}

type NullableMessageAttachment struct {
	value *MessageAttachment
	isSet bool
}

func (v NullableMessageAttachment) Get() *MessageAttachment {
	return v.value
}

func (v *NullableMessageAttachment) Set(val *MessageAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttachment(val *MessageAttachment) *NullableMessageAttachment {
	return &NullableMessageAttachment{value: val, isSet: true}
}

func (v NullableMessageAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


