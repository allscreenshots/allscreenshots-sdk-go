# ScreenshotRequest

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
**ResponseType** | Pointer to **string** |  | [optional] [default to "BINARY"]
**ScrollInterval** | Pointer to **int32** |  | [optional] [default to 150]
**Selector** | Pointer to **NullableString** |  | [optional] 
**Session** | Pointer to [**NullableBrowserSessionRequest**](BrowserSessionRequest.md) |  | [optional] 
**StealthMode** | Pointer to **bool** |  | [optional] [default to false]
**Timeout** | Pointer to **int32** |  | [optional] [default to 30000]
**Url** | **string** |  | 
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitFor** | Pointer to **NullableString** |  | [optional] 
**WaitUntil** | Pointer to **string** |  | [optional] [default to "domcontentloaded"]
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScreenshotRequest

`func NewScreenshotRequest(url string, ) *ScreenshotRequest`

NewScreenshotRequest instantiates a new ScreenshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotRequestWithDefaults

`func NewScreenshotRequestWithDefaults() *ScreenshotRequest`

NewScreenshotRequestWithDefaults instantiates a new ScreenshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ScreenshotRequest) GetActions() []PageAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ScreenshotRequest) GetActionsOk() (*[]PageAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ScreenshotRequest) SetActions(v []PageAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ScreenshotRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *ScreenshotRequest) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *ScreenshotRequest) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetBlockAds

`func (o *ScreenshotRequest) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *ScreenshotRequest) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *ScreenshotRequest) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *ScreenshotRequest) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetBlockCookieBanners

`func (o *ScreenshotRequest) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *ScreenshotRequest) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *ScreenshotRequest) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *ScreenshotRequest) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### GetBlockLevel

`func (o *ScreenshotRequest) GetBlockLevel() string`

GetBlockLevel returns the BlockLevel field if non-nil, zero value otherwise.

### GetBlockLevelOk

`func (o *ScreenshotRequest) GetBlockLevelOk() (*string, bool)`

GetBlockLevelOk returns a tuple with the BlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevel

`func (o *ScreenshotRequest) SetBlockLevel(v string)`

SetBlockLevel sets BlockLevel field to given value.

### HasBlockLevel

`func (o *ScreenshotRequest) HasBlockLevel() bool`

HasBlockLevel returns a boolean if a field has been set.

### GetBlockPopups

`func (o *ScreenshotRequest) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *ScreenshotRequest) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *ScreenshotRequest) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *ScreenshotRequest) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### GetCustomCss

`func (o *ScreenshotRequest) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *ScreenshotRequest) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *ScreenshotRequest) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *ScreenshotRequest) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *ScreenshotRequest) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *ScreenshotRequest) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *ScreenshotRequest) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *ScreenshotRequest) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *ScreenshotRequest) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *ScreenshotRequest) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDelay

`func (o *ScreenshotRequest) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *ScreenshotRequest) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *ScreenshotRequest) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *ScreenshotRequest) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDevice

`func (o *ScreenshotRequest) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ScreenshotRequest) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ScreenshotRequest) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ScreenshotRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ScreenshotRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ScreenshotRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFormat

`func (o *ScreenshotRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScreenshotRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScreenshotRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ScreenshotRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFreezeFixed

`func (o *ScreenshotRequest) GetFreezeFixed() bool`

GetFreezeFixed returns the FreezeFixed field if non-nil, zero value otherwise.

### GetFreezeFixedOk

`func (o *ScreenshotRequest) GetFreezeFixedOk() (*bool, bool)`

GetFreezeFixedOk returns a tuple with the FreezeFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeFixed

`func (o *ScreenshotRequest) SetFreezeFixed(v bool)`

SetFreezeFixed sets FreezeFixed field to given value.

### HasFreezeFixed

`func (o *ScreenshotRequest) HasFreezeFixed() bool`

HasFreezeFixed returns a boolean if a field has been set.

### GetFullPage

`func (o *ScreenshotRequest) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *ScreenshotRequest) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *ScreenshotRequest) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *ScreenshotRequest) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetFullPageMode

`func (o *ScreenshotRequest) GetFullPageMode() string`

GetFullPageMode returns the FullPageMode field if non-nil, zero value otherwise.

### GetFullPageModeOk

`func (o *ScreenshotRequest) GetFullPageModeOk() (*string, bool)`

GetFullPageModeOk returns a tuple with the FullPageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageMode

`func (o *ScreenshotRequest) SetFullPageMode(v string)`

SetFullPageMode sets FullPageMode field to given value.

### HasFullPageMode

`func (o *ScreenshotRequest) HasFullPageMode() bool`

HasFullPageMode returns a boolean if a field has been set.

### GetHideSelectors

`func (o *ScreenshotRequest) GetHideSelectors() []string`

GetHideSelectors returns the HideSelectors field if non-nil, zero value otherwise.

### GetHideSelectorsOk

`func (o *ScreenshotRequest) GetHideSelectorsOk() (*[]string, bool)`

GetHideSelectorsOk returns a tuple with the HideSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSelectors

`func (o *ScreenshotRequest) SetHideSelectors(v []string)`

SetHideSelectors sets HideSelectors field to given value.

### HasHideSelectors

