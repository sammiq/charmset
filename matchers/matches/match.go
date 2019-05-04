package matches

import (
	"fmt"
	"strings"

	"github.com/sammiq/charmset"
)

func allMatch(matchers []charmset.Matcher, actual interface{}) (err error) {
	//report any error
	for _, matcher := range matchers {
		err = matcher.Match(actual)
		if err != nil {
			break
		}
	}
	return err
}

func anyMatch(matchers []charmset.Matcher, actual interface{}) (err error) {
	for _, matcher := range matchers {
		err = matcher.Match(actual)
		if err == nil {
			break
		}
	}
	return err
}

func describeMatchers(allMatchers []charmset.Matcher, separator string) string {
	expected := make([]string, len(allMatchers))
	for i, matcher := range allMatchers {
		expected[i] = matcher.Description()
	}
	return strings.Join(expected, separator)
}

// AllOf returns a matcher that checks whether a value matches every
// matcher given. Returns early is a matcher does not match.
func AllOf(allMatchers ...charmset.Matcher) *charmset.MatcherType {
	if len(allMatchers) == 0 {
		//panic as there is no reason to continue the test if allMatchers is invalid at construction
		panic("will never match an empty set of matchers")
	}
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(allMatchers, " and ")),
		func(actual interface{}) error { return allMatch(allMatchers, actual) },
	)
}

// AnyOf returns a matcher that checks whether a value matches any
// matcher given. Returns early is a matcher returns a match.
func AnyOf(anyMatchers ...charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(anyMatchers, " or ")),
		func(actual interface{}) error { return anyMatch(anyMatchers, actual) },
	)
}
