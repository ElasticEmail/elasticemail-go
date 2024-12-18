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
	"fmt"
)

// TrackingType HTTP or HTTPS Protocal used for link tracking.
type TrackingType string

// List of TrackingType
const (
	NONE TrackingType = "None"
	DELETE TrackingType = "Delete"
	HTTP TrackingType = "Http"
	EXTERNAL_HTTPS TrackingType = "ExternalHttps"
	INTERNAL_CERT_HTTPS TrackingType = "InternalCertHttps"
	LETS_ENCRYPT_CERT TrackingType = "LetsEncryptCert"
)

// All allowed values of TrackingType enum
var AllowedTrackingTypeEnumValues = []TrackingType{
	"None",
	"Delete",
	"Http",
	"ExternalHttps",
	"InternalCertHttps",
	"LetsEncryptCert",
}

func (v *TrackingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackingType(value)
	for _, existing := range AllowedTrackingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackingType", value)
}

// NewTrackingTypeFromValue returns a pointer to a valid TrackingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackingTypeFromValue(v string) (*TrackingType, error) {
	ev := TrackingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackingType: valid values are %v", v, AllowedTrackingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackingType) IsValid() bool {
	for _, existing := range AllowedTrackingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrackingType value
func (v TrackingType) Ptr() *TrackingType {
	return &v
}

type NullableTrackingType struct {
	value *TrackingType
	isSet bool
}

func (v NullableTrackingType) Get() *TrackingType {
	return v.value
}

func (v *NullableTrackingType) Set(val *TrackingType) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingType) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingType(val *TrackingType) *NullableTrackingType {
	return &NullableTrackingType{value: val, isSet: true}
}

func (v NullableTrackingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

