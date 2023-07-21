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

// checks if the Campaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Campaign{}

// Campaign struct for Campaign
type Campaign struct {
	// Campaign's email content. Provide multiple items to send an A/X Split Campaign
	Content []CampaignTemplate `json:"Content,omitempty"`
	// Campaign name
	Name string `json:"Name"`
	Status *CampaignStatus `json:"Status,omitempty"`
	Recipients CampaignRecipient `json:"Recipients"`
	Options *CampaignOptions `json:"Options,omitempty"`
}

// NewCampaign instantiates a new Campaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaign(name string, recipients CampaignRecipient) *Campaign {
	this := Campaign{}
	this.Name = name
	var status CampaignStatus = CAMPAIGNSTATUS_DELETED
	this.Status = &status
	this.Recipients = recipients
	return &this
}

// NewCampaignWithDefaults instantiates a new Campaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignWithDefaults() *Campaign {
	this := Campaign{}
	var status CampaignStatus = CAMPAIGNSTATUS_DELETED
	this.Status = &status
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Campaign) GetContent() []CampaignTemplate {
	if o == nil || IsNil(o.Content) {
		var ret []CampaignTemplate
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetContentOk() ([]CampaignTemplate, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Campaign) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CampaignTemplate and assigns it to the Content field.
func (o *Campaign) SetContent(v []CampaignTemplate) {
	o.Content = v
}

// GetName returns the Name field value
func (o *Campaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Campaign) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Campaign) GetStatus() CampaignStatus {
	if o == nil || IsNil(o.Status) {
		var ret CampaignStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetStatusOk() (*CampaignStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Campaign) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CampaignStatus and assigns it to the Status field.
func (o *Campaign) SetStatus(v CampaignStatus) {
	o.Status = &v
}

// GetRecipients returns the Recipients field value
func (o *Campaign) GetRecipients() CampaignRecipient {
	if o == nil {
		var ret CampaignRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *Campaign) GetRecipientsOk() (*CampaignRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *Campaign) SetRecipients(v CampaignRecipient) {
	o.Recipients = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Campaign) GetOptions() CampaignOptions {
	if o == nil || IsNil(o.Options) {
		var ret CampaignOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Campaign) GetOptionsOk() (*CampaignOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Campaign) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given CampaignOptions and assigns it to the Options field.
func (o *Campaign) SetOptions(v CampaignOptions) {
	o.Options = &v
}

func (o Campaign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Campaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	toSerialize["Name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	toSerialize["Recipients"] = o.Recipients
	if !IsNil(o.Options) {
		toSerialize["Options"] = o.Options
	}
	return toSerialize, nil
}

type NullableCampaign struct {
	value *Campaign
	isSet bool
}

func (v NullableCampaign) Get() *Campaign {
	return v.value
}

func (v *NullableCampaign) Set(val *Campaign) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaign(val *Campaign) *NullableCampaign {
	return &NullableCampaign{value: val, isSet: true}
}

func (v NullableCampaign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


