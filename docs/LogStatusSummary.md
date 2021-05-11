# LogStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to **int64** | Number of recipients | [optional] 
**EmailTotal** | Pointer to **int64** | Number of emails | [optional] 
**SmsTotal** | Pointer to **int64** | Number of SMS | [optional] 
**Delivered** | Pointer to **int64** | Number of delivered messages | [optional] 
**Bounced** | Pointer to **int64** | Number of bounced messages | [optional] 
**InProgress** | Pointer to **int64** | Number of messages in progress | [optional] 
**Opened** | Pointer to **int64** | Number of opened messages | [optional] 
**Clicked** | Pointer to **int64** | Number of clicked messages | [optional] 
**Unsubscribed** | Pointer to **int64** | Number of unsubscribed messages | [optional] 
**Complaints** | Pointer to **int64** | Number of complaint messages | [optional] 
**Inbound** | Pointer to **int64** | Number of inbound messages | [optional] 
**ManualCancel** | Pointer to **int64** | Number of manually cancelled messages | [optional] 
**NotDelivered** | Pointer to **int64** | Number of messages flagged with &#39;Not Delivered&#39; | [optional] 

## Methods

### NewLogStatusSummary

`func NewLogStatusSummary() *LogStatusSummary`

NewLogStatusSummary instantiates a new LogStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStatusSummaryWithDefaults

`func NewLogStatusSummaryWithDefaults() *LogStatusSummary`

NewLogStatusSummaryWithDefaults instantiates a new LogStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *LogStatusSummary) GetRecipients() int64`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *LogStatusSummary) GetRecipientsOk() (*int64, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *LogStatusSummary) SetRecipients(v int64)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *LogStatusSummary) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetEmailTotal

`func (o *LogStatusSummary) GetEmailTotal() int64`

GetEmailTotal returns the EmailTotal field if non-nil, zero value otherwise.

### GetEmailTotalOk

`func (o *LogStatusSummary) GetEmailTotalOk() (*int64, bool)`

GetEmailTotalOk returns a tuple with the EmailTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTotal

`func (o *LogStatusSummary) SetEmailTotal(v int64)`

SetEmailTotal sets EmailTotal field to given value.

### HasEmailTotal

`func (o *LogStatusSummary) HasEmailTotal() bool`

HasEmailTotal returns a boolean if a field has been set.

### GetSmsTotal

`func (o *LogStatusSummary) GetSmsTotal() int64`

GetSmsTotal returns the SmsTotal field if non-nil, zero value otherwise.

### GetSmsTotalOk

`func (o *LogStatusSummary) GetSmsTotalOk() (*int64, bool)`

GetSmsTotalOk returns a tuple with the SmsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsTotal

`func (o *LogStatusSummary) SetSmsTotal(v int64)`

SetSmsTotal sets SmsTotal field to given value.

### HasSmsTotal

`func (o *LogStatusSummary) HasSmsTotal() bool`

HasSmsTotal returns a boolean if a field has been set.

### GetDelivered

`func (o *LogStatusSummary) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *LogStatusSummary) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *LogStatusSummary) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *LogStatusSummary) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetBounced

`func (o *LogStatusSummary) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *LogStatusSummary) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *LogStatusSummary) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *LogStatusSummary) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetInProgress

`func (o *LogStatusSummary) GetInProgress() int64`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *LogStatusSummary) GetInProgressOk() (*int64, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *LogStatusSummary) SetInProgress(v int64)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *LogStatusSummary) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetOpened

`func (o *LogStatusSummary) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *LogStatusSummary) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *LogStatusSummary) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *LogStatusSummary) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *LogStatusSummary) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *LogStatusSummary) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *LogStatusSummary) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *LogStatusSummary) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *LogStatusSummary) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *LogStatusSummary) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *LogStatusSummary) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *LogStatusSummary) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetComplaints

`func (o *LogStatusSummary) GetComplaints() int64`

GetComplaints returns the Complaints field if non-nil, zero value otherwise.

### GetComplaintsOk

`func (o *LogStatusSummary) GetComplaintsOk() (*int64, bool)`

GetComplaintsOk returns a tuple with the Complaints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaints

`func (o *LogStatusSummary) SetComplaints(v int64)`

SetComplaints sets Complaints field to given value.

### HasComplaints

`func (o *LogStatusSummary) HasComplaints() bool`

HasComplaints returns a boolean if a field has been set.

### GetInbound

`func (o *LogStatusSummary) GetInbound() int64`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *LogStatusSummary) GetInboundOk() (*int64, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *LogStatusSummary) SetInbound(v int64)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *LogStatusSummary) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetManualCancel

`func (o *LogStatusSummary) GetManualCancel() int64`

GetManualCancel returns the ManualCancel field if non-nil, zero value otherwise.

### GetManualCancelOk

`func (o *LogStatusSummary) GetManualCancelOk() (*int64, bool)`

GetManualCancelOk returns a tuple with the ManualCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualCancel

`func (o *LogStatusSummary) SetManualCancel(v int64)`

SetManualCancel sets ManualCancel field to given value.

### HasManualCancel

`func (o *LogStatusSummary) HasManualCancel() bool`

HasManualCancel returns a boolean if a field has been set.

### GetNotDelivered

`func (o *LogStatusSummary) GetNotDelivered() int64`

GetNotDelivered returns the NotDelivered field if non-nil, zero value otherwise.

### GetNotDeliveredOk

`func (o *LogStatusSummary) GetNotDeliveredOk() (*int64, bool)`

GetNotDeliveredOk returns a tuple with the NotDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotDelivered

`func (o *LogStatusSummary) SetNotDelivered(v int64)`

SetNotDelivered sets NotDelivered field to given value.

### HasNotDelivered

`func (o *LogStatusSummary) HasNotDelivered() bool`

HasNotDelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


