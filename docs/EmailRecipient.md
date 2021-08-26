# EmailRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Proper email address. | 
**Fields** | Pointer to **map[string]string** | A key-value collection of merge fields which can be used in e-mail body. | [optional] 

## Methods

### NewEmailRecipient

`func NewEmailRecipient(email string, ) *EmailRecipient`

NewEmailRecipient instantiates a new EmailRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRecipientWithDefaults

`func NewEmailRecipientWithDefaults() *EmailRecipient`

NewEmailRecipientWithDefaults instantiates a new EmailRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFields

`func (o *EmailRecipient) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EmailRecipient) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EmailRecipient) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EmailRecipient) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


