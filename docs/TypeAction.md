# TypeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clear** | Pointer to **bool** |  | [optional] [default to true]
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | **string** |  | 
**Text** | **string** |  | 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewTypeAction

`func NewTypeAction(selector string, text string, type_ string, ) *TypeAction`

NewTypeAction instantiates a new TypeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeActionWithDefaults

`func NewTypeActionWithDefaults() *TypeAction`

NewTypeActionWithDefaults instantiates a new TypeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClear

`func (o *TypeAction) GetClear() bool`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *TypeAction) GetClearOk() (*bool, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *TypeAction) SetClear(v bool)`

SetClear sets Clear field to given value.

### HasClear

`func (o *TypeAction) HasClear() bool`

HasClear returns a boolean if a field has been set.

### GetOptional

`func (o *TypeAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *TypeAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *TypeAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *TypeAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *TypeAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TypeAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TypeAction) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetText

`func (o *TypeAction) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TypeAction) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TypeAction) SetText(v string)`

SetText sets Text field to given value.


### GetTimeout

`func (o *TypeAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TypeAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TypeAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TypeAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *TypeAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *TypeAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *TypeAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypeAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypeAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


