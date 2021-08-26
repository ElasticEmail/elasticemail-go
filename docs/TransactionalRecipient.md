# TransactionalRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** | List of recipients (visible to others) | 
**CC** | Pointer to **[]string** | List of Carbon Copy recipients (visible to others) | [optional] 
**BCC** | Pointer to **[]string** | List of Blind Carbon Copy recipients (hidden from other recipients) | [optional] 

## Methods

### NewTransactionalRecipient

`func NewTransactionalRecipient(to []string, ) *TransactionalRecipient`

NewTransactionalRecipient instantiates a new TransactionalRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionalRecipientWithDefaults

`func NewTransactionalRecipientWithDefaults() *TransactionalRecipient`

NewTransactionalRecipientWithDefaults instantiates a new TransactionalRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *TransactionalRecipient) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransactionalRecipient) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransactionalRecipient) SetTo(v []string)`

SetTo sets To field to given value.


### GetCC

`func (o *TransactionalRecipient) GetCC() []string`

GetCC returns the CC field if non-nil, zero value otherwise.

### GetCCOk

`func (o *TransactionalRecipient) GetCCOk() (*[]string, bool)`

GetCCOk returns a tuple with the CC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCC

`func (o *TransactionalRecipient) SetCC(v []string)`

SetCC sets CC field to given value.

### HasCC

`func (o *TransactionalRecipient) HasCC() bool`

HasCC returns a boolean if a field has been set.

### GetBCC

`func (o *TransactionalRecipient) GetBCC() []string`

GetBCC returns the BCC field if non-nil, zero value otherwise.

### GetBCCOk

`func (o *TransactionalRecipient) GetBCCOk() (*[]string, bool)`

GetBCCOk returns a tuple with the BCC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCC

`func (o *TransactionalRecipient) SetBCC(v []string)`

SetBCC sets BCC field to given value.

### HasBCC

`func (o *TransactionalRecipient) HasBCC() bool`

HasBCC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


