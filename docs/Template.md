# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateType** | Pointer to [**TemplateType**](TemplateType.md) |  | [optional] [default to TEMPLATETYPE_RAW_HTML]
**Name** | Pointer to **string** | Template name | [optional] 
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**Body** | Pointer to [**[]BodyPart**](BodyPart.md) | Email content of this template | [optional] 
**TemplateScope** | Pointer to [**TemplateScope**](TemplateScope.md) |  | [optional] [default to TEMPLATESCOPE_PERSONAL]

## Methods

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateType

`func (o *Template) GetTemplateType() TemplateType`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Template) GetTemplateTypeOk() (*TemplateType, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Template) SetTemplateType(v TemplateType)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *Template) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateAdded

`func (o *Template) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *Template) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *Template) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *Template) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSubject

`func (o *Template) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Template) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Template) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Template) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *Template) GetBody() []BodyPart`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Template) GetBodyOk() (*[]BodyPart, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Template) SetBody(v []BodyPart)`

SetBody sets Body field to given value.

### HasBody

`func (o *Template) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTemplateScope

`func (o *Template) GetTemplateScope() TemplateScope`

GetTemplateScope returns the TemplateScope field if non-nil, zero value otherwise.

### GetTemplateScopeOk

`func (o *Template) GetTemplateScopeOk() (*TemplateScope, bool)`

GetTemplateScopeOk returns a tuple with the TemplateScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateScope

`func (o *Template) SetTemplateScope(v TemplateScope)`

SetTemplateScope sets TemplateScope field to given value.

### HasTemplateScope

`func (o *Template) HasTemplateScope() bool`

HasTemplateScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


