# WaitAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ms** | Pointer to **NullableInt32** |  | [optional] 
**NetworkIdle** | Pointer to **bool** |  | [optional] [default to false]
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewWaitAction

`func NewWaitAction(type_ string, ) *WaitAction`

NewWaitAction instantiates a new WaitAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitActionWithDefaults

`func NewWaitActionWithDefaults() *WaitAction`

NewWaitActionWithDefaults instantiates a new WaitAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMs

`func (o *WaitAction) GetMs() int32`

GetMs returns the Ms field if non-nil, zero value otherwise.

### GetMsOk

`func (o *WaitAction) GetMsOk() (*int32, bool)`

GetMsOk returns a tuple with the Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMs

`func (o *WaitAction) SetMs(v int32)`

SetMs sets Ms field to given value.

### HasMs

`func (o *WaitAction) HasMs() bool`

HasMs returns a boolean if a field has been set.

### SetMsNil

`func (o *WaitAction) SetMsNil(b bool)`

 SetMsNil sets the value for Ms to be an explicit nil

### UnsetMs
`func (o *WaitAction) UnsetMs()`

UnsetMs ensures that no value is present for Ms, not even an explicit nil
### GetNetworkIdle

`func (o *WaitAction) GetNetworkIdle() bool`

GetNetworkIdle returns the NetworkIdle field if non-nil, zero value otherwise.

### GetNetworkIdleOk

`func (o *WaitAction) GetNetworkIdleOk() (*bool, bool)`

GetNetworkIdleOk returns a tuple with the NetworkIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdle

`func (o *WaitAction) SetNetworkIdle(v bool)`

SetNetworkIdle sets NetworkIdle field to given value.

### HasNetworkIdle

`func (o *WaitAction) HasNetworkIdle() bool`

HasNetworkIdle returns a boolean if a field has been set.

### GetOptional

`func (o *WaitAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *WaitAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *WaitAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *WaitAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *WaitAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *WaitAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *WaitAction) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *WaitAction) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *WaitAction) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *WaitAction) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetTimeout

`func (o *WaitAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WaitAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WaitAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WaitAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *WaitAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *WaitAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *WaitAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WaitAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WaitAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


