# EmailValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Local part of an email | [optional] 
**Domain** | Pointer to **string** | Name of selected domain. | [optional] 
**Email** | Pointer to **string** | Full email address that was verified | [optional] 
**SuggestedSpelling** | Pointer to **string** | Suggested spelling if a possible mistake was found | [optional] 
**Disposable** | Pointer to **bool** | Does the email have a temporary domain | [optional] 
**Role** | Pointer to **bool** | Is an email a role email (e.g. info@, noreply@ etc.) | [optional] 
**Reason** | Pointer to **string** | All detected issues | [optional] 
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**Result** | Pointer to [**EmailValidationStatus**](EmailValidationStatus.md) |  | [optional] [default to EMAILVALIDATIONSTATUS_NONE]

## Methods

### NewEmailValidationResult

`func NewEmailValidationResult() *EmailValidationResult`

NewEmailValidationResult instantiates a new EmailValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailValidationResultWithDefaults

`func NewEmailValidationResultWithDefaults() *EmailValidationResult`

NewEmailValidationResultWithDefaults instantiates a new EmailValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *EmailValidationResult) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *EmailValidationResult) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *EmailValidationResult) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *EmailValidationResult) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDomain

`func (o *EmailValidationResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailValidationResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailValidationResult) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailValidationResult) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmail

`func (o *EmailValidationResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailValidationResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailValidationResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailValidationResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSuggestedSpelling

`func (o *EmailValidationResult) GetSuggestedSpelling() string`

GetSuggestedSpelling returns the SuggestedSpelling field if non-nil, zero value otherwise.

### GetSuggestedSpellingOk

`func (o *EmailValidationResult) GetSuggestedSpellingOk() (*string, bool)`

GetSuggestedSpellingOk returns a tuple with the SuggestedSpelling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedSpelling

`func (o *EmailValidationResult) SetSuggestedSpelling(v string)`

SetSuggestedSpelling sets SuggestedSpelling field to given value.

### HasSuggestedSpelling

`func (o *EmailValidationResult) HasSuggestedSpelling() bool`

HasSuggestedSpelling returns a boolean if a field has been set.

### GetDisposable

`func (o *EmailValidationResult) GetDisposable() bool`

GetDisposable returns the Disposable field if non-nil, zero value otherwise.

### GetDisposableOk

`func (o *EmailValidationResult) GetDisposableOk() (*bool, bool)`

GetDisposableOk returns a tuple with the Disposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposable

`func (o *EmailValidationResult) SetDisposable(v bool)`

SetDisposable sets Disposable field to given value.

### HasDisposable

`func (o *EmailValidationResult) HasDisposable() bool`

HasDisposable returns a boolean if a field has been set.

### GetRole

`func (o *EmailValidationResult) GetRole() bool`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmailValidationResult) GetRoleOk() (*bool, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmailValidationResult) SetRole(v bool)`

SetRole sets Role field to given value.

### HasRole

`func (o *EmailValidationResult) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetReason

`func (o *EmailValidationResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailValidationResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailValidationResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailValidationResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDateAdded

`func (o *EmailValidationResult) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *EmailValidationResult) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *EmailValidationResult) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *EmailValidationResult) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetResult

`func (o *EmailValidationResult) GetResult() EmailValidationStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EmailValidationResult) GetResultOk() (*EmailValidationStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EmailValidationResult) SetResult(v EmailValidationStatus)`

SetResult sets Result field to given value.

### HasResult

`func (o *EmailValidationResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


