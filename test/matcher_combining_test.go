package test

import (
	"testing"

	"github.com/sammiq/matchers/core/is"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/core/matches"
	"github.com/sammiq/matchers/internal"
)

func Test_BothAndMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(42, matches.Both(is.NotNil()).And(is.EqualTo(43)))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: (not value equal to <<nil>> and value equal to <43>)\n     but: value was <42>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}

func Test_EitherOrMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &matchers.MatcherAssert{Test: tester, FailFast: false}

	assert.That(42, matches.Either(is.Nil()).Or(is.EqualTo(43)))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: (value equal to <<nil>> or value equal to <43>)\n     but: (expected value equal to <<nil>>, value was <42> and\n          expected value equal to <43>, value was <42>)" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}
