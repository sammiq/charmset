package is

import (
	"fmt"

	"github.com/sammiq/matchers"
)

func not(matcher matchers.Matcher, actual interface{}) error {
	err := matcher.Match(actual)
	if err != nil {
		return nil
	} else {
		return fmt.Errorf("value was <%v>", actual)
	}
}

func Not(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("not %s", matcher.Description()),
		func(actual interface{}) error {
			return not(matcher, actual)
		},
	)
}
