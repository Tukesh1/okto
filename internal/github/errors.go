package github

import (
	"errors"
	"fmt"
	"time"
)

// HTTPError is a non-success GitHub response that is not a rate limit.
type HTTPError struct {
	StatusCode       int
	Message          string
	DocumentationURL string
	RequestID        string
}

func (e *HTTPError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("github: HTTP %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("github: HTTP %d", e.StatusCode)
}

// RateLimitError indicates a primary or secondary GitHub rate limit.
type RateLimitError struct {
	StatusCode int
	Message    string
	ResetAt    time.Time
	RetryAfter time.Duration
	Resource   string
	RequestID  string
}

func (e *RateLimitError) Error() string {
	msg := "github: rate limit exceeded"
	if e.Resource != "" {
		msg += " for " + e.Resource
	}
	if !e.ResetAt.IsZero() {
		msg += "; resets at " + e.ResetAt.UTC().Format(time.RFC3339)
	}
	if e.RetryAfter > 0 {
		msg += "; retry after " + e.RetryAfter.String()
	}
	if e.Message != "" {
		msg += ": " + e.Message
	}
	return msg
}

// StatsPendingError indicates that GitHub accepted a statistics request but
// did not finish computing it within the client's bounded retry window.
type StatsPendingError struct {
	Resource string
	Attempts int
}

func (e *StatsPendingError) Error() string {
	if e.Attempts > 0 {
		return fmt.Sprintf("github: %s statistics are still being computed after %d attempts", e.Resource, e.Attempts)
	}
	return fmt.Sprintf("github: %s statistics are still being computed", e.Resource)
}

// NetworkError wraps a transport failure after retries are exhausted.
type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string {
	return "github: network request failed: " + e.Err.Error()
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// ErrResponseTooLarge indicates that a response exceeded the client's limit.
var ErrResponseTooLarge = errors.New("github: response body too large")

// IsHTTPStatus reports whether err is an HTTPError with the given status.
func IsHTTPStatus(err error, status int) bool {
	var httpErr *HTTPError
	return errors.As(err, &httpErr) && httpErr.StatusCode == status
}

// IsNotFound reports whether err is a 404 response.
func IsNotFound(err error) bool {
	return IsHTTPStatus(err, 404)
}

// IsAuth reports whether err is a 401 response.
func IsAuth(err error) bool {
	return IsHTTPStatus(err, 401)
}

// IsRateLimit reports whether err is a rate-limit response.
func IsRateLimit(err error) bool {
	var rateErr *RateLimitError
	return errors.As(err, &rateErr)
}

// IsStatsPending reports whether err indicates pending statistics.
func IsStatsPending(err error) bool {
	var pendingErr *StatsPendingError
	return errors.As(err, &pendingErr)
}
