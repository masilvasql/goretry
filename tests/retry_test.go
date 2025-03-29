package tests

import (
	"context"
	"errors"
	"github.com/masilvasql/goretry"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestRetry_SuccesOnFirstTry(t *testing.T) {
	ctx := context.Background()
	retryFunc := func(ctx context.Context) (string, error) {
		return "success", nil
	}

	result, err := goretry.Do(ctx, retryFunc)

	assert.NoError(t, err)
	assert.Equal(t, "success", result)

}

func TestRetry_FailsAfterMaxRetries(t *testing.T) {
	ctx := context.Background()
	retryFunc := func(ctx context.Context) (string, error) {
		return "", errors.New("falha")
	}

	_, err := goretry.Do(ctx, retryFunc, goretry.WithMaxRetries(3))
	assert.Error(t, err)
	assert.Equal(t, "falha", err.Error())
}

func TestRetry_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	retryFunc := func(ctx context.Context) (string, error) {
		time.Sleep(10 * time.Millisecond)
		return "", errors.New("falha")
	}

	_, err := goretry.Do(ctx, retryFunc)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestRetry_WithShouldRetryFalseWhenRequestFailed(t *testing.T) {
	ctx := context.Background()

	_, err := goretry.Do(ctx, getCepRequest, goretry.WithShouldRetry(func(err error) bool {
		if err == nil {
			return false
		}

		if err.Error() == "Request Failed" {
			return false
		}
		return true
	}))

	assert.Error(t, err)
	assert.Equal(t, "Request Failed", err.Error())

}

func getCepRequest(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://viacep.com.br/wss/01001000/json/", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Request Failed")
	}

	return "success", nil
}
