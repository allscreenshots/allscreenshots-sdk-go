# ComposeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** |  | 
**FileSize** | **int64** |  | 
**Format** | **string** |  | 
**Height** | **int32** |  | 
**Layout** | **string** |  | 
**Metadata** | [**ComposeMetadata**](ComposeMetadata.md) |  | 
**RenderTimeMs** | **int64** |  | 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Width** | **int32** |  | 

## Methods

### NewComposeResponse

`func NewComposeResponse(expiresAt time.Time, fileSize int64, format string, height int32, layout string, metadata ComposeMetadata, renderTimeMs int64, url string, width int32, ) *ComposeResponse`

NewComposeResponse instantiates a new ComposeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeResponseWithDefaults

`func NewComposeResponseWithDefaults() *ComposeResponse`

NewComposeResponseWithDefaults instantiates a new ComposeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ComposeResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ComposeResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ComposeResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFileSize

`func (o *ComposeResponse) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ComposeResponse) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ComposeResponse) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetFormat

`func (o *ComposeResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ComposeResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ComposeResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetHeight

`func (o *ComposeResponse) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ComposeResponse) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ComposeResponse) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetLayout

`func (o *ComposeResponse) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ComposeResponse) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ComposeResponse) SetLayout(v string)`

SetLayout sets Layout field to given value.


### GetMetadata

`func (o *ComposeResponse) GetMetadata() ComposeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComposeResponse) GetMetadataOk() (*ComposeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComposeResponse) SetMetadata(v ComposeMetadata)`

SetMetadata sets Metadata field to given value.


### GetRenderTimeMs

`func (o *ComposeResponse) GetRenderTimeMs() int64`

GetRenderTimeMs returns the RenderTimeMs field if non-nil, zero value otherwise.

### GetRenderTimeMsOk

`func (o *ComposeResponse) GetRenderTimeMsOk() (*int64, bool)`

GetRenderTimeMsOk returns a tuple with the RenderTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTimeMs

`func (o *ComposeResponse) SetRenderTimeMs(v int64)`

SetRenderTimeMs sets RenderTimeMs field to given value.


### GetStorageUrl

`func (o *ComposeResponse) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *ComposeResponse) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *ComposeResponse) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *ComposeResponse) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *ComposeResponse) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *ComposeResponse) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetUrl

`func (o *ComposeResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComposeResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComposeResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWidth

`func (o *ComposeResponse) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ComposeResponse) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ComposeResponse) SetWidth(v int32)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


