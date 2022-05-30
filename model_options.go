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
)

// Options E-mail configuration
type Options struct {
	// By how long should an e-mail be delayed (in minutes). Maximum is 35 days.
	TimeOffset NullableInt32 `json:"TimeOffset,omitempty"`
	// Name of your custom IP Pool to be used in the sending process
	PoolName *string `json:"PoolName,omitempty"`
	// Name of selected channel.
	ChannelName *string `json:"ChannelName,omitempty"`
	Encoding *EncodingType `json:"Encoding,omitempty"`
	// Should the opens be tracked? If no value has been provided, Account's default setting will be used.
	TrackOpens *bool `json:"TrackOpens,omitempty"`
	// Should the clicks be tracked? If no value has been provided, Account's default setting will be used.
	TrackClicks *bool `json:"TrackClicks,omitempty"`
}

// NewOptions instantiates a new Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptions() *Options {
	this := Options{}
	var encoding EncodingType = USER_PROVIDED
	this.Encoding = &encoding
	return &this
}

// NewOptionsWithDefaults instantiates a new Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsWithDefaults() *Options {
	this := Options{}
	var encoding EncodingType = USER_PROVIDED
	this.Encoding = &encoding
	return &this
}

// GetTimeOffset returns the TimeOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Options) GetTimeOffset() int32 {
	if o == nil || o.TimeOffset.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TimeOffset.Get()
}

// GetTimeOffsetOk returns a tuple with the TimeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Options) GetTimeOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeOffset.Get(), o.TimeOffset.IsSet()
}

// HasTimeOffset returns a boolean if a field has been set.
func (o *Options) HasTimeOffset() bool {
	if o != nil && o.TimeOffset.IsSet() {
		return true
	}

	return false
}

// SetTimeOffset gets a reference to the given NullableInt32 and assigns it to the TimeOffset field.
func (o *Options) SetTimeOffset(v int32) {
	o.TimeOffset.Set(&v)
}
// SetTimeOffsetNil sets the value for TimeOffset to be an explicit nil
func (o *Options) SetTimeOffsetNil() {
	o.TimeOffset.Set(nil)
}

// UnsetTimeOffset ensures that no value is present for TimeOffset, not even an explicit nil
func (o *Options) UnsetTimeOffset() {
	o.TimeOffset.Unset()
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *Options) GetPoolName() string {
	if o == nil || o.PoolName == nil {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetPoolNameOk() (*string, bool) {
	if o == nil || o.PoolName == nil {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *Options) HasPoolName() bool {
	if o != nil && o.PoolName != nil {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *Options) SetPoolName(v string) {
	o.PoolName = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *Options) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *Options) HasChannelName() bool {
	if o != nil && o.ChannelName != nil {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *Options) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *Options) GetEncoding() EncodingType {
	if o == nil || o.Encoding == nil {
		var ret EncodingType
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetEncodingOk() (*EncodingType, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *Options) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given EncodingType and assigns it to the Encoding field.
func (o *Options) SetEncoding(v EncodingType) {
	o.Encoding = &v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *Options) GetTrackOpens() bool {
	if o == nil || o.TrackOpens == nil {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetTrackOpensOk() (*bool, bool) {
	if o == nil || o.TrackOpens == nil {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *Options) HasTrackOpens() bool {
	if o != nil && o.TrackOpens != nil {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *Options) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *Options) GetTrackClicks() bool {
	if o == nil || o.TrackClicks == nil {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetTrackClicksOk() (*bool, bool) {
	if o == nil || o.TrackClicks == nil {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *Options) HasTrackClicks() bool {
	if o != nil && o.TrackClicks != nil {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *Options) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

func (o Options) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TimeOffset.IsSet() {
		toSerialize["TimeOffset"] = o.TimeOffset.Get()
	}
	if o.PoolName != nil {
		toSerialize["PoolName"] = o.PoolName
	}
	if o.ChannelName != nil {
		toSerialize["ChannelName"] = o.ChannelName
	}
	if o.Encoding != nil {
		toSerialize["Encoding"] = o.Encoding
	}
	if o.TrackOpens != nil {
		toSerialize["TrackOpens"] = o.TrackOpens
	}
	if o.TrackClicks != nil {
		toSerialize["TrackClicks"] = o.TrackClicks
	}
	return json.Marshal(toSerialize)
}

type NullableOptions struct {
	value *Options
	isSet bool
}

func (v NullableOptions) Get() *Options {
	return v.value
}

func (v *NullableOptions) Set(val *Options) {
	v.value = val
	v.isSet = true
}

func (v NullableOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptions(val *Options) *NullableOptions {
	return &NullableOptions{value: val, isSet: true}
}

func (v NullableOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


