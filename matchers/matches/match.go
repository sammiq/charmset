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

func AllOf(allMatchers ...charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(allMatchers, " and ")),
		func(actual interface{}) error { return allMatch(allMatchers, actual) },
	)
}

func AnyOf(anyMatchers ...charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describeMatchers(anyMatchers, " or ")),
		func(actual interface{}) error { return anyMatch(anyMatchers, actual) },
	)
}
