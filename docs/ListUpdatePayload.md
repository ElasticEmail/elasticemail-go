# ListUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewListName** | Pointer to **string** | Name of your list if you want to change it. | [optional] 
**AllowUnsubscribe** | Pointer to **bool** | True: Allow unsubscribing from this list. Otherwise, false | [optional] 

## Methods

### NewListUpdatePayload

`func NewListUpdatePayload() *ListUpdatePayload`

NewListUpdatePayload instantiates a new ListUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUpdatePayloadWithDefaults

`func NewListUpdatePayloadWithDefaults() *ListUpdatePayload`

NewListUpdatePayloadWithDefaults instantiates a new ListUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewListName

`func (o *ListUpdatePayload) GetNewListName() string`

GetNewListName returns the NewListName field if non-nil, zero value otherwise.

### GetNewListNameOk

`func (o *ListUpdatePayload) GetNewListNameOk() (*string, bool)`

GetNewListNameOk returns a tuple with the NewListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewListName

`func (o *ListUpdatePayload) SetNewListName(v string)`

SetNewListName sets NewListName field to given value.

### HasNewListName

`func (o *ListUpdatePayload) HasNewListName() bool`

HasNewListName returns a boolean if a field has been set.

### GetAllowUnsubscribe

`func (o *ListUpdatePayload) GetAllowUnsubscribe() bool`

GetAllowUnsubscribe returns the AllowUnsubscribe field if non-nil, zero value otherwise.

### GetAllowUnsubscribeOk

`func (o *ListUpdatePayload) GetAllowUnsubscribeOk() (*bool, bool)`

GetAllowUnsubscribeOk returns a tuple with the AllowUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsubscribe

`func (o *ListUpdatePayload) SetAllowUnsubscribe(v bool)`

SetAllowUnsubscribe sets AllowUnsubscribe field to given value.

### HasAllowUnsubscribe

`func (o *ListUpdatePayload) HasAllowUnsubscribe() bool`

HasAllowUnsubscribe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


