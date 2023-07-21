# SubAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicAccountID** | Pointer to **string** | Public key for limited access to your Account such as contact/add so you can use it safely on public websites. | [optional] 
**Email** | Pointer to **string** | Proper email address. | [optional] 
**Settings** | Pointer to [**SubaccountSettingsInfo**](SubaccountSettingsInfo.md) |  | [optional] 
**LastActivity** | Pointer to **time.Time** | Date of last activity on Account | [optional] 
**EmailCredits** | Pointer to **int32** | Amount of email credits | [optional] 
**TotalEmailsSent** | Pointer to **int64** | Amount of emails sent from this Account | [optional] 
**Reputation** | Pointer to **float64** | Numeric reputation | [optional] 
**Status** | Pointer to [**AccountStatusEnum**](AccountStatusEnum.md) |  | [optional] [default to ACCOUNTSTATUSENUM_DISABLED]
**ContactsCount** | Pointer to **int32** | How many contacts this SubAccount has stored | [optional] 

## Methods

### NewSubAccountInfo

`func NewSubAccountInfo() *SubAccountInfo`

NewSubAccountInfo instantiates a new SubAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountInfoWithDefaults

`func NewSubAccountInfoWithDefaults() *SubAccountInfo`

NewSubAccountInfoWithDefaults instantiates a new SubAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicAccountID

`func (o *SubAccountInfo) GetPublicAccountID() string`

GetPublicAccountID returns the PublicAccountID field if non-nil, zero value otherwise.

### GetPublicAccountIDOk

`func (o *SubAccountInfo) GetPublicAccountIDOk() (*string, bool)`

GetPublicAccountIDOk returns a tuple with the PublicAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccountID

`func (o *SubAccountInfo) SetPublicAccountID(v string)`

SetPublicAccountID sets PublicAccountID field to given value.

### HasPublicAccountID

`func (o *SubAccountInfo) HasPublicAccountID() bool`

HasPublicAccountID returns a boolean if a field has been set.

### GetEmail

`func (o *SubAccountInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubAccountInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubAccountInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubAccountInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSettings

`func (o *SubAccountInfo) GetSettings() SubaccountSettingsInfo`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SubAccountInfo) GetSettingsOk() (*SubaccountSettingsInfo, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SubAccountInfo) SetSettings(v SubaccountSettingsInfo)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SubAccountInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLastActivity

`func (o *SubAccountInfo) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *SubAccountInfo) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *SubAccountInfo) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *SubAccountInfo) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetEmailCredits

`func (o *SubAccountInfo) GetEmailCredits() int32`

GetEmailCredits returns the EmailCredits field if non-nil, zero value otherwise.

### GetEmailCreditsOk

`func (o *SubAccountInfo) GetEmailCreditsOk() (*int32, bool)`

GetEmailCreditsOk returns a tuple with the EmailCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCredits

`func (o *SubAccountInfo) SetEmailCredits(v int32)`

SetEmailCredits sets EmailCredits field to given value.

### HasEmailCredits

`func (o *SubAccountInfo) HasEmailCredits() bool`

HasEmailCredits returns a boolean if a field has been set.

### GetTotalEmailsSent

`func (o *SubAccountInfo) GetTotalEmailsSent() int64`

GetTotalEmailsSent returns the TotalEmailsSent field if non-nil, zero value otherwise.

### GetTotalEmailsSentOk

`func (o *SubAccountInfo) GetTotalEmailsSentOk() (*int64, bool)`

GetTotalEmailsSentOk returns a tuple with the TotalEmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEmailsSent

`func (o *SubAccountInfo) SetTotalEmailsSent(v int64)`

SetTotalEmailsSent sets TotalEmailsSent field to given value.

### HasTotalEmailsSent

`func (o *SubAccountInfo) HasTotalEmailsSent() bool`

HasTotalEmailsSent returns a boolean if a field has been set.

### GetReputation

`func (o *SubAccountInfo) GetReputation() float64`

GetReputation returns the Reputation field if non-nil, zero value otherwise.

### GetReputationOk

`func (o *SubAccountInfo) GetReputationOk() (*float64, bool)`

GetReputationOk returns a tuple with the Reputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputation

`func (o *SubAccountInfo) SetReputation(v float64)`

SetReputation sets Reputation field to given value.

### HasReputation

`func (o *SubAccountInfo) HasReputation() bool`

HasReputation returns a boolean if a field has been set.

### GetStatus

`func (o *SubAccountInfo) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubAccountInfo) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubAccountInfo) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubAccountInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContactsCount

`func (o *SubAccountInfo) GetContactsCount() int32`

GetContactsCount returns the ContactsCount field if non-nil, zero value otherwise.

### GetContactsCountOk

`func (o *SubAccountInfo) GetContactsCountOk() (*int32, bool)`

GetContactsCountOk returns a tuple with the ContactsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactsCount

`func (o *SubAccountInfo) SetContactsCount(v int32)`

SetContactsCount sets ContactsCount field to given value.

### HasContactsCount

`func (o *SubAccountInfo) HasContactsCount() bool`

HasContactsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


