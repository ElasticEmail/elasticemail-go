# EmailStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Email address this email was sent from. | [optional] 
**To** | Pointer to **string** | Email address this email was sent to. | [optional] 
**Date** | Pointer to **time.Time** | Date the email was submitted. | [optional] 
**Status** | Pointer to [**LogJobStatus**](LogJobStatus.md) |  | [optional] [default to LOGJOBSTATUS_ALL]
**StatusName** | Pointer to **string** | Name of email&#39;s status | [optional] 
**StatusChangeDate** | Pointer to **time.Time** | Date of last status change. | [optional] 
**DateSent** | Pointer to **time.Time** | Date when the email was sent | [optional] 
**DateOpened** | Pointer to **NullableTime** | Date when the email changed the status to &#39;opened&#39; | [optional] 
**DateClicked** | Pointer to **NullableTime** | Date when the email changed the status to &#39;clicked&#39; | [optional] 
**ErrorMessage** | Pointer to **string** | Detailed error or bounced message. | [optional] 
**TransactionID** | Pointer to **string** | ID number of transaction | [optional] 
**EnvelopeFrom** | Pointer to **string** | Envelope from address | [optional] 

## Methods

### NewEmailStatus

`func NewEmailStatus() *EmailStatus`

NewEmailStatus instantiates a new EmailStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatusWithDefaults

`func NewEmailStatusWithDefaults() *EmailStatus`

NewEmailStatusWithDefaults instantiates a new EmailStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *EmailStatus) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailStatus) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailStatus) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailStatus) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *EmailStatus) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailStatus) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailStatus) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailStatus) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDate

`func (o *EmailStatus) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EmailStatus) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EmailStatus) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *EmailStatus) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStatus

`func (o *EmailStatus) GetStatus() LogJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailStatus) GetStatusOk() (*LogJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailStatus) SetStatus(v LogJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusName

`func (o *EmailStatus) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *EmailStatus) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *EmailStatus) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *EmailStatus) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetStatusChangeDate

`func (o *EmailStatus) GetStatusChangeDate() time.Time`

GetStatusChangeDate returns the StatusChangeDate field if non-nil, zero value otherwise.

### GetStatusChangeDateOk

`func (o *EmailStatus) GetStatusChangeDateOk() (*time.Time, bool)`

GetStatusChangeDateOk returns a tuple with the StatusChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeDate

`func (o *EmailStatus) SetStatusChangeDate(v time.Time)`

SetStatusChangeDate sets StatusChangeDate field to given value.

### HasStatusChangeDate

`func (o *EmailStatus) HasStatusChangeDate() bool`

HasStatusChangeDate returns a boolean if a field has been set.

### GetDateSent

`func (o *EmailStatus) GetDateSent() time.Time`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *EmailStatus) GetDateSentOk() (*time.Time, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *EmailStatus) SetDateSent(v time.Time)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *EmailStatus) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### GetDateOpened

`func (o *EmailStatus) GetDateOpened() time.Time`

GetDateOpened returns the DateOpened field if non-nil, zero value otherwise.

### GetDateOpenedOk

`func (o *EmailStatus) GetDateOpenedOk() (*time.Time, bool)`

GetDateOpenedOk returns a tuple with the DateOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOpened

`func (o *EmailStatus) SetDateOpened(v time.Time)`

SetDateOpened sets DateOpened field to given value.

### HasDateOpened

`func (o *EmailStatus) HasDateOpened() bool`

HasDateOpened returns a boolean if a field has been set.

### SetDateOpenedNil

`func (o *EmailStatus) SetDateOpenedNil(b bool)`

 SetDateOpenedNil sets the value for DateOpened to be an explicit nil

### UnsetDateOpened
`func (o *EmailStatus) UnsetDateOpened()`

UnsetDateOpened ensures that no value is present for DateOpened, not even an explicit nil
### GetDateClicked

`func (o *EmailStatus) GetDateClicked() time.Time`

GetDateClicked returns the DateClicked field if non-nil, zero value otherwise.

### GetDateClickedOk

`func (o *EmailStatus) GetDateClickedOk() (*time.Time, bool)`

GetDateClickedOk returns a tuple with the DateClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClicked

`func (o *EmailStatus) SetDateClicked(v time.Time)`

SetDateClicked sets DateClicked field to given value.

### HasDateClicked

`func (o *EmailStatus) HasDateClicked() bool`

HasDateClicked returns a boolean if a field has been set.

### SetDateClickedNil

`func (o *EmailStatus) SetDateClickedNil(b bool)`

 SetDateClickedNil sets the value for DateClicked to be an explicit nil

### UnsetDateClicked
`func (o *EmailStatus) UnsetDateClicked()`

UnsetDateClicked ensures that no value is present for DateClicked, not even an explicit nil
### GetErrorMessage

`func (o *EmailStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EmailStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EmailStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EmailStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetTransactionID

`func (o *EmailStatus) GetTransactionID() string`

GetTransactionID returns the TransactionID field if non-nil, zero value otherwise.

### GetTransactionIDOk

`func (o *EmailStatus) GetTransactionIDOk() (*string, bool)`

GetTransactionIDOk returns a tuple with the TransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionID

`func (o *EmailStatus) SetTransactionID(v string)`

SetTransactionID sets TransactionID field to given value.

### HasTransactionID

`func (o *EmailStatus) HasTransactionID() bool`

HasTransactionID returns a boolean if a field has been set.

### GetEnvelopeFrom

`func (o *EmailStatus) GetEnvelopeFrom() string`

GetEnvelopeFrom returns the EnvelopeFrom field if non-nil, zero value otherwise.

### GetEnvelopeFromOk

`func (o *EmailStatus) GetEnvelopeFromOk() (*string, bool)`

GetEnvelopeFromOk returns a tuple with the EnvelopeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelopeFrom

`func (o *EmailStatus) SetEnvelopeFrom(v string)`

SetEnvelopeFrom sets EnvelopeFrom field to given value.

### HasEnvelopeFrom

`func (o *EmailStatus) HasEnvelopeFrom() bool`

HasEnvelopeFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


