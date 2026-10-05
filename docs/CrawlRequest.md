# CrawlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockAds** | Pointer to **bool** |  | [optional] [default to true]
**BlockCookieBanners** | Pointer to **bool** |  | [optional] [default to true]
**BlockPopups** | Pointer to **bool** |  | [optional] [default to true]
**CrawlDelayMs** | Pointer to **int32** |  | [optional] [default to 500]
**DarkMode** | Pointer to **bool** |  | [optional] [default to false]
**Depth** | Pointer to **int32** |  | [optional] [default to 2]
**ExcludePatterns** | Pointer to **[]string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "png"]
**FullPage** | Pointer to **bool** |  | [optional] [default to true]
**IncludePatterns** | Pointer to **[]string** |  | [optional] 
**IncludeSubdomains** | Pointer to **bool** |  | [optional] [default to false]
**Limit** | Pointer to **int32** |  | [optional] [default to 25]
**Outputs** | Pointer to [**[]OutputSpec**](OutputSpec.md) |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**RenderDelay** | Pointer to **int32** |  | [optional] [default to 0]
**Timeout** | Pointer to **int32** |  | [optional] [default to 30000]
**Url** | **string** |  | 
**Viewport** | Pointer to [**ViewportConfig**](ViewportConfig.md) |  | [optional] 
**WaitUntil** | Pointer to **string** |  | [optional] [default to "domcontentloaded"]

## Methods

### NewCrawlRequest

`func NewCrawlRequest(url string, ) *CrawlRequest`

NewCrawlRequest instantiates a new CrawlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlRequestWithDefaults

`func NewCrawlRequestWithDefaults() *CrawlRequest`

NewCrawlRequestWithDefaults instantiates a new CrawlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockAds

`func (o *CrawlRequest) GetBlockAds() bool`

GetBlockAds returns the BlockAds field if non-nil, zero value otherwise.

### GetBlockAdsOk

`func (o *CrawlRequest) GetBlockAdsOk() (*bool, bool)`

GetBlockAdsOk returns a tuple with the BlockAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAds

`func (o *CrawlRequest) SetBlockAds(v bool)`

SetBlockAds sets BlockAds field to given value.

### HasBlockAds

`func (o *CrawlRequest) HasBlockAds() bool`

HasBlockAds returns a boolean if a field has been set.

### GetBlockCookieBanners

`func (o *CrawlRequest) GetBlockCookieBanners() bool`

GetBlockCookieBanners returns the BlockCookieBanners field if non-nil, zero value otherwise.

### GetBlockCookieBannersOk

`func (o *CrawlRequest) GetBlockCookieBannersOk() (*bool, bool)`

GetBlockCookieBannersOk returns a tuple with the BlockCookieBanners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCookieBanners

`func (o *CrawlRequest) SetBlockCookieBanners(v bool)`

SetBlockCookieBanners sets BlockCookieBanners field to given value.

### HasBlockCookieBanners

`func (o *CrawlRequest) HasBlockCookieBanners() bool`

HasBlockCookieBanners returns a boolean if a field has been set.

### GetBlockPopups

`func (o *CrawlRequest) GetBlockPopups() bool`

GetBlockPopups returns the BlockPopups field if non-nil, zero value otherwise.

### GetBlockPopupsOk

`func (o *CrawlRequest) GetBlockPopupsOk() (*bool, bool)`

GetBlockPopupsOk returns a tuple with the BlockPopups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPopups

`func (o *CrawlRequest) SetBlockPopups(v bool)`

SetBlockPopups sets BlockPopups field to given value.

### HasBlockPopups

`func (o *CrawlRequest) HasBlockPopups() bool`

HasBlockPopups returns a boolean if a field has been set.

### GetCrawlDelayMs

`func (o *CrawlRequest) GetCrawlDelayMs() int32`

GetCrawlDelayMs returns the CrawlDelayMs field if non-nil, zero value otherwise.

### GetCrawlDelayMsOk

`func (o *CrawlRequest) GetCrawlDelayMsOk() (*int32, bool)`

GetCrawlDelayMsOk returns a tuple with the CrawlDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlDelayMs

`func (o *CrawlRequest) SetCrawlDelayMs(v int32)`

SetCrawlDelayMs sets CrawlDelayMs field to given value.

### HasCrawlDelayMs

`func (o *CrawlRequest) HasCrawlDelayMs() bool`

HasCrawlDelayMs returns a boolean if a field has been set.

### GetDarkMode

