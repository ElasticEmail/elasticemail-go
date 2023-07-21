# InboundRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of this route | [optional] 
**FilterType** | Pointer to [**InboundRouteFilterType**](InboundRouteFilterType.md) |  | [optional] [default to INBOUNDROUTEFILTERTYPE_EMAIL_ADDRESS]
**Filter** | Pointer to **string** | Filter of the inbound data | [optional] 
**ActionType** | Pointer to [**InboundRouteActionType**](InboundRouteActionType.md) |  | [optional] [default to INBOUNDROUTEACTIONTYPE_FORWARD_TO_EMAIL]
**ActionParameter** | Pointer to **string** | URL address or Email to notify about the inbound | [optional] 
**SortOrder** | Pointer to **int32** | Place of this route in your routes queue&#39;s order | [optional] 

## Methods

### NewInboundRoute

`func NewInboundRoute() *InboundRoute`

NewInboundRoute instantiates a new InboundRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundRouteWithDefaults

`func NewInboundRouteWithDefaults() *InboundRoute`

NewInboundRouteWithDefaults instantiates a new InboundRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicId

`func (o *InboundRoute) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *InboundRoute) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *InboundRoute) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *InboundRoute) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetName

`func (o *InboundRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InboundRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InboundRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InboundRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilterType

`func (o *InboundRoute) GetFilterType() InboundRouteFilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *InboundRoute) GetFilterTypeOk() (*InboundRouteFilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *InboundRoute) SetFilterType(v InboundRouteFilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *InboundRoute) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetFilter

`func (o *InboundRoute) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InboundRoute) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InboundRoute) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InboundRoute) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetActionType

`func (o *InboundRoute) GetActionType() InboundRouteActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *InboundRoute) GetActionTypeOk() (*InboundRouteActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *InboundRoute) SetActionType(v InboundRouteActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *InboundRoute) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionParameter

`func (o *InboundRoute) GetActionParameter() string`

GetActionParameter returns the ActionParameter field if non-nil, zero value otherwise.

### GetActionParameterOk

`func (o *InboundRoute) GetActionParameterOk() (*string, bool)`

GetActionParameterOk returns a tuple with the ActionParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionParameter

`func (o *InboundRoute) SetActionParameter(v string)`

SetActionParameter sets ActionParameter field to given value.

### HasActionParameter

`func (o *InboundRoute) HasActionParameter() bool`

HasActionParameter returns a boolean if a field has been set.

### GetSortOrder

`func (o *InboundRoute) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InboundRoute) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InboundRoute) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InboundRoute) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


