package tabletest

import (
	"testing"
	"time"
)

// min coverage: . 95%

func TestParseDuration(t *testing.T) {
	var tests = []struct {
		in  string
		out time.Duration
	}{
		{"+1s", time.Second},
		{"-1.2s", -1200 * time.Millisecond},
		{"0", 0},
		{"a", 0},
		{"", 0},
		{".", 0},
		{"10mp", 0},
		{"0.000000000000000000000000329999999999999999999999999999999999999999999999999", 0},
		{"999999999999999999999999999.2", 0},
		{".123ms", 123 * time.Microsecond},
		{"-1.2s.", 0},
		{"1234567899999h", 0},
		{"9223372036854775808", 0},
		{"0.9223372036854775808", 0},
	}

	for _, tt := range tests {
		time, _ := ParseDuration(tt.in)
		if time != tt.out {
			t.Fatalf("TimeError: expected \"%s\",got \"%s\"\n", tt.out.String(), time.String())
		}
	}
}
