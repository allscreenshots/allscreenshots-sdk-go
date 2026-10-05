# BulkDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]PageAction**](PageAction.md) |  | [optional] 
**BlockAds** | Pointer to **bool** |  | [optional] [default to true]
**BlockCookieBanners** | Pointer to **bool** |  | [optional] [default to true]
**BlockLevel** | Pointer to **string** |  | [optional] [default to "none"]
**BlockPopups** | Pointer to **bool** |  | [optional] [default to true]
**CustomCss** | Pointer to **NullableString** |  | [optional] 
**DarkMode** | Pointer to **bool** |  | [optional] [default to false]
**Delay** | Pointer to **int32** |  | [optional] [default to 0]
**Device** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "png"]
**FullPage** | Pointer to **bool** |  | [optional] [default to false]
**Outputs** | Pointer to [**[]OutputSpec**](OutputSpec.md) |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**StealthMode** | Pointer to **bool** |  | [optional] [default to false]
**Timeout** | Pointer to **int32** |  | [optional] [default to 30000]
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitFor** | Pointer to **NullableString** |  | [optional] 
**WaitUntil** | Pointer to **string** |  | [optional] [default to "domcontentloaded"]

## Methods

### NewBulkDefaults

`func NewBulkDefaults() *BulkDefaults`

NewBulkDefaults instantiates a new BulkDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDefaultsWithDefaults

`func NewBulkDefaultsWithDefaults() *BulkDefaults`

NewBulkDefaultsWithDefaults instantiates a new BulkDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BulkDefaults) GetActions() []PageAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BulkDefaults) GetActionsOk() (*[]PageAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BulkDefaults) SetActions(v []PageAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BulkDefaults) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *BulkDefaults) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *BulkDefaults) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetBlockAds

`func (o *BulkDefaults) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *BulkDefaults) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *BulkDefaults) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *BulkDefaults) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetBlockCookieBanners

`func (o *BulkDefaults) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *BulkDefaults) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *BulkDefaults) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *BulkDefaults) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### GetBlockLevel

`func (o *BulkDefaults) GetBlockLevel() string`

GetBlockLevel returns the BlockLevel field if non-nil, zero value otherwise.

### GetBlockLevelOk

`func (o *BulkDefaults) GetBlockLevelOk() (*string, bool)`

GetBlockLevelOk returns a tuple with the BlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevel

`func (o *BulkDefaults) SetBlockLevel(v string)`

SetBlockLevel sets BlockLevel field to given value.

### HasBlockLevel

`func (o *BulkDefaults) HasBlockLevel() bool`

HasBlockLevel returns a boolean if a field has been set.

### GetBlockPopups

`func (o *BulkDefaults) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *BulkDefaults) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *BulkDefaults) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *BulkDefaults) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### GetCustomCss

`func (o *BulkDefaults) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *BulkDefaults) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *BulkDefaults) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *BulkDefaults) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *BulkDefaults) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *BulkDefaults) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *BulkDefaults) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *BulkDefaults) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *BulkDefaults) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *BulkDefaults) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDelay

`func (o *BulkDefaults) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *BulkDefaults) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *BulkDefaults) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *BulkDefaults) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDevice

`func (o *BulkDefaults) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkDefaults) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkDefaults) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BulkDefaults) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BulkDefaults) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BulkDefaults) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFormat

`func (o *BulkDefaults) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BulkDefaults) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BulkDefaults) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BulkDefaults) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFullPage

`func (o *BulkDefaults) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *BulkDefaults) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *BulkDefaults) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *BulkDefaults) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetOutputs

`func (o *BulkDefaults) GetOutputs() []OutputSpec`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BulkDefaults) GetOutputsOk() (*[]OutputSpec, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BulkDefaults) SetOutputs(v []OutputSpec)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BulkDefaults) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *BulkDefaults) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *BulkDefaults) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetQuality

`func (o *BulkDefaults) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BulkDefaults) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BulkDefaults) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BulkDefaults) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetStealthMode

`func (o *BulkDefaults) GetStealthMode() bool`

GetStealthMode returns the StealthMode field if non-nil, zero value otherwise.

### GetStealthModeOk

`func (o *BulkDefaults) GetStealthModeOk() (*bool, bool)`

GetStealthModeOk returns a tuple with the StealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthMode

`func (o *BulkDefaults) SetStealthMode(v bool)`

SetStealthMode sets StealthMode field to given value.

### HasStealthMode

`func (o *BulkDefaults) HasStealthMode() bool`

HasStealthMode returns a boolean if a field has been set.

### GetTimeout

`func (o *BulkDefaults) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BulkDefaults) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BulkDefaults) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BulkDefaults) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetViewport

`func (o *BulkDefaults) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *BulkDefaults) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *BulkDefaults) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *BulkDefaults) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *BulkDefaults) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *BulkDefaults) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil
### GetWaitFor

`func (o *BulkDefaults) GetWaitFor() string`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *BulkDefaults) GetWaitForOk() (*string, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *BulkDefaults) SetWaitFor(v string)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *BulkDefaults) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### SetWaitForNil

`func (o *BulkDefaults) SetWaitForNil(b bool)`

 SetWaitForNil sets the value for WaitFor to be an explicit nil

### UnsetWaitFor
`func (o *BulkDefaults) UnsetWaitFor()`

UnsetWaitFor ensures that no value is present for WaitFor, not even an explicit nil
### GetWaitUntil

`func (o *BulkDefaults) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *BulkDefaults) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *BulkDefaults) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *BulkDefaults) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


