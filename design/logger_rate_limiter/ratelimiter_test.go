package logger_rate_limiter_test

import (
	"testing"

	"github.com/andream16/algorithms/design/logger_rate_limiter"
)

func TestLogger_ShouldPrintMessage(t *testing.T) {
	logger := logger_rate_limiter.Constructor()

	for _, tt := range []struct {
		timestamp int
		message   string
		expected  bool
	}{
		{
			timestamp: 1,
			message:   "foo",
			expected:  true,
		},
		{
			timestamp: 2,
			message:   "bar",
			expected:  true,
		},
		{
			timestamp: 3,
			message:   "foo",
			expected:  false,
		},
		{
			timestamp: 8,
			message:   "bar",
			expected:  false,
		},
		{
			timestamp: 10,
			message:   "foo",
			expected:  false,
		},
		{
			timestamp: 11,
			message:   "foo",
			expected:  true,
		},
	} {
		if res := logger.ShouldPrintMessage(tt.timestamp, tt.message); tt.expected != res {
			t.Fatalf("expected %v, got %v instead", tt.expected, res)
		}
	}
}
