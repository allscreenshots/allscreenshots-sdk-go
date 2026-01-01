// Package allscreenshots provides a Go SDK for the Allscreenshots API.
//
// The Allscreenshots API allows you to capture screenshots of web pages
// with various options for viewport, device emulation, and output format.
package allscreenshots

import "time"

// ViewportConfig represents viewport dimensions and scale factor.
type ViewportConfig struct {
	// Width of the viewport (100-4096 pixels)
	Width int `json:"width,omitempty"`
	// Height of the viewport (100-4096 pixels)
	Height int `json:"height,omitempty"`
	// DeviceScaleFactor is the device scale factor (1-3)
	DeviceScaleFactor int `json:"deviceScaleFactor,omitempty"`
}

// ScreenshotRequest represents a request to capture a screenshot.
type ScreenshotRequest struct {
	// URL is the target URL to capture (required, must start with http:// or https://)
	URL string `json:"url"`
	// Viewport configuration for custom dimensions
	Viewport *ViewportConfig `json:"viewport,omitempty"`
	// Device preset name (e.g., "Desktop HD", "iPhone 14", "iPad")
	Device string `json:"device,omitempty"`
	// Format of the output image: png, jpeg, jpg, webp, or pdf
	Format string `json:"format,omitempty"`
	// FullPage captures the entire scrollable page
	FullPage bool `json:"fullPage,omitempty"`
	// Quality of the output image (1-100, for jpeg/webp)
	Quality int `json:"quality,omitempty"`
	// Delay in milliseconds before capture (0-30000)
	Delay int `json:"delay,omitempty"`
	// WaitFor is a CSS selector to wait for before capture
	WaitFor string `json:"waitFor,omitempty"`
	// WaitUntil specifies when to consider navigation complete: load, domcontentloaded, networkidle
	WaitUntil string `json:"waitUntil,omitempty"`
	// Timeout in milliseconds (1000-60000)
	Timeout int `json:"timeout,omitempty"`
	// DarkMode enables dark mode for the capture
	DarkMode bool `json:"darkMode,omitempty"`
	// CustomCSS to inject into the page (max 10000 chars)
	CustomCSS string `json:"customCss,omitempty"`
	// HideSelectors is a list of CSS selectors to hide (max 50)
	HideSelectors []string `json:"hideSelectors,omitempty"`
	// Selector targets a specific element to capture (max 500 chars)
	Selector string `json:"selector,omitempty"`
	// BlockAds enables ad blocking
	BlockAds bool `json:"blockAds,omitempty"`
	// BlockCookieBanners enables cookie banner blocking
	BlockCookieBanners bool `json:"blockCookieBanners,omitempty"`
	// BlockLevel sets the blocking level: none, light, normal, pro, pro_plus, ultimate
	BlockLevel string `json:"blockLevel,omitempty"`
	// WebhookURL for async notification
	WebhookURL string `json:"webhookUrl,omitempty"`
	// WebhookSecret for webhook authentication (max 255 chars)
	WebhookSecret string `json:"webhookSecret,omitempty"`
	// ResponseType specifies the response format: BINARY or JSON
	ResponseType string `json:"responseType,omitempty"`
}

// JobStatus represents the status of an async job.
type JobStatus string

const (
	JobStatusQueued     JobStatus = "QUEUED"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
	JobStatusCancelled  JobStatus = "CANCELLED"
)

