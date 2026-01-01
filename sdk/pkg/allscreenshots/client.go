package allscreenshots

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Allscreenshots API.
	DefaultBaseURL = "https://api.allscreenshots.com"
	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 120 * time.Second
	// DefaultMaxRetries is the default number of retry attempts.
	DefaultMaxRetries = 3
	// DefaultRetryWaitMin is the minimum wait time between retries.
	DefaultRetryWaitMin = 1 * time.Second
	// DefaultRetryWaitMax is the maximum wait time between retries.
	DefaultRetryWaitMax = 30 * time.Second
	// EnvAPIKey is the environment variable name for the API key.
	EnvAPIKey = "ALLSCREENSHOTS_API_KEY"

	userAgent = "allscreenshots-sdk-go/1.0.0"
)

// Client is the Allscreenshots API client.
type Client struct {
	baseURL      string
	apiKey       string
	httpClient   *http.Client
	maxRetries   int
	retryWaitMin time.Duration
	retryWaitMax time.Duration
	userAgent    string
}

// ClientOption is a function that configures the client.
type ClientOption func(*Client)

// NewClient creates a new Allscreenshots API client.
//
// If no API key is provided via WithAPIKey, the client will attempt to read
// from the ALLSCREENSHOTS_API_KEY environment variable.
//
// Example:
//
//	client := allscreenshots.NewClient(
//	    allscreenshots.WithAPIKey("your-api-key"),
//	    allscreenshots.WithTimeout(60 * time.Second),
//	)
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		baseURL:      DefaultBaseURL,
		apiKey:       os.Getenv(EnvAPIKey),
		httpClient:   &http.Client{Timeout: DefaultTimeout},
		maxRetries:   DefaultMaxRetries,
		retryWaitMin: DefaultRetryWaitMin,
		retryWaitMax: DefaultRetryWaitMax,
		userAgent:    userAgent,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithBaseURL sets a custom base URL for the API.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithMaxRetries sets the maximum number of retry attempts.
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

// WithRetryWait sets the minimum and maximum wait times between retries.
func WithRetryWait(min, max time.Duration) ClientOption {
	return func(c *Client) {
		c.retryWaitMin = min
		c.retryWaitMax = max
	}
}

// WithUserAgent sets a custom user agent string.
func WithUserAgent(ua string) ClientOption {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// request performs an HTTP request with retries.
func (c *Client) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	return c.requestRaw(ctx, method, path, body, func(resp *http.Response) error {
		if result == nil {
			return nil
		}
		return json.NewDecoder(resp.Body).Decode(result)
	})
}

// requestBinary performs an HTTP request and returns raw bytes.
func (c *Client) requestBinary(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var data []byte
	err := c.requestRaw(ctx, method, path, body, func(resp *http.Response) error {
		var readErr error
		data, readErr = io.ReadAll(resp.Body)
		return readErr
	})
	return data, err
}

// requestRaw performs an HTTP request with a custom response handler.
func (c *Client) requestRaw(ctx context.Context, method, path string, body interface{}, handler func(*http.Response) error) error {
	if c.apiKey == "" {
		return &ValidationError{Field: "apiKey", Message: "API key is required"}
	}

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("allscreenshots: failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	reqURL := c.baseURL + path

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			// Calculate exponential backoff with jitter
			wait := c.calculateBackoff(attempt)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(wait):
			}

			// Reset body reader for retry
			if body != nil {
				jsonData, _ := json.Marshal(body)
				bodyReader = bytes.NewReader(jsonData)
			}
		}

		req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
		if err != nil {
			return fmt.Errorf("allscreenshots: failed to create request: %w", err)
		}

		req.Header.Set("X-API-Key", c.apiKey)
		req.Header.Set("User-Agent", c.userAgent)
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = &NetworkError{Message: "request failed", Cause: err}
			if isRetryableError(err) {
				continue
			}
			return lastErr
		}

		// Handle response
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			err := handler(resp)
			resp.Body.Close()
			return err
		}

		// Parse error response
		apiErr := c.parseErrorResponse(resp)
		resp.Body.Close()

		if isRetryableStatus(resp.StatusCode) {
			lastErr = apiErr
			continue
		}

		return apiErr
	}

	return &RetryError{Attempts: c.maxRetries + 1, LastErr: lastErr}
}

