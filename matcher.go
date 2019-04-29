package charmset

import (
	"errors"
	"fmt"
	"strings"
)

// MatcherFunc is a function that takes a value and returns an error
type MatcherFunc func(actual interface{}) error

// Matcher abstracts away the MatcherType struct for use in input
type Matcher interface {
	Description() string
	Match(actual interface{}) error
}

// MatcherType struct is the concrete opaque type used when building matcher
type MatcherType struct {
	expected  string
	matchFunc MatcherFunc
}

// Description returns a description of the expected outcome in present tense
func (mt *MatcherType) Description() string {
	return mt.expected
}

// Match calls with the actual data to match and returns either nil for a match,
// or an error containing a description of why it did not pass in the past tense
func (mt *MatcherType) Match(actual interface{}) error {
	return mt.matchFunc(actual)
}

func mismatch(matcher Matcher, err error) string {
	return fmt.Sprintf("expected %s, %s", matcher.Description(), err)
}

func formatMismatches(mismatches ...string) error {
	return errors.New(strings.Join(mismatches, " and\n          "))
}

func and(first Matcher, second Matcher, actual interface{}) (err error) {
	err = first.Match(actual)
	if err != nil {
		return err
	}
	return second.Match(actual)
}

func or(first Matcher, second Matcher, actual interface{}) (err error) {
	err1 := first.Match(actual)
	if err1 == nil {
		return nil
	}
	err2 := second.Match(actual)
	if err2 == nil {
		return nil
	}
	return formatMismatches(
		mismatch(first, err1),
		mismatch(second, err2),
	)
}

// And composes this matcher and another matcher into a new matcher that
// performs a logical AND on both the matchers
func (mt *MatcherType) And(matcher Matcher) *MatcherType {
	return &MatcherType{
		expected: fmt.Sprintf("%s and %s", mt.expected, matcher.Description()),
		matchFunc: func(actual interface{}) error {
			return and(mt, matcher, actual)
		},
	}
}

// Or composes this matcher and another matcher into a new matcher that
// performs a logical OR on both the matchers
func (mt *MatcherType) Or(matcher Matcher) *MatcherType {
	return &MatcherType{
		expected: fmt.Sprintf("%s or %s", mt.expected, matcher.Description()),
		matchFunc: func(actual interface{}) error {
			return or(mt, matcher, actual)
		},
	}
}

// NewMatcher creates a matcher given a description and a match function
func NewMatcher(description string, match MatcherFunc) *MatcherType {
	return &MatcherType{description, match}
}
