package fastgo

import "fmt"

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func HTTPErrorf(code int, format string, args ...any) *HTTPError {
	return &HTTPError{
		StatusCode: code,
		Message:    fmt.Sprintf(format, args...),
	}
}
