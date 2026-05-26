// Package tollbooth provides rate-limiting logic to HTTP request handler.
package tollbooth

import (
	"net/http"

	"github.com/didip/tollbooth/v8/errors"
	"github.com/didip/tollbooth/v8/limiter"
)

// setResponseHeaders configures X-Rate-Limit-Limit and X-Rate-Limit-Duration
func setResponseHeaders(lmt *limiter.Limiter, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// setRateLimitResponseHeaders configures RateLimit-Limit, RateLimit-Remaining and RateLimit-Reset
// as seen at https://datatracker.ietf.org/doc/html/draft-ietf-httpapi-ratelimit-headers
func setRateLimitResponseHeaders(lmt *limiter.Limiter, w http.ResponseWriter, tokensLeft int) {
	_ = "STUB: not implemented"
	return
}

// NewLimiter is a convenience function to limiter.New.
func NewLimiter(max float64, tbOptions *limiter.ExpirableOptions) *limiter.Limiter {
	_ = "STUB: not implemented"
	return nil
}

// LimitByKeys keeps track number of request made by keys separated by pipe.
// It returns HTTPError when limit is exceeded.
func LimitByKeys(lmt *limiter.Limiter, keys []string) *errors.HTTPError {
	_ = "STUB: not implemented"
	return nil
}

// LimitByKeysAndReturn keeps track number of request made by keys separated by pipe.
// It returns HTTPError when limit is exceeded, and also returns the current limit value.
func LimitByKeysAndReturn(lmt *limiter.Limiter, keys []string) (*errors.HTTPError, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// ShouldSkipLimiter is a series of filter that decides if request should be limited or not.
func ShouldSkipLimiter(lmt *limiter.Limiter, r *http.Request) bool {
	_ = "STUB: not implemented"
	// ---------------------------------
	// Filter by remote ip
	// If we are unable to find remoteIP, skip limiter
	return false
}

// ---------------------------------
// Filter by request method

// If request does not contain all of the methods in limiter,
// skip limiter

// ---------------------------------
// Filter by request headers

// If request does not contain all of the headers in limiter,
// skip limiter

// ------------------------------
// If request contains the header key but not the values,
// skip limiter

// ---------------------------------
// Filter by context values

// If request does not contain all of the contexts in limiter,
// skip limiter

// ------------------------------
// If request contains the context key but not the values,
// skip limiter

// ---------------------------------
// Filter by basic auth usernames

// If request does not contain all of the basic auth users in limiter,
// skip limiter

// BuildKeys generates a slice of keys to rate-limit by given limiter and request structs.
func BuildKeys(lmt *limiter.Limiter, r *http.Request) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// If header values are empty, rate-limit all request containing headerKey.

// If header values are not empty, rate-limit all request with headerKey and headerValues.

// If context values are empty, rate-limit all request containing contextKey.

// If context values are not empty, rate-limit all request with contextKey and contextValues.

// LimitByRequest builds keys based on http.Request struct,
// loops through all the keys, and check if any one of them returns HTTPError.
func LimitByRequest(lmt *limiter.Limiter, w http.ResponseWriter, r *http.Request) *errors.HTTPError {
	_ = "STUB: not implemented"
	return nil
}

// Get the lowest value over all keys to return in headers.
// Start with high arbitrary number so that any limit returned would be lower and would
// overwrite the value we start with.

// Loop sliceKeys and check if one of them has error.

// LimitHandler is a middleware that performs rate-limiting given http.Handler struct.
func LimitHandler(lmt *limiter.Limiter, next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// There's no rate-limit error, serve the next handler.

// LimitFuncHandler is a middleware that performs rate-limiting given request handler function.
func LimitFuncHandler(lmt *limiter.Limiter, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// HTTPMiddleware wraps http.Handler with tollbooth limiter
func HTTPMiddleware(lmt *limiter.Limiter) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	// set IP lookup only if not set
	return nil
}

//nolint:gosec // not much we can do here with failed write
