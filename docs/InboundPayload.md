# InboundPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Filter of the inbound data | [optional] 
**Name** | Pointer to **string** | Name of this route | [optional] 
**FilterType** | Pointer to [**InboundRouteFilterType**](InboundRouteFilterType.md) | Type of the filter | [optional] 
**ActionType** | Pointer to [**InboundRouteActionType**](InboundRouteActionType.md) | Type of action to take | [optional] 
**EmailAddress** | Pointer to **string** | Email to forward the inbound to | [optional] 
**HttpAddress** | Pointer to **string** | Address to notify about the inbound | [optional] 

## Methods

### NewInboundPayload

`func NewInboundPayload() *InboundPayload`

NewInboundPayload instantiates a new InboundPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundPayloadWithDefaults

`func NewInboundPayloadWithDefaults() *InboundPayload`

NewInboundPayloadWithDefaults instantiates a new InboundPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *InboundPayload) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InboundPayload) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InboundPayload) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InboundPayload) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetName

`func (o *InboundPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InboundPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InboundPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InboundPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilterType

`func (o *InboundPayload) GetFilterType() InboundRouteFilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *InboundPayload) GetFilterTypeOk() (*InboundRouteFilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *InboundPayload) SetFilterType(v InboundRouteFilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *InboundPayload) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetActionType

`func (o *InboundPayload) GetActionType() InboundRouteActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *InboundPayload) GetActionTypeOk() (*InboundRouteActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *InboundPayload) SetActionType(v InboundRouteActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *InboundPayload) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InboundPayload) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InboundPayload) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InboundPayload) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InboundPayload) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetHttpAddress

`func (o *InboundPayload) GetHttpAddress() string`

GetHttpAddress returns the HttpAddress field if non-nil, zero value otherwise.

### GetHttpAddressOk

`func (o *InboundPayload) GetHttpAddressOk() (*string, bool)`

GetHttpAddressOk returns a tuple with the HttpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAddress

`func (o *InboundPayload) SetHttpAddress(v string)`

SetHttpAddress sets HttpAddress field to given value.

### HasHttpAddress

`func (o *InboundPayload) HasHttpAddress() bool`

HasHttpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


