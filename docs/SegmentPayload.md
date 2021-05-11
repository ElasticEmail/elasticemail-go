# SegmentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Segment name | [optional] 
**Rule** | Pointer to **string** | SQL-like rule to determine which Contacts belong to this Segment. Help for building a segment rule can be found here: https://help.elasticemail.com/en/articles/5162182-segment-rules | [optional] 

## Methods

### NewSegmentPayload

`func NewSegmentPayload() *SegmentPayload`

NewSegmentPayload instantiates a new SegmentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentPayloadWithDefaults

`func NewSegmentPayloadWithDefaults() *SegmentPayload`

NewSegmentPayloadWithDefaults instantiates a new SegmentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SegmentPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRule

`func (o *SegmentPayload) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *SegmentPayload) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *SegmentPayload) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *SegmentPayload) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


