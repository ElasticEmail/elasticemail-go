# ContactActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSent** | Pointer to **int32** | Total emails sent. | [optional] 
**TotalOpened** | Pointer to **int32** | Total emails opened. | [optional] 
**TotalClicked** | Pointer to **int32** | Total emails clicked | [optional] 
**TotalFailed** | Pointer to **int32** | Total emails failed. | [optional] 
**LastSent** | Pointer to **NullableTime** | Last date when an email was sent to this contact | [optional] 
**LastOpened** | Pointer to **NullableTime** | Date this contact last opened an email | [optional] 
**LastClicked** | Pointer to **NullableTime** | Date this contact last clicked an email | [optional] 
**LastFailed** | Pointer to **NullableTime** | Last date when an email sent to this contact bounced | [optional] 
**LastIP** | Pointer to **string** | IP from which this contact opened or clicked their email last time | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | Last RFC Error code if any occurred | [optional] 
**FriendlyErrorMessage** | Pointer to **string** | Last RFC error message if any occurred | [optional] 

## Methods

### NewContactActivity

`func NewContactActivity() *ContactActivity`

NewContactActivity instantiates a new ContactActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactActivityWithDefaults

`func NewContactActivityWithDefaults() *ContactActivity`

NewContactActivityWithDefaults instantiates a new ContactActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSent

`func (o *ContactActivity) GetTotalSent() int32`

GetTotalSent returns the TotalSent field if non-nil, zero value otherwise.

### GetTotalSentOk

`func (o *ContactActivity) GetTotalSentOk() (*int32, bool)`

GetTotalSentOk returns a tuple with the TotalSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSent

`func (o *ContactActivity) SetTotalSent(v int32)`

SetTotalSent sets TotalSent field to given value.

### HasTotalSent

`func (o *ContactActivity) HasTotalSent() bool`

HasTotalSent returns a boolean if a field has been set.

### GetTotalOpened

`func (o *ContactActivity) GetTotalOpened() int32`

GetTotalOpened returns the TotalOpened field if non-nil, zero value otherwise.

### GetTotalOpenedOk

`func (o *ContactActivity) GetTotalOpenedOk() (*int32, bool)`

GetTotalOpenedOk returns a tuple with the TotalOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpened

`func (o *ContactActivity) SetTotalOpened(v int32)`

SetTotalOpened sets TotalOpened field to given value.

### HasTotalOpened

`func (o *ContactActivity) HasTotalOpened() bool`

HasTotalOpened returns a boolean if a field has been set.

### GetTotalClicked

`func (o *ContactActivity) GetTotalClicked() int32`

GetTotalClicked returns the TotalClicked field if non-nil, zero value otherwise.

### GetTotalClickedOk

`func (o *ContactActivity) GetTotalClickedOk() (*int32, bool)`

GetTotalClickedOk returns a tuple with the TotalClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClicked

`func (o *ContactActivity) SetTotalClicked(v int32)`

SetTotalClicked sets TotalClicked field to given value.

### HasTotalClicked

`func (o *ContactActivity) HasTotalClicked() bool`

HasTotalClicked returns a boolean if a field has been set.

### GetTotalFailed

`func (o *ContactActivity) GetTotalFailed() int32`

GetTotalFailed returns the TotalFailed field if non-nil, zero value otherwise.

### GetTotalFailedOk

`func (o *ContactActivity) GetTotalFailedOk() (*int32, bool)`

GetTotalFailedOk returns a tuple with the TotalFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailed

`func (o *ContactActivity) SetTotalFailed(v int32)`

SetTotalFailed sets TotalFailed field to given value.

### HasTotalFailed

`func (o *ContactActivity) HasTotalFailed() bool`

HasTotalFailed returns a boolean if a field has been set.

### GetLastSent

`func (o *ContactActivity) GetLastSent() time.Time`

GetLastSent returns the LastSent field if non-nil, zero value otherwise.

### GetLastSentOk

`func (o *ContactActivity) GetLastSentOk() (*time.Time, bool)`

GetLastSentOk returns a tuple with the LastSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSent

`func (o *ContactActivity) SetLastSent(v time.Time)`

SetLastSent sets LastSent field to given value.

### HasLastSent

`func (o *ContactActivity) HasLastSent() bool`

HasLastSent returns a boolean if a field has been set.

### SetLastSentNil

`func (o *ContactActivity) SetLastSentNil(b bool)`

 SetLastSentNil sets the value for LastSent to be an explicit nil

