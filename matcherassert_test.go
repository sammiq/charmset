package charmset

import (
	"testing"

	"github.com/sammiq/charmset/internal"
)

func TestMatcherAssert_That(t *testing.T) {
	tests := []struct {
		match bool
		fatal bool
	}{
		{
			true, false,
		}, {
			false, false,
		}, {
			false, true,
		}, {
			true, true,
		},
	}

	for _, tt := range tests {
		tester := &internal.MockTester{}
		matcher := &internal.MockMatcher{Matches: tt.match}
		assert := &MatcherAssert{Test: tester, FailFast: tt.fatal}

		assert.That(42, matcher)

		if tt.match {
			if tester.ErrorCount != 0 || tester.FatalCount != 0 {
				t.Errorf("Expected no error to trigger when not failing")
			}
		} else {
			if tt.fatal {
				if tester.ErrorCount != 0 || tester.FatalCount != 1 {
					t.Errorf("Expected fatal error to trigger when fail with an error")
				}
			} else {
				if tester.ErrorCount != 1 || tester.FatalCount != 0 {
					t.Errorf("Expected normal error to trigger when fail with an error")
				}
			}
			if tester.LastMessage != "Expected: might match\n     but: did not match" {
				t.Errorf("Expected error message to contain correct information")
			}
		}
	}

}
