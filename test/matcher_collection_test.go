package test

import (
	"testing"

	"github.com/sammiq/charmset"
	"github.com/sammiq/charmset/internal"
	"github.com/sammiq/charmset/matchers/has"
	"github.com/sammiq/charmset/matchers/is"
)

func Test_EveryItemMatchingMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &charmset.MatcherAssert{Test: tester, FailFast: false}

	assert.That([]int{42, 42, 43}, has.EveryItemMatching(is.EqualTo(42)))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: every item to have value equal to <42>\n     but: contained an item where value was <43>" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}

func Test_AnyItemMatchingMatcher(t *testing.T) {
	tester := &internal.MockTester{}
	assert := &charmset.MatcherAssert{Test: tester, FailFast: false}

	assert.That([]int{43, 43, 43}, has.AnyItemMatching(is.EqualTo(42)))
	if tester.ErrorCount != 1 {
		t.Error("expected error to fire")
	}
	if tester.LastMessage != "Expected: any item to have value equal to <42>\n     but: no item matched where [\n          value was <43>,\n          value was <43>,\n          value was <43>\n          ]" {
		t.Errorf("Expected error message to be formatted correctly (was %s)", tester.LastMessage)
	}
}
