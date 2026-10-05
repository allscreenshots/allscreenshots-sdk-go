# \ScreenshotAPI

All URIs are relative to *https://api.allscreenshots.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaptureSync**](ScreenshotAPI.md#CaptureSync) | **Post** /v1/screenshots | 
[**GetSyncCaptureOutputResult**](ScreenshotAPI.md#GetSyncCaptureOutputResult) | **Get** /v1/screenshots/captures/{id}/result/{outputId} | 
[**GetSyncCaptureResult**](ScreenshotAPI.md#GetSyncCaptureResult) | **Get** /v1/screenshots/captures/{id}/result | 
[**SubmissionStatus**](ScreenshotAPI.md#SubmissionStatus) | **Get** /v1/screenshots/submissions/{key} | 



## CaptureSync

> *os.File CaptureSync(ctx).ScreenshotRequest(screenshotRequest).IdempotencyKey(idempotencyKey).Execute()



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
	screenshotRequest := *openapiclient.NewScreenshotRequest("Url_example") // ScreenshotRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Unique submission key. Supported for responseType=url; reuse with the same body to recover the original result. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenshotAPI.CaptureSync(context.Background()).ScreenshotRequest(screenshotRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotAPI.CaptureSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaptureSync`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ScreenshotAPI.CaptureSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaptureSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **screenshotRequest** | [**ScreenshotRequest**](ScreenshotRequest.md) |  | 
 **idempotencyKey** | **string** | Unique submission key. Supported for responseType&#x3D;url; reuse with the same body to recover the original result. | 

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


## GetSyncCaptureOutputResult

> *os.File GetSyncCaptureOutputResult(ctx, id, outputId).Execute()



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
	id := "id_example" // string | 
	outputId := "outputId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenshotAPI.GetSyncCaptureOutputResult(context.Background(), id, outputId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotAPI.GetSyncCaptureOutputResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncCaptureOutputResult`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ScreenshotAPI.GetSyncCaptureOutputResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**outputId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncCaptureOutputResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncCaptureResult

> *os.File GetSyncCaptureResult(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenshotAPI.GetSyncCaptureResult(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotAPI.GetSyncCaptureResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncCaptureResult`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ScreenshotAPI.GetSyncCaptureResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncCaptureResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmissionStatus

> map[string]interface{} SubmissionStatus(ctx, key).Execute()



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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenshotAPI.SubmissionStatus(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenshotAPI.SubmissionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmissionStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScreenshotAPI.SubmissionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmissionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

