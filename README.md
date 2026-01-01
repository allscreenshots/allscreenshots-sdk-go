# Allscreenshots SDK for Go

A Go SDK for the [Allscreenshots](https://allscreenshots.com) screenshot API. Capture screenshots of web pages with various options for viewport, device emulation, and output formats.

## Installation

```bash
go get github.com/allscreenshots/allscreenshots-sdk-go
```

## Quick start

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/allscreenshots/allscreenshots-sdk-go/pkg/allscreenshots"
)

func main() {
    // Create client (reads ALLSCREENSHOTS_API_KEY from env by default)
    client := allscreenshots.NewClient()

    // Take a screenshot
    imageData, err := client.Screenshot(context.Background(), &allscreenshots.ScreenshotRequest{
        URL:    "https://github.com",
        Device: "Desktop HD",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Save the screenshot
    if err := os.WriteFile("screenshot.png", imageData, 0644); err != nil {
        log.Fatal(err)
    }
}
```

## Configuration

### Client options

```go
client := allscreenshots.NewClient(
    // Explicitly set API key (otherwise reads from ALLSCREENSHOTS_API_KEY env var)
    allscreenshots.WithAPIKey("your-api-key"),

    // Custom base URL (default: https://api.allscreenshots.com)
    allscreenshots.WithBaseURL("https://custom.api.com"),

    // HTTP timeout (default: 120s)
    allscreenshots.WithTimeout(60 * time.Second),

    // Retry configuration (default: 3 retries, 1-30s wait)
    allscreenshots.WithMaxRetries(5),
    allscreenshots.WithRetryWait(2*time.Second, 60*time.Second),

    // Custom HTTP client
    allscreenshots.WithHTTPClient(&http.Client{}),
)
```

### Environment variables

| Variable | Description |
|----------|-------------|
| `ALLSCREENSHOTS_API_KEY` | API key for authentication (used if not set via `WithAPIKey`) |

## API reference

### Screenshots

#### Synchronous screenshot

```go
imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
    URL:      "https://example.com",
    Device:   "Desktop HD",        // Device preset
    FullPage: true,                // Capture entire page
    Format:   "png",               // Output format: png, jpeg, webp, pdf
    Quality:  90,                  // Quality (1-100, for jpeg/webp)
    DarkMode: true,                // Enable dark mode
    Delay:    1000,                // Wait before capture (ms)
    Timeout:  30000,               // Timeout (ms)
})
```

#### Asynchronous screenshot

```go
// Start async capture
job, err := client.ScreenshotAsync(ctx, &allscreenshots.ScreenshotRequest{
    URL:    "https://example.com",
    Device: "Desktop HD",
})
if err != nil {
    log.Fatal(err)
}

// Poll for completion
for {
    status, err := client.GetJob(ctx, job.ID)
    if err != nil {
        log.Fatal(err)
    }

    if status.Status == allscreenshots.JobStatusCompleted {
        // Download result
        imageData, err := client.GetJobResult(ctx, job.ID)
        break
    } else if status.Status == allscreenshots.JobStatusFailed {
        log.Fatalf("Job failed: %s", status.ErrorMessage)
    }

    time.Sleep(1 * time.Second)
}
```

#### Job management

```go
// List all jobs
jobs, err := client.ListJobs(ctx)

// Get specific job
job, err := client.GetJob(ctx, "job-id")

// Get job result (image data)
imageData, err := client.GetJobResult(ctx, "job-id")

// Cancel a job
job, err := client.CancelJob(ctx, "job-id")
```

### Bulk screenshots

```go
// Create bulk job
bulk, err := client.CreateBulkJob(ctx, &allscreenshots.BulkRequest{
    URLs: []allscreenshots.BulkURLRequest{
        {URL: "https://example.com"},
        {URL: "https://github.com"},
        {URL: "https://google.com"},
    },
    Defaults: &allscreenshots.BulkDefaults{
        Device:   "Desktop HD",
        FullPage: false,
    },
})

// List bulk jobs
bulkJobs, err := client.ListBulkJobs(ctx)

// Get bulk job status
status, err := client.GetBulkJob(ctx, "bulk-id")

// Cancel bulk job
cancelled, err := client.CancelBulkJob(ctx, "bulk-id")
```

### Compose (multi-screenshot layouts)

```go
// Compose multiple screenshots into one image
result, err := client.Compose(ctx, &allscreenshots.ComposeRequest{
    Captures: []allscreenshots.CaptureItem{
        {URL: "https://example.com", Device: "Desktop HD"},
        {URL: "https://example.com", Device: "iPhone 14"},
        {URL: "https://example.com", Device: "iPad"},
    },
    Output: &allscreenshots.ComposeOutputConfig{
        Layout:  "HORIZONTAL",
        Spacing: 20,
        Padding: 40,
    },
})

// Preview layout placement
preview, err := client.GetComposeLayoutPreview(ctx, &allscreenshots.ComposeLayoutPreviewParams{
    Layout:     "GRID",
    ImageCount: 4,
})

// Async compose
job, err := client.ComposeAsync(ctx, &allscreenshots.ComposeRequest{...})

// List compose jobs
jobs, err := client.ListComposeJobs(ctx)

