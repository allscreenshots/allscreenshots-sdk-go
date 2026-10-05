# CrawlListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crawls** | [**[]CrawlResponse**](CrawlResponse.md) |  | 
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 
**Total** | **int64** |  | 
**TotalPages** | **int32** |  | 

## Methods

### NewCrawlListResponse

`func NewCrawlListResponse(crawls []CrawlResponse, page int32, pageSize int32, total int64, totalPages int32, ) *CrawlListResponse`

NewCrawlListResponse instantiates a new CrawlListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlListResponseWithDefaults

`func NewCrawlListResponseWithDefaults() *CrawlListResponse`

NewCrawlListResponseWithDefaults instantiates a new CrawlListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrawls

`func (o *CrawlListResponse) GetCrawls() []CrawlResponse`

GetCrawls returns the Crawls field if non-nil, zero value otherwise.

### GetCrawlsOk

`func (o *CrawlListResponse) GetCrawlsOk() (*[]CrawlResponse, bool)`

GetCrawlsOk returns a tuple with the Crawls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawls

`func (o *CrawlListResponse) SetCrawls(v []CrawlResponse)`

SetCrawls sets Crawls field to given value.


### GetPage

`func (o *CrawlListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CrawlListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CrawlListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *CrawlListResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CrawlListResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CrawlListResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotal

`func (o *CrawlListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CrawlListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CrawlListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *CrawlListResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *CrawlListResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *CrawlListResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


