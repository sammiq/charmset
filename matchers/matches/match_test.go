package matches

import (
	"testing"

	"github.com/sammiq/charmset"
	"github.com/sammiq/charmset/internal"
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
		{"should match if multiple charmset match", []bool{true, true, true}, true},
		{"should not match if first of multiple charmset match", []bool{false, true, true}, false},
		{"should not match if last of multiple charmset match", []bool{true, true, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charmset := make([]charmset.Matcher, len(tt.matches))
			for i, mm := range tt.matches {
				charmset[i] = &internal.MockMatcher{Matches: mm}
			}
			err := allMatch(charmset, 42)
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
		{"should match if multiple charmset match", []bool{true, true, true}, true},
		{"should match if first of multiple charmset match", []bool{false, true, true}, true},
		{"should match if last of multiple charmset match", []bool{true, true, false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charmset := make([]charmset.Matcher, len(tt.matches))
			for i, mm := range tt.matches {
				charmset[i] = &internal.MockMatcher{Matches: mm}
			}
			err := anyMatch(charmset, 42)
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
