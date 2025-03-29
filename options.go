package goretry

import (
	"context"
	"errors"
	"time"
)

type Option func(*retryOptions)

type retryOptions struct {
	maxRetries      int
	backoffStrategy func(int) time.Duration
	shouldRetry     func(error) bool
}

func defaultOptions() retryOptions {
	return retryOptions{
		maxRetries:      3,
		backoffStrategy: ConstantBackoff(500 * time.Millisecond),
		shouldRetry: func(err error) bool {
			return !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded)
		},
	}
}

func WithMaxRetries(maxRetries int) Option {
	return func(o *retryOptions) {
		o.maxRetries = maxRetries
	}
}

func WithBackoffStrategy(backoffStrategy func(int) time.Duration) Option {
	return func(o *retryOptions) {
		o.backoffStrategy = backoffStrategy
	}
}

func WithShouldRetry(shouldRetry func(error) bool) Option {
	return func(o *retryOptions) {
		o.shouldRetry = shouldRetry
	}
}
