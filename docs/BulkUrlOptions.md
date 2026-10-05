# BulkUrlOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]PageAction**](PageAction.md) |  | [optional] 
**BlockAds** | Pointer to **NullableBool** |  | [optional] 
**BlockCookieBanners** | Pointer to **NullableBool** |  | [optional] 
**BlockLevel** | Pointer to **NullableString** |  | [optional] 
**BlockPopups** | Pointer to **NullableBool** |  | [optional] 
**CustomCss** | Pointer to **NullableString** |  | [optional] 
**DarkMode** | Pointer to **NullableBool** |  | [optional] 
**Delay** | Pointer to **NullableInt32** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**FullPage** | Pointer to **NullableBool** |  | [optional] 
**Outputs** | Pointer to [**[]OutputSpec**](OutputSpec.md) |  | [optional] 
**Quality** | Pointer to **NullableInt32** |  | [optional] 
**StealthMode** | Pointer to **NullableBool** |  | [optional] 
**Viewport** | Pointer to [**NullableViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitFor** | Pointer to **NullableString** |  | [optional] 
**WaitUntil** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBulkUrlOptions

`func NewBulkUrlOptions() *BulkUrlOptions`

NewBulkUrlOptions instantiates a new BulkUrlOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUrlOptionsWithDefaults

`func NewBulkUrlOptionsWithDefaults() *BulkUrlOptions`

NewBulkUrlOptionsWithDefaults instantiates a new BulkUrlOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *BulkUrlOptions) GetActions() []PageAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *BulkUrlOptions) GetActionsOk() (*[]PageAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *BulkUrlOptions) SetActions(v []PageAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *BulkUrlOptions) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *BulkUrlOptions) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *BulkUrlOptions) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetBlockAds

`func (o *BulkUrlOptions) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *BulkUrlOptions) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *BulkUrlOptions) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *BulkUrlOptions) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### SetBlockAdsNil

`func (o *BulkUrlOptions) SetBlockAdsNil(b bool)`

 SetBlockAdsNil sets the value for BlockAds to be an explicit nil

### UnsetBlockAds
`func (o *BulkUrlOptions) UnsetBlockAds()`

UnsetBlockAds ensures that no value is present for BlockAds, not even an explicit nil
### GetBlockCookieBanners

`func (o *BulkUrlOptions) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *BulkUrlOptions) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *BulkUrlOptions) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *BulkUrlOptions) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### SetBlockCookieBannersNil

`func (o *BulkUrlOptions) SetBlockCookieBannersNil(b bool)`

 SetBlockCookieBannersNil sets the value for BlockCookieBanners to be an explicit nil

### UnsetBlockCookieBanners
`func (o *BulkUrlOptions) UnsetBlockCookieBanners()`

UnsetBlockCookieBanners ensures that no value is present for BlockCookieBanners, not even an explicit nil
### GetBlockLevel

`func (o *BulkUrlOptions) GetBlockLevel() string`

GetBlockLevel returns the BlockLevel field if non-nil, zero value otherwise.

### GetBlockLevelOk

`func (o *BulkUrlOptions) GetBlockLevelOk() (*string, bool)`

GetBlockLevelOk returns a tuple with the BlockLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLevel

`func (o *BulkUrlOptions) SetBlockLevel(v string)`

SetBlockLevel sets BlockLevel field to given value.

### HasBlockLevel

`func (o *BulkUrlOptions) HasBlockLevel() bool`

HasBlockLevel returns a boolean if a field has been set.

### SetBlockLevelNil

`func (o *BulkUrlOptions) SetBlockLevelNil(b bool)`

 SetBlockLevelNil sets the value for BlockLevel to be an explicit nil

### UnsetBlockLevel
`func (o *BulkUrlOptions) UnsetBlockLevel()`

UnsetBlockLevel ensures that no value is present for BlockLevel, not even an explicit nil
### GetBlockPopups

`func (o *BulkUrlOptions) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *BulkUrlOptions) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *BulkUrlOptions) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *BulkUrlOptions) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### SetBlockPopupsNil

`func (o *BulkUrlOptions) SetBlockPopupsNil(b bool)`

 SetBlockPopupsNil sets the value for BlockPopups to be an explicit nil

### UnsetBlockPopups
`func (o *BulkUrlOptions) UnsetBlockPopups()`

UnsetBlockPopups ensures that no value is present for BlockPopups, not even an explicit nil
### GetCustomCss

`func (o *BulkUrlOptions) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *BulkUrlOptions) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *BulkUrlOptions) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *BulkUrlOptions) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *BulkUrlOptions) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *BulkUrlOptions) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetDarkMode

