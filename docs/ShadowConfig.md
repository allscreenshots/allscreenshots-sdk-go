# ShadowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blur** | Pointer to **int32** |  | [optional] [default to 10]
**Color** | Pointer to **string** |  | [optional] [default to "rgba(0,0,0,0.2)"]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**OffsetX** | Pointer to **int32** |  | [optional] [default to 0]
**OffsetY** | Pointer to **int32** |  | [optional] [default to 4]

## Methods

### NewShadowConfig

`func NewShadowConfig() *ShadowConfig`

NewShadowConfig instantiates a new ShadowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShadowConfigWithDefaults

`func NewShadowConfigWithDefaults() *ShadowConfig`

NewShadowConfigWithDefaults instantiates a new ShadowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlur

`func (o *ShadowConfig) GetBlur() int32`

GetBlur returns the Blur field if non-nil, zero value otherwise.

### GetBlurOk

`func (o *ShadowConfig) GetBlurOk() (*int32, bool)`

GetBlurOk returns a tuple with the Blur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlur

`func (o *ShadowConfig) SetBlur(v int32)`

SetBlur sets Blur field to given value.

### HasBlur

`func (o *ShadowConfig) HasBlur() bool`

HasBlur returns a boolean if a field has been set.

### GetColor

`func (o *ShadowConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ShadowConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ShadowConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ShadowConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetEnabled

`func (o *ShadowConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ShadowConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ShadowConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ShadowConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOffsetX

`func (o *ShadowConfig) GetOffsetX() int32`

GetOffsetX returns the OffsetX field if non-nil, zero value otherwise.

### GetOffsetXOk

`func (o *ShadowConfig) GetOffsetXOk() (*int32, bool)`

GetOffsetXOk returns a tuple with the OffsetX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetX

`func (o *ShadowConfig) SetOffsetX(v int32)`

SetOffsetX sets OffsetX field to given value.

### HasOffsetX

`func (o *ShadowConfig) HasOffsetX() bool`

HasOffsetX returns a boolean if a field has been set.

### GetOffsetY

`func (o *ShadowConfig) GetOffsetY() int32`

GetOffsetY returns the OffsetY field if non-nil, zero value otherwise.

### GetOffsetYOk

`func (o *ShadowConfig) GetOffsetYOk() (*int32, bool)`

GetOffsetYOk returns a tuple with the OffsetY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetY

`func (o *ShadowConfig) SetOffsetY(v int32)`

SetOffsetY sets OffsetY field to given value.

### HasOffsetY

`func (o *ShadowConfig) HasOffsetY() bool`

HasOffsetY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


