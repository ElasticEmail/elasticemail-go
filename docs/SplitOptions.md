# SplitOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptimizeFor** | Pointer to [**SplitOptimizationType**](SplitOptimizationType.md) |  | [optional] [default to SPLITOPTIMIZATIONTYPE_OPENS]
**OptimizePeriodMinutes** | Pointer to **int32** | For how long should the results be measured until determining the winner template (content) | [optional] 

## Methods

### NewSplitOptions

`func NewSplitOptions() *SplitOptions`

NewSplitOptions instantiates a new SplitOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitOptionsWithDefaults

`func NewSplitOptionsWithDefaults() *SplitOptions`

NewSplitOptionsWithDefaults instantiates a new SplitOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptimizeFor

`func (o *SplitOptions) GetOptimizeFor() SplitOptimizationType`

GetOptimizeFor returns the OptimizeFor field if non-nil, zero value otherwise.

### GetOptimizeForOk

`func (o *SplitOptions) GetOptimizeForOk() (*SplitOptimizationType, bool)`

GetOptimizeForOk returns a tuple with the OptimizeFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeFor

`func (o *SplitOptions) SetOptimizeFor(v SplitOptimizationType)`

SetOptimizeFor sets OptimizeFor field to given value.

### HasOptimizeFor

`func (o *SplitOptions) HasOptimizeFor() bool`

HasOptimizeFor returns a boolean if a field has been set.

### GetOptimizePeriodMinutes

`func (o *SplitOptions) GetOptimizePeriodMinutes() int32`

GetOptimizePeriodMinutes returns the OptimizePeriodMinutes field if non-nil, zero value otherwise.

### GetOptimizePeriodMinutesOk

`func (o *SplitOptions) GetOptimizePeriodMinutesOk() (*int32, bool)`

GetOptimizePeriodMinutesOk returns a tuple with the OptimizePeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizePeriodMinutes

`func (o *SplitOptions) SetOptimizePeriodMinutes(v int32)`

SetOptimizePeriodMinutes sets OptimizePeriodMinutes field to given value.

### HasOptimizePeriodMinutes

`func (o *SplitOptions) HasOptimizePeriodMinutes() bool`

HasOptimizePeriodMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


