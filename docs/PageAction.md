# PageAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | **NullableString** |  | 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 
**Clear** | Pointer to **bool** |  | [optional] [default to true]
**Text** | **string** |  | 
**Key** | **string** |  | 
**Value** | **string** |  | 
**ToBottom** | Pointer to **bool** |  | [optional] [default to false]
**Y** | Pointer to **NullableInt32** |  | [optional] 
**Ms** | Pointer to **NullableInt32** |  | [optional] 
**NetworkIdle** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPageAction

`func NewPageAction(selector NullableString, type_ string, text string, key string, value string, ) *PageAction`

NewPageAction instantiates a new PageAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageActionWithDefaults

`func NewPageActionWithDefaults() *PageAction`

NewPageActionWithDefaults instantiates a new PageAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptional

`func (o *PageAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *PageAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *PageAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *PageAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *PageAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *PageAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *PageAction) SetSelector(v string)`

SetSelector sets Selector field to given value.


### SetSelectorNil

`func (o *PageAction) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *PageAction) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetTimeout

`func (o *PageAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PageAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PageAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PageAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PageAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PageAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *PageAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageAction) SetType(v string)`

SetType sets Type field to given value.


### GetClear

`func (o *PageAction) GetClear() bool`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *PageAction) GetClearOk() (*bool, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *PageAction) SetClear(v bool)`

SetClear sets Clear field to given value.

### HasClear

`func (o *PageAction) HasClear() bool`

HasClear returns a boolean if a field has been set.

### GetText

`func (o *PageAction) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PageAction) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PageAction) SetText(v string)`

SetText sets Text field to given value.


### GetKey

`func (o *PageAction) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PageAction) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PageAction) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *PageAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PageAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PageAction) SetValue(v string)`

SetValue sets Value field to given value.


### GetToBottom

`func (o *PageAction) GetToBottom() bool`

GetToBottom returns the ToBottom field if non-nil, zero value otherwise.

### GetToBottomOk

`func (o *PageAction) GetToBottomOk() (*bool, bool)`

GetToBottomOk returns a tuple with the ToBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBottom

`func (o *PageAction) SetToBottom(v bool)`

SetToBottom sets ToBottom field to given value.

### HasToBottom

`func (o *PageAction) HasToBottom() bool`

HasToBottom returns a boolean if a field has been set.

### GetY

`func (o *PageAction) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PageAction) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PageAction) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *PageAction) HasY() bool`

HasY returns a boolean if a field has been set.

### SetYNil

`func (o *PageAction) SetYNil(b bool)`

 SetYNil sets the value for Y to be an explicit nil

### UnsetY
`func (o *PageAction) UnsetY()`

UnsetY ensures that no value is present for Y, not even an explicit nil
### GetMs

`func (o *PageAction) GetMs() int32`

GetMs returns the Ms field if non-nil, zero value otherwise.

### GetMsOk

`func (o *PageAction) GetMsOk() (*int32, bool)`

GetMsOk returns a tuple with the Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMs

`func (o *PageAction) SetMs(v int32)`

SetMs sets Ms field to given value.

### HasMs

`func (o *PageAction) HasMs() bool`

HasMs returns a boolean if a field has been set.

### SetMsNil

`func (o *PageAction) SetMsNil(b bool)`

 SetMsNil sets the value for Ms to be an explicit nil

### UnsetMs
`func (o *PageAction) UnsetMs()`

UnsetMs ensures that no value is present for Ms, not even an explicit nil
### GetNetworkIdle

`func (o *PageAction) GetNetworkIdle() bool`

GetNetworkIdle returns the NetworkIdle field if non-nil, zero value otherwise.

### GetNetworkIdleOk

`func (o *PageAction) GetNetworkIdleOk() (*bool, bool)`

GetNetworkIdleOk returns a tuple with the NetworkIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdle

`func (o *PageAction) SetNetworkIdle(v bool)`

SetNetworkIdle sets NetworkIdle field to given value.

### HasNetworkIdle

`func (o *PageAction) HasNetworkIdle() bool`

HasNetworkIdle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


