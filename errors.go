package lztapi

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("lztapi: HTTP %d: %s", e.StatusCode, e.Message)
}

type RateLimitError struct {
	RetryAfter float64
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("lztapi: rate limited, retry after %.1fs", e.RetryAfter)
}

type AuthError struct{}

func (e *AuthError) Error() string {
	return "lztapi: authentication failed (401)"
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "lztapi: resource not found (404)"
}
