# RecipientEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionID** | Pointer to **string** | ID number of transaction | [optional] 
**MsgID** | Pointer to **string** | ID number of selected message. | [optional] 
**FromEmail** | Pointer to **string** | Default From: email address. | [optional] 
**To** | Pointer to **string** | Ending date for search in YYYY-MM-DDThh:mm:ss format. | [optional] 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**EventType** | Pointer to [**EventType**](EventType.md) |  | [optional] [default to EVENTTYPE_SUBMISSION]
**EventDate** | Pointer to **time.Time** | Creation date | [optional] 
**ChannelName** | Pointer to **string** | Name of selected channel. | [optional] 
**MessageCategory** | Pointer to [**MessageCategory**](MessageCategory.md) |  | [optional] [default to MESSAGECATEGORY_UNKNOWN]
**NextTryOn** | Pointer to **NullableTime** | Date of next try | [optional] 
**Message** | Pointer to **string** | Content of message, HTML encoded | [optional] 
**IPAddress** | Pointer to **string** | IP which this email was sent through | [optional] 
**PoolName** | Pointer to **string** | Name of an IP pool this email was sent through | [optional] 

## Methods

### NewRecipientEvent

`func NewRecipientEvent() *RecipientEvent`

NewRecipientEvent instantiates a new RecipientEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientEventWithDefaults

`func NewRecipientEventWithDefaults() *RecipientEvent`

NewRecipientEventWithDefaults instantiates a new RecipientEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionID

`func (o *RecipientEvent) GetTransactionID() string`

GetTransactionID returns the TransactionID field if non-nil, zero value otherwise.

### GetTransactionIDOk

`func (o *RecipientEvent) GetTransactionIDOk() (*string, bool)`

GetTransactionIDOk returns a tuple with the TransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionID

`func (o *RecipientEvent) SetTransactionID(v string)`

SetTransactionID sets TransactionID field to given value.

### HasTransactionID

`func (o *RecipientEvent) HasTransactionID() bool`

HasTransactionID returns a boolean if a field has been set.

### GetMsgID

`func (o *RecipientEvent) GetMsgID() string`

GetMsgID returns the MsgID field if non-nil, zero value otherwise.

### GetMsgIDOk

`func (o *RecipientEvent) GetMsgIDOk() (*string, bool)`

GetMsgIDOk returns a tuple with the MsgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgID

`func (o *RecipientEvent) SetMsgID(v string)`

SetMsgID sets MsgID field to given value.

### HasMsgID

`func (o *RecipientEvent) HasMsgID() bool`

HasMsgID returns a boolean if a field has been set.

### GetFromEmail

`func (o *RecipientEvent) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *RecipientEvent) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *RecipientEvent) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *RecipientEvent) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetTo

`func (o *RecipientEvent) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RecipientEvent) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RecipientEvent) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RecipientEvent) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSubject

`func (o *RecipientEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RecipientEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RecipientEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RecipientEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetEventType

`func (o *RecipientEvent) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *RecipientEvent) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *RecipientEvent) SetEventType(v EventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *RecipientEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventDate

`func (o *RecipientEvent) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *RecipientEvent) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *RecipientEvent) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *RecipientEvent) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### GetChannelName

`func (o *RecipientEvent) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *RecipientEvent) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *RecipientEvent) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *RecipientEvent) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetMessageCategory

`func (o *RecipientEvent) GetMessageCategory() MessageCategory`

GetMessageCategory returns the MessageCategory field if non-nil, zero value otherwise.

### GetMessageCategoryOk

`func (o *RecipientEvent) GetMessageCategoryOk() (*MessageCategory, bool)`

GetMessageCategoryOk returns a tuple with the MessageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCategory

`func (o *RecipientEvent) SetMessageCategory(v MessageCategory)`

SetMessageCategory sets MessageCategory field to given value.

### HasMessageCategory

`func (o *RecipientEvent) HasMessageCategory() bool`

HasMessageCategory returns a boolean if a field has been set.

### GetNextTryOn

`func (o *RecipientEvent) GetNextTryOn() time.Time`

GetNextTryOn returns the NextTryOn field if non-nil, zero value otherwise.

### GetNextTryOnOk

`func (o *RecipientEvent) GetNextTryOnOk() (*time.Time, bool)`

GetNextTryOnOk returns a tuple with the NextTryOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTryOn

`func (o *RecipientEvent) SetNextTryOn(v time.Time)`

SetNextTryOn sets NextTryOn field to given value.

### HasNextTryOn

`func (o *RecipientEvent) HasNextTryOn() bool`

HasNextTryOn returns a boolean if a field has been set.

### SetNextTryOnNil

`func (o *RecipientEvent) SetNextTryOnNil(b bool)`

 SetNextTryOnNil sets the value for NextTryOn to be an explicit nil

### UnsetNextTryOn
`func (o *RecipientEvent) UnsetNextTryOn()`

UnsetNextTryOn ensures that no value is present for NextTryOn, not even an explicit nil
### GetMessage

`func (o *RecipientEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecipientEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecipientEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RecipientEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIPAddress

`func (o *RecipientEvent) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *RecipientEvent) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *RecipientEvent) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *RecipientEvent) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetPoolName

`func (o *RecipientEvent) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *RecipientEvent) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *RecipientEvent) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *RecipientEvent) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