// Get compose job status
status, err := client.GetComposeJob(ctx, "job-id")
```

### Schedules

```go
// Create a schedule
schedule, err := client.CreateSchedule(ctx, &allscreenshots.CreateScheduleRequest{
    Name:     "Daily GitHub Snapshot",
    URL:      "https://github.com",
    Schedule: "0 9 * * *",  // Every day at 9 AM
    Timezone: "America/New_York",
    Options: &allscreenshots.ScheduleScreenshotOptions{
        Device:   "Desktop HD",
        FullPage: true,
    },
})

// List schedules
list, err := client.ListSchedules(ctx)

// Get schedule
schedule, err := client.GetSchedule(ctx, "schedule-id")

// Update schedule
schedule, err := client.UpdateSchedule(ctx, "schedule-id", &allscreenshots.UpdateScheduleRequest{
    Name: "Updated Name",
})

// Pause/resume schedule
schedule, err := client.PauseSchedule(ctx, "schedule-id")
schedule, err := client.ResumeSchedule(ctx, "schedule-id")

// Manually trigger
schedule, err := client.TriggerSchedule(ctx, "schedule-id")

// Get execution history
history, err := client.GetScheduleHistory(ctx, "schedule-id", 10)

// Delete schedule
err := client.DeleteSchedule(ctx, "schedule-id")
```

### Usage and quota

```go
// Get usage statistics
usage, err := client.GetUsage(ctx)
fmt.Printf("Tier: %s\n", usage.Tier)
fmt.Printf("Screenshots this period: %d\n", usage.CurrentPeriod.ScreenshotsCount)

// Get quota status
quota, err := client.GetQuotaStatus(ctx)
fmt.Printf("Screenshots remaining: %d\n", quota.Screenshots.Remaining)
```

## Device presets

The API supports various device presets:

| Preset | Viewport |
|--------|----------|
| `Desktop HD` | 1920x1080 |
| `Desktop` | 1440x900 |
| `Laptop` | 1366x768 |
| `iPhone 14` | 390x844 |
| `iPhone 14 Pro Max` | 430x932 |
| `iPad` | 820x1180 |
| `iPad Pro` | 1024x1366 |

You can also specify custom viewports:

```go
&allscreenshots.ScreenshotRequest{
    URL: "https://example.com",
    Viewport: &allscreenshots.ViewportConfig{
        Width:             1920,
        Height:            1080,
        DeviceScaleFactor: 2,
    },
}
```

## Error handling

The SDK provides typed errors for different failure scenarios:

```go
imageData, err := client.Screenshot(ctx, req)
if err != nil {
    // Check error type
    if allscreenshots.IsValidationError(err) {
        // Invalid request parameters
        log.Printf("Validation error: %v", err)
    } else if allscreenshots.IsUnauthorized(err) {
        // Invalid or missing API key
        log.Printf("Unauthorized: %v", err)
    } else if allscreenshots.IsRateLimited(err) {
        // Too many requests
        log.Printf("Rate limited: %v", err)
    } else if allscreenshots.IsNotFound(err) {
        // Resource not found
        log.Printf("Not found: %v", err)
    } else if allscreenshots.IsServerError(err) {
        // Server-side error
        log.Printf("Server error: %v", err)
    } else if allscreenshots.IsRetryError(err) {
        // All retry attempts exhausted
        log.Printf("Retry exhausted: %v", err)
    }

    // Get detailed API error info
    if apiErr, ok := allscreenshots.AsAPIError(err); ok {
        log.Printf("Status: %d, Code: %s, Message: %s",
            apiErr.StatusCode, apiErr.Code, apiErr.Message)
    }
}
```

### Error types

| Type | Description |
|------|-------------|
| `*ValidationError` | Invalid request parameters (client-side validation) |
| `*APIError` | Error response from the API |
| `*NetworkError` | Network connectivity issues |
| `*TimeoutError` | Request timeout |
| `*RetryError` | All retry attempts exhausted |

### Helper functions

| Function | Description |
|----------|-------------|
| `IsValidationError(err)` | Check if error is a validation error |
| `IsAPIError(err)` | Check if error is an API error |
| `IsNetworkError(err)` | Check if error is a network error |
| `IsRetryError(err)` | Check if error is a retry error |
| `IsBadRequest(err)` | Check if error is 400 Bad Request |
| `IsUnauthorized(err)` | Check if error is 401 Unauthorized |
| `IsForbidden(err)` | Check if error is 403 Forbidden |
| `IsNotFound(err)` | Check if error is 404 Not Found |
| `IsRateLimited(err)` | Check if error is 429 Too Many Requests |
| `IsServerError(err)` | Check if error is 5xx Server Error |

## Retry behavior

The SDK automatically retries requests on transient failures:

- **Retried statuses**: 429 (Too Many Requests), 502, 503, 504
- **Default retries**: 3 attempts
- **Backoff**: Exponential with jitter (1s to 30s)

Customize retry behavior:

```go
client := allscreenshots.NewClient(
    allscreenshots.WithMaxRetries(5),
    allscreenshots.WithRetryWait(2*time.Second, 60*time.Second),
)
```

To disable retries:

```go
client := allscreenshots.NewClient(
    allscreenshots.WithMaxRetries(0),
)
```

## Testing

Run unit tests:

```bash
go test ./pkg/allscreenshots/...
```

Run integration tests (requires API key):

```bash
export ALLSCREENSHOTS_API_KEY="your-api-key"
go test ./tests/integration/...
```

## License

Apache License 2.0. See [LICENSE](LICENSE) for details.
