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
		err error
	}{
		{"+1s", time.Second, nil},
		{"aboba",time.Second,nil},
	}

	for _, tt := range tests {
		time, err := ParseDuration(tt.in)
		if time != tt.out {
			t.Fatalf("TimeError: expected %s,got %s\n", tt.out.String(), time.String())
		} else if err != tt.err {
			t.Fatalf("ErrorProblem: expected %v,got %v\n", tt.err.Error(), err)
		}
	}
}
