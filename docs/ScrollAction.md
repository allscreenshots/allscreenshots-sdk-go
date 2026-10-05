# ScrollAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**ToBottom** | Pointer to **bool** |  | [optional] [default to false]
**Type** | **string** |  | 
**Y** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewScrollAction

`func NewScrollAction(type_ string, ) *ScrollAction`

NewScrollAction instantiates a new ScrollAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrollActionWithDefaults

`func NewScrollActionWithDefaults() *ScrollAction`

NewScrollActionWithDefaults instantiates a new ScrollAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptional

`func (o *ScrollAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ScrollAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ScrollAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ScrollAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *ScrollAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ScrollAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ScrollAction) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ScrollAction) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *ScrollAction) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *ScrollAction) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetTimeout

`func (o *ScrollAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ScrollAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ScrollAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ScrollAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ScrollAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ScrollAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetToBottom

`func (o *ScrollAction) GetToBottom() bool`

GetToBottom returns the ToBottom field if non-nil, zero value otherwise.

### GetToBottomOk

`func (o *ScrollAction) GetToBottomOk() (*bool, bool)`

GetToBottomOk returns a tuple with the ToBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBottom

`func (o *ScrollAction) SetToBottom(v bool)`

SetToBottom sets ToBottom field to given value.

### HasToBottom

`func (o *ScrollAction) HasToBottom() bool`

HasToBottom returns a boolean if a field has been set.

### GetType

`func (o *ScrollAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScrollAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScrollAction) SetType(v string)`

SetType sets Type field to given value.


### GetY

`func (o *ScrollAction) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ScrollAction) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ScrollAction) SetY(v int32)`

SetY sets Y field to given value.

### HasY

`func (o *ScrollAction) HasY() bool`

HasY returns a boolean if a field has been set.

### SetYNil

`func (o *ScrollAction) SetYNil(b bool)`

 SetYNil sets the value for Y to be an explicit nil

### UnsetY
`func (o *ScrollAction) UnsetY()`

UnsetY ensures that no value is present for Y, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


