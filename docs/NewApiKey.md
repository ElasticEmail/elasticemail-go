# NewApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Unique token to be used in the system | [optional] 
**AccessLevel** | Pointer to [**[]AccessLevel**](AccessLevel.md) | Access level or permission to be assigned to this ApiKey. | [optional] 
**Name** | Pointer to **string** | Name of the ApiKey. | [optional] 
**DateCreated** | Pointer to **time.Time** | Date this ApiKey was created. | [optional] 
**LastUse** | Pointer to **NullableTime** | Date this ApiKey was last used. | [optional] 
**Expires** | Pointer to **NullableTime** | Date this ApiKey expires. | [optional] 
**RestrictAccessToIPRange** | Pointer to **[]string** | Which IPs can use this ApiKey | [optional] 

## Methods

### NewNewApiKey

`func NewNewApiKey() *NewApiKey`

NewNewApiKey instantiates a new NewApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApiKeyWithDefaults

`func NewNewApiKeyWithDefaults() *NewApiKey`

NewNewApiKeyWithDefaults instantiates a new NewApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *NewApiKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NewApiKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NewApiKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *NewApiKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccessLevel

`func (o *NewApiKey) GetAccessLevel() []AccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *NewApiKey) GetAccessLevelOk() (*[]AccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *NewApiKey) SetAccessLevel(v []AccessLevel)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *NewApiKey) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetName

`func (o *NewApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NewApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *NewApiKey) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NewApiKey) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NewApiKey) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NewApiKey) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUse

`func (o *NewApiKey) GetLastUse() time.Time`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *NewApiKey) GetLastUseOk() (*time.Time, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *NewApiKey) SetLastUse(v time.Time)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *NewApiKey) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### SetLastUseNil

`func (o *NewApiKey) SetLastUseNil(b bool)`

 SetLastUseNil sets the value for LastUse to be an explicit nil

### UnsetLastUse
`func (o *NewApiKey) UnsetLastUse()`

UnsetLastUse ensures that no value is present for LastUse, not even an explicit nil
### GetExpires

`func (o *NewApiKey) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *NewApiKey) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *NewApiKey) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *NewApiKey) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *NewApiKey) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *NewApiKey) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetRestrictAccessToIPRange

`func (o *NewApiKey) GetRestrictAccessToIPRange() []string`

GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field if non-nil, zero value otherwise.

### GetRestrictAccessToIPRangeOk

`func (o *NewApiKey) GetRestrictAccessToIPRangeOk() (*[]string, bool)`

GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAccessToIPRange

`func (o *NewApiKey) SetRestrictAccessToIPRange(v []string)`

SetRestrictAccessToIPRange sets RestrictAccessToIPRange field to given value.

### HasRestrictAccessToIPRange

`func (o *NewApiKey) HasRestrictAccessToIPRange() bool`

HasRestrictAccessToIPRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


