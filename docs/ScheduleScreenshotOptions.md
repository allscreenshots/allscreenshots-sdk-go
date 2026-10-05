# ScheduleScreenshotOptions

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
**FreezeFixed** | Pointer to **bool** |  | [optional] [default to true]
**FullPage** | Pointer to **bool** |  | [optional] [default to false]
**FullPageMode** | Pointer to **string** |  | [optional] [default to "stitch"]
**HideSelectors** | Pointer to **[]string** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxSections** | Pointer to **int32** |  | [optional] [default to 50]
**Outputs** | Pointer to [**[]OutputSpec**](OutputSpec.md) |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**ScrollInterval** | Pointer to **int32** |  | [optional] [default to 150]
**Selector** | Pointer to **NullableString** |  | [optional] 
**StealthMode** | Pointer to **bool** |  | [optional] [default to false]
**Timeout** | Pointer to **int32** |  | [optional] [default to 30000]
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitFor** | Pointer to **NullableString** |  | [optional] 
**WaitUntil** | Pointer to **string** |  | [optional] [default to "domcontentloaded"]

## Methods

### NewScheduleScreenshotOptions

`func NewScheduleScreenshotOptions() *ScheduleScreenshotOptions`

NewScheduleScreenshotOptions instantiates a new ScheduleScreenshotOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleScreenshotOptionsWithDefaults

`func NewScheduleScreenshotOptionsWithDefaults() *ScheduleScreenshotOptions`

NewScheduleScreenshotOptionsWithDefaults instantiates a new ScheduleScreenshotOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ScheduleScreenshotOptions) GetActions() []PageAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScheduleScreenshotOptions) GetActionsOk() (*[]PageAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScheduleScreenshotOptions) SetActions(v []PageAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ScheduleScreenshotOptions) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *ScheduleScreenshotOptions) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *ScheduleScreenshotOptions) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetBlockAds

`func (o *ScheduleScreenshotOptions) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *ScheduleScreenshotOptions) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *ScheduleScreenshotOptions) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *ScheduleScreenshotOptions) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetBlockCookieBanners

`func (o *ScheduleScreenshotOptions) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *ScheduleScreenshotOptions) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *ScheduleScreenshotOptions) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *ScheduleScreenshotOptions) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### GetBlockLevel

`func (o *ScheduleScreenshotOptions) GetBlockLevel() string`

GetBlockLevel returns the BlockLevel field if non-nil, zero value otherwise.

### GetBlockLevelOk

`func (o *ScheduleScreenshotOptions) GetBlockLevelOk() (*string, bool)`

GetBlockLevelOk returns a tuple with the BlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevel

`func (o *ScheduleScreenshotOptions) SetBlockLevel(v string)`

SetBlockLevel sets BlockLevel field to given value.

### HasBlockLevel

`func (o *ScheduleScreenshotOptions) HasBlockLevel() bool`

HasBlockLevel returns a boolean if a field has been set.

### GetBlockPopups

`func (o *ScheduleScreenshotOptions) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *ScheduleScreenshotOptions) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *ScheduleScreenshotOptions) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *ScheduleScreenshotOptions) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### GetCustomCss

`func (o *ScheduleScreenshotOptions) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *ScheduleScreenshotOptions) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *ScheduleScreenshotOptions) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *ScheduleScreenshotOptions) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *ScheduleScreenshotOptions) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *ScheduleScreenshotOptions) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *ScheduleScreenshotOptions) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *ScheduleScreenshotOptions) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *ScheduleScreenshotOptions) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *ScheduleScreenshotOptions) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDelay

`func (o *ScheduleScreenshotOptions) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *ScheduleScreenshotOptions) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *ScheduleScreenshotOptions) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *ScheduleScreenshotOptions) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDevice

`func (o *ScheduleScreenshotOptions) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ScheduleScreenshotOptions) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ScheduleScreenshotOptions) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ScheduleScreenshotOptions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ScheduleScreenshotOptions) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ScheduleScreenshotOptions) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFormat

`func (o *ScheduleScreenshotOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScheduleScreenshotOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScheduleScreenshotOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ScheduleScreenshotOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFreezeFixed

`func (o *ScheduleScreenshotOptions) GetFreezeFixed() bool`

GetFreezeFixed returns the FreezeFixed field if non-nil, zero value otherwise.

### GetFreezeFixedOk

`func (o *ScheduleScreenshotOptions) GetFreezeFixedOk() (*bool, bool)`

GetFreezeFixedOk returns a tuple with the FreezeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeFixed

`func (o *ScheduleScreenshotOptions) SetFreezeFixed(v bool)`

SetFreezeFixed sets FreezeFixed field to given value.

### HasFreezeFixed

`func (o *ScheduleScreenshotOptions) HasFreezeFixed() bool`

HasFreezeFixed returns a boolean if a field has been set.

### GetFullPage

`func (o *ScheduleScreenshotOptions) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *ScheduleScreenshotOptions) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *ScheduleScreenshotOptions) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *ScheduleScreenshotOptions) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetFullPageMode

`func (o *ScheduleScreenshotOptions) GetFullPageMode() string`

GetFullPageMode returns the FullPageMode field if non-nil, zero value otherwise.

### GetFullPageModeOk

`func (o *ScheduleScreenshotOptions) GetFullPageModeOk() (*string, bool)`

GetFullPageModeOk returns a tuple with the FullPageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageMode

`func (o *ScheduleScreenshotOptions) SetFullPageMode(v string)`

SetFullPageMode sets FullPageMode field to given value.

### HasFullPageMode