// calculateBackoff calculates the backoff duration for a retry attempt.
func (c *Client) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: min * 2^attempt
	backoff := float64(c.retryWaitMin) * math.Pow(2, float64(attempt-1))

	// Add jitter (up to 25% of backoff)
	jitter := backoff * 0.25 * rand.Float64()
	backoff += jitter

	// Cap at max
	if backoff > float64(c.retryWaitMax) {
		backoff = float64(c.retryWaitMax)
	}

	return time.Duration(backoff)
}

// parseErrorResponse parses an error response from the API.
func (c *Client) parseErrorResponse(resp *http.Response) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Message:    http.StatusText(resp.StatusCode),
	}

	var errResp struct {
		Error   string                 `json:"error"`
		Code    string                 `json:"code"`
		Message string                 `json:"message"`
		Details map[string]interface{} `json:"details"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
		if errResp.Message != "" {
			apiErr.Message = errResp.Message
		} else if errResp.Error != "" {
			apiErr.Message = errResp.Error
		}
		apiErr.Code = errResp.Code
		apiErr.Details = errResp.Details
	}

	return apiErr
}

// isRetryableError checks if an error is retryable.
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	// Check for timeout or temporary network errors
	if os.IsTimeout(err) {
		return true
	}
	errStr := err.Error()
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "no such host") ||
		strings.Contains(errStr, "timeout")
}

// isRetryableStatus checks if an HTTP status code is retryable.
func isRetryableStatus(status int) bool {
	return status == 429 || status == 502 || status == 503 || status == 504
}

// Screenshot captures a screenshot synchronously and returns the image bytes.
//
// Example:
//
//	imageData, err := client.Screenshot(ctx, &allscreenshots.ScreenshotRequest{
//	    URL:    "https://github.com",
//	    Device: "Desktop HD",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	os.WriteFile("screenshot.png", imageData, 0644)
func (c *Client) Screenshot(ctx context.Context, req *ScreenshotRequest) ([]byte, error) {
	if err := validateScreenshotRequest(req); err != nil {
		return nil, err
	}

	return c.requestBinary(ctx, http.MethodPost, "/v1/screenshots", req)
}

// ScreenshotAsync starts an asynchronous screenshot capture.
//
// Example:
//
//	job, err := client.ScreenshotAsync(ctx, &allscreenshots.ScreenshotRequest{
//	    URL:      "https://github.com",
//	    Device:   "Desktop HD",
//	    FullPage: true,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job created: %s\n", job.ID)
func (c *Client) ScreenshotAsync(ctx context.Context, req *ScreenshotRequest) (*AsyncJobCreatedResponse, error) {
	if err := validateScreenshotRequest(req); err != nil {
		return nil, err
	}

	var result AsyncJobCreatedResponse
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/async", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListJobs returns all screenshot jobs.
//
// Example:
//
//	jobs, err := client.ListJobs(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, job := range jobs {
//	    fmt.Printf("Job %s: %s\n", job.ID, job.Status)
//	}
func (c *Client) ListJobs(ctx context.Context) ([]JobResponse, error) {
	var result []JobResponse
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/jobs", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetJob returns the status of a specific job.
//
// Example:
//
//	job, err := client.GetJob(ctx, "job-123")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Status: %s\n", job.Status)
func (c *Client) GetJob(ctx context.Context, id string) (*JobResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "job ID is required"}
	}

	var result JobResponse
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/jobs/"+url.PathEscape(id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetJobResult returns the screenshot image for a completed job.
//
// Example:
//
//	imageData, err := client.GetJobResult(ctx, "job-123")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	os.WriteFile("screenshot.png", imageData, 0644)
func (c *Client) GetJobResult(ctx context.Context, id string) ([]byte, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "job ID is required"}
	}

	return c.requestBinary(ctx, http.MethodGet, "/v1/screenshots/jobs/"+url.PathEscape(id)+"/result", nil)
}

// CancelJob cancels a pending or processing job.
//
// Example:
//
//	job, err := client.CancelJob(ctx, "job-123")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job %s cancelled\n", job.ID)
func (c *Client) CancelJob(ctx context.Context, id string) (*JobResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "job ID is required"}
	}

	var result JobResponse
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/jobs/"+url.PathEscape(id)+"/cancel", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateBulkJob creates a bulk screenshot job.
//
// Example:
//
//	bulk, err := client.CreateBulkJob(ctx, &allscreenshots.BulkRequest{
//	    URLs: []allscreenshots.BulkURLRequest{
//	        {URL: "https://github.com"},
//	        {URL: "https://google.com"},
//	    },
//	    Defaults: &allscreenshots.BulkDefaults{
//	        Device: "Desktop HD",
//	    },
//	})
func (c *Client) CreateBulkJob(ctx context.Context, req *BulkRequest) (*BulkResponse, error) {
	if err := validateBulkRequest(req); err != nil {
		return nil, err
	}

	var result BulkResponse
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/bulk", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListBulkJobs returns all bulk screenshot jobs.
func (c *Client) ListBulkJobs(ctx context.Context) ([]BulkJobSummary, error) {
	var result []BulkJobSummary
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/bulk", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBulkJob returns the status of a bulk job.
func (c *Client) GetBulkJob(ctx context.Context, id string) (*BulkStatusResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "bulk job ID is required"}
	}

	var result BulkStatusResponse
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/bulk/"+url.PathEscape(id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelBulkJob cancels a bulk job.
func (c *Client) CancelBulkJob(ctx context.Context, id string) (*BulkJobSummary, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "bulk job ID is required"}
	}

	var result BulkJobSummary
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/bulk/"+url.PathEscape(id)+"/cancel", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Compose creates a composed image from multiple screenshots.
//
// Example:
//
//	compose, err := client.Compose(ctx, &allscreenshots.ComposeRequest{
//	    Captures: []allscreenshots.CaptureItem{
//	        {URL: "https://github.com", Device: "Desktop HD"},
//	        {URL: "https://github.com", Device: "iPhone 14"},
//	    },
//	    Output: &allscreenshots.ComposeOutputConfig{
//	        Layout: "HORIZONTAL",
//	    },
//	})
func (c *Client) Compose(ctx context.Context, req *ComposeRequest) (*ComposeResponse, error) {
	if err := validateComposeRequest(req); err != nil {
		return nil, err
	}

	var result ComposeResponse
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/compose", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ComposeAsync creates a composed image asynchronously.
func (c *Client) ComposeAsync(ctx context.Context, req *ComposeRequest) (*ComposeJobStatusResponse, error) {
	if err := validateComposeRequest(req); err != nil {
		return nil, err
	}
	req.Async = true

	var result ComposeJobStatusResponse
	err := c.request(ctx, http.MethodPost, "/v1/screenshots/compose", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ComposeLayoutPreviewParams represents parameters for layout preview.
type ComposeLayoutPreviewParams struct {
	Layout       string
	ImageCount   int
	CanvasWidth  int
	CanvasHeight int
	AspectRatios string
}

// GetComposeLayoutPreview returns a preview of a compose layout.
func (c *Client) GetComposeLayoutPreview(ctx context.Context, params *ComposeLayoutPreviewParams) (*LayoutPreviewResponse, error) {
	path := "/v1/screenshots/compose/preview"

	query := url.Values{}
	if params.Layout != "" {
		query.Set("layout", params.Layout)
	}
	if params.ImageCount > 0 {
		query.Set("image_count", strconv.Itoa(params.ImageCount))
	}
	if params.CanvasWidth > 0 {
		query.Set("canvas_width", strconv.Itoa(params.CanvasWidth))
	}
	if params.CanvasHeight > 0 {
		query.Set("canvas_height", strconv.Itoa(params.CanvasHeight))
	}
	if params.AspectRatios != "" {
		query.Set("aspect_ratios", params.AspectRatios)
	}

	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var result LayoutPreviewResponse
	err := c.request(ctx, http.MethodGet, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListComposeJobs returns all compose jobs.
func (c *Client) ListComposeJobs(ctx context.Context) ([]ComposeJobSummaryResponse, error) {
	var result []ComposeJobSummaryResponse
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/compose/jobs", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetComposeJob returns the status of a compose job.
func (c *Client) GetComposeJob(ctx context.Context, jobID string) (*ComposeJobStatusResponse, error) {
	if jobID == "" {
		return nil, &ValidationError{Field: "jobId", Message: "job ID is required"}
	}

	var result ComposeJobStatusResponse
	err := c.request(ctx, http.MethodGet, "/v1/screenshots/compose/jobs/"+url.PathEscape(jobID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSchedule creates a new scheduled screenshot.
//
// Example:
//
//	schedule, err := client.CreateSchedule(ctx, &allscreenshots.CreateScheduleRequest{
//	    Name:     "Daily GitHub Snapshot",
//	    URL:      "https://github.com",
//	    Schedule: "0 9 * * *",  // Every day at 9 AM
//	    Timezone: "America/New_York",
//	})
func (c *Client) CreateSchedule(ctx context.Context, req *CreateScheduleRequest) (*ScheduleResponse, error) {
	if err := validateCreateScheduleRequest(req); err != nil {
		return nil, err
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodPost, "/v1/schedules", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListSchedules returns all schedules.
func (c *Client) ListSchedules(ctx context.Context) (*ScheduleListResponse, error) {
	var result ScheduleListResponse
	err := c.request(ctx, http.MethodGet, "/v1/schedules", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSchedule returns a specific schedule.
func (c *Client) GetSchedule(ctx context.Context, id string) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodGet, "/v1/schedules/"+url.PathEscape(id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSchedule updates a schedule.
func (c *Client) UpdateSchedule(ctx context.Context, id string, req *UpdateScheduleRequest) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodPut, "/v1/schedules/"+url.PathEscape(id), req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteSchedule deletes a schedule.
func (c *Client) DeleteSchedule(ctx context.Context, id string) error {
	if id == "" {
		return &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	return c.request(ctx, http.MethodDelete, "/v1/schedules/"+url.PathEscape(id), nil, nil)
}

// PauseSchedule pauses a schedule.
func (c *Client) PauseSchedule(ctx context.Context, id string) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodPost, "/v1/schedules/"+url.PathEscape(id)+"/pause", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ResumeSchedule resumes a paused schedule.
func (c *Client) ResumeSchedule(ctx context.Context, id string) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodPost, "/v1/schedules/"+url.PathEscape(id)+"/resume", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TriggerSchedule manually triggers a schedule execution.
func (c *Client) TriggerSchedule(ctx context.Context, id string) (*ScheduleResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	var result ScheduleResponse
	err := c.request(ctx, http.MethodPost, "/v1/schedules/"+url.PathEscape(id)+"/trigger", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetScheduleHistory returns execution history for a schedule.
func (c *Client) GetScheduleHistory(ctx context.Context, id string, limit int) (*ScheduleHistoryResponse, error) {
	if id == "" {
		return nil, &ValidationError{Field: "id", Message: "schedule ID is required"}
	}

	path := "/v1/schedules/" + url.PathEscape(id) + "/history"
	if limit > 0 {
		path += "?limit=" + strconv.Itoa(limit)
	}

	var result ScheduleHistoryResponse
	err := c.request(ctx, http.MethodGet, path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUsage returns usage statistics.
//
// Example:
//
//	usage, err := client.GetUsage(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Tier: %s\n", usage.Tier)
//	fmt.Printf("Screenshots this period: %d\n", usage.CurrentPeriod.ScreenshotsCount)
func (c *Client) GetUsage(ctx context.Context) (*UsageResponse, error) {
	var result UsageResponse
	err := c.request(ctx, http.MethodGet, "/v1/usage", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetQuotaStatus returns quota status.
//
// Example:
//
//	quota, err := client.GetQuotaStatus(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Screenshots remaining: %d\n", quota.Screenshots.Remaining)
func (c *Client) GetQuotaStatus(ctx context.Context) (*QuotaStatusResponse, error) {
	var result QuotaStatusResponse
	err := c.request(ctx, http.MethodGet, "/v1/usage/quota", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// validateScreenshotRequest validates a screenshot request.
func validateScreenshotRequest(req *ScreenshotRequest) error {
	if req == nil {
		return &ValidationError{Field: "request", Message: "request cannot be nil"}
	}
	if req.URL == "" {
		return &ValidationError{Field: "url", Message: "URL is required"}
	}
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return &ValidationError{Field: "url", Message: "URL must start with http:// or https://"}
	}
	if req.Quality != 0 && (req.Quality < 1 || req.Quality > 100) {
		return &ValidationError{Field: "quality", Message: "quality must be between 1 and 100"}
	}
	if req.Delay != 0 && (req.Delay < 0 || req.Delay > 30000) {
		return &ValidationError{Field: "delay", Message: "delay must be between 0 and 30000"}
	}
	if req.Timeout != 0 && (req.Timeout < 1000 || req.Timeout > 60000) {
		return &ValidationError{Field: "timeout", Message: "timeout must be between 1000 and 60000"}
	}
	if req.Viewport != nil {
		if err := validateViewport(req.Viewport); err != nil {
			return err
		}
	}
	return nil
}

// validateViewport validates viewport configuration.
func validateViewport(v *ViewportConfig) error {
	if v.Width != 0 && (v.Width < 100 || v.Width > 4096) {
		return &ValidationError{Field: "viewport.width", Message: "width must be between 100 and 4096"}
	}
	if v.Height != 0 && (v.Height < 100 || v.Height > 4096) {
		return &ValidationError{Field: "viewport.height", Message: "height must be between 100 and 4096"}
	}
	if v.DeviceScaleFactor != 0 && (v.DeviceScaleFactor < 1 || v.DeviceScaleFactor > 3) {
		return &ValidationError{Field: "viewport.deviceScaleFactor", Message: "deviceScaleFactor must be between 1 and 3"}
	}
	return nil
}

// validateBulkRequest validates a bulk request.
func validateBulkRequest(req *BulkRequest) error {
	if req == nil {
		return &ValidationError{Field: "request", Message: "request cannot be nil"}
	}
	if len(req.URLs) == 0 {
		return &ValidationError{Field: "urls", Message: "at least one URL is required"}
	}
	if len(req.URLs) > 100 {
		return &ValidationError{Field: "urls", Message: "maximum 100 URLs allowed"}
	}
	for i, u := range req.URLs {
		if u.URL == "" {
			return &ValidationError{Field: fmt.Sprintf("urls[%d].url", i), Message: "URL is required"}
		}
		if !strings.HasPrefix(u.URL, "http://") && !strings.HasPrefix(u.URL, "https://") {
			return &ValidationError{Field: fmt.Sprintf("urls[%d].url", i), Message: "URL must start with http:// or https://"}
		}
	}
	return nil
}

// validateComposeRequest validates a compose request.
func validateComposeRequest(req *ComposeRequest) error {
	if req == nil {
		return &ValidationError{Field: "request", Message: "request cannot be nil"}
	}
	if len(req.Captures) == 0 && req.URL == "" {
		return &ValidationError{Field: "captures", Message: "either captures or url is required"}
	}
	if len(req.Captures) > 20 {
		return &ValidationError{Field: "captures", Message: "maximum 20 captures allowed"}
	}
	if len(req.Variants) > 20 {
		return &ValidationError{Field: "variants", Message: "maximum 20 variants allowed"}
	}
	for i, c := range req.Captures {
		if c.URL == "" {
			return &ValidationError{Field: fmt.Sprintf("captures[%d].url", i), Message: "URL is required"}
		}
		if !strings.HasPrefix(c.URL, "http://") && !strings.HasPrefix(c.URL, "https://") {
			return &ValidationError{Field: fmt.Sprintf("captures[%d].url", i), Message: "URL must start with http:// or https://"}
		}
	}
	return nil
}

// validateCreateScheduleRequest validates a create schedule request.
func validateCreateScheduleRequest(req *CreateScheduleRequest) error {
	if req == nil {
		return &ValidationError{Field: "request", Message: "request cannot be nil"}
	}
	if req.Name == "" {
		return &ValidationError{Field: "name", Message: "name is required"}
	}
	if len(req.Name) > 255 {
		return &ValidationError{Field: "name", Message: "name must be at most 255 characters"}
	}
	if req.URL == "" {
		return &ValidationError{Field: "url", Message: "URL is required"}
	}
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		return &ValidationError{Field: "url", Message: "URL must start with http:// or https://"}
	}
	if req.Schedule == "" {
		return &ValidationError{Field: "schedule", Message: "schedule is required"}
	}
	if req.RetentionDays != 0 && (req.RetentionDays < 1 || req.RetentionDays > 365) {
		return &ValidationError{Field: "retentionDays", Message: "retentionDays must be between 1 and 365"}
	}
	return nil
}
