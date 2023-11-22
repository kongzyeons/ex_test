package models

import "errors"

var (
	ErrHttpRequest        = errors.New("error making HTTP request")
	ErrStatusCode         = errors.New("error: Status code")
	ErrReadResult         = errors.New("error reading response result")
	ErrMaxRetriesExceeded = errors.New("error  HTTP request")
)
