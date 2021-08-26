# ListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListName** | **string** | Name of your list. | 
**AllowUnsubscribe** | Pointer to **bool** | True: Allow unsubscribing from this list. Otherwise, false | [optional] 
**Emails** | Pointer to **[]string** | Comma delimited list of existing contact emails that should be added to this list. Leave empty for all contacts | [optional] 

## Methods

### NewListPayload

`func NewListPayload(listName string, ) *ListPayload`

NewListPayload instantiates a new ListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPayloadWithDefaults

`func NewListPayloadWithDefaults() *ListPayload`

NewListPayloadWithDefaults instantiates a new ListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListName

`func (o *ListPayload) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *ListPayload) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *ListPayload) SetListName(v string)`

SetListName sets ListName field to given value.


### GetAllowUnsubscribe

`func (o *ListPayload) GetAllowUnsubscribe() bool`

GetAllowUnsubscribe returns the AllowUnsubscribe field if non-nil, zero value otherwise.

### GetAllowUnsubscribeOk

`func (o *ListPayload) GetAllowUnsubscribeOk() (*bool, bool)`

GetAllowUnsubscribeOk returns a tuple with the AllowUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsubscribe

`func (o *ListPayload) SetAllowUnsubscribe(v bool)`

SetAllowUnsubscribe sets AllowUnsubscribe field to given value.

### HasAllowUnsubscribe

`func (o *ListPayload) HasAllowUnsubscribe() bool`

HasAllowUnsubscribe returns a boolean if a field has been set.

### GetEmails

`func (o *ListPayload) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ListPayload) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ListPayload) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ListPayload) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


