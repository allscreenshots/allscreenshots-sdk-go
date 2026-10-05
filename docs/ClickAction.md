# ClickAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optional** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | **string** |  | 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewClickAction

`func NewClickAction(selector string, type_ string, ) *ClickAction`

NewClickAction instantiates a new ClickAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickActionWithDefaults

`func NewClickActionWithDefaults() *ClickAction`

NewClickActionWithDefaults instantiates a new ClickAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptional

`func (o *ClickAction) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *ClickAction) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *ClickAction) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *ClickAction) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSelector

`func (o *ClickAction) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ClickAction) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ClickAction) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetTimeout

`func (o *ClickAction) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ClickAction) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ClickAction) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ClickAction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ClickAction) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ClickAction) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *ClickAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClickAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClickAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


