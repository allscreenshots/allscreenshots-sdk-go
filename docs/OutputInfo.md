# OutputInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**ResultUrl** | Pointer to **NullableString** |  | [optional] 
**Size** | **int64** |  | 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewOutputInfo

`func NewOutputInfo(contentType string, size int64, type_ string, ) *OutputInfo`

NewOutputInfo instantiates a new OutputInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputInfoWithDefaults

`func NewOutputInfoWithDefaults() *OutputInfo`

NewOutputInfoWithDefaults instantiates a new OutputInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *OutputInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OutputInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OutputInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetResultUrl

`func (o *OutputInfo) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *OutputInfo) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *OutputInfo) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.

### HasResultUrl

`func (o *OutputInfo) HasResultUrl() bool`

HasResultUrl returns a boolean if a field has been set.

### SetResultUrlNil

`func (o *OutputInfo) SetResultUrlNil(b bool)`

 SetResultUrlNil sets the value for ResultUrl to be an explicit nil

### UnsetResultUrl
`func (o *OutputInfo) UnsetResultUrl()`

UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
### GetSize

`func (o *OutputInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OutputInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OutputInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStorageUrl

`func (o *OutputInfo) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *OutputInfo) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *OutputInfo) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *OutputInfo) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *OutputInfo) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *OutputInfo) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetType

`func (o *OutputInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


