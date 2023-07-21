# Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOffset** | Pointer to **NullableInt32** | By how long should an e-mail be delayed (in minutes). Maximum is 35 days. | [optional] 
**PoolName** | Pointer to **string** | Name of your custom IP Pool to be used in the sending process | [optional] 
**ChannelName** | Pointer to **string** | Name of selected channel. | [optional] 
**Encoding** | Pointer to [**EncodingType**](EncodingType.md) |  | [optional] [default to ENCODINGTYPE_USER_PROVIDED]
**TrackOpens** | Pointer to **bool** | Should the opens be tracked? If no value has been provided, Account&#39;s default setting will be used. | [optional] 
**TrackClicks** | Pointer to **bool** | Should the clicks be tracked? If no value has been provided, Account&#39;s default setting will be used. | [optional] 

## Methods

### NewOptions

`func NewOptions() *Options`

NewOptions instantiates a new Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsWithDefaults

`func NewOptionsWithDefaults() *Options`

NewOptionsWithDefaults instantiates a new Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOffset

`func (o *Options) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *Options) GetTimeOffsetOk() (*int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *Options) SetTimeOffset(v int32)`

SetTimeOffset sets TimeOffset field to given value.

### HasTimeOffset

`func (o *Options) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### SetTimeOffsetNil

`func (o *Options) SetTimeOffsetNil(b bool)`

 SetTimeOffsetNil sets the value for TimeOffset to be an explicit nil

### UnsetTimeOffset
`func (o *Options) UnsetTimeOffset()`

UnsetTimeOffset ensures that no value is present for TimeOffset, not even an explicit nil
### GetPoolName

`func (o *Options) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *Options) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *Options) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *Options) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetChannelName

`func (o *Options) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *Options) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *Options) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *Options) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetEncoding

`func (o *Options) GetEncoding() EncodingType`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *Options) GetEncodingOk() (*EncodingType, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *Options) SetEncoding(v EncodingType)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *Options) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetTrackOpens

`func (o *Options) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *Options) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *Options) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *Options) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetTrackClicks

`func (o *Options) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *Options) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *Options) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *Options) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


