# ContactsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListName** | Pointer to **string** | Name of your list. | [optional] 
**PublicListID** | Pointer to **NullableString** | ID code of list. Please note that this is different from the listid field. | [optional] 
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**AllowUnsubscribe** | Pointer to **bool** | True: Allow unsubscribing from this list. Otherwise, false | [optional] 

## Methods

### NewContactsList

`func NewContactsList() *ContactsList`

NewContactsList instantiates a new ContactsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsListWithDefaults

`func NewContactsListWithDefaults() *ContactsList`

NewContactsListWithDefaults instantiates a new ContactsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListName

`func (o *ContactsList) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *ContactsList) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *ContactsList) SetListName(v string)`

SetListName sets ListName field to given value.

### HasListName

`func (o *ContactsList) HasListName() bool`

HasListName returns a boolean if a field has been set.

### GetPublicListID

`func (o *ContactsList) GetPublicListID() string`

GetPublicListID returns the PublicListID field if non-nil, zero value otherwise.

### GetPublicListIDOk

`func (o *ContactsList) GetPublicListIDOk() (*string, bool)`

GetPublicListIDOk returns a tuple with the PublicListID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicListID

`func (o *ContactsList) SetPublicListID(v string)`

SetPublicListID sets PublicListID field to given value.

### HasPublicListID

`func (o *ContactsList) HasPublicListID() bool`

HasPublicListID returns a boolean if a field has been set.

### SetPublicListIDNil

`func (o *ContactsList) SetPublicListIDNil(b bool)`

 SetPublicListIDNil sets the value for PublicListID to be an explicit nil

### UnsetPublicListID
`func (o *ContactsList) UnsetPublicListID()`

UnsetPublicListID ensures that no value is present for PublicListID, not even an explicit nil
### GetDateAdded

`func (o *ContactsList) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ContactsList) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ContactsList) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ContactsList) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetAllowUnsubscribe

`func (o *ContactsList) GetAllowUnsubscribe() bool`

GetAllowUnsubscribe returns the AllowUnsubscribe field if non-nil, zero value otherwise.

### GetAllowUnsubscribeOk

`func (o *ContactsList) GetAllowUnsubscribeOk() (*bool, bool)`

GetAllowUnsubscribeOk returns a tuple with the AllowUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsubscribe

`func (o *ContactsList) SetAllowUnsubscribe(v bool)`

SetAllowUnsubscribe sets AllowUnsubscribe field to given value.

### HasAllowUnsubscribe

`func (o *ContactsList) HasAllowUnsubscribe() bool`

HasAllowUnsubscribe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


