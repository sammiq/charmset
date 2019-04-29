package is

import (
	"testing"

	"github.com/sammiq/charmset/internal"
)

func Test_not(t *testing.T) {
	tests := []struct {
		name    string
		matches bool
	}{
		{"should not return error when no match", false},
		{"should return error when match", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Matches: tt.matches}

			err := not(matcher, 42)
			if tt.matches {
				if err == nil {
					t.Error("Expected err not nil when matcher does match")
				}
			} else {
				if err != nil {
					t.Error("Expected err nil when matcher does not match")
				}
			}
		})
	}
}
