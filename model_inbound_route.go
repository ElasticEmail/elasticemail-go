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

// checks if the InboundRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundRoute{}

// InboundRoute struct for InboundRoute
type InboundRoute struct {
	PublicId *string `json:"PublicId,omitempty"`
	// Name of this route
	Name *string `json:"Name,omitempty"`
	FilterType *InboundRouteFilterType `json:"FilterType,omitempty"`
	// Filter of the inbound data
	Filter *string `json:"Filter,omitempty"`
	ActionType *InboundRouteActionType `json:"ActionType,omitempty"`
	// URL address or Email to notify about the inbound
	ActionParameter *string `json:"ActionParameter,omitempty"`
	// Place of this route in your routes queue's order
	SortOrder *int32 `json:"SortOrder,omitempty"`
}

// NewInboundRoute instantiates a new InboundRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundRoute() *InboundRoute {
	this := InboundRoute{}
	var filterType InboundRouteFilterType = EMAIL_ADDRESS
	this.FilterType = &filterType
	var actionType InboundRouteActionType = FORWARD_TO_EMAIL
	this.ActionType = &actionType
	return &this
}

// NewInboundRouteWithDefaults instantiates a new InboundRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundRouteWithDefaults() *InboundRoute {
	this := InboundRoute{}
	var filterType InboundRouteFilterType = EMAIL_ADDRESS
	this.FilterType = &filterType
	var actionType InboundRouteActionType = FORWARD_TO_EMAIL
	this.ActionType = &actionType
	return &this
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *InboundRoute) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *InboundRoute) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *InboundRoute) SetPublicId(v string) {
	o.PublicId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InboundRoute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InboundRoute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InboundRoute) SetName(v string) {
	o.Name = &v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *InboundRoute) GetFilterType() InboundRouteFilterType {
	if o == nil || IsNil(o.FilterType) {
		var ret InboundRouteFilterType
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetFilterTypeOk() (*InboundRouteFilterType, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *InboundRoute) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given InboundRouteFilterType and assigns it to the FilterType field.
func (o *InboundRoute) SetFilterType(v InboundRouteFilterType) {
	o.FilterType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *InboundRoute) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *InboundRoute) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *InboundRoute) SetFilter(v string) {
	o.Filter = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *InboundRoute) GetActionType() InboundRouteActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret InboundRouteActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetActionTypeOk() (*InboundRouteActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *InboundRoute) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given InboundRouteActionType and assigns it to the ActionType field.
func (o *InboundRoute) SetActionType(v InboundRouteActionType) {
	o.ActionType = &v
}

// GetActionParameter returns the ActionParameter field value if set, zero value otherwise.
func (o *InboundRoute) GetActionParameter() string {
	if o == nil || IsNil(o.ActionParameter) {
		var ret string
		return ret
	}
	return *o.ActionParameter
}

// GetActionParameterOk returns a tuple with the ActionParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetActionParameterOk() (*string, bool) {
	if o == nil || IsNil(o.ActionParameter) {
		return nil, false
	}
	return o.ActionParameter, true
}

// HasActionParameter returns a boolean if a field has been set.
func (o *InboundRoute) HasActionParameter() bool {
	if o != nil && !IsNil(o.ActionParameter) {
		return true
	}

	return false
}

// SetActionParameter gets a reference to the given string and assigns it to the ActionParameter field.
func (o *InboundRoute) SetActionParameter(v string) {
	o.ActionParameter = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *InboundRoute) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundRoute) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *InboundRoute) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *InboundRoute) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o InboundRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicId) {
		toSerialize["PublicId"] = o.PublicId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.FilterType) {
		toSerialize["FilterType"] = o.FilterType
	}
	if !IsNil(o.Filter) {
		toSerialize["Filter"] = o.Filter
	}
	if !IsNil(o.ActionType) {
		toSerialize["ActionType"] = o.ActionType
	}
	if !IsNil(o.ActionParameter) {
		toSerialize["ActionParameter"] = o.ActionParameter
	}
	if !IsNil(o.SortOrder) {
		toSerialize["SortOrder"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableInboundRoute struct {
	value *InboundRoute
	isSet bool
}

func (v NullableInboundRoute) Get() *InboundRoute {
	return v.value
}

func (v *NullableInboundRoute) Set(val *InboundRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundRoute(val *InboundRoute) *NullableInboundRoute {
	return &NullableInboundRoute{value: val, isSet: true}
}

func (v NullableInboundRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


