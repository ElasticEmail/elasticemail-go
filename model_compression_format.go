/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// CompressionFormat FileResponse compression format
type CompressionFormat string

// List of CompressionFormat
const (
	NONE CompressionFormat = "None"
	ZIP CompressionFormat = "Zip"
)

// All allowed values of CompressionFormat enum
var AllowedCompressionFormatEnumValues = []CompressionFormat{
	"None",
	"Zip",
}

func (v *CompressionFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompressionFormat(value)
	for _, existing := range AllowedCompressionFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompressionFormat", value)
}

// NewCompressionFormatFromValue returns a pointer to a valid CompressionFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompressionFormatFromValue(v string) (*CompressionFormat, error) {
	ev := CompressionFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompressionFormat: valid values are %v", v, AllowedCompressionFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompressionFormat) IsValid() bool {
	for _, existing := range AllowedCompressionFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompressionFormat value
func (v CompressionFormat) Ptr() *CompressionFormat {
	return &v
}

type NullableCompressionFormat struct {
	value *CompressionFormat
	isSet bool
}

func (v NullableCompressionFormat) Get() *CompressionFormat {
	return v.value
}

func (v *NullableCompressionFormat) Set(val *CompressionFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCompressionFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCompressionFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompressionFormat(val *CompressionFormat) *NullableCompressionFormat {
	return &NullableCompressionFormat{value: val, isSet: true}
}

func (v NullableCompressionFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompressionFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

