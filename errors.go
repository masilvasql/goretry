package goretry

import "errors"

var (
	ErrMaxRetriesExceeded = errors.New("max retries exceeded")
)
