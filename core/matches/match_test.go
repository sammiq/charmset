package matches

import (
	"testing"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/internal"
)

func Test_allMatch(t *testing.T) {
	tests := []struct {
		name        string
		matches     []bool
		shouldMatch bool
	}{
		{"should match if empty", []bool{}, true},
		{"should match if single matcher matches", []bool{true}, true},
		{"should not match if single matcher does not match", []bool{false}, false},
		{"should match if multiple matchers match", []bool{true, true, true}, true},
		{"should not match if first of multiple matchers match", []bool{false, true, true}, false},
		{"should not match if last of multiple matchers match", []bool{true, true, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchers := make([]matchers.Matcher, len(tt.matches))
			for i, mm := range tt.matches {
				matchers[i] = &internal.MockMatcher{Matches: mm}
			}
			err := allMatch(matchers, 42)
			if tt.shouldMatch {
				if err != nil {
					t.Errorf("always expect no error if should match. error was: %s", err)
				}
			} else {
				if err == nil {
					t.Errorf("always expect an error if should not match")
				}
			}
		})
	}
}

func Test_anyMatch(t *testing.T) {
	tests := []struct {
		name        string
		matches     []bool
		shouldMatch bool
	}{
		{"should match if empty", []bool{}, true},
		{"should match if single matcher matches", []bool{true}, true},
		{"should not match if single matcher does not match", []bool{false}, false},
		{"should match if multiple matchers match", []bool{true, true, true}, true},
		{"should match if first of multiple matchers match", []bool{false, true, true}, true},
		{"should match if last of multiple matchers match", []bool{true, true, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matchers := make([]matchers.Matcher, len(tt.matches))
			for i, mm := range tt.matches {
				matchers[i] = &internal.MockMatcher{Matches: mm}
			}
			err := anyMatch(matchers, 42)
			if tt.shouldMatch {
				if err != nil {
					t.Errorf("always expect no error if should match. error was: %s", err)
				}
			} else {
				if err == nil {
					t.Errorf("always expect an error if should not match")
				}
			}
		})
	}
}