// JobResponse represents the response for an async screenshot job.
type JobResponse struct {
	// ID is the unique job identifier
	ID string `json:"id"`
	// Status of the job
	Status JobStatus `json:"status"`
	// URL that was captured
	URL string `json:"url"`
	// ResultURL where the screenshot can be downloaded
	ResultURL string `json:"resultUrl,omitempty"`
	// ErrorCode if the job failed
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorMessage if the job failed
	ErrorMessage string `json:"errorMessage,omitempty"`
	// CreatedAt timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// StartedAt timestamp
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// CompletedAt timestamp
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// ExpiresAt timestamp
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Metadata contains additional job information
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// AsyncJobCreatedResponse represents the response when creating an async job.
type AsyncJobCreatedResponse struct {
	// ID is the unique job identifier
	ID string `json:"id"`
	// Status of the job
	Status JobStatus `json:"status"`
	// StatusURL where the job status can be polled
	StatusURL string `json:"statusUrl"`
	// CreatedAt timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// BulkURLRequest represents a single URL in a bulk request.
type BulkURLRequest struct {
	// URL to capture (required)
	URL string `json:"url"`
	// Options for this specific URL
	Options *BulkURLOptions `json:"options,omitempty"`
}

// BulkURLOptions represents options for a single URL in a bulk request.
type BulkURLOptions struct {
	Viewport           *ViewportConfig `json:"viewport,omitempty"`
	Device             string          `json:"device,omitempty"`
	Format             string          `json:"format,omitempty"`
	FullPage           bool            `json:"fullPage,omitempty"`
	Quality            int             `json:"quality,omitempty"`
	Delay              int             `json:"delay,omitempty"`
	WaitFor            string          `json:"waitFor,omitempty"`
	WaitUntil          string          `json:"waitUntil,omitempty"`
	Timeout            int             `json:"timeout,omitempty"`
	DarkMode           bool            `json:"darkMode,omitempty"`
	CustomCSS          string          `json:"customCss,omitempty"`
	HideSelectors      []string        `json:"hideSelectors,omitempty"`
	Selector           string          `json:"selector,omitempty"`
	BlockAds           bool            `json:"blockAds,omitempty"`
	BlockCookieBanners bool            `json:"blockCookieBanners,omitempty"`
	BlockLevel         string          `json:"blockLevel,omitempty"`
}

// BulkDefaults represents default options for bulk screenshot requests.
type BulkDefaults struct {
	Viewport           *ViewportConfig `json:"viewport,omitempty"`
	Device             string          `json:"device,omitempty"`
	Format             string          `json:"format,omitempty"`
	FullPage           bool            `json:"fullPage,omitempty"`
	Quality            int             `json:"quality,omitempty"`
	Delay              int             `json:"delay,omitempty"`
	WaitFor            string          `json:"waitFor,omitempty"`
	WaitUntil          string          `json:"waitUntil,omitempty"`
	Timeout            int             `json:"timeout,omitempty"`
	DarkMode           bool            `json:"darkMode,omitempty"`
	CustomCSS          string          `json:"customCss,omitempty"`
	BlockAds           bool            `json:"blockAds,omitempty"`
	BlockCookieBanners bool            `json:"blockCookieBanners,omitempty"`
	BlockLevel         string          `json:"blockLevel,omitempty"`
}

// BulkRequest represents a request to capture multiple screenshots.
type BulkRequest struct {
	// URLs to capture (required, max 100)
	URLs []BulkURLRequest `json:"urls"`
	// Defaults to apply to all URLs
	Defaults *BulkDefaults `json:"defaults,omitempty"`
	// WebhookURL for notifications
	WebhookURL string `json:"webhookUrl,omitempty"`
	// WebhookSecret for webhook authentication
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// BulkJobInfo represents info about a single job in a bulk request.
type BulkJobInfo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	ResultURL string `json:"resultUrl,omitempty"`
}

// BulkResponse represents the response from creating a bulk job.
type BulkResponse struct {
	ID            string        `json:"id"`
	Status        string        `json:"status"`
	TotalJobs     int           `json:"totalJobs"`
	CompletedJobs int           `json:"completedJobs"`
	FailedJobs    int           `json:"failedJobs"`
	Progress      int           `json:"progress"`
	Jobs          []BulkJobInfo `json:"jobs,omitempty"`
	CreatedAt     *time.Time    `json:"createdAt,omitempty"`
	CompletedAt   *time.Time    `json:"completedAt,omitempty"`
}

// BulkJobSummary represents a summary of a bulk job.
type BulkJobSummary struct {
	ID            string     `json:"id"`
	Status        string     `json:"status"`
	TotalJobs     int        `json:"totalJobs"`
	CompletedJobs int        `json:"completedJobs"`
	FailedJobs    int        `json:"failedJobs"`
	Progress      int        `json:"progress"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	CompletedAt   *time.Time `json:"completedAt,omitempty"`
}

// BulkJobDetailInfo represents detailed info about a single job in a bulk request.
type BulkJobDetailInfo struct {
	ID           string     `json:"id"`
	URL          string     `json:"url"`
	Status       string     `json:"status"`
	ResultURL    string     `json:"resultUrl,omitempty"`
	StorageURL   string     `json:"storageUrl,omitempty"`
	Format       string     `json:"format,omitempty"`
	Width        int        `json:"width,omitempty"`
	Height       int        `json:"height,omitempty"`
	FileSize     int64      `json:"fileSize,omitempty"`
	RenderTimeMs int64      `json:"renderTimeMs,omitempty"`
	ErrorCode    string     `json:"errorCode,omitempty"`
	ErrorMessage string     `json:"errorMessage,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	CompletedAt  *time.Time `json:"completedAt,omitempty"`
}

// BulkStatusResponse represents the status of a bulk job with details.
type BulkStatusResponse struct {
	ID            string              `json:"id"`
	Status        string              `json:"status"`
	TotalJobs     int                 `json:"totalJobs"`
	CompletedJobs int                 `json:"completedJobs"`
	FailedJobs    int                 `json:"failedJobs"`
	Progress      int                 `json:"progress"`
	Jobs          []BulkJobDetailInfo `json:"jobs,omitempty"`
	CreatedAt     *time.Time          `json:"createdAt,omitempty"`
	CompletedAt   *time.Time          `json:"completedAt,omitempty"`
}

// CaptureItem represents a single capture in a compose request.
type CaptureItem struct {
	URL      string          `json:"url"`
	ID       string          `json:"id,omitempty"`
	Label    string          `json:"label,omitempty"`
	Viewport *ViewportConfig `json:"viewport,omitempty"`
	Device   string          `json:"device,omitempty"`
	FullPage bool            `json:"fullPage,omitempty"`
	DarkMode bool            `json:"darkMode,omitempty"`
	Delay    int             `json:"delay,omitempty"`
}

// VariantConfig represents a variant configuration for compose.
type VariantConfig struct {
	ID        string          `json:"id,omitempty"`
	Label     string          `json:"label,omitempty"`
	Viewport  *ViewportConfig `json:"viewport,omitempty"`
	Device    string          `json:"device,omitempty"`
	FullPage  bool            `json:"fullPage,omitempty"`
	DarkMode  bool            `json:"darkMode,omitempty"`
	Delay     int             `json:"delay,omitempty"`
	CustomCSS string          `json:"customCss,omitempty"`
}

// CaptureDefaults represents default capture options for compose.
type CaptureDefaults struct {
	Viewport           *ViewportConfig `json:"viewport,omitempty"`
	Device             string          `json:"device,omitempty"`
	Format             string          `json:"format,omitempty"`
	FullPage           bool            `json:"fullPage,omitempty"`
	Quality            int             `json:"quality,omitempty"`
	Delay              int             `json:"delay,omitempty"`
	WaitFor            string          `json:"waitFor,omitempty"`
	WaitUntil          string          `json:"waitUntil,omitempty"`
	Timeout            int             `json:"timeout,omitempty"`
	DarkMode           bool            `json:"darkMode,omitempty"`
	CustomCSS          string          `json:"customCss,omitempty"`
	HideSelectors      []string        `json:"hideSelectors,omitempty"`
	BlockAds           bool            `json:"blockAds,omitempty"`
	BlockCookieBanners bool            `json:"blockCookieBanners,omitempty"`
	BlockLevel         string          `json:"blockLevel,omitempty"`
}

// LabelConfig represents label styling for compose output.
type LabelConfig struct {
	Show            bool   `json:"show,omitempty"`
	Position        string `json:"position,omitempty"`
	FontSize        int    `json:"fontSize,omitempty"`
	FontColor       string `json:"fontColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Padding         int    `json:"padding,omitempty"`
}

// BorderConfig represents border styling for compose output.
type BorderConfig struct {
	Width  int    `json:"width,omitempty"`
	Color  string `json:"color,omitempty"`
	Radius int    `json:"radius,omitempty"`
}

// ShadowConfig represents shadow styling for compose output.
type ShadowConfig struct {
	Enabled bool   `json:"enabled,omitempty"`
	OffsetX int    `json:"offsetX,omitempty"`
	OffsetY int    `json:"offsetY,omitempty"`
	Blur    int    `json:"blur,omitempty"`
	Color   string `json:"color,omitempty"`
}

// ComposeOutputConfig represents output configuration for compose.
type ComposeOutputConfig struct {
	// Layout type: GRID, HORIZONTAL, VERTICAL, MASONRY, MONDRIAN, PARTITIONING, AUTO
	Layout string `json:"layout,omitempty"`
	// Format of output: png, jpeg, jpg, webp
	Format string `json:"format,omitempty"`
	// Quality of output (1-100)
	Quality int `json:"quality,omitempty"`
	// Columns for grid layout (1-10)
	Columns int `json:"columns,omitempty"`
	// Spacing between images (0-100)
	Spacing int `json:"spacing,omitempty"`
	// Padding around the canvas (0-100)
	Padding int `json:"padding,omitempty"`
	// Background color (#RRGGBB, #RRGGBBAA, or "transparent")
	Background string `json:"background,omitempty"`
	// Alignment: top, center, bottom
	Alignment string `json:"alignment,omitempty"`
	// MaxWidth of the output (100-10000)
	MaxWidth int `json:"maxWidth,omitempty"`
	// MaxHeight of the output (100-10000)
	MaxHeight int `json:"maxHeight,omitempty"`
	// ThumbnailWidth for thumbnails (50-2000)
	ThumbnailWidth int `json:"thumbnailWidth,omitempty"`
	// Labels configuration
	Labels *LabelConfig `json:"labels,omitempty"`
	// Border configuration
	Border *BorderConfig `json:"border,omitempty"`
	// Shadow configuration
	Shadow *ShadowConfig `json:"shadow,omitempty"`
}

// ComposeRequest represents a request to compose multiple screenshots.
type ComposeRequest struct {
	// Captures is a list of URLs to capture (max 20)
	Captures []CaptureItem `json:"captures,omitempty"`
	// URL for single URL with variants
	URL string `json:"url,omitempty"`
	// Variants for the single URL (max 20)
	Variants []VariantConfig `json:"variants,omitempty"`
	// Defaults for all captures
	Defaults *CaptureDefaults `json:"defaults,omitempty"`
	// Output configuration
	Output *ComposeOutputConfig `json:"output,omitempty"`
	// Async mode
	Async bool `json:"async,omitempty"`
	// WebhookURL for notifications
	WebhookURL string `json:"webhookUrl,omitempty"`
	// WebhookSecret for authentication
	WebhookSecret string `json:"webhookSecret,omitempty"`
	// CapturesMode indicates multiple captures mode
	CapturesMode bool `json:"capturesMode,omitempty"`
	// VariantsMode indicates variants mode
	VariantsMode bool `json:"variantsMode,omitempty"`
}

// ComposeMetadata represents metadata for a composed image.
type ComposeMetadata struct {
	CaptureCount int                    `json:"captureCount,omitempty"`
	Layout       string                 `json:"layout,omitempty"`
	Additional   map[string]interface{} `json:"additional,omitempty"`
}

// ComposeResponse represents the response from a compose request.
type ComposeResponse struct {
	URL          string           `json:"url"`
	StorageURL   string           `json:"storageUrl,omitempty"`
	ExpiresAt    *time.Time       `json:"expiresAt,omitempty"`
	Width        int              `json:"width"`
	Height       int              `json:"height"`
	Format       string           `json:"format"`
	FileSize     int64            `json:"fileSize"`
	RenderTimeMs int64            `json:"renderTimeMs"`
	Layout       string           `json:"layout"`
	Metadata     *ComposeMetadata `json:"metadata,omitempty"`
}

// ComposeJobStatusResponse represents the status of a compose job.
type ComposeJobStatusResponse struct {
	JobID             string           `json:"jobId"`
	Status            string           `json:"status"`
	Progress          int              `json:"progress"`
	TotalCaptures     int              `json:"totalCaptures"`
	CompletedCaptures int              `json:"completedCaptures"`
	Result            *ComposeResponse `json:"result,omitempty"`
	ErrorCode         string           `json:"errorCode,omitempty"`
	ErrorMessage      string           `json:"errorMessage,omitempty"`
	CreatedAt         *time.Time       `json:"createdAt,omitempty"`
	CompletedAt       *time.Time       `json:"completedAt,omitempty"`
}

// ComposeJobSummaryResponse represents a summary of a compose job.
type ComposeJobSummaryResponse struct {
	JobID             string     `json:"jobId"`
	Status            string     `json:"status"`
	TotalCaptures     int        `json:"totalCaptures"`
	CompletedCaptures int        `json:"completedCaptures"`
	FailedCaptures    int        `json:"failedCaptures"`
	Progress          int        `json:"progress"`
	LayoutType        string     `json:"layoutType,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	CompletedAt       *time.Time `json:"completedAt,omitempty"`
}

// PlacementPreview represents a placement preview for compose.
type PlacementPreview struct {
	Index  int    `json:"index"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Label  string `json:"label,omitempty"`
}

// LayoutPreviewResponse represents a layout preview response.
type LayoutPreviewResponse struct {
	Layout         string                 `json:"layout"`
	ResolvedLayout string                 `json:"resolvedLayout,omitempty"`
	CanvasWidth    int                    `json:"canvasWidth"`
	CanvasHeight   int                    `json:"canvasHeight"`
	Placements     []PlacementPreview     `json:"placements"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// ScheduleScreenshotOptions represents options for scheduled screenshots.
type ScheduleScreenshotOptions struct {
	Viewport           *ViewportConfig `json:"viewport,omitempty"`
	Device             string          `json:"device,omitempty"`
	Format             string          `json:"format,omitempty"`
	FullPage           bool            `json:"fullPage,omitempty"`
	Quality            int             `json:"quality,omitempty"`
	Delay              int             `json:"delay,omitempty"`
	WaitFor            string          `json:"waitFor,omitempty"`
	WaitUntil          string          `json:"waitUntil,omitempty"`
	Timeout            int             `json:"timeout,omitempty"`
	DarkMode           bool            `json:"darkMode,omitempty"`
	CustomCSS          string          `json:"customCss,omitempty"`
	HideSelectors      []string        `json:"hideSelectors,omitempty"`
	BlockAds           bool            `json:"blockAds,omitempty"`
	BlockCookieBanners bool            `json:"blockCookieBanners,omitempty"`
	BlockLevel         string          `json:"blockLevel,omitempty"`
}

// CreateScheduleRequest represents a request to create a schedule.
type CreateScheduleRequest struct {
	// Name of the schedule (required, max 255)
	Name string `json:"name"`
	// URL to capture (required)
	URL string `json:"url"`
	// Schedule is a cron expression (required)
	Schedule string `json:"schedule"`
	// Timezone for the schedule
	Timezone string `json:"timezone,omitempty"`
	// Options for the screenshot
	Options *ScheduleScreenshotOptions `json:"options,omitempty"`
	// WebhookURL for notifications
	WebhookURL string `json:"webhookUrl,omitempty"`
	// WebhookSecret for authentication
	WebhookSecret string `json:"webhookSecret,omitempty"`
	// RetentionDays (1-365)
	RetentionDays int `json:"retentionDays,omitempty"`
	// StartsAt timestamp
	StartsAt *time.Time `json:"startsAt,omitempty"`
	// EndsAt timestamp
	EndsAt *time.Time `json:"endsAt,omitempty"`
}

// UpdateScheduleRequest represents a request to update a schedule.
type UpdateScheduleRequest struct {
	Name          string                     `json:"name,omitempty"`
	URL           string                     `json:"url,omitempty"`
	Schedule      string                     `json:"schedule,omitempty"`
	Timezone      string                     `json:"timezone,omitempty"`
	Options       *ScheduleScreenshotOptions `json:"options,omitempty"`
	WebhookURL    string                     `json:"webhookUrl,omitempty"`
	WebhookSecret string                     `json:"webhookSecret,omitempty"`
	RetentionDays int                        `json:"retentionDays,omitempty"`
	StartsAt      *time.Time                 `json:"startsAt,omitempty"`
	EndsAt        *time.Time                 `json:"endsAt,omitempty"`
}

// ScheduleResponse represents a schedule.
type ScheduleResponse struct {
	ID                  string                 `json:"id"`
	Name                string                 `json:"name"`
	URL                 string                 `json:"url"`
	Schedule            string                 `json:"schedule"`
	ScheduleDescription string                 `json:"scheduleDescription,omitempty"`
	Timezone            string                 `json:"timezone,omitempty"`
	Status              string                 `json:"status"`
	Options             map[string]interface{} `json:"options,omitempty"`
	WebhookURL          string                 `json:"webhookUrl,omitempty"`
	RetentionDays       int                    `json:"retentionDays,omitempty"`
	StartsAt            *time.Time             `json:"startsAt,omitempty"`
	EndsAt              *time.Time             `json:"endsAt,omitempty"`
	LastExecutedAt      *time.Time             `json:"lastExecutedAt,omitempty"`
	NextExecutionAt     *time.Time             `json:"nextExecutionAt,omitempty"`
	ExecutionCount      int                    `json:"executionCount"`
	SuccessCount        int                    `json:"successCount"`
	FailureCount        int                    `json:"failureCount"`
	CreatedAt           *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time             `json:"updatedAt,omitempty"`
}

// ScheduleListResponse represents a list of schedules.
type ScheduleListResponse struct {
	Schedules []ScheduleResponse `json:"schedules"`
	Total     int                `json:"total"`
}

// ScheduleExecutionResponse represents a schedule execution.
type ScheduleExecutionResponse struct {
	ID           string     `json:"id"`
	ExecutedAt   *time.Time `json:"executedAt,omitempty"`
	Status       string     `json:"status"`
	ResultURL    string     `json:"resultUrl,omitempty"`
	StorageURL   string     `json:"storageUrl,omitempty"`
	FileSize     int64      `json:"fileSize,omitempty"`
	RenderTimeMs int64      `json:"renderTimeMs,omitempty"`
	ErrorCode    string     `json:"errorCode,omitempty"`
	ErrorMessage string     `json:"errorMessage,omitempty"`
	ExpiresAt    *time.Time `json:"expiresAt,omitempty"`
}

// ScheduleHistoryResponse represents schedule execution history.
type ScheduleHistoryResponse struct {
	ScheduleID      string                      `json:"scheduleId"`
	TotalExecutions int64                       `json:"totalExecutions"`
	Executions      []ScheduleExecutionResponse `json:"executions"`
}

// QuotaDetailResponse represents quota details.
type QuotaDetailResponse struct {
	Limit       int `json:"limit"`
	Used        int `json:"used"`
	Remaining   int `json:"remaining"`
	PercentUsed int `json:"percentUsed"`
}

// BandwidthQuotaResponse represents bandwidth quota details.
type BandwidthQuotaResponse struct {
	LimitBytes         int64  `json:"limitBytes"`
	LimitFormatted     string `json:"limitFormatted"`
	UsedBytes          int64  `json:"usedBytes"`
	UsedFormatted      string `json:"usedFormatted"`
	RemainingBytes     int64  `json:"remainingBytes"`
	RemainingFormatted string `json:"remainingFormatted"`
	PercentUsed        int    `json:"percentUsed"`
}

// QuotaResponse represents quota information.
type QuotaResponse struct {
	Screenshots *QuotaDetailResponse    `json:"screenshots,omitempty"`
	Bandwidth   *BandwidthQuotaResponse `json:"bandwidth,omitempty"`
}

// PeriodUsageResponse represents usage for a period.
type PeriodUsageResponse struct {
	PeriodStart        string `json:"periodStart"`
	PeriodEnd          string `json:"periodEnd"`
	ScreenshotsCount   int    `json:"screenshotsCount"`
	BandwidthBytes     int64  `json:"bandwidthBytes"`
	BandwidthFormatted string `json:"bandwidthFormatted"`
}

// TotalsResponse represents total usage.
type TotalsResponse struct {
	ScreenshotsCount   int64  `json:"screenshotsCount"`
	BandwidthBytes     int64  `json:"bandwidthBytes"`
	BandwidthFormatted string `json:"bandwidthFormatted"`
}

// UsageResponse represents usage statistics.
type UsageResponse struct {
	Tier          string                `json:"tier"`
	CurrentPeriod *PeriodUsageResponse  `json:"currentPeriod,omitempty"`
	Quota         *QuotaResponse        `json:"quota,omitempty"`
	History       []PeriodUsageResponse `json:"history,omitempty"`
	Totals        *TotalsResponse       `json:"totals,omitempty"`
}

// QuotaStatusResponse represents quota status.
type QuotaStatusResponse struct {
	Tier        string                  `json:"tier"`
	Screenshots *QuotaDetailResponse    `json:"screenshots,omitempty"`
	Bandwidth   *BandwidthQuotaResponse `json:"bandwidth,omitempty"`
	PeriodEnds  string                  `json:"periodEnds,omitempty"`
}
