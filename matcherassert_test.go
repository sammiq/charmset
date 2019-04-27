package matchers

import (
	"testing"

	"github.com/sammiq/matchers/internal"
)

func TestMatcherAssert_AssertThat(t *testing.T) {
	tests := []struct {
		fail  bool
		fatal bool
	}{
		{
			false, false,
		}, {
			true, false,
		}, {
			true, true,
		}, {
			false, true,
		},
	}

	for _, tt := range tests {
		tester := &internal.MockTester{}
		matcher := &internal.MockMatcher{Fail: tt.fail}
		asserter := &MatcherAssert{Test: tester, FailFast: tt.fatal}

		asserter.AssertThat(42, matcher)

		if tt.fail {
			if tt.fatal {
				if tester.ErrorCount != 0 || tester.FatalCount != 1 {
					t.Errorf("Expected fatal error to trigger when fail with an error")
				}
			} else {
				if tester.ErrorCount != 1 || tester.FatalCount != 0 {
					t.Errorf("Expected normal error to trigger when fail with an error")
				}
			}
			if tester.LastMessage != "Expected: should pass\n     but: did fail" {
				t.Errorf("Expected error message to contain correct information")
			}
		} else {
			if tester.ErrorCount != 0 || tester.FatalCount != 0 {
				t.Errorf("Expected no error to trigger when not failing")
			}
		}
	}

}
