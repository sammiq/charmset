package matches

import (
	"fmt"
	"strings"

	"github.com/sammiq/matchers"
)

func mismatch(matcher matchers.Matcher, err error) string {
	return fmt.Sprintf("expected %s, %s", matcher.Description(), err)
}

func formatMismatches(mismatches ...string) error {
	return fmt.Errorf("(%s)", strings.Join(mismatches, "\n     and "))
}

func and(first matchers.Matcher, second matchers.Matcher, actual interface{}) (err error) {
	err = first.Match(actual)
	if err != nil {
		return err
	}
	return second.Match(actual)
}

func or(first matchers.Matcher, second matchers.Matcher, actual interface{}) (err error) {
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

type both struct {
	matcher matchers.Matcher
}

func (x both) And(matcher matchers.Matcher) matchers.Matcher {
	return matchers.NewMatcher(
		fmt.Sprintf("(%s and %s)", x.matcher.Description(), matcher.Description()),
		func(actual interface{}) error {
			return and(x.matcher, matcher, actual)
		},
	)
}

func Both(matcher matchers.Matcher) *both {
	return &both{matcher}
}

type either struct {
	matcher matchers.Matcher
}

func (x either) Or(matcher matchers.Matcher) matchers.Matcher {
	return matchers.NewMatcher(
		fmt.Sprintf("(%s or %s)", x.matcher.Description(), matcher.Description()),
		func(actual interface{}) error {
			return or(x.matcher, matcher, actual)
		},
	)
}

func Either(matcher matchers.Matcher) *either {
	return &either{matcher}
}
