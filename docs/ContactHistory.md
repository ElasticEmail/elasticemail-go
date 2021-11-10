# ContactHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to [**ContactHistEventType**](ContactHistEventType.md) |  | [optional] [default to OPENED]
**EventDate** | Pointer to **time.Time** | Formatted date of event. | [optional] 
**ChannelName** | Pointer to **string** | Name of channel this event occured on | [optional] 
**TemplateName** | Pointer to **string** | Name of template this event occured on | [optional] 
**IPAddress** | Pointer to **string** | IP Address of the event. | [optional] 
**Country** | Pointer to **string** | Country of the event. | [optional] 
**Data** | Pointer to **string** | Additional information about the event | [optional] 

## Methods

### NewContactHistory

`func NewContactHistory() *ContactHistory`

NewContactHistory instantiates a new ContactHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactHistoryWithDefaults

`func NewContactHistoryWithDefaults() *ContactHistory`

NewContactHistoryWithDefaults instantiates a new ContactHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ContactHistory) GetEventType() ContactHistEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ContactHistory) GetEventTypeOk() (*ContactHistEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ContactHistory) SetEventType(v ContactHistEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ContactHistory) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventDate

`func (o *ContactHistory) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *ContactHistory) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *ContactHistory) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *ContactHistory) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### GetChannelName

`func (o *ContactHistory) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ContactHistory) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ContactHistory) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ContactHistory) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetTemplateName

`func (o *ContactHistory) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ContactHistory) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ContactHistory) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ContactHistory) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetIPAddress

`func (o *ContactHistory) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *ContactHistory) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *ContactHistory) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *ContactHistory) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetCountry

`func (o *ContactHistory) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ContactHistory) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ContactHistory) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ContactHistory) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetData

`func (o *ContactHistory) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContactHistory) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContactHistory) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ContactHistory) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


