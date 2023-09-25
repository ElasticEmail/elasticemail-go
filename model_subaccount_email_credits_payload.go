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

// checks if the SubaccountEmailCreditsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountEmailCreditsPayload{}

// SubaccountEmailCreditsPayload A change to SubAccount email credits pool, with an additional note.
type SubaccountEmailCreditsPayload struct {
	// Positive or negative value; this will be added or subtracted from Subaccount's current email Credits pool.
	Credits int32 `json:"Credits"`
	// Note to append to this credits change, for history.
	Notes *string `json:"Notes,omitempty"`
}

// NewSubaccountEmailCreditsPayload instantiates a new SubaccountEmailCreditsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountEmailCreditsPayload(credits int32) *SubaccountEmailCreditsPayload {
	this := SubaccountEmailCreditsPayload{}
	this.Credits = credits
	return &this
}

// NewSubaccountEmailCreditsPayloadWithDefaults instantiates a new SubaccountEmailCreditsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountEmailCreditsPayloadWithDefaults() *SubaccountEmailCreditsPayload {
	this := SubaccountEmailCreditsPayload{}
	return &this
}

// GetCredits returns the Credits field value
func (o *SubaccountEmailCreditsPayload) GetCredits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *SubaccountEmailCreditsPayload) GetCreditsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *SubaccountEmailCreditsPayload) SetCredits(v int32) {
	o.Credits = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SubaccountEmailCreditsPayload) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailCreditsPayload) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SubaccountEmailCreditsPayload) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SubaccountEmailCreditsPayload) SetNotes(v string) {
	o.Notes = &v
}

func (o SubaccountEmailCreditsPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountEmailCreditsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Credits"] = o.Credits
	if !isNil(o.Notes) {
		toSerialize["Notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableSubaccountEmailCreditsPayload struct {
	value *SubaccountEmailCreditsPayload
	isSet bool
}

func (v NullableSubaccountEmailCreditsPayload) Get() *SubaccountEmailCreditsPayload {
	return v.value
}

func (v *NullableSubaccountEmailCreditsPayload) Set(val *SubaccountEmailCreditsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountEmailCreditsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountEmailCreditsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountEmailCreditsPayload(val *SubaccountEmailCreditsPayload) *NullableSubaccountEmailCreditsPayload {
	return &NullableSubaccountEmailCreditsPayload{value: val, isSet: true}
}

func (v NullableSubaccountEmailCreditsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountEmailCreditsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


