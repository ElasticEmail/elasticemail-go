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

// checks if the Utm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Utm{}

// Utm Utm marketing data to be attached to every link in this e-mail.
type Utm struct {
	// utmsource value
	Source *string `json:"Source,omitempty"`
	// utmmedium value
	Medium *string `json:"Medium,omitempty"`
	// utmcampaign value
	Campaign *string `json:"Campaign,omitempty"`
	// utmcontent value
	Content *string `json:"Content,omitempty"`
}

// NewUtm instantiates a new Utm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtm() *Utm {
	this := Utm{}
	return &this
}

// NewUtmWithDefaults instantiates a new Utm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtmWithDefaults() *Utm {
	this := Utm{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Utm) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utm) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Utm) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Utm) SetSource(v string) {
	o.Source = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *Utm) GetMedium() string {
	if o == nil || IsNil(o.Medium) {
		var ret string
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utm) GetMediumOk() (*string, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *Utm) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given string and assigns it to the Medium field.
func (o *Utm) SetMedium(v string) {
	o.Medium = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *Utm) GetCampaign() string {
	if o == nil || IsNil(o.Campaign) {
		var ret string
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utm) GetCampaignOk() (*string, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *Utm) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given string and assigns it to the Campaign field.
func (o *Utm) SetCampaign(v string) {
	o.Campaign = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Utm) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Utm) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Utm) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Utm) SetContent(v string) {
	o.Content = &v
}

func (o Utm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Utm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["Source"] = o.Source
	}
	if !IsNil(o.Medium) {
		toSerialize["Medium"] = o.Medium
	}
	if !IsNil(o.Campaign) {
		toSerialize["Campaign"] = o.Campaign
	}
	if !IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	return toSerialize, nil
}

type NullableUtm struct {
	value *Utm
	isSet bool
}

func (v NullableUtm) Get() *Utm {
	return v.value
}

func (v *NullableUtm) Set(val *Utm) {
	v.value = val
	v.isSet = true
}

func (v NullableUtm) IsSet() bool {
	return v.isSet
}

func (v *NullableUtm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtm(val *Utm) *NullableUtm {
	return &NullableUtm{value: val, isSet: true}
}

func (v NullableUtm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


