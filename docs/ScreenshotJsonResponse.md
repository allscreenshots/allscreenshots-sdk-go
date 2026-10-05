# ScreenshotJsonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cached** | Pointer to **bool** |  | [optional] [default to false]
**ContentType** | **string** |  | 
**Data** | Pointer to **NullableString** |  | [optional] 
**Encoding** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**Format** | **string** |  | 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**RenderTimeMs** | **int64** |  | 
**ResultUrl** | Pointer to **NullableString** |  | [optional] 
**Size** | **int32** |  | 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Width** | **int32** |  | 

## Methods

### NewScreenshotJsonResponse

`func NewScreenshotJsonResponse(contentType string, format string, renderTimeMs int64, size int32, url string, width int32, ) *ScreenshotJsonResponse`

NewScreenshotJsonResponse instantiates a new ScreenshotJsonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotJsonResponseWithDefaults

`func NewScreenshotJsonResponseWithDefaults() *ScreenshotJsonResponse`

NewScreenshotJsonResponseWithDefaults instantiates a new ScreenshotJsonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCached

`func (o *ScreenshotJsonResponse) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *ScreenshotJsonResponse) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *ScreenshotJsonResponse) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *ScreenshotJsonResponse) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetContentType

`func (o *ScreenshotJsonResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ScreenshotJsonResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ScreenshotJsonResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetData

`func (o *ScreenshotJsonResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScreenshotJsonResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScreenshotJsonResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ScreenshotJsonResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ScreenshotJsonResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ScreenshotJsonResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEncoding

`func (o *ScreenshotJsonResponse) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ScreenshotJsonResponse) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ScreenshotJsonResponse) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ScreenshotJsonResponse) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### SetEncodingNil

`func (o *ScreenshotJsonResponse) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *ScreenshotJsonResponse) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetExpiresAt

`func (o *ScreenshotJsonResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ScreenshotJsonResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ScreenshotJsonResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ScreenshotJsonResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ScreenshotJsonResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ScreenshotJsonResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFormat

`func (o *ScreenshotJsonResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScreenshotJsonResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScreenshotJsonResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetHeight

`func (o *ScreenshotJsonResponse) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ScreenshotJsonResponse) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ScreenshotJsonResponse) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ScreenshotJsonResponse) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ScreenshotJsonResponse) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ScreenshotJsonResponse) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetRenderTimeMs

`func (o *ScreenshotJsonResponse) GetRenderTimeMs() int64`

GetRenderTimeMs returns the RenderTimeMs field if non-nil, zero value otherwise.

### GetRenderTimeMsOk

`func (o *ScreenshotJsonResponse) GetRenderTimeMsOk() (*int64, bool)`

GetRenderTimeMsOk returns a tuple with the RenderTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTimeMs

`func (o *ScreenshotJsonResponse) SetRenderTimeMs(v int64)`

SetRenderTimeMs sets RenderTimeMs field to given value.


### GetResultUrl

`func (o *ScreenshotJsonResponse) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *ScreenshotJsonResponse) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *ScreenshotJsonResponse) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.

### HasResultUrl

`func (o *ScreenshotJsonResponse) HasResultUrl() bool`

HasResultUrl returns a boolean if a field has been set.

### SetResultUrlNil

`func (o *ScreenshotJsonResponse) SetResultUrlNil(b bool)`

 SetResultUrlNil sets the value for ResultUrl to be an explicit nil

### UnsetResultUrl
`func (o *ScreenshotJsonResponse) UnsetResultUrl()`

UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
### GetSize

`func (o *ScreenshotJsonResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ScreenshotJsonResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ScreenshotJsonResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStorageUrl

`func (o *ScreenshotJsonResponse) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *ScreenshotJsonResponse) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *ScreenshotJsonResponse) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *ScreenshotJsonResponse) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *ScreenshotJsonResponse) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *ScreenshotJsonResponse) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetUrl

`func (o *ScreenshotJsonResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScreenshotJsonResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScreenshotJsonResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWidth

`func (o *ScreenshotJsonResponse) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ScreenshotJsonResponse) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ScreenshotJsonResponse) SetWidth(v int32)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


