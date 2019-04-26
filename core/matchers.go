package core

import (
	"fmt"
	"strings"

	"github.com/sammiq/matchers"
)

func describeMatchers(allMatchers []matchers.Matcher, separator string) string {
	expected := make([]string, len(allMatchers))
	for i, matcher := range allMatchers {
		expected[i] = matcher.Description()
	}
	return strings.Join(expected, separator)
}

func AllOf(allMatchers ...matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(allMatchers, "and")),
		func(actual interface{}) error { return allMatch(allMatchers, actual) },
	)
}

func AnyOf(anyMatchers ...matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(anyMatchers, "or")),
		func(actual interface{}) error { return anyMatch(anyMatchers, actual) },
	)
}

func EqualTo(expected interface{}) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("value equal to <%v>", expected),
		func(actual interface{}) error { return isEqual(expected, actual) },
	)
}
