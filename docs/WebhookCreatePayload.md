# WebhookCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Filename | 
**URL** | **string** | URL of notification. | 
**NotifyOncePerEmail** | Pointer to **bool** |  | [optional] 
**NotificationForSent** | Pointer to **bool** |  | [optional] 
**NotificationForOpened** | Pointer to **bool** |  | [optional] 
**NotificationForClicked** | Pointer to **bool** |  | [optional] 
**NotificationForUnsubscribed** | Pointer to **bool** |  | [optional] 
**NotificationForAbuseReport** | Pointer to **bool** |  | [optional] 
**NotificationForError** | Pointer to **bool** |  | [optional] 

## Methods

### NewWebhookCreatePayload

`func NewWebhookCreatePayload(name string, uRL string, ) *WebhookCreatePayload`

NewWebhookCreatePayload instantiates a new WebhookCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreatePayloadWithDefaults

`func NewWebhookCreatePayloadWithDefaults() *WebhookCreatePayload`

NewWebhookCreatePayloadWithDefaults instantiates a new WebhookCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookCreatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookCreatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookCreatePayload) SetName(v string)`

SetName sets Name field to given value.


### GetURL

`func (o *WebhookCreatePayload) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *WebhookCreatePayload) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *WebhookCreatePayload) SetURL(v string)`

SetURL sets URL field to given value.


### GetNotifyOncePerEmail

`func (o *WebhookCreatePayload) GetNotifyOncePerEmail() bool`

GetNotifyOncePerEmail returns the NotifyOncePerEmail field if non-nil, zero value otherwise.

### GetNotifyOncePerEmailOk

`func (o *WebhookCreatePayload) GetNotifyOncePerEmailOk() (*bool, bool)`

GetNotifyOncePerEmailOk returns a tuple with the NotifyOncePerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOncePerEmail

`func (o *WebhookCreatePayload) SetNotifyOncePerEmail(v bool)`

SetNotifyOncePerEmail sets NotifyOncePerEmail field to given value.

### HasNotifyOncePerEmail

`func (o *WebhookCreatePayload) HasNotifyOncePerEmail() bool`

HasNotifyOncePerEmail returns a boolean if a field has been set.

### GetNotificationForSent

`func (o *WebhookCreatePayload) GetNotificationForSent() bool`

GetNotificationForSent returns the NotificationForSent field if non-nil, zero value otherwise.

### GetNotificationForSentOk

`func (o *WebhookCreatePayload) GetNotificationForSentOk() (*bool, bool)`

GetNotificationForSentOk returns a tuple with the NotificationForSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForSent

`func (o *WebhookCreatePayload) SetNotificationForSent(v bool)`

SetNotificationForSent sets NotificationForSent field to given value.

### HasNotificationForSent

`func (o *WebhookCreatePayload) HasNotificationForSent() bool`

HasNotificationForSent returns a boolean if a field has been set.

### GetNotificationForOpened

`func (o *WebhookCreatePayload) GetNotificationForOpened() bool`

GetNotificationForOpened returns the NotificationForOpened field if non-nil, zero value otherwise.

### GetNotificationForOpenedOk

`func (o *WebhookCreatePayload) GetNotificationForOpenedOk() (*bool, bool)`

GetNotificationForOpenedOk returns a tuple with the NotificationForOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForOpened

`func (o *WebhookCreatePayload) SetNotificationForOpened(v bool)`

SetNotificationForOpened sets NotificationForOpened field to given value.

### HasNotificationForOpened

`func (o *WebhookCreatePayload) HasNotificationForOpened() bool`

HasNotificationForOpened returns a boolean if a field has been set.

### GetNotificationForClicked

`func (o *WebhookCreatePayload) GetNotificationForClicked() bool`

GetNotificationForClicked returns the NotificationForClicked field if non-nil, zero value otherwise.

### GetNotificationForClickedOk

`func (o *WebhookCreatePayload) GetNotificationForClickedOk() (*bool, bool)`

GetNotificationForClickedOk returns a tuple with the NotificationForClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForClicked

`func (o *WebhookCreatePayload) SetNotificationForClicked(v bool)`

SetNotificationForClicked sets NotificationForClicked field to given value.

### HasNotificationForClicked

`func (o *WebhookCreatePayload) HasNotificationForClicked() bool`

HasNotificationForClicked returns a boolean if a field has been set.

### GetNotificationForUnsubscribed

`func (o *WebhookCreatePayload) GetNotificationForUnsubscribed() bool`

GetNotificationForUnsubscribed returns the NotificationForUnsubscribed field if non-nil, zero value otherwise.

### GetNotificationForUnsubscribedOk

`func (o *WebhookCreatePayload) GetNotificationForUnsubscribedOk() (*bool, bool)`

GetNotificationForUnsubscribedOk returns a tuple with the NotificationForUnsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForUnsubscribed

`func (o *WebhookCreatePayload) SetNotificationForUnsubscribed(v bool)`

SetNotificationForUnsubscribed sets NotificationForUnsubscribed field to given value.

### HasNotificationForUnsubscribed

`func (o *WebhookCreatePayload) HasNotificationForUnsubscribed() bool`

HasNotificationForUnsubscribed returns a boolean if a field has been set.

### GetNotificationForAbuseReport

`func (o *WebhookCreatePayload) GetNotificationForAbuseReport() bool`

GetNotificationForAbuseReport returns the NotificationForAbuseReport field if non-nil, zero value otherwise.

### GetNotificationForAbuseReportOk

`func (o *WebhookCreatePayload) GetNotificationForAbuseReportOk() (*bool, bool)`

GetNotificationForAbuseReportOk returns a tuple with the NotificationForAbuseReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForAbuseReport

`func (o *WebhookCreatePayload) SetNotificationForAbuseReport(v bool)`

SetNotificationForAbuseReport sets NotificationForAbuseReport field to given value.

### HasNotificationForAbuseReport

`func (o *WebhookCreatePayload) HasNotificationForAbuseReport() bool`

HasNotificationForAbuseReport returns a boolean if a field has been set.

### GetNotificationForError

`func (o *WebhookCreatePayload) GetNotificationForError() bool`

GetNotificationForError returns the NotificationForError field if non-nil, zero value otherwise.

### GetNotificationForErrorOk

`func (o *WebhookCreatePayload) GetNotificationForErrorOk() (*bool, bool)`

GetNotificationForErrorOk returns a tuple with the NotificationForError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForError

`func (o *WebhookCreatePayload) SetNotificationForError(v bool)`

SetNotificationForError sets NotificationForError field to given value.

### HasNotificationForError

`func (o *WebhookCreatePayload) HasNotificationForError() bool`

HasNotificationForError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