`func (o *CrawlRequest) GetDarkMode() bool`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *CrawlRequest) GetDarkModeOk() (*bool, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *CrawlRequest) SetDarkMode(v bool)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *CrawlRequest) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDepth

`func (o *CrawlRequest) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *CrawlRequest) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *CrawlRequest) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *CrawlRequest) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetExcludePatterns

`func (o *CrawlRequest) GetExcludePatterns() []string`

GetExcludePatterns returns the ExcludePatterns field if non-nil, zero value otherwise.

### GetExcludePatternsOk

`func (o *CrawlRequest) GetExcludePatternsOk() (*[]string, bool)`

GetExcludePatternsOk returns a tuple with the ExcludePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePatterns

`func (o *CrawlRequest) SetExcludePatterns(v []string)`

SetExcludePatterns sets ExcludePatterns field to given value.

### HasExcludePatterns

`func (o *CrawlRequest) HasExcludePatterns() bool`

HasExcludePatterns returns a boolean if a field has been set.

### GetFormat

`func (o *CrawlRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CrawlRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CrawlRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CrawlRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFullPage

`func (o *CrawlRequest) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *CrawlRequest) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *CrawlRequest) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *CrawlRequest) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetIncludePatterns

`func (o *CrawlRequest) GetIncludePatterns() []string`

GetIncludePatterns returns the IncludePatterns field if non-nil, zero value otherwise.

### GetIncludePatternsOk

`func (o *CrawlRequest) GetIncludePatternsOk() (*[]string, bool)`

GetIncludePatternsOk returns a tuple with the IncludePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePatterns

`func (o *CrawlRequest) SetIncludePatterns(v []string)`

SetIncludePatterns sets IncludePatterns field to given value.

### HasIncludePatterns

`func (o *CrawlRequest) HasIncludePatterns() bool`

HasIncludePatterns returns a boolean if a field has been set.

### GetIncludeSubdomains

`func (o *CrawlRequest) GetIncludeSubdomains() bool`

GetIncludeSubdomains returns the IncludeSubdomains field if non-nil, zero value otherwise.

### GetIncludeSubdomainsOk

`func (o *CrawlRequest) GetIncludeSubdomainsOk() (*bool, bool)`

GetIncludeSubdomainsOk returns a tuple with the IncludeSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubdomains

`func (o *CrawlRequest) SetIncludeSubdomains(v bool)`

SetIncludeSubdomains sets IncludeSubdomains field to given value.

### HasIncludeSubdomains

`func (o *CrawlRequest) HasIncludeSubdomains() bool`

HasIncludeSubdomains returns a boolean if a field has been set.

### GetLimit

`func (o *CrawlRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CrawlRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CrawlRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CrawlRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOutputs

`func (o *CrawlRequest) GetOutputs() []OutputSpec`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *CrawlRequest) GetOutputsOk() (*[]OutputSpec, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *CrawlRequest) SetOutputs(v []OutputSpec)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *CrawlRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetQuality

`func (o *CrawlRequest) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *CrawlRequest) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *CrawlRequest) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *CrawlRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetRenderDelay

`func (o *CrawlRequest) GetRenderDelay() int32`

GetRenderDelay returns the RenderDelay field if non-nil, zero value otherwise.

### GetRenderDelayOk

`func (o *CrawlRequest) GetRenderDelayOk() (*int32, bool)`

GetRenderDelayOk returns a tuple with the RenderDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderDelay

`func (o *CrawlRequest) SetRenderDelay(v int32)`

SetRenderDelay sets RenderDelay field to given value.

### HasRenderDelay

`func (o *CrawlRequest) HasRenderDelay() bool`

HasRenderDelay returns a boolean if a field has been set.

### GetTimeout

`func (o *CrawlRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CrawlRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CrawlRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CrawlRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *CrawlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CrawlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CrawlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetViewport

`func (o *CrawlRequest) GetViewport() ViewportConfig`

GetViewport returns the Viewport field if non-nil, zero value otherwise.

### GetViewportOk

`func (o *CrawlRequest) GetViewportOk() (*ViewportConfig, bool)`

GetViewportOk returns a tuple with the Viewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewport

`func (o *CrawlRequest) SetViewport(v ViewportConfig)`

SetViewport sets Viewport field to given value.

### HasViewport

`func (o *CrawlRequest) HasViewport() bool`

HasViewport returns a boolean if a field has been set.

### GetWaitUntil

`func (o *CrawlRequest) GetWaitUntil() string`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *CrawlRequest) GetWaitUntilOk() (*string, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *CrawlRequest) SetWaitUntil(v string)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *CrawlRequest) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


