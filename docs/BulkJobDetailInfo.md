# BulkJobDetailInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**FileSize** | Pointer to **NullableInt64** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** |  | 
**RenderTimeMs** | Pointer to **NullableInt64** |  | [optional] 
**ResultUrl** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Width** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBulkJobDetailInfo

`func NewBulkJobDetailInfo(createdAt time.Time, id string, status string, url string, ) *BulkJobDetailInfo`

NewBulkJobDetailInfo instantiates a new BulkJobDetailInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobDetailInfoWithDefaults

`func NewBulkJobDetailInfoWithDefaults() *BulkJobDetailInfo`

NewBulkJobDetailInfoWithDefaults instantiates a new BulkJobDetailInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BulkJobDetailInfo) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BulkJobDetailInfo) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BulkJobDetailInfo) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BulkJobDetailInfo) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BulkJobDetailInfo) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BulkJobDetailInfo) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *BulkJobDetailInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkJobDetailInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkJobDetailInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorCode

`func (o *BulkJobDetailInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BulkJobDetailInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BulkJobDetailInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BulkJobDetailInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BulkJobDetailInfo) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BulkJobDetailInfo) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BulkJobDetailInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BulkJobDetailInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BulkJobDetailInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BulkJobDetailInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BulkJobDetailInfo) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BulkJobDetailInfo) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFileSize

`func (o *BulkJobDetailInfo) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *BulkJobDetailInfo) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *BulkJobDetailInfo) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *BulkJobDetailInfo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *BulkJobDetailInfo) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *BulkJobDetailInfo) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetFormat

`func (o *BulkJobDetailInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BulkJobDetailInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BulkJobDetailInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BulkJobDetailInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *BulkJobDetailInfo) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *BulkJobDetailInfo) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetHeight

`func (o *BulkJobDetailInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BulkJobDetailInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BulkJobDetailInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BulkJobDetailInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *BulkJobDetailInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *BulkJobDetailInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetId

`func (o *BulkJobDetailInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobDetailInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobDetailInfo) SetId(v string)`

SetId sets Id field to given value.


### GetRenderTimeMs

`func (o *BulkJobDetailInfo) GetRenderTimeMs() int64`

GetRenderTimeMs returns the RenderTimeMs field if non-nil, zero value otherwise.

### GetRenderTimeMsOk

`func (o *BulkJobDetailInfo) GetRenderTimeMsOk() (*int64, bool)`

GetRenderTimeMsOk returns a tuple with the RenderTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTimeMs

`func (o *BulkJobDetailInfo) SetRenderTimeMs(v int64)`

SetRenderTimeMs sets RenderTimeMs field to given value.

### HasRenderTimeMs

`func (o *BulkJobDetailInfo) HasRenderTimeMs() bool`

HasRenderTimeMs returns a boolean if a field has been set.

### SetRenderTimeMsNil

`func (o *BulkJobDetailInfo) SetRenderTimeMsNil(b bool)`

 SetRenderTimeMsNil sets the value for RenderTimeMs to be an explicit nil

### UnsetRenderTimeMs
`func (o *BulkJobDetailInfo) UnsetRenderTimeMs()`

UnsetRenderTimeMs ensures that no value is present for RenderTimeMs, not even an explicit nil
### GetResultUrl

`func (o *BulkJobDetailInfo) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *BulkJobDetailInfo) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *BulkJobDetailInfo) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.

### HasResultUrl

`func (o *BulkJobDetailInfo) HasResultUrl() bool`

HasResultUrl returns a boolean if a field has been set.

### SetResultUrlNil

`func (o *BulkJobDetailInfo) SetResultUrlNil(b bool)`

 SetResultUrlNil sets the value for ResultUrl to be an explicit nil

### UnsetResultUrl
`func (o *BulkJobDetailInfo) UnsetResultUrl()`

UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
### GetStatus

`func (o *BulkJobDetailInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkJobDetailInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkJobDetailInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageUrl

`func (o *BulkJobDetailInfo) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *BulkJobDetailInfo) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *BulkJobDetailInfo) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *BulkJobDetailInfo) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *BulkJobDetailInfo) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *BulkJobDetailInfo) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetUrl

`func (o *BulkJobDetailInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkJobDetailInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkJobDetailInfo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWidth

`func (o *BulkJobDetailInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BulkJobDetailInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BulkJobDetailInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BulkJobDetailInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *BulkJobDetailInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *BulkJobDetailInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


