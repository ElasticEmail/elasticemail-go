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

// checks if the EmailJobFailedStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailJobFailedStatus{}

// EmailJobFailedStatus struct for EmailJobFailedStatus
type EmailJobFailedStatus struct {
	Address *string `json:"Address,omitempty"`
	Error *string `json:"Error,omitempty"`
	// RFC Error code
	ErrorCode *int32 `json:"ErrorCode,omitempty"`
	Category *string `json:"Category,omitempty"`
}

// NewEmailJobFailedStatus instantiates a new EmailJobFailedStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailJobFailedStatus() *EmailJobFailedStatus {
	this := EmailJobFailedStatus{}
	return &this
}

// NewEmailJobFailedStatusWithDefaults instantiates a new EmailJobFailedStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailJobFailedStatusWithDefaults() *EmailJobFailedStatus {
	this := EmailJobFailedStatus{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *EmailJobFailedStatus) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailJobFailedStatus) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *EmailJobFailedStatus) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *EmailJobFailedStatus) SetAddress(v string) {
	o.Address = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *EmailJobFailedStatus) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailJobFailedStatus) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *EmailJobFailedStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *EmailJobFailedStatus) SetError(v string) {
	o.Error = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *EmailJobFailedStatus) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailJobFailedStatus) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *EmailJobFailedStatus) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *EmailJobFailedStatus) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *EmailJobFailedStatus) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailJobFailedStatus) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *EmailJobFailedStatus) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *EmailJobFailedStatus) SetCategory(v string) {
	o.Category = &v
}

func (o EmailJobFailedStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailJobFailedStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	return toSerialize, nil
}

type NullableEmailJobFailedStatus struct {
	value *EmailJobFailedStatus
	isSet bool
}

func (v NullableEmailJobFailedStatus) Get() *EmailJobFailedStatus {
	return v.value
}

func (v *NullableEmailJobFailedStatus) Set(val *EmailJobFailedStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailJobFailedStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailJobFailedStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailJobFailedStatus(val *EmailJobFailedStatus) *NullableEmailJobFailedStatus {
	return &NullableEmailJobFailedStatus{value: val, isSet: true}
}

func (v NullableEmailJobFailedStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailJobFailedStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

