# EmailJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | ID number of your attachment | [optional] 
**Status** | Pointer to **string** | Name of status: submitted, complete, in_progress | [optional] 
**RecipientsCount** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to [**[]EmailJobFailedStatus**](EmailJobFailedStatus.md) |  | [optional] 
**FailedCount** | Pointer to **int32** | Total emails failed. | [optional] 
**Sent** | Pointer to **[]string** |  | [optional] 
**SentCount** | Pointer to **int32** | Total emails sent. | [optional] 
**Delivered** | Pointer to **[]string** | Number of delivered messages | [optional] 
**DeliveredCount** | Pointer to **int32** |  | [optional] 
**Pending** | Pointer to **[]string** |  | [optional] 
**PendingCount** | Pointer to **int32** |  | [optional] 
**Opened** | Pointer to **[]string** | Number of opened messages | [optional] 
**OpenedCount** | Pointer to **int32** | Total emails opened. | [optional] 
**Clicked** | Pointer to **[]string** | Number of clicked messages | [optional] 
**ClickedCount** | Pointer to **int32** | Total emails clicked | [optional] 
**Unsubscribed** | Pointer to **[]string** | Number of unsubscribed messages | [optional] 
**UnsubscribedCount** | Pointer to **int32** | Total emails unsubscribed | [optional] 
**AbuseReports** | Pointer to **[]string** |  | [optional] 
**AbuseReportsCount** | Pointer to **int32** |  | [optional] 
**MessageIDs** | Pointer to **[]string** | List of all MessageIDs for this job. | [optional] 

## Methods

### NewEmailJobStatus

`func NewEmailJobStatus() *EmailJobStatus`

NewEmailJobStatus instantiates a new EmailJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailJobStatusWithDefaults

`func NewEmailJobStatusWithDefaults() *EmailJobStatus`

NewEmailJobStatusWithDefaults instantiates a new EmailJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *EmailJobStatus) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *EmailJobStatus) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *EmailJobStatus) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *EmailJobStatus) HasID() bool`

HasID returns a boolean if a field has been set.

### GetStatus

`func (o *EmailJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailJobStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRecipientsCount

`func (o *EmailJobStatus) GetRecipientsCount() int32`

GetRecipientsCount returns the RecipientsCount field if non-nil, zero value otherwise.

### GetRecipientsCountOk

`func (o *EmailJobStatus) GetRecipientsCountOk() (*int32, bool)`

GetRecipientsCountOk returns a tuple with the RecipientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsCount

`func (o *EmailJobStatus) SetRecipientsCount(v int32)`

SetRecipientsCount sets RecipientsCount field to given value.

### HasRecipientsCount

`func (o *EmailJobStatus) HasRecipientsCount() bool`

HasRecipientsCount returns a boolean if a field has been set.

### GetFailed

`func (o *EmailJobStatus) GetFailed() []EmailJobFailedStatus`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *EmailJobStatus) GetFailedOk() (*[]EmailJobFailedStatus, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *EmailJobStatus) SetFailed(v []EmailJobFailedStatus)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *EmailJobStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFailedCount

`func (o *EmailJobStatus) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *EmailJobStatus) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *EmailJobStatus) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *EmailJobStatus) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetSent

`func (o *EmailJobStatus) GetSent() []string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *EmailJobStatus) GetSentOk() (*[]string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *EmailJobStatus) SetSent(v []string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *EmailJobStatus) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetSentCount

`func (o *EmailJobStatus) GetSentCount() int32`

GetSentCount returns the SentCount field if non-nil, zero value otherwise.

### GetSentCountOk

`func (o *EmailJobStatus) GetSentCountOk() (*int32, bool)`

GetSentCountOk returns a tuple with the SentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentCount

`func (o *EmailJobStatus) SetSentCount(v int32)`

SetSentCount sets SentCount field to given value.

### HasSentCount

`func (o *EmailJobStatus) HasSentCount() bool`

HasSentCount returns a boolean if a field has been set.

### GetDelivered

`func (o *EmailJobStatus) GetDelivered() []string`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *EmailJobStatus) GetDeliveredOk() (*[]string, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *EmailJobStatus) SetDelivered(v []string)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *EmailJobStatus) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDeliveredCount

`func (o *EmailJobStatus) GetDeliveredCount() int32`

GetDeliveredCount returns the DeliveredCount field if non-nil, zero value otherwise.

### GetDeliveredCountOk

`func (o *EmailJobStatus) GetDeliveredCountOk() (*int32, bool)`

GetDeliveredCountOk returns a tuple with the DeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCount

`func (o *EmailJobStatus) SetDeliveredCount(v int32)`

SetDeliveredCount sets DeliveredCount field to given value.

### HasDeliveredCount

`func (o *EmailJobStatus) HasDeliveredCount() bool`

HasDeliveredCount returns a boolean if a field has been set.

### GetPending

