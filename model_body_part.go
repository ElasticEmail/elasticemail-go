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
	"bytes"
	"fmt"
)

// checks if the BodyPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BodyPart{}

// BodyPart Email body part with user-provided MIME type (text/html, text/plain, etc)
type BodyPart struct {
	ContentType BodyContentType `json:"ContentType"`
	// Actual content of the body part
	Content *string `json:"Content,omitempty"`
	// Text value of charset encoding for example: iso-8859-1, windows-1251, utf-8, us-ascii, windows-1250 and more...
	Charset *string `json:"Charset,omitempty"`
}

type _BodyPart BodyPart

// NewBodyPart instantiates a new BodyPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodyPart(contentType BodyContentType) *BodyPart {
	this := BodyPart{}
	this.ContentType = contentType
	return &this
}

// NewBodyPartWithDefaults instantiates a new BodyPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodyPartWithDefaults() *BodyPart {
	this := BodyPart{}
	var contentType BodyContentType = HTML
	this.ContentType = contentType
	return &this
}

// GetContentType returns the ContentType field value
func (o *BodyPart) GetContentType() BodyContentType {
	if o == nil {
		var ret BodyContentType
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *BodyPart) GetContentTypeOk() (*BodyContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *BodyPart) SetContentType(v BodyContentType) {
	o.ContentType = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BodyPart) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BodyPart) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *BodyPart) SetContent(v string) {
	o.Content = &v
}

// GetCharset returns the Charset field value if set, zero value otherwise.
func (o *BodyPart) GetCharset() string {
	if o == nil || IsNil(o.Charset) {
		var ret string
		return ret
	}
	return *o.Charset
}

// GetCharsetOk returns a tuple with the Charset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodyPart) GetCharsetOk() (*string, bool) {
	if o == nil || IsNil(o.Charset) {
		return nil, false
	}
	return o.Charset, true
}

// HasCharset returns a boolean if a field has been set.
func (o *BodyPart) HasCharset() bool {
	if o != nil && !IsNil(o.Charset) {
		return true
	}

	return false
}

// SetCharset gets a reference to the given string and assigns it to the Charset field.
func (o *BodyPart) SetCharset(v string) {
	o.Charset = &v
}

func (o BodyPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BodyPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ContentType"] = o.ContentType
	if !IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !IsNil(o.Charset) {
		toSerialize["Charset"] = o.Charset
	}
	return toSerialize, nil
}

func (o *BodyPart) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ContentType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBodyPart := _BodyPart{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBodyPart)

	if err != nil {
		return err
	}

	*o = BodyPart(varBodyPart)

	return err
}

type NullableBodyPart struct {
	value *BodyPart
	isSet bool
}

func (v NullableBodyPart) Get() *BodyPart {
	return v.value
}

func (v *NullableBodyPart) Set(val *BodyPart) {
	v.value = val
	v.isSet = true
}

func (v NullableBodyPart) IsSet() bool {
	return v.isSet
}

func (v *NullableBodyPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodyPart(val *BodyPart) *NullableBodyPart {
	return &NullableBodyPart{value: val, isSet: true}
}

func (v NullableBodyPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodyPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


