package allscreenshots

import (
	"fmt"
)

// APIError represents an error returned by the Allscreenshots API.
type APIError struct {
	// StatusCode is the HTTP status code returned by the API
	StatusCode int
	// Code is the error code from the API response
	Code string
	// Message is the human-readable error message
	Message string
	// Details contains additional error information
	Details map[string]interface{}
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("allscreenshots: API error %d (%s): %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("allscreenshots: API error %d: %s", e.StatusCode, e.Message)
}

// IsAPIError checks if an error is an APIError.
func IsAPIError(err error) bool {
	_, ok := err.(*APIError)
	return ok
}

// AsAPIError converts an error to an APIError if possible.
func AsAPIError(err error) (*APIError, bool) {
	apiErr, ok := err.(*APIError)
	return apiErr, ok
}

// ValidationError represents a validation error.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("allscreenshots: validation error for field '%s': %s", e.Field, e.Message)
}

// IsValidationError checks if an error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// NetworkError represents a network-related error.
type NetworkError struct {
	Message string
	Cause   error
}

// Error implements the error interface.
func (e *NetworkError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("allscreenshots: network error: %s: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("allscreenshots: network error: %s", e.Message)
}

// Unwrap returns the underlying cause.
func (e *NetworkError) Unwrap() error {
	return e.Cause
}

// IsNetworkError checks if an error is a NetworkError.
func IsNetworkError(err error) bool {
	_, ok := err.(*NetworkError)
	return ok
}

// TimeoutError represents a timeout error.
type TimeoutError struct {
	Message string
	Cause   error
}

// Error implements the error interface.
func (e *TimeoutError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("allscreenshots: timeout: %s: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("allscreenshots: timeout: %s", e.Message)
}

// Unwrap returns the underlying cause.
func (e *TimeoutError) Unwrap() error {
	return e.Cause
}

// IsTimeoutError checks if an error is a TimeoutError.
func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}

// RetryError represents an error that occurred after all retries were exhausted.
type RetryError struct {
	Attempts int
	LastErr  error
}

// Error implements the error interface.
func (e *RetryError) Error() string {
	return fmt.Sprintf("allscreenshots: failed after %d attempts: %v", e.Attempts, e.LastErr)
}

// Unwrap returns the last error that occurred.
func (e *RetryError) Unwrap() error {
	return e.LastErr
}

// IsRetryError checks if an error is a RetryError.
func IsRetryError(err error) bool {
	_, ok := err.(*RetryError)
	return ok
}

// Common error codes returned by the API.
const (
	ErrCodeInvalidURL         = "INVALID_URL"
	ErrCodeInvalidDevice      = "INVALID_DEVICE"
	ErrCodeInvalidFormat      = "INVALID_FORMAT"
	ErrCodeInvalidViewport    = "INVALID_VIEWPORT"
	ErrCodeInvalidTimeout     = "INVALID_TIMEOUT"
	ErrCodeInvalidDelay       = "INVALID_DELAY"
	ErrCodeInvalidQuality     = "INVALID_QUALITY"
	ErrCodeURLUnreachable     = "URL_UNREACHABLE"
	ErrCodeTimeout            = "TIMEOUT"
	ErrCodeRateLimitExceeded  = "RATE_LIMIT_EXCEEDED"
	ErrCodeQuotaExceeded      = "QUOTA_EXCEEDED"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
)

// IsBadRequest returns true if the error is a 400 Bad Request error.
func IsBadRequest(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode == 400
	}
	return false
}

// IsUnauthorized returns true if the error is a 401 Unauthorized error.
func IsUnauthorized(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode == 401
	}
	return false
}

// IsForbidden returns true if the error is a 403 Forbidden error.
func IsForbidden(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode == 403
	}
	return false
}

// IsNotFound returns true if the error is a 404 Not Found error.
func IsNotFound(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode == 404
	}
	return false
}

// IsRateLimited returns true if the error is a 429 Too Many Requests error.
func IsRateLimited(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode == 429
	}
	return false
}

// IsServerError returns true if the error is a 5xx server error.
func IsServerError(err error) bool {
	if apiErr, ok := AsAPIError(err); ok {
		return apiErr.StatusCode >= 500 && apiErr.StatusCode < 600
	}
	return false
}
