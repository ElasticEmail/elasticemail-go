# EmailsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** | SQL-like rule. Sending &#39;All&#39; as a value loads all resources of the given type. Help for building a segment rule can be found here: https://help.elasticemail.com/en/articles/5162182-segment-rules | [optional] 
**Emails** | Pointer to **[]string** | Comma delimited list of contact emails | [optional] 

## Methods

### NewEmailsPayload

`func NewEmailsPayload() *EmailsPayload`

NewEmailsPayload instantiates a new EmailsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailsPayloadWithDefaults

`func NewEmailsPayloadWithDefaults() *EmailsPayload`

NewEmailsPayloadWithDefaults instantiates a new EmailsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *EmailsPayload) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *EmailsPayload) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *EmailsPayload) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *EmailsPayload) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetEmails

`func (o *EmailsPayload) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *EmailsPayload) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *EmailsPayload) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *EmailsPayload) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


