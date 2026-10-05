# CrawlPageOutputResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**ResultUrl** | **string** |  | 
**Size** | **int64** |  | 
**Type** | **string** |  | 

## Methods

### NewCrawlPageOutputResponse

`func NewCrawlPageOutputResponse(contentType string, expiresAt time.Time, resultUrl string, size int64, type_ string, ) *CrawlPageOutputResponse`

NewCrawlPageOutputResponse instantiates a new CrawlPageOutputResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlPageOutputResponseWithDefaults

`func NewCrawlPageOutputResponseWithDefaults() *CrawlPageOutputResponse`

NewCrawlPageOutputResponseWithDefaults instantiates a new CrawlPageOutputResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *CrawlPageOutputResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CrawlPageOutputResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CrawlPageOutputResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetExpiresAt

`func (o *CrawlPageOutputResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CrawlPageOutputResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CrawlPageOutputResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetResultUrl

`func (o *CrawlPageOutputResponse) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *CrawlPageOutputResponse) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *CrawlPageOutputResponse) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.


### GetSize

`func (o *CrawlPageOutputResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CrawlPageOutputResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CrawlPageOutputResponse) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *CrawlPageOutputResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrawlPageOutputResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrawlPageOutputResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