`func (o *BulkUrlOptions) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *BulkUrlOptions) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *BulkUrlOptions) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *BulkUrlOptions) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### SetDarkModeNil

`func (o *BulkUrlOptions) SetDarkModeNil(b bool)`

 SetDarkModeNil sets the value for DarkMode to be an explicit nil

### UnsetDarkMode
`func (o *BulkUrlOptions) UnsetDarkMode()`

UnsetDarkMode ensures that no value is present for DarkMode, not even an explicit nil
### GetDelay

`func (o *BulkUrlOptions) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *BulkUrlOptions) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *BulkUrlOptions) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *BulkUrlOptions) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *BulkUrlOptions) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *BulkUrlOptions) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetDevice

`func (o *BulkUrlOptions) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BulkUrlOptions) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BulkUrlOptions) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BulkUrlOptions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *BulkUrlOptions) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *BulkUrlOptions) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetFormat

`func (o *BulkUrlOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BulkUrlOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BulkUrlOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BulkUrlOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *BulkUrlOptions) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *BulkUrlOptions) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetFullPage

`func (o *BulkUrlOptions) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *BulkUrlOptions) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *BulkUrlOptions) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *BulkUrlOptions) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### SetFullPageNil

`func (o *BulkUrlOptions) SetFullPageNil(b bool)`

 SetFullPageNil sets the value for FullPage to be an explicit nil

### UnsetFullPage
`func (o *BulkUrlOptions) UnsetFullPage()`

UnsetFullPage ensures that no value is present for FullPage, not even an explicit nil
### GetOutputs

`func (o *BulkUrlOptions) GetOutputs() []OutputSpec`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BulkUrlOptions) GetOutputsOk() (*[]OutputSpec, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BulkUrlOptions) SetOutputs(v []OutputSpec)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BulkUrlOptions) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *BulkUrlOptions) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *BulkUrlOptions) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetQuality

`func (o *BulkUrlOptions) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BulkUrlOptions) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BulkUrlOptions) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BulkUrlOptions) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *BulkUrlOptions) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *BulkUrlOptions) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetStealthMode

`func (o *BulkUrlOptions) GetStealthMode() bool`

GetStealthMode returns the StealthMode field if non-nil, zero value otherwise.

### GetStealthModeOk

`func (o *BulkUrlOptions) GetStealthModeOk() (*bool, bool)`

GetStealthModeOk returns a tuple with the StealthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealthMode

`func (o *BulkUrlOptions) SetStealthMode(v bool)`

SetStealthMode sets StealthMode field to given value.

### HasStealthMode

`func (o *BulkUrlOptions) HasStealthMode() bool`

HasStealthMode returns a boolean if a field has been set.

### SetStealthModeNil

`func (o *BulkUrlOptions) SetStealthModeNil(b bool)`

 SetStealthModeNil sets the value for StealthMode to be an explicit nil

### UnsetStealthMode
`func (o *BulkUrlOptions) UnsetStealthMode()`

UnsetStealthMode ensures that no value is present for StealthMode, not even an explicit nil
### GetViewport

`func (o *BulkUrlOptions) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *BulkUrlOptions) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *BulkUrlOptions) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *BulkUrlOptions) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### SetViewportNil

`func (o *BulkUrlOptions) SetViewportNil(b bool)`

 SetViewportNil sets the value for Viewport to be an explicit nil

### UnsetViewport
`func (o *BulkUrlOptions) UnsetViewport()`

UnsetViewport ensures that no value is present for Viewport, not even an explicit nil
### GetWaitFor

`func (o *BulkUrlOptions) GetWaitFor() string`

GetWaitFor returns the WaitFor field if non-nil, zero value otherwise.

### GetWaitForOk

`func (o *BulkUrlOptions) GetWaitForOk() (*string, bool)`

GetWaitForOk returns a tuple with the WaitFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitFor

`func (o *BulkUrlOptions) SetWaitFor(v string)`

SetWaitFor sets WaitFor field to given value.

### HasWaitFor

`func (o *BulkUrlOptions) HasWaitFor() bool`

HasWaitFor returns a boolean if a field has been set.

### SetWaitForNil

`func (o *BulkUrlOptions) SetWaitForNil(b bool)`

 SetWaitForNil sets the value for WaitFor to be an explicit nil

### UnsetWaitFor
`func (o *BulkUrlOptions) UnsetWaitFor()`

UnsetWaitFor ensures that no value is present for WaitFor, not even an explicit nil
### GetWaitUntil

`func (o *BulkUrlOptions) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *BulkUrlOptions) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *BulkUrlOptions) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *BulkUrlOptions) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.

### SetWaitUntilNil

`func (o *BulkUrlOptions) SetWaitUntilNil(b bool)`

 SetWaitUntilNil sets the value for WaitUntil to be an explicit nil

### UnsetWaitUntil
`func (o *BulkUrlOptions) UnsetWaitUntil()`

UnsetWaitUntil ensures that no value is present for WaitUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


