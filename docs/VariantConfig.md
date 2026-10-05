# VariantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomCss** | Pointer to **NullableString** |  | [optional] 
**DarkMode** | Pointer to **NullableBool** |  | [optional] 
**Delay** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**FullPage** | Pointer to **NullableBool** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 

## Methods

### NewVariantConfig

`func NewVariantConfig() *VariantConfig`

NewVariantConfig instantiates a new VariantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantConfigWithDefaults

`func NewVariantConfigWithDefaults() *VariantConfig`

NewVariantConfigWithDefaults instantiates a new VariantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomCss

`func (o *VariantConfig) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *VariantConfig) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *VariantConfig) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *VariantConfig) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *VariantConfig) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *VariantConfig) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *VariantConfig) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *VariantConfig) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *VariantConfig) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *VariantConfig) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### SetDarkModeNil

`func (o *VariantConfig) SetDarkModeNil(b bool)`

 SetDarkModeNil sets the value for DarkMode to be an explicit nil

### UnsetDarkMode
`func (o *VariantConfig) UnsetDarkMode()`

UnsetDarkMode ensures that no value is present for DarkMode, not even an explicit nil
### GetDelay

`func (o *VariantConfig) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *VariantConfig) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *VariantConfig) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *VariantConfig) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *VariantConfig) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *VariantConfig) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetDevice

`func (o *VariantConfig) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VariantConfig) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VariantConfig) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VariantConfig) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VariantConfig) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VariantConfig) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFullPage

`func (o *VariantConfig) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *VariantConfig) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *VariantConfig) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *VariantConfig) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### SetFullPageNil

`func (o *VariantConfig) SetFullPageNil(b bool)`

 SetFullPageNil sets the value for FullPage to be an explicit nil

### UnsetFullPage
`func (o *VariantConfig) UnsetFullPage()`

UnsetFullPage ensures that no value is present for FullPage, not even an explicit nil
### GetId

`func (o *VariantConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariantConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariantConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VariantConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VariantConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VariantConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *VariantConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VariantConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VariantConfig) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VariantConfig) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *VariantConfig) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *VariantConfig) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetViewport

`func (o *VariantConfig) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *VariantConfig) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *VariantConfig) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *VariantConfig) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *VariantConfig) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *VariantConfig) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


