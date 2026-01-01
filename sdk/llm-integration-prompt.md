# Allscreenshots Go SDK - LLM integration prompt

Use this prompt to help LLMs understand how to use the Allscreenshots Go SDK.

---

## SDK overview

The Allscreenshots Go SDK allows you to capture screenshots of web pages programmatically. It provides a simple, idiomatic Go interface with support for synchronous and asynchronous operations, retry logic, and typed errors.

## Installation

```bash
go get github.com/allscreenshots/allscreenshots-sdk-go
```

## Import

```go
import "github.com/allscreenshots/allscreenshots-sdk-go/pkg/allscreenshots"
```

## Authentication

The SDK reads the API key from the `ALLSCREENSHOTS_API_KEY` environment variable by default. You can also set it explicitly:

```go
client := allscreenshots.NewClient(
    allscreenshots.WithAPIKey("your-api-key"),
)
```

## Common operations

### Take a screenshot

```go
imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
    URL:    "https://example.com",
    Device: "Desktop HD",
})
if err != nil {
    log.Fatal(err)
}
os.WriteFile("screenshot.png", imageData, 0644)
```

### Take a full-page screenshot

```go
imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
    URL:      "https://example.com",
    Device:   "Desktop HD",
    FullPage: true,
})
```

### Take a mobile screenshot

```go
imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
    URL:    "https://example.com",
    Device: "iPhone 14",
})
```

### Take a screenshot with custom viewport

```go
imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
    URL: "https://example.com",
    Viewport: &allscreenshots.ViewportConfig{
        Width:  1920,
        Height: 1080,
    },
})
```

### Take an async screenshot

```go
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
    if status.Status == allscreenshots.JobStatusCompleted {
        imageData, _ := client.GetJobResult(ctx, job.ID)
        break
    }
    time.Sleep(time.Second)
}
```

### Take bulk screenshots

```go
bulk, err := client.CreateBulkJob(ctx, &allscreenshots.BulkRequest{
    URLs: []allscreenshots.BulkURLRequest{
        {URL: "https://example.com"},
        {URL: "https://github.com"},
    },
    Defaults: &allscreenshots.BulkDefaults{
        Device: "Desktop HD",
    },
})
```

### Compose multiple screenshots

```go
result, err := client.Compose(ctx, &allscreenshots.ComposeRequest{
    Captures: []allscreenshots.CaptureItem{
        {URL: "https://example.com", Device: "Desktop HD"},
        {URL: "https://example.com", Device: "iPhone 14"},
    },
    Output: &allscreenshots.ComposeOutputConfig{
        Layout: "HORIZONTAL",
    },
})
```

### Create a scheduled screenshot

```go
schedule, err := client.CreateSchedule(ctx, &allscreenshots.CreateScheduleRequest{
    Name:     "Daily Snapshot",
    URL:      "https://example.com",
    Schedule: "0 9 * * *",  // 9 AM daily
    Timezone: "UTC",
})
```

### Check usage

```go
usage, err := client.GetUsage(ctx)
fmt.Printf("Screenshots used: %d\n", usage.CurrentPeriod.ScreenshotsCount)
```

## Device presets

Available device presets: `Desktop HD`, `Desktop`, `Laptop`, `iPhone 14`, `iPhone 14 Pro Max`, `iPad`, `iPad Pro`

## Error handling

```go
imageData, err := client.Screenshot(ctx, req)
if err != nil {
    if allscreenshots.IsValidationError(err) {
        // Invalid request
    } else if allscreenshots.IsUnauthorized(err) {
        // Invalid API key
    } else if allscreenshots.IsRateLimited(err) {
        // Too many requests
    }
}
```

## Configuration options

```go
client := allscreenshots.NewClient(
    allscreenshots.WithAPIKey("key"),           // API key
    allscreenshots.WithBaseURL("url"),          // Custom base URL
    allscreenshots.WithTimeout(60*time.Second), // HTTP timeout
    allscreenshots.WithMaxRetries(5),           // Retry attempts
)
```
