package matches

import (
	"fmt"
	"strings"

	"github.com/sammiq/charmset"
)

func allMatch(charmset []charmset.Matcher, actual interface{}) (err error) {
	//report any error
	for _, matcher := range charmset {
		err = matcher.Match(actual)
		if err != nil {
			break
		}
	}
	return err
}

func anyMatch(charmset []charmset.Matcher, actual interface{}) (err error) {
	for _, matcher := range charmset {
		err = matcher.Match(actual)
		if err == nil {
			break
		}
	}
	return err
}

func describecharmset(allcharmset []charmset.Matcher, separator string) string {
	expected := make([]string, len(allcharmset))
	for i, matcher := range allcharmset {
		expected[i] = matcher.Description()
	}
	return strings.Join(expected, separator)
}

func AllOf(allcharmset ...charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describecharmset(allcharmset, " and ")),
		func(actual interface{}) error { return allMatch(allcharmset, actual) },
	)
}

func AnyOf(anycharmset ...charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("(%s)", describecharmset(anycharmset, " or ")),
		func(actual interface{}) error { return anyMatch(anycharmset, actual) },
	)
}
