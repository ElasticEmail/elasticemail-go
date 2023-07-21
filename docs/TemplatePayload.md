# TemplatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Template name | 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**Body** | Pointer to [**[]BodyPart**](BodyPart.md) | Email content of this template | [optional] 
**TemplateScope** | Pointer to [**TemplateScope**](TemplateScope.md) |  | [optional] [default to TEMPLATESCOPE_PERSONAL]

## Methods

### NewTemplatePayload

`func NewTemplatePayload(name string, ) *TemplatePayload`

NewTemplatePayload instantiates a new TemplatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePayloadWithDefaults

`func NewTemplatePayloadWithDefaults() *TemplatePayload`

NewTemplatePayloadWithDefaults instantiates a new TemplatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatePayload) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *TemplatePayload) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplatePayload) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplatePayload) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplatePayload) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *TemplatePayload) GetBody() []BodyPart`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplatePayload) GetBodyOk() (*[]BodyPart, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplatePayload) SetBody(v []BodyPart)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplatePayload) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTemplateScope

`func (o *TemplatePayload) GetTemplateScope() TemplateScope`

GetTemplateScope returns the TemplateScope field if non-nil, zero value otherwise.

### GetTemplateScopeOk

`func (o *TemplatePayload) GetTemplateScopeOk() (*TemplateScope, bool)`

GetTemplateScopeOk returns a tuple with the TemplateScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateScope

`func (o *TemplatePayload) SetTemplateScope(v TemplateScope)`

SetTemplateScope sets TemplateScope field to given value.

### HasTemplateScope

`func (o *TemplatePayload) HasTemplateScope() bool`

HasTemplateScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


