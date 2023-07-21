# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Proper email address. | [optional] 
**Status** | Pointer to [**ContactStatus**](ContactStatus.md) |  | [optional] [default to CONTACTSTATUS_TRANSACTIONAL]
**FirstName** | Pointer to **string** | First name. | [optional] 
**LastName** | Pointer to **string** | Last name. | [optional] 
**CustomFields** | Pointer to **map[string]string** | A key-value collection of custom contact fields which can be used in the system. | [optional] 
**Consent** | Pointer to [**ConsentData**](ConsentData.md) |  | [optional] 
**Source** | Pointer to [**ContactSource**](ContactSource.md) |  | [optional] [default to CONTACTSOURCE_DELIVERY_API]
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**DateUpdated** | Pointer to **NullableTime** | Last change date | [optional] 
**StatusChangeDate** | Pointer to **NullableTime** | Date of last status change. | [optional] 
**Activity** | Pointer to [**ContactActivity**](ContactActivity.md) |  | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *Contact) GetStatus() ContactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Contact) GetStatusOk() (*ContactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Contact) SetStatus(v ContactStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Contact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCustomFields

`func (o *Contact) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Contact) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Contact) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Contact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConsent

`func (o *Contact) GetConsent() ConsentData`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *Contact) GetConsentOk() (*ConsentData, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *Contact) SetConsent(v ConsentData)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *Contact) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetSource

`func (o *Contact) GetSource() ContactSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Contact) GetSourceOk() (*ContactSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Contact) SetSource(v ContactSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Contact) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDateAdded

`func (o *Contact) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *Contact) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *Contact) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *Contact) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateUpdated

`func (o *Contact) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *Contact) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *Contact) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *Contact) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *Contact) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *Contact) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetStatusChangeDate

`func (o *Contact) GetStatusChangeDate() time.Time`

GetStatusChangeDate returns the StatusChangeDate field if non-nil, zero value otherwise.

### GetStatusChangeDateOk

`func (o *Contact) GetStatusChangeDateOk() (*time.Time, bool)`

GetStatusChangeDateOk returns a tuple with the StatusChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangeDate

`func (o *Contact) SetStatusChangeDate(v time.Time)`

SetStatusChangeDate sets StatusChangeDate field to given value.

### HasStatusChangeDate

`func (o *Contact) HasStatusChangeDate() bool`

HasStatusChangeDate returns a boolean if a field has been set.

### SetStatusChangeDateNil

`func (o *Contact) SetStatusChangeDateNil(b bool)`

 SetStatusChangeDateNil sets the value for StatusChangeDate to be an explicit nil

### UnsetStatusChangeDate
`func (o *Contact) UnsetStatusChangeDate()`

UnsetStatusChangeDate ensures that no value is present for StatusChangeDate, not even an explicit nil
### GetActivity

`func (o *Contact) GetActivity() ContactActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Contact) GetActivityOk() (*ContactActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Contact) SetActivity(v ContactActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Contact) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


