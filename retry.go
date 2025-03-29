package goretry

import (
	"context"
	"log"
	"time"
)

type RetryFunc[T any] func(ctx context.Context) (T, error)

func Do[T any](ctx context.Context, fn RetryFunc[T], opts ...Option) (T, error) {
	config := defaultOptions()
	for _, opt := range opts {
		opt(&config)
	}

	var lastErr error
	var result T

	for attempt := 1; attempt <= config.maxRetries; attempt++ {
		log.Println("Trying attempt", attempt)
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		result, lastErr = fn(ctx)
		if lastErr == nil {
			return result, nil
		}

		if !config.shouldRetry(lastErr) {
			break
		}
		
		if config.backoffStrategy != nil {
			sleepDuration := config.backoffStrategy(attempt)
			select {
			case <-ctx.Done():
				return result, ctx.Err()
			case <-time.After(sleepDuration):
			}
		}
	}

	return result, lastErr
}
