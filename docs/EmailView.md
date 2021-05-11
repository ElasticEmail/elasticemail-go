# EmailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body (HTML, otherwise plain text) of email | [optional] 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**From** | Pointer to **string** | From email address | [optional] 

## Methods

### NewEmailView

`func NewEmailView() *EmailView`

NewEmailView instantiates a new EmailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailViewWithDefaults

`func NewEmailViewWithDefaults() *EmailView`

NewEmailViewWithDefaults instantiates a new EmailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSubject

`func (o *EmailView) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailView) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailView) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailView) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFrom

`func (o *EmailView) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailView) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailView) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailView) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


