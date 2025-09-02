package juicewrld

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.StatusCode == 0 {
		return e.Message
	}
	return fmt.Sprintf("api error: %d - %s", e.StatusCode, e.Message)
}

type RateLimitError struct{ APIError }
type NotFoundError struct{ APIError }
type AuthenticationError struct{ APIError }
type ValidationError struct{ APIError }
