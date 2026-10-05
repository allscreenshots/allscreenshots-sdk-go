# CrawlPagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PageSize** | **int32** |  | 
**Pages** | [**[]CrawlPageResponse**](CrawlPageResponse.md) |  | 
**Total** | **int64** |  | 
**TotalPages** | **int32** |  | 

## Methods

### NewCrawlPagesResponse

`func NewCrawlPagesResponse(page int32, pageSize int32, pages []CrawlPageResponse, total int64, totalPages int32, ) *CrawlPagesResponse`

NewCrawlPagesResponse instantiates a new CrawlPagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlPagesResponseWithDefaults

`func NewCrawlPagesResponseWithDefaults() *CrawlPagesResponse`

NewCrawlPagesResponseWithDefaults instantiates a new CrawlPagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CrawlPagesResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CrawlPagesResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CrawlPagesResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *CrawlPagesResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CrawlPagesResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CrawlPagesResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPages

`func (o *CrawlPagesResponse) GetPages() []CrawlPageResponse`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *CrawlPagesResponse) GetPagesOk() (*[]CrawlPageResponse, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *CrawlPagesResponse) SetPages(v []CrawlPageResponse)`

SetPages sets Pages field to given value.


### GetTotal

`func (o *CrawlPagesResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CrawlPagesResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CrawlPagesResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *CrawlPagesResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *CrawlPagesResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *CrawlPagesResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


