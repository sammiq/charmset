package is

import (
	"fmt"

	"github.com/sammiq/charmset"
)

func not(matcher charmset.Matcher, actual interface{}) error {
	err := matcher.Match(actual)
	if err != nil {
		return nil
	} else {
		return fmt.Errorf("value was <%v>", actual)
	}
}

func Not(matcher charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("not %s", matcher.Description()),
		func(actual interface{}) error {
			return not(matcher, actual)
		},
	)
}
