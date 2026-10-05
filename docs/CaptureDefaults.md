# CaptureDefaults

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
**HideSelectors** | Pointer to **[]string** |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**StealthMode** | Pointer to **bool** |  | [optional] [default to false]
**Timeout** | Pointer to **int32** |  | [optional] [default to 30000]
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitFor** | Pointer to **NullableString** |  | [optional] 
**WaitUntil** | Pointer to **string** |  | [optional] [default to "domcontentloaded"]

## Methods

### NewCaptureDefaults

`func NewCaptureDefaults() *CaptureDefaults`

NewCaptureDefaults instantiates a new CaptureDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureDefaultsWithDefaults

`func NewCaptureDefaultsWithDefaults() *CaptureDefaults`

NewCaptureDefaultsWithDefaults instantiates a new CaptureDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CaptureDefaults) GetActions() []PageAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CaptureDefaults) GetActionsOk() (*[]PageAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CaptureDefaults) SetActions(v []PageAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CaptureDefaults) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *CaptureDefaults) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *CaptureDefaults) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetBlockAds

`func (o *CaptureDefaults) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *CaptureDefaults) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *CaptureDefaults) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *CaptureDefaults) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetBlockCookieBanners

`func (o *CaptureDefaults) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *CaptureDefaults) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *CaptureDefaults) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *CaptureDefaults) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### GetBlockLevel

`func (o *CaptureDefaults) GetBlockLevel() string`

GetBlockLevel returns the BlockLevel field if non-nil, zero value otherwise.

### GetBlockLevelOk

`func (o *CaptureDefaults) GetBlockLevelOk() (*string, bool)`

GetBlockLevelOk returns a tuple with the BlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevel

`func (o *CaptureDefaults) SetBlockLevel(v string)`

SetBlockLevel sets BlockLevel field to given value.

### HasBlockLevel

`func (o *CaptureDefaults) HasBlockLevel() bool`

HasBlockLevel returns a boolean if a field has been set.

### GetBlockPopups

`func (o *CaptureDefaults) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *CaptureDefaults) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *CaptureDefaults) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *CaptureDefaults) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### GetCustomCss

`func (o *CaptureDefaults) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *CaptureDefaults) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *CaptureDefaults) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *CaptureDefaults) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *CaptureDefaults) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *CaptureDefaults) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *CaptureDefaults) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *CaptureDefaults) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *CaptureDefaults) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *CaptureDefaults) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDelay

`func (o *CaptureDefaults) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *CaptureDefaults) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *CaptureDefaults) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *CaptureDefaults) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDevice

`func (o *CaptureDefaults) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CaptureDefaults) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CaptureDefaults) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CaptureDefaults) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *CaptureDefaults) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *CaptureDefaults) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFormat

`func (o *CaptureDefaults) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureDefaults) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureDefaults) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureDefaults) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFullPage

`func (o *CaptureDefaults) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *CaptureDefaults) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *CaptureDefaults) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *CaptureDefaults) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetHideSelectors

`func (o *CaptureDefaults) GetHideSelectors() []string`

GetHideSelectors returns the HideSelectors field if non-nil, zero value otherwise.

### GetHideSelectorsOk

`func (o *CaptureDefaults) GetHideSelectorsOk() (*[]string, bool)`

GetHideSelectorsOk returns a tuple with the HideSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSelectors

`func (o *CaptureDefaults) SetHideSelectors(v []string)`

SetHideSelectors sets HideSelectors field to given value.

### HasHideSelectors

`func (o *CaptureDefaults) HasHideSelectors() bool`

HasHideSelectors returns a boolean if a field has been set.

### SetHideSelectorsNil

`func (o *CaptureDefaults) SetHideSelectorsNil(b bool)`

 SetHideSelectorsNil sets the value for HideSelectors to be an explicit nil

### UnsetHideSelectors
`func (o *CaptureDefaults) UnsetHideSelectors()`

UnsetHideSelectors ensures that no value is present for HideSelectors, not even an explicit nil
### GetQuality

`func (o *CaptureDefaults) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *CaptureDefaults) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *CaptureDefaults) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *CaptureDefaults) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetStealthMode

`func (o *CaptureDefaults) GetStealthMode() bool`

GetStealthMode returns the StealthMode field if non-nil, zero value otherwise.

### GetStealthModeOk

`func (o *CaptureDefaults) GetStealthModeOk() (*bool, bool)`

GetStealthModeOk returns a tuple with the StealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthMode

`func (o *CaptureDefaults) SetStealthMode(v bool)`

SetStealthMode sets StealthMode field to given value.

### HasStealthMode

`func (o *CaptureDefaults) HasStealthMode() bool`

HasStealthMode returns a boolean if a field has been set.

### GetTimeout

`func (o *CaptureDefaults) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CaptureDefaults) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CaptureDefaults) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CaptureDefaults) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetViewport

`func (o *CaptureDefaults) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *CaptureDefaults) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *CaptureDefaults) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *CaptureDefaults) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *CaptureDefaults) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *CaptureDefaults) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil
### GetWaitFor

`func (o *CaptureDefaults) GetWaitFor() string`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *CaptureDefaults) GetWaitForOk() (*string, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *CaptureDefaults) SetWaitFor(v string)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *CaptureDefaults) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### SetWaitForNil

`func (o *CaptureDefaults) SetWaitForNil(b bool)`

 SetWaitForNil sets the value for WaitFor to be an explicit nil

### UnsetWaitFor
`func (o *CaptureDefaults) UnsetWaitFor()`

UnsetWaitFor ensures that no value is present for WaitFor, not even an explicit nil
### GetWaitUntil

`func (o *CaptureDefaults) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *CaptureDefaults) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *CaptureDefaults) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *CaptureDefaults) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


