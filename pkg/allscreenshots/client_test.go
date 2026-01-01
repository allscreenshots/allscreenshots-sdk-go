package allscreenshots

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Run("creates client with defaults", func(t *testing.T) {
		client := NewClient(WithAPIKey("test-key"))
		assert.NotNil(t, client)
		assert.Equal(t, DefaultBaseURL, client.baseURL)
		assert.Equal(t, "test-key", client.apiKey)
		assert.Equal(t, DefaultMaxRetries, client.maxRetries)
	})

	t.Run("creates client with custom options", func(t *testing.T) {
		client := NewClient(
			WithAPIKey("custom-key"),
			WithBaseURL("https://custom.api.com"),
			WithTimeout(60*time.Second),
			WithMaxRetries(5),
			WithRetryWait(2*time.Second, 60*time.Second),
		)
		assert.Equal(t, "https://custom.api.com", client.baseURL)
		assert.Equal(t, "custom-key", client.apiKey)
		assert.Equal(t, 5, client.maxRetries)
		assert.Equal(t, 2*time.Second, client.retryWaitMin)
		assert.Equal(t, 60*time.Second, client.retryWaitMax)
	})

	t.Run("trims trailing slash from base URL", func(t *testing.T) {
		client := NewClient(WithBaseURL("https://api.example.com/"))
		assert.Equal(t, "https://api.example.com", client.baseURL)
	})
}

func TestScreenshotRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		req     *ScreenshotRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name:    "empty URL",
			req:     &ScreenshotRequest{},
			wantErr: "URL is required",
		},
		{
			name:    "invalid URL scheme",
			req:     &ScreenshotRequest{URL: "ftp://example.com"},
			wantErr: "URL must start with http:// or https://",
		},
		{
			name:    "quality too low",
			req:     &ScreenshotRequest{URL: "https://example.com", Quality: 0},
			wantErr: "",
		},
		{
			name:    "quality too high",
			req:     &ScreenshotRequest{URL: "https://example.com", Quality: 101},
			wantErr: "quality must be between 1 and 100",
		},
		{
			name:    "delay too high",
			req:     &ScreenshotRequest{URL: "https://example.com", Delay: 30001},
			wantErr: "delay must be between 0 and 30000",
		},
		{
			name:    "timeout too low",
			req:     &ScreenshotRequest{URL: "https://example.com", Timeout: 500},
			wantErr: "timeout must be between 1000 and 60000",
		},
		{
			name:    "valid request",
			req:     &ScreenshotRequest{URL: "https://example.com", Device: "Desktop HD"},
			wantErr: "",
		},
		{
			name: "valid viewport",
			req: &ScreenshotRequest{
				URL:      "https://example.com",
				Viewport: &ViewportConfig{Width: 1920, Height: 1080},
			},
			wantErr: "",
		},
		{
			name: "invalid viewport width",
			req: &ScreenshotRequest{
				URL:      "https://example.com",
				Viewport: &ViewportConfig{Width: 50},
			},
			wantErr: "width must be between 100 and 4096",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateScreenshotRequest(tt.req)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestBulkRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		req     *BulkRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name:    "empty URLs",
			req:     &BulkRequest{},
			wantErr: "at least one URL is required",
		},
		{
			name: "invalid URL in list",
			req: &BulkRequest{
				URLs: []BulkURLRequest{{URL: "not-a-url"}},
			},
			wantErr: "URL must start with http:// or https://",
		},
		{
			name: "valid request",
			req: &BulkRequest{
				URLs: []BulkURLRequest{
					{URL: "https://example.com"},
					{URL: "https://github.com"},
				},
			},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBulkRequest(tt.req)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestComposeRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		req     *ComposeRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name:    "no captures or URL",
			req:     &ComposeRequest{},
			wantErr: "either captures or url is required",
		},
		{
			name: "valid with URL",
			req: &ComposeRequest{
				URL: "https://example.com",
			},
			wantErr: "",
		},
		{
			name: "valid with captures",
			req: &ComposeRequest{
				Captures: []CaptureItem{{URL: "https://example.com"}},
			},
			wantErr: "",
		},
		{
			name: "invalid capture URL",
			req: &ComposeRequest{
				Captures: []CaptureItem{{URL: "not-valid"}},
			},
			wantErr: "URL must start with http:// or https://",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateComposeRequest(tt.req)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestCreateScheduleRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		req     *CreateScheduleRequest
		wantErr string
	}{
		{
			name:    "nil request",
			req:     nil,
			wantErr: "request cannot be nil",
		},
		{
			name:    "missing name",
			req:     &CreateScheduleRequest{URL: "https://example.com", Schedule: "0 9 * * *"},
			wantErr: "name is required",
		},
		{
			name:    "missing URL",
			req:     &CreateScheduleRequest{Name: "Test", Schedule: "0 9 * * *"},
			wantErr: "URL is required",
		},
		{
			name:    "missing schedule",
			req:     &CreateScheduleRequest{Name: "Test", URL: "https://example.com"},
			wantErr: "schedule is required",
		},
		{
			name:    "invalid retention days",
			req:     &CreateScheduleRequest{Name: "Test", URL: "https://example.com", Schedule: "0 9 * * *", RetentionDays: 400},
			wantErr: "retentionDays must be between 1 and 365",
		},
		{
			name:    "valid request",
			req:     &CreateScheduleRequest{Name: "Test", URL: "https://example.com", Schedule: "0 9 * * *"},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateScheduleRequest(tt.req)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			}
		})
	}
}

