# SmtpCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to [**AccessLevel**](AccessLevel.md) |  | [optional] [default to ACCESSLEVEL_NONE]
**Name** | Pointer to **string** | Name of the key. | [optional] 
**DateCreated** | Pointer to **time.Time** | Date this SmtpCredential was created. | [optional] 
**LastUse** | Pointer to **NullableTime** | Date this SmtpCredential was last used. | [optional] 
**Expires** | Pointer to **NullableTime** | Date this SmtpCredential expires. | [optional] 
**RestrictAccessToIPRange** | Pointer to **[]string** | Which IPs can use this SmtpCredential | [optional] 

## Methods

### NewSmtpCredentials

`func NewSmtpCredentials() *SmtpCredentials`

NewSmtpCredentials instantiates a new SmtpCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpCredentialsWithDefaults

`func NewSmtpCredentialsWithDefaults() *SmtpCredentials`

NewSmtpCredentialsWithDefaults instantiates a new SmtpCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *SmtpCredentials) GetAccessLevel() AccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *SmtpCredentials) GetAccessLevelOk() (*AccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *SmtpCredentials) SetAccessLevel(v AccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *SmtpCredentials) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetName

`func (o *SmtpCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmtpCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmtpCredentials) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmtpCredentials) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *SmtpCredentials) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SmtpCredentials) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SmtpCredentials) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SmtpCredentials) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUse

`func (o *SmtpCredentials) GetLastUse() time.Time`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *SmtpCredentials) GetLastUseOk() (*time.Time, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *SmtpCredentials) SetLastUse(v time.Time)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *SmtpCredentials) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### SetLastUseNil

`func (o *SmtpCredentials) SetLastUseNil(b bool)`

 SetLastUseNil sets the value for LastUse to be an explicit nil

### UnsetLastUse
`func (o *SmtpCredentials) UnsetLastUse()`

UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
### GetExpires

`func (o *SmtpCredentials) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SmtpCredentials) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SmtpCredentials) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SmtpCredentials) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *SmtpCredentials) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *SmtpCredentials) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetRestrictAccessToIPRange

`func (o *SmtpCredentials) GetRestrictAccessToIPRange() []string`

GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field if non-nil, zero value otherwise.

### GetRestrictAccessToIPRangeOk

`func (o *SmtpCredentials) GetRestrictAccessToIPRangeOk() (*[]string, bool)`

GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAccessToIPRange

`func (o *SmtpCredentials) SetRestrictAccessToIPRange(v []string)`

SetRestrictAccessToIPRange sets RestrictAccessToIPRange field to given value.

### HasRestrictAccessToIPRange

`func (o *SmtpCredentials) HasRestrictAccessToIPRange() bool`

HasRestrictAccessToIPRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


