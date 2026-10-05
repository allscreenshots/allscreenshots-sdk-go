# \ComposeAPI

All URIs are relative to *https://api.allscreenshots.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Compose**](ComposeAPI.md#Compose) | **Post** /v1/screenshots/compose | 
[**GetComposeJobStatus**](ComposeAPI.md#GetComposeJobStatus) | **Get** /v1/screenshots/compose/jobs/{jobId} | 
[**GetLayoutPreview**](ComposeAPI.md#GetLayoutPreview) | **Get** /v1/screenshots/compose/preview | 
[**ListComposeJobs**](ComposeAPI.md#ListComposeJobs) | **Get** /v1/screenshots/compose/jobs | 



## Compose

> *os.File Compose(ctx).ComposeRequest(composeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/allscreenshots/allscreenshots-sdk-go/v2"
)

func main() {
	composeRequest := *openapiclient.NewComposeRequest() // ComposeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.Compose(context.Background()).ComposeRequest(composeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.Compose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Compose`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.Compose`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **composeRequest** | [**ComposeRequest**](ComposeRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComposeJobStatus

> ComposeJobStatusResponse GetComposeJobStatus(ctx, jobId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/allscreenshots/allscreenshots-sdk-go/v2"
)

func main() {
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.GetComposeJobStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.GetComposeJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComposeJobStatus`: ComposeJobStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.GetComposeJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComposeJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComposeJobStatusResponse**](ComposeJobStatusResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayoutPreview

> LayoutPreviewResponse GetLayoutPreview(ctx).Layout(layout).ImageCount(imageCount).CanvasWidth(canvasWidth).CanvasHeight(canvasHeight).AspectRatios(aspectRatios).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/allscreenshots/allscreenshots-sdk-go/v2"
)

func main() {
	layout := "layout_example" // string | 
	imageCount := int32(56) // int32 | 
	canvasWidth := int32(56) // int32 |  (optional) (default to 1200)
	canvasHeight := int32(56) // int32 |  (optional) (default to 800)
	aspectRatios := []float64{float64(123)} // []float64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.GetLayoutPreview(context.Background()).Layout(layout).ImageCount(imageCount).CanvasWidth(canvasWidth).CanvasHeight(canvasHeight).AspectRatios(aspectRatios).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.GetLayoutPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLayoutPreview`: LayoutPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.GetLayoutPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLayoutPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layout** | **string** |  | 
 **imageCount** | **int32** |  | 
 **canvasWidth** | **int32** |  | [default to 1200]
 **canvasHeight** | **int32** |  | [default to 800]
 **aspectRatios** | **[]float64** |  | 

### Return type

[**LayoutPreviewResponse**](LayoutPreviewResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComposeJobs

> []ComposeJobSummaryResponse ListComposeJobs(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/allscreenshots/allscreenshots-sdk-go/v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.ListComposeJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.ListComposeJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComposeJobs`: []ComposeJobSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.ListComposeJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListComposeJobsRequest struct via the builder pattern


### Return type

[**[]ComposeJobSummaryResponse**](ComposeJobSummaryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

