package matchers

import (
	"testing"

	"github.com/sammiq/matchers/internal"
)

func Test_and(t *testing.T) {
	tests := []struct {
		name        string
		firstMatch  bool
		secondMatch bool
		shouldMatch bool
	}{
		{"should not match if both do not match", false, false, false},
		{"should not match if first does not match", true, false, false},
		{"not match if second does not match", false, true, false},
		{"should match if both match", true, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first := &internal.MockMatcher{Matches: tt.firstMatch}
			second := &internal.MockMatcher{Matches: tt.secondMatch}

			err := and(first, second, 42)
			if tt.shouldMatch {
				if err != nil {
					t.Errorf("expected no error from %v and %v", tt.firstMatch, tt.secondMatch)
				}
			} else {
				if err == nil {
					t.Errorf("expected error from %v and %v", tt.firstMatch, tt.secondMatch)
				}
			}
		})
	}
}

func Test_or(t *testing.T) {
	tests := []struct {
		name        string
		firstMatch  bool
		secondMatch bool
		shouldMatch bool
	}{
		{"should not match if both do not match", false, false, false},
		{"should not match if first does not match", true, false, true},
		{"not match if second does not match", false, true, true},
		{"should match if both match", true, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first := &internal.MockMatcher{Matches: tt.firstMatch}
			second := &internal.MockMatcher{Matches: tt.secondMatch}

			err := or(first, second, 42)
			if tt.shouldMatch {
				if err != nil {
					t.Errorf("expected no error from %v or %v", tt.firstMatch, tt.secondMatch)
				}
			} else {
				if err == nil {
					t.Errorf("expected error from %v or %v", tt.firstMatch, tt.secondMatch)
				}
			}
		})
	}
}
