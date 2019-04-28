package matchers

import (
	"errors"
	"fmt"
	"strings"
)

//abstracts away the MatcherType struct for use in input
type Matcher interface {
	Description() string
	Match(actual interface{}) error
}

type MatcherType struct {
	expected  string
	matchFunc func(actual interface{}) error
}

func (mt *MatcherType) Description() string {
	return mt.expected
}

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

func (mt *MatcherType) And(matcher Matcher) *MatcherType {
	return &MatcherType{
		expected: fmt.Sprintf("%s and %s", mt.expected, matcher.Description()),
		matchFunc: func(actual interface{}) error {
			return and(mt, matcher, actual)
		},
	}
}

func (mt *MatcherType) Or(matcher Matcher) *MatcherType {
	return &MatcherType{
		expected: fmt.Sprintf("%s or %s", mt.expected, matcher.Description()),
		matchFunc: func(actual interface{}) error {
			return or(mt, matcher, actual)
		},
	}
}

func NewMatcher(description string, match func(actual interface{}) error) *MatcherType {
	return &MatcherType{description, match}
}