`func (o *ScheduleScreenshotOptions) HasFullPageMode() bool`

HasFullPageMode returns a boolean if a field has been set.

### GetHideSelectors

`func (o *ScheduleScreenshotOptions) GetHideSelectors() []string`

GetHideSelectors returns the HideSelectors field if non-nil, zero value otherwise.

### GetHideSelectorsOk

`func (o *ScheduleScreenshotOptions) GetHideSelectorsOk() (*[]string, bool)`

GetHideSelectorsOk returns a tuple with the HideSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSelectors

`func (o *ScheduleScreenshotOptions) SetHideSelectors(v []string)`

SetHideSelectors sets HideSelectors field to given value.

### HasHideSelectors

`func (o *ScheduleScreenshotOptions) HasHideSelectors() bool`

HasHideSelectors returns a boolean if a field has been set.

### SetHideSelectorsNil

`func (o *ScheduleScreenshotOptions) SetHideSelectorsNil(b bool)`

 SetHideSelectorsNil sets the value for HideSelectors to be an explicit nil

### UnsetHideSelectors
`func (o *ScheduleScreenshotOptions) UnsetHideSelectors()`

UnsetHideSelectors ensures that no value is present for HideSelectors, not even an explicit nil
### GetMaxHeight

`func (o *ScheduleScreenshotOptions) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ScheduleScreenshotOptions) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ScheduleScreenshotOptions) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ScheduleScreenshotOptions) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ScheduleScreenshotOptions) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ScheduleScreenshotOptions) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetMaxSections

`func (o *ScheduleScreenshotOptions) GetMaxSections() int32`

GetMaxSections returns the MaxSections field if non-nil, zero value otherwise.

### GetMaxSectionsOk

`func (o *ScheduleScreenshotOptions) GetMaxSectionsOk() (*int32, bool)`

GetMaxSectionsOk returns a tuple with the MaxSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSections

`func (o *ScheduleScreenshotOptions) SetMaxSections(v int32)`

SetMaxSections sets MaxSections field to given value.

### HasMaxSections

`func (o *ScheduleScreenshotOptions) HasMaxSections() bool`

HasMaxSections returns a boolean if a field has been set.

### GetOutputs

`func (o *ScheduleScreenshotOptions) GetOutputs() []OutputSpec`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ScheduleScreenshotOptions) GetOutputsOk() (*[]OutputSpec, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ScheduleScreenshotOptions) SetOutputs(v []OutputSpec)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ScheduleScreenshotOptions) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *ScheduleScreenshotOptions) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *ScheduleScreenshotOptions) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetQuality

`func (o *ScheduleScreenshotOptions) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ScheduleScreenshotOptions) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ScheduleScreenshotOptions) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ScheduleScreenshotOptions) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetScrollInterval

`func (o *ScheduleScreenshotOptions) GetScrollInterval() int32`

GetScrollInterval returns the ScrollInterval field if non-nil, zero value otherwise.

### GetScrollIntervalOk

`func (o *ScheduleScreenshotOptions) GetScrollIntervalOk() (*int32, bool)`

GetScrollIntervalOk returns a tuple with the ScrollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollInterval

`func (o *ScheduleScreenshotOptions) SetScrollInterval(v int32)`

SetScrollInterval sets ScrollInterval field to given value.

### HasScrollInterval

`func (o *ScheduleScreenshotOptions) HasScrollInterval() bool`

HasScrollInterval returns a boolean if a field has been set.

### GetSelector

`func (o *ScheduleScreenshotOptions) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ScheduleScreenshotOptions) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ScheduleScreenshotOptions) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ScheduleScreenshotOptions) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *ScheduleScreenshotOptions) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *ScheduleScreenshotOptions) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetStealthMode

`func (o *ScheduleScreenshotOptions) GetStealthMode() bool`

GetStealthMode returns the StealthMode field if non-nil, zero value otherwise.

### GetStealthModeOk

`func (o *ScheduleScreenshotOptions) GetStealthModeOk() (*bool, bool)`

GetStealthModeOk returns a tuple with the StealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthMode

`func (o *ScheduleScreenshotOptions) SetStealthMode(v bool)`

SetStealthMode sets StealthMode field to given value.

### HasStealthMode

`func (o *ScheduleScreenshotOptions) HasStealthMode() bool`

HasStealthMode returns a boolean if a field has been set.

### GetTimeout

`func (o *ScheduleScreenshotOptions) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ScheduleScreenshotOptions) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ScheduleScreenshotOptions) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ScheduleScreenshotOptions) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetViewport

`func (o *ScheduleScreenshotOptions) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *ScheduleScreenshotOptions) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *ScheduleScreenshotOptions) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *ScheduleScreenshotOptions) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *ScheduleScreenshotOptions) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *ScheduleScreenshotOptions) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil
### GetWaitFor

`func (o *ScheduleScreenshotOptions) GetWaitFor() string`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *ScheduleScreenshotOptions) GetWaitForOk() (*string, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *ScheduleScreenshotOptions) SetWaitFor(v string)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *ScheduleScreenshotOptions) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### SetWaitForNil

`func (o *ScheduleScreenshotOptions) SetWaitForNil(b bool)`

 SetWaitForNil sets the value for WaitFor to be an explicit nil

### UnsetWaitFor
`func (o *ScheduleScreenshotOptions) UnsetWaitFor()`

UnsetWaitFor ensures that no value is present for WaitFor, not even an explicit nil
### GetWaitUntil

`func (o *ScheduleScreenshotOptions) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *ScheduleScreenshotOptions) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *ScheduleScreenshotOptions) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *ScheduleScreenshotOptions) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