func TestClient_Screenshot(t *testing.T) {
	imageData := []byte{0x89, 0x50, 0x4E, 0x47} // PNG magic bytes

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/screenshots", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "test-api-key", r.Header.Get("X-API-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		var req ScreenshotRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		assert.Equal(t, "https://example.com", req.URL)
		assert.Equal(t, "Desktop HD", req.Device)

		w.WriteHeader(http.StatusOK)
		w.Write(imageData)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-api-key"),
		WithBaseURL(server.URL),
	)

	result, err := client.Screenshot(context.Background(), &ScreenshotRequest{
		URL:    "https://example.com",
		Device: "Desktop HD",
	})

	require.NoError(t, err)
	assert.Equal(t, imageData, result)
}

func TestClient_ScreenshotAsync(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/screenshots/async", r.URL.Path)

		resp := AsyncJobCreatedResponse{
			ID:        "job-123",
			Status:    JobStatusQueued,
			StatusURL: "https://api.example.com/v1/screenshots/jobs/job-123",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-api-key"),
		WithBaseURL(server.URL),
	)

	result, err := client.ScreenshotAsync(context.Background(), &ScreenshotRequest{
		URL:    "https://example.com",
		Device: "Desktop HD",
	})

	require.NoError(t, err)
	assert.Equal(t, "job-123", result.ID)
	assert.Equal(t, JobStatusQueued, result.Status)
}

func TestClient_GetJob(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/screenshots/jobs/job-123", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		resp := JobResponse{
			ID:     "job-123",
			Status: JobStatusCompleted,
			URL:    "https://example.com",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-api-key"),
		WithBaseURL(server.URL),
	)

	result, err := client.GetJob(context.Background(), "job-123")

	require.NoError(t, err)
	assert.Equal(t, "job-123", result.ID)
	assert.Equal(t, JobStatusCompleted, result.Status)
}

func TestClient_ErrorHandling(t *testing.T) {
	t.Run("handles 400 error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    "INVALID_URL",
				"message": "Invalid URL provided",
			})
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		_, err := client.Screenshot(context.Background(), &ScreenshotRequest{
			URL: "https://example.com",
		})

		require.Error(t, err)
		apiErr, ok := AsAPIError(err)
		require.True(t, ok)
		assert.Equal(t, 400, apiErr.StatusCode)
		assert.Equal(t, "INVALID_URL", apiErr.Code)
		assert.Equal(t, "Invalid URL provided", apiErr.Message)
		assert.True(t, IsBadRequest(err))
	})

	t.Run("handles 401 error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("invalid-key"),
			WithBaseURL(server.URL),
		)

		_, err := client.Screenshot(context.Background(), &ScreenshotRequest{
			URL: "https://example.com",
		})

		require.Error(t, err)
		assert.True(t, IsUnauthorized(err))
	})

	t.Run("handles 404 error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		_, err := client.GetJob(context.Background(), "nonexistent")

		require.Error(t, err)
		assert.True(t, IsNotFound(err))
	})

	t.Run("handles 429 with retries", func(t *testing.T) {
		attempts := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attempts++
			if attempts < 3 {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{0x89, 0x50, 0x4E, 0x47})
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
			WithRetryWait(1*time.Millisecond, 10*time.Millisecond),
		)

		result, err := client.Screenshot(context.Background(), &ScreenshotRequest{
			URL: "https://example.com",
		})

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, attempts)
	})

	t.Run("fails after max retries", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
			WithMaxRetries(2),
			WithRetryWait(1*time.Millisecond, 10*time.Millisecond),
		)

		_, err := client.Screenshot(context.Background(), &ScreenshotRequest{
			URL: "https://example.com",
		})

		require.Error(t, err)
		assert.True(t, IsRetryError(err))
	})

	t.Run("requires API key", func(t *testing.T) {
		client := NewClient(
			WithBaseURL("https://api.example.com"),
		)
		client.apiKey = "" // Override any env var

		_, err := client.Screenshot(context.Background(), &ScreenshotRequest{
			URL: "https://example.com",
		})

		require.Error(t, err)
		assert.True(t, IsValidationError(err))
		assert.Contains(t, err.Error(), "API key is required")
	})
}

func TestClient_ListJobs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/screenshots/jobs", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		jobs := []JobResponse{
			{ID: "job-1", Status: JobStatusCompleted},
			{ID: "job-2", Status: JobStatusProcessing},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jobs)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-api-key"),
		WithBaseURL(server.URL),
	)

	result, err := client.ListJobs(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "job-1", result[0].ID)
}

