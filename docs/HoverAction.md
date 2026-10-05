# HoverAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | **string** |  | 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewHoverAction

`func NewHoverAction(selector string, type_ string, ) *HoverAction`

NewHoverAction instantiates a new HoverAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoverActionWithDefaults

`func NewHoverActionWithDefaults() *HoverAction`

NewHoverActionWithDefaults instantiates a new HoverAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptional

`func (o *HoverAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *HoverAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *HoverAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *HoverAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *HoverAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *HoverAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *HoverAction) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetTimeout

`func (o *HoverAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HoverAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HoverAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HoverAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *HoverAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *HoverAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *HoverAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HoverAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HoverAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


