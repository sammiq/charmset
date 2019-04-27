package core

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

type both struct {
	matcher matchers.Matcher
}

func (x both) And(matcher matchers.Matcher) matchers.Matcher {
	return matchers.NewMatcher(
		fmt.Sprintf("(%s and %s)", x.matcher.Description(), matcher.Description()),
		func(actual interface{}) error {
			err1 := x.matcher.Match(actual)
			if err1 != nil {
				return err1
			}
			return x.matcher.Match(actual)
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
			err1 := x.matcher.Match(actual)
			if err1 == nil {
				return nil
			}
			err2 := matcher.Match(actual)
			if err2 == nil {
				return nil
			}
			return formatMismatches(
				mismatch(x.matcher, err1),
				mismatch(matcher, err2),
			)
		},
	)
}

func Either(matcher matchers.Matcher) *either {
	return &either{matcher}
}
