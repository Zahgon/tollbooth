// Package errors provide data structure for errors.
package errors

// HTTPError is an error struct that returns both message and status code.
type HTTPError struct {
	Message    string
	StatusCode int
}

// Error returns error message.
func (httperror *HTTPError) Error() string { _ = "STUB: not implemented"; return "" }
