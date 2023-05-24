package clashofclans

import "errors"

var (
	ErrFailedToBuildRequest      = errors.New("failed to build request")
	ErrRequestFailed             = errors.New("request failed")
	ErrFailedToReadResponseBody  = errors.New("failed to read response body")
	ErrFailedToParseResponseBody = errors.New("failed to parse response body")
)
