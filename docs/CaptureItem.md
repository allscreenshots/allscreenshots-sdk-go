# CaptureItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DarkMode** | Pointer to **NullableBool** |  | [optional] 
**Delay** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**FullPage** | Pointer to **NullableBool** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 

## Methods

### NewCaptureItem

`func NewCaptureItem(url string, ) *CaptureItem`

NewCaptureItem instantiates a new CaptureItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureItemWithDefaults

`func NewCaptureItemWithDefaults() *CaptureItem`

NewCaptureItemWithDefaults instantiates a new CaptureItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDarkMode

`func (o *CaptureItem) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *CaptureItem) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *CaptureItem) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *CaptureItem) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### SetDarkModeNil

`func (o *CaptureItem) SetDarkModeNil(b bool)`

 SetDarkModeNil sets the value for DarkMode to be an explicit nil

### UnsetDarkMode
`func (o *CaptureItem) UnsetDarkMode()`

UnsetDarkMode ensures that no value is present for DarkMode, not even an explicit nil
### GetDelay

`func (o *CaptureItem) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *CaptureItem) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *CaptureItem) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *CaptureItem) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *CaptureItem) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *CaptureItem) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetDevice

`func (o *CaptureItem) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CaptureItem) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CaptureItem) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CaptureItem) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *CaptureItem) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *CaptureItem) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFullPage

`func (o *CaptureItem) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *CaptureItem) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *CaptureItem) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *CaptureItem) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### SetFullPageNil

`func (o *CaptureItem) SetFullPageNil(b bool)`

 SetFullPageNil sets the value for FullPage to be an explicit nil

### UnsetFullPage
`func (o *CaptureItem) UnsetFullPage()`

UnsetFullPage ensures that no value is present for FullPage, not even an explicit nil
### GetId

`func (o *CaptureItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptureItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptureItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptureItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CaptureItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CaptureItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *CaptureItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CaptureItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CaptureItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CaptureItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CaptureItem) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CaptureItem) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetUrl

`func (o *CaptureItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CaptureItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CaptureItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetViewport

`func (o *CaptureItem) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *CaptureItem) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *CaptureItem) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *CaptureItem) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *CaptureItem) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *CaptureItem) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


