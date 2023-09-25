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

// checks if the ListPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPayload{}

// ListPayload struct for ListPayload
type ListPayload struct {
	// Name of your list.
	ListName string `json:"ListName"`
	// True: Allow unsubscribing from this list. Otherwise, false
	AllowUnsubscribe *bool `json:"AllowUnsubscribe,omitempty"`
	// Comma delimited list of existing contact emails that should be added to this list. Leave empty for all contacts
	Emails []string `json:"Emails,omitempty"`
}

// NewListPayload instantiates a new ListPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPayload(listName string) *ListPayload {
	this := ListPayload{}
	this.ListName = listName
	return &this
}

// NewListPayloadWithDefaults instantiates a new ListPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPayloadWithDefaults() *ListPayload {
	this := ListPayload{}
	return &this
}

// GetListName returns the ListName field value
func (o *ListPayload) GetListName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListName
}

// GetListNameOk returns a tuple with the ListName field value
// and a boolean to check if the value has been set.
func (o *ListPayload) GetListNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListName, true
}

// SetListName sets field value
func (o *ListPayload) SetListName(v string) {
	o.ListName = v
}

// GetAllowUnsubscribe returns the AllowUnsubscribe field value if set, zero value otherwise.
func (o *ListPayload) GetAllowUnsubscribe() bool {
	if o == nil || isNil(o.AllowUnsubscribe) {
		var ret bool
		return ret
	}
	return *o.AllowUnsubscribe
}

// GetAllowUnsubscribeOk returns a tuple with the AllowUnsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPayload) GetAllowUnsubscribeOk() (*bool, bool) {
	if o == nil || isNil(o.AllowUnsubscribe) {
		return nil, false
	}
	return o.AllowUnsubscribe, true
}

// HasAllowUnsubscribe returns a boolean if a field has been set.
func (o *ListPayload) HasAllowUnsubscribe() bool {
	if o != nil && !isNil(o.AllowUnsubscribe) {
		return true
	}

	return false
}

// SetAllowUnsubscribe gets a reference to the given bool and assigns it to the AllowUnsubscribe field.
func (o *ListPayload) SetAllowUnsubscribe(v bool) {
	o.AllowUnsubscribe = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ListPayload) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPayload) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ListPayload) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *ListPayload) SetEmails(v []string) {
	o.Emails = v
}

func (o ListPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ListName"] = o.ListName
	if !isNil(o.AllowUnsubscribe) {
		toSerialize["AllowUnsubscribe"] = o.AllowUnsubscribe
	}
	if !isNil(o.Emails) {
		toSerialize["Emails"] = o.Emails
	}
	return toSerialize, nil
}

type NullableListPayload struct {
	value *ListPayload
	isSet bool
}

func (v NullableListPayload) Get() *ListPayload {
	return v.value
}

func (v *NullableListPayload) Set(val *ListPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableListPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableListPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPayload(val *ListPayload) *NullableListPayload {
	return &NullableListPayload{value: val, isSet: true}
}

func (v NullableListPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