`func (o *ScreenshotRequest) HasHideSelectors() bool`

HasHideSelectors returns a boolean if a field has been set.

### SetHideSelectorsNil

`func (o *ScreenshotRequest) SetHideSelectorsNil(b bool)`

 SetHideSelectorsNil sets the value for HideSelectors to be an explicit nil

### UnsetHideSelectors
`func (o *ScreenshotRequest) UnsetHideSelectors()`

UnsetHideSelectors ensures that no value is present for HideSelectors, not even an explicit nil
### GetMaxHeight

`func (o *ScreenshotRequest) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ScreenshotRequest) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ScreenshotRequest) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ScreenshotRequest) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ScreenshotRequest) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ScreenshotRequest) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetMaxSections

`func (o *ScreenshotRequest) GetMaxSections() int32`

GetMaxSections returns the MaxSections field if non-nil, zero value otherwise.

### GetMaxSectionsOk

`func (o *ScreenshotRequest) GetMaxSectionsOk() (*int32, bool)`

GetMaxSectionsOk returns a tuple with the MaxSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSections

`func (o *ScreenshotRequest) SetMaxSections(v int32)`

SetMaxSections sets MaxSections field to given value.

### HasMaxSections

`func (o *ScreenshotRequest) HasMaxSections() bool`

HasMaxSections returns a boolean if a field has been set.

### GetOutputs

`func (o *ScreenshotRequest) GetOutputs() []OutputSpec`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ScreenshotRequest) GetOutputsOk() (*[]OutputSpec, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ScreenshotRequest) SetOutputs(v []OutputSpec)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ScreenshotRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *ScreenshotRequest) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *ScreenshotRequest) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetQuality

`func (o *ScreenshotRequest) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ScreenshotRequest) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ScreenshotRequest) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ScreenshotRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResponseType

`func (o *ScreenshotRequest) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *ScreenshotRequest) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *ScreenshotRequest) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *ScreenshotRequest) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetScrollInterval

`func (o *ScreenshotRequest) GetScrollInterval() int32`

GetScrollInterval returns the ScrollInterval field if non-nil, zero value otherwise.

### GetScrollIntervalOk

`func (o *ScreenshotRequest) GetScrollIntervalOk() (*int32, bool)`

GetScrollIntervalOk returns a tuple with the ScrollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollInterval

`func (o *ScreenshotRequest) SetScrollInterval(v int32)`

SetScrollInterval sets ScrollInterval field to given value.

### HasScrollInterval

`func (o *ScreenshotRequest) HasScrollInterval() bool`

HasScrollInterval returns a boolean if a field has been set.

### GetSelector

`func (o *ScreenshotRequest) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ScreenshotRequest) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ScreenshotRequest) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ScreenshotRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *ScreenshotRequest) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *ScreenshotRequest) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetSession

`func (o *ScreenshotRequest) GetSession() BrowserSessionRequest`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *ScreenshotRequest) GetSessionOk() (*BrowserSessionRequest, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *ScreenshotRequest) SetSession(v BrowserSessionRequest)`

SetSession sets Session field to given value.

### HasSession

`func (o *ScreenshotRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.

### SetSessionNil

`func (o *ScreenshotRequest) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *ScreenshotRequest) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetStealthMode

`func (o *ScreenshotRequest) GetStealthMode() bool`

GetStealthMode returns the StealthMode field if non-nil, zero value otherwise.

### GetStealthModeOk

`func (o *ScreenshotRequest) GetStealthModeOk() (*bool, bool)`

GetStealthModeOk returns a tuple with the StealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthMode

`func (o *ScreenshotRequest) SetStealthMode(v bool)`

SetStealthMode sets StealthMode field to given value.

### HasStealthMode

`func (o *ScreenshotRequest) HasStealthMode() bool`

HasStealthMode returns a boolean if a field has been set.

### GetTimeout

`func (o *ScreenshotRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ScreenshotRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ScreenshotRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ScreenshotRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *ScreenshotRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScreenshotRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScreenshotRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetViewport

`func (o *ScreenshotRequest) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *ScreenshotRequest) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *ScreenshotRequest) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *ScreenshotRequest) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *ScreenshotRequest) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *ScreenshotRequest) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil
### GetWaitFor

`func (o *ScreenshotRequest) GetWaitFor() string`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *ScreenshotRequest) GetWaitForOk() (*string, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *ScreenshotRequest) SetWaitFor(v string)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *ScreenshotRequest) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### SetWaitForNil

`func (o *ScreenshotRequest) SetWaitForNil(b bool)`

 SetWaitForNil sets the value for WaitFor to be an explicit nil

### UnsetWaitFor
`func (o *ScreenshotRequest) UnsetWaitFor()`

UnsetWaitFor ensures that no value is present for WaitFor, not even an explicit nil
### GetWaitUntil

`func (o *ScreenshotRequest) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *ScreenshotRequest) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *ScreenshotRequest) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *ScreenshotRequest) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *ScreenshotRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *ScreenshotRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *ScreenshotRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *ScreenshotRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *ScreenshotRequest) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *ScreenshotRequest) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetWebhookUrl

`func (o *ScreenshotRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ScreenshotRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ScreenshotRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ScreenshotRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *ScreenshotRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *ScreenshotRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


