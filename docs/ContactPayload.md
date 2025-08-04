# ContactPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Proper email address. | 
**Status** | Pointer to [**ContactStatus**](ContactStatus.md) |  | [optional] [default to CONTACTSTATUS_TRANSACTIONAL]
**FirstName** | Pointer to **string** | First name. | [optional] 
**LastName** | Pointer to **string** | Last name. | [optional] 
**CustomFields** | Pointer to **map[string]string** | A key-value collection of custom contact fields which can be used in the system. Only already existing custom fields will be saved. | [optional] 
**Consent** | Pointer to [**ConsentData**](ConsentData.md) |  | [optional] 

## Methods

### NewContactPayload

`func NewContactPayload(email string, ) *ContactPayload`

NewContactPayload instantiates a new ContactPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPayloadWithDefaults

`func NewContactPayloadWithDefaults() *ContactPayload`

NewContactPayloadWithDefaults instantiates a new ContactPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ContactPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactPayload) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStatus

`func (o *ContactPayload) GetStatus() ContactStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContactPayload) GetStatusOk() (*ContactStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContactPayload) SetStatus(v ContactStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContactPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFirstName

`func (o *ContactPayload) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactPayload) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactPayload) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactPayload) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ContactPayload) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactPayload) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactPayload) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactPayload) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCustomFields

`func (o *ContactPayload) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ContactPayload) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ContactPayload) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ContactPayload) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConsent

`func (o *ContactPayload) GetConsent() ConsentData`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *ContactPayload) GetConsentOk() (*ConsentData, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *ContactPayload) SetConsent(v ConsentData)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *ContactPayload) HasConsent() bool`

HasConsent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