func TestClient_CancelJob(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/screenshots/jobs/job-123/cancel", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		resp := JobResponse{
			ID:     "job-123",
			Status: JobStatusCancelled,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIKey("test-api-key"),
		WithBaseURL(server.URL),
	)

	result, err := client.CancelJob(context.Background(), "job-123")

	require.NoError(t, err)
	assert.Equal(t, JobStatusCancelled, result.Status)
}

func TestClient_BulkOperations(t *testing.T) {
	t.Run("CreateBulkJob", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/screenshots/bulk", r.URL.Path)
			assert.Equal(t, "POST", r.Method)

			resp := BulkResponse{
				ID:        "bulk-123",
				Status:    "PROCESSING",
				TotalJobs: 3,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		result, err := client.CreateBulkJob(context.Background(), &BulkRequest{
			URLs: []BulkURLRequest{
				{URL: "https://example.com"},
				{URL: "https://github.com"},
			},
		})

		require.NoError(t, err)
		assert.Equal(t, "bulk-123", result.ID)
		assert.Equal(t, 3, result.TotalJobs)
	})

	t.Run("ListBulkJobs", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jobs := []BulkJobSummary{
				{ID: "bulk-1", Status: "COMPLETED"},
				{ID: "bulk-2", Status: "PROCESSING"},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(jobs)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		result, err := client.ListBulkJobs(context.Background())

		require.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("GetBulkJob", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/screenshots/bulk/bulk-123", r.URL.Path)

			resp := BulkStatusResponse{
				ID:            "bulk-123",
				Status:        "COMPLETED",
				TotalJobs:     3,
				CompletedJobs: 3,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		result, err := client.GetBulkJob(context.Background(), "bulk-123")

		require.NoError(t, err)
		assert.Equal(t, "bulk-123", result.ID)
		assert.Equal(t, 3, result.CompletedJobs)
	})
}

func TestClient_Usage(t *testing.T) {
	t.Run("GetUsage", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/usage", r.URL.Path)

			resp := UsageResponse{
				Tier: "pro",
				CurrentPeriod: &PeriodUsageResponse{
					ScreenshotsCount: 150,
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		result, err := client.GetUsage(context.Background())

		require.NoError(t, err)
		assert.Equal(t, "pro", result.Tier)
		assert.Equal(t, 150, result.CurrentPeriod.ScreenshotsCount)
	})

	t.Run("GetQuotaStatus", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/v1/usage/quota", r.URL.Path)

			resp := QuotaStatusResponse{
				Tier: "pro",
				Screenshots: &QuotaDetailResponse{
					Limit:     1000,
					Used:      150,
					Remaining: 850,
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(
			WithAPIKey("test-api-key"),
			WithBaseURL(server.URL),
		)

		result, err := client.GetQuotaStatus(context.Background())

		require.NoError(t, err)
		assert.Equal(t, 850, result.Screenshots.Remaining)
	})
}

func TestCalculateBackoff(t *testing.T) {
	client := NewClient(
		WithRetryWait(1*time.Second, 30*time.Second),
	)

	// Test that backoff increases with attempts
	backoff1 := client.calculateBackoff(1)
	backoff2 := client.calculateBackoff(2)
	backoff3 := client.calculateBackoff(3)

	// Backoff should generally increase (allowing for jitter)
	assert.Less(t, backoff1, 2*time.Second)
	assert.LessOrEqual(t, backoff1, backoff2+500*time.Millisecond) // Account for jitter
	assert.LessOrEqual(t, backoff2, backoff3+1*time.Second)

	// Backoff should not exceed max
	backoff10 := client.calculateBackoff(10)
	assert.LessOrEqual(t, backoff10, 30*time.Second)
}

func TestErrorTypes(t *testing.T) {
	t.Run("APIError", func(t *testing.T) {
		err := &APIError{
			StatusCode: 400,
			Code:       "INVALID_URL",
			Message:    "Invalid URL",
		}
		assert.Contains(t, err.Error(), "400")
		assert.Contains(t, err.Error(), "INVALID_URL")
		assert.True(t, IsAPIError(err))
	})

	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{
			Field:   "url",
			Message: "URL is required",
		}
		assert.Contains(t, err.Error(), "url")
		assert.Contains(t, err.Error(), "URL is required")
		assert.True(t, IsValidationError(err))
	})

	t.Run("NetworkError", func(t *testing.T) {
		err := &NetworkError{
			Message: "connection refused",
		}
		assert.Contains(t, err.Error(), "connection refused")
		assert.True(t, IsNetworkError(err))
	})

	t.Run("RetryError", func(t *testing.T) {
		innerErr := &NetworkError{Message: "timeout"}
		err := &RetryError{
			Attempts: 3,
			LastErr:  innerErr,
		}
		assert.Contains(t, err.Error(), "3 attempts")
		assert.True(t, IsRetryError(err))
		assert.Equal(t, innerErr, err.Unwrap())
	})
}
