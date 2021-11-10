# EmailTransactionalMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | [**TransactionalRecipient**](TransactionalRecipient.md) |  | 
**Content** | Pointer to [**EmailContent**](EmailContent.md) |  | [optional] 
**Options** | Pointer to [**Options**](Options.md) |  | [optional] 

## Methods

### NewEmailTransactionalMessageData

`func NewEmailTransactionalMessageData(recipients TransactionalRecipient, ) *EmailTransactionalMessageData`

NewEmailTransactionalMessageData instantiates a new EmailTransactionalMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTransactionalMessageDataWithDefaults

`func NewEmailTransactionalMessageDataWithDefaults() *EmailTransactionalMessageData`

NewEmailTransactionalMessageDataWithDefaults instantiates a new EmailTransactionalMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *EmailTransactionalMessageData) GetRecipients() TransactionalRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailTransactionalMessageData) GetRecipientsOk() (*TransactionalRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailTransactionalMessageData) SetRecipients(v TransactionalRecipient)`

SetRecipients sets Recipients field to given value.


### GetContent

`func (o *EmailTransactionalMessageData) GetContent() EmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailTransactionalMessageData) GetContentOk() (*EmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailTransactionalMessageData) SetContent(v EmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *EmailTransactionalMessageData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetOptions

`func (o *EmailTransactionalMessageData) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EmailTransactionalMessageData) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EmailTransactionalMessageData) SetOptions(v Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EmailTransactionalMessageData) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


