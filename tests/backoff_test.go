package tests

import (
	"github.com/masilvasql/goretry"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	base := 500 * time.Millisecond
	factor := 2.0
	backoff := goretry.ExponentialBackoff(base, factor)

	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{1, 1000 * time.Millisecond}, // 500ms * 2 * 1
		{2, 2000 * time.Millisecond}, // 500ms * 2 * 2
		{3, 3000 * time.Millisecond}, // 500ms * 2 * 3
		{4, 4000 * time.Millisecond}, // 500ms * 2 * 4
		{5, 5000 * time.Millisecond}, // 500ms * 2 * 5
	}

	for _, tt := range tests {
		t.Run("Attempt "+string(rune(tt.attempt)), func(t *testing.T) {
			assert.Equal(t, tt.expected, backoff(tt.attempt))
		})
	}
}

func TestExponentialBackoff_WithZeroBase(t *testing.T) {
	backoff := goretry.ExponentialBackoff(0, 2.0)

	assert.Equal(t, time.Duration(0), backoff(1))
	assert.Equal(t, time.Duration(0), backoff(5))
}

func TestExponentialBackoff_WithOneFactor(t *testing.T) {
	base := 300 * time.Millisecond
	backoff := goretry.ExponentialBackoff(base, 1.0)

	assert.Equal(t, 300*time.Millisecond, backoff(1))
	assert.Equal(t, 600*time.Millisecond, backoff(2))
	assert.Equal(t, 900*time.Millisecond, backoff(3))
}

func TestLinearBackoff(t *testing.T) {
	base := 500 * time.Millisecond
	backoff := goretry.LinearBackoff(base)

	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{1, 500 * time.Millisecond},  // 500ms * 1
		{2, 1000 * time.Millisecond}, // 500ms * 2
		{3, 1500 * time.Millisecond}, // 500ms * 3
		{4, 2000 * time.Millisecond}, // 500ms * 4
		{5, 2500 * time.Millisecond}, // 500ms * 5
	}

	for _, tt := range tests {
		t.Run("Attempt "+string(rune(tt.attempt)), func(t *testing.T) {
			assert.Equal(t, tt.expected, backoff(tt.attempt))
		})
	}
}

func TestLinearBackoff_WithZeroBase(t *testing.T) {
	backoff := goretry.LinearBackoff(0)

	assert.Equal(t, time.Duration(0), backoff(1))
	assert.Equal(t, time.Duration(0), backoff(5))
}

func TestLinearBackoff_WithNegativeAttempts(t *testing.T) {
	backoff := goretry.LinearBackoff(500 * time.Millisecond)

	assert.Equal(t, time.Duration(0), backoff(0))
	assert.Equal(t, time.Duration(0), backoff(-1))
}
