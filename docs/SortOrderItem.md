# SortOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicInboundId** | **string** | ID of the route to change the order of | 
**SortOrder** | **int32** | 1 - route will be used first | 

## Methods

### NewSortOrderItem

`func NewSortOrderItem(publicInboundId string, sortOrder int32, ) *SortOrderItem`

NewSortOrderItem instantiates a new SortOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortOrderItemWithDefaults

`func NewSortOrderItemWithDefaults() *SortOrderItem`

NewSortOrderItemWithDefaults instantiates a new SortOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicInboundId

`func (o *SortOrderItem) GetPublicInboundId() string`

GetPublicInboundId returns the PublicInboundId field if non-nil, zero value otherwise.

### GetPublicInboundIdOk

`func (o *SortOrderItem) GetPublicInboundIdOk() (*string, bool)`

GetPublicInboundIdOk returns a tuple with the PublicInboundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInboundId

`func (o *SortOrderItem) SetPublicInboundId(v string)`

SetPublicInboundId sets PublicInboundId field to given value.


### GetSortOrder

`func (o *SortOrderItem) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SortOrderItem) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SortOrderItem) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


