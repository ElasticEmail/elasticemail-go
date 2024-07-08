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

// checks if the CampaignTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignTemplate{}

// CampaignTemplate Content of a Campaign
type CampaignTemplate struct {
	// Name of your custom IP Pool to be used in the sending process
	Poolname *string `json:"Poolname,omitempty"`
	// Your e-mail with an optional name (e.g.: John Doe <email@domain.com>)
	From string `json:"From"`
	// To what address should the recipients reply to (e.g. John Doe <email@domain.com>)
	ReplyTo *string `json:"ReplyTo,omitempty"`
	// Default subject of email.
	Subject *string `json:"Subject,omitempty"`
	// Name of template.
	TemplateName *string `json:"TemplateName,omitempty"`
	// Names of previously uploaded files that should be sent as downloadable attachments
	AttachFiles []string `json:"AttachFiles,omitempty"`
	Utm *Utm `json:"Utm,omitempty"`
}

type _CampaignTemplate CampaignTemplate

// NewCampaignTemplate instantiates a new CampaignTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignTemplate(from string) *CampaignTemplate {
	this := CampaignTemplate{}
	this.From = from
	return &this
}

// NewCampaignTemplateWithDefaults instantiates a new CampaignTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignTemplateWithDefaults() *CampaignTemplate {
	this := CampaignTemplate{}
	return &this
}

// GetPoolname returns the Poolname field value if set, zero value otherwise.
func (o *CampaignTemplate) GetPoolname() string {
	if o == nil || IsNil(o.Poolname) {
		var ret string
		return ret
	}
	return *o.Poolname
}

// GetPoolnameOk returns a tuple with the Poolname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetPoolnameOk() (*string, bool) {
	if o == nil || IsNil(o.Poolname) {
		return nil, false
	}
	return o.Poolname, true
}

// HasPoolname returns a boolean if a field has been set.
func (o *CampaignTemplate) HasPoolname() bool {
	if o != nil && !IsNil(o.Poolname) {
		return true
	}

	return false
}

// SetPoolname gets a reference to the given string and assigns it to the Poolname field.
func (o *CampaignTemplate) SetPoolname(v string) {
	o.Poolname = &v
}

// GetFrom returns the From field value
func (o *CampaignTemplate) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CampaignTemplate) SetFrom(v string) {
	o.From = v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *CampaignTemplate) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *CampaignTemplate) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given string and assigns it to the ReplyTo field.
func (o *CampaignTemplate) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CampaignTemplate) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CampaignTemplate) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CampaignTemplate) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *CampaignTemplate) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *CampaignTemplate) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *CampaignTemplate) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetAttachFiles returns the AttachFiles field value if set, zero value otherwise.
func (o *CampaignTemplate) GetAttachFiles() []string {
	if o == nil || IsNil(o.AttachFiles) {
		var ret []string
		return ret
	}
	return o.AttachFiles
}

// GetAttachFilesOk returns a tuple with the AttachFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetAttachFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachFiles) {
		return nil, false
	}
	return o.AttachFiles, true
}

// HasAttachFiles returns a boolean if a field has been set.
func (o *CampaignTemplate) HasAttachFiles() bool {
	if o != nil && !IsNil(o.AttachFiles) {
		return true
	}

	return false
}

// SetAttachFiles gets a reference to the given []string and assigns it to the AttachFiles field.
func (o *CampaignTemplate) SetAttachFiles(v []string) {
	o.AttachFiles = v
}

// GetUtm returns the Utm field value if set, zero value otherwise.
func (o *CampaignTemplate) GetUtm() Utm {
	if o == nil || IsNil(o.Utm) {
		var ret Utm
		return ret
	}
	return *o.Utm
}

// GetUtmOk returns a tuple with the Utm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplate) GetUtmOk() (*Utm, bool) {
	if o == nil || IsNil(o.Utm) {
		return nil, false
	}
	return o.Utm, true
}

// HasUtm returns a boolean if a field has been set.
func (o *CampaignTemplate) HasUtm() bool {
	if o != nil && !IsNil(o.Utm) {
		return true
	}

	return false
}

// SetUtm gets a reference to the given Utm and assigns it to the Utm field.
func (o *CampaignTemplate) SetUtm(v Utm) {
	o.Utm = &v
}

func (o CampaignTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Poolname) {
		toSerialize["Poolname"] = o.Poolname
	}
	toSerialize["From"] = o.From
	if !IsNil(o.ReplyTo) {
		toSerialize["ReplyTo"] = o.ReplyTo
	}
	if !IsNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !IsNil(o.TemplateName) {
		toSerialize["TemplateName"] = o.TemplateName
	}
	if !IsNil(o.AttachFiles) {
		toSerialize["AttachFiles"] = o.AttachFiles
	}
	if !IsNil(o.Utm) {
		toSerialize["Utm"] = o.Utm
	}
	return toSerialize, nil
}

func (o *CampaignTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"From",
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

	varCampaignTemplate := _CampaignTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignTemplate)

	if err != nil {
		return err
	}

	*o = CampaignTemplate(varCampaignTemplate)

	return err
}

type NullableCampaignTemplate struct {
	value *CampaignTemplate
	isSet bool
}

func (v NullableCampaignTemplate) Get() *CampaignTemplate {
	return v.value
}

func (v *NullableCampaignTemplate) Set(val *CampaignTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignTemplate(val *CampaignTemplate) *NullableCampaignTemplate {
	return &NullableCampaignTemplate{value: val, isSet: true}
}

func (v NullableCampaignTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


