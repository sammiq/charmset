package core

import (
	"testing"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/internal"
)

func Test_EqualToMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(42, EqualTo(43))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: value equal to <43>\n     but: value was <42>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}

func Test_NotEqualToMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(42, NotEqualTo(42))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: not value equal to <42>\n     but: value was <42>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}

func Test_NilMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(42, Nil())
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: value equal to <<nil>>\n     but: value was <42>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}

func Test_NotNilMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(nil, NotNil())
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: not value equal to <<nil>>\n     but: value was <<nil>>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}