`func (o *EmailJobStatus) GetPending() []string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *EmailJobStatus) GetPendingOk() (*[]string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *EmailJobStatus) SetPending(v []string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *EmailJobStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPendingCount

`func (o *EmailJobStatus) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *EmailJobStatus) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *EmailJobStatus) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *EmailJobStatus) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetOpened

`func (o *EmailJobStatus) GetOpened() []string`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *EmailJobStatus) GetOpenedOk() (*[]string, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *EmailJobStatus) SetOpened(v []string)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *EmailJobStatus) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetOpenedCount

`func (o *EmailJobStatus) GetOpenedCount() int32`

GetOpenedCount returns the OpenedCount field if non-nil, zero value otherwise.

### GetOpenedCountOk

`func (o *EmailJobStatus) GetOpenedCountOk() (*int32, bool)`

GetOpenedCountOk returns a tuple with the OpenedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedCount

`func (o *EmailJobStatus) SetOpenedCount(v int32)`

SetOpenedCount sets OpenedCount field to given value.

### HasOpenedCount

`func (o *EmailJobStatus) HasOpenedCount() bool`

HasOpenedCount returns a boolean if a field has been set.

### GetClicked

`func (o *EmailJobStatus) GetClicked() []string`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *EmailJobStatus) GetClickedOk() (*[]string, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *EmailJobStatus) SetClicked(v []string)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *EmailJobStatus) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetClickedCount

`func (o *EmailJobStatus) GetClickedCount() int32`

GetClickedCount returns the ClickedCount field if non-nil, zero value otherwise.

### GetClickedCountOk

`func (o *EmailJobStatus) GetClickedCountOk() (*int32, bool)`

GetClickedCountOk returns a tuple with the ClickedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedCount

`func (o *EmailJobStatus) SetClickedCount(v int32)`

SetClickedCount sets ClickedCount field to given value.

### HasClickedCount

`func (o *EmailJobStatus) HasClickedCount() bool`

HasClickedCount returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *EmailJobStatus) GetUnsubscribed() []string`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *EmailJobStatus) GetUnsubscribedOk() (*[]string, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *EmailJobStatus) SetUnsubscribed(v []string)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *EmailJobStatus) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetUnsubscribedCount

`func (o *EmailJobStatus) GetUnsubscribedCount() int32`

GetUnsubscribedCount returns the UnsubscribedCount field if non-nil, zero value otherwise.

### GetUnsubscribedCountOk

`func (o *EmailJobStatus) GetUnsubscribedCountOk() (*int32, bool)`

GetUnsubscribedCountOk returns a tuple with the UnsubscribedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribedCount

`func (o *EmailJobStatus) SetUnsubscribedCount(v int32)`

SetUnsubscribedCount sets UnsubscribedCount field to given value.

### HasUnsubscribedCount

`func (o *EmailJobStatus) HasUnsubscribedCount() bool`

HasUnsubscribedCount returns a boolean if a field has been set.

### GetAbuseReports

`func (o *EmailJobStatus) GetAbuseReports() []string`

GetAbuseReports returns the AbuseReports field if non-nil, zero value otherwise.

### GetAbuseReportsOk

`func (o *EmailJobStatus) GetAbuseReportsOk() (*[]string, bool)`

GetAbuseReportsOk returns a tuple with the AbuseReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseReports

`func (o *EmailJobStatus) SetAbuseReports(v []string)`

SetAbuseReports sets AbuseReports field to given value.

### HasAbuseReports

`func (o *EmailJobStatus) HasAbuseReports() bool`

HasAbuseReports returns a boolean if a field has been set.

### GetAbuseReportsCount

`func (o *EmailJobStatus) GetAbuseReportsCount() int32`

GetAbuseReportsCount returns the AbuseReportsCount field if non-nil, zero value otherwise.

### GetAbuseReportsCountOk

`func (o *EmailJobStatus) GetAbuseReportsCountOk() (*int32, bool)`

GetAbuseReportsCountOk returns a tuple with the AbuseReportsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuseReportsCount

`func (o *EmailJobStatus) SetAbuseReportsCount(v int32)`

SetAbuseReportsCount sets AbuseReportsCount field to given value.

### HasAbuseReportsCount

`func (o *EmailJobStatus) HasAbuseReportsCount() bool`

HasAbuseReportsCount returns a boolean if a field has been set.

### GetMessageIDs

`func (o *EmailJobStatus) GetMessageIDs() []string`

GetMessageIDs returns the MessageIDs field if non-nil, zero value otherwise.

### GetMessageIDsOk

`func (o *EmailJobStatus) GetMessageIDsOk() (*[]string, bool)`

GetMessageIDsOk returns a tuple with the MessageIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageIDs

`func (o *EmailJobStatus) SetMessageIDs(v []string)`

SetMessageIDs sets MessageIDs field to given value.

### HasMessageIDs

`func (o *EmailJobStatus) HasMessageIDs() bool`

HasMessageIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


