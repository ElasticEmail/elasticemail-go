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

// DeliveryOptimizationType the model 'DeliveryOptimizationType'
type DeliveryOptimizationType string

// List of DeliveryOptimizationType
const (
	NONE DeliveryOptimizationType = "None"
	TO_ENGAGED_FIRST DeliveryOptimizationType = "ToEngagedFirst"
	BY_OPEN_TIME DeliveryOptimizationType = "ByOpenTime"
)

// All allowed values of DeliveryOptimizationType enum
var AllowedDeliveryOptimizationTypeEnumValues = []DeliveryOptimizationType{
	"None",
	"ToEngagedFirst",
	"ByOpenTime",
}

func (v *DeliveryOptimizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryOptimizationType(value)
	for _, existing := range AllowedDeliveryOptimizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryOptimizationType", value)
}

// NewDeliveryOptimizationTypeFromValue returns a pointer to a valid DeliveryOptimizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryOptimizationTypeFromValue(v string) (*DeliveryOptimizationType, error) {
	ev := DeliveryOptimizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryOptimizationType: valid values are %v", v, AllowedDeliveryOptimizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryOptimizationType) IsValid() bool {
	for _, existing := range AllowedDeliveryOptimizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryOptimizationType value
func (v DeliveryOptimizationType) Ptr() *DeliveryOptimizationType {
	return &v
}

type NullableDeliveryOptimizationType struct {
	value *DeliveryOptimizationType
	isSet bool
}

func (v NullableDeliveryOptimizationType) Get() *DeliveryOptimizationType {
	return v.value
}

func (v *NullableDeliveryOptimizationType) Set(val *DeliveryOptimizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOptimizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOptimizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOptimizationType(val *DeliveryOptimizationType) *NullableDeliveryOptimizationType {
	return &NullableDeliveryOptimizationType{value: val, isSet: true}
}

func (v NullableDeliveryOptimizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOptimizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