### UnsetLastSent
`func (o *ContactActivity) UnsetLastSent()`

UnsetLastSent ensures that no value is present for LastSent, not even an explicit nil
### GetLastOpened

`func (o *ContactActivity) GetLastOpened() time.Time`

GetLastOpened returns the LastOpened field if non-nil, zero value otherwise.

### GetLastOpenedOk

`func (o *ContactActivity) GetLastOpenedOk() (*time.Time, bool)`

GetLastOpenedOk returns a tuple with the LastOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpened

`func (o *ContactActivity) SetLastOpened(v time.Time)`

SetLastOpened sets LastOpened field to given value.

### HasLastOpened

`func (o *ContactActivity) HasLastOpened() bool`

HasLastOpened returns a boolean if a field has been set.

### SetLastOpenedNil

`func (o *ContactActivity) SetLastOpenedNil(b bool)`

 SetLastOpenedNil sets the value for LastOpened to be an explicit nil

### UnsetLastOpened
`func (o *ContactActivity) UnsetLastOpened()`

UnsetLastOpened ensures that no value is present for LastOpened, not even an explicit nil
### GetLastClicked

`func (o *ContactActivity) GetLastClicked() time.Time`

GetLastClicked returns the LastClicked field if non-nil, zero value otherwise.

### GetLastClickedOk

`func (o *ContactActivity) GetLastClickedOk() (*time.Time, bool)`

GetLastClickedOk returns a tuple with the LastClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClicked

`func (o *ContactActivity) SetLastClicked(v time.Time)`

SetLastClicked sets LastClicked field to given value.

### HasLastClicked

`func (o *ContactActivity) HasLastClicked() bool`

HasLastClicked returns a boolean if a field has been set.

### SetLastClickedNil

`func (o *ContactActivity) SetLastClickedNil(b bool)`

 SetLastClickedNil sets the value for LastClicked to be an explicit nil

### UnsetLastClicked
`func (o *ContactActivity) UnsetLastClicked()`

UnsetLastClicked ensures that no value is present for LastClicked, not even an explicit nil
### GetLastFailed

`func (o *ContactActivity) GetLastFailed() time.Time`

GetLastFailed returns the LastFailed field if non-nil, zero value otherwise.

### GetLastFailedOk

`func (o *ContactActivity) GetLastFailedOk() (*time.Time, bool)`

GetLastFailedOk returns a tuple with the LastFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailed

`func (o *ContactActivity) SetLastFailed(v time.Time)`

SetLastFailed sets LastFailed field to given value.

### HasLastFailed

`func (o *ContactActivity) HasLastFailed() bool`

HasLastFailed returns a boolean if a field has been set.

### SetLastFailedNil

`func (o *ContactActivity) SetLastFailedNil(b bool)`

 SetLastFailedNil sets the value for LastFailed to be an explicit nil

### UnsetLastFailed
`func (o *ContactActivity) UnsetLastFailed()`

UnsetLastFailed ensures that no value is present for LastFailed, not even an explicit nil
### GetLastIP

`func (o *ContactActivity) GetLastIP() string`

GetLastIP returns the LastIP field if non-nil, zero value otherwise.

### GetLastIPOk

`func (o *ContactActivity) GetLastIPOk() (*string, bool)`

GetLastIPOk returns a tuple with the LastIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIP

`func (o *ContactActivity) SetLastIP(v string)`

SetLastIP sets LastIP field to given value.

### HasLastIP

`func (o *ContactActivity) HasLastIP() bool`

HasLastIP returns a boolean if a field has been set.

### GetErrorCode

`func (o *ContactActivity) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ContactActivity) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ContactActivity) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ContactActivity) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ContactActivity) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ContactActivity) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetFriendlyErrorMessage

`func (o *ContactActivity) GetFriendlyErrorMessage() string`

GetFriendlyErrorMessage returns the FriendlyErrorMessage field if non-nil, zero value otherwise.

### GetFriendlyErrorMessageOk

`func (o *ContactActivity) GetFriendlyErrorMessageOk() (*string, bool)`

GetFriendlyErrorMessageOk returns a tuple with the FriendlyErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyErrorMessage

`func (o *ContactActivity) SetFriendlyErrorMessage(v string)`

SetFriendlyErrorMessage sets FriendlyErrorMessage field to given value.

### HasFriendlyErrorMessage

`func (o *ContactActivity) HasFriendlyErrorMessage() bool`

HasFriendlyErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


