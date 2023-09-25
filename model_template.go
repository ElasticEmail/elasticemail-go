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
	"time"
)

// checks if the Template type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Template{}

// Template Template info
type Template struct {
	TemplateType *TemplateType `json:"TemplateType,omitempty"`
	// Template name
	Name *string `json:"Name,omitempty"`
	// Date of creation in YYYY-MM-DDThh:ii:ss format
	DateAdded *time.Time `json:"DateAdded,omitempty"`
	// Default subject of email.
	Subject *string `json:"Subject,omitempty"`
	// Email content of this template
	Body []BodyPart `json:"Body,omitempty"`
	TemplateScope *TemplateScope `json:"TemplateScope,omitempty"`
}

// NewTemplate instantiates a new Template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplate() *Template {
	this := Template{}
	var templateType TemplateType = RAW_HTML
	this.TemplateType = &templateType
	var templateScope TemplateScope = PERSONAL
	this.TemplateScope = &templateScope
	return &this
}

// NewTemplateWithDefaults instantiates a new Template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateWithDefaults() *Template {
	this := Template{}
	var templateType TemplateType = RAW_HTML
	this.TemplateType = &templateType
	var templateScope TemplateScope = PERSONAL
	this.TemplateScope = &templateScope
	return &this
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *Template) GetTemplateType() TemplateType {
	if o == nil || isNil(o.TemplateType) {
		var ret TemplateType
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetTemplateTypeOk() (*TemplateType, bool) {
	if o == nil || isNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *Template) HasTemplateType() bool {
	if o != nil && !isNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given TemplateType and assigns it to the TemplateType field.
func (o *Template) SetTemplateType(v TemplateType) {
	o.TemplateType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Template) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Template) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Template) SetName(v string) {
	o.Name = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *Template) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *Template) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *Template) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Template) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Template) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Template) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Template) GetBody() []BodyPart {
	if o == nil || isNil(o.Body) {
		var ret []BodyPart
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetBodyOk() ([]BodyPart, bool) {
	if o == nil || isNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Template) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given []BodyPart and assigns it to the Body field.
func (o *Template) SetBody(v []BodyPart) {
	o.Body = v
}

// GetTemplateScope returns the TemplateScope field value if set, zero value otherwise.
func (o *Template) GetTemplateScope() TemplateScope {
	if o == nil || isNil(o.TemplateScope) {
		var ret TemplateScope
		return ret
	}
	return *o.TemplateScope
}

// GetTemplateScopeOk returns a tuple with the TemplateScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetTemplateScopeOk() (*TemplateScope, bool) {
	if o == nil || isNil(o.TemplateScope) {
		return nil, false
	}
	return o.TemplateScope, true
}

// HasTemplateScope returns a boolean if a field has been set.
func (o *Template) HasTemplateScope() bool {
	if o != nil && !isNil(o.TemplateScope) {
		return true
	}

	return false
}

// SetTemplateScope gets a reference to the given TemplateScope and assigns it to the TemplateScope field.
func (o *Template) SetTemplateScope(v TemplateScope) {
	o.TemplateScope = &v
}

func (o Template) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Template) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TemplateType) {
		toSerialize["TemplateType"] = o.TemplateType
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.DateAdded) {
		toSerialize["DateAdded"] = o.DateAdded
	}
	if !isNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !isNil(o.Body) {
		toSerialize["Body"] = o.Body
	}
	if !isNil(o.TemplateScope) {
		toSerialize["TemplateScope"] = o.TemplateScope
	}
	return toSerialize, nil
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


