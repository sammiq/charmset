package matchers

import (
	"errors"
	"fmt"
	"testing"
)

type mockTester struct {
	errorCount  int
	fatalCount  int
	lastMessage string
}

func (x *mockTester) Logf(format string, args ...interface{}) {
	x.lastMessage = fmt.Sprintf(format, args...)
}

func (x *mockTester) Fail() {
	x.errorCount++
}

func (x *mockTester) FailNow() {
	x.fatalCount++
}

type mockMatcher struct {
	fail bool
}

func (x *mockMatcher) Description() string {
	return "expectation"
}

func (x *mockMatcher) Match(actual interface{}) error {
	if x.fail {
		return errors.New("actual")
	}
	return nil
}

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
		tester := &mockTester{}
		matcher := &mockMatcher{tt.fail}
		asserter := &MatcherAssert{tester, tt.fatal}

		asserter.AssertThat("reason", 42, matcher)

		if tt.fail {
			if tt.fatal {
				if tester.errorCount != 0 || tester.fatalCount != 1 {
					t.Errorf("Expected fatal error to trigger when fail with an error")
				}
			} else {
				if tester.errorCount != 1 || tester.fatalCount != 0 {
					t.Errorf("Expected normal error to trigger when fail with an error")
				}
			}
			if tester.lastMessage != "reason\nExpected: expectation\nbut: actual" {
				t.Errorf("Expected error message to contain correct information")
			}
		} else {
			if tester.errorCount != 0 || tester.fatalCount != 0 {
				t.Errorf("Expected no error to trigger when not failing")
			}
		}
	}

}
