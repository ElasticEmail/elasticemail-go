# EmailSend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionID** | Pointer to **string** | ID number of transaction | [optional] 
**MessageID** | Pointer to **string** | Unique identifier for this email. | [optional] 

## Methods

### NewEmailSend

`func NewEmailSend() *EmailSend`

NewEmailSend instantiates a new EmailSend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSendWithDefaults

`func NewEmailSendWithDefaults() *EmailSend`

NewEmailSendWithDefaults instantiates a new EmailSend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionID

`func (o *EmailSend) GetTransactionID() string`

GetTransactionID returns the TransactionID field if non-nil, zero value otherwise.

### GetTransactionIDOk

`func (o *EmailSend) GetTransactionIDOk() (*string, bool)`

GetTransactionIDOk returns a tuple with the TransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionID

`func (o *EmailSend) SetTransactionID(v string)`

SetTransactionID sets TransactionID field to given value.

### HasTransactionID

`func (o *EmailSend) HasTransactionID() bool`

HasTransactionID returns a boolean if a field has been set.

### GetMessageID

`func (o *EmailSend) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *EmailSend) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *EmailSend) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *EmailSend) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


